# WeeklyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | Pointer to **[]string** | Days of the week to schedule backups.  | [optional] [default to ["SUNDAY"]]
**TimeOfDay** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**StartWindowMinutes** | Pointer to **int32** | The window of time after the start time you want the backup to start, in minutes. Defaults to &#x60;240&#x60; (4 hours).  | [optional] [default to 240]

## Methods

### NewWeeklyConfig

`func NewWeeklyConfig() *WeeklyConfig`

NewWeeklyConfig instantiates a new WeeklyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyConfigWithDefaults

`func NewWeeklyConfigWithDefaults() *WeeklyConfig`

NewWeeklyConfigWithDefaults instantiates a new WeeklyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfWeek

`func (o *WeeklyConfig) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *WeeklyConfig) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *WeeklyConfig) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *WeeklyConfig) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *WeeklyConfig) GetTimeOfDay() TimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *WeeklyConfig) GetTimeOfDayOk() (*TimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *WeeklyConfig) SetTimeOfDay(v TimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *WeeklyConfig) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetStartWindowMinutes

`func (o *WeeklyConfig) GetStartWindowMinutes() int32`

GetStartWindowMinutes returns the StartWindowMinutes field if non-nil, zero value otherwise.

### GetStartWindowMinutesOk

`func (o *WeeklyConfig) GetStartWindowMinutesOk() (*int32, bool)`

GetStartWindowMinutesOk returns a tuple with the StartWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowMinutes

`func (o *WeeklyConfig) SetStartWindowMinutes(v int32)`

SetStartWindowMinutes sets StartWindowMinutes field to given value.

### HasStartWindowMinutes

`func (o *WeeklyConfig) HasStartWindowMinutes() bool`

HasStartWindowMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


