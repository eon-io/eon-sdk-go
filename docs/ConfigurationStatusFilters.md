# ConfigurationStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]ConfigurationStatus**](ConfigurationStatus.md) | Matches if any value in this list equals &#x60;configurationStatus&#x60;. | [optional] 
**NotIn** | Pointer to [**[]ConfigurationStatus**](ConfigurationStatus.md) | Matches if no value in this list equals &#x60;configurationStatus&#x60;. A resource the status recalculation has not reached yet has no configuration status, and matches every &#x60;notIn&#x60; list.  | [optional] 

## Methods

### NewConfigurationStatusFilters

`func NewConfigurationStatusFilters() *ConfigurationStatusFilters`

NewConfigurationStatusFilters instantiates a new ConfigurationStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationStatusFiltersWithDefaults

`func NewConfigurationStatusFiltersWithDefaults() *ConfigurationStatusFilters`

NewConfigurationStatusFiltersWithDefaults instantiates a new ConfigurationStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *ConfigurationStatusFilters) GetIn() []ConfigurationStatus`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ConfigurationStatusFilters) GetInOk() (*[]ConfigurationStatus, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ConfigurationStatusFilters) SetIn(v []ConfigurationStatus)`

SetIn sets In field to given value.

### HasIn

`func (o *ConfigurationStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *ConfigurationStatusFilters) GetNotIn() []ConfigurationStatus`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *ConfigurationStatusFilters) GetNotInOk() (*[]ConfigurationStatus, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *ConfigurationStatusFilters) SetNotIn(v []ConfigurationStatus)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *ConfigurationStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


