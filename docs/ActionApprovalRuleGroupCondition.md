# ActionApprovalRuleGroupCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**LogicalOperator**](LogicalOperator.md) |  | 
**Operands** | [**[]ActionApprovalRuleExpression**](ActionApprovalRuleExpression.md) | Must contain at least two expressions. | 

## Methods

### NewActionApprovalRuleGroupCondition

`func NewActionApprovalRuleGroupCondition(operator LogicalOperator, operands []ActionApprovalRuleExpression, ) *ActionApprovalRuleGroupCondition`

NewActionApprovalRuleGroupCondition instantiates a new ActionApprovalRuleGroupCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionApprovalRuleGroupConditionWithDefaults

`func NewActionApprovalRuleGroupConditionWithDefaults() *ActionApprovalRuleGroupCondition`

NewActionApprovalRuleGroupConditionWithDefaults instantiates a new ActionApprovalRuleGroupCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ActionApprovalRuleGroupCondition) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ActionApprovalRuleGroupCondition) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ActionApprovalRuleGroupCondition) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.


### GetOperands

`func (o *ActionApprovalRuleGroupCondition) GetOperands() []ActionApprovalRuleExpression`

GetOperands returns the Operands field if non-nil, zero value otherwise.

### GetOperandsOk

`func (o *ActionApprovalRuleGroupCondition) GetOperandsOk() (*[]ActionApprovalRuleExpression, bool)`

GetOperandsOk returns a tuple with the Operands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperands

`func (o *ActionApprovalRuleGroupCondition) SetOperands(v []ActionApprovalRuleExpression)`

SetOperands sets Operands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


