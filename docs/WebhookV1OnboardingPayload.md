# WebhookV1OnboardingPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Always &#x60;onboarding&#x60;. Also sent in the &#x60;X-Eon-Notification-Type&#x60; header.  | 
**Message** | **string** | Human-readable confirmation message. | 
**ProjectId** | **string** | ID of the project the notification policy is in. | 
**Timestamp** | **time.Time** | Time the message was sent, in RFC 3339 format with second precision. Also sent in the &#x60;X-Eon-Timestamp&#x60; header.  | 

## Methods

### NewWebhookV1OnboardingPayload

`func NewWebhookV1OnboardingPayload(type_ string, message string, projectId string, timestamp time.Time, ) *WebhookV1OnboardingPayload`

NewWebhookV1OnboardingPayload instantiates a new WebhookV1OnboardingPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1OnboardingPayloadWithDefaults

`func NewWebhookV1OnboardingPayloadWithDefaults() *WebhookV1OnboardingPayload`

NewWebhookV1OnboardingPayloadWithDefaults instantiates a new WebhookV1OnboardingPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookV1OnboardingPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1OnboardingPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1OnboardingPayload) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *WebhookV1OnboardingPayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookV1OnboardingPayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookV1OnboardingPayload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetProjectId

`func (o *WebhookV1OnboardingPayload) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1OnboardingPayload) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1OnboardingPayload) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTimestamp

`func (o *WebhookV1OnboardingPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1OnboardingPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1OnboardingPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


