# CostDataRecordCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeteringDimension** | [**MeteringDimension**](MeteringDimension.md) |  | 
**Cost** | [**Cost**](Cost.md) |  | 
**Usage** | [**Usage**](Usage.md) |  | 

## Methods

### NewCostDataRecordCost

`func NewCostDataRecordCost(meteringDimension MeteringDimension, cost Cost, usage Usage, ) *CostDataRecordCost`

NewCostDataRecordCost instantiates a new CostDataRecordCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataRecordCostWithDefaults

`func NewCostDataRecordCostWithDefaults() *CostDataRecordCost`

NewCostDataRecordCostWithDefaults instantiates a new CostDataRecordCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeteringDimension

`func (o *CostDataRecordCost) GetMeteringDimension() MeteringDimension`

GetMeteringDimension returns the MeteringDimension field if non-nil, zero value otherwise.

### GetMeteringDimensionOk

`func (o *CostDataRecordCost) GetMeteringDimensionOk() (*MeteringDimension, bool)`

GetMeteringDimensionOk returns a tuple with the MeteringDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringDimension

`func (o *CostDataRecordCost) SetMeteringDimension(v MeteringDimension)`

SetMeteringDimension sets MeteringDimension field to given value.


### GetCost

`func (o *CostDataRecordCost) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CostDataRecordCost) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CostDataRecordCost) SetCost(v Cost)`

SetCost sets Cost field to given value.


### GetUsage

`func (o *CostDataRecordCost) GetUsage() Usage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CostDataRecordCost) GetUsageOk() (*Usage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CostDataRecordCost) SetUsage(v Usage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


