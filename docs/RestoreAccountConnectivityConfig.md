# RestoreAccountConnectivityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Config** | [**ProviderRestoreAccountConnectivityConfig**](ProviderRestoreAccountConnectivityConfig.md) |  | 

## Methods

### NewRestoreAccountConnectivityConfig

`func NewRestoreAccountConnectivityConfig(restoreAccountId string, config ProviderRestoreAccountConnectivityConfig, ) *RestoreAccountConnectivityConfig`

NewRestoreAccountConnectivityConfig instantiates a new RestoreAccountConnectivityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAccountConnectivityConfigWithDefaults

`func NewRestoreAccountConnectivityConfigWithDefaults() *RestoreAccountConnectivityConfig`

NewRestoreAccountConnectivityConfigWithDefaults instantiates a new RestoreAccountConnectivityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *RestoreAccountConnectivityConfig) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreAccountConnectivityConfig) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreAccountConnectivityConfig) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetConfig

`func (o *RestoreAccountConnectivityConfig) GetConfig() ProviderRestoreAccountConnectivityConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RestoreAccountConnectivityConfig) GetConfigOk() (*ProviderRestoreAccountConnectivityConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RestoreAccountConnectivityConfig) SetConfig(v ProviderRestoreAccountConnectivityConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


