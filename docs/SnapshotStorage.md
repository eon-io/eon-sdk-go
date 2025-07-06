# SnapshotStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeBytes** | Pointer to **int64** | Total size of the resource&#39;s Eon snapshots, in bytes. | [optional] 
**EonSnapshotCount** | Pointer to **int32** | Number of Eon snapshots backing up the resource. | [optional] 
**NonEonSnapshotCount** | Pointer to **int32** | Number of non-Eon snapshots backing up the resource. Only cloud-provider-native snapshots in the source account are included in this count.  | [optional] 

## Methods

### NewSnapshotStorage

`func NewSnapshotStorage() *SnapshotStorage`

NewSnapshotStorage instantiates a new SnapshotStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotStorageWithDefaults

`func NewSnapshotStorageWithDefaults() *SnapshotStorage`

NewSnapshotStorageWithDefaults instantiates a new SnapshotStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeBytes

`func (o *SnapshotStorage) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SnapshotStorage) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SnapshotStorage) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *SnapshotStorage) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetEonSnapshotCount

`func (o *SnapshotStorage) GetEonSnapshotCount() int32`

GetEonSnapshotCount returns the EonSnapshotCount field if non-nil, zero value otherwise.

### GetEonSnapshotCountOk

`func (o *SnapshotStorage) GetEonSnapshotCountOk() (*int32, bool)`

GetEonSnapshotCountOk returns a tuple with the EonSnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEonSnapshotCount

`func (o *SnapshotStorage) SetEonSnapshotCount(v int32)`

SetEonSnapshotCount sets EonSnapshotCount field to given value.

### HasEonSnapshotCount

`func (o *SnapshotStorage) HasEonSnapshotCount() bool`

HasEonSnapshotCount returns a boolean if a field has been set.

### GetNonEonSnapshotCount

`func (o *SnapshotStorage) GetNonEonSnapshotCount() int32`

GetNonEonSnapshotCount returns the NonEonSnapshotCount field if non-nil, zero value otherwise.

### GetNonEonSnapshotCountOk

`func (o *SnapshotStorage) GetNonEonSnapshotCountOk() (*int32, bool)`

GetNonEonSnapshotCountOk returns a tuple with the NonEonSnapshotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonEonSnapshotCount

`func (o *SnapshotStorage) SetNonEonSnapshotCount(v int32)`

SetNonEonSnapshotCount sets NonEonSnapshotCount field to given value.

### HasNonEonSnapshotCount

`func (o *SnapshotStorage) HasNonEonSnapshotCount() bool`

HasNonEonSnapshotCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


