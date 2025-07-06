# RestoreJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobExecutionDetails** | [**JobExecutionDetails**](JobExecutionDetails.md) |  | 
**ResourceDetails** | Pointer to [**NullableResourceDetails**](ResourceDetails.md) |  | [optional] 
**SnapshotDetails** | Pointer to [**NullableJobSnapshotDetails**](JobSnapshotDetails.md) |  | [optional] 
**DestinationDetails** | [**DestinationDetails**](DestinationDetails.md) |  | 
**Vault** | Pointer to [**NullableBackupVault**](BackupVault.md) |  | [optional] 
**RestoreType** | [**RestoreJobType**](RestoreJobType.md) |  | 

## Methods

### NewRestoreJob

`func NewRestoreJob(jobExecutionDetails JobExecutionDetails, destinationDetails DestinationDetails, restoreType RestoreJobType, ) *RestoreJob`

NewRestoreJob instantiates a new RestoreJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobWithDefaults

`func NewRestoreJobWithDefaults() *RestoreJob`

NewRestoreJobWithDefaults instantiates a new RestoreJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobExecutionDetails

`func (o *RestoreJob) GetJobExecutionDetails() JobExecutionDetails`

GetJobExecutionDetails returns the JobExecutionDetails field if non-nil, zero value otherwise.

### GetJobExecutionDetailsOk

`func (o *RestoreJob) GetJobExecutionDetailsOk() (*JobExecutionDetails, bool)`

GetJobExecutionDetailsOk returns a tuple with the JobExecutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutionDetails

`func (o *RestoreJob) SetJobExecutionDetails(v JobExecutionDetails)`

SetJobExecutionDetails sets JobExecutionDetails field to given value.


### GetResourceDetails

`func (o *RestoreJob) GetResourceDetails() ResourceDetails`

GetResourceDetails returns the ResourceDetails field if non-nil, zero value otherwise.

### GetResourceDetailsOk

`func (o *RestoreJob) GetResourceDetailsOk() (*ResourceDetails, bool)`

GetResourceDetailsOk returns a tuple with the ResourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDetails

`func (o *RestoreJob) SetResourceDetails(v ResourceDetails)`

SetResourceDetails sets ResourceDetails field to given value.

### HasResourceDetails

`func (o *RestoreJob) HasResourceDetails() bool`

HasResourceDetails returns a boolean if a field has been set.

### SetResourceDetailsNil

`func (o *RestoreJob) SetResourceDetailsNil(b bool)`

 SetResourceDetailsNil sets the value for ResourceDetails to be an explicit nil

### UnsetResourceDetails
`func (o *RestoreJob) UnsetResourceDetails()`

UnsetResourceDetails ensures that no value is present for ResourceDetails, not even an explicit nil
### GetSnapshotDetails

`func (o *RestoreJob) GetSnapshotDetails() JobSnapshotDetails`

GetSnapshotDetails returns the SnapshotDetails field if non-nil, zero value otherwise.

### GetSnapshotDetailsOk

`func (o *RestoreJob) GetSnapshotDetailsOk() (*JobSnapshotDetails, bool)`

GetSnapshotDetailsOk returns a tuple with the SnapshotDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDetails

`func (o *RestoreJob) SetSnapshotDetails(v JobSnapshotDetails)`

SetSnapshotDetails sets SnapshotDetails field to given value.

### HasSnapshotDetails

`func (o *RestoreJob) HasSnapshotDetails() bool`

HasSnapshotDetails returns a boolean if a field has been set.

### SetSnapshotDetailsNil

`func (o *RestoreJob) SetSnapshotDetailsNil(b bool)`

 SetSnapshotDetailsNil sets the value for SnapshotDetails to be an explicit nil

### UnsetSnapshotDetails
`func (o *RestoreJob) UnsetSnapshotDetails()`

UnsetSnapshotDetails ensures that no value is present for SnapshotDetails, not even an explicit nil
### GetDestinationDetails

`func (o *RestoreJob) GetDestinationDetails() DestinationDetails`

GetDestinationDetails returns the DestinationDetails field if non-nil, zero value otherwise.

### GetDestinationDetailsOk

`func (o *RestoreJob) GetDestinationDetailsOk() (*DestinationDetails, bool)`

GetDestinationDetailsOk returns a tuple with the DestinationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDetails

`func (o *RestoreJob) SetDestinationDetails(v DestinationDetails)`

SetDestinationDetails sets DestinationDetails field to given value.


### GetVault

`func (o *RestoreJob) GetVault() BackupVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *RestoreJob) GetVaultOk() (*BackupVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *RestoreJob) SetVault(v BackupVault)`

SetVault sets Vault field to given value.

### HasVault

`func (o *RestoreJob) HasVault() bool`

HasVault returns a boolean if a field has been set.

### SetVaultNil

`func (o *RestoreJob) SetVaultNil(b bool)`

 SetVaultNil sets the value for Vault to be an explicit nil

### UnsetVault
`func (o *RestoreJob) UnsetVault()`

UnsetVault ensures that no value is present for Vault, not even an explicit nil
### GetRestoreType

`func (o *RestoreJob) GetRestoreType() RestoreJobType`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *RestoreJob) GetRestoreTypeOk() (*RestoreJobType, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *RestoreJob) SetRestoreType(v RestoreJobType)`

SetRestoreType sets RestoreType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


