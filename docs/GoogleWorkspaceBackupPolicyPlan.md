# GoogleWorkspaceBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | **string** | Vault the Google Workspace backups are written to. Must be a GCP vault.  | 
**BackupRetentionDays** | **int32** | Backup retention period, in days.  | 

## Methods

### NewGoogleWorkspaceBackupPolicyPlan

`func NewGoogleWorkspaceBackupPolicyPlan(vaultId string, backupRetentionDays int32, ) *GoogleWorkspaceBackupPolicyPlan`

NewGoogleWorkspaceBackupPolicyPlan instantiates a new GoogleWorkspaceBackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceBackupPolicyPlanWithDefaults

`func NewGoogleWorkspaceBackupPolicyPlanWithDefaults() *GoogleWorkspaceBackupPolicyPlan`

NewGoogleWorkspaceBackupPolicyPlanWithDefaults instantiates a new GoogleWorkspaceBackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *GoogleWorkspaceBackupPolicyPlan) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *GoogleWorkspaceBackupPolicyPlan) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *GoogleWorkspaceBackupPolicyPlan) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetBackupRetentionDays

`func (o *GoogleWorkspaceBackupPolicyPlan) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *GoogleWorkspaceBackupPolicyPlan) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *GoogleWorkspaceBackupPolicyPlan) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


