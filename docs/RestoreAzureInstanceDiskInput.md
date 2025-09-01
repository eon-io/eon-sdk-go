# RestoreAzureInstanceDiskInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore | 
**Name** | **string** | The name of the disk to restore | 
**Type** | **string** | The type of the disk to restore | 
**Tier** | **string** | The tier of the disk to restore | 
**HyperVGeneration** | Pointer to **string** | The Hyper-V generation of the disk to restore | [optional] 
**SizeBytes** | Pointer to **int64** | The size of the disk to restore in bytes | [optional] 
**Tags** | Pointer to **map[string]string** | Optional tags to apply to the output instance | [optional] 

## Methods

### NewRestoreAzureInstanceDiskInput

`func NewRestoreAzureInstanceDiskInput(providerDiskId string, name string, type_ string, tier string, ) *RestoreAzureInstanceDiskInput`

NewRestoreAzureInstanceDiskInput instantiates a new RestoreAzureInstanceDiskInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAzureInstanceDiskInputWithDefaults

`func NewRestoreAzureInstanceDiskInputWithDefaults() *RestoreAzureInstanceDiskInput`

NewRestoreAzureInstanceDiskInputWithDefaults instantiates a new RestoreAzureInstanceDiskInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderDiskId

`func (o *RestoreAzureInstanceDiskInput) GetProviderDiskId() string`

GetProviderDiskId returns the ProviderDiskId field if non-nil, zero value otherwise.

### GetProviderDiskIdOk

`func (o *RestoreAzureInstanceDiskInput) GetProviderDiskIdOk() (*string, bool)`

GetProviderDiskIdOk returns a tuple with the ProviderDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDiskId

`func (o *RestoreAzureInstanceDiskInput) SetProviderDiskId(v string)`

SetProviderDiskId sets ProviderDiskId field to given value.


### GetName

`func (o *RestoreAzureInstanceDiskInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreAzureInstanceDiskInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreAzureInstanceDiskInput) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RestoreAzureInstanceDiskInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreAzureInstanceDiskInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreAzureInstanceDiskInput) SetType(v string)`

SetType sets Type field to given value.


### GetTier

`func (o *RestoreAzureInstanceDiskInput) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RestoreAzureInstanceDiskInput) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RestoreAzureInstanceDiskInput) SetTier(v string)`

SetTier sets Tier field to given value.


### GetHyperVGeneration

`func (o *RestoreAzureInstanceDiskInput) GetHyperVGeneration() string`

GetHyperVGeneration returns the HyperVGeneration field if non-nil, zero value otherwise.

### GetHyperVGenerationOk

`func (o *RestoreAzureInstanceDiskInput) GetHyperVGenerationOk() (*string, bool)`

GetHyperVGenerationOk returns a tuple with the HyperVGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGeneration

`func (o *RestoreAzureInstanceDiskInput) SetHyperVGeneration(v string)`

SetHyperVGeneration sets HyperVGeneration field to given value.

### HasHyperVGeneration

`func (o *RestoreAzureInstanceDiskInput) HasHyperVGeneration() bool`

HasHyperVGeneration returns a boolean if a field has been set.

### GetSizeBytes

`func (o *RestoreAzureInstanceDiskInput) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *RestoreAzureInstanceDiskInput) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *RestoreAzureInstanceDiskInput) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *RestoreAzureInstanceDiskInput) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetTags

`func (o *RestoreAzureInstanceDiskInput) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestoreAzureInstanceDiskInput) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestoreAzureInstanceDiskInput) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestoreAzureInstanceDiskInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


