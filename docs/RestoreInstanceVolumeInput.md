# RestoreInstanceVolumeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderVolumeId** | **string** | Cloud-provider-assigned ID of the volume to restore. | 
**Description** | Pointer to **string** | Optional description. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored volume as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;primary\&quot;: \&quot;\&quot;}&#x60;  | [optional] 
**VolumeEncryptionKeyId** | **string** | ARN of the KMS key for encrypting the restored volume. | 
**VolumeSettings** | [**VolumeSettings**](VolumeSettings.md) |  | 

## Methods

### NewRestoreInstanceVolumeInput

`func NewRestoreInstanceVolumeInput(providerVolumeId string, volumeEncryptionKeyId string, volumeSettings VolumeSettings, ) *RestoreInstanceVolumeInput`

NewRestoreInstanceVolumeInput instantiates a new RestoreInstanceVolumeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInstanceVolumeInputWithDefaults

`func NewRestoreInstanceVolumeInputWithDefaults() *RestoreInstanceVolumeInput`

NewRestoreInstanceVolumeInputWithDefaults instantiates a new RestoreInstanceVolumeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderVolumeId

`func (o *RestoreInstanceVolumeInput) GetProviderVolumeId() string`

GetProviderVolumeId returns the ProviderVolumeId field if non-nil, zero value otherwise.

### GetProviderVolumeIdOk

`func (o *RestoreInstanceVolumeInput) GetProviderVolumeIdOk() (*string, bool)`

GetProviderVolumeIdOk returns a tuple with the ProviderVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVolumeId

`func (o *RestoreInstanceVolumeInput) SetProviderVolumeId(v string)`

SetProviderVolumeId sets ProviderVolumeId field to given value.


### GetDescription

`func (o *RestoreInstanceVolumeInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestoreInstanceVolumeInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestoreInstanceVolumeInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestoreInstanceVolumeInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *RestoreInstanceVolumeInput) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestoreInstanceVolumeInput) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestoreInstanceVolumeInput) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestoreInstanceVolumeInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeEncryptionKeyId

`func (o *RestoreInstanceVolumeInput) GetVolumeEncryptionKeyId() string`

GetVolumeEncryptionKeyId returns the VolumeEncryptionKeyId field if non-nil, zero value otherwise.

### GetVolumeEncryptionKeyIdOk

`func (o *RestoreInstanceVolumeInput) GetVolumeEncryptionKeyIdOk() (*string, bool)`

GetVolumeEncryptionKeyIdOk returns a tuple with the VolumeEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeEncryptionKeyId

`func (o *RestoreInstanceVolumeInput) SetVolumeEncryptionKeyId(v string)`

SetVolumeEncryptionKeyId sets VolumeEncryptionKeyId field to given value.


### GetVolumeSettings

`func (o *RestoreInstanceVolumeInput) GetVolumeSettings() VolumeSettings`

GetVolumeSettings returns the VolumeSettings field if non-nil, zero value otherwise.

### GetVolumeSettingsOk

`func (o *RestoreInstanceVolumeInput) GetVolumeSettingsOk() (*VolumeSettings, bool)`

GetVolumeSettingsOk returns a tuple with the VolumeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSettings

`func (o *RestoreInstanceVolumeInput) SetVolumeSettings(v VolumeSettings)`

SetVolumeSettings sets VolumeSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


