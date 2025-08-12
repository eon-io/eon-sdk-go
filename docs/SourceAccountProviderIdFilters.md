# SourceAccountProviderIdFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches if any value in this list equals &#x60;sourceAccountProviderId&#x60;. | [optional] 
**NotIn** | Pointer to **[]string** | Matches if none of the values in this list equal &#x60;sourceAccountProviderId&#x60;. | [optional] 

## Methods

### NewSourceAccountProviderIdFilters

`func NewSourceAccountProviderIdFilters() *SourceAccountProviderIdFilters`

NewSourceAccountProviderIdFilters instantiates a new SourceAccountProviderIdFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountProviderIdFiltersWithDefaults

`func NewSourceAccountProviderIdFiltersWithDefaults() *SourceAccountProviderIdFilters`

NewSourceAccountProviderIdFiltersWithDefaults instantiates a new SourceAccountProviderIdFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *SourceAccountProviderIdFilters) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *SourceAccountProviderIdFilters) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *SourceAccountProviderIdFilters) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *SourceAccountProviderIdFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *SourceAccountProviderIdFilters) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *SourceAccountProviderIdFilters) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *SourceAccountProviderIdFilters) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *SourceAccountProviderIdFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


