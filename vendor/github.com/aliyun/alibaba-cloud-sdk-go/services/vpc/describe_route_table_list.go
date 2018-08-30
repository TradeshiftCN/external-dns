package vpc

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

// DescribeRouteTableList invokes the vpc.DescribeRouteTableList API synchronously
// api document: https://help.aliyun.com/api/vpc/describeroutetablelist.html
func (client *Client) DescribeRouteTableList(request *DescribeRouteTableListRequest) (response *DescribeRouteTableListResponse, err error) {
	response = CreateDescribeRouteTableListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRouteTableListWithChan invokes the vpc.DescribeRouteTableList API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeroutetablelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteTableListWithChan(request *DescribeRouteTableListRequest) (<-chan *DescribeRouteTableListResponse, <-chan error) {
	responseChan := make(chan *DescribeRouteTableListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouteTableList(request)
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

// DescribeRouteTableListWithCallback invokes the vpc.DescribeRouteTableList API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeroutetablelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRouteTableListWithCallback(request *DescribeRouteTableListRequest, callback func(response *DescribeRouteTableListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouteTableListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouteTableList(request)
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

// DescribeRouteTableListRequest is the request struct for api DescribeRouteTableList
type DescribeRouteTableListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	RouterType           string           `position:"Query" name:"RouterType"`
	RouteTableName       string           `position:"Query" name:"RouteTableName"`
	RouterId             string           `position:"Query" name:"RouterId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
}

// DescribeRouteTableListResponse is the response struct for api DescribeRouteTableList
type DescribeRouteTableListResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	Success         bool            `json:"Success" xml:"Success"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	RouterTableList RouterTableList `json:"RouterTableList" xml:"RouterTableList"`
}

// CreateDescribeRouteTableListRequest creates a request to invoke DescribeRouteTableList API
func CreateDescribeRouteTableListRequest() (request *DescribeRouteTableListRequest) {
	request = &DescribeRouteTableListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouteTableList", "vpc", "openAPI")
	return
}

// CreateDescribeRouteTableListResponse creates a response to parse from DescribeRouteTableList response
func CreateDescribeRouteTableListResponse() (response *DescribeRouteTableListResponse) {
	response = &DescribeRouteTableListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
