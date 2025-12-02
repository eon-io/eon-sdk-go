# JobStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | [**JobStepName**](JobStepName.md) |  | 
**Status** | [**JobStepStatus**](JobStepStatus.md) |  | 
**StartTime** | Pointer to **NullableTime** | Date and time the step started. | [optional] 
**EndTime** | Pointer to **NullableTime** | Date and time the step completed. | [optional] 

## Methods

### NewJobStep

`func NewJobStep(stepName JobStepName, status JobStepStatus, ) *JobStep`

NewJobStep instantiates a new JobStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStepWithDefaults

`func NewJobStepWithDefaults() *JobStep`

NewJobStepWithDefaults instantiates a new JobStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepName

`func (o *JobStep) GetStepName() JobStepName`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *JobStep) GetStepNameOk() (*JobStepName, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *JobStep) SetStepName(v JobStepName)`

SetStepName sets StepName field to given value.


### GetStatus

`func (o *JobStep) GetStatus() JobStepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobStep) GetStatusOk() (*JobStepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobStep) SetStatus(v JobStepStatus)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *JobStep) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobStep) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobStep) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobStep) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *JobStep) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *JobStep) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *JobStep) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *JobStep) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *JobStep) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *JobStep) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *JobStep) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *JobStep) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


