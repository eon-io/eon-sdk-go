# GcpDiskSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Disk name. | 
**Type** | **string** | Disk type. | 
**SizeBytes** | **int64** | Size of the disk, in bytes. | 
**Iops** | Pointer to **int64** | Provisioned IOPS for the disk. Applicable only when &#x60;type&#x60; is &#x60;pd-extreme&#x60;. When restoring, defaults to the original IOPS captured by the snapshot.  | [optional] 
**Throughput** | Pointer to **int64** | Disk throughput. When restoring, defaults to the original throughput captured by the snapshot.  | [optional] 
**Description** | Pointer to **string** | Disk description. | [optional] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the restored disk as key-value pairs, where key and value are both strings. These labels are always applied: &#x60;\&quot;eon-restore\&quot;: \&quot;true\&quot;&#x60;, &#x60;\&quot;eon-job-id\&quot;&#x60;, &#x60;\&quot;eon-original-disk-id\&quot;&#x60;, &#x60;\&quot;eon-original-disk-name\&quot;&#x60;.  **Example:** &#x60;{\&quot;primary\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**EncryptionKeyId** | Pointer to **string** | ID of the customer-managed encryption key to use for the disk. | [optional] 

## Methods

### NewGcpDiskSettings

`func NewGcpDiskSettings(name string, type_ string, sizeBytes int64, ) *GcpDiskSettings`

NewGcpDiskSettings instantiates a new GcpDiskSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDiskSettingsWithDefaults

`func NewGcpDiskSettingsWithDefaults() *GcpDiskSettings`

NewGcpDiskSettingsWithDefaults instantiates a new GcpDiskSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GcpDiskSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpDiskSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpDiskSettings) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GcpDiskSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GcpDiskSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GcpDiskSettings) SetType(v string)`

SetType sets Type field to given value.


### GetSizeBytes

`func (o *GcpDiskSettings) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *GcpDiskSettings) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *GcpDiskSettings) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.


### GetIops

`func (o *GcpDiskSettings) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *GcpDiskSettings) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *GcpDiskSettings) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *GcpDiskSettings) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetThroughput

`func (o *GcpDiskSettings) GetThroughput() int64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *GcpDiskSettings) GetThroughputOk() (*int64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *GcpDiskSettings) SetThroughput(v int64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *GcpDiskSettings) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetDescription

`func (o *GcpDiskSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GcpDiskSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GcpDiskSettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GcpDiskSettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *GcpDiskSettings) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpDiskSettings) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpDiskSettings) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpDiskSettings) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetEncryptionKeyId

`func (o *GcpDiskSettings) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *GcpDiskSettings) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *GcpDiskSettings) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *GcpDiskSettings) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


