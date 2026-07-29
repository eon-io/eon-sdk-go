# AwsNativeStandardBackupSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetRegion** | **string** | Target region for the native backup. | 
**ScheduleConfig** | [**AwsNativeStandardBackupScheduleConfig**](AwsNativeStandardBackupScheduleConfig.md) |  | 
**BackupRetentionDays** | **int32** | Backup retention period, in days. | 

## Methods

### NewAwsNativeStandardBackupSchedules

`func NewAwsNativeStandardBackupSchedules(targetRegion string, scheduleConfig AwsNativeStandardBackupScheduleConfig, backupRetentionDays int32, ) *AwsNativeStandardBackupSchedules`

NewAwsNativeStandardBackupSchedules instantiates a new AwsNativeStandardBackupSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativeStandardBackupSchedulesWithDefaults

`func NewAwsNativeStandardBackupSchedulesWithDefaults() *AwsNativeStandardBackupSchedules`

NewAwsNativeStandardBackupSchedulesWithDefaults instantiates a new AwsNativeStandardBackupSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetRegion

`func (o *AwsNativeStandardBackupSchedules) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *AwsNativeStandardBackupSchedules) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *AwsNativeStandardBackupSchedules) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


### GetScheduleConfig

`func (o *AwsNativeStandardBackupSchedules) GetScheduleConfig() AwsNativeStandardBackupScheduleConfig`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *AwsNativeStandardBackupSchedules) GetScheduleConfigOk() (*AwsNativeStandardBackupScheduleConfig, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *AwsNativeStandardBackupSchedules) SetScheduleConfig(v AwsNativeStandardBackupScheduleConfig)`

SetScheduleConfig sets ScheduleConfig field to given value.


### GetBackupRetentionDays

`func (o *AwsNativeStandardBackupSchedules) GetBackupRetentionDays() int32`

GetBackupRetentionDays returns the BackupRetentionDays field if non-nil, zero value otherwise.

### GetBackupRetentionDaysOk

`func (o *AwsNativeStandardBackupSchedules) GetBackupRetentionDaysOk() (*int32, bool)`

GetBackupRetentionDaysOk returns a tuple with the BackupRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionDays

`func (o *AwsNativeStandardBackupSchedules) SetBackupRetentionDays(v int32)`

SetBackupRetentionDays sets BackupRetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


