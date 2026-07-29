# AwsNativeStandardBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSchedules** | [**[]AwsNativeStandardBackupSchedules**](AwsNativeStandardBackupSchedules.md) | List of backup schedules. Each schedule specifies a backup frequency, retention period, and target region.  | 
**VaultProtectionMode** | Pointer to [**VaultProtectionMode**](VaultProtectionMode.md) |  | [optional] 

## Methods

### NewAwsNativeStandardBackupPolicyPlan

`func NewAwsNativeStandardBackupPolicyPlan(backupSchedules []AwsNativeStandardBackupSchedules, ) *AwsNativeStandardBackupPolicyPlan`

NewAwsNativeStandardBackupPolicyPlan instantiates a new AwsNativeStandardBackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativeStandardBackupPolicyPlanWithDefaults

`func NewAwsNativeStandardBackupPolicyPlanWithDefaults() *AwsNativeStandardBackupPolicyPlan`

NewAwsNativeStandardBackupPolicyPlanWithDefaults instantiates a new AwsNativeStandardBackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSchedules

`func (o *AwsNativeStandardBackupPolicyPlan) GetBackupSchedules() []AwsNativeStandardBackupSchedules`

GetBackupSchedules returns the BackupSchedules field if non-nil, zero value otherwise.

### GetBackupSchedulesOk

`func (o *AwsNativeStandardBackupPolicyPlan) GetBackupSchedulesOk() (*[]AwsNativeStandardBackupSchedules, bool)`

GetBackupSchedulesOk returns a tuple with the BackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedules

`func (o *AwsNativeStandardBackupPolicyPlan) SetBackupSchedules(v []AwsNativeStandardBackupSchedules)`

SetBackupSchedules sets BackupSchedules field to given value.


### GetVaultProtectionMode

`func (o *AwsNativeStandardBackupPolicyPlan) GetVaultProtectionMode() VaultProtectionMode`

GetVaultProtectionMode returns the VaultProtectionMode field if non-nil, zero value otherwise.

### GetVaultProtectionModeOk

`func (o *AwsNativeStandardBackupPolicyPlan) GetVaultProtectionModeOk() (*VaultProtectionMode, bool)`

GetVaultProtectionModeOk returns a tuple with the VaultProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultProtectionMode

`func (o *AwsNativeStandardBackupPolicyPlan) SetVaultProtectionMode(v VaultProtectionMode)`

SetVaultProtectionMode sets VaultProtectionMode field to given value.

### HasVaultProtectionMode

`func (o *AwsNativeStandardBackupPolicyPlan) HasVaultProtectionMode() bool`

HasVaultProtectionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


