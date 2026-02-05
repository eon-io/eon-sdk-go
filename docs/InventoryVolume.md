# InventoryVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderVolumeId** | **string** | The volume&#39;s cloud-provider-assigned resource ID. | 
**Path** | **string** | Path of the volume on the machine it&#39;s attached to. | 
**Region** | **string** | Region the volume is hosted in. | 
**EncryptionKeyId** | Pointer to **string** | ID of the encryption key used to encrypt the volume. | [optional] 
**IsAwsManagedKey** | Pointer to **bool** | Whether the encryption key is an AWS-managed KMS key. If &#x60;true&#x60;, the key is AWS-managed. If &#x60;false&#x60;, the key is customer-managed.  | [optional] 
**Tags** | **map[string]string** | Volume tags as key-value pairs. Both keys and values are strings. If a tag is a key with no value, the value is presented as an empty string.  **Example:** &#x60;{\&quot;primary\&quot;: \&quot;\&quot;}&#x60;  | 
**AvailabilityZone** | **string** | Volume availability zone. | 
**VolumeSettings** | [**VolumeSettings**](VolumeSettings.md) |  | 

## Methods

### NewInventoryVolume

`func NewInventoryVolume(providerVolumeId string, path string, region string, tags map[string]string, availabilityZone string, volumeSettings VolumeSettings, ) *InventoryVolume`

NewInventoryVolume instantiates a new InventoryVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryVolumeWithDefaults

`func NewInventoryVolumeWithDefaults() *InventoryVolume`

NewInventoryVolumeWithDefaults instantiates a new InventoryVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderVolumeId

`func (o *InventoryVolume) GetProviderVolumeId() string`

GetProviderVolumeId returns the ProviderVolumeId field if non-nil, zero value otherwise.

### GetProviderVolumeIdOk

`func (o *InventoryVolume) GetProviderVolumeIdOk() (*string, bool)`

GetProviderVolumeIdOk returns a tuple with the ProviderVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVolumeId

`func (o *InventoryVolume) SetProviderVolumeId(v string)`

SetProviderVolumeId sets ProviderVolumeId field to given value.


### GetPath

`func (o *InventoryVolume) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InventoryVolume) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InventoryVolume) SetPath(v string)`

SetPath sets Path field to given value.


### GetRegion

`func (o *InventoryVolume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InventoryVolume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InventoryVolume) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetEncryptionKeyId

`func (o *InventoryVolume) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *InventoryVolume) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *InventoryVolume) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *InventoryVolume) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### GetIsAwsManagedKey

`func (o *InventoryVolume) GetIsAwsManagedKey() bool`

GetIsAwsManagedKey returns the IsAwsManagedKey field if non-nil, zero value otherwise.

### GetIsAwsManagedKeyOk

`func (o *InventoryVolume) GetIsAwsManagedKeyOk() (*bool, bool)`

GetIsAwsManagedKeyOk returns a tuple with the IsAwsManagedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAwsManagedKey

`func (o *InventoryVolume) SetIsAwsManagedKey(v bool)`

SetIsAwsManagedKey sets IsAwsManagedKey field to given value.

### HasIsAwsManagedKey

`func (o *InventoryVolume) HasIsAwsManagedKey() bool`

HasIsAwsManagedKey returns a boolean if a field has been set.

### GetTags

`func (o *InventoryVolume) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryVolume) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryVolume) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetAvailabilityZone

`func (o *InventoryVolume) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *InventoryVolume) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *InventoryVolume) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeSettings

`func (o *InventoryVolume) GetVolumeSettings() VolumeSettings`

GetVolumeSettings returns the VolumeSettings field if non-nil, zero value otherwise.

### GetVolumeSettingsOk

`func (o *InventoryVolume) GetVolumeSettingsOk() (*VolumeSettings, bool)`

GetVolumeSettingsOk returns a tuple with the VolumeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSettings

`func (o *InventoryVolume) SetVolumeSettings(v VolumeSettings)`

SetVolumeSettings sets VolumeSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


