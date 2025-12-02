# \AccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectRestoreAccount**](AccountsAPI.md#ConnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts | Connect Restore Account
[**ConnectSourceAccount**](AccountsAPI.md#ConnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts | Connect Source Account
[**DeleteRestoreAccountConnectivityConfig**](AccountsAPI.md#DeleteRestoreAccountConnectivityConfig) | **Delete** /v1/projects/{projectId}/restore-accounts/{accountId}/connectivity-config | Delete Restore Account Connectivity Configuration
[**DisableRestoreAccountMetricsConfig**](AccountsAPI.md#DisableRestoreAccountMetricsConfig) | **Delete** /v1/projects/{projectId}/restore-accounts/{accountId}/metrics-config | Disable Restore Account Metrics
[**DisableSourceAccountMetricsConfig**](AccountsAPI.md#DisableSourceAccountMetricsConfig) | **Delete** /v1/projects/{projectId}/source-accounts/{accountId}/metrics-config | Disable Source Account Metrics
[**DisconnectRestoreAccount**](AccountsAPI.md#DisconnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts/{accountId}/disconnect | Disconnect Restore Account
[**DisconnectSourceAccount**](AccountsAPI.md#DisconnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts/{accountId}/disconnect | Disconnect Source Account
[**EnableRestoreAccountMetricsConfig**](AccountsAPI.md#EnableRestoreAccountMetricsConfig) | **Put** /v1/projects/{projectId}/restore-accounts/{accountId}/metrics-config | Configure Restore Account Metrics
[**EnableSourceAccountMetricsConfig**](AccountsAPI.md#EnableSourceAccountMetricsConfig) | **Put** /v1/projects/{projectId}/source-accounts/{accountId}/metrics-config | Configure Source Account Metrics
[**GetAzureOnboardedTenants**](AccountsAPI.md#GetAzureOnboardedTenants) | **Get** /projects/{projectId}/azure/onboarded-tenants | Get onboarded Azure tenants for an Eon account
[**GetRestoreAccountConnectivityConfig**](AccountsAPI.md#GetRestoreAccountConnectivityConfig) | **Get** /v1/projects/{projectId}/restore-accounts/{accountId}/connectivity-config | Get Restore Account Connectivity Configuration
[**GetRestoreAccountMetricsConfig**](AccountsAPI.md#GetRestoreAccountMetricsConfig) | **Get** /v1/projects/{projectId}/restore-accounts/{accountId}/metrics-config | Get Restore Account Metrics Configuration
[**GetSourceAccountMetricsConfig**](AccountsAPI.md#GetSourceAccountMetricsConfig) | **Get** /v1/projects/{projectId}/source-accounts/{accountId}/metrics-config | Get Source Account Metrics Configuration
[**ListRestoreAccounts**](AccountsAPI.md#ListRestoreAccounts) | **Post** /v1/projects/{projectId}/restore-accounts/list | List Restore Accounts
[**ListSourceAccounts**](AccountsAPI.md#ListSourceAccounts) | **Post** /v1/projects/{projectId}/source-accounts/list | List Source Accounts
[**ReconnectRestoreAccount**](AccountsAPI.md#ReconnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts/{accountId}/reconnect | Reconnect Restore Account
[**ReconnectSourceAccount**](AccountsAPI.md#ReconnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts/{accountId}/reconnect | Reconnect Source Account
[**UpdateRestoreAccountConnectivityConfig**](AccountsAPI.md#UpdateRestoreAccountConnectivityConfig) | **Put** /v1/projects/{projectId}/restore-accounts/{accountId}/connectivity-config | Update Restore Account Connectivity Configuration



## ConnectRestoreAccount

> ConnectRestoreAccountResponse ConnectRestoreAccount(ctx, projectId).ConnectRestoreAccountRequest(connectRestoreAccountRequest).Execute()

Connect Restore Account



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	connectRestoreAccountRequest := *openapiclient.NewConnectRestoreAccountRequest(*openapiclient.NewRestoreAccountAttributesInput(openapiclient.Provider("AWS"))) // ConnectRestoreAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ConnectRestoreAccount(context.Background(), projectId).ConnectRestoreAccountRequest(connectRestoreAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ConnectRestoreAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectRestoreAccount`: ConnectRestoreAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ConnectRestoreAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectRestoreAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectRestoreAccountRequest** | [**ConnectRestoreAccountRequest**](ConnectRestoreAccountRequest.md) |  | 

### Return type

[**ConnectRestoreAccountResponse**](ConnectRestoreAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectSourceAccount

> ConnectSourceAccountResponse ConnectSourceAccount(ctx, projectId).ConnectSourceAccountRequest(connectSourceAccountRequest).Execute()

Connect Source Account



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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	connectSourceAccountRequest := *openapiclient.NewConnectSourceAccountRequest(*openapiclient.NewSourceAccountAttributesInput(openapiclient.Provider("AWS"))) // ConnectSourceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ConnectSourceAccount(context.Background(), projectId).ConnectSourceAccountRequest(connectSourceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ConnectSourceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectSourceAccount`: ConnectSourceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ConnectSourceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectSourceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectSourceAccountRequest** | [**ConnectSourceAccountRequest**](ConnectSourceAccountRequest.md) |  | 

### Return type

[**ConnectSourceAccountResponse**](ConnectSourceAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRestoreAccountConnectivityConfig

> DeleteRestoreAccountConnectivityConfig(ctx, accountId, projectId).Execute()

Delete Restore Account Connectivity Configuration



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Eon-assigned ID of the restore account.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose restore account connectivity configuration you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.DeleteRestoreAccountConnectivityConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DeleteRestoreAccountConnectivityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**projectId** | **string** | ID of the project whose restore account connectivity configuration you want to delete. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRestoreAccountConnectivityConfigRequest struct via the builder pattern


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


## DisableRestoreAccountMetricsConfig

> DisableRestoreAccountMetricsConfig(ctx, accountId, projectId).Execute()

Disable Restore Account Metrics



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the restore account.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose restore account metrics you want to disable. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.DisableRestoreAccountMetricsConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DisableRestoreAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**projectId** | **string** | ID of the project whose restore account metrics you want to disable. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRestoreAccountMetricsConfigRequest struct via the builder pattern


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


## DisableSourceAccountMetricsConfig

> DisableSourceAccountMetricsConfig(ctx, accountId, projectId).Execute()

Disable Source Account Metrics



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the source account.
	projectId := "d593811b-bfbc-515b-a21d-9fb95eb1071b" // string | ID of the project whose source account metrics configuration you want to disable. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.DisableSourceAccountMetricsConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DisableSourceAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the source account. | 
**projectId** | **string** | ID of the project whose source account metrics configuration you want to disable. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSourceAccountMetricsConfigRequest struct via the builder pattern


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


## DisconnectRestoreAccount

> DisconnectRestoreAccountResponse DisconnectRestoreAccount(ctx, projectId, accountId).Execute()

Disconnect Restore Account



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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	accountId := "64a698c5-540d-5d09-80f7-f4c39c7045c1" // string | Eon-assigned ID of the restore account to disconnect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.DisconnectRestoreAccount(context.Background(), projectId, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DisconnectRestoreAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectRestoreAccount`: DisconnectRestoreAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DisconnectRestoreAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 
**accountId** | **string** | Eon-assigned ID of the restore account to disconnect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectRestoreAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DisconnectRestoreAccountResponse**](DisconnectRestoreAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectSourceAccount

> DisconnectSourceAccountResponse DisconnectSourceAccount(ctx, projectId, accountId).Execute()

Disconnect Source Account



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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	accountId := "72d29280-a0be-59df-b33c-59f9015606c3" // string | Eon-assigned ID of the source account to disconnect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.DisconnectSourceAccount(context.Background(), projectId, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.DisconnectSourceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectSourceAccount`: DisconnectSourceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.DisconnectSourceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 
**accountId** | **string** | Eon-assigned ID of the source account to disconnect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectSourceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DisconnectSourceAccountResponse**](DisconnectSourceAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableRestoreAccountMetricsConfig

> EnableRestoreAccountMetricsConfigResponse EnableRestoreAccountMetricsConfig(ctx, accountId, projectId).EnableRestoreAccountMetricsConfigRequest(enableRestoreAccountMetricsConfigRequest).Execute()

Configure Restore Account Metrics



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the restore account.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose restore account metrics you want to configure. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	enableRestoreAccountMetricsConfigRequest := *openapiclient.NewEnableRestoreAccountMetricsConfigRequest() // EnableRestoreAccountMetricsConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.EnableRestoreAccountMetricsConfig(context.Background(), accountId, projectId).EnableRestoreAccountMetricsConfigRequest(enableRestoreAccountMetricsConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.EnableRestoreAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableRestoreAccountMetricsConfig`: EnableRestoreAccountMetricsConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.EnableRestoreAccountMetricsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**projectId** | **string** | ID of the project whose restore account metrics you want to configure. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRestoreAccountMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableRestoreAccountMetricsConfigRequest** | [**EnableRestoreAccountMetricsConfigRequest**](EnableRestoreAccountMetricsConfigRequest.md) |  | 

### Return type

[**EnableRestoreAccountMetricsConfigResponse**](EnableRestoreAccountMetricsConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableSourceAccountMetricsConfig

> EnableSourceAccountMetricsConfigResponse EnableSourceAccountMetricsConfig(ctx, accountId, projectId).EnableSourceAccountMetricsConfigRequest(enableSourceAccountMetricsConfigRequest).Execute()

Configure Source Account Metrics



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the source account.
	projectId := "b9c79c3f-399c-5e32-a8fb-762f87bebd7b" // string | ID of the project whose source account metrics you want to configure. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	enableSourceAccountMetricsConfigRequest := *openapiclient.NewEnableSourceAccountMetricsConfigRequest() // EnableSourceAccountMetricsConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.EnableSourceAccountMetricsConfig(context.Background(), accountId, projectId).EnableSourceAccountMetricsConfigRequest(enableSourceAccountMetricsConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.EnableSourceAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableSourceAccountMetricsConfig`: EnableSourceAccountMetricsConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.EnableSourceAccountMetricsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the source account. | 
**projectId** | **string** | ID of the project whose source account metrics you want to configure. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSourceAccountMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enableSourceAccountMetricsConfigRequest** | [**EnableSourceAccountMetricsConfigRequest**](EnableSourceAccountMetricsConfigRequest.md) |  | 

### Return type

[**EnableSourceAccountMetricsConfigResponse**](EnableSourceAccountMetricsConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureOnboardedTenants

> AzureOnboardedTenantsResponse GetAzureOnboardedTenants(ctx, projectId).Execute()

Get onboarded Azure tenants for an Eon account

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAzureOnboardedTenants(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAzureOnboardedTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureOnboardedTenants`: AzureOnboardedTenantsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAzureOnboardedTenants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureOnboardedTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureOnboardedTenantsResponse**](AzureOnboardedTenantsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreAccountConnectivityConfig

> GetRestoreAccountConnectivityConfigResponse GetRestoreAccountConnectivityConfig(ctx, accountId, projectId).Execute()

Get Restore Account Connectivity Configuration



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Eon-assigned ID of the restore account.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose restore account connectivity configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetRestoreAccountConnectivityConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetRestoreAccountConnectivityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreAccountConnectivityConfig`: GetRestoreAccountConnectivityConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetRestoreAccountConnectivityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**projectId** | **string** | ID of the project whose restore account connectivity configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreAccountConnectivityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRestoreAccountConnectivityConfigResponse**](GetRestoreAccountConnectivityConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreAccountMetricsConfig

> GetRestoreAccountMetricsConfigResponse GetRestoreAccountMetricsConfig(ctx, accountId, projectId).Execute()

Get Restore Account Metrics Configuration



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the restore account.
	projectId := "4bd9ca0f-b6cd-5a00-b6ce-76960e76c4ba" // string | ID of the project whose restore account metrics configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetRestoreAccountMetricsConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetRestoreAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreAccountMetricsConfig`: GetRestoreAccountMetricsConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetRestoreAccountMetricsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**projectId** | **string** | ID of the project whose restore account metrics configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreAccountMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRestoreAccountMetricsConfigResponse**](GetRestoreAccountMetricsConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountMetricsConfig

> GetSourceAccountMetricsConfigResponse GetSourceAccountMetricsConfig(ctx, accountId, projectId).Execute()

Get Source Account Metrics Configuration



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
	accountId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | Eon-assigned ID of the source account.
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose source account metrics configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetSourceAccountMetricsConfig(context.Background(), accountId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetSourceAccountMetricsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceAccountMetricsConfig`: GetSourceAccountMetricsConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetSourceAccountMetricsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Eon-assigned ID of the source account. | 
**projectId** | **string** | ID of the project whose source account metrics configuration you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountMetricsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSourceAccountMetricsConfigResponse**](GetSourceAccountMetricsConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRestoreAccounts

> ListRestoreAccountsResponse ListRestoreAccounts(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListRestoreAccountsRequest(listRestoreAccountsRequest).Execute()

List Restore Accounts



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
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project whose restore accounts you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)
	listRestoreAccountsRequest := *openapiclient.NewListRestoreAccountsRequest() // ListRestoreAccountsRequest | Filter options.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListRestoreAccounts(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListRestoreAccountsRequest(listRestoreAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListRestoreAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRestoreAccounts`: ListRestoreAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListRestoreAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose restore accounts you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRestoreAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]
 **listRestoreAccountsRequest** | [**ListRestoreAccountsRequest**](ListRestoreAccountsRequest.md) | Filter options.  | 

### Return type

[**ListRestoreAccountsResponse**](ListRestoreAccountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceAccounts

> ListSourceAccountsResponse ListSourceAccounts(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListSourceAccountsRequest(listSourceAccountsRequest).Execute()

List Source Accounts



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
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project whose source accounts you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)
	listSourceAccountsRequest := *openapiclient.NewListSourceAccountsRequest() // ListSourceAccountsRequest | Filter options.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListSourceAccounts(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListSourceAccountsRequest(listSourceAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListSourceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSourceAccounts`: ListSourceAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListSourceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose source accounts you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]
 **listSourceAccountsRequest** | [**ListSourceAccountsRequest**](ListSourceAccountsRequest.md) | Filter options.  | 

### Return type

[**ListSourceAccountsResponse**](ListSourceAccountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReconnectRestoreAccount

> ReconnectRestoreAccountResponse ReconnectRestoreAccount(ctx, projectId, accountId).Execute()

Reconnect Restore Account



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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project whose restore account you want to reconnect. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	accountId := "72d29280-a0be-59df-b33c-59f9015606c3" // string | Eon-assigned ID of the restore account to reconnect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ReconnectRestoreAccount(context.Background(), projectId, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReconnectRestoreAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconnectRestoreAccount`: ReconnectRestoreAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ReconnectRestoreAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose restore account you want to reconnect. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 
**accountId** | **string** | Eon-assigned ID of the restore account to reconnect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconnectRestoreAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReconnectRestoreAccountResponse**](ReconnectRestoreAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReconnectSourceAccount

> ReconnectSourceAccountResponse ReconnectSourceAccount(ctx, projectId, accountId).Execute()

Reconnect Source Account



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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project whose source account you want to reconnect. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	accountId := "72d29280-a0be-59df-b33c-59f9015606c3" // string | Eon-assigned ID of the source account to reconnect.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ReconnectSourceAccount(context.Background(), projectId, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReconnectSourceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconnectSourceAccount`: ReconnectSourceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ReconnectSourceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose source account you want to reconnect. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 
**accountId** | **string** | Eon-assigned ID of the source account to reconnect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconnectSourceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReconnectSourceAccountResponse**](ReconnectSourceAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRestoreAccountConnectivityConfig

> UpdateRestoreAccountConnectivityConfigResponse UpdateRestoreAccountConnectivityConfig(ctx, accountId, projectId).UpdateRestoreAccountConnectivityConfigRequest(updateRestoreAccountConnectivityConfigRequest).Execute()

Update Restore Account Connectivity Configuration



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
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the restore account
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project whose restore account connectivity configuration you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	updateRestoreAccountConnectivityConfigRequest := *openapiclient.NewUpdateRestoreAccountConnectivityConfigRequest() // UpdateRestoreAccountConnectivityConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateRestoreAccountConnectivityConfig(context.Background(), accountId, projectId).UpdateRestoreAccountConnectivityConfigRequest(updateRestoreAccountConnectivityConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateRestoreAccountConnectivityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRestoreAccountConnectivityConfig`: UpdateRestoreAccountConnectivityConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateRestoreAccountConnectivityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | ID of the restore account | 
**projectId** | **string** | ID of the project whose restore account connectivity configuration you want to update. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRestoreAccountConnectivityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRestoreAccountConnectivityConfigRequest** | [**UpdateRestoreAccountConnectivityConfigRequest**](UpdateRestoreAccountConnectivityConfigRequest.md) |  | 

### Return type

[**UpdateRestoreAccountConnectivityConfigResponse**](UpdateRestoreAccountConnectivityConfigResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

