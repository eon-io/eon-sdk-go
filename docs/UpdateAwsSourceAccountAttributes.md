# UpdateAwsSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **NullableString** | ARN of the role Eon assumes to access the account in AWS. Only the role name portion of the ARN can be changed; the AWS account ID must remain the same. To use a different AWS account, disconnect and connect a new source account.  | [optional] 

## Methods

### NewUpdateAwsSourceAccountAttributes

`func NewUpdateAwsSourceAccountAttributes() *UpdateAwsSourceAccountAttributes`

NewUpdateAwsSourceAccountAttributes instantiates a new UpdateAwsSourceAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsSourceAccountAttributesWithDefaults

`func NewUpdateAwsSourceAccountAttributesWithDefaults() *UpdateAwsSourceAccountAttributes`

NewUpdateAwsSourceAccountAttributesWithDefaults instantiates a new UpdateAwsSourceAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *UpdateAwsSourceAccountAttributes) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *UpdateAwsSourceAccountAttributes) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *UpdateAwsSourceAccountAttributes) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *UpdateAwsSourceAccountAttributes) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *UpdateAwsSourceAccountAttributes) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *UpdateAwsSourceAccountAttributes) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


