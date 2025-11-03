# \DashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyStorageSummaries**](DashboardAPI.md#GetDailyStorageSummaries) | **Get** /v1/projects/{projectId}/dashboard/daily-storage-summary | Get Daily Storage Summaries



## GetDailyStorageSummaries

> DailyStorageSummaries GetDailyStorageSummaries(ctx, projectId).StartDate(startDate).EndDate(endDate).Execute()

Get Daily Storage Summaries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "f9304613-dddb-52fe-b883-f5e671a868a3" // string | ID of the project whose daily storage summaries you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings. 
	startDate := time.Now() // string | First day of the date range in `YYYY-MM-DD` format.
	endDate := time.Now() // string | Last day of the date range in `YYYY-MM-DD` format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetDailyStorageSummaries(context.Background(), projectId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetDailyStorageSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDailyStorageSummaries`: DailyStorageSummaries
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetDailyStorageSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose daily storage summaries you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-settings/api-credentials) page in your global settings.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyStorageSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | First day of the date range in &#x60;YYYY-MM-DD&#x60; format. | 
 **endDate** | **string** | Last day of the date range in &#x60;YYYY-MM-DD&#x60; format. | 

### Return type

[**DailyStorageSummaries**](DailyStorageSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

