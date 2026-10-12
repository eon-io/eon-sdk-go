# WebhookV1ProtobufTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **int64** | Seconds since the Unix epoch. | [optional] 
**Nanos** | Pointer to **int32** | Nanosecond fraction of the second. | [optional] 

## Methods

### NewWebhookV1ProtobufTimestamp

`func NewWebhookV1ProtobufTimestamp() *WebhookV1ProtobufTimestamp`

NewWebhookV1ProtobufTimestamp instantiates a new WebhookV1ProtobufTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1ProtobufTimestampWithDefaults

`func NewWebhookV1ProtobufTimestampWithDefaults() *WebhookV1ProtobufTimestamp`

NewWebhookV1ProtobufTimestampWithDefaults instantiates a new WebhookV1ProtobufTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *WebhookV1ProtobufTimestamp) GetSeconds() int64`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *WebhookV1ProtobufTimestamp) GetSecondsOk() (*int64, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *WebhookV1ProtobufTimestamp) SetSeconds(v int64)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *WebhookV1ProtobufTimestamp) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### GetNanos

`func (o *WebhookV1ProtobufTimestamp) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *WebhookV1ProtobufTimestamp) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *WebhookV1ProtobufTimestamp) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *WebhookV1ProtobufTimestamp) HasNanos() bool`

HasNanos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


