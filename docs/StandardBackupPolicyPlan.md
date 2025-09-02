# StandardBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSchedules** | [**[]StandardBackupSchedules**](StandardBackupSchedules.md) | List of backup schedules. Each schedule specifies a backup frequency, retention period, and vault.  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


