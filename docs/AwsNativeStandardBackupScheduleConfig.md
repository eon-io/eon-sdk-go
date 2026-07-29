# AwsNativeStandardBackupScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | [**StandardBackupScheduleFrequency**](StandardBackupScheduleFrequency.md) |  | 
**IntervalConfig** | Pointer to [**NullableAwsNativeStandardIntervalConfig**](AwsNativeStandardIntervalConfig.md) |  | [optional] 
**DailyConfig** | Pointer to [**NullableDailyConfig**](DailyConfig.md) |  | [optional] 
**WeeklyConfig** | Pointer to [**NullableWeeklyConfig**](WeeklyConfig.md) |  | [optional] 
**MonthlyConfig** | Pointer to [**NullableMonthlyConfig**](MonthlyConfig.md) |  | [optional] 
**AnnuallyConfig** | Pointer to [**NullableAnnuallyConfig**](AnnuallyConfig.md) |  | [optional] 

## Methods

### NewAwsNativeStandardBackupScheduleConfig

`func NewAwsNativeStandardBackupScheduleConfig(frequency StandardBackupScheduleFrequency, ) *AwsNativeStandardBackupScheduleConfig`

NewAwsNativeStandardBackupScheduleConfig instantiates a new AwsNativeStandardBackupScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativeStandardBackupScheduleConfigWithDefaults

`func NewAwsNativeStandardBackupScheduleConfigWithDefaults() *AwsNativeStandardBackupScheduleConfig`

NewAwsNativeStandardBackupScheduleConfigWithDefaults instantiates a new AwsNativeStandardBackupScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *AwsNativeStandardBackupScheduleConfig) GetFrequency() StandardBackupScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetFrequencyOk() (*StandardBackupScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AwsNativeStandardBackupScheduleConfig) SetFrequency(v StandardBackupScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetIntervalConfig

`func (o *AwsNativeStandardBackupScheduleConfig) GetIntervalConfig() AwsNativeStandardIntervalConfig`

GetIntervalConfig returns the IntervalConfig field if non-nil, zero value otherwise.

### GetIntervalConfigOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetIntervalConfigOk() (*AwsNativeStandardIntervalConfig, bool)`

GetIntervalConfigOk returns a tuple with the IntervalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalConfig

`func (o *AwsNativeStandardBackupScheduleConfig) SetIntervalConfig(v AwsNativeStandardIntervalConfig)`

SetIntervalConfig sets IntervalConfig field to given value.

### HasIntervalConfig

`func (o *AwsNativeStandardBackupScheduleConfig) HasIntervalConfig() bool`

HasIntervalConfig returns a boolean if a field has been set.

### SetIntervalConfigNil

`func (o *AwsNativeStandardBackupScheduleConfig) SetIntervalConfigNil(b bool)`

 SetIntervalConfigNil sets the value for IntervalConfig to be an explicit nil

### UnsetIntervalConfig
`func (o *AwsNativeStandardBackupScheduleConfig) UnsetIntervalConfig()`

UnsetIntervalConfig ensures that no value is present for IntervalConfig, not even an explicit nil
### GetDailyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) GetDailyConfig() DailyConfig`

GetDailyConfig returns the DailyConfig field if non-nil, zero value otherwise.

### GetDailyConfigOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetDailyConfigOk() (*DailyConfig, bool)`

GetDailyConfigOk returns a tuple with the DailyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) SetDailyConfig(v DailyConfig)`

SetDailyConfig sets DailyConfig field to given value.

### HasDailyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) HasDailyConfig() bool`

HasDailyConfig returns a boolean if a field has been set.

### SetDailyConfigNil

`func (o *AwsNativeStandardBackupScheduleConfig) SetDailyConfigNil(b bool)`

 SetDailyConfigNil sets the value for DailyConfig to be an explicit nil

### UnsetDailyConfig
`func (o *AwsNativeStandardBackupScheduleConfig) UnsetDailyConfig()`

UnsetDailyConfig ensures that no value is present for DailyConfig, not even an explicit nil
### GetWeeklyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) GetWeeklyConfig() WeeklyConfig`

GetWeeklyConfig returns the WeeklyConfig field if non-nil, zero value otherwise.

### GetWeeklyConfigOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetWeeklyConfigOk() (*WeeklyConfig, bool)`

GetWeeklyConfigOk returns a tuple with the WeeklyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) SetWeeklyConfig(v WeeklyConfig)`

SetWeeklyConfig sets WeeklyConfig field to given value.

### HasWeeklyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) HasWeeklyConfig() bool`

HasWeeklyConfig returns a boolean if a field has been set.

### SetWeeklyConfigNil

`func (o *AwsNativeStandardBackupScheduleConfig) SetWeeklyConfigNil(b bool)`

 SetWeeklyConfigNil sets the value for WeeklyConfig to be an explicit nil

### UnsetWeeklyConfig
`func (o *AwsNativeStandardBackupScheduleConfig) UnsetWeeklyConfig()`

UnsetWeeklyConfig ensures that no value is present for WeeklyConfig, not even an explicit nil
### GetMonthlyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) GetMonthlyConfig() MonthlyConfig`

GetMonthlyConfig returns the MonthlyConfig field if non-nil, zero value otherwise.

### GetMonthlyConfigOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetMonthlyConfigOk() (*MonthlyConfig, bool)`

GetMonthlyConfigOk returns a tuple with the MonthlyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) SetMonthlyConfig(v MonthlyConfig)`

SetMonthlyConfig sets MonthlyConfig field to given value.

### HasMonthlyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) HasMonthlyConfig() bool`

HasMonthlyConfig returns a boolean if a field has been set.

### SetMonthlyConfigNil

`func (o *AwsNativeStandardBackupScheduleConfig) SetMonthlyConfigNil(b bool)`

 SetMonthlyConfigNil sets the value for MonthlyConfig to be an explicit nil

### UnsetMonthlyConfig
`func (o *AwsNativeStandardBackupScheduleConfig) UnsetMonthlyConfig()`

UnsetMonthlyConfig ensures that no value is present for MonthlyConfig, not even an explicit nil
### GetAnnuallyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) GetAnnuallyConfig() AnnuallyConfig`

GetAnnuallyConfig returns the AnnuallyConfig field if non-nil, zero value otherwise.

### GetAnnuallyConfigOk

`func (o *AwsNativeStandardBackupScheduleConfig) GetAnnuallyConfigOk() (*AnnuallyConfig, bool)`

GetAnnuallyConfigOk returns a tuple with the AnnuallyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnuallyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) SetAnnuallyConfig(v AnnuallyConfig)`

SetAnnuallyConfig sets AnnuallyConfig field to given value.

### HasAnnuallyConfig

`func (o *AwsNativeStandardBackupScheduleConfig) HasAnnuallyConfig() bool`

HasAnnuallyConfig returns a boolean if a field has been set.

### SetAnnuallyConfigNil

`func (o *AwsNativeStandardBackupScheduleConfig) SetAnnuallyConfigNil(b bool)`

 SetAnnuallyConfigNil sets the value for AnnuallyConfig to be an explicit nil

### UnsetAnnuallyConfig
`func (o *AwsNativeStandardBackupScheduleConfig) UnsetAnnuallyConfig()`

UnsetAnnuallyConfig ensures that no value is present for AnnuallyConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


