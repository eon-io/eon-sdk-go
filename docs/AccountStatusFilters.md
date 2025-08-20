# AccountStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]AccountState**](AccountState.md) |  | [optional] 
**NotIn** | Pointer to [**[]AccountState**](AccountState.md) |  | [optional] 

## Methods

### NewAccountStatusFilters

`func NewAccountStatusFilters() *AccountStatusFilters`

NewAccountStatusFilters instantiates a new AccountStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusFiltersWithDefaults

`func NewAccountStatusFiltersWithDefaults() *AccountStatusFilters`

NewAccountStatusFiltersWithDefaults instantiates a new AccountStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *AccountStatusFilters) GetIn() []AccountState`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *AccountStatusFilters) GetInOk() (*[]AccountState, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *AccountStatusFilters) SetIn(v []AccountState)`

SetIn sets In field to given value.

### HasIn

`func (o *AccountStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *AccountStatusFilters) GetNotIn() []AccountState`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *AccountStatusFilters) GetNotInOk() (*[]AccountState, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *AccountStatusFilters) SetNotIn(v []AccountState)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *AccountStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


