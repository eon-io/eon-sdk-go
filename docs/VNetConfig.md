# VNetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Whether to use an existing virtual network or create a new one. | 
**Name** | Pointer to **string** | Virtual network name. Required when &#x60;mode&#x60; is set to use an existing virtual network.  Optional when &#x60;mode&#x60; is set to create a new virtual network. New virtual networks default to &#x60;eon-default-vnet-&lt;region&gt;&#x60;.  The virtual network must be in the same subscription and region as the VM.  | [optional] 
**ResourceGroupName** | Pointer to **string** | Resource group that contains the new or existing virtual network.  Defaults to the restored VM&#39;s resource group.  | [optional] 
**AddressSpace** | Pointer to **string** | Address space in CIDR notation. Used only when &#x60;mode&#x60; is set to create a new virtual network.  Defaults to &#x60;10.0.0.0/16&#x60;.  | [optional] 

## Methods

### NewVNetConfig

`func NewVNetConfig(mode string, ) *VNetConfig`

NewVNetConfig instantiates a new VNetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetConfigWithDefaults

`func NewVNetConfigWithDefaults() *VNetConfig`

NewVNetConfigWithDefaults instantiates a new VNetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *VNetConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VNetConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VNetConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetName

`func (o *VNetConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VNetConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VNetConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VNetConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceGroupName

`func (o *VNetConfig) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *VNetConfig) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *VNetConfig) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *VNetConfig) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### GetAddressSpace

`func (o *VNetConfig) GetAddressSpace() string`

GetAddressSpace returns the AddressSpace field if non-nil, zero value otherwise.

### GetAddressSpaceOk

`func (o *VNetConfig) GetAddressSpaceOk() (*string, bool)`

GetAddressSpaceOk returns a tuple with the AddressSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSpace

`func (o *VNetConfig) SetAddressSpace(v string)`

SetAddressSpace sets AddressSpace field to given value.

### HasAddressSpace

`func (o *VNetConfig) HasAddressSpace() bool`

HasAddressSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


