# AzureVmInstanceRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to restore to. | 
**ResourceGroupName** | **string** | Name of the resource group to restore to. | 
**VmName** | **string** | Restored VM name. | 
**VmSize** | **string** | Size of the restored VM. | 
**NetworkInterface** | **string** | Name of the network interface to use. | 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings. If not provided, defaults to an empty object, with no tags applied.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**Disks** | [**[]RestoreAzureInstanceDiskInput**](RestoreAzureInstanceDiskInput.md) |  | 
**StartInstanceAfterRestore** | Pointer to **bool** | Whether to start the VM instance after restoring it. If set to &#x60;false&#x60;, the VM will be created in a stopped state.  | [optional] [default to true]

## Methods

### NewAzureVmInstanceRestoreTarget

`func NewAzureVmInstanceRestoreTarget(region string, resourceGroupName string, vmName string, vmSize string, networkInterface string, disks []RestoreAzureInstanceDiskInput, ) *AzureVmInstanceRestoreTarget`

NewAzureVmInstanceRestoreTarget instantiates a new AzureVmInstanceRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVmInstanceRestoreTargetWithDefaults

`func NewAzureVmInstanceRestoreTargetWithDefaults() *AzureVmInstanceRestoreTarget`

NewAzureVmInstanceRestoreTargetWithDefaults instantiates a new AzureVmInstanceRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AzureVmInstanceRestoreTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureVmInstanceRestoreTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureVmInstanceRestoreTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


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

### GetDisks

`func (o *AzureVmInstanceRestoreTarget) GetDisks() []RestoreAzureInstanceDiskInput`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *AzureVmInstanceRestoreTarget) GetDisksOk() (*[]RestoreAzureInstanceDiskInput, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *AzureVmInstanceRestoreTarget) SetDisks(v []RestoreAzureInstanceDiskInput)`

SetDisks sets Disks field to given value.


### GetStartInstanceAfterRestore

`func (o *AzureVmInstanceRestoreTarget) GetStartInstanceAfterRestore() bool`

GetStartInstanceAfterRestore returns the StartInstanceAfterRestore field if non-nil, zero value otherwise.

### GetStartInstanceAfterRestoreOk

`func (o *AzureVmInstanceRestoreTarget) GetStartInstanceAfterRestoreOk() (*bool, bool)`

GetStartInstanceAfterRestoreOk returns a tuple with the StartInstanceAfterRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInstanceAfterRestore

`func (o *AzureVmInstanceRestoreTarget) SetStartInstanceAfterRestore(v bool)`

SetStartInstanceAfterRestore sets StartInstanceAfterRestore field to given value.

### HasStartInstanceAfterRestore

`func (o *AzureVmInstanceRestoreTarget) HasStartInstanceAfterRestore() bool`

HasStartInstanceAfterRestore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


