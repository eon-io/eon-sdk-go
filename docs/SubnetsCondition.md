# SubnetsCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**Subnets** | **[]string** |  | 

## Methods

### NewSubnetsCondition

`func NewSubnetsCondition(operator ListOperators, subnets []string, ) *SubnetsCondition`

NewSubnetsCondition instantiates a new SubnetsCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetsConditionWithDefaults

`func NewSubnetsConditionWithDefaults() *SubnetsCondition`

NewSubnetsConditionWithDefaults instantiates a new SubnetsCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SubnetsCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SubnetsCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SubnetsCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetSubnets

`func (o *SubnetsCondition) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *SubnetsCondition) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *SubnetsCondition) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


