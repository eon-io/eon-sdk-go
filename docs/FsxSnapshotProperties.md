# FsxSnapshotProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryPointArn** | Pointer to **string** | AWS Backup recovery point ARN (primary for file systems, or first volume for ONTAP). | [optional] 
**BackupVaultName** | Pointer to **string** | AWS Backup vault name where the backup is stored. | [optional] 
**Region** | Pointer to **string** | AWS region where the backup is stored. | [optional] 
**SourceAccountId** | Pointer to **string** | AWS account ID where the backup is stored. | [optional] 
**VolumeBackupArns** | Pointer to **[]string** | For ONTAP file systems, the list of recovery point ARNs for each volume backup. | [optional] 
**BackedUpVolumes** | Pointer to [**[]FsxVolumeBackupInfo**](FsxVolumeBackupInfo.md) | Details of volumes that were backed up in this snapshot. | [optional] 

## Methods

### NewFsxSnapshotProperties

`func NewFsxSnapshotProperties() *FsxSnapshotProperties`

NewFsxSnapshotProperties instantiates a new FsxSnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFsxSnapshotPropertiesWithDefaults

`func NewFsxSnapshotPropertiesWithDefaults() *FsxSnapshotProperties`

NewFsxSnapshotPropertiesWithDefaults instantiates a new FsxSnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryPointArn

`func (o *FsxSnapshotProperties) GetRecoveryPointArn() string`

GetRecoveryPointArn returns the RecoveryPointArn field if non-nil, zero value otherwise.

### GetRecoveryPointArnOk

`func (o *FsxSnapshotProperties) GetRecoveryPointArnOk() (*string, bool)`

GetRecoveryPointArnOk returns a tuple with the RecoveryPointArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointArn

`func (o *FsxSnapshotProperties) SetRecoveryPointArn(v string)`

SetRecoveryPointArn sets RecoveryPointArn field to given value.

### HasRecoveryPointArn

`func (o *FsxSnapshotProperties) HasRecoveryPointArn() bool`

HasRecoveryPointArn returns a boolean if a field has been set.

### GetBackupVaultName

`func (o *FsxSnapshotProperties) GetBackupVaultName() string`

GetBackupVaultName returns the BackupVaultName field if non-nil, zero value otherwise.

### GetBackupVaultNameOk

`func (o *FsxSnapshotProperties) GetBackupVaultNameOk() (*string, bool)`

GetBackupVaultNameOk returns a tuple with the BackupVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupVaultName

`func (o *FsxSnapshotProperties) SetBackupVaultName(v string)`

SetBackupVaultName sets BackupVaultName field to given value.

### HasBackupVaultName

`func (o *FsxSnapshotProperties) HasBackupVaultName() bool`

HasBackupVaultName returns a boolean if a field has been set.

### GetRegion

`func (o *FsxSnapshotProperties) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FsxSnapshotProperties) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FsxSnapshotProperties) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FsxSnapshotProperties) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *FsxSnapshotProperties) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *FsxSnapshotProperties) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *FsxSnapshotProperties) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *FsxSnapshotProperties) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetVolumeBackupArns

`func (o *FsxSnapshotProperties) GetVolumeBackupArns() []string`

GetVolumeBackupArns returns the VolumeBackupArns field if non-nil, zero value otherwise.

### GetVolumeBackupArnsOk

`func (o *FsxSnapshotProperties) GetVolumeBackupArnsOk() (*[]string, bool)`

GetVolumeBackupArnsOk returns a tuple with the VolumeBackupArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBackupArns

`func (o *FsxSnapshotProperties) SetVolumeBackupArns(v []string)`

SetVolumeBackupArns sets VolumeBackupArns field to given value.

### HasVolumeBackupArns

`func (o *FsxSnapshotProperties) HasVolumeBackupArns() bool`

HasVolumeBackupArns returns a boolean if a field has been set.

### GetBackedUpVolumes

`func (o *FsxSnapshotProperties) GetBackedUpVolumes() []FsxVolumeBackupInfo`

GetBackedUpVolumes returns the BackedUpVolumes field if non-nil, zero value otherwise.

### GetBackedUpVolumesOk

`func (o *FsxSnapshotProperties) GetBackedUpVolumesOk() (*[]FsxVolumeBackupInfo, bool)`

GetBackedUpVolumesOk returns a tuple with the BackedUpVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUpVolumes

`func (o *FsxSnapshotProperties) SetBackedUpVolumes(v []FsxVolumeBackupInfo)`

SetBackedUpVolumes sets BackedUpVolumes field to given value.

### HasBackedUpVolumes

`func (o *FsxSnapshotProperties) HasBackedUpVolumes() bool`

HasBackedUpVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


