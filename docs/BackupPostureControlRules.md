# BackupPostureControlRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumRetention** | Pointer to [**[]MinRetentionRule**](MinRetentionRule.md) | Minimum retention required per backup frequency (daily, weekly, monthly, annual). | [optional] 
**MaximumRetention** | Pointer to [**NullableMaxRetentionRule**](MaxRetentionRule.md) |  | [optional] 
**NumberOfCopies** | Pointer to [**NullableNumberOfCopiesRule**](NumberOfCopiesRule.md) |  | [optional] 
**CrossRegion** | Pointer to **bool** | Whether at least one backup copy must be stored in a different region. | [optional] 
**CrossAccount** | Pointer to **bool** | Whether at least one backup copy must be stored in a different cloud account. | [optional] 
**CrossCloudProvider** | Pointer to **bool** | Whether at least one backup copy must be stored in a different cloud provider. | [optional] 

## Methods

### NewBackupPostureControlRules

`func NewBackupPostureControlRules() *BackupPostureControlRules`

NewBackupPostureControlRules instantiates a new BackupPostureControlRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPostureControlRulesWithDefaults

`func NewBackupPostureControlRulesWithDefaults() *BackupPostureControlRules`

NewBackupPostureControlRulesWithDefaults instantiates a new BackupPostureControlRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumRetention

`func (o *BackupPostureControlRules) GetMinimumRetention() []MinRetentionRule`

GetMinimumRetention returns the MinimumRetention field if non-nil, zero value otherwise.

### GetMinimumRetentionOk

`func (o *BackupPostureControlRules) GetMinimumRetentionOk() (*[]MinRetentionRule, bool)`

GetMinimumRetentionOk returns a tuple with the MinimumRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRetention

`func (o *BackupPostureControlRules) SetMinimumRetention(v []MinRetentionRule)`

SetMinimumRetention sets MinimumRetention field to given value.

### HasMinimumRetention

`func (o *BackupPostureControlRules) HasMinimumRetention() bool`

HasMinimumRetention returns a boolean if a field has been set.

### GetMaximumRetention

`func (o *BackupPostureControlRules) GetMaximumRetention() MaxRetentionRule`

GetMaximumRetention returns the MaximumRetention field if non-nil, zero value otherwise.

### GetMaximumRetentionOk

`func (o *BackupPostureControlRules) GetMaximumRetentionOk() (*MaxRetentionRule, bool)`

GetMaximumRetentionOk returns a tuple with the MaximumRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRetention

`func (o *BackupPostureControlRules) SetMaximumRetention(v MaxRetentionRule)`

SetMaximumRetention sets MaximumRetention field to given value.

### HasMaximumRetention

`func (o *BackupPostureControlRules) HasMaximumRetention() bool`

HasMaximumRetention returns a boolean if a field has been set.

### SetMaximumRetentionNil

`func (o *BackupPostureControlRules) SetMaximumRetentionNil(b bool)`

 SetMaximumRetentionNil sets the value for MaximumRetention to be an explicit nil

### UnsetMaximumRetention
`func (o *BackupPostureControlRules) UnsetMaximumRetention()`

UnsetMaximumRetention ensures that no value is present for MaximumRetention, not even an explicit nil
### GetNumberOfCopies

`func (o *BackupPostureControlRules) GetNumberOfCopies() NumberOfCopiesRule`

GetNumberOfCopies returns the NumberOfCopies field if non-nil, zero value otherwise.

### GetNumberOfCopiesOk

`func (o *BackupPostureControlRules) GetNumberOfCopiesOk() (*NumberOfCopiesRule, bool)`

GetNumberOfCopiesOk returns a tuple with the NumberOfCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCopies

`func (o *BackupPostureControlRules) SetNumberOfCopies(v NumberOfCopiesRule)`

SetNumberOfCopies sets NumberOfCopies field to given value.

### HasNumberOfCopies

`func (o *BackupPostureControlRules) HasNumberOfCopies() bool`

HasNumberOfCopies returns a boolean if a field has been set.

### SetNumberOfCopiesNil

`func (o *BackupPostureControlRules) SetNumberOfCopiesNil(b bool)`

 SetNumberOfCopiesNil sets the value for NumberOfCopies to be an explicit nil

### UnsetNumberOfCopies
`func (o *BackupPostureControlRules) UnsetNumberOfCopies()`

UnsetNumberOfCopies ensures that no value is present for NumberOfCopies, not even an explicit nil
### GetCrossRegion

`func (o *BackupPostureControlRules) GetCrossRegion() bool`

GetCrossRegion returns the CrossRegion field if non-nil, zero value otherwise.

### GetCrossRegionOk

`func (o *BackupPostureControlRules) GetCrossRegionOk() (*bool, bool)`

GetCrossRegionOk returns a tuple with the CrossRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossRegion

`func (o *BackupPostureControlRules) SetCrossRegion(v bool)`

SetCrossRegion sets CrossRegion field to given value.

### HasCrossRegion

`func (o *BackupPostureControlRules) HasCrossRegion() bool`

HasCrossRegion returns a boolean if a field has been set.

### GetCrossAccount

`func (o *BackupPostureControlRules) GetCrossAccount() bool`

GetCrossAccount returns the CrossAccount field if non-nil, zero value otherwise.

### GetCrossAccountOk

`func (o *BackupPostureControlRules) GetCrossAccountOk() (*bool, bool)`

GetCrossAccountOk returns a tuple with the CrossAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossAccount

`func (o *BackupPostureControlRules) SetCrossAccount(v bool)`

SetCrossAccount sets CrossAccount field to given value.

### HasCrossAccount

`func (o *BackupPostureControlRules) HasCrossAccount() bool`

HasCrossAccount returns a boolean if a field has been set.

### GetCrossCloudProvider

`func (o *BackupPostureControlRules) GetCrossCloudProvider() bool`

GetCrossCloudProvider returns the CrossCloudProvider field if non-nil, zero value otherwise.

### GetCrossCloudProviderOk

`func (o *BackupPostureControlRules) GetCrossCloudProviderOk() (*bool, bool)`

GetCrossCloudProviderOk returns a tuple with the CrossCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudProvider

`func (o *BackupPostureControlRules) SetCrossCloudProvider(v bool)`

SetCrossCloudProvider sets CrossCloudProvider field to given value.

### HasCrossCloudProvider

`func (o *BackupPostureControlRules) HasCrossCloudProvider() bool`

HasCrossCloudProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


