# CreateBackupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Backup policy display name. | 
**Enabled** | Pointer to **bool** | Whether the backup policy is enabled. | [optional] [default to true]
**ResourceSelector** | [**BackupPolicyResourceSelector**](BackupPolicyResourceSelector.md) |  | 
**BackupPlan** | [**BackupPolicyPlan**](BackupPolicyPlan.md) |  | 

## Methods

### NewCreateBackupPolicyRequest

`func NewCreateBackupPolicyRequest(name string, resourceSelector BackupPolicyResourceSelector, backupPlan BackupPolicyPlan, ) *CreateBackupPolicyRequest`

NewCreateBackupPolicyRequest instantiates a new CreateBackupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupPolicyRequestWithDefaults

`func NewCreateBackupPolicyRequestWithDefaults() *CreateBackupPolicyRequest`

NewCreateBackupPolicyRequestWithDefaults instantiates a new CreateBackupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBackupPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBackupPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBackupPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateBackupPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateBackupPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateBackupPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateBackupPolicyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetResourceSelector

`func (o *CreateBackupPolicyRequest) GetResourceSelector() BackupPolicyResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *CreateBackupPolicyRequest) GetResourceSelectorOk() (*BackupPolicyResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *CreateBackupPolicyRequest) SetResourceSelector(v BackupPolicyResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### GetBackupPlan

`func (o *CreateBackupPolicyRequest) GetBackupPlan() BackupPolicyPlan`

GetBackupPlan returns the BackupPlan field if non-nil, zero value otherwise.

### GetBackupPlanOk

`func (o *CreateBackupPolicyRequest) GetBackupPlanOk() (*BackupPolicyPlan, bool)`

GetBackupPlanOk returns a tuple with the BackupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPlan

`func (o *CreateBackupPolicyRequest) SetBackupPlan(v BackupPolicyPlan)`

SetBackupPlan sets BackupPlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


