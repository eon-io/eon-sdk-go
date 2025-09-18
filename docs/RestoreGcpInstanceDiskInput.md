# RestoreGcpInstanceDiskInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore. | 
**Settings** | [**GcpDiskSettings**](GcpDiskSettings.md) |  | 

## Methods

### NewRestoreGcpInstanceDiskInput

`func NewRestoreGcpInstanceDiskInput(providerDiskId string, settings GcpDiskSettings, ) *RestoreGcpInstanceDiskInput`

NewRestoreGcpInstanceDiskInput instantiates a new RestoreGcpInstanceDiskInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreGcpInstanceDiskInputWithDefaults

`func NewRestoreGcpInstanceDiskInputWithDefaults() *RestoreGcpInstanceDiskInput`

NewRestoreGcpInstanceDiskInputWithDefaults instantiates a new RestoreGcpInstanceDiskInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderDiskId

`func (o *RestoreGcpInstanceDiskInput) GetProviderDiskId() string`

GetProviderDiskId returns the ProviderDiskId field if non-nil, zero value otherwise.

### GetProviderDiskIdOk

`func (o *RestoreGcpInstanceDiskInput) GetProviderDiskIdOk() (*string, bool)`

GetProviderDiskIdOk returns a tuple with the ProviderDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDiskId

`func (o *RestoreGcpInstanceDiskInput) SetProviderDiskId(v string)`

SetProviderDiskId sets ProviderDiskId field to given value.


### GetSettings

`func (o *RestoreGcpInstanceDiskInput) GetSettings() GcpDiskSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *RestoreGcpInstanceDiskInput) GetSettingsOk() (*GcpDiskSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *RestoreGcpInstanceDiskInput) SetSettings(v GcpDiskSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


