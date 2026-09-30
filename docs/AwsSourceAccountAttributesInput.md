# AwsSourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the role Eon assumes to access the account in AWS. | 
**Regions** | Pointer to **[]string** | AWS regions Eon discovers in. Omit or leave empty to discover in all supported regions. | [optional] 

## Methods

### NewAwsSourceAccountAttributesInput

`func NewAwsSourceAccountAttributesInput(roleArn string, ) *AwsSourceAccountAttributesInput`

NewAwsSourceAccountAttributesInput instantiates a new AwsSourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSourceAccountAttributesInputWithDefaults

`func NewAwsSourceAccountAttributesInputWithDefaults() *AwsSourceAccountAttributesInput`

NewAwsSourceAccountAttributesInputWithDefaults instantiates a new AwsSourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsSourceAccountAttributesInput) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsSourceAccountAttributesInput) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsSourceAccountAttributesInput) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetRegions

`func (o *AwsSourceAccountAttributesInput) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AwsSourceAccountAttributesInput) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AwsSourceAccountAttributesInput) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AwsSourceAccountAttributesInput) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


