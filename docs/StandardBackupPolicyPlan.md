# StandardBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTypes** | Pointer to [**[]ResourceType**](ResourceType.md) | Resource types to backup. | [optional] 
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


