# \SnapshotsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnapshot**](SnapshotsAPI.md#GetSnapshot) | **Get** /v1/projects/{projectId}/snapshots/{id} | Get Snapshot
[**ListResourceSnapshots**](SnapshotsAPI.md#ListResourceSnapshots) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots | List Resource Snapshots
[**RestoreAzureVmInstance**](SnapshotsAPI.md#RestoreAzureVmInstance) | **Post** /v1/projects/{projectId}/inventory/{id}/snapshots/{snapshotId}/restore-azure-vm-instance | Restore Azure VM Instance
[**RestoreBucket**](SnapshotsAPI.md#RestoreBucket) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-bucket | Restore Bucket
[**RestoreDatabase**](SnapshotsAPI.md#RestoreDatabase) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-rds-instance | Restore RDS Instance
[**RestoreDynamoDBTable**](SnapshotsAPI.md#RestoreDynamoDBTable) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-dynamo-db-table | Restore DynamoDB Table
[**RestoreEbsVolume**](SnapshotsAPI.md#RestoreEbsVolume) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-ec2-ebs-volume | Restore EBS Volume
[**RestoreEc2Instance**](SnapshotsAPI.md#RestoreEc2Instance) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-ec2-instance | Restore EC2 Instance
[**RestoreFiles**](SnapshotsAPI.md#RestoreFiles) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/restore-files | Restore Files
[**RestoreToEbsSnapshot**](SnapshotsAPI.md#RestoreToEbsSnapshot) | **Post** /v1/projects/{projectId}/resources/{id}/snapshots/{snapshotId}/convert-ec2-ebs-snapshot | Restore to EBS Snapshot



## GetSnapshot

> GetSnapshotResponse GetSnapshot(ctx, id, projectId).Execute()

Get Snapshot



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
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | ID of the Eon snapshot to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.GetSnapshot(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: GetSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Eon snapshot to retrieve. | 
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSnapshotResponse**](GetSnapshotResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceSnapshots

> ListInventorySnapshotsResponse ListResourceSnapshots(ctx, id, projectId).PageToken(pageToken).PageSize(pageSize).ListInventorySnapshotsRequest(listInventorySnapshotsRequest).Execute()

List Resource Snapshots



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
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned ID of the resource whose Eon snapshots you want to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	pageToken := "OGRjZmRkNjYtZjJiMy00MmQ3LWIyMjMtNzU2M2NmMThiYWM2fDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional)
	listInventorySnapshotsRequest := *openapiclient.NewListInventorySnapshotsRequest() // ListInventorySnapshotsRequest | Filter options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ListResourceSnapshots(context.Background(), id, projectId).PageToken(pageToken).PageSize(pageSize).ListInventorySnapshotsRequest(listInventorySnapshotsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ListResourceSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceSnapshots`: ListInventorySnapshotsResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ListResourceSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Eon-assigned ID of the resource whose Eon snapshots you want to retrieve. | 
**projectId** | **string** | ID of the project the resource is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | 
 **listInventorySnapshotsRequest** | [**ListInventorySnapshotsRequest**](ListInventorySnapshotsRequest.md) | Filter options. | 

### Return type

[**ListInventorySnapshotsResponse**](ListInventorySnapshotsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreAzureVmInstance

> RestoreJobInitiationResponse RestoreAzureVmInstance(ctx, projectId, id, snapshotId).RestoreAzureVmInstanceRequest(restoreAzureVmInstanceRequest).Execute()

Restore Azure VM Instance

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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "fb806ce1-1cd3-5034-928a-33a87be714da" // string | Eon-assigned resource ID.
	snapshotId := "3dc6c0c6-f94d-5e85-a174-4b981a4bb262" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreAzureVmInstanceRequest := *openapiclient.NewRestoreAzureVmInstanceRequest("1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewAzureVmInstanceRestoreDestination()) // RestoreAzureVmInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreAzureVmInstance(context.Background(), projectId, id, snapshotId).RestoreAzureVmInstanceRequest(restoreAzureVmInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreAzureVmInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreAzureVmInstance`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreAzureVmInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreAzureVmInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreAzureVmInstanceRequest** | [**RestoreAzureVmInstanceRequest**](RestoreAzureVmInstanceRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBucket

> RestoreJobInitiationResponse RestoreBucket(ctx, projectId, id, snapshotId).RestoreBucketRequest(restoreBucketRequest).Execute()

Restore Bucket



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned resource ID.
	snapshotId := "c11d3c11-7be5-4ee4-9eb8-2024d9c04904" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreBucketRequest := *openapiclient.NewRestoreBucketRequest(*openapiclient.NewObjectStorageDestination(), "1ee34dc5-0a7c-4e56-a820-917371e05c8d") // RestoreBucketRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreBucket(context.Background(), projectId, id, snapshotId).RestoreBucketRequest(restoreBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreBucket`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreBucketRequest** | [**RestoreBucketRequest**](RestoreBucketRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDatabase

> RestoreJobInitiationResponse RestoreDatabase(ctx, projectId, id, snapshotId).RestoreDbToRdsInstanceRequest(restoreDbToRdsInstanceRequest).Execute()

Restore RDS Instance



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned resource ID.
	snapshotId := "c11d3c11-7be5-4ee4-9eb8-2024d9c04904" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreDbToRdsInstanceRequest := *openapiclient.NewRestoreDbToRdsInstanceRequest("1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewDatabaseDestination()) // RestoreDbToRdsInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreDatabase(context.Background(), projectId, id, snapshotId).RestoreDbToRdsInstanceRequest(restoreDbToRdsInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreDatabase`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreDbToRdsInstanceRequest** | [**RestoreDbToRdsInstanceRequest**](RestoreDbToRdsInstanceRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDynamoDBTable

> RestoreJobInitiationResponse RestoreDynamoDBTable(ctx, projectId, id, snapshotId).RestoreDynamoDBTableRequest(restoreDynamoDBTableRequest).Execute()

Restore DynamoDB Table



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "2f97ca76-6a78-55d8-94d3-66c2f2cfff23" // string | Eon-assigned resource ID.
	snapshotId := "ac3014c2-9ab3-5d7f-ab4c-73412d6b9ef5" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreDynamoDBTableRequest := *openapiclient.NewRestoreDynamoDBTableRequest("1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewDynamodbTableRestoreDestination()) // RestoreDynamoDBTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreDynamoDBTable(context.Background(), projectId, id, snapshotId).RestoreDynamoDBTableRequest(restoreDynamoDBTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreDynamoDBTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreDynamoDBTable`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreDynamoDBTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDynamoDBTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreDynamoDBTableRequest** | [**RestoreDynamoDBTableRequest**](RestoreDynamoDBTableRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreEbsVolume

> RestoreJobInitiationResponse RestoreEbsVolume(ctx, projectId, id, snapshotId).RestoreVolumeToEbsRequest(restoreVolumeToEbsRequest).Execute()

Restore EBS Volume



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned resource ID.
	snapshotId := "c11d3c11-7be5-4ee4-9eb8-2024d9c04904" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreVolumeToEbsRequest := *openapiclient.NewRestoreVolumeToEbsRequest("vol-01a29e3ba811d4613", "1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewEbsRestoreDestination()) // RestoreVolumeToEbsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreEbsVolume(context.Background(), projectId, id, snapshotId).RestoreVolumeToEbsRequest(restoreVolumeToEbsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreEbsVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreEbsVolume`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreEbsVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreEbsVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreVolumeToEbsRequest** | [**RestoreVolumeToEbsRequest**](RestoreVolumeToEbsRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreEc2Instance

> RestoreJobInitiationResponse RestoreEc2Instance(ctx, projectId, id, snapshotId).RestoreAwsEc2InstanceRequest(restoreAwsEc2InstanceRequest).Execute()

Restore EC2 Instance



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "fb806ce1-1cd3-5034-928a-33a87be714da" // string | Eon-assigned resource ID.
	snapshotId := "3dc6c0c6-f94d-5e85-a174-4b981a4bb262" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreAwsEc2InstanceRequest := *openapiclient.NewRestoreAwsEc2InstanceRequest("1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewAwsEc2InstanceRestoreDestination()) // RestoreAwsEc2InstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreEc2Instance(context.Background(), projectId, id, snapshotId).RestoreAwsEc2InstanceRequest(restoreAwsEc2InstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreEc2Instance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreEc2Instance`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreEc2Instance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreEc2InstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreAwsEc2InstanceRequest** | [**RestoreAwsEc2InstanceRequest**](RestoreAwsEc2InstanceRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFiles

> RestoreJobInitiationResponse RestoreFiles(ctx, projectId, id, snapshotId).RestoreFilesRequest(restoreFilesRequest).Execute()

Restore Files



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned resource ID.
	snapshotId := "c11d3c11-7be5-4ee4-9eb8-2024d9c04904" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreFilesRequest := *openapiclient.NewRestoreFilesRequest(*openapiclient.NewObjectStorageDestination(), []openapiclient.FilePath{*openapiclient.NewFilePath("Path_example", false)}, "1ee34dc5-0a7c-4e56-a820-917371e05c8d") // RestoreFilesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreFiles(context.Background(), projectId, id, snapshotId).RestoreFilesRequest(restoreFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFiles`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreFilesRequest** | [**RestoreFilesRequest**](RestoreFilesRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreToEbsSnapshot

> RestoreJobInitiationResponse RestoreToEbsSnapshot(ctx, projectId, id, snapshotId).RestoreVolumeToEbsSnapshotRequest(restoreVolumeToEbsSnapshotRequest).Execute()

Restore to EBS Snapshot



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
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
	id := "043090df-9fe5-4f89-9859-45db589c2936" // string | Eon-assigned resource ID.
	snapshotId := "c11d3c11-7be5-4ee4-9eb8-2024d9c04904" // string | ID of the Eon [snapshot](./list-resource-snapshots) to restore.
	restoreVolumeToEbsSnapshotRequest := *openapiclient.NewRestoreVolumeToEbsSnapshotRequest("vol-01a29e3ba811d4613", "1ee34dc5-0a7c-4e56-a820-917371e05c8d", *openapiclient.NewEbsSnapshotRestoreDestination()) // RestoreVolumeToEbsSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.RestoreToEbsSnapshot(context.Background(), projectId, id, snapshotId).RestoreVolumeToEbsSnapshotRequest(restoreVolumeToEbsSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RestoreToEbsSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreToEbsSnapshot`: RestoreJobInitiationResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RestoreToEbsSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project the snapshot is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings.  | 
**id** | **string** | Eon-assigned resource ID. | 
**snapshotId** | **string** | ID of the Eon [snapshot](./list-resource-snapshots) to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreToEbsSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restoreVolumeToEbsSnapshotRequest** | [**RestoreVolumeToEbsSnapshotRequest**](RestoreVolumeToEbsSnapshotRequest.md) |  | 

### Return type

[**RestoreJobInitiationResponse**](RestoreJobInitiationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

