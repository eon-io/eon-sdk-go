# BackupPolicyGroupCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**LogicalOperator**](LogicalOperator.md) |  | 
**Operands** | [**[]BackupPolicyExpression**](BackupPolicyExpression.md) | List of expressions to evaluate using the logical operator specified in &#x60;operator&#x60;. Each item in the list can be either a grouped expression of type &#x60;BackupPolicyLogicalExpression&#x60; or a basic selector of type &#x60;BackupPolicyBasicSelector&#x60;.  Must contain at least 2 items.  | 

## Methods

### NewBackupPolicyGroupCondition

`func NewBackupPolicyGroupCondition(operator LogicalOperator, operands []BackupPolicyExpression, ) *BackupPolicyGroupCondition`

NewBackupPolicyGroupCondition instantiates a new BackupPolicyGroupCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyGroupConditionWithDefaults

`func NewBackupPolicyGroupConditionWithDefaults() *BackupPolicyGroupCondition`

NewBackupPolicyGroupConditionWithDefaults instantiates a new BackupPolicyGroupCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *BackupPolicyGroupCondition) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BackupPolicyGroupCondition) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BackupPolicyGroupCondition) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.


### GetOperands

`func (o *BackupPolicyGroupCondition) GetOperands() []BackupPolicyExpression`

GetOperands returns the Operands field if non-nil, zero value otherwise.

### GetOperandsOk

`func (o *BackupPolicyGroupCondition) GetOperandsOk() (*[]BackupPolicyExpression, bool)`

GetOperandsOk returns a tuple with the Operands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperands

`func (o *BackupPolicyGroupCondition) SetOperands(v []BackupPolicyExpression)`

SetOperands sets Operands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


