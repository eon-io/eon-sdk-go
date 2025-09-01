# AwsEc2InstanceRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to restore the instance to. | 
**InstanceType** | **string** | Instance type to use for the restored instance. | 
**SubnetId** | **string** | Subnet ID to associate with the restored instance. | 
**SecurityGroupIds** | Pointer to **[]string** | List of security group IDs to associate with the restored instance. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**VolumeRestoreParameters** | [**[]RestoreInstanceVolumeInput**](RestoreInstanceVolumeInput.md) | Volumes to restore and attach to the restored instance. Each item in the list corresponds to a volume to be restored, where &#x60;providerVolumeId&#x60; matches the volume&#39;s ID at the time of the snapshot. The root volume must be present in the list.  | 

## Methods

### NewAwsEc2InstanceRestoreTarget

`func NewAwsEc2InstanceRestoreTarget(region string, instanceType string, subnetId string, volumeRestoreParameters []RestoreInstanceVolumeInput, ) *AwsEc2InstanceRestoreTarget`

NewAwsEc2InstanceRestoreTarget instantiates a new AwsEc2InstanceRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEc2InstanceRestoreTargetWithDefaults

`func NewAwsEc2InstanceRestoreTargetWithDefaults() *AwsEc2InstanceRestoreTarget`

NewAwsEc2InstanceRestoreTargetWithDefaults instantiates a new AwsEc2InstanceRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsEc2InstanceRestoreTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsEc2InstanceRestoreTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsEc2InstanceRestoreTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstanceType

`func (o *AwsEc2InstanceRestoreTarget) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AwsEc2InstanceRestoreTarget) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AwsEc2InstanceRestoreTarget) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetSubnetId

`func (o *AwsEc2InstanceRestoreTarget) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsEc2InstanceRestoreTarget) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsEc2InstanceRestoreTarget) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSecurityGroupIds

`func (o *AwsEc2InstanceRestoreTarget) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AwsEc2InstanceRestoreTarget) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AwsEc2InstanceRestoreTarget) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AwsEc2InstanceRestoreTarget) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetTags

`func (o *AwsEc2InstanceRestoreTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEc2InstanceRestoreTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEc2InstanceRestoreTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEc2InstanceRestoreTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeRestoreParameters

`func (o *AwsEc2InstanceRestoreTarget) GetVolumeRestoreParameters() []RestoreInstanceVolumeInput`

GetVolumeRestoreParameters returns the VolumeRestoreParameters field if non-nil, zero value otherwise.

### GetVolumeRestoreParametersOk

`func (o *AwsEc2InstanceRestoreTarget) GetVolumeRestoreParametersOk() (*[]RestoreInstanceVolumeInput, bool)`

GetVolumeRestoreParametersOk returns a tuple with the VolumeRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRestoreParameters

`func (o *AwsEc2InstanceRestoreTarget) SetVolumeRestoreParameters(v []RestoreInstanceVolumeInput)`

SetVolumeRestoreParameters sets VolumeRestoreParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


