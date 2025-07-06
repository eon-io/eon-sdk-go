# TagKeyValuesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainsAllOf** | Pointer to **[]string** | Matches if any &#x60;{key}&#x3D;{value}&#x60; pair in this list is in the &#x60;tags&#x60; list. | [optional] 
**ContainsAnyOf** | Pointer to **[]string** | Matches if none of the &#x60;{key}&#x3D;{value}&#x60; pairs in this list are in the &#x60;tags&#x60; list. | [optional] 
**ContainsNoneOf** | Pointer to **[]string** | Matches if all &#x60;{key}&#x3D;{value}&#x60; pairs in this list are in the &#x60;tags&#x60; list. | [optional] 

## Methods

### NewTagKeyValuesFilters

`func NewTagKeyValuesFilters() *TagKeyValuesFilters`

NewTagKeyValuesFilters instantiates a new TagKeyValuesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagKeyValuesFiltersWithDefaults

`func NewTagKeyValuesFiltersWithDefaults() *TagKeyValuesFilters`

NewTagKeyValuesFiltersWithDefaults instantiates a new TagKeyValuesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainsAllOf

`func (o *TagKeyValuesFilters) GetContainsAllOf() []string`

GetContainsAllOf returns the ContainsAllOf field if non-nil, zero value otherwise.

### GetContainsAllOfOk

`func (o *TagKeyValuesFilters) GetContainsAllOfOk() (*[]string, bool)`

GetContainsAllOfOk returns a tuple with the ContainsAllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAllOf

`func (o *TagKeyValuesFilters) SetContainsAllOf(v []string)`

SetContainsAllOf sets ContainsAllOf field to given value.

### HasContainsAllOf

`func (o *TagKeyValuesFilters) HasContainsAllOf() bool`

HasContainsAllOf returns a boolean if a field has been set.

### GetContainsAnyOf

`func (o *TagKeyValuesFilters) GetContainsAnyOf() []string`

GetContainsAnyOf returns the ContainsAnyOf field if non-nil, zero value otherwise.

### GetContainsAnyOfOk

`func (o *TagKeyValuesFilters) GetContainsAnyOfOk() (*[]string, bool)`

GetContainsAnyOfOk returns a tuple with the ContainsAnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAnyOf

`func (o *TagKeyValuesFilters) SetContainsAnyOf(v []string)`

SetContainsAnyOf sets ContainsAnyOf field to given value.

### HasContainsAnyOf

`func (o *TagKeyValuesFilters) HasContainsAnyOf() bool`

HasContainsAnyOf returns a boolean if a field has been set.

### GetContainsNoneOf

`func (o *TagKeyValuesFilters) GetContainsNoneOf() []string`

GetContainsNoneOf returns the ContainsNoneOf field if non-nil, zero value otherwise.

### GetContainsNoneOfOk

`func (o *TagKeyValuesFilters) GetContainsNoneOfOk() (*[]string, bool)`

GetContainsNoneOfOk returns a tuple with the ContainsNoneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsNoneOf

`func (o *TagKeyValuesFilters) SetContainsNoneOf(v []string)`

SetContainsNoneOf sets ContainsNoneOf field to given value.

### HasContainsNoneOf

`func (o *TagKeyValuesFilters) HasContainsNoneOf() bool`

HasContainsNoneOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


