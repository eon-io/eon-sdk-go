# UpdateAwsSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | Pointer to **NullableString** | ARN of the role Eon assumes to access the account in AWS. Only the role name portion of the ARN can be changed. The AWS account ID must remain the same.  | [optional] 
**Regions** | Pointer to **[]string** | AWS regions Eon discovers in. Omit to leave the current regions unchanged. Pass an empty list to discover in all supported regions.  | [optional] 

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
### GetRegions

`func (o *UpdateAwsSourceAccountAttributes) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *UpdateAwsSourceAccountAttributes) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *UpdateAwsSourceAccountAttributes) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *UpdateAwsSourceAccountAttributes) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


