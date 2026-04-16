# AwsNativePitrBackupPolicyPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionDays** | **int32** | Number of days to retain continuous backups using AWS Backup. AWS allows 1-35 days for RDS/Aurora continuous backups.  | 
**ResourceType** | [**AwsNativePitrBackupResourceType**](AwsNativePitrBackupResourceType.md) |  | 

## Methods

### NewAwsNativePitrBackupPolicyPlan

`func NewAwsNativePitrBackupPolicyPlan(retentionDays int32, resourceType AwsNativePitrBackupResourceType, ) *AwsNativePitrBackupPolicyPlan`

NewAwsNativePitrBackupPolicyPlan instantiates a new AwsNativePitrBackupPolicyPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsNativePitrBackupPolicyPlanWithDefaults

`func NewAwsNativePitrBackupPolicyPlanWithDefaults() *AwsNativePitrBackupPolicyPlan`

NewAwsNativePitrBackupPolicyPlanWithDefaults instantiates a new AwsNativePitrBackupPolicyPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionDays

`func (o *AwsNativePitrBackupPolicyPlan) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *AwsNativePitrBackupPolicyPlan) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *AwsNativePitrBackupPolicyPlan) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetResourceType

`func (o *AwsNativePitrBackupPolicyPlan) GetResourceType() AwsNativePitrBackupResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AwsNativePitrBackupPolicyPlan) GetResourceTypeOk() (*AwsNativePitrBackupResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AwsNativePitrBackupPolicyPlan) SetResourceType(v AwsNativePitrBackupResourceType)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


