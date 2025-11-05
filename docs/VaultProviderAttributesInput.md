# VaultProviderAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Aws** | Pointer to [**NullableAwsVaultConfigInput**](AwsVaultConfigInput.md) |  | [optional] 

## Methods

### NewVaultProviderAttributesInput

`func NewVaultProviderAttributesInput(cloudProvider Provider, ) *VaultProviderAttributesInput`

NewVaultProviderAttributesInput instantiates a new VaultProviderAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultProviderAttributesInputWithDefaults

`func NewVaultProviderAttributesInputWithDefaults() *VaultProviderAttributesInput`

NewVaultProviderAttributesInputWithDefaults instantiates a new VaultProviderAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *VaultProviderAttributesInput) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VaultProviderAttributesInput) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VaultProviderAttributesInput) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetAws

`func (o *VaultProviderAttributesInput) GetAws() AwsVaultConfigInput`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *VaultProviderAttributesInput) GetAwsOk() (*AwsVaultConfigInput, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *VaultProviderAttributesInput) SetAws(v AwsVaultConfigInput)`

SetAws sets Aws field to given value.

### HasAws

`func (o *VaultProviderAttributesInput) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *VaultProviderAttributesInput) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *VaultProviderAttributesInput) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


