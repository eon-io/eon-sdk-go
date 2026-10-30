# \ListAwsVpcsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAwsVpcs**](ListAwsVpcsAPI.md#ListAwsVpcs) | **Get** /projects/{projectId}/restore-accounts/{accountId}/regions/{region}/aws-vpcs | List VPCs



## ListAwsVpcs

> ListAwsVpcsResponse ListAwsVpcs(ctx, projectId, accountId, region).Execute()

List VPCs



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the project the restore account is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console. 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Eon-assigned ID of the restore account.
	region := "region_example" // string | Region where the VPCs are located.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAwsVpcsAPI.ListAwsVpcs(context.Background(), projectId, accountId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAwsVpcsAPI.ListAwsVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAwsVpcs`: ListAwsVpcsResponse
	fmt.Fprintf(os.Stdout, "Response from `ListAwsVpcsAPI.ListAwsVpcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the restore account is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**region** | **string** | Region where the VPCs are located. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListAwsVpcsResponse**](ListAwsVpcsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

