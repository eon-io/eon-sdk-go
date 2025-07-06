# SourceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned account ID. | 
**Name** | **string** | Account display name in Eon. | 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**Status** | [**AccountState**](AccountState.md) |  | 
**SourceAccountConfig** | [**AccountConfig**](AccountConfig.md) |  | 

## Methods

### NewSourceAccount

`func NewSourceAccount(id string, name string, providerAccountId string, status AccountState, sourceAccountConfig AccountConfig, ) *SourceAccount`

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


### GetSourceAccountConfig

`func (o *SourceAccount) GetSourceAccountConfig() AccountConfig`

GetSourceAccountConfig returns the SourceAccountConfig field if non-nil, zero value otherwise.

### GetSourceAccountConfigOk

`func (o *SourceAccount) GetSourceAccountConfigOk() (*AccountConfig, bool)`

GetSourceAccountConfigOk returns a tuple with the SourceAccountConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountConfig

`func (o *SourceAccount) SetSourceAccountConfig(v AccountConfig)`

SetSourceAccountConfig sets SourceAccountConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


