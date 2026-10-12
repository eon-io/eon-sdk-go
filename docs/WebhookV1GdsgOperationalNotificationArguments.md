# WebhookV1GdsgOperationalNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | Environment ID. | [optional] 
**EnvironmentName** | Pointer to **string** | Environment name. | [optional] 
**Region** | Pointer to **string** | Region of the environment. | [optional] 
**Stage** | Pointer to **string** | Pipeline stage the alert refers to. | [optional] 
**Signal** | Pointer to **string** | Signal that triggered the alert. | [optional] 
**CurrentValueMs** | Pointer to **int64** | Current value of the signal, in milliseconds. | [optional] 
**SlaSeconds** | Pointer to **int64** | SLA threshold, in seconds. | [optional] 
**LastReportTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**ReportIntervalMs** | Pointer to **int64** | Expected report interval, in milliseconds. | [optional] 
**SampleError** | Pointer to **string** | Sample error message. | [optional] 
**ConsoleLink** | Pointer to **string** | Link to the environment in the Eon console. | [optional] 

## Methods

### NewWebhookV1GdsgOperationalNotificationArguments

`func NewWebhookV1GdsgOperationalNotificationArguments() *WebhookV1GdsgOperationalNotificationArguments`

NewWebhookV1GdsgOperationalNotificationArguments instantiates a new WebhookV1GdsgOperationalNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1GdsgOperationalNotificationArgumentsWithDefaults

`func NewWebhookV1GdsgOperationalNotificationArgumentsWithDefaults() *WebhookV1GdsgOperationalNotificationArguments`

NewWebhookV1GdsgOperationalNotificationArgumentsWithDefaults instantiates a new WebhookV1GdsgOperationalNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetRegion

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStage

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetSignal

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetCurrentValueMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetCurrentValueMs() int64`

GetCurrentValueMs returns the CurrentValueMs field if non-nil, zero value otherwise.

### GetCurrentValueMsOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetCurrentValueMsOk() (*int64, bool)`

GetCurrentValueMsOk returns a tuple with the CurrentValueMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValueMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetCurrentValueMs(v int64)`

SetCurrentValueMs sets CurrentValueMs field to given value.

### HasCurrentValueMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasCurrentValueMs() bool`

HasCurrentValueMs returns a boolean if a field has been set.

### GetSlaSeconds

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSlaSeconds() int64`

GetSlaSeconds returns the SlaSeconds field if non-nil, zero value otherwise.

### GetSlaSecondsOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSlaSecondsOk() (*int64, bool)`

GetSlaSecondsOk returns a tuple with the SlaSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaSeconds

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetSlaSeconds(v int64)`

SetSlaSeconds sets SlaSeconds field to given value.

### HasSlaSeconds

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasSlaSeconds() bool`

HasSlaSeconds returns a boolean if a field has been set.

### GetLastReportTime

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetLastReportTime() WebhookV1ProtobufTimestamp`

GetLastReportTime returns the LastReportTime field if non-nil, zero value otherwise.

### GetLastReportTimeOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetLastReportTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetLastReportTimeOk returns a tuple with the LastReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportTime

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetLastReportTime(v WebhookV1ProtobufTimestamp)`

SetLastReportTime sets LastReportTime field to given value.

### HasLastReportTime

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasLastReportTime() bool`

HasLastReportTime returns a boolean if a field has been set.

### GetReportIntervalMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetReportIntervalMs() int64`

GetReportIntervalMs returns the ReportIntervalMs field if non-nil, zero value otherwise.

### GetReportIntervalMsOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetReportIntervalMsOk() (*int64, bool)`

GetReportIntervalMsOk returns a tuple with the ReportIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetReportIntervalMs(v int64)`

SetReportIntervalMs sets ReportIntervalMs field to given value.

### HasReportIntervalMs

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasReportIntervalMs() bool`

HasReportIntervalMs returns a boolean if a field has been set.

### GetSampleError

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSampleError() string`

GetSampleError returns the SampleError field if non-nil, zero value otherwise.

### GetSampleErrorOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetSampleErrorOk() (*string, bool)`

GetSampleErrorOk returns a tuple with the SampleError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleError

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetSampleError(v string)`

SetSampleError sets SampleError field to given value.

### HasSampleError

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasSampleError() bool`

HasSampleError returns a boolean if a field has been set.

### GetConsoleLink

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetConsoleLink() string`

GetConsoleLink returns the ConsoleLink field if non-nil, zero value otherwise.

### GetConsoleLinkOk

`func (o *WebhookV1GdsgOperationalNotificationArguments) GetConsoleLinkOk() (*string, bool)`

GetConsoleLinkOk returns a tuple with the ConsoleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLink

`func (o *WebhookV1GdsgOperationalNotificationArguments) SetConsoleLink(v string)`

SetConsoleLink sets ConsoleLink field to given value.

### HasConsoleLink

`func (o *WebhookV1GdsgOperationalNotificationArguments) HasConsoleLink() bool`

HasConsoleLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


