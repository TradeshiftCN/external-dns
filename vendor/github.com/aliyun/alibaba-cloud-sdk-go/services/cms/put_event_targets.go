package cms

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

// PutEventTargets invokes the cms.PutEventTargets API synchronously
// api document: https://help.aliyun.com/api/cms/puteventtargets.html
func (client *Client) PutEventTargets(request *PutEventTargetsRequest) (response *PutEventTargetsResponse, err error) {
	response = CreatePutEventTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// PutEventTargetsWithChan invokes the cms.PutEventTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/puteventtargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutEventTargetsWithChan(request *PutEventTargetsRequest) (<-chan *PutEventTargetsResponse, <-chan error) {
	responseChan := make(chan *PutEventTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutEventTargets(request)
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

// PutEventTargetsWithCallback invokes the cms.PutEventTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/puteventtargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutEventTargetsWithCallback(request *PutEventTargetsRequest, callback func(response *PutEventTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutEventTargetsResponse
		var err error
		defer close(result)
		response, err = client.PutEventTargets(request)
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

// PutEventTargetsRequest is the request struct for api PutEventTargets
type PutEventTargetsRequest struct {
	*requests.RpcRequest
	WebhookParameters *[]PutEventTargetsWebhookParameters `position:"Query" name:"WebhookParameters"  type:"Repeated"`
	ContactParameters *[]PutEventTargetsContactParameters `position:"Query" name:"ContactParameters"  type:"Repeated"`
	FcParameters      *[]PutEventTargetsFcParameters      `position:"Query" name:"FcParameters"  type:"Repeated"`
	RuleName          string                              `position:"Query" name:"RuleName"`
	MnsParameters     *[]PutEventTargetsMnsParameters     `position:"Query" name:"MnsParameters"  type:"Repeated"`
}

// PutEventTargetsWebhookParameters is a repeated param struct in PutEventTargetsRequest
type PutEventTargetsWebhookParameters struct {
	Id       string `name:"Id"`
	Protocol string `name:"Protocol"`
	Url      string `name:"Url"`
	Method   string `name:"Method"`
}

// PutEventTargetsContactParameters is a repeated param struct in PutEventTargetsRequest
type PutEventTargetsContactParameters struct {
	Id               string `name:"Id"`
	ContactGroupName string `name:"ContactGroupName"`
	Level            string `name:"Level"`
}

// PutEventTargetsFcParameters is a repeated param struct in PutEventTargetsRequest
type PutEventTargetsFcParameters struct {
	Id           string `name:"Id"`
	Region       string `name:"Region"`
	ServiceName  string `name:"ServiceName"`
	FunctionName string `name:"FunctionName"`
}

// PutEventTargetsMnsParameters is a repeated param struct in PutEventTargetsRequest
type PutEventTargetsMnsParameters struct {
	Id     string `name:"Id"`
	Region string `name:"Region"`
	Queue  string `name:"Queue"`
}

// PutEventTargetsResponse is the response struct for api PutEventTargets
type PutEventTargetsResponse struct {
	*responses.BaseResponse
	Success                 bool                               `json:"Success" xml:"Success"`
	Code                    string                             `json:"Code" xml:"Code"`
	Message                 string                             `json:"Message" xml:"Message"`
	RequestId               string                             `json:"RequestId" xml:"RequestId"`
	ParameterCount          string                             `json:"ParameterCount" xml:"ParameterCount"`
	FailedParameterCount    string                             `json:"FailedParameterCount" xml:"FailedParameterCount"`
	ContactParameters       ContactParametersInPutEventTargets `json:"ContactParameters" xml:"ContactParameters"`
	MnsParameters           MnsParametersInPutEventTargets     `json:"MnsParameters" xml:"MnsParameters"`
	FcParameters            FcParametersInPutEventTargets      `json:"FcParameters" xml:"FcParameters"`
	FailedContactParameters FailedContactParameters            `json:"FailedContactParameters" xml:"FailedContactParameters"`
	FailedMnsParameters     FailedMnsParameters                `json:"FailedMnsParameters" xml:"FailedMnsParameters"`
	FailedFcParameters      FailedFcParameters                 `json:"FailedFcParameters" xml:"FailedFcParameters"`
}

// CreatePutEventTargetsRequest creates a request to invoke PutEventTargets API
func CreatePutEventTargetsRequest() (request *PutEventTargetsRequest) {
	request = &PutEventTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "PutEventTargets", "cms", "openAPI")
	return
}

// CreatePutEventTargetsResponse creates a response to parse from PutEventTargets response
func CreatePutEventTargetsResponse() (response *PutEventTargetsResponse) {
	response = &PutEventTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
