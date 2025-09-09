# RestoreAzureInstanceDiskInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore | 
**Settings** | [**AzureDiskSettings**](AzureDiskSettings.md) |  | 

## Methods

### NewRestoreAzureInstanceDiskInput

`func NewRestoreAzureInstanceDiskInput(providerDiskId string, settings AzureDiskSettings, ) *RestoreAzureInstanceDiskInput`

NewRestoreAzureInstanceDiskInput instantiates a new RestoreAzureInstanceDiskInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAzureInstanceDiskInputWithDefaults

`func NewRestoreAzureInstanceDiskInputWithDefaults() *RestoreAzureInstanceDiskInput`

NewRestoreAzureInstanceDiskInputWithDefaults instantiates a new RestoreAzureInstanceDiskInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderDiskId

`func (o *RestoreAzureInstanceDiskInput) GetProviderDiskId() string`

GetProviderDiskId returns the ProviderDiskId field if non-nil, zero value otherwise.

### GetProviderDiskIdOk

`func (o *RestoreAzureInstanceDiskInput) GetProviderDiskIdOk() (*string, bool)`

GetProviderDiskIdOk returns a tuple with the ProviderDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDiskId

`func (o *RestoreAzureInstanceDiskInput) SetProviderDiskId(v string)`

SetProviderDiskId sets ProviderDiskId field to given value.


### GetSettings

`func (o *RestoreAzureInstanceDiskInput) GetSettings() AzureDiskSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *RestoreAzureInstanceDiskInput) GetSettingsOk() (*AzureDiskSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *RestoreAzureInstanceDiskInput) SetSettings(v AzureDiskSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


