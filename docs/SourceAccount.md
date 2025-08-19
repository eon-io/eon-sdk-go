# SourceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned source account ID. | 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**Name** | **string** | Account display name in Eon. | 
**ProviderAccountName** | Pointer to **string** | The display name of the account from the cloud provider | [optional] 
**Status** | [**AccountState**](AccountState.md) |  | 
**Version** | Pointer to [**AccountVersion**](AccountVersion.md) |  | [optional] 
**ConnectionTime** | Pointer to **time.Time** | The date and time the account was connected to Eon | [optional] 
**SourceAccountAttributes** | Pointer to [**AccountConfig**](AccountConfig.md) |  | [optional] 

## Methods

### NewSourceAccount

`func NewSourceAccount(id string, providerAccountId string, name string, status AccountState, ) *SourceAccount`

NewSourceAccount instantiates a new SourceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountWithDefaults

`func NewSourceAccountWithDefaults() *SourceAccount`

NewSourceAccountWithDefaults instantiates a new SourceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAccount) SetId(v string)`

SetId sets Id field to given value.


### GetProviderAccountId

`func (o *SourceAccount) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *SourceAccount) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *SourceAccount) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetName

`func (o *SourceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetProviderAccountName

`func (o *SourceAccount) GetProviderAccountName() string`

GetProviderAccountName returns the ProviderAccountName field if non-nil, zero value otherwise.

### GetProviderAccountNameOk

`func (o *SourceAccount) GetProviderAccountNameOk() (*string, bool)`

GetProviderAccountNameOk returns a tuple with the ProviderAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountName

`func (o *SourceAccount) SetProviderAccountName(v string)`

SetProviderAccountName sets ProviderAccountName field to given value.

### HasProviderAccountName

`func (o *SourceAccount) HasProviderAccountName() bool`

HasProviderAccountName returns a boolean if a field has been set.

### GetStatus

`func (o *SourceAccount) GetStatus() AccountState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceAccount) GetStatusOk() (*AccountState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceAccount) SetStatus(v AccountState)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *SourceAccount) GetVersion() AccountVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SourceAccount) GetVersionOk() (*AccountVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SourceAccount) SetVersion(v AccountVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SourceAccount) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConnectionTime

`func (o *SourceAccount) GetConnectionTime() time.Time`

GetConnectionTime returns the ConnectionTime field if non-nil, zero value otherwise.

### GetConnectionTimeOk

`func (o *SourceAccount) GetConnectionTimeOk() (*time.Time, bool)`

GetConnectionTimeOk returns a tuple with the ConnectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTime

`func (o *SourceAccount) SetConnectionTime(v time.Time)`

SetConnectionTime sets ConnectionTime field to given value.

### HasConnectionTime

`func (o *SourceAccount) HasConnectionTime() bool`

HasConnectionTime returns a boolean if a field has been set.

### GetSourceAccountAttributes

`func (o *SourceAccount) GetSourceAccountAttributes() AccountConfig`

GetSourceAccountAttributes returns the SourceAccountAttributes field if non-nil, zero value otherwise.

### GetSourceAccountAttributesOk

`func (o *SourceAccount) GetSourceAccountAttributesOk() (*AccountConfig, bool)`

GetSourceAccountAttributesOk returns a tuple with the SourceAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountAttributes

`func (o *SourceAccount) SetSourceAccountAttributes(v AccountConfig)`

SetSourceAccountAttributes sets SourceAccountAttributes field to given value.

### HasSourceAccountAttributes

`func (o *SourceAccount) HasSourceAccountAttributes() bool`

HasSourceAccountAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


