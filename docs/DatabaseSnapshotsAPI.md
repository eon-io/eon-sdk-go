# \DatabaseSnapshotsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryResult**](DatabaseSnapshotsAPI.md#GetQueryResult) | **Get** /v1/projects/{projectId}/queries/{queryId}/results | Get Query Result
[**GetQueryStatus**](DatabaseSnapshotsAPI.md#GetQueryStatus) | **Get** /v1/projects/{projectId}/queries/{queryId}/status | Get Query Status
[**RunQuery**](DatabaseSnapshotsAPI.md#RunQuery) | **Post** /v1/projects/{projectId}/snapshots/{snapshotId}/databases/query | Run Query



## GetQueryResult

> QueryDBResultResponse GetQueryResult(ctx, projectId, queryId).PageSize(pageSize).PageToken(pageToken).Execute()

Get Query Result



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
	projectId := "330e9e09-6cb8-52af-b532-a476f598cf94" // string | ID of the project that contains the query you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  
	queryId := "330e9e09-6cb8-52af-b532-a476f598cf94:3fdf3419-366b-5e58-847d-c4671f79ce8f" // string | Query ID returned by [Run Query](./run-query).
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 500)
	pageToken := "pageToken_example" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseSnapshotsAPI.GetQueryResult(context.Background(), projectId, queryId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSnapshotsAPI.GetQueryResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryResult`: QueryDBResultResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabaseSnapshotsAPI.GetQueryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the query you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.   | 
**queryId** | **string** | Query ID returned by [Run Query](./run-query). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 500]
 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response.  | 

### Return type

[**QueryDBResultResponse**](QueryDBResultResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryStatus

> QueryDBStatusResponse GetQueryStatus(ctx, projectId, queryId).Execute()

Get Query Status



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
	projectId := "330e9e09-6cb8-52af-b532-a476f598cf94" // string | ID of the project that contains the query status you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  
	queryId := "330e9e09-6cb8-52af-b532-a476f598cf94:3fdf3419-366b-5e58-847d-c4671f79ce8f" // string | Query ID returned by [Run Query](./run-query).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseSnapshotsAPI.GetQueryStatus(context.Background(), projectId, queryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSnapshotsAPI.GetQueryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryStatus`: QueryDBStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabaseSnapshotsAPI.GetQueryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the query status you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.   | 
**queryId** | **string** | Query ID returned by [Run Query](./run-query). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueryDBStatusResponse**](QueryDBStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunQuery

> QueryDBResponse RunQuery(ctx, projectId, snapshotId).QueryDBRequest(queryDBRequest).Execute()

Run Query



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
	projectId := "330e9e09-6cb8-52af-b532-a476f598cf94" // string | ID of the project that contains the database snapshot you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  
	snapshotId := "9888b6d9-8d8c-4c25-bcd9-7c71298eba0f" // string | ID of the database [snapshot](./list-resource-snapshots) you want to query. Snapshot must belong to an RDS resource. 
	queryDBRequest := *openapiclient.NewQueryDBRequest("postgres", "SELECT * FROM customers LIMIT 50;", "1e0b9aa3-a942-5e86-afe6-57606e952622", *openapiclient.NewQueryDBRestoreDestination()) // QueryDBRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabaseSnapshotsAPI.RunQuery(context.Background(), projectId, snapshotId).QueryDBRequest(queryDBRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabaseSnapshotsAPI.RunQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunQuery`: QueryDBResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabaseSnapshotsAPI.RunQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the database snapshot you want to retrieve. You can see your project ID in the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.   | 
**snapshotId** | **string** | ID of the database [snapshot](./list-resource-snapshots) you want to query. Snapshot must belong to an RDS resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryDBRequest** | [**QueryDBRequest**](QueryDBRequest.md) |  | 

### Return type

[**QueryDBResponse**](QueryDBResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

