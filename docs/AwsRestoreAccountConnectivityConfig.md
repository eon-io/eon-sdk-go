# AwsRestoreAccountConnectivityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcConfigs** | Pointer to [**[]AwsVpcConnectivityConfig**](AwsVpcConnectivityConfig.md) | VPCs to configure. | [optional] 

## Methods

### NewAwsRestoreAccountConnectivityConfig

`func NewAwsRestoreAccountConnectivityConfig() *AwsRestoreAccountConnectivityConfig`

NewAwsRestoreAccountConnectivityConfig instantiates a new AwsRestoreAccountConnectivityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRestoreAccountConnectivityConfigWithDefaults

`func NewAwsRestoreAccountConnectivityConfigWithDefaults() *AwsRestoreAccountConnectivityConfig`

NewAwsRestoreAccountConnectivityConfigWithDefaults instantiates a new AwsRestoreAccountConnectivityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcConfigs

`func (o *AwsRestoreAccountConnectivityConfig) GetVpcConfigs() []AwsVpcConnectivityConfig`

GetVpcConfigs returns the VpcConfigs field if non-nil, zero value otherwise.

### GetVpcConfigsOk

`func (o *AwsRestoreAccountConnectivityConfig) GetVpcConfigsOk() (*[]AwsVpcConnectivityConfig, bool)`

GetVpcConfigsOk returns a tuple with the VpcConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcConfigs

`func (o *AwsRestoreAccountConnectivityConfig) SetVpcConfigs(v []AwsVpcConnectivityConfig)`

SetVpcConfigs sets VpcConfigs field to given value.

### HasVpcConfigs

`func (o *AwsRestoreAccountConnectivityConfig) HasVpcConfigs() bool`

HasVpcConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


