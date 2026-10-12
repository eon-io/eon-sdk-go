# WebhookV1RestoreJobRejectedArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Job ID. | [optional] 
**StartTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**EndTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**UserMessage** | Pointer to **string** | Explanation of the job outcome. | [optional] 

## Methods

### NewWebhookV1RestoreJobRejectedArguments

`func NewWebhookV1RestoreJobRejectedArguments() *WebhookV1RestoreJobRejectedArguments`

NewWebhookV1RestoreJobRejectedArguments instantiates a new WebhookV1RestoreJobRejectedArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1RestoreJobRejectedArgumentsWithDefaults

`func NewWebhookV1RestoreJobRejectedArgumentsWithDefaults() *WebhookV1RestoreJobRejectedArguments`

NewWebhookV1RestoreJobRejectedArgumentsWithDefaults instantiates a new WebhookV1RestoreJobRejectedArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1RestoreJobRejectedArguments) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1RestoreJobRejectedArguments) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1RestoreJobRejectedArguments) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookV1RestoreJobRejectedArguments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *WebhookV1RestoreJobRejectedArguments) GetStartTime() WebhookV1ProtobufTimestamp`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WebhookV1RestoreJobRejectedArguments) GetStartTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WebhookV1RestoreJobRejectedArguments) SetStartTime(v WebhookV1ProtobufTimestamp)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WebhookV1RestoreJobRejectedArguments) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WebhookV1RestoreJobRejectedArguments) GetEndTime() WebhookV1ProtobufTimestamp`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WebhookV1RestoreJobRejectedArguments) GetEndTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WebhookV1RestoreJobRejectedArguments) SetEndTime(v WebhookV1ProtobufTimestamp)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WebhookV1RestoreJobRejectedArguments) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUserMessage

`func (o *WebhookV1RestoreJobRejectedArguments) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *WebhookV1RestoreJobRejectedArguments) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *WebhookV1RestoreJobRejectedArguments) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *WebhookV1RestoreJobRejectedArguments) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


