# FindingExclusionValueFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | Pointer to **[]string** | Matches exclusions whose &#x60;value&#x60; contains every one of these substrings. An empty list applies no filter. | [optional] 

## Methods

### NewFindingExclusionValueFilters

`func NewFindingExclusionValueFilters() *FindingExclusionValueFilters`

NewFindingExclusionValueFilters instantiates a new FindingExclusionValueFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionValueFiltersWithDefaults

`func NewFindingExclusionValueFiltersWithDefaults() *FindingExclusionValueFilters`

NewFindingExclusionValueFiltersWithDefaults instantiates a new FindingExclusionValueFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *FindingExclusionValueFilters) GetContains() []string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *FindingExclusionValueFilters) GetContainsOk() (*[]string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *FindingExclusionValueFilters) SetContains(v []string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *FindingExclusionValueFilters) HasContains() bool`

HasContains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


