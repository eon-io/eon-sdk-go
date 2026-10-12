# WebhookV1AuditPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**WebhookV1AuditEvent**](WebhookV1AuditEvent.md) |  | 

## Methods

### NewWebhookV1AuditPayload

`func NewWebhookV1AuditPayload(event WebhookV1AuditEvent, ) *WebhookV1AuditPayload`

NewWebhookV1AuditPayload instantiates a new WebhookV1AuditPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1AuditPayloadWithDefaults

`func NewWebhookV1AuditPayloadWithDefaults() *WebhookV1AuditPayload`

NewWebhookV1AuditPayloadWithDefaults instantiates a new WebhookV1AuditPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookV1AuditPayload) GetEvent() WebhookV1AuditEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookV1AuditPayload) GetEventOk() (*WebhookV1AuditEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookV1AuditPayload) SetEvent(v WebhookV1AuditEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


