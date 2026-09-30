# FindingExclusionScopeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]FindingExclusionScope**](FindingExclusionScope.md) | Matches exclusions with any of these scopes. An empty list applies no filter. | [optional] 
**NotIn** | Pointer to [**[]FindingExclusionScope**](FindingExclusionScope.md) | Matches exclusions with none of these scopes. An empty list applies no filter. | [optional] 

## Methods

### NewFindingExclusionScopeFilters

`func NewFindingExclusionScopeFilters() *FindingExclusionScopeFilters`

NewFindingExclusionScopeFilters instantiates a new FindingExclusionScopeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionScopeFiltersWithDefaults

`func NewFindingExclusionScopeFiltersWithDefaults() *FindingExclusionScopeFilters`

NewFindingExclusionScopeFiltersWithDefaults instantiates a new FindingExclusionScopeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *FindingExclusionScopeFilters) GetIn() []FindingExclusionScope`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *FindingExclusionScopeFilters) GetInOk() (*[]FindingExclusionScope, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *FindingExclusionScopeFilters) SetIn(v []FindingExclusionScope)`

SetIn sets In field to given value.

### HasIn

`func (o *FindingExclusionScopeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *FindingExclusionScopeFilters) GetNotIn() []FindingExclusionScope`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *FindingExclusionScopeFilters) GetNotInOk() (*[]FindingExclusionScope, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *FindingExclusionScopeFilters) SetNotIn(v []FindingExclusionScope)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *FindingExclusionScopeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


