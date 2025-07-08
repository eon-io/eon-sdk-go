# StandardBackupSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | **string** | Vault ID. | 
**ScheduleConfig** | [**StandardBackupScheduleConfig**](StandardBackupScheduleConfig.md) |  | 
**BackupRetentionDays** | **int32** | Backup retention period, in days. | 

## Methods

### NewStandardBackupSchedules

`func NewStandardBackupSchedules(vaultId string, scheduleConfig StandardBackupScheduleConfig, backupRetentionDays int32, ) *StandardBackupSchedules`

NewStandardBackupSchedules instantiates a new StandardBackupSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardBackupSchedulesWithDefaults

`func NewStandardBackupSchedulesWithDefaults() *StandardBackupSchedules`

NewStandardBackupSchedulesWithDefaults instantiates a new StandardBackupSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *StandardBackupSchedules) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *StandardBackupSchedules) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *StandardBackupSchedules) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetScheduleConfig

`func (o *StandardBackupSchedules) GetScheduleConfig() StandardBackupScheduleConfig`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *StandardBackupSchedules) GetScheduleConfigOk() (*StandardBackupScheduleConfig, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *StandardBackupSchedules) SetScheduleConfig(v StandardBackupScheduleConfig)`

SetScheduleConfig sets ScheduleConfig field to given value.


### GetBackupRetentionDays

`func (o *StandardBackupSchedules) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *StandardBackupSchedules) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *StandardBackupSchedules) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


