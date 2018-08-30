package scdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScdnService invokes the scdn.DescribeScdnService API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdnservice.html
func (client *Client) DescribeScdnService(request *DescribeScdnServiceRequest) (response *DescribeScdnServiceResponse, err error) {
	response = CreateDescribeScdnServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnServiceWithChan invokes the scdn.DescribeScdnService API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnServiceWithChan(request *DescribeScdnServiceRequest) (<-chan *DescribeScdnServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnService(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeScdnServiceWithCallback invokes the scdn.DescribeScdnService API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnServiceWithCallback(request *DescribeScdnServiceRequest, callback func(response *DescribeScdnServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnService(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeScdnServiceRequest is the request struct for api DescribeScdnService
type DescribeScdnServiceRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnServiceResponse is the response struct for api DescribeScdnService
type DescribeScdnServiceResponse struct {
	*responses.BaseResponse
	RequestId              string         `json:"RequestId" xml:"RequestId"`
	InstanceId             string         `json:"InstanceId" xml:"InstanceId"`
	OpenTime               string         `json:"OpenTime" xml:"OpenTime"`
	EndTime                string         `json:"EndTime" xml:"EndTime"`
	ProtectType            string         `json:"ProtectType" xml:"ProtectType"`
	ProtectTypeValue       string         `json:"ProtectTypeValue" xml:"ProtectTypeValue"`
	Bandwidth              string         `json:"Bandwidth" xml:"Bandwidth"`
	CcProtection           string         `json:"CcProtection" xml:"CcProtection"`
	DDoSBasic              string         `json:"DDoSBasic" xml:"DDoSBasic"`
	DomainCount            string         `json:"DomainCount" xml:"DomainCount"`
	ElasticProtection      string         `json:"ElasticProtection" xml:"ElasticProtection"`
	BandwidthValue         string         `json:"BandwidthValue" xml:"BandwidthValue"`
	CcProtectionValue      string         `json:"CcProtectionValue" xml:"CcProtectionValue"`
	DDoSBasicValue         string         `json:"DDoSBasicValue" xml:"DDoSBasicValue"`
	DomainCountValue       string         `json:"DomainCountValue" xml:"DomainCountValue"`
	ElasticProtectionValue string         `json:"ElasticProtectionValue" xml:"ElasticProtectionValue"`
	PriceType              string         `json:"PriceType" xml:"PriceType"`
	PricingCycle           string         `json:"PricingCycle" xml:"PricingCycle"`
	OperationLocks         OperationLocks `json:"OperationLocks" xml:"OperationLocks"`
}

// CreateDescribeScdnServiceRequest creates a request to invoke DescribeScdnService API
func CreateDescribeScdnServiceRequest() (request *DescribeScdnServiceRequest) {
	request = &DescribeScdnServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnService", "", "")
	return
}

// CreateDescribeScdnServiceResponse creates a response to parse from DescribeScdnService response
func CreateDescribeScdnServiceResponse() (response *DescribeScdnServiceResponse) {
	response = &DescribeScdnServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
