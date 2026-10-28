# \RestoreTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAttachRestoreTemplate**](RestoreTemplatesAPI.md#BulkAttachRestoreTemplate) | **Post** /v1/projects/{projectId}/restore-template/{restoreTemplateId}/bulk-attach | Bulk Attach Restore Template
[**CreateRestoreTemplate**](RestoreTemplatesAPI.md#CreateRestoreTemplate) | **Post** /v1/projects/{projectId}/inventory/restore-templates | Create Restore Template



## BulkAttachRestoreTemplate

> BulkAttachRestoreTemplate(ctx, projectId, restoreTemplateId).BulkAttachRestoreTemplateRequest(bulkAttachRestoreTemplateRequest).Execute()

Bulk Attach Restore Template



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project the resources belong to. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	restoreTemplateId := "de308488-df3d-5b82-8f37-477ef529f491" // string | ID of the restore template to attach to the resources.
	bulkAttachRestoreTemplateRequest := *openapiclient.NewBulkAttachRestoreTemplateRequest([]string{"ResourceIds_example"}) // BulkAttachRestoreTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RestoreTemplatesAPI.BulkAttachRestoreTemplate(context.Background(), projectId, restoreTemplateId).BulkAttachRestoreTemplateRequest(bulkAttachRestoreTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreTemplatesAPI.BulkAttachRestoreTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the resources belong to. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**restoreTemplateId** | **string** | ID of the restore template to attach to the resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkAttachRestoreTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkAttachRestoreTemplateRequest** | [**BulkAttachRestoreTemplateRequest**](BulkAttachRestoreTemplateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRestoreTemplate

> CreateRestoreTemplateResponse CreateRestoreTemplate(ctx, projectId).CreateRestoreTemplateRequest(createRestoreTemplateRequest).Execute()

Create Restore Template



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project you want to create a restore template in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	createRestoreTemplateRequest := *openapiclient.NewCreateRestoreTemplateRequest("EC2 production restore template", openapiclient.RestoreTemplateResourceType("RESTORE_TEMPLATE_RESOURCE_TYPE_UNSPECIFIED")) // CreateRestoreTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestoreTemplatesAPI.CreateRestoreTemplate(context.Background(), projectId).CreateRestoreTemplateRequest(createRestoreTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreTemplatesAPI.CreateRestoreTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRestoreTemplate`: CreateRestoreTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `RestoreTemplatesAPI.CreateRestoreTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project you want to create a restore template in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestoreTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRestoreTemplateRequest** | [**CreateRestoreTemplateRequest**](CreateRestoreTemplateRequest.md) |  | 

### Return type

[**CreateRestoreTemplateResponse**](CreateRestoreTemplateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

