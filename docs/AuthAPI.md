# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessToken**](AuthAPI.md#GetAccessToken) | **Post** /v1/token | Get Access Token
[**GetAccessTokenOAuth2**](AuthAPI.md#GetAccessTokenOAuth2) | **Post** /v1/oauth2/token | Get Access Token (OAuth 2)
[**RotateApiClientSecret**](AuthAPI.md#RotateApiClientSecret) | **Post** /v1/api-credentials/{clientId}/rotate | Rotate API Client Secret
[**RotateCurrentApiClientSecret**](AuthAPI.md#RotateCurrentApiClientSecret) | **Post** /v1/api-credentials/current/rotate | Rotate Current API Client Secret



## GetAccessToken

> TokenResponse GetAccessToken(ctx).ApiCredentials(apiCredentials).Execute()

Get Access Token



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
	apiCredentials := *openapiclient.NewApiCredentials("yaeiq73af5cbxibLiu22yw2t5ezeozxu5itv43uw2tyn3xpqprta", "eon_sample620a571c00caf772b8e1b69de703792ba12c9eadef689725e0489324a093a64b4459ffecbaaa441c377950628b6b9042af728056d1d060b8f5d5sample") // ApiCredentials | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAccessToken(context.Background()).ApiCredentials(apiCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCredentials** | [**ApiCredentials**](ApiCredentials.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessTokenOAuth2

> Oauth2TokenResponse GetAccessTokenOAuth2(ctx).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Execute()

Get Access Token (OAuth 2)



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
	clientId := "clientId_example" // string | API client ID.
	clientSecret := "clientSecret_example" // string | API client secret.
	grantType := "grantType_example" // string | The OAuth 2.0 grant type you're using. Must be `client_credentials`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAccessTokenOAuth2(context.Background()).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAccessTokenOAuth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessTokenOAuth2`: Oauth2TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAccessTokenOAuth2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenOAuth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | API client ID. | 
 **clientSecret** | **string** | API client secret. | 
 **grantType** | **string** | The OAuth 2.0 grant type you&#39;re using. Must be &#x60;client_credentials&#x60;.  | 

### Return type

[**Oauth2TokenResponse**](Oauth2TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateApiClientSecret

> ApiCredentials RotateApiClientSecret(ctx, clientId).Execute()

Rotate API Client Secret



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
	clientId := "yaeiq73af5cbxibLiu22yw2t5ezeozxu5itv43uw2tyn3xpqprta" // string | API client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RotateApiClientSecret(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RotateApiClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateApiClientSecret`: ApiCredentials
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RotateApiClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | API client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateApiClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCredentials**](ApiCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateCurrentApiClientSecret

> ApiCredentials RotateCurrentApiClientSecret(ctx).Execute()

Rotate Current API Client Secret



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RotateCurrentApiClientSecret(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RotateCurrentApiClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateCurrentApiClientSecret`: ApiCredentials
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RotateCurrentApiClientSecret`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateCurrentApiClientSecretRequest struct via the builder pattern


### Return type

[**ApiCredentials**](ApiCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

