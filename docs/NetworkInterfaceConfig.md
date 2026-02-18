# NetworkInterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Whether to use an existing NIC or create a new one | 
**NicId** | Pointer to **string** | Full Azure resource ID of the existing NIC. Required when mode&#x3D;existing. Format: /subscriptions/{sub}/resourceGroups/{rg}/providers/Microsoft.Network/networkInterfaces/{name}  | [optional] 
**NicName** | Pointer to **string** | Name for the new NIC. Only used when mode&#x3D;create. Defaults to &lt;vmName&gt;-nic if not specified.  | [optional] 
**Vnet** | Pointer to [**VNetConfig**](VNetConfig.md) |  | [optional] 
**Subnet** | Pointer to [**SubnetConfig**](SubnetConfig.md) |  | [optional] 

## Methods

### NewNetworkInterfaceConfig

`func NewNetworkInterfaceConfig(mode string, ) *NetworkInterfaceConfig`

NewNetworkInterfaceConfig instantiates a new NetworkInterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceConfigWithDefaults

`func NewNetworkInterfaceConfigWithDefaults() *NetworkInterfaceConfig`

NewNetworkInterfaceConfigWithDefaults instantiates a new NetworkInterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *NetworkInterfaceConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NetworkInterfaceConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NetworkInterfaceConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNicId

`func (o *NetworkInterfaceConfig) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *NetworkInterfaceConfig) GetNicIdOk() (*string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicId

`func (o *NetworkInterfaceConfig) SetNicId(v string)`

SetNicId sets NicId field to given value.

### HasNicId

`func (o *NetworkInterfaceConfig) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### GetNicName

`func (o *NetworkInterfaceConfig) GetNicName() string`

GetNicName returns the NicName field if non-nil, zero value otherwise.

### GetNicNameOk

`func (o *NetworkInterfaceConfig) GetNicNameOk() (*string, bool)`

GetNicNameOk returns a tuple with the NicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicName

`func (o *NetworkInterfaceConfig) SetNicName(v string)`

SetNicName sets NicName field to given value.

### HasNicName

`func (o *NetworkInterfaceConfig) HasNicName() bool`

HasNicName returns a boolean if a field has been set.

### GetVnet

`func (o *NetworkInterfaceConfig) GetVnet() VNetConfig`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NetworkInterfaceConfig) GetVnetOk() (*VNetConfig, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NetworkInterfaceConfig) SetVnet(v VNetConfig)`

SetVnet sets Vnet field to given value.

### HasVnet

`func (o *NetworkInterfaceConfig) HasVnet() bool`

HasVnet returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkInterfaceConfig) GetSubnet() SubnetConfig`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkInterfaceConfig) GetSubnetOk() (*SubnetConfig, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkInterfaceConfig) SetSubnet(v SubnetConfig)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkInterfaceConfig) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


