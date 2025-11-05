# CreateVaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vault** | [**BackupVault**](BackupVault.md) |  | 

## Methods

### NewCreateVaultResponse

`func NewCreateVaultResponse(vault BackupVault, ) *CreateVaultResponse`

NewCreateVaultResponse instantiates a new CreateVaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVaultResponseWithDefaults

`func NewCreateVaultResponseWithDefaults() *CreateVaultResponse`

NewCreateVaultResponseWithDefaults instantiates a new CreateVaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVault

`func (o *CreateVaultResponse) GetVault() BackupVault`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *CreateVaultResponse) GetVaultOk() (*BackupVault, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *CreateVaultResponse) SetVault(v BackupVault)`

SetVault sets Vault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


