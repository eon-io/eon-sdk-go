# UpdateAwsRestoreAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **NullableString** | ARN of the role Eon assumes to access the account in AWS. Only the role name portion of the ARN can be changed; the AWS account ID must remain the same. To use a different AWS account, disconnect and connect a new restore account.  | [optional] 

## Methods

### NewUpdateAwsRestoreAccountAttributes

`func NewUpdateAwsRestoreAccountAttributes() *UpdateAwsRestoreAccountAttributes`

NewUpdateAwsRestoreAccountAttributes instantiates a new UpdateAwsRestoreAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsRestoreAccountAttributesWithDefaults

`func NewUpdateAwsRestoreAccountAttributesWithDefaults() *UpdateAwsRestoreAccountAttributes`

NewUpdateAwsRestoreAccountAttributesWithDefaults instantiates a new UpdateAwsRestoreAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *UpdateAwsRestoreAccountAttributes) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *UpdateAwsRestoreAccountAttributes) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *UpdateAwsRestoreAccountAttributes) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *UpdateAwsRestoreAccountAttributes) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *UpdateAwsRestoreAccountAttributes) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *UpdateAwsRestoreAccountAttributes) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


