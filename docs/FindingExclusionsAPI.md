# \FindingExclusionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFindingExclusion**](FindingExclusionsAPI.md#CreateFindingExclusion) | **Post** /v1/projects/{projectId}/finding-exclusions | Create Finding Exclusion
[**DeleteFindingExclusion**](FindingExclusionsAPI.md#DeleteFindingExclusion) | **Delete** /v1/projects/{projectId}/finding-exclusions/{findingExclusionId} | Delete Finding Exclusion
[**GetFindingExclusion**](FindingExclusionsAPI.md#GetFindingExclusion) | **Get** /v1/projects/{projectId}/finding-exclusions/{findingExclusionId} | Get Finding Exclusion
[**ListFindingExclusions**](FindingExclusionsAPI.md#ListFindingExclusions) | **Post** /v1/projects/{projectId}/finding-exclusions/list | List Finding Exclusions
[**UpdateFindingExclusion**](FindingExclusionsAPI.md#UpdateFindingExclusion) | **Put** /v1/projects/{projectId}/finding-exclusions/{findingExclusionId} | Update Finding Exclusion



## CreateFindingExclusion

> CreateFindingExclusionResponse CreateFindingExclusion(ctx, projectId).CreateFindingExclusionRequest(createFindingExclusionRequest).Execute()

Create Finding Exclusion



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	createFindingExclusionRequest := *openapiclient.NewCreateFindingExclusionRequest(openapiclient.FindingExclusionScope("RESOURCE"), "/var/cache", openapiclient.FindingObjectType("PATH"), openapiclient.FindingExclusionDetectorType("RANSOMWARE_BEHAVIOR")) // CreateFindingExclusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingExclusionsAPI.CreateFindingExclusion(context.Background(), projectId).CreateFindingExclusionRequest(createFindingExclusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingExclusionsAPI.CreateFindingExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFindingExclusion`: CreateFindingExclusionResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingExclusionsAPI.CreateFindingExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFindingExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFindingExclusionRequest** | [**CreateFindingExclusionRequest**](CreateFindingExclusionRequest.md) |  | 

### Return type

[**CreateFindingExclusionResponse**](CreateFindingExclusionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFindingExclusion

> DeleteFindingExclusion(ctx, projectId, findingExclusionId).Execute()

Delete Finding Exclusion



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	findingExclusionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Finding exclusion ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingExclusionsAPI.DeleteFindingExclusion(context.Background(), projectId, findingExclusionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingExclusionsAPI.DeleteFindingExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**findingExclusionId** | **string** | Finding exclusion ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFindingExclusionRequest struct via the builder pattern


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


## GetFindingExclusion

> GetFindingExclusionResponse GetFindingExclusion(ctx, projectId, findingExclusionId).Execute()

Get Finding Exclusion



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	findingExclusionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Finding exclusion ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingExclusionsAPI.GetFindingExclusion(context.Background(), projectId, findingExclusionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingExclusionsAPI.GetFindingExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFindingExclusion`: GetFindingExclusionResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingExclusionsAPI.GetFindingExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**findingExclusionId** | **string** | Finding exclusion ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFindingExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFindingExclusionResponse**](GetFindingExclusionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFindingExclusions

> ListFindingExclusionsResponse ListFindingExclusions(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListFindingExclusionsRequest(listFindingExclusionsRequest).Execute()

List Finding Exclusions



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	pageToken := "pageToken_example" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response, and send it with the same `pageSize` and request body.  (optional)
	pageSize := int32(56) // int32 | Maximum number of finding exclusions to return. Keep the same value for every page of a listing: the page token counts pages of this size.  (optional) (default to 50)
	listFindingExclusionsRequest := *openapiclient.NewListFindingExclusionsRequest() // ListFindingExclusionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingExclusionsAPI.ListFindingExclusions(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListFindingExclusionsRequest(listFindingExclusionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingExclusionsAPI.ListFindingExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFindingExclusions`: ListFindingExclusionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingExclusionsAPI.ListFindingExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFindingExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response, and send it with the same &#x60;pageSize&#x60; and request body.  | 
 **pageSize** | **int32** | Maximum number of finding exclusions to return. Keep the same value for every page of a listing: the page token counts pages of this size.  | [default to 50]
 **listFindingExclusionsRequest** | [**ListFindingExclusionsRequest**](ListFindingExclusionsRequest.md) |  | 

### Return type

[**ListFindingExclusionsResponse**](ListFindingExclusionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFindingExclusion

> UpdateFindingExclusionResponse UpdateFindingExclusion(ctx, projectId, findingExclusionId).UpdateFindingExclusionRequest(updateFindingExclusionRequest).Execute()

Update Finding Exclusion



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	findingExclusionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Finding exclusion ID.
	updateFindingExclusionRequest := *openapiclient.NewUpdateFindingExclusionRequest(openapiclient.FindingExclusionScope("RESOURCE"), "/var/cache", openapiclient.FindingObjectType("PATH"), openapiclient.FindingExclusionDetectorType("RANSOMWARE_BEHAVIOR")) // UpdateFindingExclusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingExclusionsAPI.UpdateFindingExclusion(context.Background(), projectId, findingExclusionId).UpdateFindingExclusionRequest(updateFindingExclusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingExclusionsAPI.UpdateFindingExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFindingExclusion`: UpdateFindingExclusionResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingExclusionsAPI.UpdateFindingExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the finding exclusions. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**findingExclusionId** | **string** | Finding exclusion ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFindingExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFindingExclusionRequest** | [**UpdateFindingExclusionRequest**](UpdateFindingExclusionRequest.md) |  | 

### Return type

[**UpdateFindingExclusionResponse**](UpdateFindingExclusionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

