# Cost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | Cost amount | 
**Unit** | [**CostUnit**](CostUnit.md) |  | [default to COST_UNIT_CREDITS]

## Methods

### NewCost

`func NewCost(amount float64, unit CostUnit, ) *Cost`

NewCost instantiates a new Cost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostWithDefaults

`func NewCostWithDefaults() *Cost`

NewCostWithDefaults instantiates a new Cost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Cost) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Cost) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Cost) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *Cost) GetUnit() CostUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Cost) GetUnitOk() (*CostUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Cost) SetUnit(v CostUnit)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


