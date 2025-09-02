# QueryCostDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]CostDataRecord**](CostDataRecord.md) | Array of cost data records matching the requested query criteria. Each record contains cost and usage information grouped by the specified dimensions and time granularity.  | 
**TotalCount** | **int32** | Total number of cost records that matched the time frame and filter options.  | 
**ResourceCount** | **int32** | Total number of resources that incurred costs during the specified time frame.  | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 

## Methods

### NewQueryCostDataResponse

`func NewQueryCostDataResponse(records []CostDataRecord, totalCount int32, resourceCount int32, ) *QueryCostDataResponse`

NewQueryCostDataResponse instantiates a new QueryCostDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCostDataResponseWithDefaults

`func NewQueryCostDataResponseWithDefaults() *QueryCostDataResponse`

NewQueryCostDataResponseWithDefaults instantiates a new QueryCostDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *QueryCostDataResponse) GetRecords() []CostDataRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueryCostDataResponse) GetRecordsOk() (*[]CostDataRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueryCostDataResponse) SetRecords(v []CostDataRecord)`

SetRecords sets Records field to given value.


### GetTotalCount

`func (o *QueryCostDataResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QueryCostDataResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QueryCostDataResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResourceCount

`func (o *QueryCostDataResponse) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *QueryCostDataResponse) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *QueryCostDataResponse) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.


### GetNextToken

`func (o *QueryCostDataResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *QueryCostDataResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *QueryCostDataResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *QueryCostDataResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


