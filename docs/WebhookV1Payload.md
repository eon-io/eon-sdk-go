# WebhookV1Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**WebhookV1AuditEvent**](WebhookV1AuditEvent.md) |  | 
**Type** | **string** | Always &#x60;onboarding&#x60;. Also sent in the &#x60;X-Eon-Notification-Type&#x60; header. | 
**Message** | **string** | Human-readable confirmation message. | 
**ProjectId** | **string** | ID of the project the webhook target belongs to. | 
**Timestamp** | **time.Time** | Time the message was sent, in RFC 3339 format with second precision. Also sent in the &#x60;X-Eon-Timestamp&#x60; header.  | 

## Methods

### NewWebhookV1Payload

`func NewWebhookV1Payload(event WebhookV1AuditEvent, type_ string, message string, projectId string, timestamp time.Time, ) *WebhookV1Payload`

NewWebhookV1Payload instantiates a new WebhookV1Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1PayloadWithDefaults

`func NewWebhookV1PayloadWithDefaults() *WebhookV1Payload`

NewWebhookV1PayloadWithDefaults instantiates a new WebhookV1Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookV1Payload) GetEvent() WebhookV1AuditEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookV1Payload) GetEventOk() (*WebhookV1AuditEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookV1Payload) SetEvent(v WebhookV1AuditEvent)`

SetEvent sets Event field to given value.


### GetType

`func (o *WebhookV1Payload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1Payload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1Payload) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *WebhookV1Payload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookV1Payload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookV1Payload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetProjectId

`func (o *WebhookV1Payload) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1Payload) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1Payload) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTimestamp

`func (o *WebhookV1Payload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1Payload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1Payload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


