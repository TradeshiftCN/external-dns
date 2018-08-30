package mts

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

// UpdateMediaCover invokes the mts.UpdateMediaCover API synchronously
// api document: https://help.aliyun.com/api/mts/updatemediacover.html
func (client *Client) UpdateMediaCover(request *UpdateMediaCoverRequest) (response *UpdateMediaCoverResponse, err error) {
	response = CreateUpdateMediaCoverResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMediaCoverWithChan invokes the mts.UpdateMediaCover API asynchronously
// api document: https://help.aliyun.com/api/mts/updatemediacover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMediaCoverWithChan(request *UpdateMediaCoverRequest) (<-chan *UpdateMediaCoverResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaCoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMediaCover(request)
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

// UpdateMediaCoverWithCallback invokes the mts.UpdateMediaCover API asynchronously
// api document: https://help.aliyun.com/api/mts/updatemediacover.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMediaCoverWithCallback(request *UpdateMediaCoverRequest, callback func(response *UpdateMediaCoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaCoverResponse
		var err error
		defer close(result)
		response, err = client.UpdateMediaCover(request)
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

// UpdateMediaCoverRequest is the request struct for api UpdateMediaCover
type UpdateMediaCoverRequest struct {
	*requests.RpcRequest
	CoverURL             string           `position:"Query" name:"CoverURL"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
}

// UpdateMediaCoverResponse is the response struct for api UpdateMediaCover
type UpdateMediaCoverResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateMediaCoverRequest creates a request to invoke UpdateMediaCover API
func CreateUpdateMediaCoverRequest() (request *UpdateMediaCoverRequest) {
	request = &UpdateMediaCoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCover", "mts", "openAPI")
	return
}

// CreateUpdateMediaCoverResponse creates a response to parse from UpdateMediaCover response
func CreateUpdateMediaCoverResponse() (response *UpdateMediaCoverResponse) {
	response = &UpdateMediaCoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
