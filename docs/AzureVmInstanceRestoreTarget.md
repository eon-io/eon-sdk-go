# AzureVmInstanceRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location to restore the instance to. | 
**ResourceGroupName** | **string** | Name of the resource group to restore the vm to | 
**VmName** | **string** | The name of the VM resource to restore | 
**VmSize** | **string** | The size of the VM to restore | 
**NetworkInterface** | **string** | The Name of the network interface to use | 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**DiskParameters** | [**[]RestoreAzureInstanceDiskInput**](RestoreAzureInstanceDiskInput.md) |  | 

## Methods

### NewAzureVmInstanceRestoreTarget

`func NewAzureVmInstanceRestoreTarget(location string, resourceGroupName string, vmName string, vmSize string, networkInterface string, diskParameters []RestoreAzureInstanceDiskInput, ) *AzureVmInstanceRestoreTarget`

NewAzureVmInstanceRestoreTarget instantiates a new AzureVmInstanceRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVmInstanceRestoreTargetWithDefaults

`func NewAzureVmInstanceRestoreTargetWithDefaults() *AzureVmInstanceRestoreTarget`

NewAzureVmInstanceRestoreTargetWithDefaults instantiates a new AzureVmInstanceRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *AzureVmInstanceRestoreTarget) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureVmInstanceRestoreTarget) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureVmInstanceRestoreTarget) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetResourceGroupName

`func (o *AzureVmInstanceRestoreTarget) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureVmInstanceRestoreTarget) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureVmInstanceRestoreTarget) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetVmName

`func (o *AzureVmInstanceRestoreTarget) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *AzureVmInstanceRestoreTarget) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *AzureVmInstanceRestoreTarget) SetVmName(v string)`

SetVmName sets VmName field to given value.


### GetVmSize

`func (o *AzureVmInstanceRestoreTarget) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *AzureVmInstanceRestoreTarget) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *AzureVmInstanceRestoreTarget) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.


### GetNetworkInterface

`func (o *AzureVmInstanceRestoreTarget) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *AzureVmInstanceRestoreTarget) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *AzureVmInstanceRestoreTarget) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.


### GetTags

`func (o *AzureVmInstanceRestoreTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureVmInstanceRestoreTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureVmInstanceRestoreTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureVmInstanceRestoreTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDiskParameters

`func (o *AzureVmInstanceRestoreTarget) GetDiskParameters() []RestoreAzureInstanceDiskInput`

GetDiskParameters returns the DiskParameters field if non-nil, zero value otherwise.

### GetDiskParametersOk

`func (o *AzureVmInstanceRestoreTarget) GetDiskParametersOk() (*[]RestoreAzureInstanceDiskInput, bool)`

GetDiskParametersOk returns a tuple with the DiskParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskParameters

`func (o *AzureVmInstanceRestoreTarget) SetDiskParameters(v []RestoreAzureInstanceDiskInput)`

SetDiskParameters sets DiskParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


