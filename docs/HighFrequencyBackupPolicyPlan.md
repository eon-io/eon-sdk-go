# HighFrequencyBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTypes** | [**[]HighFrequencyBackupResourceType**](HighFrequencyBackupResourceType.md) | Resource types to backup. | 
**BackupSchedules** | [**[]HighFrequencyBackupSchedules**](HighFrequencyBackupSchedules.md) | List of backup schedules. Each schedule specifies a backup frequency, retention period, and vault.  | 
**ScheduleTimezone** | Pointer to [**ScheduleTimezone**](ScheduleTimezone.md) |  | [optional] [default to SCHEDULE_TIMEZONE_UTC]

## Methods

### NewHighFrequencyBackupPolicyPlan

`func NewHighFrequencyBackupPolicyPlan(resourceTypes []HighFrequencyBackupResourceType, backupSchedules []HighFrequencyBackupSchedules, ) *HighFrequencyBackupPolicyPlan`

NewHighFrequencyBackupPolicyPlan instantiates a new HighFrequencyBackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighFrequencyBackupPolicyPlanWithDefaults

`func NewHighFrequencyBackupPolicyPlanWithDefaults() *HighFrequencyBackupPolicyPlan`

NewHighFrequencyBackupPolicyPlanWithDefaults instantiates a new HighFrequencyBackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceTypes

`func (o *HighFrequencyBackupPolicyPlan) GetResourceTypes() []HighFrequencyBackupResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *HighFrequencyBackupPolicyPlan) GetResourceTypesOk() (*[]HighFrequencyBackupResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *HighFrequencyBackupPolicyPlan) SetResourceTypes(v []HighFrequencyBackupResourceType)`

SetResourceTypes sets ResourceTypes field to given value.


### GetBackupSchedules

`func (o *HighFrequencyBackupPolicyPlan) GetBackupSchedules() []HighFrequencyBackupSchedules`

GetBackupSchedules returns the BackupSchedules field if non-nil, zero value otherwise.

### GetBackupSchedulesOk

`func (o *HighFrequencyBackupPolicyPlan) GetBackupSchedulesOk() (*[]HighFrequencyBackupSchedules, bool)`

GetBackupSchedulesOk returns a tuple with the BackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedules

`func (o *HighFrequencyBackupPolicyPlan) SetBackupSchedules(v []HighFrequencyBackupSchedules)`

SetBackupSchedules sets BackupSchedules field to given value.


### GetScheduleTimezone

`func (o *HighFrequencyBackupPolicyPlan) GetScheduleTimezone() ScheduleTimezone`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *HighFrequencyBackupPolicyPlan) GetScheduleTimezoneOk() (*ScheduleTimezone, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *HighFrequencyBackupPolicyPlan) SetScheduleTimezone(v ScheduleTimezone)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *HighFrequencyBackupPolicyPlan) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


