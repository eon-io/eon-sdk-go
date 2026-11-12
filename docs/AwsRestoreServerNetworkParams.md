# AwsRestoreServerNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | **string** | VPC ID used by the restore server. | 
**SubnetId** | **string** | Subnet ID used by the restore server. | 
**SecurityGroupIds** | Pointer to **[]string** | Security group IDs used by the restore server. | [optional] 

## Methods

### NewAwsRestoreServerNetworkParams

`func NewAwsRestoreServerNetworkParams(vpcId string, subnetId string, ) *AwsRestoreServerNetworkParams`

NewAwsRestoreServerNetworkParams instantiates a new AwsRestoreServerNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRestoreServerNetworkParamsWithDefaults

`func NewAwsRestoreServerNetworkParamsWithDefaults() *AwsRestoreServerNetworkParams`

NewAwsRestoreServerNetworkParamsWithDefaults instantiates a new AwsRestoreServerNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *AwsRestoreServerNetworkParams) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsRestoreServerNetworkParams) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsRestoreServerNetworkParams) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetId

`func (o *AwsRestoreServerNetworkParams) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsRestoreServerNetworkParams) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsRestoreServerNetworkParams) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSecurityGroupIds

`func (o *AwsRestoreServerNetworkParams) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AwsRestoreServerNetworkParams) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AwsRestoreServerNetworkParams) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AwsRestoreServerNetworkParams) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


