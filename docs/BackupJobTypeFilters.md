# BackupJobTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]BackupJobType**](BackupJobType.md) | Matches if any item in this list equals &#x60;backupType&#x60;. | [optional] 
**NotIn** | Pointer to [**[]BackupJobType**](BackupJobType.md) | Matches if no item in this list equals &#x60;backupType&#x60;. | [optional] 

## Methods

### NewBackupJobTypeFilters

`func NewBackupJobTypeFilters() *BackupJobTypeFilters`

NewBackupJobTypeFilters instantiates a new BackupJobTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobTypeFiltersWithDefaults

`func NewBackupJobTypeFiltersWithDefaults() *BackupJobTypeFilters`

NewBackupJobTypeFiltersWithDefaults instantiates a new BackupJobTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *BackupJobTypeFilters) GetIn() []BackupJobType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *BackupJobTypeFilters) GetInOk() (*[]BackupJobType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *BackupJobTypeFilters) SetIn(v []BackupJobType)`

SetIn sets In field to given value.

### HasIn

`func (o *BackupJobTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *BackupJobTypeFilters) GetNotIn() []BackupJobType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *BackupJobTypeFilters) GetNotInOk() (*[]BackupJobType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *BackupJobTypeFilters) SetNotIn(v []BackupJobType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *BackupJobTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


