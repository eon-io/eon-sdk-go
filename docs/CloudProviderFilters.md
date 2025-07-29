# CloudProviderFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]Provider**](Provider.md) | Matches if any value in this list equals the cloud provider. | [optional] 
**NotIn** | Pointer to [**[]Provider**](Provider.md) | Matches if none of the values in this list equal the cloud provider. | [optional] 

## Methods

### NewCloudProviderFilters

`func NewCloudProviderFilters() *CloudProviderFilters`

NewCloudProviderFilters instantiates a new CloudProviderFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderFiltersWithDefaults

`func NewCloudProviderFiltersWithDefaults() *CloudProviderFilters`

NewCloudProviderFiltersWithDefaults instantiates a new CloudProviderFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *CloudProviderFilters) GetIn() []Provider`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *CloudProviderFilters) GetInOk() (*[]Provider, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *CloudProviderFilters) SetIn(v []Provider)`

SetIn sets In field to given value.

### HasIn

`func (o *CloudProviderFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *CloudProviderFilters) GetNotIn() []Provider`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *CloudProviderFilters) GetNotInOk() (*[]Provider, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *CloudProviderFilters) SetNotIn(v []Provider)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *CloudProviderFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


