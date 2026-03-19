# \OauthAppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuthApp**](OauthAppsAPI.md#CreateOAuthApp) | **Post** /v1/oauth-apps | Create an OAuth App
[**ListOAuthApps**](OauthAppsAPI.md#ListOAuthApps) | **Get** /v1/oauth-apps | List OAuth Apps



## CreateOAuthApp

> CreateOAuthAppResponse CreateOAuthApp(ctx).CreateOAuthAppRequest(createOAuthAppRequest).Execute()

Create an OAuth App



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
	createOAuthAppRequest := *openapiclient.NewCreateOAuthAppRequest("gcp-gemini") // CreateOAuthAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.CreateOAuthApp(context.Background()).CreateOAuthAppRequest(createOAuthAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.CreateOAuthApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOAuthApp`: CreateOAuthAppResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.CreateOAuthApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuthAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOAuthAppRequest** | [**CreateOAuthAppRequest**](CreateOAuthAppRequest.md) |  | 

### Return type

[**CreateOAuthAppResponse**](CreateOAuthAppResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuthApps

> ListOAuthAppsResponse ListOAuthApps(ctx).Execute()

List OAuth Apps



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
	resp, r, err := apiClient.OauthAppsAPI.ListOAuthApps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.ListOAuthApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOAuthApps`: ListOAuthAppsResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.ListOAuthApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuthAppsRequest struct via the builder pattern


### Return type

[**ListOAuthAppsResponse**](ListOAuthAppsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

