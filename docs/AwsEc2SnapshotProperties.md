# AwsEc2SnapshotProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | Instance type. | [optional] 
**SubnetId** | Pointer to **string** | ID of the subnet the instance is in. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | IDs of security groups associated with the resource. | [optional] 
**InstanceProfileName** | Pointer to **string** | The resource&#39;s instance profile. | [optional] 
**Volumes** | Pointer to [**[]InventorySnapshotVolume**](InventorySnapshotVolume.md) | List of volumes attached to the resource. | [optional] 

## Methods

### NewAwsEc2SnapshotProperties

`func NewAwsEc2SnapshotProperties() *AwsEc2SnapshotProperties`

NewAwsEc2SnapshotProperties instantiates a new AwsEc2SnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEc2SnapshotPropertiesWithDefaults

`func NewAwsEc2SnapshotPropertiesWithDefaults() *AwsEc2SnapshotProperties`

NewAwsEc2SnapshotPropertiesWithDefaults instantiates a new AwsEc2SnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *AwsEc2SnapshotProperties) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AwsEc2SnapshotProperties) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AwsEc2SnapshotProperties) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AwsEc2SnapshotProperties) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSubnetId

`func (o *AwsEc2SnapshotProperties) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsEc2SnapshotProperties) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsEc2SnapshotProperties) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsEc2SnapshotProperties) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSecurityGroupIds

`func (o *AwsEc2SnapshotProperties) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AwsEc2SnapshotProperties) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AwsEc2SnapshotProperties) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AwsEc2SnapshotProperties) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetInstanceProfileName

`func (o *AwsEc2SnapshotProperties) GetInstanceProfileName() string`

GetInstanceProfileName returns the InstanceProfileName field if non-nil, zero value otherwise.

### GetInstanceProfileNameOk

`func (o *AwsEc2SnapshotProperties) GetInstanceProfileNameOk() (*string, bool)`

GetInstanceProfileNameOk returns a tuple with the InstanceProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfileName

`func (o *AwsEc2SnapshotProperties) SetInstanceProfileName(v string)`

SetInstanceProfileName sets InstanceProfileName field to given value.

### HasInstanceProfileName

`func (o *AwsEc2SnapshotProperties) HasInstanceProfileName() bool`

HasInstanceProfileName returns a boolean if a field has been set.

### GetVolumes

`func (o *AwsEc2SnapshotProperties) GetVolumes() []InventorySnapshotVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AwsEc2SnapshotProperties) GetVolumesOk() (*[]InventorySnapshotVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AwsEc2SnapshotProperties) SetVolumes(v []InventorySnapshotVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AwsEc2SnapshotProperties) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


