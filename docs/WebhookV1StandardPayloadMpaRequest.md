# WebhookV1StandardPayloadMpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Notification ID. Also sent in the &#x60;X-Eon-Notification-Id&#x60; header. | 
**ProjectId** | **string** | ID of the project the notification policy is in. | 
**Timestamp** | **time.Time** | Time the notification was created, in RFC 3339 format with second precision. Also sent in the &#x60;X-Eon-Timestamp&#x60; header.  | 
**Type** | **string** | Notification type. Also sent in the &#x60;X-Eon-Notification-Type&#x60; header. | 
**Arguments** | [**WebhookV1StandardPayloadMpaRequestArguments**](WebhookV1StandardPayloadMpaRequestArguments.md) |  | 

## Methods

### NewWebhookV1StandardPayloadMpaRequest

`func NewWebhookV1StandardPayloadMpaRequest(id string, projectId string, timestamp time.Time, type_ string, arguments WebhookV1StandardPayloadMpaRequestArguments, ) *WebhookV1StandardPayloadMpaRequest`

NewWebhookV1StandardPayloadMpaRequest instantiates a new WebhookV1StandardPayloadMpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1StandardPayloadMpaRequestWithDefaults

`func NewWebhookV1StandardPayloadMpaRequestWithDefaults() *WebhookV1StandardPayloadMpaRequest`

NewWebhookV1StandardPayloadMpaRequestWithDefaults instantiates a new WebhookV1StandardPayloadMpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1StandardPayloadMpaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1StandardPayloadMpaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1StandardPayloadMpaRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *WebhookV1StandardPayloadMpaRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1StandardPayloadMpaRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1StandardPayloadMpaRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTimestamp

`func (o *WebhookV1StandardPayloadMpaRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1StandardPayloadMpaRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1StandardPayloadMpaRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *WebhookV1StandardPayloadMpaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1StandardPayloadMpaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1StandardPayloadMpaRequest) SetType(v string)`

SetType sets Type field to given value.


### GetArguments

`func (o *WebhookV1StandardPayloadMpaRequest) GetArguments() WebhookV1StandardPayloadMpaRequestArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *WebhookV1StandardPayloadMpaRequest) GetArgumentsOk() (*WebhookV1StandardPayloadMpaRequestArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *WebhookV1StandardPayloadMpaRequest) SetArguments(v WebhookV1StandardPayloadMpaRequestArguments)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


