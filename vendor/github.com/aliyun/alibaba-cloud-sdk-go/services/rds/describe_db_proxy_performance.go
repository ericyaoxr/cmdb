package rds

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

// DescribeDBProxyPerformance invokes the rds.DescribeDBProxyPerformance API synchronously
func (client *Client) DescribeDBProxyPerformance(request *DescribeDBProxyPerformanceRequest) (response *DescribeDBProxyPerformanceResponse, err error) {
	response = CreateDescribeDBProxyPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBProxyPerformanceWithChan invokes the rds.DescribeDBProxyPerformance API asynchronously
func (client *Client) DescribeDBProxyPerformanceWithChan(request *DescribeDBProxyPerformanceRequest) (<-chan *DescribeDBProxyPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBProxyPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBProxyPerformance(request)
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

// DescribeDBProxyPerformanceWithCallback invokes the rds.DescribeDBProxyPerformance API asynchronously
func (client *Client) DescribeDBProxyPerformanceWithCallback(request *DescribeDBProxyPerformanceRequest, callback func(response *DescribeDBProxyPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBProxyPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBProxyPerformance(request)
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

// DescribeDBProxyPerformanceRequest is the request struct for api DescribeDBProxyPerformance
type DescribeDBProxyPerformanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MetricsName          string           `position:"Query" name:"MetricsName"`
	StartTime            string           `position:"Query" name:"StartTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBProxyInstanceType  string           `position:"Query" name:"DBProxyInstanceType"`
}

// DescribeDBProxyPerformanceResponse is the response struct for api DescribeDBProxyPerformance
type DescribeDBProxyPerformanceResponse struct {
	*responses.BaseResponse
	RequestId       string                                      `json:"RequestId" xml:"RequestId"`
	DBInstanceId    string                                      `json:"DBInstanceId" xml:"DBInstanceId"`
	StartTime       string                                      `json:"StartTime" xml:"StartTime"`
	EndTime         string                                      `json:"EndTime" xml:"EndTime"`
	PerformanceKeys PerformanceKeysInDescribeDBProxyPerformance `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

// CreateDescribeDBProxyPerformanceRequest creates a request to invoke DescribeDBProxyPerformance API
func CreateDescribeDBProxyPerformanceRequest() (request *DescribeDBProxyPerformanceRequest) {
	request = &DescribeDBProxyPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBProxyPerformance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBProxyPerformanceResponse creates a response to parse from DescribeDBProxyPerformance response
func CreateDescribeDBProxyPerformanceResponse() (response *DescribeDBProxyPerformanceResponse) {
	response = &DescribeDBProxyPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
