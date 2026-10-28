# \ListAwsInstanceTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAwsInstanceTypes**](ListAwsInstanceTypesAPI.md#ListAwsInstanceTypes) | **Get** /projects/{projectId}/restore-accounts/{accountId}/regions/{region}/aws-instance-types | List Instance Types



## ListAwsInstanceTypes

> ListInstanceTypesResponse ListAwsInstanceTypes(ctx, projectId, accountId, region).SubnetId(subnetId).VolumeCount(volumeCount).TotalVolumeBytes(totalVolumeBytes).Execute()

List Instance Types



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
	subnetId := "subnetId_example" // string | The subnet ID to get instance types for
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the restore account
	region := "region_example" // string | The region to get instance types for
	volumeCount := int32(56) // int32 | Number of guest volumes the operator plans to restore. Used to size the host-memory reserve for instant-restore host selection. Optional; the server uses a default when absent.  (optional)
	totalVolumeBytes := int64(789) // int64 | Sum of guest-volume sizes in bytes. Used to size the host-memory reserve (mmap working-set term). Optional; the server uses a default when absent.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAwsInstanceTypesAPI.ListAwsInstanceTypes(context.Background(), projectId, accountId, region).SubnetId(subnetId).VolumeCount(volumeCount).TotalVolumeBytes(totalVolumeBytes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAwsInstanceTypesAPI.ListAwsInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAwsInstanceTypes`: ListInstanceTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `ListAwsInstanceTypesAPI.ListAwsInstanceTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID | 
**accountId** | **string** | ID of the restore account | 
**region** | **string** | The region to get instance types for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetId** | **string** | The subnet ID to get instance types for | 


 **volumeCount** | **int32** | Number of guest volumes the operator plans to restore. Used to size the host-memory reserve for instant-restore host selection. Optional; the server uses a default when absent.  | 
 **totalVolumeBytes** | **int64** | Sum of guest-volume sizes in bytes. Used to size the host-memory reserve (mmap working-set term). Optional; the server uses a default when absent.  | 

### Return type

[**ListInstanceTypesResponse**](ListInstanceTypesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

