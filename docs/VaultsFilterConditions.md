# VaultsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**FreeTextFilters**](FreeTextFilters.md) |  | [optional] 
**CloudProvider** | Pointer to [**CloudProviderFilters**](CloudProviderFilters.md) |  | [optional] 
**Region** | Pointer to [**VaultRegionFilters**](VaultRegionFilters.md) |  | [optional] 
**Status** | Pointer to [**VaultStatusFilters**](VaultStatusFilters.md) |  | [optional] 
**VaultType** | Pointer to [**VaultTypeFilters**](VaultTypeFilters.md) |  | [optional] 

## Methods

### NewVaultsFilterConditions

`func NewVaultsFilterConditions() *VaultsFilterConditions`

NewVaultsFilterConditions instantiates a new VaultsFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultsFilterConditionsWithDefaults

`func NewVaultsFilterConditionsWithDefaults() *VaultsFilterConditions`

NewVaultsFilterConditionsWithDefaults instantiates a new VaultsFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VaultsFilterConditions) GetName() FreeTextFilters`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultsFilterConditions) GetNameOk() (*FreeTextFilters, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultsFilterConditions) SetName(v FreeTextFilters)`

SetName sets Name field to given value.

### HasName

`func (o *VaultsFilterConditions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *VaultsFilterConditions) GetCloudProvider() CloudProviderFilters`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VaultsFilterConditions) GetCloudProviderOk() (*CloudProviderFilters, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VaultsFilterConditions) SetCloudProvider(v CloudProviderFilters)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *VaultsFilterConditions) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetRegion

`func (o *VaultsFilterConditions) GetRegion() VaultRegionFilters`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VaultsFilterConditions) GetRegionOk() (*VaultRegionFilters, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VaultsFilterConditions) SetRegion(v VaultRegionFilters)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VaultsFilterConditions) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *VaultsFilterConditions) GetStatus() VaultStatusFilters`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VaultsFilterConditions) GetStatusOk() (*VaultStatusFilters, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VaultsFilterConditions) SetStatus(v VaultStatusFilters)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VaultsFilterConditions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVaultType

`func (o *VaultsFilterConditions) GetVaultType() VaultTypeFilters`

GetVaultType returns the VaultType field if non-nil, zero value otherwise.

### GetVaultTypeOk

`func (o *VaultsFilterConditions) GetVaultTypeOk() (*VaultTypeFilters, bool)`

GetVaultTypeOk returns a tuple with the VaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultType

`func (o *VaultsFilterConditions) SetVaultType(v VaultTypeFilters)`

SetVaultType sets VaultType field to given value.

### HasVaultType

`func (o *VaultsFilterConditions) HasVaultType() bool`

HasVaultType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


