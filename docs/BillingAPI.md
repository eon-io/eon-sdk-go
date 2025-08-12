# \BillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryCostData**](BillingAPI.md#QueryCostData) | **Post** /v1/cost-data | Query cost data



## QueryCostData

> QueryCostDataResponse QueryCostData(ctx).QueryCostDataRequest(queryCostDataRequest).PageSize(pageSize).PageToken(pageToken).Execute()

Query cost data



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
	queryCostDataRequest := *openapiclient.NewQueryCostDataRequest(*openapiclient.NewTimeFrame(time.Now(), time.Now())) // QueryCostDataRequest | 
	pageSize := int32(100) // int32 | Maximum number of records to return (optional)
	pageToken := "eyJwYWdlIjoxfQ==" // string | Token for retrieving next page of results (returned from previous request) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.QueryCostData(context.Background()).QueryCostDataRequest(queryCostDataRequest).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.QueryCostData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCostData`: QueryCostDataResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.QueryCostData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCostDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryCostDataRequest** | [**QueryCostDataRequest**](QueryCostDataRequest.md) |  | 
 **pageSize** | **int32** | Maximum number of records to return | 
 **pageToken** | **string** | Token for retrieving next page of results (returned from previous request) | 

### Return type

[**QueryCostDataResponse**](QueryCostDataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

