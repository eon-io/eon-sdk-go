# WebhookV1BackupJobDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | Name of the resource. | [optional] 
**ResourceId** | Pointer to **string** | Cloud-provider-assigned ID of the resource. | [optional] 
**PolicyName** | Pointer to **string** | Name of the backup policy that triggered the job. | [optional] 
**SourceAccountName** | Pointer to **string** | Display name of the source account. | [optional] 
**Duration** | Pointer to [**WebhookV1ProtobufDuration**](WebhookV1ProtobufDuration.md) |  | [optional] 
**ResourceType** | Pointer to **string** | Resource type. | [optional] 

## Methods

### NewWebhookV1BackupJobDetails

`func NewWebhookV1BackupJobDetails() *WebhookV1BackupJobDetails`

NewWebhookV1BackupJobDetails instantiates a new WebhookV1BackupJobDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1BackupJobDetailsWithDefaults

`func NewWebhookV1BackupJobDetailsWithDefaults() *WebhookV1BackupJobDetails`

NewWebhookV1BackupJobDetailsWithDefaults instantiates a new WebhookV1BackupJobDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *WebhookV1BackupJobDetails) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *WebhookV1BackupJobDetails) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *WebhookV1BackupJobDetails) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *WebhookV1BackupJobDetails) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceId

`func (o *WebhookV1BackupJobDetails) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *WebhookV1BackupJobDetails) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *WebhookV1BackupJobDetails) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *WebhookV1BackupJobDetails) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetPolicyName

`func (o *WebhookV1BackupJobDetails) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *WebhookV1BackupJobDetails) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *WebhookV1BackupJobDetails) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *WebhookV1BackupJobDetails) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *WebhookV1BackupJobDetails) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *WebhookV1BackupJobDetails) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *WebhookV1BackupJobDetails) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *WebhookV1BackupJobDetails) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetDuration

`func (o *WebhookV1BackupJobDetails) GetDuration() WebhookV1ProtobufDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WebhookV1BackupJobDetails) GetDurationOk() (*WebhookV1ProtobufDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WebhookV1BackupJobDetails) SetDuration(v WebhookV1ProtobufDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WebhookV1BackupJobDetails) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResourceType

`func (o *WebhookV1BackupJobDetails) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebhookV1BackupJobDetails) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebhookV1BackupJobDetails) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WebhookV1BackupJobDetails) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


