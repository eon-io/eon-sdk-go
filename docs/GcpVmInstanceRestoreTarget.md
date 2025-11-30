# GcpVmInstanceRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | Zone to restore to. | 
**MachineType** | **string** | Machine type to use. | 
**Name** | **string** | VM name. | 
**NetworkName** | **string** | Name of the VPC network to use. | 
**SubnetName** | **string** | Name of the subnet to use. | 
**NetworkHostProject** | Pointer to **string** | ID of the project that hosts the VPC network. Applicable only when restoring to a shared VPC network.  | [optional] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the restored VM as key-value pairs, where key and value are both strings. These labels are always applied: &#x60;\&quot;eon-restore\&quot;: \&quot;true\&quot;&#x60;.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**Disks** | [**[]RestoreGcpInstanceDiskInput**](RestoreGcpInstanceDiskInput.md) | Disks to restore and attach to the restored instance. Each item in the list corresponds to a disk to be restored, where &#x60;providerDiskId&#x60; matches the disk&#39;s ID at the time of the snapshot. The boot disk must be in the list.  | 
**StartInstanceAfterRestore** | Pointer to **bool** | Whether to start the VM instance after restoring it. If set to &#x60;false&#x60;, the VM will be created in a stopped state.  | [optional] [default to true]

## Methods

### NewGcpVmInstanceRestoreTarget

`func NewGcpVmInstanceRestoreTarget(zone string, machineType string, name string, networkName string, subnetName string, disks []RestoreGcpInstanceDiskInput, ) *GcpVmInstanceRestoreTarget`

NewGcpVmInstanceRestoreTarget instantiates a new GcpVmInstanceRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpVmInstanceRestoreTargetWithDefaults

`func NewGcpVmInstanceRestoreTargetWithDefaults() *GcpVmInstanceRestoreTarget`

NewGcpVmInstanceRestoreTargetWithDefaults instantiates a new GcpVmInstanceRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *GcpVmInstanceRestoreTarget) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GcpVmInstanceRestoreTarget) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GcpVmInstanceRestoreTarget) SetZone(v string)`

SetZone sets Zone field to given value.


### GetMachineType

`func (o *GcpVmInstanceRestoreTarget) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *GcpVmInstanceRestoreTarget) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *GcpVmInstanceRestoreTarget) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.


### GetName

`func (o *GcpVmInstanceRestoreTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpVmInstanceRestoreTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpVmInstanceRestoreTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkName

`func (o *GcpVmInstanceRestoreTarget) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GcpVmInstanceRestoreTarget) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GcpVmInstanceRestoreTarget) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetSubnetName

`func (o *GcpVmInstanceRestoreTarget) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *GcpVmInstanceRestoreTarget) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *GcpVmInstanceRestoreTarget) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetNetworkHostProject

`func (o *GcpVmInstanceRestoreTarget) GetNetworkHostProject() string`

GetNetworkHostProject returns the NetworkHostProject field if non-nil, zero value otherwise.

### GetNetworkHostProjectOk

`func (o *GcpVmInstanceRestoreTarget) GetNetworkHostProjectOk() (*string, bool)`

GetNetworkHostProjectOk returns a tuple with the NetworkHostProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHostProject

`func (o *GcpVmInstanceRestoreTarget) SetNetworkHostProject(v string)`

SetNetworkHostProject sets NetworkHostProject field to given value.

### HasNetworkHostProject

`func (o *GcpVmInstanceRestoreTarget) HasNetworkHostProject() bool`

HasNetworkHostProject returns a boolean if a field has been set.

### GetLabels

`func (o *GcpVmInstanceRestoreTarget) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpVmInstanceRestoreTarget) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpVmInstanceRestoreTarget) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpVmInstanceRestoreTarget) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDisks

`func (o *GcpVmInstanceRestoreTarget) GetDisks() []RestoreGcpInstanceDiskInput`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *GcpVmInstanceRestoreTarget) GetDisksOk() (*[]RestoreGcpInstanceDiskInput, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *GcpVmInstanceRestoreTarget) SetDisks(v []RestoreGcpInstanceDiskInput)`

SetDisks sets Disks field to given value.


### GetStartInstanceAfterRestore

`func (o *GcpVmInstanceRestoreTarget) GetStartInstanceAfterRestore() bool`

GetStartInstanceAfterRestore returns the StartInstanceAfterRestore field if non-nil, zero value otherwise.

### GetStartInstanceAfterRestoreOk

`func (o *GcpVmInstanceRestoreTarget) GetStartInstanceAfterRestoreOk() (*bool, bool)`

GetStartInstanceAfterRestoreOk returns a tuple with the StartInstanceAfterRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInstanceAfterRestore

`func (o *GcpVmInstanceRestoreTarget) SetStartInstanceAfterRestore(v bool)`

SetStartInstanceAfterRestore sets StartInstanceAfterRestore field to given value.

### HasStartInstanceAfterRestore

`func (o *GcpVmInstanceRestoreTarget) HasStartInstanceAfterRestore() bool`

HasStartInstanceAfterRestore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


