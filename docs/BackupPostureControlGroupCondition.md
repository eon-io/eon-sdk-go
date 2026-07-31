# BackupPostureControlGroupCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**LogicalOperator**](LogicalOperator.md) |  | 
**Operands** | [**[]BackupPostureControlExpression**](BackupPostureControlExpression.md) | The expressions to combine with &#x60;operator&#x60;. | 

## Methods

### NewBackupPostureControlGroupCondition

`func NewBackupPostureControlGroupCondition(operator LogicalOperator, operands []BackupPostureControlExpression, ) *BackupPostureControlGroupCondition`

NewBackupPostureControlGroupCondition instantiates a new BackupPostureControlGroupCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPostureControlGroupConditionWithDefaults

`func NewBackupPostureControlGroupConditionWithDefaults() *BackupPostureControlGroupCondition`

NewBackupPostureControlGroupConditionWithDefaults instantiates a new BackupPostureControlGroupCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *BackupPostureControlGroupCondition) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BackupPostureControlGroupCondition) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BackupPostureControlGroupCondition) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.


### GetOperands

`func (o *BackupPostureControlGroupCondition) GetOperands() []BackupPostureControlExpression`

GetOperands returns the Operands field if non-nil, zero value otherwise.

### GetOperandsOk

`func (o *BackupPostureControlGroupCondition) GetOperandsOk() (*[]BackupPostureControlExpression, bool)`

GetOperandsOk returns a tuple with the Operands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperands

`func (o *BackupPostureControlGroupCondition) SetOperands(v []BackupPostureControlExpression)`

SetOperands sets Operands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


