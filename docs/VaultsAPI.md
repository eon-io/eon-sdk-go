# \VaultsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVault**](VaultsAPI.md#CreateVault) | **Post** /v1/projects/{projectId}/vaults | Create Vault
[**GetVault**](VaultsAPI.md#GetVault) | **Get** /v1/projects/{projectId}/vaults/{vaultId} | Get Vault
[**ListVaults**](VaultsAPI.md#ListVaults) | **Post** /v1/projects/{projectId}/vaults/list | List Vaults
[**UpdateVault**](VaultsAPI.md#UpdateVault) | **Patch** /v1/projects/{projectId}/vaults/{vaultId} | Update Vault



## CreateVault

> CreateVaultResponse CreateVault(ctx, projectId).CreateVaultRequest(createVaultRequest).Execute()

Create Vault



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the vault is created in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	createVaultRequest := *openapiclient.NewCreateVaultRequest("EU central", "eu-central-1", *openapiclient.NewVaultProviderAttributesInput(openapiclient.Provider("AWS"))) // CreateVaultRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.CreateVault(context.Background(), projectId).CreateVaultRequest(createVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.CreateVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVault`: CreateVaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.CreateVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the vault is created in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVaultRequest** | [**CreateVaultRequest**](CreateVaultRequest.md) |  | 

### Return type

[**CreateVaultResponse**](CreateVaultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVault

> GetVaultResponse GetVault(ctx, vaultId, projectId).Execute()

Get Vault



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
	vaultId := "36388572-9c69-5309-b362-04e7d85ae503" // string | ID of the vault to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the vault is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.GetVault(context.Background(), vaultId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.GetVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVault`: GetVaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.GetVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | ID of the vault to retrieve. | 
**projectId** | **string** | ID of the project the vault is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVaultResponse**](GetVaultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVaults

> ListBackupVaultResponse ListVaults(ctx, projectId).PageToken(pageToken).PageSize(pageSize).Execute()

List Vaults



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project whose vaults you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.ListVaults(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.ListVaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVaults`: ListBackupVaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.ListVaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose vaults you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]

### Return type

[**ListBackupVaultResponse**](ListBackupVaultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVault

> UpdateVaultResponse UpdateVault(ctx, vaultId, projectId).UpdateVaultRequest(updateVaultRequest).Execute()

Update Vault



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
	vaultId := "36388572-9c69-5309-b362-04e7d85ae503" // string | ID of the vault to update.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the vault is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	updateVaultRequest := *openapiclient.NewUpdateVaultRequest("My new vault name") // UpdateVaultRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultsAPI.UpdateVault(context.Background(), vaultId, projectId).UpdateVaultRequest(updateVaultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultsAPI.UpdateVault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVault`: UpdateVaultResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultsAPI.UpdateVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultId** | **string** | ID of the vault to update. | 
**projectId** | **string** | ID of the project the vault is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVaultRequest** | [**UpdateVaultRequest**](UpdateVaultRequest.md) |  | 

### Return type

[**UpdateVaultResponse**](UpdateVaultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

