# ObjectStoreScanMethodSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Whether the scan method is enabled. Required only when systemControlled is false. | [optional] 
**SystemControlled** | **bool** | Whether Eon should decide when to use this scan method. | 

## Methods

### NewObjectStoreScanMethodSetting

`func NewObjectStoreScanMethodSetting(systemControlled bool, ) *ObjectStoreScanMethodSetting`

NewObjectStoreScanMethodSetting instantiates a new ObjectStoreScanMethodSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoreScanMethodSettingWithDefaults

`func NewObjectStoreScanMethodSettingWithDefaults() *ObjectStoreScanMethodSetting`

NewObjectStoreScanMethodSettingWithDefaults instantiates a new ObjectStoreScanMethodSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ObjectStoreScanMethodSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObjectStoreScanMethodSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObjectStoreScanMethodSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ObjectStoreScanMethodSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ObjectStoreScanMethodSetting) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ObjectStoreScanMethodSetting) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetSystemControlled

`func (o *ObjectStoreScanMethodSetting) GetSystemControlled() bool`

GetSystemControlled returns the SystemControlled field if non-nil, zero value otherwise.

### GetSystemControlledOk

`func (o *ObjectStoreScanMethodSetting) GetSystemControlledOk() (*bool, bool)`

GetSystemControlledOk returns a tuple with the SystemControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemControlled

`func (o *ObjectStoreScanMethodSetting) SetSystemControlled(v bool)`

SetSystemControlled sets SystemControlled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


