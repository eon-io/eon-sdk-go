# HighFrequencyBackupScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to [**HighFrequencyBackupScheduleFrequency**](HighFrequencyBackupScheduleFrequency.md) |  | [optional] 
**IntervalConfig** | Pointer to [**NullableHighFrequencyIntervalConfig**](HighFrequencyIntervalConfig.md) |  | [optional] 
**DailyConfig** | Pointer to [**NullableDailyConfig**](DailyConfig.md) |  | [optional] 
**WeeklyConfig** | Pointer to [**NullableWeeklyConfig**](WeeklyConfig.md) |  | [optional] 
**MonthlyConfig** | Pointer to [**NullableMonthlyConfig**](MonthlyConfig.md) |  | [optional] 
**AnnuallyConfig** | Pointer to [**NullableAnnuallyConfig**](AnnuallyConfig.md) |  | [optional] 

## Methods

### NewHighFrequencyBackupScheduleConfig

`func NewHighFrequencyBackupScheduleConfig() *HighFrequencyBackupScheduleConfig`

NewHighFrequencyBackupScheduleConfig instantiates a new HighFrequencyBackupScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighFrequencyBackupScheduleConfigWithDefaults

`func NewHighFrequencyBackupScheduleConfigWithDefaults() *HighFrequencyBackupScheduleConfig`

NewHighFrequencyBackupScheduleConfigWithDefaults instantiates a new HighFrequencyBackupScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *HighFrequencyBackupScheduleConfig) GetFrequency() HighFrequencyBackupScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *HighFrequencyBackupScheduleConfig) GetFrequencyOk() (*HighFrequencyBackupScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *HighFrequencyBackupScheduleConfig) SetFrequency(v HighFrequencyBackupScheduleFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *HighFrequencyBackupScheduleConfig) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetIntervalConfig

`func (o *HighFrequencyBackupScheduleConfig) GetIntervalConfig() HighFrequencyIntervalConfig`

GetIntervalConfig returns the IntervalConfig field if non-nil, zero value otherwise.

### GetIntervalConfigOk

`func (o *HighFrequencyBackupScheduleConfig) GetIntervalConfigOk() (*HighFrequencyIntervalConfig, bool)`

GetIntervalConfigOk returns a tuple with the IntervalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalConfig

`func (o *HighFrequencyBackupScheduleConfig) SetIntervalConfig(v HighFrequencyIntervalConfig)`

SetIntervalConfig sets IntervalConfig field to given value.

### HasIntervalConfig

`func (o *HighFrequencyBackupScheduleConfig) HasIntervalConfig() bool`

HasIntervalConfig returns a boolean if a field has been set.

### SetIntervalConfigNil

`func (o *HighFrequencyBackupScheduleConfig) SetIntervalConfigNil(b bool)`

 SetIntervalConfigNil sets the value for IntervalConfig to be an explicit nil

### UnsetIntervalConfig
`func (o *HighFrequencyBackupScheduleConfig) UnsetIntervalConfig()`

UnsetIntervalConfig ensures that no value is present for IntervalConfig, not even an explicit nil
### GetDailyConfig

`func (o *HighFrequencyBackupScheduleConfig) GetDailyConfig() DailyConfig`

GetDailyConfig returns the DailyConfig field if non-nil, zero value otherwise.

### GetDailyConfigOk

`func (o *HighFrequencyBackupScheduleConfig) GetDailyConfigOk() (*DailyConfig, bool)`

GetDailyConfigOk returns a tuple with the DailyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyConfig

`func (o *HighFrequencyBackupScheduleConfig) SetDailyConfig(v DailyConfig)`

SetDailyConfig sets DailyConfig field to given value.

### HasDailyConfig

`func (o *HighFrequencyBackupScheduleConfig) HasDailyConfig() bool`

HasDailyConfig returns a boolean if a field has been set.

### SetDailyConfigNil

`func (o *HighFrequencyBackupScheduleConfig) SetDailyConfigNil(b bool)`

 SetDailyConfigNil sets the value for DailyConfig to be an explicit nil

### UnsetDailyConfig
`func (o *HighFrequencyBackupScheduleConfig) UnsetDailyConfig()`

UnsetDailyConfig ensures that no value is present for DailyConfig, not even an explicit nil
### GetWeeklyConfig

`func (o *HighFrequencyBackupScheduleConfig) GetWeeklyConfig() WeeklyConfig`

GetWeeklyConfig returns the WeeklyConfig field if non-nil, zero value otherwise.

### GetWeeklyConfigOk

`func (o *HighFrequencyBackupScheduleConfig) GetWeeklyConfigOk() (*WeeklyConfig, bool)`

GetWeeklyConfigOk returns a tuple with the WeeklyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyConfig

`func (o *HighFrequencyBackupScheduleConfig) SetWeeklyConfig(v WeeklyConfig)`

SetWeeklyConfig sets WeeklyConfig field to given value.

### HasWeeklyConfig

`func (o *HighFrequencyBackupScheduleConfig) HasWeeklyConfig() bool`

HasWeeklyConfig returns a boolean if a field has been set.

### SetWeeklyConfigNil

`func (o *HighFrequencyBackupScheduleConfig) SetWeeklyConfigNil(b bool)`

 SetWeeklyConfigNil sets the value for WeeklyConfig to be an explicit nil

### UnsetWeeklyConfig
`func (o *HighFrequencyBackupScheduleConfig) UnsetWeeklyConfig()`

UnsetWeeklyConfig ensures that no value is present for WeeklyConfig, not even an explicit nil
### GetMonthlyConfig

`func (o *HighFrequencyBackupScheduleConfig) GetMonthlyConfig() MonthlyConfig`

GetMonthlyConfig returns the MonthlyConfig field if non-nil, zero value otherwise.

### GetMonthlyConfigOk

`func (o *HighFrequencyBackupScheduleConfig) GetMonthlyConfigOk() (*MonthlyConfig, bool)`

GetMonthlyConfigOk returns a tuple with the MonthlyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConfig

`func (o *HighFrequencyBackupScheduleConfig) SetMonthlyConfig(v MonthlyConfig)`

SetMonthlyConfig sets MonthlyConfig field to given value.

### HasMonthlyConfig

`func (o *HighFrequencyBackupScheduleConfig) HasMonthlyConfig() bool`

HasMonthlyConfig returns a boolean if a field has been set.

### SetMonthlyConfigNil

`func (o *HighFrequencyBackupScheduleConfig) SetMonthlyConfigNil(b bool)`

 SetMonthlyConfigNil sets the value for MonthlyConfig to be an explicit nil

### UnsetMonthlyConfig
`func (o *HighFrequencyBackupScheduleConfig) UnsetMonthlyConfig()`

UnsetMonthlyConfig ensures that no value is present for MonthlyConfig, not even an explicit nil
### GetAnnuallyConfig

`func (o *HighFrequencyBackupScheduleConfig) GetAnnuallyConfig() AnnuallyConfig`

GetAnnuallyConfig returns the AnnuallyConfig field if non-nil, zero value otherwise.

### GetAnnuallyConfigOk

`func (o *HighFrequencyBackupScheduleConfig) GetAnnuallyConfigOk() (*AnnuallyConfig, bool)`

GetAnnuallyConfigOk returns a tuple with the AnnuallyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuallyConfig

`func (o *HighFrequencyBackupScheduleConfig) SetAnnuallyConfig(v AnnuallyConfig)`

SetAnnuallyConfig sets AnnuallyConfig field to given value.

### HasAnnuallyConfig

`func (o *HighFrequencyBackupScheduleConfig) HasAnnuallyConfig() bool`

HasAnnuallyConfig returns a boolean if a field has been set.

### SetAnnuallyConfigNil

`func (o *HighFrequencyBackupScheduleConfig) SetAnnuallyConfigNil(b bool)`

 SetAnnuallyConfigNil sets the value for AnnuallyConfig to be an explicit nil

### UnsetAnnuallyConfig
`func (o *HighFrequencyBackupScheduleConfig) UnsetAnnuallyConfig()`

UnsetAnnuallyConfig ensures that no value is present for AnnuallyConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


