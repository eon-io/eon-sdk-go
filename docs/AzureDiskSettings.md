# AzureDiskSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the disk to restore. | 
**Type** | **string** | Type of the disk to restore. | 
**Tier** | **string** | Tier of the disk to restore. | 
**HyperVGeneration** | Pointer to **string** | Hyper-V generation of the disk to restore. Defaults to the original Hyper-V generation captured by the snapshot.  | [optional] 
**SizeBytes** | Pointer to **int64** | Size of the disk to restore, in bytes. Defaults to the original disk size captured by the snapshot.  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored disk as key-value pairs, where key and value are both strings. If not provided, defaults to an empty object, with no tags applied.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 

## Methods

### NewAzureDiskSettings

`func NewAzureDiskSettings(name string, type_ string, tier string, ) *AzureDiskSettings`

NewAzureDiskSettings instantiates a new AzureDiskSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDiskSettingsWithDefaults

`func NewAzureDiskSettingsWithDefaults() *AzureDiskSettings`

NewAzureDiskSettingsWithDefaults instantiates a new AzureDiskSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureDiskSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureDiskSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureDiskSettings) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AzureDiskSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureDiskSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureDiskSettings) SetType(v string)`

SetType sets Type field to given value.


### GetTier

`func (o *AzureDiskSettings) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AzureDiskSettings) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AzureDiskSettings) SetTier(v string)`

SetTier sets Tier field to given value.


### GetHyperVGeneration

`func (o *AzureDiskSettings) GetHyperVGeneration() string`

GetHyperVGeneration returns the HyperVGeneration field if non-nil, zero value otherwise.

### GetHyperVGenerationOk

`func (o *AzureDiskSettings) GetHyperVGenerationOk() (*string, bool)`

GetHyperVGenerationOk returns a tuple with the HyperVGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGeneration

`func (o *AzureDiskSettings) SetHyperVGeneration(v string)`

SetHyperVGeneration sets HyperVGeneration field to given value.

### HasHyperVGeneration

`func (o *AzureDiskSettings) HasHyperVGeneration() bool`

HasHyperVGeneration returns a boolean if a field has been set.

### GetSizeBytes

`func (o *AzureDiskSettings) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *AzureDiskSettings) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *AzureDiskSettings) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *AzureDiskSettings) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetTags

`func (o *AzureDiskSettings) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureDiskSettings) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureDiskSettings) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureDiskSettings) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


