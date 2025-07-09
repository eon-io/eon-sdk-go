# StandardBackupScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | [**StandardBackupScheduleFrequency**](StandardBackupScheduleFrequency.md) |  | 
**IntervalConfig** | Pointer to [**NullableStandardIntervalConfig**](StandardIntervalConfig.md) |  | [optional] 
**DailyConfig** | Pointer to [**NullableDailyConfig**](DailyConfig.md) |  | [optional] 
**WeeklyConfig** | Pointer to [**NullableWeeklyConfig**](WeeklyConfig.md) |  | [optional] 
**MonthlyConfig** | Pointer to [**NullableMonthlyConfig**](MonthlyConfig.md) |  | [optional] 
**AnnuallyConfig** | Pointer to [**NullableAnnuallyConfig**](AnnuallyConfig.md) |  | [optional] 

## Methods

### NewStandardBackupScheduleConfig

`func NewStandardBackupScheduleConfig(frequency StandardBackupScheduleFrequency, ) *StandardBackupScheduleConfig`

NewStandardBackupScheduleConfig instantiates a new StandardBackupScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardBackupScheduleConfigWithDefaults

`func NewStandardBackupScheduleConfigWithDefaults() *StandardBackupScheduleConfig`

NewStandardBackupScheduleConfigWithDefaults instantiates a new StandardBackupScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *StandardBackupScheduleConfig) GetFrequency() StandardBackupScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StandardBackupScheduleConfig) GetFrequencyOk() (*StandardBackupScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StandardBackupScheduleConfig) SetFrequency(v StandardBackupScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetIntervalConfig

`func (o *StandardBackupScheduleConfig) GetIntervalConfig() StandardIntervalConfig`

GetIntervalConfig returns the IntervalConfig field if non-nil, zero value otherwise.

### GetIntervalConfigOk

`func (o *StandardBackupScheduleConfig) GetIntervalConfigOk() (*StandardIntervalConfig, bool)`

GetIntervalConfigOk returns a tuple with the IntervalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalConfig

`func (o *StandardBackupScheduleConfig) SetIntervalConfig(v StandardIntervalConfig)`

SetIntervalConfig sets IntervalConfig field to given value.

### HasIntervalConfig

`func (o *StandardBackupScheduleConfig) HasIntervalConfig() bool`

HasIntervalConfig returns a boolean if a field has been set.

### SetIntervalConfigNil

`func (o *StandardBackupScheduleConfig) SetIntervalConfigNil(b bool)`

 SetIntervalConfigNil sets the value for IntervalConfig to be an explicit nil

### UnsetIntervalConfig
`func (o *StandardBackupScheduleConfig) UnsetIntervalConfig()`

UnsetIntervalConfig ensures that no value is present for IntervalConfig, not even an explicit nil
### GetDailyConfig

`func (o *StandardBackupScheduleConfig) GetDailyConfig() DailyConfig`

GetDailyConfig returns the DailyConfig field if non-nil, zero value otherwise.

### GetDailyConfigOk

`func (o *StandardBackupScheduleConfig) GetDailyConfigOk() (*DailyConfig, bool)`

GetDailyConfigOk returns a tuple with the DailyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyConfig

`func (o *StandardBackupScheduleConfig) SetDailyConfig(v DailyConfig)`

SetDailyConfig sets DailyConfig field to given value.

### HasDailyConfig

`func (o *StandardBackupScheduleConfig) HasDailyConfig() bool`

HasDailyConfig returns a boolean if a field has been set.

### SetDailyConfigNil

`func (o *StandardBackupScheduleConfig) SetDailyConfigNil(b bool)`

 SetDailyConfigNil sets the value for DailyConfig to be an explicit nil

### UnsetDailyConfig
`func (o *StandardBackupScheduleConfig) UnsetDailyConfig()`

UnsetDailyConfig ensures that no value is present for DailyConfig, not even an explicit nil
### GetWeeklyConfig

`func (o *StandardBackupScheduleConfig) GetWeeklyConfig() WeeklyConfig`

GetWeeklyConfig returns the WeeklyConfig field if non-nil, zero value otherwise.

### GetWeeklyConfigOk

`func (o *StandardBackupScheduleConfig) GetWeeklyConfigOk() (*WeeklyConfig, bool)`

GetWeeklyConfigOk returns a tuple with the WeeklyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyConfig

`func (o *StandardBackupScheduleConfig) SetWeeklyConfig(v WeeklyConfig)`

SetWeeklyConfig sets WeeklyConfig field to given value.

### HasWeeklyConfig

`func (o *StandardBackupScheduleConfig) HasWeeklyConfig() bool`

HasWeeklyConfig returns a boolean if a field has been set.

### SetWeeklyConfigNil

`func (o *StandardBackupScheduleConfig) SetWeeklyConfigNil(b bool)`

 SetWeeklyConfigNil sets the value for WeeklyConfig to be an explicit nil

### UnsetWeeklyConfig
`func (o *StandardBackupScheduleConfig) UnsetWeeklyConfig()`

UnsetWeeklyConfig ensures that no value is present for WeeklyConfig, not even an explicit nil
### GetMonthlyConfig

`func (o *StandardBackupScheduleConfig) GetMonthlyConfig() MonthlyConfig`

GetMonthlyConfig returns the MonthlyConfig field if non-nil, zero value otherwise.

### GetMonthlyConfigOk

`func (o *StandardBackupScheduleConfig) GetMonthlyConfigOk() (*MonthlyConfig, bool)`

GetMonthlyConfigOk returns a tuple with the MonthlyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConfig

`func (o *StandardBackupScheduleConfig) SetMonthlyConfig(v MonthlyConfig)`

SetMonthlyConfig sets MonthlyConfig field to given value.

### HasMonthlyConfig

`func (o *StandardBackupScheduleConfig) HasMonthlyConfig() bool`

HasMonthlyConfig returns a boolean if a field has been set.

### SetMonthlyConfigNil

`func (o *StandardBackupScheduleConfig) SetMonthlyConfigNil(b bool)`

 SetMonthlyConfigNil sets the value for MonthlyConfig to be an explicit nil

### UnsetMonthlyConfig
`func (o *StandardBackupScheduleConfig) UnsetMonthlyConfig()`

UnsetMonthlyConfig ensures that no value is present for MonthlyConfig, not even an explicit nil
### GetAnnuallyConfig

`func (o *StandardBackupScheduleConfig) GetAnnuallyConfig() AnnuallyConfig`

GetAnnuallyConfig returns the AnnuallyConfig field if non-nil, zero value otherwise.

### GetAnnuallyConfigOk

`func (o *StandardBackupScheduleConfig) GetAnnuallyConfigOk() (*AnnuallyConfig, bool)`

GetAnnuallyConfigOk returns a tuple with the AnnuallyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuallyConfig

`func (o *StandardBackupScheduleConfig) SetAnnuallyConfig(v AnnuallyConfig)`

SetAnnuallyConfig sets AnnuallyConfig field to given value.

### HasAnnuallyConfig

`func (o *StandardBackupScheduleConfig) HasAnnuallyConfig() bool`

HasAnnuallyConfig returns a boolean if a field has been set.

### SetAnnuallyConfigNil

`func (o *StandardBackupScheduleConfig) SetAnnuallyConfigNil(b bool)`

 SetAnnuallyConfigNil sets the value for AnnuallyConfig to be an explicit nil

### UnsetAnnuallyConfig
`func (o *StandardBackupScheduleConfig) UnsetAnnuallyConfig()`

UnsetAnnuallyConfig ensures that no value is present for AnnuallyConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


