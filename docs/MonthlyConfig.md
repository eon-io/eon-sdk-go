# MonthlyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfMonth** | Pointer to **[]int32** | Days of the month to schedule backups.  | [optional] 
**TimeOfDay** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**StartWindowMinutes** | Pointer to **int32** | The window of time after the start time you want the backup to start, in minutes. Defaults to &#x60;240&#x60; (4 hours).  | [optional] [default to 240]

## Methods

### NewMonthlyConfig

`func NewMonthlyConfig() *MonthlyConfig`

NewMonthlyConfig instantiates a new MonthlyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyConfigWithDefaults

`func NewMonthlyConfigWithDefaults() *MonthlyConfig`

NewMonthlyConfigWithDefaults instantiates a new MonthlyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfMonth

`func (o *MonthlyConfig) GetDaysOfMonth() []int32`

GetDaysOfMonth returns the DaysOfMonth field if non-nil, zero value otherwise.

### GetDaysOfMonthOk

`func (o *MonthlyConfig) GetDaysOfMonthOk() (*[]int32, bool)`

GetDaysOfMonthOk returns a tuple with the DaysOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfMonth

`func (o *MonthlyConfig) SetDaysOfMonth(v []int32)`

SetDaysOfMonth sets DaysOfMonth field to given value.

### HasDaysOfMonth

`func (o *MonthlyConfig) HasDaysOfMonth() bool`

HasDaysOfMonth returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *MonthlyConfig) GetTimeOfDay() TimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *MonthlyConfig) GetTimeOfDayOk() (*TimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *MonthlyConfig) SetTimeOfDay(v TimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *MonthlyConfig) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetStartWindowMinutes

`func (o *MonthlyConfig) GetStartWindowMinutes() int32`

GetStartWindowMinutes returns the StartWindowMinutes field if non-nil, zero value otherwise.

### GetStartWindowMinutesOk

`func (o *MonthlyConfig) GetStartWindowMinutesOk() (*int32, bool)`

GetStartWindowMinutesOk returns a tuple with the StartWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowMinutes

`func (o *MonthlyConfig) SetStartWindowMinutes(v int32)`

SetStartWindowMinutes sets StartWindowMinutes field to given value.

### HasStartWindowMinutes

`func (o *MonthlyConfig) HasStartWindowMinutes() bool`

HasStartWindowMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


