# \AccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectRestoreAccount**](AccountsAPI.md#ConnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts | Connect Restore Account
[**ConnectSourceAccount**](AccountsAPI.md#ConnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts | Connect Source Account
[**DisconnectRestoreAccount**](AccountsAPI.md#DisconnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts/{accountId}/disconnect | Disconnect Restore Account
[**DisconnectSourceAccount**](AccountsAPI.md#DisconnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts/{accountId}/disconnect | Disconnect Source Account
[**ListRestoreAccounts**](AccountsAPI.md#ListRestoreAccounts) | **Post** /v1/projects/{projectId}/restore-accounts/list | List Restore Accounts
[**ListSourceAccounts**](AccountsAPI.md#ListSourceAccounts) | **Post** /v1/projects/{projectId}/source-accounts/list | List Source Accounts
[**ReconnectRestoreAccount**](AccountsAPI.md#ReconnectRestoreAccount) | **Post** /v1/projects/{projectId}/restore-accounts/{accountId}/reconnect | Reconnect Restore Account
[**ReconnectSourceAccount**](AccountsAPI.md#ReconnectSourceAccount) | **Post** /v1/projects/{projectId}/source-accounts/{accountId}/reconnect | Reconnect Source Account



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
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	connectRestoreAccountRequest := *openapiclient.NewConnectRestoreAccountRequest("restore-sandbox", *openapiclient.NewAccountConfigInput(openapiclient.Provider("AWS"))) // ConnectRestoreAccountRequest | 

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
**projectId** | **string** | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	connectSourceAccountRequest := *openapiclient.NewConnectSourceAccountRequest("prod-eu", *openapiclient.NewAccountConfigInput(openapiclient.Provider("AWS"))) // ConnectSourceAccountRequest | 

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
**projectId** | **string** | ID of the project you want to connect the source account to. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
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
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
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
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
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
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project whose restore accounts you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional)
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
**projectId** | **string** | ID of the project whose restore accounts you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRestoreAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | 
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
	projectId := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the project whose source accounts you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional)
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
**projectId** | **string** | ID of the project whose source accounts you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | 
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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
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
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
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
	projectId := "cc4fe8ee-0c62-56f2-9fda-f27bc7753e55" // string | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
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
**projectId** | **string** | ID of the project. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
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

