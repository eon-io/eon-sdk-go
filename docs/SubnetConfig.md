# SubnetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Whether to use an existing subnet or create a new one. | 
**Name** | Pointer to **string** | Subnet name. Required when &#x60;mode&#x60; is set to use an existing subnet.  Optional when &#x60;mode&#x60; is set to create a new subnet. New subnets default to &#x60;eon-default-subnet-&lt;region&gt;&#x60;.  | [optional] 
**AddressPrefix** | Pointer to **string** | Address prefix in CIDR notation. Used only when &#x60;mode&#x60; is set to create a new subnet.  Defaults to &#x60;10.0.0.0/24&#x60;.  | [optional] 

## Methods

### NewSubnetConfig

`func NewSubnetConfig(mode string, ) *SubnetConfig`

NewSubnetConfig instantiates a new SubnetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetConfigWithDefaults

`func NewSubnetConfigWithDefaults() *SubnetConfig`

NewSubnetConfigWithDefaults instantiates a new SubnetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *SubnetConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SubnetConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SubnetConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetName

`func (o *SubnetConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubnetConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubnetConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubnetConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressPrefix

`func (o *SubnetConfig) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *SubnetConfig) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *SubnetConfig) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.

### HasAddressPrefix

`func (o *SubnetConfig) HasAddressPrefix() bool`

HasAddressPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


