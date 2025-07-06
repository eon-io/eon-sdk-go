# BackupStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]BackupStatus**](BackupStatus.md) | Matches if any value in this list equals &#x60;backupStatus&#x60;. | [optional] 
**NotIn** | Pointer to [**[]BackupStatus**](BackupStatus.md) | Matches if no value in this list equals &#x60;backupStatus&#x60;. | [optional] 

## Methods

### NewBackupStatusFilters

`func NewBackupStatusFilters() *BackupStatusFilters`

NewBackupStatusFilters instantiates a new BackupStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStatusFiltersWithDefaults

`func NewBackupStatusFiltersWithDefaults() *BackupStatusFilters`

NewBackupStatusFiltersWithDefaults instantiates a new BackupStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *BackupStatusFilters) GetIn() []BackupStatus`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *BackupStatusFilters) GetInOk() (*[]BackupStatus, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *BackupStatusFilters) SetIn(v []BackupStatus)`

SetIn sets In field to given value.

### HasIn

`func (o *BackupStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *BackupStatusFilters) GetNotIn() []BackupStatus`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *BackupStatusFilters) GetNotInOk() (*[]BackupStatus, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *BackupStatusFilters) SetNotIn(v []BackupStatus)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *BackupStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


