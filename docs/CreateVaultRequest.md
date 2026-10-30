# CreateVaultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Vault display name. | 
**Region** | **string** | Region where the vault is hosted. | 
**VaultAttributes** | [**VaultProviderAttributesInput**](VaultProviderAttributesInput.md) |  | 
**Labels** | Pointer to **[]string** | List of labels for filtering and searching for vaults. Doesn&#39;t affect backup behavior, access control, or the tags Eon sets on your cloud resources.  | [optional] 

## Methods

### NewCreateVaultRequest

`func NewCreateVaultRequest(name string, region string, vaultAttributes VaultProviderAttributesInput, ) *CreateVaultRequest`

NewCreateVaultRequest instantiates a new CreateVaultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVaultRequestWithDefaults

`func NewCreateVaultRequestWithDefaults() *CreateVaultRequest`

NewCreateVaultRequestWithDefaults instantiates a new CreateVaultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVaultRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVaultRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVaultRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CreateVaultRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateVaultRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateVaultRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVaultAttributes

`func (o *CreateVaultRequest) GetVaultAttributes() VaultProviderAttributesInput`

GetVaultAttributes returns the VaultAttributes field if non-nil, zero value otherwise.

### GetVaultAttributesOk

`func (o *CreateVaultRequest) GetVaultAttributesOk() (*VaultProviderAttributesInput, bool)`

GetVaultAttributesOk returns a tuple with the VaultAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAttributes

`func (o *CreateVaultRequest) SetVaultAttributes(v VaultProviderAttributesInput)`

SetVaultAttributes sets VaultAttributes field to given value.


### GetLabels

`func (o *CreateVaultRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateVaultRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateVaultRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateVaultRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


