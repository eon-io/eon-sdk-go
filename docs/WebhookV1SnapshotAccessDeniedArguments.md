# WebhookV1SnapshotAccessDeniedArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | ID of the backup job that created the snapshot. | [optional] 
**ResourceName** | Pointer to **string** | Name of the resource. | [optional] 
**ResourceProviderId** | Pointer to **string** | Cloud-provider-assigned ID of the resource. | [optional] 
**AccessDeniedItems** | Pointer to **int64** | Number of items the backup couldn&#39;t access. | [optional] 
**AccessDeniedBytes** | Pointer to **int64** | Total size, in bytes, of the items the backup couldn&#39;t access. | [optional] 
**ResourceType** | Pointer to **string** | Resource type. | [optional] 

## Methods

### NewWebhookV1SnapshotAccessDeniedArguments

`func NewWebhookV1SnapshotAccessDeniedArguments() *WebhookV1SnapshotAccessDeniedArguments`

NewWebhookV1SnapshotAccessDeniedArguments instantiates a new WebhookV1SnapshotAccessDeniedArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1SnapshotAccessDeniedArgumentsWithDefaults

`func NewWebhookV1SnapshotAccessDeniedArgumentsWithDefaults() *WebhookV1SnapshotAccessDeniedArguments`

NewWebhookV1SnapshotAccessDeniedArgumentsWithDefaults instantiates a new WebhookV1SnapshotAccessDeniedArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetResourceName

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceProviderId

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceProviderId() string`

GetResourceProviderId returns the ResourceProviderId field if non-nil, zero value otherwise.

### GetResourceProviderIdOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceProviderIdOk() (*string, bool)`

GetResourceProviderIdOk returns a tuple with the ResourceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProviderId

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetResourceProviderId(v string)`

SetResourceProviderId sets ResourceProviderId field to given value.

### HasResourceProviderId

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasResourceProviderId() bool`

HasResourceProviderId returns a boolean if a field has been set.

### GetAccessDeniedItems

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetAccessDeniedItems() int64`

GetAccessDeniedItems returns the AccessDeniedItems field if non-nil, zero value otherwise.

### GetAccessDeniedItemsOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetAccessDeniedItemsOk() (*int64, bool)`

GetAccessDeniedItemsOk returns a tuple with the AccessDeniedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDeniedItems

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetAccessDeniedItems(v int64)`

SetAccessDeniedItems sets AccessDeniedItems field to given value.

### HasAccessDeniedItems

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasAccessDeniedItems() bool`

HasAccessDeniedItems returns a boolean if a field has been set.

### GetAccessDeniedBytes

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetAccessDeniedBytes() int64`

GetAccessDeniedBytes returns the AccessDeniedBytes field if non-nil, zero value otherwise.

### GetAccessDeniedBytesOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetAccessDeniedBytesOk() (*int64, bool)`

GetAccessDeniedBytesOk returns a tuple with the AccessDeniedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDeniedBytes

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetAccessDeniedBytes(v int64)`

SetAccessDeniedBytes sets AccessDeniedBytes field to given value.

### HasAccessDeniedBytes

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasAccessDeniedBytes() bool`

HasAccessDeniedBytes returns a boolean if a field has been set.

### GetResourceType

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebhookV1SnapshotAccessDeniedArguments) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebhookV1SnapshotAccessDeniedArguments) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WebhookV1SnapshotAccessDeniedArguments) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


