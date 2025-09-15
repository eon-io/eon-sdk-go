# AzureDiskTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to restore to. | 
**ResourceGroupName** | **string** | Name of the resource group to restore to. | 
**Settings** | [**AzureDiskSettings**](AzureDiskSettings.md) |  | 

## Methods

### NewAzureDiskTarget

`func NewAzureDiskTarget(region string, resourceGroupName string, settings AzureDiskSettings, ) *AzureDiskTarget`

NewAzureDiskTarget instantiates a new AzureDiskTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDiskTargetWithDefaults

`func NewAzureDiskTargetWithDefaults() *AzureDiskTarget`

NewAzureDiskTargetWithDefaults instantiates a new AzureDiskTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AzureDiskTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureDiskTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureDiskTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceGroupName

`func (o *AzureDiskTarget) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDiskTarget) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDiskTarget) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetSettings

`func (o *AzureDiskTarget) GetSettings() AzureDiskSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AzureDiskTarget) GetSettingsOk() (*AzureDiskSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AzureDiskTarget) SetSettings(v AzureDiskSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


