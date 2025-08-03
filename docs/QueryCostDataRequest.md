# QueryCostDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostUnit** | Pointer to [**CostUnit**](CostUnit.md) |  | [optional] [default to CREDITS]
**TimePeriod** | [**TimePeriod**](TimePeriod.md) |  | 
**Granularity** | Pointer to **string** | Granularity for cost aggregation | [optional] [default to "MONTHLY"]
**Filters** | Pointer to [**CostDataFilters**](CostDataFilters.md) |  | [optional] 
**GroupBy** | Pointer to **string** | Group results by specified dimensions | [optional] 

## Methods

### NewQueryCostDataRequest

`func NewQueryCostDataRequest(timePeriod TimePeriod, ) *QueryCostDataRequest`

NewQueryCostDataRequest instantiates a new QueryCostDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCostDataRequestWithDefaults

`func NewQueryCostDataRequestWithDefaults() *QueryCostDataRequest`

NewQueryCostDataRequestWithDefaults instantiates a new QueryCostDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostUnit

`func (o *QueryCostDataRequest) GetCostUnit() CostUnit`

GetCostUnit returns the CostUnit field if non-nil, zero value otherwise.

### GetCostUnitOk

`func (o *QueryCostDataRequest) GetCostUnitOk() (*CostUnit, bool)`

GetCostUnitOk returns a tuple with the CostUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUnit

`func (o *QueryCostDataRequest) SetCostUnit(v CostUnit)`

SetCostUnit sets CostUnit field to given value.

### HasCostUnit

`func (o *QueryCostDataRequest) HasCostUnit() bool`

HasCostUnit returns a boolean if a field has been set.

### GetTimePeriod

`func (o *QueryCostDataRequest) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *QueryCostDataRequest) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *QueryCostDataRequest) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.


### GetGranularity

`func (o *QueryCostDataRequest) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *QueryCostDataRequest) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *QueryCostDataRequest) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *QueryCostDataRequest) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetFilters

`func (o *QueryCostDataRequest) GetFilters() CostDataFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QueryCostDataRequest) GetFiltersOk() (*CostDataFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QueryCostDataRequest) SetFilters(v CostDataFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QueryCostDataRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *QueryCostDataRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *QueryCostDataRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *QueryCostDataRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *QueryCostDataRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


