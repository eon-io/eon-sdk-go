# AccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**Aws** | Pointer to [**NullableAwsAccountConfig**](AwsAccountConfig.md) |  | [optional] 

## Methods

### NewAccountConfig

`func NewAccountConfig() *AccountConfig`

NewAccountConfig instantiates a new AccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigWithDefaults

`func NewAccountConfigWithDefaults() *AccountConfig`

NewAccountConfigWithDefaults instantiates a new AccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AccountConfig) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AccountConfig) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AccountConfig) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AccountConfig) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAws

`func (o *AccountConfig) GetAws() AwsAccountConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *AccountConfig) GetAwsOk() (*AwsAccountConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *AccountConfig) SetAws(v AwsAccountConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *AccountConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *AccountConfig) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *AccountConfig) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


