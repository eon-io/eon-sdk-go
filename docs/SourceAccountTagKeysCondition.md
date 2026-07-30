# SourceAccountTagKeysCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**SourceAccountTagKeys** | **[]string** |  | 

## Methods

### NewSourceAccountTagKeysCondition

`func NewSourceAccountTagKeysCondition(operator ListOperators, sourceAccountTagKeys []string, ) *SourceAccountTagKeysCondition`

NewSourceAccountTagKeysCondition instantiates a new SourceAccountTagKeysCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountTagKeysConditionWithDefaults

`func NewSourceAccountTagKeysConditionWithDefaults() *SourceAccountTagKeysCondition`

NewSourceAccountTagKeysConditionWithDefaults instantiates a new SourceAccountTagKeysCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SourceAccountTagKeysCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SourceAccountTagKeysCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SourceAccountTagKeysCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetSourceAccountTagKeys

`func (o *SourceAccountTagKeysCondition) GetSourceAccountTagKeys() []string`

GetSourceAccountTagKeys returns the SourceAccountTagKeys field if non-nil, zero value otherwise.

### GetSourceAccountTagKeysOk

`func (o *SourceAccountTagKeysCondition) GetSourceAccountTagKeysOk() (*[]string, bool)`

GetSourceAccountTagKeysOk returns a tuple with the SourceAccountTagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountTagKeys

`func (o *SourceAccountTagKeysCondition) SetSourceAccountTagKeys(v []string)`

SetSourceAccountTagKeys sets SourceAccountTagKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


