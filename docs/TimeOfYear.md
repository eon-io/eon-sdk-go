# TimeOfYear

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **int32** |  | [default to 1]
**DayOfMonth** | **int32** |  | [default to 1]

## Methods

### NewTimeOfYear

`func NewTimeOfYear(month int32, dayOfMonth int32, ) *TimeOfYear`

NewTimeOfYear instantiates a new TimeOfYear object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOfYearWithDefaults

`func NewTimeOfYearWithDefaults() *TimeOfYear`

NewTimeOfYearWithDefaults instantiates a new TimeOfYear object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *TimeOfYear) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *TimeOfYear) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *TimeOfYear) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDayOfMonth

`func (o *TimeOfYear) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *TimeOfYear) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *TimeOfYear) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


