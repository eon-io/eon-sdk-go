# AzureDiskRestoreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | **string** | The ID of the restored Azure disk. | 

## Methods

### NewAzureDiskRestoreResult

`func NewAzureDiskRestoreResult(diskId string, ) *AzureDiskRestoreResult`

NewAzureDiskRestoreResult instantiates a new AzureDiskRestoreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDiskRestoreResultWithDefaults

`func NewAzureDiskRestoreResultWithDefaults() *AzureDiskRestoreResult`

NewAzureDiskRestoreResultWithDefaults instantiates a new AzureDiskRestoreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskId

`func (o *AzureDiskRestoreResult) GetDiskId() string`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *AzureDiskRestoreResult) GetDiskIdOk() (*string, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *AzureDiskRestoreResult) SetDiskId(v string)`

SetDiskId sets DiskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


