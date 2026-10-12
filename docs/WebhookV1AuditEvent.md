# WebhookV1AuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Event type. | 
**Payload** | **interface{}** | The audit log entry as arbitrary JSON. Its structure isn&#39;t part of the webhook contract. | 

## Methods

### NewWebhookV1AuditEvent

`func NewWebhookV1AuditEvent(type_ string, payload interface{}, ) *WebhookV1AuditEvent`

NewWebhookV1AuditEvent instantiates a new WebhookV1AuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1AuditEventWithDefaults

`func NewWebhookV1AuditEventWithDefaults() *WebhookV1AuditEvent`

NewWebhookV1AuditEventWithDefaults instantiates a new WebhookV1AuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookV1AuditEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1AuditEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1AuditEvent) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *WebhookV1AuditEvent) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookV1AuditEvent) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookV1AuditEvent) SetPayload(v interface{})`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *WebhookV1AuditEvent) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *WebhookV1AuditEvent) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


