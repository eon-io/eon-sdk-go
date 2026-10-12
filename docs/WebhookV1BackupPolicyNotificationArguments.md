# WebhookV1BackupPolicyNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy ID. | [optional] 
**ResourceIds** | Pointer to **[]string** | IDs of the affected resources. | [optional] 
**TotalResources** | Pointer to **int32** | Total number of affected resources. | [optional] 
**Name** | Pointer to **string** | Policy name. | [optional] 
**ResourceNames** | Pointer to **[]string** | Names of the affected resources. | [optional] 

## Methods

### NewWebhookV1BackupPolicyNotificationArguments

`func NewWebhookV1BackupPolicyNotificationArguments() *WebhookV1BackupPolicyNotificationArguments`

NewWebhookV1BackupPolicyNotificationArguments instantiates a new WebhookV1BackupPolicyNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1BackupPolicyNotificationArgumentsWithDefaults

`func NewWebhookV1BackupPolicyNotificationArgumentsWithDefaults() *WebhookV1BackupPolicyNotificationArguments`

NewWebhookV1BackupPolicyNotificationArgumentsWithDefaults instantiates a new WebhookV1BackupPolicyNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1BackupPolicyNotificationArguments) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1BackupPolicyNotificationArguments) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1BackupPolicyNotificationArguments) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookV1BackupPolicyNotificationArguments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceIds

`func (o *WebhookV1BackupPolicyNotificationArguments) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *WebhookV1BackupPolicyNotificationArguments) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *WebhookV1BackupPolicyNotificationArguments) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *WebhookV1BackupPolicyNotificationArguments) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetTotalResources

`func (o *WebhookV1BackupPolicyNotificationArguments) GetTotalResources() int32`

GetTotalResources returns the TotalResources field if non-nil, zero value otherwise.

### GetTotalResourcesOk

`func (o *WebhookV1BackupPolicyNotificationArguments) GetTotalResourcesOk() (*int32, bool)`

GetTotalResourcesOk returns a tuple with the TotalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResources

`func (o *WebhookV1BackupPolicyNotificationArguments) SetTotalResources(v int32)`

SetTotalResources sets TotalResources field to given value.

### HasTotalResources

`func (o *WebhookV1BackupPolicyNotificationArguments) HasTotalResources() bool`

HasTotalResources returns a boolean if a field has been set.

### GetName

`func (o *WebhookV1BackupPolicyNotificationArguments) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookV1BackupPolicyNotificationArguments) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookV1BackupPolicyNotificationArguments) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookV1BackupPolicyNotificationArguments) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceNames

`func (o *WebhookV1BackupPolicyNotificationArguments) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *WebhookV1BackupPolicyNotificationArguments) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *WebhookV1BackupPolicyNotificationArguments) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *WebhookV1BackupPolicyNotificationArguments) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


