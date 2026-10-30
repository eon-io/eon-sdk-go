# VaultStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]VaultStatus**](VaultStatus.md) | Matches if any value in this list equals the vault status. | [optional] 
**NotIn** | Pointer to [**[]VaultStatus**](VaultStatus.md) | Matches if none of the values in this list equal the vault status. | [optional] 

## Methods

### NewVaultStatusFilters

`func NewVaultStatusFilters() *VaultStatusFilters`

NewVaultStatusFilters instantiates a new VaultStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultStatusFiltersWithDefaults

`func NewVaultStatusFiltersWithDefaults() *VaultStatusFilters`

NewVaultStatusFiltersWithDefaults instantiates a new VaultStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *VaultStatusFilters) GetIn() []VaultStatus`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *VaultStatusFilters) GetInOk() (*[]VaultStatus, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *VaultStatusFilters) SetIn(v []VaultStatus)`

SetIn sets In field to given value.

### HasIn

`func (o *VaultStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *VaultStatusFilters) GetNotIn() []VaultStatus`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *VaultStatusFilters) GetNotInOk() (*[]VaultStatus, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *VaultStatusFilters) SetNotIn(v []VaultStatus)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *VaultStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


