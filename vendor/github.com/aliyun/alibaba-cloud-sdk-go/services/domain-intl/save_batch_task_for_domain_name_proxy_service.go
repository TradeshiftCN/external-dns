package domain_intl

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

// SaveBatchTaskForDomainNameProxyService invokes the domain_intl.SaveBatchTaskForDomainNameProxyService API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskfordomainnameproxyservice.html
func (client *Client) SaveBatchTaskForDomainNameProxyService(request *SaveBatchTaskForDomainNameProxyServiceRequest) (response *SaveBatchTaskForDomainNameProxyServiceResponse, err error) {
	response = CreateSaveBatchTaskForDomainNameProxyServiceResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForDomainNameProxyServiceWithChan invokes the domain_intl.SaveBatchTaskForDomainNameProxyService API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskfordomainnameproxyservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForDomainNameProxyServiceWithChan(request *SaveBatchTaskForDomainNameProxyServiceRequest) (<-chan *SaveBatchTaskForDomainNameProxyServiceResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForDomainNameProxyServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForDomainNameProxyService(request)
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

// SaveBatchTaskForDomainNameProxyServiceWithCallback invokes the domain_intl.SaveBatchTaskForDomainNameProxyService API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskfordomainnameproxyservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForDomainNameProxyServiceWithCallback(request *SaveBatchTaskForDomainNameProxyServiceRequest, callback func(response *SaveBatchTaskForDomainNameProxyServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForDomainNameProxyServiceResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForDomainNameProxyService(request)
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

// SaveBatchTaskForDomainNameProxyServiceRequest is the request struct for api SaveBatchTaskForDomainNameProxyService
type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	*requests.RpcRequest
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
	DomainName   *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	Status       requests.Boolean `position:"Query" name:"status"`
}

// SaveBatchTaskForDomainNameProxyServiceResponse is the response struct for api SaveBatchTaskForDomainNameProxyService
type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForDomainNameProxyServiceRequest creates a request to invoke SaveBatchTaskForDomainNameProxyService API
func CreateSaveBatchTaskForDomainNameProxyServiceRequest() (request *SaveBatchTaskForDomainNameProxyServiceRequest) {
	request = &SaveBatchTaskForDomainNameProxyServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForDomainNameProxyService", "domain", "openAPI")
	return
}

// CreateSaveBatchTaskForDomainNameProxyServiceResponse creates a response to parse from SaveBatchTaskForDomainNameProxyService response
func CreateSaveBatchTaskForDomainNameProxyServiceResponse() (response *SaveBatchTaskForDomainNameProxyServiceResponse) {
	response = &SaveBatchTaskForDomainNameProxyServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
