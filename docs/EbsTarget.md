# EbsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description to apply to the restored volume.  Default: No description.  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored volume as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**VolumeEncryptionKeyId** | **string** | ID of the key you want Eon to use for encrypting the restored volume. | 
**AvailabilityZone** | **string** | Availability zone to restore the volume to. | 
**VolumeSettings** | [**VolumeSettings**](VolumeSettings.md) |  | 

## Methods

### NewEbsTarget

`func NewEbsTarget(volumeEncryptionKeyId string, availabilityZone string, volumeSettings VolumeSettings, ) *EbsTarget`

NewEbsTarget instantiates a new EbsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbsTargetWithDefaults

`func NewEbsTargetWithDefaults() *EbsTarget`

NewEbsTargetWithDefaults instantiates a new EbsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EbsTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EbsTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EbsTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EbsTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EbsTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EbsTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EbsTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EbsTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeEncryptionKeyId

`func (o *EbsTarget) GetVolumeEncryptionKeyId() string`

GetVolumeEncryptionKeyId returns the VolumeEncryptionKeyId field if non-nil, zero value otherwise.

### GetVolumeEncryptionKeyIdOk

`func (o *EbsTarget) GetVolumeEncryptionKeyIdOk() (*string, bool)`

GetVolumeEncryptionKeyIdOk returns a tuple with the VolumeEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeEncryptionKeyId

`func (o *EbsTarget) SetVolumeEncryptionKeyId(v string)`

SetVolumeEncryptionKeyId sets VolumeEncryptionKeyId field to given value.


### GetAvailabilityZone

`func (o *EbsTarget) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *EbsTarget) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *EbsTarget) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeSettings

`func (o *EbsTarget) GetVolumeSettings() VolumeSettings`

GetVolumeSettings returns the VolumeSettings field if non-nil, zero value otherwise.

### GetVolumeSettingsOk

`func (o *EbsTarget) GetVolumeSettingsOk() (*VolumeSettings, bool)`

GetVolumeSettingsOk returns a tuple with the VolumeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSettings

`func (o *EbsTarget) SetVolumeSettings(v VolumeSettings)`

SetVolumeSettings sets VolumeSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


