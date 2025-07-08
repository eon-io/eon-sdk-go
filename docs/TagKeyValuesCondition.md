# TagKeyValuesCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**TagKeyValues** | [**[]TagKeyValue**](TagKeyValue.md) |  | 

## Methods

### NewTagKeyValuesCondition

`func NewTagKeyValuesCondition(operator ListOperators, tagKeyValues []TagKeyValue, ) *TagKeyValuesCondition`

NewTagKeyValuesCondition instantiates a new TagKeyValuesCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagKeyValuesConditionWithDefaults

`func NewTagKeyValuesConditionWithDefaults() *TagKeyValuesCondition`

NewTagKeyValuesConditionWithDefaults instantiates a new TagKeyValuesCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *TagKeyValuesCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TagKeyValuesCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TagKeyValuesCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetTagKeyValues

`func (o *TagKeyValuesCondition) GetTagKeyValues() []TagKeyValue`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *TagKeyValuesCondition) GetTagKeyValuesOk() (*[]TagKeyValue, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *TagKeyValuesCondition) SetTagKeyValues(v []TagKeyValue)`

SetTagKeyValues sets TagKeyValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


