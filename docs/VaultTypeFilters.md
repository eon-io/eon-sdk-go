# VaultTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]VaultType**](VaultType.md) | Matches if any value in this list equals the vault type. | [optional] 
**NotIn** | Pointer to [**[]VaultType**](VaultType.md) | Matches if none of the values in this list equal the vault type. | [optional] 

## Methods

### NewVaultTypeFilters

`func NewVaultTypeFilters() *VaultTypeFilters`

NewVaultTypeFilters instantiates a new VaultTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultTypeFiltersWithDefaults

`func NewVaultTypeFiltersWithDefaults() *VaultTypeFilters`

NewVaultTypeFiltersWithDefaults instantiates a new VaultTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *VaultTypeFilters) GetIn() []VaultType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *VaultTypeFilters) GetInOk() (*[]VaultType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *VaultTypeFilters) SetIn(v []VaultType)`

SetIn sets In field to given value.

### HasIn

`func (o *VaultTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *VaultTypeFilters) GetNotIn() []VaultType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *VaultTypeFilters) GetNotInOk() (*[]VaultType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *VaultTypeFilters) SetNotIn(v []VaultType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *VaultTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


