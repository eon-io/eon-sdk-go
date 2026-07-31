# \BackupPostureControlsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupPostureControl**](BackupPostureControlsAPI.md#CreateBackupPostureControl) | **Post** /v1/projects/{projectId}/backup-posture-controls | Create a backup posture control
[**DeleteBackupPostureControl**](BackupPostureControlsAPI.md#DeleteBackupPostureControl) | **Delete** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Delete a backup posture control
[**GetBackupPostureControl**](BackupPostureControlsAPI.md#GetBackupPostureControl) | **Get** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Get a backup posture control
[**ListBackupPostureControls**](BackupPostureControlsAPI.md#ListBackupPostureControls) | **Post** /v1/projects/{projectId}/backup-posture-controls/list | List backup posture controls
[**UpdateBackupPostureControl**](BackupPostureControlsAPI.md#UpdateBackupPostureControl) | **Put** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Update a backup posture control



## CreateBackupPostureControl

> BackupPostureControl CreateBackupPostureControl(ctx, projectId).CreateBackupPostureControlRequest(createBackupPostureControlRequest).Execute()

Create a backup posture control



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
	createBackupPostureControlRequest := *openapiclient.NewCreateBackupPostureControlRequest("Name_example", openapiclient.Severity("HIGH"), "TODO", "TODO") // CreateBackupPostureControlRequest | The definition of the backup posture control to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPostureControlsAPI.CreateBackupPostureControl(context.Background(), projectId).CreateBackupPostureControlRequest(createBackupPostureControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPostureControlsAPI.CreateBackupPostureControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupPostureControl`: BackupPostureControl
	fmt.Fprintf(os.Stdout, "Response from `BackupPostureControlsAPI.CreateBackupPostureControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBackupPostureControlRequest** | [**CreateBackupPostureControlRequest**](CreateBackupPostureControlRequest.md) | The definition of the backup posture control to create. | 

### Return type

[**BackupPostureControl**](BackupPostureControl.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupPostureControl

> DeleteBackupPostureControl(ctx, projectId, controlId).Execute()

Delete a backup posture control



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
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPostureControlsAPI.DeleteBackupPostureControl(context.Background(), projectId, controlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPostureControlsAPI.DeleteBackupPostureControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**controlId** | **string** | ID of the backup posture control to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupPostureControl

> BackupPostureControl GetBackupPostureControl(ctx, projectId, controlId).Execute()

Get a backup posture control



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
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPostureControlsAPI.GetBackupPostureControl(context.Background(), projectId, controlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPostureControlsAPI.GetBackupPostureControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupPostureControl`: BackupPostureControl
	fmt.Fprintf(os.Stdout, "Response from `BackupPostureControlsAPI.GetBackupPostureControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**controlId** | **string** | ID of the backup posture control to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BackupPostureControl**](BackupPostureControl.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupPostureControls

> ListBackupPostureControlsResponse ListBackupPostureControls(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListBackupPostureControlsRequest(listBackupPostureControlsRequest).Execute()

List backup posture controls



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
	pageToken := "pageToken_example" // string | Opaque cursor for the next page of results, returned as `nextPageToken` by a previous call. Reuse the same request body for stable pagination.  (optional)
	pageSize := int32(56) // int32 | Maximum number of backup posture controls to return. (optional) (default to 50)
	listBackupPostureControlsRequest := *openapiclient.NewListBackupPostureControlsRequest() // ListBackupPostureControlsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPostureControlsAPI.ListBackupPostureControls(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListBackupPostureControlsRequest(listBackupPostureControlsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPostureControlsAPI.ListBackupPostureControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupPostureControls`: ListBackupPostureControlsResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupPostureControlsAPI.ListBackupPostureControls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupPostureControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Opaque cursor for the next page of results, returned as &#x60;nextPageToken&#x60; by a previous call. Reuse the same request body for stable pagination.  | 
 **pageSize** | **int32** | Maximum number of backup posture controls to return. | [default to 50]
 **listBackupPostureControlsRequest** | [**ListBackupPostureControlsRequest**](ListBackupPostureControlsRequest.md) |  | 

### Return type

[**ListBackupPostureControlsResponse**](ListBackupPostureControlsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupPostureControl

> BackupPostureControl UpdateBackupPostureControl(ctx, projectId, controlId).UpdateBackupPostureControlRequest(updateBackupPostureControlRequest).Execute()

Update a backup posture control



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
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to update
	updateBackupPostureControlRequest := *openapiclient.NewUpdateBackupPostureControlRequest("Name_example", openapiclient.Severity("HIGH"), "TODO", "TODO") // UpdateBackupPostureControlRequest | The full desired state of the backup posture control.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPostureControlsAPI.UpdateBackupPostureControl(context.Background(), projectId, controlId).UpdateBackupPostureControlRequest(updateBackupPostureControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPostureControlsAPI.UpdateBackupPostureControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBackupPostureControl`: BackupPostureControl
	fmt.Fprintf(os.Stdout, "Response from `BackupPostureControlsAPI.UpdateBackupPostureControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**controlId** | **string** | ID of the backup posture control to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBackupPostureControlRequest** | [**UpdateBackupPostureControlRequest**](UpdateBackupPostureControlRequest.md) | The full desired state of the backup posture control. | 

### Return type

[**BackupPostureControl**](BackupPostureControl.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

