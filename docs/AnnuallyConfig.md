# AnnuallyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOfDay** | Pointer to [**TimeOfDay**](TimeOfDay.md) |  | [optional] 
**TimeOfYear** | Pointer to [**TimeOfYear**](TimeOfYear.md) |  | [optional] 
**StartWindowMinutes** | Pointer to **int32** | The window of time after the start time you want the backup to start, in minutes. Defaults to &#x60;240&#x60; (4 hours).  | [optional] [default to 240]

## Methods

### NewAnnuallyConfig

`func NewAnnuallyConfig() *AnnuallyConfig`

NewAnnuallyConfig instantiates a new AnnuallyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnuallyConfigWithDefaults

`func NewAnnuallyConfigWithDefaults() *AnnuallyConfig`

NewAnnuallyConfigWithDefaults instantiates a new AnnuallyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOfDay

`func (o *AnnuallyConfig) GetTimeOfDay() TimeOfDay`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *AnnuallyConfig) GetTimeOfDayOk() (*TimeOfDay, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *AnnuallyConfig) SetTimeOfDay(v TimeOfDay)`

SetTimeOfDay sets TimeOfDay field to given value.

### HasTimeOfDay

`func (o *AnnuallyConfig) HasTimeOfDay() bool`

HasTimeOfDay returns a boolean if a field has been set.

### GetTimeOfYear

`func (o *AnnuallyConfig) GetTimeOfYear() TimeOfYear`

GetTimeOfYear returns the TimeOfYear field if non-nil, zero value otherwise.

### GetTimeOfYearOk

`func (o *AnnuallyConfig) GetTimeOfYearOk() (*TimeOfYear, bool)`

GetTimeOfYearOk returns a tuple with the TimeOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfYear

`func (o *AnnuallyConfig) SetTimeOfYear(v TimeOfYear)`

SetTimeOfYear sets TimeOfYear field to given value.

### HasTimeOfYear

`func (o *AnnuallyConfig) HasTimeOfYear() bool`

HasTimeOfYear returns a boolean if a field has been set.

### GetStartWindowMinutes

`func (o *AnnuallyConfig) GetStartWindowMinutes() int32`

GetStartWindowMinutes returns the StartWindowMinutes field if non-nil, zero value otherwise.

### GetStartWindowMinutesOk

`func (o *AnnuallyConfig) GetStartWindowMinutesOk() (*int32, bool)`

GetStartWindowMinutesOk returns a tuple with the StartWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWindowMinutes

`func (o *AnnuallyConfig) SetStartWindowMinutes(v int32)`

SetStartWindowMinutes sets StartWindowMinutes field to given value.

### HasStartWindowMinutes

`func (o *AnnuallyConfig) HasStartWindowMinutes() bool`

HasStartWindowMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


