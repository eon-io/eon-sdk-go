# AwsAccountConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the role Eon assumes to access the account in AWS. | 

## Methods

### NewAwsAccountConfigInput

`func NewAwsAccountConfigInput(roleArn string, ) *AwsAccountConfigInput`

NewAwsAccountConfigInput instantiates a new AwsAccountConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAccountConfigInputWithDefaults

`func NewAwsAccountConfigInputWithDefaults() *AwsAccountConfigInput`

NewAwsAccountConfigInputWithDefaults instantiates a new AwsAccountConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsAccountConfigInput) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsAccountConfigInput) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsAccountConfigInput) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


