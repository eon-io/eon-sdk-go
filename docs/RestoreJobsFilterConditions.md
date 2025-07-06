# RestoreJobsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to [**JobIdFilters**](JobIdFilters.md) |  | [optional] 
**SnapshotId** | Pointer to [**SnapshotIdFilters**](SnapshotIdFilters.md) |  | [optional] 
**ProviderResourceId** | Pointer to [**ResourceIdFilters**](ResourceIdFilters.md) |  | [optional] 
**RestoreType** | Pointer to [**RestoreJobTypeFilters**](RestoreJobTypeFilters.md) |  | [optional] 

## Methods

### NewRestoreJobsFilterConditions

`func NewRestoreJobsFilterConditions() *RestoreJobsFilterConditions`

NewRestoreJobsFilterConditions instantiates a new RestoreJobsFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobsFilterConditionsWithDefaults

`func NewRestoreJobsFilterConditionsWithDefaults() *RestoreJobsFilterConditions`

NewRestoreJobsFilterConditionsWithDefaults instantiates a new RestoreJobsFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *RestoreJobsFilterConditions) GetJobId() JobIdFilters`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RestoreJobsFilterConditions) GetJobIdOk() (*JobIdFilters, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RestoreJobsFilterConditions) SetJobId(v JobIdFilters)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RestoreJobsFilterConditions) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *RestoreJobsFilterConditions) GetSnapshotId() SnapshotIdFilters`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreJobsFilterConditions) GetSnapshotIdOk() (*SnapshotIdFilters, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreJobsFilterConditions) SetSnapshotId(v SnapshotIdFilters)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *RestoreJobsFilterConditions) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetProviderResourceId

`func (o *RestoreJobsFilterConditions) GetProviderResourceId() ResourceIdFilters`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *RestoreJobsFilterConditions) GetProviderResourceIdOk() (*ResourceIdFilters, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *RestoreJobsFilterConditions) SetProviderResourceId(v ResourceIdFilters)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *RestoreJobsFilterConditions) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### GetRestoreType

`func (o *RestoreJobsFilterConditions) GetRestoreType() RestoreJobTypeFilters`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *RestoreJobsFilterConditions) GetRestoreTypeOk() (*RestoreJobTypeFilters, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *RestoreJobsFilterConditions) SetRestoreType(v RestoreJobTypeFilters)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *RestoreJobsFilterConditions) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


