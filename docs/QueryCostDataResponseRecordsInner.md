# QueryCostDataResponseRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimePeriod** | [**TimePeriod**](TimePeriod.md) |  | 
**TotalUniqueResources** | Pointer to **int32** | Total number of unique resources in the record | [optional] 
**Dimensions** | Pointer to [**QueryCostDataResponseRecordsInnerDimensions**](QueryCostDataResponseRecordsInnerDimensions.md) |  | [optional] 
**Costs** | [**[]QueryCostDataResponseRecordsInnerCostsInner**](QueryCostDataResponseRecordsInnerCostsInner.md) | Cost breakdown by metering dimension | 

## Methods

### NewQueryCostDataResponseRecordsInner

`func NewQueryCostDataResponseRecordsInner(timePeriod TimePeriod, costs []QueryCostDataResponseRecordsInnerCostsInner, ) *QueryCostDataResponseRecordsInner`

NewQueryCostDataResponseRecordsInner instantiates a new QueryCostDataResponseRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCostDataResponseRecordsInnerWithDefaults

`func NewQueryCostDataResponseRecordsInnerWithDefaults() *QueryCostDataResponseRecordsInner`

NewQueryCostDataResponseRecordsInnerWithDefaults instantiates a new QueryCostDataResponseRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimePeriod

`func (o *QueryCostDataResponseRecordsInner) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *QueryCostDataResponseRecordsInner) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *QueryCostDataResponseRecordsInner) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.


### GetTotalUniqueResources

`func (o *QueryCostDataResponseRecordsInner) GetTotalUniqueResources() int32`

GetTotalUniqueResources returns the TotalUniqueResources field if non-nil, zero value otherwise.

### GetTotalUniqueResourcesOk

`func (o *QueryCostDataResponseRecordsInner) GetTotalUniqueResourcesOk() (*int32, bool)`

GetTotalUniqueResourcesOk returns a tuple with the TotalUniqueResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUniqueResources

`func (o *QueryCostDataResponseRecordsInner) SetTotalUniqueResources(v int32)`

SetTotalUniqueResources sets TotalUniqueResources field to given value.

### HasTotalUniqueResources

`func (o *QueryCostDataResponseRecordsInner) HasTotalUniqueResources() bool`

HasTotalUniqueResources returns a boolean if a field has been set.

### GetDimensions

`func (o *QueryCostDataResponseRecordsInner) GetDimensions() QueryCostDataResponseRecordsInnerDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *QueryCostDataResponseRecordsInner) GetDimensionsOk() (*QueryCostDataResponseRecordsInnerDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *QueryCostDataResponseRecordsInner) SetDimensions(v QueryCostDataResponseRecordsInnerDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *QueryCostDataResponseRecordsInner) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCosts

`func (o *QueryCostDataResponseRecordsInner) GetCosts() []QueryCostDataResponseRecordsInnerCostsInner`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *QueryCostDataResponseRecordsInner) GetCostsOk() (*[]QueryCostDataResponseRecordsInnerCostsInner, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *QueryCostDataResponseRecordsInner) SetCosts(v []QueryCostDataResponseRecordsInnerCostsInner)`

SetCosts sets Costs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


