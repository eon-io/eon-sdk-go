# SecurityScanConclusionCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**SecurityScanConclusions** | [**[]SecurityScanConclusion**](SecurityScanConclusion.md) |  | 

## Methods

### NewSecurityScanConclusionCondition

`func NewSecurityScanConclusionCondition(operator ListOperators, securityScanConclusions []SecurityScanConclusion, ) *SecurityScanConclusionCondition`

NewSecurityScanConclusionCondition instantiates a new SecurityScanConclusionCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanConclusionConditionWithDefaults

`func NewSecurityScanConclusionConditionWithDefaults() *SecurityScanConclusionCondition`

NewSecurityScanConclusionConditionWithDefaults instantiates a new SecurityScanConclusionCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SecurityScanConclusionCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SecurityScanConclusionCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SecurityScanConclusionCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetSecurityScanConclusions

`func (o *SecurityScanConclusionCondition) GetSecurityScanConclusions() []SecurityScanConclusion`

GetSecurityScanConclusions returns the SecurityScanConclusions field if non-nil, zero value otherwise.

### GetSecurityScanConclusionsOk

`func (o *SecurityScanConclusionCondition) GetSecurityScanConclusionsOk() (*[]SecurityScanConclusion, bool)`

GetSecurityScanConclusionsOk returns a tuple with the SecurityScanConclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanConclusions

`func (o *SecurityScanConclusionCondition) SetSecurityScanConclusions(v []SecurityScanConclusion)`

SetSecurityScanConclusions sets SecurityScanConclusions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


