# DailyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOfDay** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**StartWindowMinutes** | Pointer to **int32** | The window of time after the start time you want the backup to start, in minutes. Defaults to &#x60;240&#x60; (4 hours).  | [optional] [default to 240]

## Methods

### NewDailyConfig

`func NewDailyConfig() *DailyConfig`

NewDailyConfig instantiates a new DailyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyConfigWithDefaults

`func NewDailyConfigWithDefaults() *DailyConfig`

NewDailyConfigWithDefaults instantiates a new DailyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOfDay

`func (o *DailyConfig) GetTimeOfDay() TimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *DailyConfig) GetTimeOfDayOk() (*TimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *DailyConfig) SetTimeOfDay(v TimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *DailyConfig) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetStartWindowMinutes

`func (o *DailyConfig) GetStartWindowMinutes() int32`

GetStartWindowMinutes returns the StartWindowMinutes field if non-nil, zero value otherwise.

### GetStartWindowMinutesOk

`func (o *DailyConfig) GetStartWindowMinutesOk() (*int32, bool)`

GetStartWindowMinutesOk returns a tuple with the StartWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowMinutes

`func (o *DailyConfig) SetStartWindowMinutes(v int32)`

SetStartWindowMinutes sets StartWindowMinutes field to given value.

### HasStartWindowMinutes

`func (o *DailyConfig) HasStartWindowMinutes() bool`

HasStartWindowMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


