# UpdateBackupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Backup policy display name. | 
**Enabled** | Pointer to **bool** | Whether the policy is enabled. | [optional] [default to true]
**ResourceSelector** | [**BackupPolicyResourceSelector**](BackupPolicyResourceSelector.md) |  | 
**BackupPlan** | [**BackupPolicyPlan**](BackupPolicyPlan.md) |  | 
**Tags** | Pointer to **map[string]string** | User-defined metadata tags, used for filtering and search only. They do not affect backup behavior, access control, or the tags Eon sets on your cloud resources.  | [optional] 

## Methods

### NewUpdateBackupPolicyRequest

`func NewUpdateBackupPolicyRequest(name string, resourceSelector BackupPolicyResourceSelector, backupPlan BackupPolicyPlan, ) *UpdateBackupPolicyRequest`

NewUpdateBackupPolicyRequest instantiates a new UpdateBackupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupPolicyRequestWithDefaults

`func NewUpdateBackupPolicyRequestWithDefaults() *UpdateBackupPolicyRequest`

NewUpdateBackupPolicyRequestWithDefaults instantiates a new UpdateBackupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBackupPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBackupPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBackupPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UpdateBackupPolicyRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBackupPolicyRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBackupPolicyRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBackupPolicyRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetResourceSelector

`func (o *UpdateBackupPolicyRequest) GetResourceSelector() BackupPolicyResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *UpdateBackupPolicyRequest) GetResourceSelectorOk() (*BackupPolicyResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *UpdateBackupPolicyRequest) SetResourceSelector(v BackupPolicyResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### GetBackupPlan

`func (o *UpdateBackupPolicyRequest) GetBackupPlan() BackupPolicyPlan`

GetBackupPlan returns the BackupPlan field if non-nil, zero value otherwise.

### GetBackupPlanOk

`func (o *UpdateBackupPolicyRequest) GetBackupPlanOk() (*BackupPolicyPlan, bool)`

GetBackupPlanOk returns a tuple with the BackupPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPlan

`func (o *UpdateBackupPolicyRequest) SetBackupPlan(v BackupPolicyPlan)`

SetBackupPlan sets BackupPlan field to given value.


### GetTags

`func (o *UpdateBackupPolicyRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateBackupPolicyRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateBackupPolicyRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateBackupPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


