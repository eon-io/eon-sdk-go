# RoleAccessGroupCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**LogicalOperator**](LogicalOperator.md) |  | 
**Operands** | [**[]AccessConditionalExpression**](AccessConditionalExpression.md) | List of expressions to evaluate using the logical operator specified in &#x60;operator&#x60;. Each item in the list can be either a single expression or a nested &#x60;group&#x60; expression.  Must contain at least 2 items.  | 

## Methods

### NewRoleAccessGroupCondition

`func NewRoleAccessGroupCondition(operator LogicalOperator, operands []AccessConditionalExpression, ) *RoleAccessGroupCondition`

NewRoleAccessGroupCondition instantiates a new RoleAccessGroupCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAccessGroupConditionWithDefaults

`func NewRoleAccessGroupConditionWithDefaults() *RoleAccessGroupCondition`

NewRoleAccessGroupConditionWithDefaults instantiates a new RoleAccessGroupCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *RoleAccessGroupCondition) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RoleAccessGroupCondition) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RoleAccessGroupCondition) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.


### GetOperands

`func (o *RoleAccessGroupCondition) GetOperands() []AccessConditionalExpression`

GetOperands returns the Operands field if non-nil, zero value otherwise.

### GetOperandsOk

`func (o *RoleAccessGroupCondition) GetOperandsOk() (*[]AccessConditionalExpression, bool)`

GetOperandsOk returns a tuple with the Operands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperands

`func (o *RoleAccessGroupCondition) SetOperands(v []AccessConditionalExpression)`

SetOperands sets Operands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


