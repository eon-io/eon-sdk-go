# AzureDiskSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Disk name. | 
**Type** | **string** | Disk type. | 
**Tier** | **string** | Disk tier. | 
**HyperVGeneration** | Pointer to **string** | Hyper-V generation of the disk. When restoring, defaults to the original Hyper-V generation captured by the snapshot.  | [optional] 
**SizeBytes** | Pointer to **int64** | Size of the disk, in bytes. When restoring, defaults to the original disk size captured by the snapshot.  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored disk as key-value pairs, where key and value are both strings. If not provided, defaults to an empty object, with no tags applied.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**Zone** | Pointer to **string** | Availability zone to create the restored disk in (for example &#x60;1&#x60;). Only for restoring a disk on its own; a VM restore rejects it. When omitted, the disk is created without a zone, except for Premium SSD v2 disks, which follow the zone of the restore. The zone must be available for the disk type in the destination region.  | [optional] 
**DiskIopsReadWrite** | Pointer to **int64** | Provisioned IOPS of the restored disk. Only valid for the &#x60;PremiumV2_LRS&#x60; and &#x60;UltraSSD_LRS&#x60; disk types. When omitted, Azure applies the default for the disk type and size.  | [optional] 
**DiskMBpsReadWrite** | Pointer to **int64** | Provisioned throughput of the restored disk, in MB/s. Only valid for the &#x60;PremiumV2_LRS&#x60; and &#x60;UltraSSD_LRS&#x60; disk types. When omitted, Azure applies the default for the disk type and size.  | [optional] 
**NetworkAccessPolicy** | Pointer to **string** | Network access policy of the restored disk, which controls how it can be exported or imported. Defaults to &#x60;AllowAll&#x60;.  | [optional] 
**PublicNetworkAccess** | Pointer to **string** | Whether the restored disk can be exported or imported over the public network. Defaults to &#x60;Enabled&#x60;.  | [optional] 

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

### GetZone

`func (o *AzureDiskSettings) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AzureDiskSettings) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AzureDiskSettings) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AzureDiskSettings) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDiskIopsReadWrite

`func (o *AzureDiskSettings) GetDiskIopsReadWrite() int64`

GetDiskIopsReadWrite returns the DiskIopsReadWrite field if non-nil, zero value otherwise.

### GetDiskIopsReadWriteOk

`func (o *AzureDiskSettings) GetDiskIopsReadWriteOk() (*int64, bool)`

GetDiskIopsReadWriteOk returns a tuple with the DiskIopsReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIopsReadWrite

`func (o *AzureDiskSettings) SetDiskIopsReadWrite(v int64)`

SetDiskIopsReadWrite sets DiskIopsReadWrite field to given value.

### HasDiskIopsReadWrite

`func (o *AzureDiskSettings) HasDiskIopsReadWrite() bool`

HasDiskIopsReadWrite returns a boolean if a field has been set.

### GetDiskMBpsReadWrite

`func (o *AzureDiskSettings) GetDiskMBpsReadWrite() int64`

GetDiskMBpsReadWrite returns the DiskMBpsReadWrite field if non-nil, zero value otherwise.

### GetDiskMBpsReadWriteOk

`func (o *AzureDiskSettings) GetDiskMBpsReadWriteOk() (*int64, bool)`

GetDiskMBpsReadWriteOk returns a tuple with the DiskMBpsReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMBpsReadWrite

`func (o *AzureDiskSettings) SetDiskMBpsReadWrite(v int64)`

SetDiskMBpsReadWrite sets DiskMBpsReadWrite field to given value.

### HasDiskMBpsReadWrite

`func (o *AzureDiskSettings) HasDiskMBpsReadWrite() bool`

HasDiskMBpsReadWrite returns a boolean if a field has been set.

### GetNetworkAccessPolicy

`func (o *AzureDiskSettings) GetNetworkAccessPolicy() string`

GetNetworkAccessPolicy returns the NetworkAccessPolicy field if non-nil, zero value otherwise.

### GetNetworkAccessPolicyOk

`func (o *AzureDiskSettings) GetNetworkAccessPolicyOk() (*string, bool)`

GetNetworkAccessPolicyOk returns a tuple with the NetworkAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessPolicy

`func (o *AzureDiskSettings) SetNetworkAccessPolicy(v string)`

SetNetworkAccessPolicy sets NetworkAccessPolicy field to given value.

### HasNetworkAccessPolicy

`func (o *AzureDiskSettings) HasNetworkAccessPolicy() bool`

HasNetworkAccessPolicy returns a boolean if a field has been set.

### GetPublicNetworkAccess

`func (o *AzureDiskSettings) GetPublicNetworkAccess() string`

GetPublicNetworkAccess returns the PublicNetworkAccess field if non-nil, zero value otherwise.

### GetPublicNetworkAccessOk

`func (o *AzureDiskSettings) GetPublicNetworkAccessOk() (*string, bool)`

GetPublicNetworkAccessOk returns a tuple with the PublicNetworkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkAccess

`func (o *AzureDiskSettings) SetPublicNetworkAccess(v string)`

SetPublicNetworkAccess sets PublicNetworkAccess field to given value.

### HasPublicNetworkAccess

`func (o *AzureDiskSettings) HasPublicNetworkAccess() bool`

HasPublicNetworkAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


