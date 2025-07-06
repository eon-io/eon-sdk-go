# ListBackupJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**BackupJobsFilterConditions**](BackupJobsFilterConditions.md) |  | [optional] 
**Sorts** | Pointer to [**[]SortBackupJobBy**](SortBackupJobBy.md) |  | [optional] 

## Methods

### NewListBackupJobsRequest

`func NewListBackupJobsRequest() *ListBackupJobsRequest`

NewListBackupJobsRequest instantiates a new ListBackupJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupJobsRequestWithDefaults

`func NewListBackupJobsRequestWithDefaults() *ListBackupJobsRequest`

NewListBackupJobsRequestWithDefaults instantiates a new ListBackupJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListBackupJobsRequest) GetFilters() BackupJobsFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListBackupJobsRequest) GetFiltersOk() (*BackupJobsFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListBackupJobsRequest) SetFilters(v BackupJobsFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListBackupJobsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorts

`func (o *ListBackupJobsRequest) GetSorts() []SortBackupJobBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListBackupJobsRequest) GetSortsOk() (*[]SortBackupJobBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListBackupJobsRequest) SetSorts(v []SortBackupJobBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListBackupJobsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


