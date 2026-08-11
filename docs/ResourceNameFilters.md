# ResourceNameFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches if any string in this list equals &#x60;resourceName&#x60;. | [optional] 
**NotIn** | Pointer to **[]string** | Matches if no string in this list equals &#x60;resourceName&#x60;. | [optional] 
**Contains** | Pointer to **[]string** | Matches if any string in this list is a substring of &#x60;resourceName&#x60;. | [optional] 
**NotContains** | Pointer to **[]string** | Matches if no string in this list is a substring of &#x60;resourceName&#x60;. | [optional] 
**StartsWith** | Pointer to **[]string** | Matches if any string in this list is a prefix of &#x60;resourceName&#x60;. | [optional] 
**NotStartsWith** | Pointer to **[]string** | Matches if no string in this list is a prefix of &#x60;resourceName&#x60;. | [optional] 
**EndsWith** | Pointer to **[]string** | Matches if any string in this list is a suffix of &#x60;resourceName&#x60;. | [optional] 
**NotEndsWith** | Pointer to **[]string** | Matches if no string in this list is a suffix of &#x60;resourceName&#x60;. | [optional] 

## Methods

### NewResourceNameFilters

`func NewResourceNameFilters() *ResourceNameFilters`

NewResourceNameFilters instantiates a new ResourceNameFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNameFiltersWithDefaults

`func NewResourceNameFiltersWithDefaults() *ResourceNameFilters`

NewResourceNameFiltersWithDefaults instantiates a new ResourceNameFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *ResourceNameFilters) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ResourceNameFilters) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ResourceNameFilters) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *ResourceNameFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *ResourceNameFilters) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *ResourceNameFilters) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *ResourceNameFilters) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *ResourceNameFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.

### GetContains

`func (o *ResourceNameFilters) GetContains() []string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *ResourceNameFilters) GetContainsOk() (*[]string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *ResourceNameFilters) SetContains(v []string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *ResourceNameFilters) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetNotContains

`func (o *ResourceNameFilters) GetNotContains() []string`

GetNotContains returns the NotContains field if non-nil, zero value otherwise.

### GetNotContainsOk

`func (o *ResourceNameFilters) GetNotContainsOk() (*[]string, bool)`

GetNotContainsOk returns a tuple with the NotContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotContains

`func (o *ResourceNameFilters) SetNotContains(v []string)`

SetNotContains sets NotContains field to given value.

### HasNotContains

`func (o *ResourceNameFilters) HasNotContains() bool`

HasNotContains returns a boolean if a field has been set.

### GetStartsWith

`func (o *ResourceNameFilters) GetStartsWith() []string`

GetStartsWith returns the StartsWith field if non-nil, zero value otherwise.

### GetStartsWithOk

`func (o *ResourceNameFilters) GetStartsWithOk() (*[]string, bool)`

GetStartsWithOk returns a tuple with the StartsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsWith

`func (o *ResourceNameFilters) SetStartsWith(v []string)`

SetStartsWith sets StartsWith field to given value.

### HasStartsWith

`func (o *ResourceNameFilters) HasStartsWith() bool`

HasStartsWith returns a boolean if a field has been set.

### GetNotStartsWith

`func (o *ResourceNameFilters) GetNotStartsWith() []string`

GetNotStartsWith returns the NotStartsWith field if non-nil, zero value otherwise.

### GetNotStartsWithOk

`func (o *ResourceNameFilters) GetNotStartsWithOk() (*[]string, bool)`

GetNotStartsWithOk returns a tuple with the NotStartsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStartsWith

`func (o *ResourceNameFilters) SetNotStartsWith(v []string)`

SetNotStartsWith sets NotStartsWith field to given value.

### HasNotStartsWith

`func (o *ResourceNameFilters) HasNotStartsWith() bool`

HasNotStartsWith returns a boolean if a field has been set.

### GetEndsWith

`func (o *ResourceNameFilters) GetEndsWith() []string`

GetEndsWith returns the EndsWith field if non-nil, zero value otherwise.

### GetEndsWithOk

`func (o *ResourceNameFilters) GetEndsWithOk() (*[]string, bool)`

GetEndsWithOk returns a tuple with the EndsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsWith

`func (o *ResourceNameFilters) SetEndsWith(v []string)`

SetEndsWith sets EndsWith field to given value.

### HasEndsWith

`func (o *ResourceNameFilters) HasEndsWith() bool`

HasEndsWith returns a boolean if a field has been set.

### GetNotEndsWith

`func (o *ResourceNameFilters) GetNotEndsWith() []string`

GetNotEndsWith returns the NotEndsWith field if non-nil, zero value otherwise.

### GetNotEndsWithOk

`func (o *ResourceNameFilters) GetNotEndsWithOk() (*[]string, bool)`

GetNotEndsWithOk returns a tuple with the NotEndsWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEndsWith

`func (o *ResourceNameFilters) SetNotEndsWith(v []string)`

SetNotEndsWith sets NotEndsWith field to given value.

### HasNotEndsWith

`func (o *ResourceNameFilters) HasNotEndsWith() bool`

HasNotEndsWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


