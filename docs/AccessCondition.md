# AccessCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Display name for the data access rule, which also serves as its ID for use in &#x60;permissionGrants&#x60;. Valid characters: &#x60;[a-zA-Z0-9-_.\\s]&#x60;.  | 
**Effect** | [**AccessConditionEffect**](AccessConditionEffect.md) |  | 
**Expression** | [**NullableAccessConditionalExpression**](AccessConditionalExpression.md) |  | 

## Methods

### NewAccessCondition

`func NewAccessCondition(id string, effect AccessConditionEffect, expression NullableAccessConditionalExpression, ) *AccessCondition`

NewAccessCondition instantiates a new AccessCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessConditionWithDefaults

`func NewAccessConditionWithDefaults() *AccessCondition`

NewAccessConditionWithDefaults instantiates a new AccessCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessCondition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessCondition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessCondition) SetId(v string)`

SetId sets Id field to given value.


### GetEffect

`func (o *AccessCondition) GetEffect() AccessConditionEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AccessCondition) GetEffectOk() (*AccessConditionEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AccessCondition) SetEffect(v AccessConditionEffect)`

SetEffect sets Effect field to given value.


### GetExpression

`func (o *AccessCondition) GetExpression() AccessConditionalExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *AccessCondition) GetExpressionOk() (*AccessConditionalExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *AccessCondition) SetExpression(v AccessConditionalExpression)`

SetExpression sets Expression field to given value.


### SetExpressionNil

`func (o *AccessCondition) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *AccessCondition) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


