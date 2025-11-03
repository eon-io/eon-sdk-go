# \JobsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupJob**](JobsAPI.md#GetBackupJob) | **Get** /v1/projects/{projectId}/backup-jobs/{jobId} | Get Backup Job
[**GetRestoreJob**](JobsAPI.md#GetRestoreJob) | **Get** /v1/projects/{projectId}/restore-jobs/{jobId} | Get Restore Job
[**ListBackupJobs**](JobsAPI.md#ListBackupJobs) | **Post** /v1/projects/{projectId}/backup-jobs | List Backup Jobs
[**ListRestoreJobs**](JobsAPI.md#ListRestoreJobs) | **Post** /v1/projects/{projectId}/restore-jobs | List Restore Jobs



## GetBackupJob

> GetBackupJobResponse GetBackupJob(ctx, jobId, projectId).Execute()

Get Backup Job



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
	jobId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | ID of the job to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the job is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetBackupJob(context.Background(), jobId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetBackupJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupJob`: GetBackupJobResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetBackupJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 
**projectId** | **string** | ID of the project the job is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBackupJobResponse**](GetBackupJobResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreJob

> GetRestoreJobResponse GetRestoreJob(ctx, jobId, projectId).Execute()

Get Restore Job



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
	jobId := "806f6781-1d66-4c7d-9f0e-e7da04e12541" // string | ID of the job to retrieve.
	projectId := "1ee34dc5-0a7c-4e56-a820-917371e05c8d" // string | ID of the project the job is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetRestoreJob(context.Background(), jobId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetRestoreJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreJob`: GetRestoreJobResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetRestoreJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | ID of the job to retrieve. | 
**projectId** | **string** | ID of the project the job is in. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRestoreJobResponse**](GetRestoreJobResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupJobs

> ListBackupJobsResponse ListBackupJobs(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListBackupJobsRequest(listBackupJobsRequest).Execute()

List Backup Jobs



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
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)
	listBackupJobsRequest := *openapiclient.NewListBackupJobsRequest() // ListBackupJobsRequest | Filter and sort options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.ListBackupJobs(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListBackupJobsRequest(listBackupJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.ListBackupJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupJobs`: ListBackupJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.ListBackupJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]
 **listBackupJobsRequest** | [**ListBackupJobsRequest**](ListBackupJobsRequest.md) | Filter and sort options. | 

### Return type

[**ListBackupJobsResponse**](ListBackupJobsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRestoreJobs

> ListRestoreJobsResponse ListRestoreJobs(ctx, projectId).PageToken(pageToken).PageSize(pageSize).ListRestoreJobsRequest(listRestoreJobsRequest).Execute()

List Restore Jobs



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
	projectId := "6b3ea428-f6a4-5bb5-8fb2-e4d5d2d920ce" // string | ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)
	listRestoreJobsRequest := *openapiclient.NewListRestoreJobsRequest() // ListRestoreJobsRequest | Filter and sort options.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.ListRestoreJobs(context.Background(), projectId).PageToken(pageToken).PageSize(pageSize).ListRestoreJobsRequest(listRestoreJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.ListRestoreJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRestoreJobs`: ListRestoreJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.ListRestoreJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRestoreJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]
 **listRestoreJobsRequest** | [**ListRestoreJobsRequest**](ListRestoreJobsRequest.md) | Filter and sort options.  | 

### Return type

[**ListRestoreJobsResponse**](ListRestoreJobsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

