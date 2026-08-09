# SensitivityAnnotationsCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**SensitivityAnnotations** | [**[]SensitivityLevel**](SensitivityLevel.md) |  | 

## Methods

### NewSensitivityAnnotationsCondition

`func NewSensitivityAnnotationsCondition(operator ScalarOperators, sensitivityAnnotations []SensitivityLevel, ) *SensitivityAnnotationsCondition`

NewSensitivityAnnotationsCondition instantiates a new SensitivityAnnotationsCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensitivityAnnotationsConditionWithDefaults

`func NewSensitivityAnnotationsConditionWithDefaults() *SensitivityAnnotationsCondition`

NewSensitivityAnnotationsConditionWithDefaults instantiates a new SensitivityAnnotationsCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SensitivityAnnotationsCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SensitivityAnnotationsCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SensitivityAnnotationsCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetSensitivityAnnotations

`func (o *SensitivityAnnotationsCondition) GetSensitivityAnnotations() []SensitivityLevel`

GetSensitivityAnnotations returns the SensitivityAnnotations field if non-nil, zero value otherwise.

### GetSensitivityAnnotationsOk

`func (o *SensitivityAnnotationsCondition) GetSensitivityAnnotationsOk() (*[]SensitivityLevel, bool)`

GetSensitivityAnnotationsOk returns a tuple with the SensitivityAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivityAnnotations

`func (o *SensitivityAnnotationsCondition) SetSensitivityAnnotations(v []SensitivityLevel)`

SetSensitivityAnnotations sets SensitivityAnnotations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


