# RestoreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsEc2Instance** | Pointer to [**NullableAwsEc2InstanceRestoreResult**](AwsEc2InstanceRestoreResult.md) |  | [optional] 
**AwsEbsVolume** | Pointer to [**NullableAwsEbsVolumeRestoreResult**](AwsEbsVolumeRestoreResult.md) |  | [optional] 
**AzureDisk** | Pointer to [**NullableAzureDiskRestoreResult**](AzureDiskRestoreResult.md) |  | [optional] 

## Methods

### NewRestoreResult

`func NewRestoreResult() *RestoreResult`

NewRestoreResult instantiates a new RestoreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreResultWithDefaults

`func NewRestoreResultWithDefaults() *RestoreResult`

NewRestoreResultWithDefaults instantiates a new RestoreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsEc2Instance

`func (o *RestoreResult) GetAwsEc2Instance() AwsEc2InstanceRestoreResult`

GetAwsEc2Instance returns the AwsEc2Instance field if non-nil, zero value otherwise.

### GetAwsEc2InstanceOk

`func (o *RestoreResult) GetAwsEc2InstanceOk() (*AwsEc2InstanceRestoreResult, bool)`

GetAwsEc2InstanceOk returns a tuple with the AwsEc2Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEc2Instance

`func (o *RestoreResult) SetAwsEc2Instance(v AwsEc2InstanceRestoreResult)`

SetAwsEc2Instance sets AwsEc2Instance field to given value.

### HasAwsEc2Instance

`func (o *RestoreResult) HasAwsEc2Instance() bool`

HasAwsEc2Instance returns a boolean if a field has been set.

### SetAwsEc2InstanceNil

`func (o *RestoreResult) SetAwsEc2InstanceNil(b bool)`

 SetAwsEc2InstanceNil sets the value for AwsEc2Instance to be an explicit nil

### UnsetAwsEc2Instance
`func (o *RestoreResult) UnsetAwsEc2Instance()`

UnsetAwsEc2Instance ensures that no value is present for AwsEc2Instance, not even an explicit nil
### GetAwsEbsVolume

`func (o *RestoreResult) GetAwsEbsVolume() AwsEbsVolumeRestoreResult`

GetAwsEbsVolume returns the AwsEbsVolume field if non-nil, zero value otherwise.

### GetAwsEbsVolumeOk

`func (o *RestoreResult) GetAwsEbsVolumeOk() (*AwsEbsVolumeRestoreResult, bool)`

GetAwsEbsVolumeOk returns a tuple with the AwsEbsVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEbsVolume

`func (o *RestoreResult) SetAwsEbsVolume(v AwsEbsVolumeRestoreResult)`

SetAwsEbsVolume sets AwsEbsVolume field to given value.

### HasAwsEbsVolume

`func (o *RestoreResult) HasAwsEbsVolume() bool`

HasAwsEbsVolume returns a boolean if a field has been set.

### SetAwsEbsVolumeNil

`func (o *RestoreResult) SetAwsEbsVolumeNil(b bool)`

 SetAwsEbsVolumeNil sets the value for AwsEbsVolume to be an explicit nil

### UnsetAwsEbsVolume
`func (o *RestoreResult) UnsetAwsEbsVolume()`

UnsetAwsEbsVolume ensures that no value is present for AwsEbsVolume, not even an explicit nil
### GetAzureDisk

`func (o *RestoreResult) GetAzureDisk() AzureDiskRestoreResult`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *RestoreResult) GetAzureDiskOk() (*AzureDiskRestoreResult, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *RestoreResult) SetAzureDisk(v AzureDiskRestoreResult)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *RestoreResult) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### SetAzureDiskNil

`func (o *RestoreResult) SetAzureDiskNil(b bool)`

 SetAzureDiskNil sets the value for AzureDisk to be an explicit nil

### UnsetAzureDisk
`func (o *RestoreResult) UnsetAzureDisk()`

UnsetAzureDisk ensures that no value is present for AzureDisk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


