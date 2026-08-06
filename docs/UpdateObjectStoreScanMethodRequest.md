# UpdateObjectStoreScanMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdcBackup** | Pointer to [**UpdateObjectStoreScanMethodSetting**](UpdateObjectStoreScanMethodSetting.md) |  | [optional] 
**InventoryBackup** | Pointer to [**UpdateObjectStoreScanMethodSetting**](UpdateObjectStoreScanMethodSetting.md) |  | [optional] 

## Methods

### NewUpdateObjectStoreScanMethodRequest

`func NewUpdateObjectStoreScanMethodRequest() *UpdateObjectStoreScanMethodRequest`

NewUpdateObjectStoreScanMethodRequest instantiates a new UpdateObjectStoreScanMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStoreScanMethodRequestWithDefaults

`func NewUpdateObjectStoreScanMethodRequestWithDefaults() *UpdateObjectStoreScanMethodRequest`

NewUpdateObjectStoreScanMethodRequestWithDefaults instantiates a new UpdateObjectStoreScanMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdcBackup

`func (o *UpdateObjectStoreScanMethodRequest) GetCdcBackup() UpdateObjectStoreScanMethodSetting`

GetCdcBackup returns the CdcBackup field if non-nil, zero value otherwise.

### GetCdcBackupOk

`func (o *UpdateObjectStoreScanMethodRequest) GetCdcBackupOk() (*UpdateObjectStoreScanMethodSetting, bool)`

GetCdcBackupOk returns a tuple with the CdcBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcBackup

`func (o *UpdateObjectStoreScanMethodRequest) SetCdcBackup(v UpdateObjectStoreScanMethodSetting)`

SetCdcBackup sets CdcBackup field to given value.

### HasCdcBackup

`func (o *UpdateObjectStoreScanMethodRequest) HasCdcBackup() bool`

HasCdcBackup returns a boolean if a field has been set.

### GetInventoryBackup

`func (o *UpdateObjectStoreScanMethodRequest) GetInventoryBackup() UpdateObjectStoreScanMethodSetting`

GetInventoryBackup returns the InventoryBackup field if non-nil, zero value otherwise.

### GetInventoryBackupOk

`func (o *UpdateObjectStoreScanMethodRequest) GetInventoryBackupOk() (*UpdateObjectStoreScanMethodSetting, bool)`

GetInventoryBackupOk returns a tuple with the InventoryBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryBackup

`func (o *UpdateObjectStoreScanMethodRequest) SetInventoryBackup(v UpdateObjectStoreScanMethodSetting)`

SetInventoryBackup sets InventoryBackup field to given value.

### HasInventoryBackup

`func (o *UpdateObjectStoreScanMethodRequest) HasInventoryBackup() bool`

HasInventoryBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


