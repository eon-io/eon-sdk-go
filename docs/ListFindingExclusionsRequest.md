# ListFindingExclusionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**NullableFindingExclusionFilterConditions**](FindingExclusionFilterConditions.md) |  | [optional] 

## Methods

### NewListFindingExclusionsRequest

`func NewListFindingExclusionsRequest() *ListFindingExclusionsRequest`

NewListFindingExclusionsRequest instantiates a new ListFindingExclusionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFindingExclusionsRequestWithDefaults

`func NewListFindingExclusionsRequestWithDefaults() *ListFindingExclusionsRequest`

NewListFindingExclusionsRequestWithDefaults instantiates a new ListFindingExclusionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListFindingExclusionsRequest) GetFilters() FindingExclusionFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListFindingExclusionsRequest) GetFiltersOk() (*FindingExclusionFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListFindingExclusionsRequest) SetFilters(v FindingExclusionFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListFindingExclusionsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *ListFindingExclusionsRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ListFindingExclusionsRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


