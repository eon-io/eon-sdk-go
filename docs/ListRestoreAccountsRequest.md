# ListRestoreAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**RestoreAccountsFilterConditions**](RestoreAccountsFilterConditions.md) |  | [optional] 
**Sorts** | Pointer to [**[]SortRestoreAccountsBy**](SortRestoreAccountsBy.md) |  | [optional] 

## Methods

### NewListRestoreAccountsRequest

`func NewListRestoreAccountsRequest() *ListRestoreAccountsRequest`

NewListRestoreAccountsRequest instantiates a new ListRestoreAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestoreAccountsRequestWithDefaults

`func NewListRestoreAccountsRequestWithDefaults() *ListRestoreAccountsRequest`

NewListRestoreAccountsRequestWithDefaults instantiates a new ListRestoreAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListRestoreAccountsRequest) GetFilters() RestoreAccountsFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListRestoreAccountsRequest) GetFiltersOk() (*RestoreAccountsFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListRestoreAccountsRequest) SetFilters(v RestoreAccountsFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListRestoreAccountsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorts

`func (o *ListRestoreAccountsRequest) GetSorts() []SortRestoreAccountsBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListRestoreAccountsRequest) GetSortsOk() (*[]SortRestoreAccountsBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListRestoreAccountsRequest) SetSorts(v []SortRestoreAccountsBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListRestoreAccountsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


