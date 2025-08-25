# AwsRestoreAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the role Eon assumes to access the account in AWS. | 
**OrganizationId** | Pointer to **string** | Cloud-provider-assigned ID of the AWS organization the account belongs to. | [optional] 

## Methods

### NewAwsRestoreAccountAttributes

`func NewAwsRestoreAccountAttributes(roleArn string, ) *AwsRestoreAccountAttributes`

NewAwsRestoreAccountAttributes instantiates a new AwsRestoreAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRestoreAccountAttributesWithDefaults

`func NewAwsRestoreAccountAttributesWithDefaults() *AwsRestoreAccountAttributes`

NewAwsRestoreAccountAttributesWithDefaults instantiates a new AwsRestoreAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *AwsRestoreAccountAttributes) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsRestoreAccountAttributes) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsRestoreAccountAttributes) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetOrganizationId

`func (o *AwsRestoreAccountAttributes) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AwsRestoreAccountAttributes) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AwsRestoreAccountAttributes) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AwsRestoreAccountAttributes) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


