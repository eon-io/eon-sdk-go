# BackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Backup policy ID. | 
**Name** | **string** | Backup policy display name. | 
**Enabled** | **bool** | Whether the backup policy is enabled. | 
**ResourceSelector** | [**BackupPolicyResourceSelector**](BackupPolicyResourceSelector.md) |  | 
**BackupPlan** | [**BackupPolicyPlan**](BackupPolicyPlan.md) |  | 

## Methods

### NewBackupPolicy

`func NewBackupPolicy(id string, name string, enabled bool, resourceSelector BackupPolicyResourceSelector, backupPlan BackupPolicyPlan, ) *BackupPolicy`

NewBackupPolicy instantiates a new BackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyWithDefaults

`func NewBackupPolicyWithDefaults() *BackupPolicy`

NewBackupPolicyWithDefaults instantiates a new BackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BackupPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *BackupPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetResourceSelector

`func (o *BackupPolicy) GetResourceSelector() BackupPolicyResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *BackupPolicy) GetResourceSelectorOk() (*BackupPolicyResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *BackupPolicy) SetResourceSelector(v BackupPolicyResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### GetBackupPlan

`func (o *BackupPolicy) GetBackupPlan() BackupPolicyPlan`

GetBackupPlan returns the BackupPlan field if non-nil, zero value otherwise.

### GetBackupPlanOk

`func (o *BackupPolicy) GetBackupPlanOk() (*BackupPolicyPlan, bool)`

GetBackupPlanOk returns a tuple with the BackupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPlan

`func (o *BackupPolicy) SetBackupPlan(v BackupPolicyPlan)`

SetBackupPlan sets BackupPlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


