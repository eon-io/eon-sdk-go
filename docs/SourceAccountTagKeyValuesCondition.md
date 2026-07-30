# SourceAccountTagKeyValuesCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**SourceAccountTagKeyValues** | [**[]TagKeyValue**](TagKeyValue.md) |  | 

## Methods

### NewSourceAccountTagKeyValuesCondition

`func NewSourceAccountTagKeyValuesCondition(operator ListOperators, sourceAccountTagKeyValues []TagKeyValue, ) *SourceAccountTagKeyValuesCondition`

NewSourceAccountTagKeyValuesCondition instantiates a new SourceAccountTagKeyValuesCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountTagKeyValuesConditionWithDefaults

`func NewSourceAccountTagKeyValuesConditionWithDefaults() *SourceAccountTagKeyValuesCondition`

NewSourceAccountTagKeyValuesConditionWithDefaults instantiates a new SourceAccountTagKeyValuesCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SourceAccountTagKeyValuesCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SourceAccountTagKeyValuesCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SourceAccountTagKeyValuesCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetSourceAccountTagKeyValues

`func (o *SourceAccountTagKeyValuesCondition) GetSourceAccountTagKeyValues() []TagKeyValue`

GetSourceAccountTagKeyValues returns the SourceAccountTagKeyValues field if non-nil, zero value otherwise.

### GetSourceAccountTagKeyValuesOk

`func (o *SourceAccountTagKeyValuesCondition) GetSourceAccountTagKeyValuesOk() (*[]TagKeyValue, bool)`

GetSourceAccountTagKeyValuesOk returns a tuple with the SourceAccountTagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountTagKeyValues

`func (o *SourceAccountTagKeyValuesCondition) SetSourceAccountTagKeyValues(v []TagKeyValue)`

SetSourceAccountTagKeyValues sets SourceAccountTagKeyValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


