# \ListEncryptionKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEncryptionKeys**](ListEncryptionKeysAPI.md#ListEncryptionKeys) | **Get** /projects/{projectId}/restore-accounts/{accountId}/regions/{region}/encryption-keys | List Encryption Keys



## ListEncryptionKeys

> EncryptionKeys ListEncryptionKeys(ctx, projectId, accountId, region).IncludeAwsManagedKeys(includeAwsManagedKeys).Execute()

List Encryption Keys



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The project ID
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the restore account
	region := "region_example" // string | Region
	includeAwsManagedKeys := true // bool | Whether to include AWS managed keys in the response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListEncryptionKeysAPI.ListEncryptionKeys(context.Background(), projectId, accountId, region).IncludeAwsManagedKeys(includeAwsManagedKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListEncryptionKeysAPI.ListEncryptionKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEncryptionKeys`: EncryptionKeys
	fmt.Fprintf(os.Stdout, "Response from `ListEncryptionKeysAPI.ListEncryptionKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**accountId** | **string** | ID of the restore account | 
**region** | **string** | Region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEncryptionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeAwsManagedKeys** | **bool** | Whether to include AWS managed keys in the response | [default to false]

### Return type

[**EncryptionKeys**](EncryptionKeys.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

