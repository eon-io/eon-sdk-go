# Ec2InstanceRestoreTarget

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

### NewEc2InstanceRestoreTarget

`func NewEc2InstanceRestoreTarget(region string, instanceType string, subnetId string, volumeRestoreParameters []RestoreInstanceVolumeInput, ) *Ec2InstanceRestoreTarget`

NewEc2InstanceRestoreTarget instantiates a new Ec2InstanceRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2InstanceRestoreTargetWithDefaults

`func NewEc2InstanceRestoreTargetWithDefaults() *Ec2InstanceRestoreTarget`

NewEc2InstanceRestoreTargetWithDefaults instantiates a new Ec2InstanceRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Ec2InstanceRestoreTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Ec2InstanceRestoreTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Ec2InstanceRestoreTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstanceType

`func (o *Ec2InstanceRestoreTarget) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Ec2InstanceRestoreTarget) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Ec2InstanceRestoreTarget) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetSubnetId

`func (o *Ec2InstanceRestoreTarget) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Ec2InstanceRestoreTarget) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Ec2InstanceRestoreTarget) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSecurityGroupIds

`func (o *Ec2InstanceRestoreTarget) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *Ec2InstanceRestoreTarget) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *Ec2InstanceRestoreTarget) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *Ec2InstanceRestoreTarget) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetTags

`func (o *Ec2InstanceRestoreTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Ec2InstanceRestoreTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Ec2InstanceRestoreTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Ec2InstanceRestoreTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeRestoreParameters

`func (o *Ec2InstanceRestoreTarget) GetVolumeRestoreParameters() []RestoreInstanceVolumeInput`

GetVolumeRestoreParameters returns the VolumeRestoreParameters field if non-nil, zero value otherwise.

### GetVolumeRestoreParametersOk

`func (o *Ec2InstanceRestoreTarget) GetVolumeRestoreParametersOk() (*[]RestoreInstanceVolumeInput, bool)`

GetVolumeRestoreParametersOk returns a tuple with the VolumeRestoreParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRestoreParameters

`func (o *Ec2InstanceRestoreTarget) SetVolumeRestoreParameters(v []RestoreInstanceVolumeInput)`

SetVolumeRestoreParameters sets VolumeRestoreParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


