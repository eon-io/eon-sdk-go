# BackupVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Vault ID. | 
**Name** | **string** | Vault display name. | 
**Region** | **string** | Region where the vault is located. | 
**VaultAttributes** | [**VaultProviderAttributes**](VaultProviderAttributes.md) |  | 

## Methods

### NewBackupVault

`func NewBackupVault(id string, name string, region string, vaultAttributes VaultProviderAttributes, ) *BackupVault`

NewBackupVault instantiates a new BackupVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupVaultWithDefaults

`func NewBackupVaultWithDefaults() *BackupVault`

NewBackupVaultWithDefaults instantiates a new BackupVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupVault) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupVault) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupVault) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BackupVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupVault) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *BackupVault) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BackupVault) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BackupVault) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVaultAttributes

`func (o *BackupVault) GetVaultAttributes() VaultProviderAttributes`

GetVaultAttributes returns the VaultAttributes field if non-nil, zero value otherwise.

### GetVaultAttributesOk

`func (o *BackupVault) GetVaultAttributesOk() (*VaultProviderAttributes, bool)`

GetVaultAttributesOk returns a tuple with the VaultAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAttributes

`func (o *BackupVault) SetVaultAttributes(v VaultProviderAttributes)`

SetVaultAttributes sets VaultAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


