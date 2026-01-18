# FsxVolumeBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | Pointer to **string** | The volume ID. | [optional] 
**Name** | Pointer to **string** | The volume name. | [optional] 
**SizeBytes** | Pointer to **int64** | The volume size in bytes. | [optional] 
**JunctionPath** | Pointer to **string** | The volume junction path. | [optional] 
**VolumeType** | Pointer to **string** | The volume type (RW, DP, LS). | [optional] 
**RecoveryPointArn** | Pointer to **string** | The AWS Backup recovery point ARN for this volume. | [optional] 

## Methods

### NewFsxVolumeBackupInfo

`func NewFsxVolumeBackupInfo() *FsxVolumeBackupInfo`

NewFsxVolumeBackupInfo instantiates a new FsxVolumeBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFsxVolumeBackupInfoWithDefaults

`func NewFsxVolumeBackupInfoWithDefaults() *FsxVolumeBackupInfo`

NewFsxVolumeBackupInfoWithDefaults instantiates a new FsxVolumeBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeId

`func (o *FsxVolumeBackupInfo) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *FsxVolumeBackupInfo) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *FsxVolumeBackupInfo) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *FsxVolumeBackupInfo) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetName

`func (o *FsxVolumeBackupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FsxVolumeBackupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FsxVolumeBackupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FsxVolumeBackupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeBytes

`func (o *FsxVolumeBackupInfo) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *FsxVolumeBackupInfo) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *FsxVolumeBackupInfo) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *FsxVolumeBackupInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetJunctionPath

`func (o *FsxVolumeBackupInfo) GetJunctionPath() string`

GetJunctionPath returns the JunctionPath field if non-nil, zero value otherwise.

### GetJunctionPathOk

`func (o *FsxVolumeBackupInfo) GetJunctionPathOk() (*string, bool)`

GetJunctionPathOk returns a tuple with the JunctionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionPath

`func (o *FsxVolumeBackupInfo) SetJunctionPath(v string)`

SetJunctionPath sets JunctionPath field to given value.

### HasJunctionPath

`func (o *FsxVolumeBackupInfo) HasJunctionPath() bool`

HasJunctionPath returns a boolean if a field has been set.

### GetVolumeType

`func (o *FsxVolumeBackupInfo) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *FsxVolumeBackupInfo) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *FsxVolumeBackupInfo) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *FsxVolumeBackupInfo) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetRecoveryPointArn

`func (o *FsxVolumeBackupInfo) GetRecoveryPointArn() string`

GetRecoveryPointArn returns the RecoveryPointArn field if non-nil, zero value otherwise.

### GetRecoveryPointArnOk

`func (o *FsxVolumeBackupInfo) GetRecoveryPointArnOk() (*string, bool)`

GetRecoveryPointArnOk returns a tuple with the RecoveryPointArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointArn

`func (o *FsxVolumeBackupInfo) SetRecoveryPointArn(v string)`

SetRecoveryPointArn sets RecoveryPointArn field to given value.

### HasRecoveryPointArn

`func (o *FsxVolumeBackupInfo) HasRecoveryPointArn() bool`

HasRecoveryPointArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


