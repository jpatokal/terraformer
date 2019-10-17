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

// ModifyInstanceCrossBackupPolicy invokes the rds.ModifyInstanceCrossBackupPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/modifyinstancecrossbackuppolicy.html
func (client *Client) ModifyInstanceCrossBackupPolicy(request *ModifyInstanceCrossBackupPolicyRequest) (response *ModifyInstanceCrossBackupPolicyResponse, err error) {
	response = CreateModifyInstanceCrossBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceCrossBackupPolicyWithChan invokes the rds.ModifyInstanceCrossBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyinstancecrossbackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceCrossBackupPolicyWithChan(request *ModifyInstanceCrossBackupPolicyRequest) (<-chan *ModifyInstanceCrossBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceCrossBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceCrossBackupPolicy(request)
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

// ModifyInstanceCrossBackupPolicyWithCallback invokes the rds.ModifyInstanceCrossBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyinstancecrossbackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceCrossBackupPolicyWithCallback(request *ModifyInstanceCrossBackupPolicyRequest, callback func(response *ModifyInstanceCrossBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceCrossBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceCrossBackupPolicy(request)
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

// ModifyInstanceCrossBackupPolicyRequest is the request struct for api ModifyInstanceCrossBackupPolicy
type ModifyInstanceCrossBackupPolicyRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	CrossBackupType      string           `position:"Query" name:"CrossBackupType"`
	LogBackupEnabled     string           `position:"Query" name:"LogBackupEnabled"`
	BackupEnabled        string           `position:"Query" name:"BackupEnabled"`
	CrossBackupRegion    string           `position:"Query" name:"CrossBackupRegion"`
	StorageOwner         string           `position:"Query" name:"StorageOwner"`
	StorageType          string           `position:"Query" name:"StorageType"`
	Endpoint             string           `position:"Query" name:"Endpoint"`
	RetentType           requests.Integer `position:"Query" name:"RetentType"`
	Retention            requests.Integer `position:"Query" name:"Retention"`
	RelService           string           `position:"Query" name:"RelService"`
}

// ModifyInstanceCrossBackupPolicyResponse is the response struct for api ModifyInstanceCrossBackupPolicy
type ModifyInstanceCrossBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DBInstanceId      string `json:"DBInstanceId" xml:"DBInstanceId"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
	CrossBackupRegion string `json:"CrossBackupRegion" xml:"CrossBackupRegion"`
	CrossBackupType   string `json:"CrossBackupType" xml:"CrossBackupType"`
	BackupEnabled     string `json:"BackupEnabled" xml:"BackupEnabled"`
	LogBackupEnabled  string `json:"LogBackupEnabled" xml:"LogBackupEnabled"`
	StorageOwner      string `json:"StorageOwner" xml:"StorageOwner"`
	StorageType       string `json:"StorageType" xml:"StorageType"`
	Endpoint          string `json:"Endpoint" xml:"Endpoint"`
	RetentType        int    `json:"RetentType" xml:"RetentType"`
	Retention         int    `json:"Retention" xml:"Retention"`
}

// CreateModifyInstanceCrossBackupPolicyRequest creates a request to invoke ModifyInstanceCrossBackupPolicy API
func CreateModifyInstanceCrossBackupPolicyRequest() (request *ModifyInstanceCrossBackupPolicyRequest) {
	request = &ModifyInstanceCrossBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyInstanceCrossBackupPolicy", "rds", "openAPI")
	return
}

// CreateModifyInstanceCrossBackupPolicyResponse creates a response to parse from ModifyInstanceCrossBackupPolicy response
func CreateModifyInstanceCrossBackupPolicyResponse() (response *ModifyInstanceCrossBackupPolicyResponse) {
	response = &ModifyInstanceCrossBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
