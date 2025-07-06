# BackupJobsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to [**JobIdFilters**](JobIdFilters.md) |  | [optional] 
**SnapshotId** | Pointer to [**SnapshotIdFilters**](SnapshotIdFilters.md) |  | [optional] 
**ProviderResourceId** | Pointer to [**ResourceIdFilters**](ResourceIdFilters.md) |  | [optional] 
**BackupType** | Pointer to [**BackupJobTypeFilters**](BackupJobTypeFilters.md) |  | [optional] 

## Methods

### NewBackupJobsFilterConditions

`func NewBackupJobsFilterConditions() *BackupJobsFilterConditions`

NewBackupJobsFilterConditions instantiates a new BackupJobsFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobsFilterConditionsWithDefaults

`func NewBackupJobsFilterConditionsWithDefaults() *BackupJobsFilterConditions`

NewBackupJobsFilterConditionsWithDefaults instantiates a new BackupJobsFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *BackupJobsFilterConditions) GetJobId() JobIdFilters`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupJobsFilterConditions) GetJobIdOk() (*JobIdFilters, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupJobsFilterConditions) SetJobId(v JobIdFilters)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupJobsFilterConditions) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *BackupJobsFilterConditions) GetSnapshotId() SnapshotIdFilters`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *BackupJobsFilterConditions) GetSnapshotIdOk() (*SnapshotIdFilters, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *BackupJobsFilterConditions) SetSnapshotId(v SnapshotIdFilters)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *BackupJobsFilterConditions) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetProviderResourceId

`func (o *BackupJobsFilterConditions) GetProviderResourceId() ResourceIdFilters`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *BackupJobsFilterConditions) GetProviderResourceIdOk() (*ResourceIdFilters, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *BackupJobsFilterConditions) SetProviderResourceId(v ResourceIdFilters)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *BackupJobsFilterConditions) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupJobsFilterConditions) GetBackupType() BackupJobTypeFilters`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupJobsFilterConditions) GetBackupTypeOk() (*BackupJobTypeFilters, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupJobsFilterConditions) SetBackupType(v BackupJobTypeFilters)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *BackupJobsFilterConditions) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


