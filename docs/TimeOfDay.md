# TimeOfDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | **int32** | Hour of the day. | [default to 0]
**Minute** | **int32** | Minute of the hour. | [default to 0]

## Methods

### NewTimeOfDay

`func NewTimeOfDay(hour int32, minute int32, ) *TimeOfDay`

NewTimeOfDay instantiates a new TimeOfDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfDayWithDefaults

`func NewTimeOfDayWithDefaults() *TimeOfDay`

NewTimeOfDayWithDefaults instantiates a new TimeOfDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHour

`func (o *TimeOfDay) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *TimeOfDay) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *TimeOfDay) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMinute

`func (o *TimeOfDay) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *TimeOfDay) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *TimeOfDay) SetMinute(v int32)`

SetMinute sets Minute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


