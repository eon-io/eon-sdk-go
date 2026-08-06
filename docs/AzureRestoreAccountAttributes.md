# AzureRestoreAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of the tenant the subscription belongs to. | 
**SubscriptionId** | **string** | ID of the Azure subscription. | 
**ResourceGroupName** | Pointer to **string** | Name of the Azure resource group to scope permissions to. When provided, restoring is limited to this specific resource group. When omitted, permissions are scoped to the subscription.  | [optional] 
**EonInternalResourceGroupName** | Pointer to **string** | Name of the Eon internal resource group for temporary restore resources.  | [optional] 
**ManagementGroupId** | Pointer to **string** | ID of the Azure management group the restore account is scoped to. | [optional] 

## Methods

### NewAzureRestoreAccountAttributes

`func NewAzureRestoreAccountAttributes(tenantId string, subscriptionId string, ) *AzureRestoreAccountAttributes`

NewAzureRestoreAccountAttributes instantiates a new AzureRestoreAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRestoreAccountAttributesWithDefaults

`func NewAzureRestoreAccountAttributesWithDefaults() *AzureRestoreAccountAttributes`

NewAzureRestoreAccountAttributesWithDefaults instantiates a new AzureRestoreAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureRestoreAccountAttributes) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureRestoreAccountAttributes) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureRestoreAccountAttributes) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetSubscriptionId

`func (o *AzureRestoreAccountAttributes) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureRestoreAccountAttributes) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureRestoreAccountAttributes) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureRestoreAccountAttributes) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureRestoreAccountAttributes) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureRestoreAccountAttributes) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureRestoreAccountAttributes) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributes) GetEonInternalResourceGroupName() string`

GetEonInternalResourceGroupName returns the EonInternalResourceGroupName field if non-nil, zero value otherwise.

### GetEonInternalResourceGroupNameOk

`func (o *AzureRestoreAccountAttributes) GetEonInternalResourceGroupNameOk() (*string, bool)`

GetEonInternalResourceGroupNameOk returns a tuple with the EonInternalResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributes) SetEonInternalResourceGroupName(v string)`

SetEonInternalResourceGroupName sets EonInternalResourceGroupName field to given value.

### HasEonInternalResourceGroupName

`func (o *AzureRestoreAccountAttributes) HasEonInternalResourceGroupName() bool`

HasEonInternalResourceGroupName returns a boolean if a field has been set.

### GetManagementGroupId

`func (o *AzureRestoreAccountAttributes) GetManagementGroupId() string`

GetManagementGroupId returns the ManagementGroupId field if non-nil, zero value otherwise.

### GetManagementGroupIdOk

`func (o *AzureRestoreAccountAttributes) GetManagementGroupIdOk() (*string, bool)`

GetManagementGroupIdOk returns a tuple with the ManagementGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementGroupId

`func (o *AzureRestoreAccountAttributes) SetManagementGroupId(v string)`

SetManagementGroupId sets ManagementGroupId field to given value.

### HasManagementGroupId

`func (o *AzureRestoreAccountAttributes) HasManagementGroupId() bool`

HasManagementGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


