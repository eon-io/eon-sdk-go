# QueryCostDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]QueryCostDataResponseRecordsInner**](QueryCostDataResponseRecordsInner.md) | Array of cost records based on requested filters and grouping | 
**TotalCount** | **int32** | Total number of records available | 
**TotalUniqueResources** | **int32** | Total number of unique resources in time range | 
**FilteredUniqueResources** | **int32** | Number of unique resources after applying filters in time range | 
**NextToken** | Pointer to **string** | Token for retrieving next page of results | [optional] 

## Methods

### NewQueryCostDataResponse

`func NewQueryCostDataResponse(records []QueryCostDataResponseRecordsInner, totalCount int32, totalUniqueResources int32, filteredUniqueResources int32, ) *QueryCostDataResponse`

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

`func (o *QueryCostDataResponse) GetRecords() []QueryCostDataResponseRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueryCostDataResponse) GetRecordsOk() (*[]QueryCostDataResponseRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueryCostDataResponse) SetRecords(v []QueryCostDataResponseRecordsInner)`

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


### GetTotalUniqueResources

`func (o *QueryCostDataResponse) GetTotalUniqueResources() int32`

GetTotalUniqueResources returns the TotalUniqueResources field if non-nil, zero value otherwise.

### GetTotalUniqueResourcesOk

`func (o *QueryCostDataResponse) GetTotalUniqueResourcesOk() (*int32, bool)`

GetTotalUniqueResourcesOk returns a tuple with the TotalUniqueResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUniqueResources

`func (o *QueryCostDataResponse) SetTotalUniqueResources(v int32)`

SetTotalUniqueResources sets TotalUniqueResources field to given value.


### GetFilteredUniqueResources

`func (o *QueryCostDataResponse) GetFilteredUniqueResources() int32`

GetFilteredUniqueResources returns the FilteredUniqueResources field if non-nil, zero value otherwise.

### GetFilteredUniqueResourcesOk

`func (o *QueryCostDataResponse) GetFilteredUniqueResourcesOk() (*int32, bool)`

GetFilteredUniqueResourcesOk returns a tuple with the FilteredUniqueResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredUniqueResources

`func (o *QueryCostDataResponse) SetFilteredUniqueResources(v int32)`

SetFilteredUniqueResources sets FilteredUniqueResources field to given value.


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


