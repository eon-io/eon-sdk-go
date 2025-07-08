# HighFrequencyBackupSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | **string** | Vault ID. | 
**ScheduleConfig** | [**HighFrequencyBackupScheduleConfig**](HighFrequencyBackupScheduleConfig.md) |  | 
**BackupRetentionDays** | **int32** | Backup retention period, in days. | 

## Methods

### NewHighFrequencyBackupSchedules

`func NewHighFrequencyBackupSchedules(vaultId string, scheduleConfig HighFrequencyBackupScheduleConfig, backupRetentionDays int32, ) *HighFrequencyBackupSchedules`

NewHighFrequencyBackupSchedules instantiates a new HighFrequencyBackupSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighFrequencyBackupSchedulesWithDefaults

`func NewHighFrequencyBackupSchedulesWithDefaults() *HighFrequencyBackupSchedules`

NewHighFrequencyBackupSchedulesWithDefaults instantiates a new HighFrequencyBackupSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *HighFrequencyBackupSchedules) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *HighFrequencyBackupSchedules) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *HighFrequencyBackupSchedules) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetScheduleConfig

`func (o *HighFrequencyBackupSchedules) GetScheduleConfig() HighFrequencyBackupScheduleConfig`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *HighFrequencyBackupSchedules) GetScheduleConfigOk() (*HighFrequencyBackupScheduleConfig, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *HighFrequencyBackupSchedules) SetScheduleConfig(v HighFrequencyBackupScheduleConfig)`

SetScheduleConfig sets ScheduleConfig field to given value.


### GetBackupRetentionDays

`func (o *HighFrequencyBackupSchedules) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *HighFrequencyBackupSchedules) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *HighFrequencyBackupSchedules) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


