# WebhookV1NotificationTargetChangeArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **string** | Notification policy ID. | [optional] 
**PolicyName** | Pointer to **string** | Notification policy name. | [optional] 
**UserName** | Pointer to **string** | Name of the user who made the change. | [optional] 
**Timestamp** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 

## Methods

### NewWebhookV1NotificationTargetChangeArguments

`func NewWebhookV1NotificationTargetChangeArguments() *WebhookV1NotificationTargetChangeArguments`

NewWebhookV1NotificationTargetChangeArguments instantiates a new WebhookV1NotificationTargetChangeArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1NotificationTargetChangeArgumentsWithDefaults

`func NewWebhookV1NotificationTargetChangeArgumentsWithDefaults() *WebhookV1NotificationTargetChangeArguments`

NewWebhookV1NotificationTargetChangeArgumentsWithDefaults instantiates a new WebhookV1NotificationTargetChangeArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *WebhookV1NotificationTargetChangeArguments) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *WebhookV1NotificationTargetChangeArguments) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *WebhookV1NotificationTargetChangeArguments) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *WebhookV1NotificationTargetChangeArguments) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *WebhookV1NotificationTargetChangeArguments) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *WebhookV1NotificationTargetChangeArguments) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *WebhookV1NotificationTargetChangeArguments) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *WebhookV1NotificationTargetChangeArguments) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetUserName

`func (o *WebhookV1NotificationTargetChangeArguments) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *WebhookV1NotificationTargetChangeArguments) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *WebhookV1NotificationTargetChangeArguments) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *WebhookV1NotificationTargetChangeArguments) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookV1NotificationTargetChangeArguments) GetTimestamp() WebhookV1ProtobufTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1NotificationTargetChangeArguments) GetTimestampOk() (*WebhookV1ProtobufTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1NotificationTargetChangeArguments) SetTimestamp(v WebhookV1ProtobufTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookV1NotificationTargetChangeArguments) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


