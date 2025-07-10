# ListInventorySnapshotsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**SnapshotFilterConditions**](SnapshotFilterConditions.md) |  | [optional] 
**Sorts** | Pointer to [**[]SortSnapshotsBy**](SortSnapshotsBy.md) | List of sorting options. Sorting is applied in the order passed in the list.  | [optional] 

## Methods

### NewListInventorySnapshotsRequest

`func NewListInventorySnapshotsRequest() *ListInventorySnapshotsRequest`

NewListInventorySnapshotsRequest instantiates a new ListInventorySnapshotsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInventorySnapshotsRequestWithDefaults

`func NewListInventorySnapshotsRequestWithDefaults() *ListInventorySnapshotsRequest`

NewListInventorySnapshotsRequestWithDefaults instantiates a new ListInventorySnapshotsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListInventorySnapshotsRequest) GetFilters() SnapshotFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListInventorySnapshotsRequest) GetFiltersOk() (*SnapshotFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListInventorySnapshotsRequest) SetFilters(v SnapshotFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListInventorySnapshotsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSorts

`func (o *ListInventorySnapshotsRequest) GetSorts() []SortSnapshotsBy`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListInventorySnapshotsRequest) GetSortsOk() (*[]SortSnapshotsBy, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListInventorySnapshotsRequest) SetSorts(v []SortSnapshotsBy)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListInventorySnapshotsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


