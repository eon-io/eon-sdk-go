# CostDataRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordTimeFrame** | [**TimeFrame**](TimeFrame.md) |  | 
**ResourceCount** | Pointer to **int32** | Total number of unique resources in the record | [optional] 
**Dimensions** | Pointer to [**CostDataRecordDimensions**](CostDataRecordDimensions.md) |  | [optional] 
**Costs** | [**[]CostDataRecordCost**](CostDataRecordCost.md) |  | 

## Methods

### NewCostDataRecord

`func NewCostDataRecord(recordTimeFrame TimeFrame, costs []CostDataRecordCost, ) *CostDataRecord`

NewCostDataRecord instantiates a new CostDataRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataRecordWithDefaults

`func NewCostDataRecordWithDefaults() *CostDataRecord`

NewCostDataRecordWithDefaults instantiates a new CostDataRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordTimeFrame

`func (o *CostDataRecord) GetRecordTimeFrame() TimeFrame`

GetRecordTimeFrame returns the RecordTimeFrame field if non-nil, zero value otherwise.

### GetRecordTimeFrameOk

`func (o *CostDataRecord) GetRecordTimeFrameOk() (*TimeFrame, bool)`

GetRecordTimeFrameOk returns a tuple with the RecordTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTimeFrame

`func (o *CostDataRecord) SetRecordTimeFrame(v TimeFrame)`

SetRecordTimeFrame sets RecordTimeFrame field to given value.


### GetResourceCount

`func (o *CostDataRecord) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *CostDataRecord) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *CostDataRecord) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *CostDataRecord) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.

### GetDimensions

`func (o *CostDataRecord) GetDimensions() CostDataRecordDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CostDataRecord) GetDimensionsOk() (*CostDataRecordDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CostDataRecord) SetDimensions(v CostDataRecordDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CostDataRecord) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCosts

`func (o *CostDataRecord) GetCosts() []CostDataRecordCost`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *CostDataRecord) GetCostsOk() (*[]CostDataRecordCost, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *CostDataRecord) SetCosts(v []CostDataRecordCost)`

SetCosts sets Costs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


