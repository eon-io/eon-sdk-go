# StandardBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTypes** | Pointer to [**[]ResourceType**](ResourceType.md) | Optional resource types covered by this policy. On creation, omission or an empty array covers all supported standard backup resource types. On update, omission keeps the existing scope; an empty array clears it. A specified set also limits manual inclusions.  | [optional] 
**BackupSchedules** | [**[]StandardBackupSchedules**](StandardBackupSchedules.md) | List of backup schedules. Each schedule specifies a backup frequency, retention period, and vault.  | 
**ScheduleTimezone** | Pointer to [**ScheduleTimezone**](ScheduleTimezone.md) |  | [optional] [default to SCHEDULE_TIMEZONE_UTC]
**VaultStorageTier** | Pointer to [**VaultStorageTier**](VaultStorageTier.md) |  | [optional] 

## Methods

### NewStandardBackupPolicyPlan

`func NewStandardBackupPolicyPlan(backupSchedules []StandardBackupSchedules, ) *StandardBackupPolicyPlan`

NewStandardBackupPolicyPlan instantiates a new StandardBackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardBackupPolicyPlanWithDefaults

`func NewStandardBackupPolicyPlanWithDefaults() *StandardBackupPolicyPlan`

NewStandardBackupPolicyPlanWithDefaults instantiates a new StandardBackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceTypes

`func (o *StandardBackupPolicyPlan) GetResourceTypes() []ResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *StandardBackupPolicyPlan) GetResourceTypesOk() (*[]ResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *StandardBackupPolicyPlan) SetResourceTypes(v []ResourceType)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *StandardBackupPolicyPlan) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### GetBackupSchedules

`func (o *StandardBackupPolicyPlan) GetBackupSchedules() []StandardBackupSchedules`

GetBackupSchedules returns the BackupSchedules field if non-nil, zero value otherwise.

### GetBackupSchedulesOk

`func (o *StandardBackupPolicyPlan) GetBackupSchedulesOk() (*[]StandardBackupSchedules, bool)`

GetBackupSchedulesOk returns a tuple with the BackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedules

`func (o *StandardBackupPolicyPlan) SetBackupSchedules(v []StandardBackupSchedules)`

SetBackupSchedules sets BackupSchedules field to given value.


### GetScheduleTimezone

`func (o *StandardBackupPolicyPlan) GetScheduleTimezone() ScheduleTimezone`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *StandardBackupPolicyPlan) GetScheduleTimezoneOk() (*ScheduleTimezone, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *StandardBackupPolicyPlan) SetScheduleTimezone(v ScheduleTimezone)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *StandardBackupPolicyPlan) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetVaultStorageTier

`func (o *StandardBackupPolicyPlan) GetVaultStorageTier() VaultStorageTier`

GetVaultStorageTier returns the VaultStorageTier field if non-nil, zero value otherwise.

### GetVaultStorageTierOk

`func (o *StandardBackupPolicyPlan) GetVaultStorageTierOk() (*VaultStorageTier, bool)`

GetVaultStorageTierOk returns a tuple with the VaultStorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultStorageTier

`func (o *StandardBackupPolicyPlan) SetVaultStorageTier(v VaultStorageTier)`

SetVaultStorageTier sets VaultStorageTier field to given value.

### HasVaultStorageTier

`func (o *StandardBackupPolicyPlan) HasVaultStorageTier() bool`

HasVaultStorageTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


