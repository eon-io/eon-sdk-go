# ListSourceAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**SourceAccountsFilterConditions**](SourceAccountsFilterConditions.md) |  | [optional] 
**Sorts** | Pointer to [**[]SortSourceAccountsBy**](SortSourceAccountsBy.md) |  | [optional] 

## Methods

### NewListSourceAccountsRequest

`func NewListSourceAccountsRequest() *ListSourceAccountsRequest`

NewListSourceAccountsRequest instantiates a new ListSourceAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceAccountsRequestWithDefaults

`func NewListSourceAccountsRequestWithDefaults() *ListSourceAccountsRequest`

NewListSourceAccountsRequestWithDefaults instantiates a new ListSourceAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListSourceAccountsRequest) GetFilters() SourceAccountsFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListSourceAccountsRequest) GetFiltersOk() (*SourceAccountsFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListSourceAccountsRequest) SetFilters(v SourceAccountsFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListSourceAccountsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorts

`func (o *ListSourceAccountsRequest) GetSorts() []SortSourceAccountsBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListSourceAccountsRequest) GetSortsOk() (*[]SortSourceAccountsBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListSourceAccountsRequest) SetSorts(v []SortSourceAccountsBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListSourceAccountsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


