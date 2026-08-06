# UpdateObjectStoreScanMethodSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Whether to enable this scan method. Required when systemControlled is false. Ignored when systemControlled is true. | [optional] 
**SystemControlled** | **bool** | Whether Eon manages this scan method. | 

## Methods

### NewUpdateObjectStoreScanMethodSetting

`func NewUpdateObjectStoreScanMethodSetting(systemControlled bool, ) *UpdateObjectStoreScanMethodSetting`

NewUpdateObjectStoreScanMethodSetting instantiates a new UpdateObjectStoreScanMethodSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStoreScanMethodSettingWithDefaults

`func NewUpdateObjectStoreScanMethodSettingWithDefaults() *UpdateObjectStoreScanMethodSetting`

NewUpdateObjectStoreScanMethodSettingWithDefaults instantiates a new UpdateObjectStoreScanMethodSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateObjectStoreScanMethodSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateObjectStoreScanMethodSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateObjectStoreScanMethodSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateObjectStoreScanMethodSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateObjectStoreScanMethodSetting) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateObjectStoreScanMethodSetting) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetSystemControlled

`func (o *UpdateObjectStoreScanMethodSetting) GetSystemControlled() bool`

GetSystemControlled returns the SystemControlled field if non-nil, zero value otherwise.

### GetSystemControlledOk

`func (o *UpdateObjectStoreScanMethodSetting) GetSystemControlledOk() (*bool, bool)`

GetSystemControlledOk returns a tuple with the SystemControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemControlled

`func (o *UpdateObjectStoreScanMethodSetting) SetSystemControlled(v bool)`

SetSystemControlled sets SystemControlled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


