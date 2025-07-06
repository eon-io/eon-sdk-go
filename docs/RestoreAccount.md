# RestoreAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned restore account ID. | 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**Status** | [**AccountState**](AccountState.md) |  | 
**RestoreAccountAttributes** | Pointer to [**AccountConfig**](AccountConfig.md) |  | [optional] 

## Methods

### NewRestoreAccount

`func NewRestoreAccount(id string, providerAccountId string, status AccountState, ) *RestoreAccount`

NewRestoreAccount instantiates a new RestoreAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAccountWithDefaults

`func NewRestoreAccountWithDefaults() *RestoreAccount`

NewRestoreAccountWithDefaults instantiates a new RestoreAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreAccount) SetId(v string)`

SetId sets Id field to given value.


### GetProviderAccountId

`func (o *RestoreAccount) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *RestoreAccount) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *RestoreAccount) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetStatus

`func (o *RestoreAccount) GetStatus() AccountState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreAccount) GetStatusOk() (*AccountState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreAccount) SetStatus(v AccountState)`

SetStatus sets Status field to given value.


### GetRestoreAccountAttributes

`func (o *RestoreAccount) GetRestoreAccountAttributes() AccountConfig`

GetRestoreAccountAttributes returns the RestoreAccountAttributes field if non-nil, zero value otherwise.

### GetRestoreAccountAttributesOk

`func (o *RestoreAccount) GetRestoreAccountAttributesOk() (*AccountConfig, bool)`

GetRestoreAccountAttributesOk returns a tuple with the RestoreAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountAttributes

`func (o *RestoreAccount) SetRestoreAccountAttributes(v AccountConfig)`

SetRestoreAccountAttributes sets RestoreAccountAttributes field to given value.

### HasRestoreAccountAttributes

`func (o *RestoreAccount) HasRestoreAccountAttributes() bool`

HasRestoreAccountAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


