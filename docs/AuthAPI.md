# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessToken**](AuthAPI.md#GetAccessToken) | **Post** /v1/token | Get Access Token
[**GetAccessTokenOAuth2**](AuthAPI.md#GetAccessTokenOAuth2) | **Post** /v1/oauth2/token | Get Access Token (OAuth2)



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

> Oauth2TokenResponse GetAccessTokenOAuth2(ctx).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).RefreshToken(refreshToken).Execute()

Get Access Token (OAuth2)



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
	clientId := "clientId_example" // string | Your integration's client ID.
	clientSecret := "clientSecret_example" // string | Your integration's client secret.
	grantType := "grantType_example" // string | The OAuth2 grant you are using. (optional)
	refreshToken := "refreshToken_example" // string | Refresh token. Required if `grant_type` is `refresh_token`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAccessTokenOAuth2(context.Background()).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).RefreshToken(refreshToken).Execute()
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
 **clientId** | **string** | Your integration&#39;s client ID. | 
 **clientSecret** | **string** | Your integration&#39;s client secret. | 
 **grantType** | **string** | The OAuth2 grant you are using. | 
 **refreshToken** | **string** | Refresh token. Required if &#x60;grant_type&#x60; is &#x60;refresh_token&#x60;.  | 

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

