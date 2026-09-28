# FindingExclusionTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]FindingObjectType**](FindingObjectType.md) | Matches exclusions of any of these types. An empty list applies no filter. | [optional] 
**NotIn** | Pointer to [**[]FindingObjectType**](FindingObjectType.md) | Matches exclusions of none of these types. An empty list applies no filter. | [optional] 

## Methods

### NewFindingExclusionTypeFilters

`func NewFindingExclusionTypeFilters() *FindingExclusionTypeFilters`

NewFindingExclusionTypeFilters instantiates a new FindingExclusionTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionTypeFiltersWithDefaults

`func NewFindingExclusionTypeFiltersWithDefaults() *FindingExclusionTypeFilters`

NewFindingExclusionTypeFiltersWithDefaults instantiates a new FindingExclusionTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *FindingExclusionTypeFilters) GetIn() []FindingObjectType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *FindingExclusionTypeFilters) GetInOk() (*[]FindingObjectType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *FindingExclusionTypeFilters) SetIn(v []FindingObjectType)`

SetIn sets In field to given value.

### HasIn

`func (o *FindingExclusionTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *FindingExclusionTypeFilters) GetNotIn() []FindingObjectType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *FindingExclusionTypeFilters) GetNotInOk() (*[]FindingObjectType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *FindingExclusionTypeFilters) SetNotIn(v []FindingObjectType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *FindingExclusionTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


