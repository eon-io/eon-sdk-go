# GlobalClusterIdentifierCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**GlobalClusterIdentifier** | **[]string** |  | 

## Methods

### NewGlobalClusterIdentifierCondition

`func NewGlobalClusterIdentifierCondition(operator ScalarOperators, globalClusterIdentifier []string, ) *GlobalClusterIdentifierCondition`

NewGlobalClusterIdentifierCondition instantiates a new GlobalClusterIdentifierCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalClusterIdentifierConditionWithDefaults

`func NewGlobalClusterIdentifierConditionWithDefaults() *GlobalClusterIdentifierCondition`

NewGlobalClusterIdentifierConditionWithDefaults instantiates a new GlobalClusterIdentifierCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *GlobalClusterIdentifierCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GlobalClusterIdentifierCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GlobalClusterIdentifierCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetGlobalClusterIdentifier

`func (o *GlobalClusterIdentifierCondition) GetGlobalClusterIdentifier() []string`

GetGlobalClusterIdentifier returns the GlobalClusterIdentifier field if non-nil, zero value otherwise.

### GetGlobalClusterIdentifierOk

`func (o *GlobalClusterIdentifierCondition) GetGlobalClusterIdentifierOk() (*[]string, bool)`

GetGlobalClusterIdentifierOk returns a tuple with the GlobalClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterIdentifier

`func (o *GlobalClusterIdentifierCondition) SetGlobalClusterIdentifier(v []string)`

SetGlobalClusterIdentifier sets GlobalClusterIdentifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


