# ListRestoreJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**RestoreJobsFilterConditions**](RestoreJobsFilterConditions.md) |  | [optional] 
**Sorts** | Pointer to [**[]SortRestoreJobBy**](SortRestoreJobBy.md) |  | [optional] 

## Methods

### NewListRestoreJobsRequest

`func NewListRestoreJobsRequest() *ListRestoreJobsRequest`

NewListRestoreJobsRequest instantiates a new ListRestoreJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestoreJobsRequestWithDefaults

`func NewListRestoreJobsRequestWithDefaults() *ListRestoreJobsRequest`

NewListRestoreJobsRequestWithDefaults instantiates a new ListRestoreJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListRestoreJobsRequest) GetFilters() RestoreJobsFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListRestoreJobsRequest) GetFiltersOk() (*RestoreJobsFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListRestoreJobsRequest) SetFilters(v RestoreJobsFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListRestoreJobsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorts

`func (o *ListRestoreJobsRequest) GetSorts() []SortRestoreJobBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListRestoreJobsRequest) GetSortsOk() (*[]SortRestoreJobBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListRestoreJobsRequest) SetSorts(v []SortRestoreJobBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListRestoreJobsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


