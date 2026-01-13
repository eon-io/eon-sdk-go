# AzureSourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of the Azure tenant the source subscription belongs to. | 
**SubscriptionId** | **string** | ID of the Azure source subscription. | 
**ResourceGroupName** | Pointer to **string** | Name of the Azure resource group to scope the source account to. When provided, discovery and permissions are limited to this specific resource group.  | [optional] 
**EonInternalResourceGroupName** | Pointer to **string** | Eon internal resource group name for temporary resources. Defaults to eon-source-internal-rg if not provided.  | [optional] 

## Methods

### NewAzureSourceAccountAttributesInput

`func NewAzureSourceAccountAttributesInput(tenantId string, subscriptionId string, ) *AzureSourceAccountAttributesInput`

NewAzureSourceAccountAttributesInput instantiates a new AzureSourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSourceAccountAttributesInputWithDefaults

`func NewAzureSourceAccountAttributesInputWithDefaults() *AzureSourceAccountAttributesInput`

NewAzureSourceAccountAttributesInputWithDefaults instantiates a new AzureSourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureSourceAccountAttributesInput) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureSourceAccountAttributesInput) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureSourceAccountAttributesInput) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *AzureSourceAccountAttributesInput) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureSourceAccountAttributesInput) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureSourceAccountAttributesInput) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureSourceAccountAttributesInput) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureSourceAccountAttributesInput) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureSourceAccountAttributesInput) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureSourceAccountAttributesInput) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetEonInternalResourceGroupName

`func (o *AzureSourceAccountAttributesInput) GetEonInternalResourceGroupName() string`

GetEonInternalResourceGroupName returns the EonInternalResourceGroupName field if non-nil, zero value otherwise.

### GetEonInternalResourceGroupNameOk

`func (o *AzureSourceAccountAttributesInput) GetEonInternalResourceGroupNameOk() (*string, bool)`

GetEonInternalResourceGroupNameOk returns a tuple with the EonInternalResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEonInternalResourceGroupName

`func (o *AzureSourceAccountAttributesInput) SetEonInternalResourceGroupName(v string)`

SetEonInternalResourceGroupName sets EonInternalResourceGroupName field to given value.

### HasEonInternalResourceGroupName

`func (o *AzureSourceAccountAttributesInput) HasEonInternalResourceGroupName() bool`

HasEonInternalResourceGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


