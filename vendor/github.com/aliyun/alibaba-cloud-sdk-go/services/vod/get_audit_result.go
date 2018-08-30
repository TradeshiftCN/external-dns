package vod

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

// GetAuditResult invokes the vod.GetAuditResult API synchronously
// api document: https://help.aliyun.com/api/vod/getauditresult.html
func (client *Client) GetAuditResult(request *GetAuditResultRequest) (response *GetAuditResultResponse, err error) {
	response = CreateGetAuditResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuditResultWithChan invokes the vod.GetAuditResult API asynchronously
// api document: https://help.aliyun.com/api/vod/getauditresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuditResultWithChan(request *GetAuditResultRequest) (<-chan *GetAuditResultResponse, <-chan error) {
	responseChan := make(chan *GetAuditResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuditResult(request)
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

// GetAuditResultWithCallback invokes the vod.GetAuditResult API asynchronously
// api document: https://help.aliyun.com/api/vod/getauditresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuditResultWithCallback(request *GetAuditResultRequest, callback func(response *GetAuditResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuditResultResponse
		var err error
		defer close(result)
		response, err = client.GetAuditResult(request)
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

// GetAuditResultRequest is the request struct for api GetAuditResult
type GetAuditResultRequest struct {
	*requests.RpcRequest
	VideoId string `position:"Query" name:"VideoId"`
}

// GetAuditResultResponse is the response struct for api GetAuditResult
type GetAuditResultResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	AIAuditResult AIAuditResult `json:"AIAuditResult" xml:"AIAuditResult"`
}

// CreateGetAuditResultRequest creates a request to invoke GetAuditResult API
func CreateGetAuditResultRequest() (request *GetAuditResultRequest) {
	request = &GetAuditResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAuditResult", "vod", "openAPI")
	return
}

// CreateGetAuditResultResponse creates a response to parse from GetAuditResult response
func CreateGetAuditResultResponse() (response *GetAuditResultResponse) {
	response = &GetAuditResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
