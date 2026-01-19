# AzureRestoreAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of the tenant the subscription belongs to. | 
**SubscriptionId** | **string** | Subscription ID.  | 
**ResourceGroupName** | Pointer to **string** | Resource group name. | [optional] 
**EonInternalResourceGroupName** | Pointer to **string** | Resource group name for Eon&#39;s temporary internal resources.  | [optional] [default to "eon-restore-internal-rg"]

## Methods

### NewAzureRestoreAccountAttributesInput

`func NewAzureRestoreAccountAttributesInput(tenantId string, subscriptionId string, ) *AzureRestoreAccountAttributesInput`

NewAzureRestoreAccountAttributesInput instantiates a new AzureRestoreAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRestoreAccountAttributesInputWithDefaults

`func NewAzureRestoreAccountAttributesInputWithDefaults() *AzureRestoreAccountAttributesInput`

NewAzureRestoreAccountAttributesInputWithDefaults instantiates a new AzureRestoreAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureRestoreAccountAttributesInput) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureRestoreAccountAttributesInput) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureRestoreAccountAttributesInput) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *AzureRestoreAccountAttributesInput) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureRestoreAccountAttributesInput) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureRestoreAccountAttributesInput) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureRestoreAccountAttributesInput) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) GetEonInternalResourceGroupName() string`

GetEonInternalResourceGroupName returns the EonInternalResourceGroupName field if non-nil, zero value otherwise.

### GetEonInternalResourceGroupNameOk

`func (o *AzureRestoreAccountAttributesInput) GetEonInternalResourceGroupNameOk() (*string, bool)`

GetEonInternalResourceGroupNameOk returns a tuple with the EonInternalResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) SetEonInternalResourceGroupName(v string)`

SetEonInternalResourceGroupName sets EonInternalResourceGroupName field to given value.

### HasEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributesInput) HasEonInternalResourceGroupName() bool`

HasEonInternalResourceGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


