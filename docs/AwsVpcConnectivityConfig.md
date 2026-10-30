# AwsVpcConnectivityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | VPC region. | 
**Vpc** | **string** | VPC ID. | 
**SubnetsPerAvailabilityZone** | Pointer to [**[]SubnetPerAvailabilityZone**](SubnetPerAvailabilityZone.md) | Subnets to configure for availability zones in the VPC. For availability zones not specified in this list, Eon attempts to use the default subnet.  | [optional] 
**SecurityGroups** | Pointer to [**ResourceTypeToSecurityGroup**](ResourceTypeToSecurityGroup.md) |  | [optional] 
**PrivateSubnetEnabled** | Pointer to **bool** | Whether to allow cross-region restores into subnets that have no internet access. When true, Eon reaches vault data from this VPC through an S3 gateway endpoint using S3 Multi-Region Access Points. Turn this on if any subnet in the VPC has no internet access.  In the Eon console, this setting is labeled **Allow cross-region restore via S3 gateway endpoint**.  | [optional] [default to false]
**RestoreServerKmsKeyArn** | Pointer to **string** | ARN of the KMS key for the temporary restore server that Eon launches in this VPC. The key encrypts the server&#39;s root volume and the SSM parameter that holds its container registry token. It applies only if the server launches in one of the subnets listed in &#x60;subnetsPerAvailabilityZone&#x60;.  If this is the only VPC in the region, the key also encrypts the SSM parameter that holds the access token for restores from an Azure vault.  The key doesn&#39;t encrypt the restored resources.  The key must be in the same region as the VPC, and its key policy must allow the restore account to use it. If you leave this out, the server uses the key chosen for the restored resource.  | [optional] 

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

### GetPrivateSubnetEnabled

`func (o *AwsVpcConnectivityConfig) GetPrivateSubnetEnabled() bool`

GetPrivateSubnetEnabled returns the PrivateSubnetEnabled field if non-nil, zero value otherwise.

### GetPrivateSubnetEnabledOk

`func (o *AwsVpcConnectivityConfig) GetPrivateSubnetEnabledOk() (*bool, bool)`

GetPrivateSubnetEnabledOk returns a tuple with the PrivateSubnetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnetEnabled

`func (o *AwsVpcConnectivityConfig) SetPrivateSubnetEnabled(v bool)`

SetPrivateSubnetEnabled sets PrivateSubnetEnabled field to given value.

### HasPrivateSubnetEnabled

`func (o *AwsVpcConnectivityConfig) HasPrivateSubnetEnabled() bool`

HasPrivateSubnetEnabled returns a boolean if a field has been set.

### GetRestoreServerKmsKeyArn

`func (o *AwsVpcConnectivityConfig) GetRestoreServerKmsKeyArn() string`

GetRestoreServerKmsKeyArn returns the RestoreServerKmsKeyArn field if non-nil, zero value otherwise.

### GetRestoreServerKmsKeyArnOk

`func (o *AwsVpcConnectivityConfig) GetRestoreServerKmsKeyArnOk() (*string, bool)`

GetRestoreServerKmsKeyArnOk returns a tuple with the RestoreServerKmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerKmsKeyArn

`func (o *AwsVpcConnectivityConfig) SetRestoreServerKmsKeyArn(v string)`

SetRestoreServerKmsKeyArn sets RestoreServerKmsKeyArn field to given value.

### HasRestoreServerKmsKeyArn

`func (o *AwsVpcConnectivityConfig) HasRestoreServerKmsKeyArn() bool`

HasRestoreServerKmsKeyArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


