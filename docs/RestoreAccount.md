# RestoreAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned restore account ID. | 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**Name** | **string** | Account display name in Eon. | 
**ProviderAccountName** | Pointer to **string** | Display name inherited from the cloud provider. | [optional] 
**Status** | [**AccountState**](AccountState.md) |  | 
**Version** | Pointer to [**AccountVersion**](AccountVersion.md) |  | [optional] 
**ConnectedTime** | Pointer to **time.Time** | Date and time the account was connected to Eon. | [optional] 
**RestoreAccountAttributes** | [**RestoreAccountCloudAttributes**](RestoreAccountCloudAttributes.md) |  | 

## Methods

### NewRestoreAccount

`func NewRestoreAccount(id string, providerAccountId string, name string, status AccountState, restoreAccountAttributes RestoreAccountCloudAttributes, ) *RestoreAccount`

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


### GetName

`func (o *RestoreAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreAccount) SetName(v string)`

SetName sets Name field to given value.


### GetProviderAccountName

`func (o *RestoreAccount) GetProviderAccountName() string`

GetProviderAccountName returns the ProviderAccountName field if non-nil, zero value otherwise.

### GetProviderAccountNameOk

`func (o *RestoreAccount) GetProviderAccountNameOk() (*string, bool)`

GetProviderAccountNameOk returns a tuple with the ProviderAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountName

`func (o *RestoreAccount) SetProviderAccountName(v string)`

SetProviderAccountName sets ProviderAccountName field to given value.

### HasProviderAccountName

`func (o *RestoreAccount) HasProviderAccountName() bool`

HasProviderAccountName returns a boolean if a field has been set.

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


### GetVersion

`func (o *RestoreAccount) GetVersion() AccountVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RestoreAccount) GetVersionOk() (*AccountVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RestoreAccount) SetVersion(v AccountVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RestoreAccount) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConnectedTime

`func (o *RestoreAccount) GetConnectedTime() time.Time`

GetConnectedTime returns the ConnectedTime field if non-nil, zero value otherwise.

### GetConnectedTimeOk

`func (o *RestoreAccount) GetConnectedTimeOk() (*time.Time, bool)`

GetConnectedTimeOk returns a tuple with the ConnectedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedTime

`func (o *RestoreAccount) SetConnectedTime(v time.Time)`

SetConnectedTime sets ConnectedTime field to given value.

### HasConnectedTime

`func (o *RestoreAccount) HasConnectedTime() bool`

HasConnectedTime returns a boolean if a field has been set.

### GetRestoreAccountAttributes

`func (o *RestoreAccount) GetRestoreAccountAttributes() RestoreAccountCloudAttributes`

GetRestoreAccountAttributes returns the RestoreAccountAttributes field if non-nil, zero value otherwise.

### GetRestoreAccountAttributesOk

`func (o *RestoreAccount) GetRestoreAccountAttributesOk() (*RestoreAccountCloudAttributes, bool)`

GetRestoreAccountAttributesOk returns a tuple with the RestoreAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountAttributes

`func (o *RestoreAccount) SetRestoreAccountAttributes(v RestoreAccountCloudAttributes)`

SetRestoreAccountAttributes sets RestoreAccountAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


