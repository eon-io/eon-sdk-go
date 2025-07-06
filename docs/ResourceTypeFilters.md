# ResourceTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]ResourceType**](ResourceType.md) | Matches if any value in this list equals &#x60;resourceType&#x60;. | [optional] 
**NotIn** | Pointer to [**[]ResourceType**](ResourceType.md) | Matches if no value in this list equals &#x60;resourceType&#x60;. | [optional] 

## Methods

### NewResourceTypeFilters

`func NewResourceTypeFilters() *ResourceTypeFilters`

NewResourceTypeFilters instantiates a new ResourceTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeFiltersWithDefaults

`func NewResourceTypeFiltersWithDefaults() *ResourceTypeFilters`

NewResourceTypeFiltersWithDefaults instantiates a new ResourceTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *ResourceTypeFilters) GetIn() []ResourceType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ResourceTypeFilters) GetInOk() (*[]ResourceType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ResourceTypeFilters) SetIn(v []ResourceType)`

SetIn sets In field to given value.

### HasIn

`func (o *ResourceTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *ResourceTypeFilters) GetNotIn() []ResourceType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *ResourceTypeFilters) GetNotInOk() (*[]ResourceType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *ResourceTypeFilters) SetNotIn(v []ResourceType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *ResourceTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


