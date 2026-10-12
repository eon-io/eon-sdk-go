# WebhookV1SnapshotArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceProviderId** | Pointer to **string** | Cloud-provider-assigned ID of the snapshot&#39;s resource. | [optional] 
**ResourceName** | Pointer to **string** | Name of the snapshot&#39;s resource. | [optional] 
**SnapshotId** | Pointer to **string** | Eon-assigned snapshot ID. | [optional] 
**SecurityScan** | Pointer to [**WebhookV1SecurityScan**](WebhookV1SecurityScan.md) |  | [optional] 
**ResourceType** | Pointer to **string** | Resource type. | [optional] 

## Methods

### NewWebhookV1SnapshotArguments

`func NewWebhookV1SnapshotArguments() *WebhookV1SnapshotArguments`

NewWebhookV1SnapshotArguments instantiates a new WebhookV1SnapshotArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1SnapshotArgumentsWithDefaults

`func NewWebhookV1SnapshotArgumentsWithDefaults() *WebhookV1SnapshotArguments`

NewWebhookV1SnapshotArgumentsWithDefaults instantiates a new WebhookV1SnapshotArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceProviderId

`func (o *WebhookV1SnapshotArguments) GetResourceProviderId() string`

GetResourceProviderId returns the ResourceProviderId field if non-nil, zero value otherwise.

### GetResourceProviderIdOk

`func (o *WebhookV1SnapshotArguments) GetResourceProviderIdOk() (*string, bool)`

GetResourceProviderIdOk returns a tuple with the ResourceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProviderId

`func (o *WebhookV1SnapshotArguments) SetResourceProviderId(v string)`

SetResourceProviderId sets ResourceProviderId field to given value.

### HasResourceProviderId

`func (o *WebhookV1SnapshotArguments) HasResourceProviderId() bool`

HasResourceProviderId returns a boolean if a field has been set.

### GetResourceName

`func (o *WebhookV1SnapshotArguments) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *WebhookV1SnapshotArguments) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *WebhookV1SnapshotArguments) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *WebhookV1SnapshotArguments) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSnapshotId

`func (o *WebhookV1SnapshotArguments) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *WebhookV1SnapshotArguments) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *WebhookV1SnapshotArguments) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *WebhookV1SnapshotArguments) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSecurityScan

`func (o *WebhookV1SnapshotArguments) GetSecurityScan() WebhookV1SecurityScan`

GetSecurityScan returns the SecurityScan field if non-nil, zero value otherwise.

### GetSecurityScanOk

`func (o *WebhookV1SnapshotArguments) GetSecurityScanOk() (*WebhookV1SecurityScan, bool)`

GetSecurityScanOk returns a tuple with the SecurityScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScan

`func (o *WebhookV1SnapshotArguments) SetSecurityScan(v WebhookV1SecurityScan)`

SetSecurityScan sets SecurityScan field to given value.

### HasSecurityScan

`func (o *WebhookV1SnapshotArguments) HasSecurityScan() bool`

HasSecurityScan returns a boolean if a field has been set.

### GetResourceType

`func (o *WebhookV1SnapshotArguments) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebhookV1SnapshotArguments) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebhookV1SnapshotArguments) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WebhookV1SnapshotArguments) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


