# TagKeysCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**TagKeys** | **[]string** |  | 

## Methods

### NewTagKeysCondition

`func NewTagKeysCondition(operator ListOperators, tagKeys []string, ) *TagKeysCondition`

NewTagKeysCondition instantiates a new TagKeysCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagKeysConditionWithDefaults

`func NewTagKeysConditionWithDefaults() *TagKeysCondition`

NewTagKeysConditionWithDefaults instantiates a new TagKeysCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *TagKeysCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TagKeysCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TagKeysCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetTagKeys

`func (o *TagKeysCondition) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *TagKeysCondition) GetTagKeysOk() (*[]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *TagKeysCondition) SetTagKeys(v []string)`

SetTagKeys sets TagKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


