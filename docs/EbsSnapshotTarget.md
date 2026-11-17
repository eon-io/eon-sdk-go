# EbsSnapshotTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to restore the EBS snapshot to. | 
**Description** | Pointer to **string** | Description to apply to the EBS snapshot. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the EBS snapshot as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**SnapshotEncryptionKeyId** | **string** | ID of the key you want Eon to use for encrypting the EBS snapshot. | 

## Methods

### NewEbsSnapshotTarget

`func NewEbsSnapshotTarget(region string, snapshotEncryptionKeyId string, ) *EbsSnapshotTarget`

NewEbsSnapshotTarget instantiates a new EbsSnapshotTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbsSnapshotTargetWithDefaults

`func NewEbsSnapshotTargetWithDefaults() *EbsSnapshotTarget`

NewEbsSnapshotTargetWithDefaults instantiates a new EbsSnapshotTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *EbsSnapshotTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EbsSnapshotTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EbsSnapshotTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDescription

`func (o *EbsSnapshotTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EbsSnapshotTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EbsSnapshotTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EbsSnapshotTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *EbsSnapshotTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EbsSnapshotTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EbsSnapshotTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EbsSnapshotTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSnapshotEncryptionKeyId

`func (o *EbsSnapshotTarget) GetSnapshotEncryptionKeyId() string`

GetSnapshotEncryptionKeyId returns the SnapshotEncryptionKeyId field if non-nil, zero value otherwise.

### GetSnapshotEncryptionKeyIdOk

`func (o *EbsSnapshotTarget) GetSnapshotEncryptionKeyIdOk() (*string, bool)`

GetSnapshotEncryptionKeyIdOk returns a tuple with the SnapshotEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotEncryptionKeyId

`func (o *EbsSnapshotTarget) SetSnapshotEncryptionKeyId(v string)`

SetSnapshotEncryptionKeyId sets SnapshotEncryptionKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


