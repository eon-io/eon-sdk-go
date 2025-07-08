# MonthlyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **[]int32** | Days of the month to schedule backups.  | [optional] 
**LastDayOfMonth** | Pointer to **bool** | Whether to schedule backups on the last day of the month. Used only when &#x60;dayOfMonth&#x60; is not specified.  | [optional] 
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

### GetDayOfMonth

`func (o *MonthlyConfig) GetDayOfMonth() []int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MonthlyConfig) GetDayOfMonthOk() (*[]int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MonthlyConfig) SetDayOfMonth(v []int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MonthlyConfig) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetLastDayOfMonth

`func (o *MonthlyConfig) GetLastDayOfMonth() bool`

GetLastDayOfMonth returns the LastDayOfMonth field if non-nil, zero value otherwise.

### GetLastDayOfMonthOk

`func (o *MonthlyConfig) GetLastDayOfMonthOk() (*bool, bool)`

GetLastDayOfMonthOk returns a tuple with the LastDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayOfMonth

`func (o *MonthlyConfig) SetLastDayOfMonth(v bool)`

SetLastDayOfMonth sets LastDayOfMonth field to given value.

### HasLastDayOfMonth

`func (o *MonthlyConfig) HasLastDayOfMonth() bool`

HasLastDayOfMonth returns a boolean if a field has been set.

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


