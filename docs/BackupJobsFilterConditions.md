# BackupJobsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to [**JobIdFilters**](JobIdFilters.md) |  | [optional] 
**SnapshotId** | Pointer to [**SnapshotIdFilters**](SnapshotIdFilters.md) |  | [optional] 
**Status** | Pointer to [**JobStatusFilters**](JobStatusFilters.md) |  | [optional] 
**ProviderResourceId** | Pointer to [**ResourceIdFilters**](ResourceIdFilters.md) |  | [optional] 
**BackupType** | Pointer to [**BackupJobTypeFilters**](BackupJobTypeFilters.md) |  | [optional] 
**StartTime** | Pointer to [**StartTimeDateFilters**](StartTimeDateFilters.md) |  | [optional] 
**EndTime** | Pointer to [**EndTimeDateFilters**](EndTimeDateFilters.md) |  | [optional] 

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

### GetStatus

`func (o *BackupJobsFilterConditions) GetStatus() JobStatusFilters`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupJobsFilterConditions) GetStatusOk() (*JobStatusFilters, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupJobsFilterConditions) SetStatus(v JobStatusFilters)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupJobsFilterConditions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetStartTime

`func (o *BackupJobsFilterConditions) GetStartTime() StartTimeDateFilters`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupJobsFilterConditions) GetStartTimeOk() (*StartTimeDateFilters, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupJobsFilterConditions) SetStartTime(v StartTimeDateFilters)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupJobsFilterConditions) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *BackupJobsFilterConditions) GetEndTime() EndTimeDateFilters`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BackupJobsFilterConditions) GetEndTimeOk() (*EndTimeDateFilters, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BackupJobsFilterConditions) SetEndTime(v EndTimeDateFilters)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BackupJobsFilterConditions) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


