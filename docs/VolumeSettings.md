# VolumeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Volume type. | 
**SizeBytes** | **int64** | Volume size in bytes. | 
**Iops** | Pointer to **int32** | Volume IOPS. | [optional] 
**Throughput** | Pointer to **int32** | Volume throughput. | [optional] 

## Methods

### NewVolumeSettings

`func NewVolumeSettings(type_ string, sizeBytes int64, ) *VolumeSettings`

NewVolumeSettings instantiates a new VolumeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSettingsWithDefaults

`func NewVolumeSettingsWithDefaults() *VolumeSettings`

NewVolumeSettingsWithDefaults instantiates a new VolumeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VolumeSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VolumeSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VolumeSettings) SetType(v string)`

SetType sets Type field to given value.


### GetSizeBytes

`func (o *VolumeSettings) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *VolumeSettings) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *VolumeSettings) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.


### GetIops

`func (o *VolumeSettings) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *VolumeSettings) GetIopsOk() (*int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *VolumeSettings) SetIops(v int32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *VolumeSettings) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetThroughput

`func (o *VolumeSettings) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *VolumeSettings) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *VolumeSettings) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *VolumeSettings) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


