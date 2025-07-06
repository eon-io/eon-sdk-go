# InventorySnapshotVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderVolumeId** | **string** | Cloud-provider-assigned volume ID. | 
**Region** | **string** | Region the volume is hosted in. | 
**AvailabilityZone** | **string** | Availability zone the volume is hosted in. | 
**Tags** | **map[string]string** | Volume tags as key-value pairs. Both keys and values are strings.  **Example**: &#x60;{\&quot;env\&quot;: \&quot;prod\&quot;, \&quot;db\&quot;: \&quot;\&quot;}&#x60;  | 
**VolumeSettings** | [**VolumeSettings**](VolumeSettings.md) |  | 

## Methods

### NewInventorySnapshotVolume

`func NewInventorySnapshotVolume(providerVolumeId string, region string, availabilityZone string, tags map[string]string, volumeSettings VolumeSettings, ) *InventorySnapshotVolume`

NewInventorySnapshotVolume instantiates a new InventorySnapshotVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventorySnapshotVolumeWithDefaults

`func NewInventorySnapshotVolumeWithDefaults() *InventorySnapshotVolume`

NewInventorySnapshotVolumeWithDefaults instantiates a new InventorySnapshotVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderVolumeId

`func (o *InventorySnapshotVolume) GetProviderVolumeId() string`

GetProviderVolumeId returns the ProviderVolumeId field if non-nil, zero value otherwise.

### GetProviderVolumeIdOk

`func (o *InventorySnapshotVolume) GetProviderVolumeIdOk() (*string, bool)`

GetProviderVolumeIdOk returns a tuple with the ProviderVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVolumeId

`func (o *InventorySnapshotVolume) SetProviderVolumeId(v string)`

SetProviderVolumeId sets ProviderVolumeId field to given value.


### GetRegion

`func (o *InventorySnapshotVolume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InventorySnapshotVolume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InventorySnapshotVolume) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAvailabilityZone

`func (o *InventorySnapshotVolume) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *InventorySnapshotVolume) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *InventorySnapshotVolume) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetTags

`func (o *InventorySnapshotVolume) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventorySnapshotVolume) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventorySnapshotVolume) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetVolumeSettings

`func (o *InventorySnapshotVolume) GetVolumeSettings() VolumeSettings`

GetVolumeSettings returns the VolumeSettings field if non-nil, zero value otherwise.

### GetVolumeSettingsOk

`func (o *InventorySnapshotVolume) GetVolumeSettingsOk() (*VolumeSettings, bool)`

GetVolumeSettingsOk returns a tuple with the VolumeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSettings

`func (o *InventorySnapshotVolume) SetVolumeSettings(v VolumeSettings)`

SetVolumeSettings sets VolumeSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


