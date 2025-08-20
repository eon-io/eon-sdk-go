# RestoreAccountsFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IdFilters**](IdFilters.md) |  | [optional] 
**ProviderAccountId** | Pointer to [**ProviderAccountIdFilters**](ProviderAccountIdFilters.md) |  | [optional] 
**AccountStatus** | Pointer to [**AccountStatusFilters**](AccountStatusFilters.md) |  | [optional] 
**CloudProvider** | Pointer to [**CloudProviderFilters**](CloudProviderFilters.md) |  | [optional] 
**AccountName** | Pointer to [**AccountNameFilters**](AccountNameFilters.md) |  | [optional] 

## Methods

### NewRestoreAccountsFilterConditions

`func NewRestoreAccountsFilterConditions() *RestoreAccountsFilterConditions`

NewRestoreAccountsFilterConditions instantiates a new RestoreAccountsFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAccountsFilterConditionsWithDefaults

`func NewRestoreAccountsFilterConditionsWithDefaults() *RestoreAccountsFilterConditions`

NewRestoreAccountsFilterConditionsWithDefaults instantiates a new RestoreAccountsFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreAccountsFilterConditions) GetId() IdFilters`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreAccountsFilterConditions) GetIdOk() (*IdFilters, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreAccountsFilterConditions) SetId(v IdFilters)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreAccountsFilterConditions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *RestoreAccountsFilterConditions) GetProviderAccountId() ProviderAccountIdFilters`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *RestoreAccountsFilterConditions) GetProviderAccountIdOk() (*ProviderAccountIdFilters, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *RestoreAccountsFilterConditions) SetProviderAccountId(v ProviderAccountIdFilters)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *RestoreAccountsFilterConditions) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.

### GetAccountStatus

`func (o *RestoreAccountsFilterConditions) GetAccountStatus() AccountStatusFilters`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *RestoreAccountsFilterConditions) GetAccountStatusOk() (*AccountStatusFilters, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *RestoreAccountsFilterConditions) SetAccountStatus(v AccountStatusFilters)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *RestoreAccountsFilterConditions) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *RestoreAccountsFilterConditions) GetCloudProvider() CloudProviderFilters`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *RestoreAccountsFilterConditions) GetCloudProviderOk() (*CloudProviderFilters, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *RestoreAccountsFilterConditions) SetCloudProvider(v CloudProviderFilters)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *RestoreAccountsFilterConditions) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetAccountName

`func (o *RestoreAccountsFilterConditions) GetAccountName() AccountNameFilters`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *RestoreAccountsFilterConditions) GetAccountNameOk() (*AccountNameFilters, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *RestoreAccountsFilterConditions) SetAccountName(v AccountNameFilters)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *RestoreAccountsFilterConditions) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


