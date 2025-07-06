# JobExecutionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | Job ID. | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**CreatedTime** | **time.Time** | Date and time the job was created. | 
**ExpectedStartTime** | Pointer to **NullableTime** | Date and time the job is expected to start. Returned only for pending jobs.  | [optional] 
**StartTime** | Pointer to **NullableTime** | Date and time the job started. | [optional] 
**EndTime** | Pointer to **NullableTime** | Date and time the job finished. | [optional] 
**DurationSeconds** | Pointer to **NullableInt64** | How long the job took, in seconds. | [optional] 
**StatusMessage** | Pointer to **string** | Message that gives additional details about the job status, if applicable. | [optional] 

## Methods

### NewJobExecutionDetails

`func NewJobExecutionDetails(jobId string, status JobStatus, createdTime time.Time, ) *JobExecutionDetails`

NewJobExecutionDetails instantiates a new JobExecutionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobExecutionDetailsWithDefaults

`func NewJobExecutionDetailsWithDefaults() *JobExecutionDetails`

NewJobExecutionDetailsWithDefaults instantiates a new JobExecutionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobExecutionDetails) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobExecutionDetails) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobExecutionDetails) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *JobExecutionDetails) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobExecutionDetails) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobExecutionDetails) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetCreatedTime

`func (o *JobExecutionDetails) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *JobExecutionDetails) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *JobExecutionDetails) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetExpectedStartTime

`func (o *JobExecutionDetails) GetExpectedStartTime() time.Time`

GetExpectedStartTime returns the ExpectedStartTime field if non-nil, zero value otherwise.

### GetExpectedStartTimeOk

`func (o *JobExecutionDetails) GetExpectedStartTimeOk() (*time.Time, bool)`

GetExpectedStartTimeOk returns a tuple with the ExpectedStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartTime

`func (o *JobExecutionDetails) SetExpectedStartTime(v time.Time)`

SetExpectedStartTime sets ExpectedStartTime field to given value.

### HasExpectedStartTime

`func (o *JobExecutionDetails) HasExpectedStartTime() bool`

HasExpectedStartTime returns a boolean if a field has been set.

### SetExpectedStartTimeNil

`func (o *JobExecutionDetails) SetExpectedStartTimeNil(b bool)`

 SetExpectedStartTimeNil sets the value for ExpectedStartTime to be an explicit nil

### UnsetExpectedStartTime
`func (o *JobExecutionDetails) UnsetExpectedStartTime()`

UnsetExpectedStartTime ensures that no value is present for ExpectedStartTime, not even an explicit nil
### GetStartTime

`func (o *JobExecutionDetails) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobExecutionDetails) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobExecutionDetails) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobExecutionDetails) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *JobExecutionDetails) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *JobExecutionDetails) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *JobExecutionDetails) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobExecutionDetails) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobExecutionDetails) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobExecutionDetails) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *JobExecutionDetails) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *JobExecutionDetails) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDurationSeconds

`func (o *JobExecutionDetails) GetDurationSeconds() int64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *JobExecutionDetails) GetDurationSecondsOk() (*int64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *JobExecutionDetails) SetDurationSeconds(v int64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *JobExecutionDetails) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *JobExecutionDetails) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *JobExecutionDetails) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetStatusMessage

`func (o *JobExecutionDetails) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *JobExecutionDetails) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *JobExecutionDetails) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *JobExecutionDetails) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


