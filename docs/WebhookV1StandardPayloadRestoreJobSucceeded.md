# WebhookV1StandardPayloadRestoreJobSucceeded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Notification ID. Also sent in the &#x60;X-Eon-Notification-Id&#x60; header. | 
**ProjectId** | **string** | ID of the project the notification policy is in. | 
**Timestamp** | **time.Time** | Time the notification was created, in RFC 3339 format with second precision. Also sent in the &#x60;X-Eon-Timestamp&#x60; header.  | 
**Type** | **string** | Notification type. Also sent in the &#x60;X-Eon-Notification-Type&#x60; header. | 
**Arguments** | [**WebhookV1StandardPayloadRestoreJobSucceededArguments**](WebhookV1StandardPayloadRestoreJobSucceededArguments.md) |  | 

## Methods

### NewWebhookV1StandardPayloadRestoreJobSucceeded

`func NewWebhookV1StandardPayloadRestoreJobSucceeded(id string, projectId string, timestamp time.Time, type_ string, arguments WebhookV1StandardPayloadRestoreJobSucceededArguments, ) *WebhookV1StandardPayloadRestoreJobSucceeded`

NewWebhookV1StandardPayloadRestoreJobSucceeded instantiates a new WebhookV1StandardPayloadRestoreJobSucceeded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1StandardPayloadRestoreJobSucceededWithDefaults

`func NewWebhookV1StandardPayloadRestoreJobSucceededWithDefaults() *WebhookV1StandardPayloadRestoreJobSucceeded`

NewWebhookV1StandardPayloadRestoreJobSucceededWithDefaults instantiates a new WebhookV1StandardPayloadRestoreJobSucceeded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTimestamp

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) SetType(v string)`

SetType sets Type field to given value.


### GetArguments

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetArguments() WebhookV1StandardPayloadRestoreJobSucceededArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) GetArgumentsOk() (*WebhookV1StandardPayloadRestoreJobSucceededArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *WebhookV1StandardPayloadRestoreJobSucceeded) SetArguments(v WebhookV1StandardPayloadRestoreJobSucceededArguments)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


