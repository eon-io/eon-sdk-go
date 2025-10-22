# \CreateApiCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiCredentials**](CreateApiCredentialsAPI.md#CreateApiCredentials) | **Post** /v1/api-credentials | Create a new api credentials



## CreateApiCredentials

> CreateApiCredentialsResponse CreateApiCredentials(ctx).CreateApiCredentialsRequest(createApiCredentialsRequest).Execute()

Create a new api credentials



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
	createApiCredentialsRequest := *openapiclient.NewCreateApiCredentialsRequest("Name_example", "RoleId_example", []string{"Projects_example"}) // CreateApiCredentialsRequest | The request body for creating api credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreateApiCredentialsAPI.CreateApiCredentials(context.Background()).CreateApiCredentialsRequest(createApiCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateApiCredentialsAPI.CreateApiCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiCredentials`: CreateApiCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CreateApiCredentialsAPI.CreateApiCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiCredentialsRequest** | [**CreateApiCredentialsRequest**](CreateApiCredentialsRequest.md) | The request body for creating api credentials | 

### Return type

[**CreateApiCredentialsResponse**](CreateApiCredentialsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

