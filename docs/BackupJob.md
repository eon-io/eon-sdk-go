# BackupJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobExecutionDetails** | [**JobExecutionDetails**](JobExecutionDetails.md) |  | 
**ResourceDetails** | Pointer to [**NullableResourceDetails**](ResourceDetails.md) |  | [optional] 
**SnapshotDetails** | Pointer to [**NullableJobSnapshotDetails**](JobSnapshotDetails.md) |  | [optional] 
**BackupType** | [**BackupJobType**](BackupJobType.md) |  | 
**Vault** | Pointer to [**NullableBackupVault**](BackupVault.md) |  | [optional] 

## Methods

### NewBackupJob

`func NewBackupJob(jobExecutionDetails JobExecutionDetails, backupType BackupJobType, ) *BackupJob`

NewBackupJob instantiates a new BackupJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobWithDefaults

`func NewBackupJobWithDefaults() *BackupJob`

NewBackupJobWithDefaults instantiates a new BackupJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobExecutionDetails

`func (o *BackupJob) GetJobExecutionDetails() JobExecutionDetails`

GetJobExecutionDetails returns the JobExecutionDetails field if non-nil, zero value otherwise.

### GetJobExecutionDetailsOk

`func (o *BackupJob) GetJobExecutionDetailsOk() (*JobExecutionDetails, bool)`

GetJobExecutionDetailsOk returns a tuple with the JobExecutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutionDetails

`func (o *BackupJob) SetJobExecutionDetails(v JobExecutionDetails)`

SetJobExecutionDetails sets JobExecutionDetails field to given value.


### GetResourceDetails

`func (o *BackupJob) GetResourceDetails() ResourceDetails`

GetResourceDetails returns the ResourceDetails field if non-nil, zero value otherwise.

### GetResourceDetailsOk

`func (o *BackupJob) GetResourceDetailsOk() (*ResourceDetails, bool)`

GetResourceDetailsOk returns a tuple with the ResourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDetails

`func (o *BackupJob) SetResourceDetails(v ResourceDetails)`

SetResourceDetails sets ResourceDetails field to given value.

### HasResourceDetails

`func (o *BackupJob) HasResourceDetails() bool`

HasResourceDetails returns a boolean if a field has been set.

### SetResourceDetailsNil

`func (o *BackupJob) SetResourceDetailsNil(b bool)`

 SetResourceDetailsNil sets the value for ResourceDetails to be an explicit nil

### UnsetResourceDetails
`func (o *BackupJob) UnsetResourceDetails()`

UnsetResourceDetails ensures that no value is present for ResourceDetails, not even an explicit nil
### GetSnapshotDetails

`func (o *BackupJob) GetSnapshotDetails() JobSnapshotDetails`

GetSnapshotDetails returns the SnapshotDetails field if non-nil, zero value otherwise.

### GetSnapshotDetailsOk

`func (o *BackupJob) GetSnapshotDetailsOk() (*JobSnapshotDetails, bool)`

GetSnapshotDetailsOk returns a tuple with the SnapshotDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDetails

`func (o *BackupJob) SetSnapshotDetails(v JobSnapshotDetails)`

SetSnapshotDetails sets SnapshotDetails field to given value.

### HasSnapshotDetails

`func (o *BackupJob) HasSnapshotDetails() bool`

HasSnapshotDetails returns a boolean if a field has been set.

### SetSnapshotDetailsNil

`func (o *BackupJob) SetSnapshotDetailsNil(b bool)`

 SetSnapshotDetailsNil sets the value for SnapshotDetails to be an explicit nil

### UnsetSnapshotDetails
`func (o *BackupJob) UnsetSnapshotDetails()`

UnsetSnapshotDetails ensures that no value is present for SnapshotDetails, not even an explicit nil
### GetBackupType

`func (o *BackupJob) GetBackupType() BackupJobType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupJob) GetBackupTypeOk() (*BackupJobType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupJob) SetBackupType(v BackupJobType)`

SetBackupType sets BackupType field to given value.


### GetVault

`func (o *BackupJob) GetVault() BackupVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *BackupJob) GetVaultOk() (*BackupVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *BackupJob) SetVault(v BackupVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *BackupJob) HasVault() bool`

HasVault returns a boolean if a field has been set.

### SetVaultNil

`func (o *BackupJob) SetVaultNil(b bool)`

 SetVaultNil sets the value for Vault to be an explicit nil

### UnsetVault
`func (o *BackupJob) UnsetVault()`

UnsetVault ensures that no value is present for Vault, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


