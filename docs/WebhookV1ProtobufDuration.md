# WebhookV1ProtobufDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int64** | Whole seconds. Omitted when &#x60;0&#x60;.  | [optional] 
**Nanos** | Pointer to **int32** | Nanosecond fraction of the second. Omitted when &#x60;0&#x60;.  | [optional] 

## Methods

### NewWebhookV1ProtobufDuration

`func NewWebhookV1ProtobufDuration() *WebhookV1ProtobufDuration`

NewWebhookV1ProtobufDuration instantiates a new WebhookV1ProtobufDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1ProtobufDurationWithDefaults

`func NewWebhookV1ProtobufDurationWithDefaults() *WebhookV1ProtobufDuration`

NewWebhookV1ProtobufDurationWithDefaults instantiates a new WebhookV1ProtobufDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *WebhookV1ProtobufDuration) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *WebhookV1ProtobufDuration) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *WebhookV1ProtobufDuration) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *WebhookV1ProtobufDuration) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetNanos

`func (o *WebhookV1ProtobufDuration) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *WebhookV1ProtobufDuration) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *WebhookV1ProtobufDuration) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *WebhookV1ProtobufDuration) HasNanos() bool`

HasNanos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


