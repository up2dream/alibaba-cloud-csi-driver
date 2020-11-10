package ecs

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

// ExportImage invokes the ecs.ExportImage API synchronously
func (client *Client) ExportImage(request *ExportImageRequest) (response *ExportImageResponse, err error) {
	response = CreateExportImageResponse()
	err = client.DoAction(request, response)
	return
}

// ExportImageWithChan invokes the ecs.ExportImage API asynchronously
func (client *Client) ExportImageWithChan(request *ExportImageRequest) (<-chan *ExportImageResponse, <-chan error) {
	responseChan := make(chan *ExportImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportImage(request)
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

// ExportImageWithCallback invokes the ecs.ExportImage API asynchronously
func (client *Client) ExportImageWithCallback(request *ExportImageRequest, callback func(response *ExportImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportImageResponse
		var err error
		defer close(result)
		response, err = client.ExportImage(request)
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

// ExportImageRequest is the request struct for api ExportImage
type ExportImageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId              string           `position:"Query" name:"ImageId"`
	ImageFormat          string           `position:"Query" name:"ImageFormat"`
	OSSBucket            string           `position:"Query" name:"OSSBucket"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RoleName             string           `position:"Query" name:"RoleName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OSSPrefix            string           `position:"Query" name:"OSSPrefix"`
}

// ExportImageResponse is the response struct for api ExportImage
type ExportImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
}

// CreateExportImageRequest creates a request to invoke ExportImage API
func CreateExportImageRequest() (request *ExportImageRequest) {
	request = &ExportImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ExportImage", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportImageResponse creates a response to parse from ExportImage response
func CreateExportImageResponse() (response *ExportImageResponse) {
	response = &ExportImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
