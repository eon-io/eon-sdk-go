# \MpaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionApprovalRequest**](MpaAPI.md#CancelActionApprovalRequest) | **Post** /v1/projects/{projectId}/mpa/my-requests/{requestId}/cancel | Cancel Action Approval Request
[**CreateActionApprovalRequest**](MpaAPI.md#CreateActionApprovalRequest) | **Post** /v1/projects/{projectId}/mpa/my-requests/{requestId}/submit | Create Action Approval Request
[**GetMyActionApprovalRequest**](MpaAPI.md#GetMyActionApprovalRequest) | **Get** /v1/projects/{projectId}/mpa/my-requests/{requestId} | Get My Action Approval Request



## CancelActionApprovalRequest

> CancelMyMPARequestResponse CancelActionApprovalRequest(ctx, projectId, requestId).Execute()

Cancel Action Approval Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project ID
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action approval request ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MpaAPI.CancelActionApprovalRequest(context.Background(), projectId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MpaAPI.CancelActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActionApprovalRequest`: CancelMyMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `MpaAPI.CancelActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**requestId** | **string** | Action approval request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CancelMyMPARequestResponse**](CancelMyMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActionApprovalRequest

> SubmitMyMPARequestResponse CreateActionApprovalRequest(ctx, projectId, requestId).SubmitMyMPARequestRequest(submitMyMPARequestRequest).Execute()

Create Action Approval Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project ID
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action approval request ID.
	submitMyMPARequestRequest := *openapiclient.NewSubmitMyMPARequestRequest(openapiclient.MPASubmitAction("MPA_SUBMIT_ACTION_UNSPECIFIED")) // SubmitMyMPARequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MpaAPI.CreateActionApprovalRequest(context.Background(), projectId, requestId).SubmitMyMPARequestRequest(submitMyMPARequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MpaAPI.CreateActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActionApprovalRequest`: SubmitMyMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `MpaAPI.CreateActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**requestId** | **string** | Action approval request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **submitMyMPARequestRequest** | [**SubmitMyMPARequestRequest**](SubmitMyMPARequestRequest.md) |  | 

### Return type

[**SubmitMyMPARequestResponse**](SubmitMyMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyActionApprovalRequest

> GetMPARequestResponse GetMyActionApprovalRequest(ctx, projectId, requestId).Execute()

Get My Action Approval Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project ID
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The action approval request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MpaAPI.GetMyActionApprovalRequest(context.Background(), projectId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MpaAPI.GetMyActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyActionApprovalRequest`: GetMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `MpaAPI.GetMyActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**requestId** | **string** | The action approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetMPARequestResponse**](GetMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

