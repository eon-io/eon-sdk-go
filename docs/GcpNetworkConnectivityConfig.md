# GcpNetworkConnectivityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Network name. | 
**SubnetsPerRegion** | Pointer to [**[]GcpSubnetPerRegion**](GcpSubnetPerRegion.md) | Subnets to configure for regions in the network. For regions not specified in this list, Eon attempts to use the default subnet.  | [optional] 
**IsSharedVPC** | **bool** | Indicates whether the VPC network is a shared VPC. If true, &#x60;networkHostProject&#x60; must be specified.  | [default to false]
**NetworkHostProject** | **string** | ID of the project that hosts the VPC network. Applicable for shared VPC network.  | 

## Methods

### NewGcpNetworkConnectivityConfig

`func NewGcpNetworkConnectivityConfig(network string, isSharedVPC bool, networkHostProject string, ) *GcpNetworkConnectivityConfig`

NewGcpNetworkConnectivityConfig instantiates a new GcpNetworkConnectivityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpNetworkConnectivityConfigWithDefaults

`func NewGcpNetworkConnectivityConfigWithDefaults() *GcpNetworkConnectivityConfig`

NewGcpNetworkConnectivityConfigWithDefaults instantiates a new GcpNetworkConnectivityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GcpNetworkConnectivityConfig) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GcpNetworkConnectivityConfig) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GcpNetworkConnectivityConfig) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSubnetsPerRegion

`func (o *GcpNetworkConnectivityConfig) GetSubnetsPerRegion() []GcpSubnetPerRegion`

GetSubnetsPerRegion returns the SubnetsPerRegion field if non-nil, zero value otherwise.

### GetSubnetsPerRegionOk

`func (o *GcpNetworkConnectivityConfig) GetSubnetsPerRegionOk() (*[]GcpSubnetPerRegion, bool)`

GetSubnetsPerRegionOk returns a tuple with the SubnetsPerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsPerRegion

`func (o *GcpNetworkConnectivityConfig) SetSubnetsPerRegion(v []GcpSubnetPerRegion)`

SetSubnetsPerRegion sets SubnetsPerRegion field to given value.

### HasSubnetsPerRegion

`func (o *GcpNetworkConnectivityConfig) HasSubnetsPerRegion() bool`

HasSubnetsPerRegion returns a boolean if a field has been set.

### GetIsSharedVPC

`func (o *GcpNetworkConnectivityConfig) GetIsSharedVPC() bool`

GetIsSharedVPC returns the IsSharedVPC field if non-nil, zero value otherwise.

### GetIsSharedVPCOk

`func (o *GcpNetworkConnectivityConfig) GetIsSharedVPCOk() (*bool, bool)`

GetIsSharedVPCOk returns a tuple with the IsSharedVPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedVPC

`func (o *GcpNetworkConnectivityConfig) SetIsSharedVPC(v bool)`

SetIsSharedVPC sets IsSharedVPC field to given value.


### GetNetworkHostProject

`func (o *GcpNetworkConnectivityConfig) GetNetworkHostProject() string`

GetNetworkHostProject returns the NetworkHostProject field if non-nil, zero value otherwise.

### GetNetworkHostProjectOk

`func (o *GcpNetworkConnectivityConfig) GetNetworkHostProjectOk() (*string, bool)`

GetNetworkHostProjectOk returns a tuple with the NetworkHostProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHostProject

`func (o *GcpNetworkConnectivityConfig) SetNetworkHostProject(v string)`

SetNetworkHostProject sets NetworkHostProject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


