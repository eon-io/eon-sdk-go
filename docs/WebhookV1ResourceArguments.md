# WebhookV1ResourceArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **string** | Cloud-provider-assigned ID of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**SecurityScan** | Pointer to [**WebhookV1SecurityScan**](WebhookV1SecurityScan.md) |  | [optional] 
**ResourceType** | Pointer to **string** | Resource type. | [optional] 

## Methods

### NewWebhookV1ResourceArguments

`func NewWebhookV1ResourceArguments() *WebhookV1ResourceArguments`

NewWebhookV1ResourceArguments instantiates a new WebhookV1ResourceArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1ResourceArgumentsWithDefaults

`func NewWebhookV1ResourceArgumentsWithDefaults() *WebhookV1ResourceArguments`

NewWebhookV1ResourceArgumentsWithDefaults instantiates a new WebhookV1ResourceArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *WebhookV1ResourceArguments) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *WebhookV1ResourceArguments) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *WebhookV1ResourceArguments) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *WebhookV1ResourceArguments) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetName

`func (o *WebhookV1ResourceArguments) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookV1ResourceArguments) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookV1ResourceArguments) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookV1ResourceArguments) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityScan

`func (o *WebhookV1ResourceArguments) GetSecurityScan() WebhookV1SecurityScan`

GetSecurityScan returns the SecurityScan field if non-nil, zero value otherwise.

### GetSecurityScanOk

`func (o *WebhookV1ResourceArguments) GetSecurityScanOk() (*WebhookV1SecurityScan, bool)`

GetSecurityScanOk returns a tuple with the SecurityScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScan

`func (o *WebhookV1ResourceArguments) SetSecurityScan(v WebhookV1SecurityScan)`

SetSecurityScan sets SecurityScan field to given value.

### HasSecurityScan

`func (o *WebhookV1ResourceArguments) HasSecurityScan() bool`

HasSecurityScan returns a boolean if a field has been set.

### GetResourceType

`func (o *WebhookV1ResourceArguments) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebhookV1ResourceArguments) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebhookV1ResourceArguments) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WebhookV1ResourceArguments) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


