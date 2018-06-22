package provider

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	log "github.com/sirupsen/logrus"
)

const (
	aliAccessKeyId = "ALICLOUD_ACCESS_KEY_ID"
	aliAccessKeySecret = "ALICLOUD_ACCESS_KEY_SECRET"
	region = "cn-beijing"
	recordIdKey = "RecordId"
	defaultAlidnsRecordTTL = 600
)

type AliAPI interface {
	AddDomainRecord(request *alidns.AddDomainRecordRequest) (response *alidns.AddDomainRecordResponse, err error)
	DeleteDomainRecord(request *alidns.DeleteDomainRecordRequest) (response *alidns.DeleteDomainRecordResponse, err error)
	DescribeDomainRecords(request *alidns.DescribeDomainRecordsRequest) (response *alidns.DescribeDomainRecordsResponse, err error)
	DescribeDomains(request *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error)
	UpdateDomainRecord(request *alidns.UpdateDomainRecordRequest) (response *alidns.UpdateDomainRecordResponse, err error)
}

type AliProvider struct {
	client AliAPI
	dryRun bool
	// only consider hosted domains managing domains ending in this suffix
	domainFilter DomainFilter
}

// NewAliProvider initializes a new Alidns based Provider.
func NewAliProvider(domainFilter DomainFilter, dryRun bool) (*AliProvider, error) {
	accessKeyId := os.Getenv(aliAccessKeyId)
	if len(accessKeyId) == 0 {
		return nil, fmt.Errorf("%s is not set", aliAccessKeyId)
	}

	accessKeySecret := os.Getenv(aliAccessKeySecret)
	if len(accessKeySecret) == 0 {
		return nil, fmt.Errorf("%s is not set", aliAccessKeySecret)
	}

	client, err := alidns.NewClientWithAccessKey(region, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	provider := &AliProvider{
		client:         client,
		domainFilter:   domainFilter,
		dryRun:         dryRun,
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
		request.SetDomain(domain.DomainName)

		response, err := p.client.DescribeDomainRecords(request)
		if err != nil {
			return nil, err
		}

		for _, record := range response.DomainRecords.Record {
			ep := endpoint.NewEndpointWithTTL(record.RR, record.Type, endpoint.TTL(record.TTL), record.Value)
			recordId, err := endpoint.NewLabelsFromString(recordIdKey + "=" + record.RecordId)
			if err != nil {
				return nil, err
			}
			ep.Labels = recordId

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
			log.Infof("Ignoring changes to '%s' because a suitable Azure DNS domain was not found.", change.DNSName)
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

func (p *AliProvider) findRecord(record endpoint.Endpoint) (recordIds []string) {
	records, err := p.Records()
	if err != nil {
		log.Error(err)
		return nil
	}

	isRecordIncluded := func(e1 endpoint.Endpoint, e2 endpoint.Endpoint) bool {
		if e1.DNSName != e2.DNSName {
			return false
		}
		if e1.RecordType != e2.RecordType {
			return false
		}
		for _, t := range e2.Targets {
			if t == e1.Targets[0] {
				return true
			}
		}
		return false
	}

	for _, r := range records {
		if isRecordIncluded(*r, record) {
			recordIds = append(recordIds, r.Labels[recordIdKey])
		}
	}
	return recordIds
}

func (p *AliProvider) deleteRecords(deleted aliChangeMap) {
	for domain, endpoints := range deleted {
		for _, endpoint := range endpoints {
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' value '%s' for Ali DNS domain '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Targets,
					domain,
				)
			} else {
				log.Infof("Deleting %s record named '%s' value '%s' for Ali DNS domain '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Targets,
					domain,
				)
				recordIds := p.findRecord(*endpoint)
				if recordIds == nil {
					log.Errorf("Failed to find %s record named '%s' value '%s' for Ali DNS domain '%s'.",
						endpoint.RecordType,
						endpoint.DNSName,
						endpoint.Targets,
						domain,
					)
					continue
				}

				for _, recordId := range recordIds {
					request := alidns.CreateDeleteDomainRecordRequest()
					request.RegionId = recordId

					_, err := p.client.DeleteDomainRecord(request)
					if err != nil {
						log.Errorf("Failed to delete %s record named '%s' value '%s' recordId '%s' for Ali DNS domain '%s': %v",
							endpoint.RecordType,
							endpoint.DNSName,
							endpoint.Targets,
							recordId,
							domain,
							err,
						)
					}
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
					endpoint.Targets,
					domain,
				)
				continue
			}

			log.Infof("Creating %s record named '%s' value '%s' for Ali DNS domain '%s'.",
				endpoint.RecordType,
				endpoint.DNSName,
				endpoint.Targets,
				domain,
			)

			request := alidns.CreateAddDomainRecordRequest()
			request.SetDomain(domain)
			request.RR = endpoint.DNSName
			request.Type = endpoint.RecordType
			if endpoint.RecordTTL.IsConfigured() {
				request.TTL = requests.NewInteger(int(endpoint.RecordTTL))
			} else {
				request.TTL = requests.NewInteger(defaultAlidnsRecordTTL)
			}

			for _, target := range endpoint.Targets {
				request.Value = target
				_, err := p.client.AddDomainRecord(request)
				if err != nil {
					log.Errorf("Failed to create %s record named '%s' value '%s' for Ali DNS domain '%s': %v",
						endpoint.RecordType,
						endpoint.DNSName,
						target,
						domain,
						err,
					)
				}
			}
		}
	}
}
