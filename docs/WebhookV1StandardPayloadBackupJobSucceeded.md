# WebhookV1StandardPayloadBackupJobSucceeded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Notification ID. Also sent in the &#x60;X-Eon-Notification-Id&#x60; header. | 
**ProjectId** | **string** | ID of the project the notification policy is in. | 
**Timestamp** | **time.Time** | Time the notification was created, in RFC 3339 format with second precision. Also sent in the &#x60;X-Eon-Timestamp&#x60; header.  | 
**Type** | **string** | Notification type. Also sent in the &#x60;X-Eon-Notification-Type&#x60; header. | 
**Arguments** | [**WebhookV1StandardPayloadBackupJobSucceededArguments**](WebhookV1StandardPayloadBackupJobSucceededArguments.md) |  | 

## Methods

### NewWebhookV1StandardPayloadBackupJobSucceeded

`func NewWebhookV1StandardPayloadBackupJobSucceeded(id string, projectId string, timestamp time.Time, type_ string, arguments WebhookV1StandardPayloadBackupJobSucceededArguments, ) *WebhookV1StandardPayloadBackupJobSucceeded`

NewWebhookV1StandardPayloadBackupJobSucceeded instantiates a new WebhookV1StandardPayloadBackupJobSucceeded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1StandardPayloadBackupJobSucceededWithDefaults

`func NewWebhookV1StandardPayloadBackupJobSucceededWithDefaults() *WebhookV1StandardPayloadBackupJobSucceeded`

NewWebhookV1StandardPayloadBackupJobSucceededWithDefaults instantiates a new WebhookV1StandardPayloadBackupJobSucceeded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTimestamp

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) SetType(v string)`

SetType sets Type field to given value.


### GetArguments

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetArguments() WebhookV1StandardPayloadBackupJobSucceededArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) GetArgumentsOk() (*WebhookV1StandardPayloadBackupJobSucceededArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *WebhookV1StandardPayloadBackupJobSucceeded) SetArguments(v WebhookV1StandardPayloadBackupJobSucceededArguments)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


