# AccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**Aws** | Pointer to [**NullableAwsAccountConfig**](AwsAccountConfig.md) |  | [optional] 
**Gcp** | Pointer to [**NullableGcpAccountConfig**](GcpAccountConfig.md) |  | [optional] 
**Azure** | Pointer to [**NullableAzureAccountConfig**](AzureAccountConfig.md) |  | [optional] 

## Methods

### NewAccountConfig

`func NewAccountConfig() *AccountConfig`

NewAccountConfig instantiates a new AccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigWithDefaults

`func NewAccountConfigWithDefaults() *AccountConfig`

NewAccountConfigWithDefaults instantiates a new AccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AccountConfig) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AccountConfig) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AccountConfig) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AccountConfig) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAws

`func (o *AccountConfig) GetAws() AwsAccountConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *AccountConfig) GetAwsOk() (*AwsAccountConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *AccountConfig) SetAws(v AwsAccountConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *AccountConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *AccountConfig) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *AccountConfig) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *AccountConfig) GetGcp() GcpAccountConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *AccountConfig) GetGcpOk() (*GcpAccountConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *AccountConfig) SetGcp(v GcpAccountConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *AccountConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *AccountConfig) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *AccountConfig) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil
### GetAzure

`func (o *AccountConfig) GetAzure() AzureAccountConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *AccountConfig) GetAzureOk() (*AzureAccountConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *AccountConfig) SetAzure(v AzureAccountConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *AccountConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *AccountConfig) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *AccountConfig) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


