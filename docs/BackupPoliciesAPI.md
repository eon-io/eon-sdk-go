# \BackupPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupPolicy**](BackupPoliciesAPI.md#CreateBackupPolicy) | **Post** /v1/projects/{projectId}/backup-policies | Create Backup Policy
[**DeleteBackupPolicy**](BackupPoliciesAPI.md#DeleteBackupPolicy) | **Delete** /v1/projects/{projectId}/backup-policies/{backupPolicyId} | Delete Backup Policy
[**GetBackupPolicy**](BackupPoliciesAPI.md#GetBackupPolicy) | **Get** /v1/projects/{projectId}/backup-policies/{backupPolicyId} | Get Backup Policy
[**ListBackupPolicies**](BackupPoliciesAPI.md#ListBackupPolicies) | **Post** /v1/projects/{projectId}/backup-policies/list | List Backup Policies
[**UpdateBackupPolicy**](BackupPoliciesAPI.md#UpdateBackupPolicy) | **Put** /v1/projects/{projectId}/backup-policies/{backupPolicyId} | Update Backup Policy



## CreateBackupPolicy

> CreateBackupPolicyResponse CreateBackupPolicy(ctx, projectId).CreateBackupPolicyRequest(createBackupPolicyRequest).Execute()

Create Backup Policy



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project you want to create a backup policy in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	createBackupPolicyRequest := *openapiclient.NewCreateBackupPolicyRequest("Production with PII", *openapiclient.NewBackupPolicyResourceSelector(openapiclient.ResourceSelectorMode("ALL")), *openapiclient.NewBackupPolicyPlan(openapiclient.BackupPolicyType("UNSPECIFIED"))) // CreateBackupPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPoliciesAPI.CreateBackupPolicy(context.Background(), projectId).CreateBackupPolicyRequest(createBackupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesAPI.CreateBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupPolicy`: CreateBackupPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesAPI.CreateBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project you want to create a backup policy in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBackupPolicyRequest** | [**CreateBackupPolicyRequest**](CreateBackupPolicyRequest.md) |  | 

### Return type

[**CreateBackupPolicyResponse**](CreateBackupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupPolicy

> DeleteBackupPolicy(ctx, backupPolicyId, projectId).Execute()

Delete Backup Policy



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
	backupPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup policy to delete
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose backup policy you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupPoliciesAPI.DeleteBackupPolicy(context.Background(), backupPolicyId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesAPI.DeleteBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupPolicyId** | **string** | ID of the backup policy to delete | 
**projectId** | **string** | ID of the project whose backup policy you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupPolicyRequest struct via the builder pattern


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


## GetBackupPolicy

> GetBackupPolicyResponse GetBackupPolicy(ctx, backupPolicyId, projectId).Execute()

Get Backup Policy



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
	backupPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup policy to retrieve
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose backup policy you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPoliciesAPI.GetBackupPolicy(context.Background(), backupPolicyId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesAPI.GetBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupPolicy`: GetBackupPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesAPI.GetBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupPolicyId** | **string** | ID of the backup policy to retrieve | 
**projectId** | **string** | ID of the project whose backup policy you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBackupPolicyResponse**](GetBackupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupPolicies

> ListBackupPoliciesResponse ListBackupPolicies(ctx, projectId).Execute()

List Backup Policies



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
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose backup policies you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPoliciesAPI.ListBackupPolicies(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesAPI.ListBackupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupPolicies`: ListBackupPoliciesResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesAPI.ListBackupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose backup policies you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBackupPoliciesResponse**](ListBackupPoliciesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupPolicy

> UpdateBackupPolicyResponse UpdateBackupPolicy(ctx, backupPolicyId, projectId).UpdateBackupPolicyRequest(updateBackupPolicyRequest).Execute()

Update Backup Policy



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
	backupPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup policy to update
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose backup policy you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	updateBackupPolicyRequest := *openapiclient.NewUpdateBackupPolicyRequest("Production with PII", *openapiclient.NewBackupPolicyResourceSelector(openapiclient.ResourceSelectorMode("ALL")), *openapiclient.NewBackupPolicyPlan(openapiclient.BackupPolicyType("UNSPECIFIED"))) // UpdateBackupPolicyRequest | The request body for updating a policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupPoliciesAPI.UpdateBackupPolicy(context.Background(), backupPolicyId, projectId).UpdateBackupPolicyRequest(updateBackupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupPoliciesAPI.UpdateBackupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBackupPolicy`: UpdateBackupPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupPoliciesAPI.UpdateBackupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupPolicyId** | **string** | ID of the backup policy to update | 
**projectId** | **string** | ID of the project whose backup policy you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBackupPolicyRequest** | [**UpdateBackupPolicyRequest**](UpdateBackupPolicyRequest.md) | The request body for updating a policy | 

### Return type

[**UpdateBackupPolicyResponse**](UpdateBackupPolicyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

