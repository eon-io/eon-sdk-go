# AwsVpcConnectivityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | VPC region. | 
**Vpc** | **string** | VPC ID. | 
**SubnetsPerAvailabilityZone** | Pointer to [**[]SubnetPerAvailabilityZone**](SubnetPerAvailabilityZone.md) | Subnets to configure for availability zones in the VPC. For availability zones not specified in this list, Eon attempts to use the default subnet.  | [optional] 
**SecurityGroups** | Pointer to [**ResourceTypeToSecurityGroup**](ResourceTypeToSecurityGroup.md) |  | [optional] 

## Methods

### NewAwsVpcConnectivityConfig

`func NewAwsVpcConnectivityConfig(region string, vpc string, ) *AwsVpcConnectivityConfig`

NewAwsVpcConnectivityConfig instantiates a new AwsVpcConnectivityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcConnectivityConfigWithDefaults

`func NewAwsVpcConnectivityConfigWithDefaults() *AwsVpcConnectivityConfig`

NewAwsVpcConnectivityConfigWithDefaults instantiates a new AwsVpcConnectivityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsVpcConnectivityConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsVpcConnectivityConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsVpcConnectivityConfig) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVpc

`func (o *AwsVpcConnectivityConfig) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AwsVpcConnectivityConfig) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AwsVpcConnectivityConfig) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetSubnetsPerAvailabilityZone

`func (o *AwsVpcConnectivityConfig) GetSubnetsPerAvailabilityZone() []SubnetPerAvailabilityZone`

GetSubnetsPerAvailabilityZone returns the SubnetsPerAvailabilityZone field if non-nil, zero value otherwise.

### GetSubnetsPerAvailabilityZoneOk

`func (o *AwsVpcConnectivityConfig) GetSubnetsPerAvailabilityZoneOk() (*[]SubnetPerAvailabilityZone, bool)`

GetSubnetsPerAvailabilityZoneOk returns a tuple with the SubnetsPerAvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsPerAvailabilityZone

`func (o *AwsVpcConnectivityConfig) SetSubnetsPerAvailabilityZone(v []SubnetPerAvailabilityZone)`

SetSubnetsPerAvailabilityZone sets SubnetsPerAvailabilityZone field to given value.

### HasSubnetsPerAvailabilityZone

`func (o *AwsVpcConnectivityConfig) HasSubnetsPerAvailabilityZone() bool`

HasSubnetsPerAvailabilityZone returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AwsVpcConnectivityConfig) GetSecurityGroups() ResourceTypeToSecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AwsVpcConnectivityConfig) GetSecurityGroupsOk() (*ResourceTypeToSecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AwsVpcConnectivityConfig) SetSecurityGroups(v ResourceTypeToSecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AwsVpcConnectivityConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


