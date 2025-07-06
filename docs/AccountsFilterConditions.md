# AccountsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IdFilters**](IdFilters.md) |  | [optional] 
**ProviderAccountId** | Pointer to [**ProviderAccountIdFilters**](ProviderAccountIdFilters.md) |  | [optional] 

## Methods

### NewAccountsFilterConditions

`func NewAccountsFilterConditions() *AccountsFilterConditions`

NewAccountsFilterConditions instantiates a new AccountsFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsFilterConditionsWithDefaults

`func NewAccountsFilterConditionsWithDefaults() *AccountsFilterConditions`

NewAccountsFilterConditionsWithDefaults instantiates a new AccountsFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountsFilterConditions) GetId() IdFilters`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountsFilterConditions) GetIdOk() (*IdFilters, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountsFilterConditions) SetId(v IdFilters)`

SetId sets Id field to given value.

### HasId

`func (o *AccountsFilterConditions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *AccountsFilterConditions) GetProviderAccountId() ProviderAccountIdFilters`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *AccountsFilterConditions) GetProviderAccountIdOk() (*ProviderAccountIdFilters, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *AccountsFilterConditions) SetProviderAccountId(v ProviderAccountIdFilters)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *AccountsFilterConditions) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


