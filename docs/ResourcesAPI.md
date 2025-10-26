# \ResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelResourceBackupExclusion**](ResourcesAPI.md#CancelResourceBackupExclusion) | **Patch** /v1/projects/{projectId}/resources/{id}/include | Cancel Resource Backup Exclusion
[**ExcludeResourceFromBackup**](ResourcesAPI.md#ExcludeResourceFromBackup) | **Patch** /v1/projects/{projectId}/resources/{id}/exclude | Exclude Resource from Backup
[**GetResource**](ResourcesAPI.md#GetResource) | **Get** /v1/projects/{projectId}/resources/{id} | Get Resource
[**ListResources**](ResourcesAPI.md#ListResources) | **Post** /v1/projects/{projectId}/resources | List Resources
[**OverrideDataClasses**](ResourcesAPI.md#OverrideDataClasses) | **Patch** /v1/projects/{projectId}/resources/{id}/data-classifications | Override Data Classes
[**OverrideEnvironment**](ResourcesAPI.md#OverrideEnvironment) | **Patch** /v1/projects/{projectId}/resources/{id}/environments | Override Environment
[**RemoveDataClassesOverride**](ResourcesAPI.md#RemoveDataClassesOverride) | **Delete** /v1/projects/{projectId}/resources/{id}/data-classifications | Remove Data Classes Override
[**RemoveEnvironmentOverride**](ResourcesAPI.md#RemoveEnvironmentOverride) | **Delete** /v1/projects/{projectId}/resources/{id}/environments | Remove Environment Override



## CancelResourceBackupExclusion

> CancelExclusionFromBackupResponse CancelResourceBackupExclusion(ctx, projectId, id).Execute()

Cancel Resource Backup Exclusion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to cancel the backup exclusion for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.CancelResourceBackupExclusion(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CancelResourceBackupExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelResourceBackupExclusion`: CancelExclusionFromBackupResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CancelResourceBackupExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to cancel the backup exclusion for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelResourceBackupExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CancelExclusionFromBackupResponse**](CancelExclusionFromBackupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExcludeResourceFromBackup

> ExcludeFromBackupResponse ExcludeResourceFromBackup(ctx, projectId, id).Execute()

Exclude Resource from Backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to exclude from backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ExcludeResourceFromBackup(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ExcludeResourceFromBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExcludeResourceFromBackup`: ExcludeFromBackupResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ExcludeResourceFromBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to exclude from backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExcludeResourceFromBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExcludeFromBackupResponse**](ExcludeFromBackupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResource

> GetResourceResponse GetResource(ctx, id, projectId).Execute()

Get Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned ID of the resource to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetResource(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResource`: GetResourceResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Eon-assigned ID of the resource to retrieve. | 
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetResourceResponse**](GetResourceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> ListResourcesResponse ListResources(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListInventoryRequest(listInventoryRequest).Execute()

List Resources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "f9304613-dddb-52fe-b883-f5e671a868a3" // string | ID of the project whose resources you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	pageToken := "pageToken_example" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)
	listInventoryRequest := *openapiclient.NewListInventoryRequest() // ListInventoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ListResources(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListInventoryRequest(listInventoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResources`: ListResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose resources you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]
 **listInventoryRequest** | [**ListInventoryRequest**](ListInventoryRequest.md) |  | 

### Return type

[**ListResourcesResponse**](ListResourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideDataClasses

> OverrideDataClassificationsResponse OverrideDataClasses(ctx, projectId, id).OverrideDataClassificationsRequest(overrideDataClassificationsRequest).Execute()

Override Data Classes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to override.
	overrideDataClassificationsRequest := *openapiclient.NewOverrideDataClassificationsRequest() // OverrideDataClassificationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.OverrideDataClasses(context.Background(), projectId, id).OverrideDataClassificationsRequest(overrideDataClassificationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.OverrideDataClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideDataClasses`: OverrideDataClassificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.OverrideDataClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to override. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideDataClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **overrideDataClassificationsRequest** | [**OverrideDataClassificationsRequest**](OverrideDataClassificationsRequest.md) |  | 

### Return type

[**OverrideDataClassificationsResponse**](OverrideDataClassificationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideEnvironment

> OverrideEnvironmentResponse OverrideEnvironment(ctx, projectId, id).OverrideEnvironmentRequest(overrideEnvironmentRequest).Execute()

Override Environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to override.
	overrideEnvironmentRequest := *openapiclient.NewOverrideEnvironmentRequest() // OverrideEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.OverrideEnvironment(context.Background(), projectId, id).OverrideEnvironmentRequest(overrideEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.OverrideEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideEnvironment`: OverrideEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.OverrideEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to override. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **overrideEnvironmentRequest** | [**OverrideEnvironmentRequest**](OverrideEnvironmentRequest.md) |  | 

### Return type

[**OverrideEnvironmentResponse**](OverrideEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDataClassesOverride

> RemoveDataClassesOverride(ctx, projectId, id).Execute()

Remove Data Classes Override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to remove the override from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.RemoveDataClassesOverride(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.RemoveDataClassesOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to remove the override from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDataClassesOverrideRequest struct via the builder pattern


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


## RemoveEnvironmentOverride

> RemoveEnvironmentOverride(ctx, projectId, id).Execute()

Remove Environment Override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | Eon-assigned ID of the resource to remove the override from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.RemoveEnvironmentOverride(context.Background(), projectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.RemoveEnvironmentOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned ID of the resource to remove the override from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEnvironmentOverrideRequest struct via the builder pattern


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

