# \ListAwsSecurityGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAwsSecurityGroups**](ListAwsSecurityGroupsAPI.md#ListAwsSecurityGroups) | **Get** /projects/{projectId}/restore-accounts/{accountId}/regions/{region}/aws-security-groups | List Security Groups



## ListAwsSecurityGroups

> RestoreSecurityGroups ListAwsSecurityGroups(ctx, projectId, accountId, region).VpcId(vpcId).SubnetId(subnetId).Execute()

List Security Groups



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
	region := "region_example" // string | Region where the security groups are located.
	vpcId := "vpcId_example" // string | ID of the VPC to filter security groups by. (optional)
	subnetId := "subnetId_example" // string | ID of the subnet to filter security groups by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAwsSecurityGroupsAPI.ListAwsSecurityGroups(context.Background(), projectId, accountId, region).VpcId(vpcId).SubnetId(subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAwsSecurityGroupsAPI.ListAwsSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAwsSecurityGroups`: RestoreSecurityGroups
	fmt.Fprintf(os.Stdout, "Response from `ListAwsSecurityGroupsAPI.ListAwsSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the restore account is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page in your global management console.  | 
**accountId** | **string** | Eon-assigned ID of the restore account. | 
**region** | **string** | Region where the security groups are located. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vpcId** | **string** | ID of the VPC to filter security groups by. | 
 **subnetId** | **string** | ID of the subnet to filter security groups by. | 

### Return type

[**RestoreSecurityGroups**](RestoreSecurityGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

