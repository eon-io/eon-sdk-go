# ListInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sorts** | Pointer to [**[]SortResourceBy**](SortResourceBy.md) | List of sorting options. Sorting is applied in the order passed in the list.  | [optional] 
**Filters** | Pointer to [**InventoryFilterConditions**](InventoryFilterConditions.md) |  | [optional] 

## Methods

### NewListInventoryRequest

`func NewListInventoryRequest() *ListInventoryRequest`

NewListInventoryRequest instantiates a new ListInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInventoryRequestWithDefaults

`func NewListInventoryRequestWithDefaults() *ListInventoryRequest`

NewListInventoryRequestWithDefaults instantiates a new ListInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSorts

`func (o *ListInventoryRequest) GetSorts() []SortResourceBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListInventoryRequest) GetSortsOk() (*[]SortResourceBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListInventoryRequest) SetSorts(v []SortResourceBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListInventoryRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetFilters

`func (o *ListInventoryRequest) GetFilters() InventoryFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListInventoryRequest) GetFiltersOk() (*InventoryFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListInventoryRequest) SetFilters(v InventoryFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListInventoryRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


