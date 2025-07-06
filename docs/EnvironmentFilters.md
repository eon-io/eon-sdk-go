# EnvironmentFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]Environment**](Environment.md) | Matches if any value in this list equals &#x60;environment&#x60;. | [optional] 
**NotIn** | Pointer to [**[]Environment**](Environment.md) | Matches if no value in this list equals &#x60;environment&#x60;. | [optional] 

## Methods

### NewEnvironmentFilters

`func NewEnvironmentFilters() *EnvironmentFilters`

NewEnvironmentFilters instantiates a new EnvironmentFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentFiltersWithDefaults

`func NewEnvironmentFiltersWithDefaults() *EnvironmentFilters`

NewEnvironmentFiltersWithDefaults instantiates a new EnvironmentFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *EnvironmentFilters) GetIn() []Environment`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *EnvironmentFilters) GetInOk() (*[]Environment, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *EnvironmentFilters) SetIn(v []Environment)`

SetIn sets In field to given value.

### HasIn

`func (o *EnvironmentFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *EnvironmentFilters) GetNotIn() []Environment`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *EnvironmentFilters) GetNotInOk() (*[]Environment, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *EnvironmentFilters) SetNotIn(v []Environment)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *EnvironmentFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


