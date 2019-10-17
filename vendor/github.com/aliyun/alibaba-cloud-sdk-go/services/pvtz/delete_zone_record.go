package pvtz

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

// DeleteZoneRecord invokes the pvtz.DeleteZoneRecord API synchronously
// api document: https://help.aliyun.com/api/pvtz/deletezonerecord.html
func (client *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) (response *DeleteZoneRecordResponse, err error) {
	response = CreateDeleteZoneRecordResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteZoneRecordWithChan invokes the pvtz.DeleteZoneRecord API asynchronously
// api document: https://help.aliyun.com/api/pvtz/deletezonerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteZoneRecordWithChan(request *DeleteZoneRecordRequest) (<-chan *DeleteZoneRecordResponse, <-chan error) {
	responseChan := make(chan *DeleteZoneRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteZoneRecord(request)
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

// DeleteZoneRecordWithCallback invokes the pvtz.DeleteZoneRecord API asynchronously
// api document: https://help.aliyun.com/api/pvtz/deletezonerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteZoneRecordWithCallback(request *DeleteZoneRecordRequest, callback func(response *DeleteZoneRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteZoneRecordResponse
		var err error
		defer close(result)
		response, err = client.DeleteZoneRecord(request)
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

// DeleteZoneRecordRequest is the request struct for api DeleteZoneRecord
type DeleteZoneRecordRequest struct {
	*requests.RpcRequest
	RecordId     requests.Integer `position:"Query" name:"RecordId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DeleteZoneRecordResponse is the response struct for api DeleteZoneRecord
type DeleteZoneRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RecordId  int    `json:"RecordId" xml:"RecordId"`
}

// CreateDeleteZoneRecordRequest creates a request to invoke DeleteZoneRecord API
func CreateDeleteZoneRecordRequest() (request *DeleteZoneRecordRequest) {
	request = &DeleteZoneRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DeleteZoneRecord", "pvtz", "openAPI")
	return
}

// CreateDeleteZoneRecordResponse creates a response to parse from DeleteZoneRecord response
func CreateDeleteZoneRecordResponse() (response *DeleteZoneRecordResponse) {
	response = &DeleteZoneRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
