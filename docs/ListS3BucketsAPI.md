# \ListS3BucketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListS3Buckets**](ListS3BucketsAPI.md#ListS3Buckets) | **Get** /projects/{projectId}/restore-accounts/{accountId}/aws/s3-buckets | List S3 Buckets



## ListS3Buckets

> S3Buckets ListS3Buckets(ctx, projectId, accountId).Execute()

List S3 Buckets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListS3BucketsAPI.ListS3Buckets(context.Background(), projectId, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListS3BucketsAPI.ListS3Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListS3Buckets`: S3Buckets
	fmt.Fprintf(os.Stdout, "Response from `ListS3BucketsAPI.ListS3Buckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**accountId** | **string** | ID of the restore account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListS3BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**S3Buckets**](S3Buckets.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

