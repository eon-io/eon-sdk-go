# AzureAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | The tenant ID | 
**ResourceGroupName** | Pointer to **string** | The resource group name | [optional] 

## Methods

### NewAzureAccountConfig

`func NewAzureAccountConfig(tenantId string, ) *AzureAccountConfig`

NewAzureAccountConfig instantiates a new AzureAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAccountConfigWithDefaults

`func NewAzureAccountConfigWithDefaults() *AzureAccountConfig`

NewAzureAccountConfigWithDefaults instantiates a new AzureAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureAccountConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureAccountConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureAccountConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetResourceGroupName

`func (o *AzureAccountConfig) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureAccountConfig) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureAccountConfig) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureAccountConfig) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


