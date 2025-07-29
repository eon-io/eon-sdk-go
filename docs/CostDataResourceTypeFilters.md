# CostDataResourceTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]ResourceType**](ResourceType.md) | Matches if any value in this list equals the resource type. | [optional] 
**NotIn** | Pointer to [**[]ResourceType**](ResourceType.md) | Matches if none of the values in this list equal the resource type. | [optional] 

## Methods

### NewCostDataResourceTypeFilters

`func NewCostDataResourceTypeFilters() *CostDataResourceTypeFilters`

NewCostDataResourceTypeFilters instantiates a new CostDataResourceTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataResourceTypeFiltersWithDefaults

`func NewCostDataResourceTypeFiltersWithDefaults() *CostDataResourceTypeFilters`

NewCostDataResourceTypeFiltersWithDefaults instantiates a new CostDataResourceTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *CostDataResourceTypeFilters) GetIn() []ResourceType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *CostDataResourceTypeFilters) GetInOk() (*[]ResourceType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *CostDataResourceTypeFilters) SetIn(v []ResourceType)`

SetIn sets In field to given value.

### HasIn

`func (o *CostDataResourceTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *CostDataResourceTypeFilters) GetNotIn() []ResourceType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *CostDataResourceTypeFilters) GetNotInOk() (*[]ResourceType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *CostDataResourceTypeFilters) SetNotIn(v []ResourceType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *CostDataResourceTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


