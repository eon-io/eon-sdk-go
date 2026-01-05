# AzureRestoreAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of the tenant the subscription belongs to. | 
**SubscriptionId** | **string** | ID of the Azure subscription. | 
**ResourceGroupName** | Pointer to **string** | Name of the target restore resource group. | [optional] 
**Location** | Pointer to **string** | Azure region for restore operations. | [optional] 

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

### GetLocation

`func (o *AzureRestoreAccountAttributes) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureRestoreAccountAttributes) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureRestoreAccountAttributes) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureRestoreAccountAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


