# EnvironmentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**Environments** | [**[]Environment**](Environment.md) |  | 

## Methods

### NewEnvironmentCondition

`func NewEnvironmentCondition(operator ScalarOperators, environments []Environment, ) *EnvironmentCondition`

NewEnvironmentCondition instantiates a new EnvironmentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentConditionWithDefaults

`func NewEnvironmentConditionWithDefaults() *EnvironmentCondition`

NewEnvironmentConditionWithDefaults instantiates a new EnvironmentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *EnvironmentCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EnvironmentCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EnvironmentCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetEnvironments

`func (o *EnvironmentCondition) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EnvironmentCondition) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EnvironmentCondition) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


