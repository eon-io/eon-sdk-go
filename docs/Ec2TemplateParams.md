# Ec2TemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetId** | **string** | Cloud-provider-assigned ID of the subnet for the restored EC2 instance. | 
**SecurityGroupIds** | Pointer to **[]string** | Cloud-provider-assigned IDs of the security groups to attach to the restored EC2 instance. | [optional] 
**InstanceType** | Pointer to **string** | EC2 instance type for the restored instance. Required unless &#x60;inheritInstanceType&#x60; is set to &#x60;true&#x60;, in which case, must be empty.  | [optional] 
**InheritInstanceType** | Pointer to **bool** | Whether to restore each instance with the instance type recorded in its own snapshot, instead of &#x60;instanceType&#x60;. The template&#39;s subnet must be in an availability zone that offers the specified type.  | [optional] [default to false]
**KmsKeyArn** | Pointer to **string** | ARN of the KMS key to use for encrypting the restored instance. | [optional] 
**KeepOriginalTags** | **bool** | Whether to keep the original tags from the source resource. | [default to true]
**AdditionalTags** | Pointer to [**[]Ec2TemplateParamsAdditionalTagsInner**](Ec2TemplateParamsAdditionalTagsInner.md) | Additional tags to apply to the restored EC2 instance. | [optional] 
**NameTemplate** | **string** | Template for naming the restored EC2 instance. | 
**StartAfterRestore** | **bool** | Whether to start the restored EC2 instance after restoration. | 
**TargetAccountId** | **string** | Eon-assigned ID of the destination restore account. | 
**TargetRegion** | **string** | Target region for the restore. | 

## Methods

### NewEc2TemplateParams

`func NewEc2TemplateParams(subnetId string, keepOriginalTags bool, nameTemplate string, startAfterRestore bool, targetAccountId string, targetRegion string, ) *Ec2TemplateParams`

NewEc2TemplateParams instantiates a new Ec2TemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2TemplateParamsWithDefaults

`func NewEc2TemplateParamsWithDefaults() *Ec2TemplateParams`

NewEc2TemplateParamsWithDefaults instantiates a new Ec2TemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetId

`func (o *Ec2TemplateParams) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Ec2TemplateParams) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Ec2TemplateParams) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSecurityGroupIds

`func (o *Ec2TemplateParams) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *Ec2TemplateParams) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *Ec2TemplateParams) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *Ec2TemplateParams) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetInstanceType

`func (o *Ec2TemplateParams) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Ec2TemplateParams) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Ec2TemplateParams) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Ec2TemplateParams) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInheritInstanceType

`func (o *Ec2TemplateParams) GetInheritInstanceType() bool`

GetInheritInstanceType returns the InheritInstanceType field if non-nil, zero value otherwise.

### GetInheritInstanceTypeOk

`func (o *Ec2TemplateParams) GetInheritInstanceTypeOk() (*bool, bool)`

GetInheritInstanceTypeOk returns a tuple with the InheritInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritInstanceType

`func (o *Ec2TemplateParams) SetInheritInstanceType(v bool)`

SetInheritInstanceType sets InheritInstanceType field to given value.

### HasInheritInstanceType

`func (o *Ec2TemplateParams) HasInheritInstanceType() bool`

HasInheritInstanceType returns a boolean if a field has been set.

### GetKmsKeyArn

`func (o *Ec2TemplateParams) GetKmsKeyArn() string`

GetKmsKeyArn returns the KmsKeyArn field if non-nil, zero value otherwise.

### GetKmsKeyArnOk

`func (o *Ec2TemplateParams) GetKmsKeyArnOk() (*string, bool)`

GetKmsKeyArnOk returns a tuple with the KmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyArn

`func (o *Ec2TemplateParams) SetKmsKeyArn(v string)`

SetKmsKeyArn sets KmsKeyArn field to given value.

### HasKmsKeyArn

`func (o *Ec2TemplateParams) HasKmsKeyArn() bool`

HasKmsKeyArn returns a boolean if a field has been set.

### GetKeepOriginalTags

`func (o *Ec2TemplateParams) GetKeepOriginalTags() bool`

GetKeepOriginalTags returns the KeepOriginalTags field if non-nil, zero value otherwise.

### GetKeepOriginalTagsOk

`func (o *Ec2TemplateParams) GetKeepOriginalTagsOk() (*bool, bool)`

GetKeepOriginalTagsOk returns a tuple with the KeepOriginalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOriginalTags

`func (o *Ec2TemplateParams) SetKeepOriginalTags(v bool)`

SetKeepOriginalTags sets KeepOriginalTags field to given value.


### GetAdditionalTags

`func (o *Ec2TemplateParams) GetAdditionalTags() []Ec2TemplateParamsAdditionalTagsInner`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *Ec2TemplateParams) GetAdditionalTagsOk() (*[]Ec2TemplateParamsAdditionalTagsInner, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *Ec2TemplateParams) SetAdditionalTags(v []Ec2TemplateParamsAdditionalTagsInner)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *Ec2TemplateParams) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetNameTemplate

`func (o *Ec2TemplateParams) GetNameTemplate() string`

GetNameTemplate returns the NameTemplate field if non-nil, zero value otherwise.

### GetNameTemplateOk

`func (o *Ec2TemplateParams) GetNameTemplateOk() (*string, bool)`

GetNameTemplateOk returns a tuple with the NameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTemplate

`func (o *Ec2TemplateParams) SetNameTemplate(v string)`

SetNameTemplate sets NameTemplate field to given value.


### GetStartAfterRestore

`func (o *Ec2TemplateParams) GetStartAfterRestore() bool`

GetStartAfterRestore returns the StartAfterRestore field if non-nil, zero value otherwise.

### GetStartAfterRestoreOk

`func (o *Ec2TemplateParams) GetStartAfterRestoreOk() (*bool, bool)`

GetStartAfterRestoreOk returns a tuple with the StartAfterRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAfterRestore

`func (o *Ec2TemplateParams) SetStartAfterRestore(v bool)`

SetStartAfterRestore sets StartAfterRestore field to given value.


### GetTargetAccountId

`func (o *Ec2TemplateParams) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *Ec2TemplateParams) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *Ec2TemplateParams) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.


### GetTargetRegion

`func (o *Ec2TemplateParams) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *Ec2TemplateParams) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *Ec2TemplateParams) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


