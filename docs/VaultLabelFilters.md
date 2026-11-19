# VaultLabelFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainsAnyOf** | Pointer to **[]string** | Matches if any value in this list is in the &#x60;labels&#x60; list. | [optional] 
**ContainsAllOf** | Pointer to **[]string** | Matches if all values in this list are in the &#x60;labels&#x60; list. | [optional] 
**ContainsNoneOf** | Pointer to **[]string** | Matches if none of the values in this list are in the &#x60;labels&#x60; list. | [optional] 

## Methods

### NewVaultLabelFilters

`func NewVaultLabelFilters() *VaultLabelFilters`

NewVaultLabelFilters instantiates a new VaultLabelFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultLabelFiltersWithDefaults

`func NewVaultLabelFiltersWithDefaults() *VaultLabelFilters`

NewVaultLabelFiltersWithDefaults instantiates a new VaultLabelFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainsAnyOf

`func (o *VaultLabelFilters) GetContainsAnyOf() []string`

GetContainsAnyOf returns the ContainsAnyOf field if non-nil, zero value otherwise.

### GetContainsAnyOfOk

`func (o *VaultLabelFilters) GetContainsAnyOfOk() (*[]string, bool)`

GetContainsAnyOfOk returns a tuple with the ContainsAnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAnyOf

`func (o *VaultLabelFilters) SetContainsAnyOf(v []string)`

SetContainsAnyOf sets ContainsAnyOf field to given value.

### HasContainsAnyOf

`func (o *VaultLabelFilters) HasContainsAnyOf() bool`

HasContainsAnyOf returns a boolean if a field has been set.

### GetContainsAllOf

`func (o *VaultLabelFilters) GetContainsAllOf() []string`

GetContainsAllOf returns the ContainsAllOf field if non-nil, zero value otherwise.

### GetContainsAllOfOk

`func (o *VaultLabelFilters) GetContainsAllOfOk() (*[]string, bool)`

GetContainsAllOfOk returns a tuple with the ContainsAllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAllOf

`func (o *VaultLabelFilters) SetContainsAllOf(v []string)`

SetContainsAllOf sets ContainsAllOf field to given value.

### HasContainsAllOf

`func (o *VaultLabelFilters) HasContainsAllOf() bool`

HasContainsAllOf returns a boolean if a field has been set.

### GetContainsNoneOf

`func (o *VaultLabelFilters) GetContainsNoneOf() []string`

GetContainsNoneOf returns the ContainsNoneOf field if non-nil, zero value otherwise.

### GetContainsNoneOfOk

`func (o *VaultLabelFilters) GetContainsNoneOfOk() (*[]string, bool)`

GetContainsNoneOfOk returns a tuple with the ContainsNoneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsNoneOf

`func (o *VaultLabelFilters) SetContainsNoneOf(v []string)`

SetContainsNoneOf sets ContainsNoneOf field to given value.

### HasContainsNoneOf

`func (o *VaultLabelFilters) HasContainsNoneOf() bool`

HasContainsNoneOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


