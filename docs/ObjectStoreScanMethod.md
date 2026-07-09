# ObjectStoreScanMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdcBackup** | [**ObjectStoreScanMethodSetting**](ObjectStoreScanMethodSetting.md) |  | 
**InventoryBackup** | [**ObjectStoreScanMethodSetting**](ObjectStoreScanMethodSetting.md) |  | 

## Methods

### NewObjectStoreScanMethod

`func NewObjectStoreScanMethod(cdcBackup ObjectStoreScanMethodSetting, inventoryBackup ObjectStoreScanMethodSetting, ) *ObjectStoreScanMethod`

NewObjectStoreScanMethod instantiates a new ObjectStoreScanMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoreScanMethodWithDefaults

`func NewObjectStoreScanMethodWithDefaults() *ObjectStoreScanMethod`

NewObjectStoreScanMethodWithDefaults instantiates a new ObjectStoreScanMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdcBackup

`func (o *ObjectStoreScanMethod) GetCdcBackup() ObjectStoreScanMethodSetting`

GetCdcBackup returns the CdcBackup field if non-nil, zero value otherwise.

### GetCdcBackupOk

`func (o *ObjectStoreScanMethod) GetCdcBackupOk() (*ObjectStoreScanMethodSetting, bool)`

GetCdcBackupOk returns a tuple with the CdcBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcBackup

`func (o *ObjectStoreScanMethod) SetCdcBackup(v ObjectStoreScanMethodSetting)`

SetCdcBackup sets CdcBackup field to given value.


### GetInventoryBackup

`func (o *ObjectStoreScanMethod) GetInventoryBackup() ObjectStoreScanMethodSetting`

GetInventoryBackup returns the InventoryBackup field if non-nil, zero value otherwise.

### GetInventoryBackupOk

`func (o *ObjectStoreScanMethod) GetInventoryBackupOk() (*ObjectStoreScanMethodSetting, bool)`

GetInventoryBackupOk returns a tuple with the InventoryBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryBackup

`func (o *ObjectStoreScanMethod) SetInventoryBackup(v ObjectStoreScanMethodSetting)`

SetInventoryBackup sets InventoryBackup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


