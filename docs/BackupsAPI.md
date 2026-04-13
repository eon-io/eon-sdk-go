# \BackupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TakeSnapshot**](BackupsAPI.md#TakeSnapshot) | **Post** /v1/projects/{projectId}/resources/{id}/take-snapshot | Take Snapshot



## TakeSnapshot

> TakeSnapshotResponse TakeSnapshot(ctx, projectId, id).TakeSnapshotRequest(takeSnapshotRequest).Execute()

Take Snapshot



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned ID of the resource to back up.
	takeSnapshotRequest := *openapiclient.NewTakeSnapshotRequest("0d79d713-78df-5870-882f-0089d05396b6", int32(30)) // TakeSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.TakeSnapshot(context.Background(), projectId, id).TakeSnapshotRequest(takeSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.TakeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TakeSnapshot`: TakeSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.TakeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**id** | **string** | Eon-assigned ID of the resource to back up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **takeSnapshotRequest** | [**TakeSnapshotRequest**](TakeSnapshotRequest.md) |  | 

### Return type

[**TakeSnapshotResponse**](TakeSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

