# BackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicyType** | [**BackupPolicyType**](BackupPolicyType.md) |  | 
**StandardPlan** | Pointer to [**NullableStandardBackupPolicyPlan**](StandardBackupPolicyPlan.md) |  | [optional] 
**HighFrequencyPlan** | Pointer to [**NullableHighFrequencyBackupPolicyPlan**](HighFrequencyBackupPolicyPlan.md) |  | [optional] 

## Methods

### NewBackupPolicyPlan

`func NewBackupPolicyPlan(backupPolicyType BackupPolicyType, ) *BackupPolicyPlan`

NewBackupPolicyPlan instantiates a new BackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyPlanWithDefaults

`func NewBackupPolicyPlanWithDefaults() *BackupPolicyPlan`

NewBackupPolicyPlanWithDefaults instantiates a new BackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicyType

`func (o *BackupPolicyPlan) GetBackupPolicyType() BackupPolicyType`

GetBackupPolicyType returns the BackupPolicyType field if non-nil, zero value otherwise.

### GetBackupPolicyTypeOk

`func (o *BackupPolicyPlan) GetBackupPolicyTypeOk() (*BackupPolicyType, bool)`

GetBackupPolicyTypeOk returns a tuple with the BackupPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicyType

`func (o *BackupPolicyPlan) SetBackupPolicyType(v BackupPolicyType)`

SetBackupPolicyType sets BackupPolicyType field to given value.


### GetStandardPlan

`func (o *BackupPolicyPlan) GetStandardPlan() StandardBackupPolicyPlan`

GetStandardPlan returns the StandardPlan field if non-nil, zero value otherwise.

### GetStandardPlanOk

`func (o *BackupPolicyPlan) GetStandardPlanOk() (*StandardBackupPolicyPlan, bool)`

GetStandardPlanOk returns a tuple with the StandardPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPlan

`func (o *BackupPolicyPlan) SetStandardPlan(v StandardBackupPolicyPlan)`

SetStandardPlan sets StandardPlan field to given value.

### HasStandardPlan

`func (o *BackupPolicyPlan) HasStandardPlan() bool`

HasStandardPlan returns a boolean if a field has been set.

### SetStandardPlanNil

`func (o *BackupPolicyPlan) SetStandardPlanNil(b bool)`

 SetStandardPlanNil sets the value for StandardPlan to be an explicit nil

### UnsetStandardPlan
`func (o *BackupPolicyPlan) UnsetStandardPlan()`

UnsetStandardPlan ensures that no value is present for StandardPlan, not even an explicit nil
### GetHighFrequencyPlan

`func (o *BackupPolicyPlan) GetHighFrequencyPlan() HighFrequencyBackupPolicyPlan`

GetHighFrequencyPlan returns the HighFrequencyPlan field if non-nil, zero value otherwise.

### GetHighFrequencyPlanOk

`func (o *BackupPolicyPlan) GetHighFrequencyPlanOk() (*HighFrequencyBackupPolicyPlan, bool)`

GetHighFrequencyPlanOk returns a tuple with the HighFrequencyPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighFrequencyPlan

`func (o *BackupPolicyPlan) SetHighFrequencyPlan(v HighFrequencyBackupPolicyPlan)`

SetHighFrequencyPlan sets HighFrequencyPlan field to given value.

### HasHighFrequencyPlan

`func (o *BackupPolicyPlan) HasHighFrequencyPlan() bool`

HasHighFrequencyPlan returns a boolean if a field has been set.

### SetHighFrequencyPlanNil

`func (o *BackupPolicyPlan) SetHighFrequencyPlanNil(b bool)`

 SetHighFrequencyPlanNil sets the value for HighFrequencyPlan to be an explicit nil

### UnsetHighFrequencyPlan
`func (o *BackupPolicyPlan) UnsetHighFrequencyPlan()`

UnsetHighFrequencyPlan ensures that no value is present for HighFrequencyPlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


