# \BackupPostureControlsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupPostureControl**](BackupPostureControlsAPI.md#CreateBackupPostureControl) | **Post** /v1/projects/{projectId}/backup-posture-controls | Create Backup Posture Control
[**DeleteBackupPostureControl**](BackupPostureControlsAPI.md#DeleteBackupPostureControl) | **Delete** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Delete Backup Posture Control
[**GetBackupPostureControl**](BackupPostureControlsAPI.md#GetBackupPostureControl) | **Get** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Get Backup Posture Control
[**ListBackupPostureControls**](BackupPostureControlsAPI.md#ListBackupPostureControls) | **Post** /v1/projects/{projectId}/backup-posture-controls/list | List Backup Posture Controls
[**UpdateBackupPostureControl**](BackupPostureControlsAPI.md#UpdateBackupPostureControl) | **Put** /v1/projects/{projectId}/backup-posture-controls/{controlId} | Update Backup Posture Control



## CreateBackupPostureControl

> BackupPostureControl CreateBackupPostureControl(ctx, projectId).CreateBackupPostureControlRequest(createBackupPostureControlRequest).Execute()

Create Backup Posture Control



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project you want to create a backup posture control in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	createBackupPostureControlRequest := *openapiclient.NewCreateBackupPostureControlRequest("Name_example", openapiclient.Severity("HIGH"), "TODO", "TODO") // CreateBackupPostureControlRequest | Backup posture control to create.

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
**projectId** | **string** | ID of the project you want to create a backup posture control in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBackupPostureControlRequest** | [**CreateBackupPostureControlRequest**](CreateBackupPostureControlRequest.md) | Backup posture control to create. | 

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

Delete Backup Posture Control



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose backup posture control you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to delete.

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
**projectId** | **string** | ID of the project whose backup posture control you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**controlId** | **string** | ID of the backup posture control to delete. | 

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

Get Backup Posture Control



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose backup posture control you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to retrieve.

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
**projectId** | **string** | ID of the project whose backup posture control you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**controlId** | **string** | ID of the backup posture control to retrieve. | 

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

List Backup Posture Controls



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose backup posture controls you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
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
**projectId** | **string** | ID of the project whose backup posture controls you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 

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

Update Backup Posture Control



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose backup posture control you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	controlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup posture control to update.
	updateBackupPostureControlRequest := *openapiclient.NewUpdateBackupPostureControlRequest("Name_example", openapiclient.Severity("HIGH"), "TODO", "TODO") // UpdateBackupPostureControlRequest | 

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
**projectId** | **string** | ID of the project whose backup posture control you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**controlId** | **string** | ID of the backup posture control to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupPostureControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBackupPostureControlRequest** | [**UpdateBackupPostureControlRequest**](UpdateBackupPostureControlRequest.md) |  | 

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

