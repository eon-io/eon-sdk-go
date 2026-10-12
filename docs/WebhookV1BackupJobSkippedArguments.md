# WebhookV1BackupJobSkippedArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Job ID. | [optional] 
**StartTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**EndTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**UserMessage** | Pointer to **string** | Explanation of the job outcome. | [optional] 
**Details** | Pointer to [**WebhookV1BackupJobDetails**](WebhookV1BackupJobDetails.md) |  | [optional] 

## Methods

### NewWebhookV1BackupJobSkippedArguments

`func NewWebhookV1BackupJobSkippedArguments() *WebhookV1BackupJobSkippedArguments`

NewWebhookV1BackupJobSkippedArguments instantiates a new WebhookV1BackupJobSkippedArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1BackupJobSkippedArgumentsWithDefaults

`func NewWebhookV1BackupJobSkippedArgumentsWithDefaults() *WebhookV1BackupJobSkippedArguments`

NewWebhookV1BackupJobSkippedArgumentsWithDefaults instantiates a new WebhookV1BackupJobSkippedArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV1BackupJobSkippedArguments) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV1BackupJobSkippedArguments) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV1BackupJobSkippedArguments) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookV1BackupJobSkippedArguments) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *WebhookV1BackupJobSkippedArguments) GetStartTime() WebhookV1ProtobufTimestamp`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WebhookV1BackupJobSkippedArguments) GetStartTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WebhookV1BackupJobSkippedArguments) SetStartTime(v WebhookV1ProtobufTimestamp)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WebhookV1BackupJobSkippedArguments) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WebhookV1BackupJobSkippedArguments) GetEndTime() WebhookV1ProtobufTimestamp`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WebhookV1BackupJobSkippedArguments) GetEndTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WebhookV1BackupJobSkippedArguments) SetEndTime(v WebhookV1ProtobufTimestamp)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WebhookV1BackupJobSkippedArguments) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUserMessage

`func (o *WebhookV1BackupJobSkippedArguments) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *WebhookV1BackupJobSkippedArguments) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *WebhookV1BackupJobSkippedArguments) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *WebhookV1BackupJobSkippedArguments) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.

### GetDetails

`func (o *WebhookV1BackupJobSkippedArguments) GetDetails() WebhookV1BackupJobDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WebhookV1BackupJobSkippedArguments) GetDetailsOk() (*WebhookV1BackupJobDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WebhookV1BackupJobSkippedArguments) SetDetails(v WebhookV1BackupJobDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WebhookV1BackupJobSkippedArguments) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


