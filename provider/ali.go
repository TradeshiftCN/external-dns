package provider

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	aliAccessKeyID         = "ALICLOUD_ACCESS_KEY_ID"
	aliAccessKeySecret     = "ALICLOUD_ACCESS_KEY_SECRET"
	region                 = "cn-beijing"
	recordIDKey            = "RecordId"
	defaultAlidnsRecordTTL = 600
)

// AliAPI is alibaba cloud dns service
type AliAPI interface {
	AddDomainRecord(request *alidns.AddDomainRecordRequest) (response *alidns.AddDomainRecordResponse, err error)
	DeleteDomainRecord(request *alidns.DeleteDomainRecordRequest) (response *alidns.DeleteDomainRecordResponse, err error)
	DescribeDomainRecords(request *alidns.DescribeDomainRecordsRequest) (response *alidns.DescribeDomainRecordsResponse, err error)
	DescribeDomains(request *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error)
	UpdateDomainRecord(request *alidns.UpdateDomainRecordRequest) (response *alidns.UpdateDomainRecordResponse, err error)
}

// AliProvider is an implementation of Provider for Ali DNS
type AliProvider struct {
	client AliAPI
	dryRun bool
	// only consider hosted domains managing domains ending in this suffix
	domainFilter DomainFilter
}

// NewAliProvider initializes a new Alidns based Provider.
func NewAliProvider(domainFilter DomainFilter, dryRun bool) (*AliProvider, error) {
	accessKeyID := os.Getenv(aliAccessKeyID)
	if len(accessKeyID) == 0 {
		return nil, fmt.Errorf("%s is not set", aliAccessKeyID)
	}

	accessKeySecret := os.Getenv(aliAccessKeySecret)
	if len(accessKeySecret) == 0 {
		return nil, fmt.Errorf("%s is not set", aliAccessKeySecret)
	}

	client, err := alidns.NewClientWithAccessKey(region, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}

	provider := &AliProvider{
		client:       client,
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}

	return provider, nil
}

// Domains returns the list of hosted domain names.
func (p *AliProvider) Domains() ([]alidns.Domain, error) {
	response, err := p.client.DescribeDomains(alidns.CreateDescribeDomainsRequest())
	if err != nil {
		return nil, err
	}

	var domains []alidns.Domain
	for _, domain := range response.Domains.Domain {
		if !p.domainFilter.Match(domain.DomainName) {
			continue
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

// Records returns the list of records in given hosted domains.
func (p *AliProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	domains, err := p.Domains()
	if err != nil {
		return nil, err
	}

	for _, domain := range domains {
		request := alidns.CreateDescribeDomainRecordsRequest()
		request.DomainName = domain.DomainName

		response, err := p.client.DescribeDomainRecords(request)
		if err != nil {
			return nil, err
		}

		for _, record := range response.DomainRecords.Record {
			target := record.Value
			// Add double quotation marks for TXT record target as the regex needs it to match the owner id
			if record.Type == endpoint.RecordTypeTXT {
				target = fmt.Sprintf("\"%s\"", record.Value)
			}

			ep := endpoint.NewEndpointWithTTL(record.RR + "." + domain.DomainName, target, record.Type, endpoint.TTL(record.TTL))
			recordID := map[string]string{
				recordIDKey: record.RecordId,
			}
			ep.MergeLabels(recordID)

			endpoints = append(endpoints, ep)
		}
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given domain.
func (p *AliProvider) ApplyChanges(changes *plan.Changes) error {
	domains, err := p.Domains()
	if err != nil {
		return err
	}

	deleted, created := p.mapChanges(domains, changes)
	p.deleteRecords(deleted)
	p.createRecords(created)

	return nil
}

type aliChangeMap map[string][]*endpoint.Endpoint

func (p *AliProvider) mapChanges(domains []alidns.Domain, changes *plan.Changes) (aliChangeMap, aliChangeMap) {
	deleted := aliChangeMap{}
	created := aliChangeMap{}

	zoneNameIDMapper := zoneIDName{}
	for _, d := range domains {
		if d.DomainName != "" {
			zoneNameIDMapper.Add(d.DomainName, d.DomainName)
		}
	}

	mapChange := func(changeMap aliChangeMap, change *endpoint.Endpoint) {
		domain, _ := zoneNameIDMapper.FindZone(change.DNSName)
		if domain == "" {
			log.Infof("Ignoring changes to '%s' because a suitable Ali DNS domain was not found.", change.DNSName)
			return
		}

		changeMap[domain] = append(changeMap[domain], change)
	}

	for _, change := range changes.Delete {
		mapChange(deleted, change)
	}
	for _, change := range changes.UpdateOld {
		mapChange(deleted, change)
	}
	for _, change := range changes.UpdateNew {
		mapChange(created, change)
	}
	for _, change := range changes.Create {
		mapChange(created, change)
	}

	return deleted, created
}

func (p *AliProvider) findRecord(record endpoint.Endpoint) (recordID string) {
	records, err := p.Records()
	if err != nil {
		log.Error(err)
		return ""
	}

	recordEquals := func(e1 endpoint.Endpoint, e2 endpoint.Endpoint) bool {
		if e1.DNSName != e2.DNSName {
			return false
		}
		if e1.Target != e2.Target {
			return false
		}
		if e1.RecordType != e2.RecordType {
			return false
		}
		return true
	}

	for _, r := range records {
		if recordEquals(*r, record) {
			recordID = r.Labels[recordIDKey]
			return recordID
		}
	}

	return ""
}

func (p *AliProvider) deleteRecords(deleted aliChangeMap) {
	for domain, endpoints := range deleted {
		for _, endpoint := range endpoints {
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' value '%s' for Ali DNS domain '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					domain,
				)
			} else {
				log.Infof("Deleting %s record named '%s' value '%s' for Ali DNS domain '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					domain,
				)
				recordID := p.findRecord(*endpoint)
				if recordID == "" {
					log.Errorf("Failed to find %s record named '%s' for Ali DNS domain '%s'.",
						endpoint.RecordType,
						endpoint.DNSName,
						endpoint.Target,
						domain,
					)
					continue
				}

				request := alidns.CreateDeleteDomainRecordRequest()
				request.RecordId = recordID

				_, err := p.client.DeleteDomainRecord(request)
				if err != nil {
					log.Errorf("Failed to delete %s record named '%s' value '%s' recordID '%s' for Ali DNS domain '%s': %v",
						endpoint.RecordType,
						endpoint.DNSName,
						endpoint.Target,
						recordID,
						domain,
						err,
					)
				}
			}
		}
	}
}

func (p *AliProvider) createRecords(created aliChangeMap) {
	for domain, endpoints := range created {
		for _, endpoint := range endpoints {
			if p.dryRun {
				log.Infof("Would create %s record named '%s' value '%s' for Ali DNS domain '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					domain,
				)
				continue
			}

			log.Infof("Creating %s record named '%s' value '%s' for Ali DNS domain '%s'.",
				endpoint.RecordType,
				endpoint.DNSName,
				endpoint.Target,
				domain,
			)

			request := alidns.CreateAddDomainRecordRequest()
			request.DomainName = domain
			request.RR = strings.TrimSuffix(endpoint.DNSName, "." + domain)
			request.Type = endpoint.RecordType
			request.Value = endpoint.Target
			if endpoint.RecordTTL.IsConfigured() {
				request.TTL = requests.NewInteger(int(endpoint.RecordTTL))
			} else {
				request.TTL = requests.NewInteger(defaultAlidnsRecordTTL)
			}

			_, err := p.client.AddDomainRecord(request)
			if err != nil {
				log.Errorf("Failed to create %s record named '%s' value '%s' for Ali DNS domain '%s': %v",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					domain,
					err,
				)
			}
		}
	}
}
