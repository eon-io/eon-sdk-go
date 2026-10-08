# RestoreAzureInstanceDiskInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore. | 
**Settings** | [**AzureDiskSettings**](AzureDiskSettings.md) |  | 
**DiskEncryptionSetId** | Pointer to **string** | ARM resource ID of the disk encryption set to encrypt the restored disk with a customer-managed key. The disk encryption set must be in the restore account&#39;s subscription and in the target region. If not provided, the restored disk is encrypted with a platform-managed key.  | [optional] 

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


### GetDiskEncryptionSetId

`func (o *RestoreAzureInstanceDiskInput) GetDiskEncryptionSetId() string`

GetDiskEncryptionSetId returns the DiskEncryptionSetId field if non-nil, zero value otherwise.

### GetDiskEncryptionSetIdOk

`func (o *RestoreAzureInstanceDiskInput) GetDiskEncryptionSetIdOk() (*string, bool)`

GetDiskEncryptionSetIdOk returns a tuple with the DiskEncryptionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionSetId

`func (o *RestoreAzureInstanceDiskInput) SetDiskEncryptionSetId(v string)`

SetDiskEncryptionSetId sets DiskEncryptionSetId field to given value.

### HasDiskEncryptionSetId

`func (o *RestoreAzureInstanceDiskInput) HasDiskEncryptionSetId() bool`

HasDiskEncryptionSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


