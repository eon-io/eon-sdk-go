# AccountIdCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**AccountIds** | **[]string** |  | 

## Methods

### NewAccountIdCondition

`func NewAccountIdCondition(operator ScalarOperators, accountIds []string, ) *AccountIdCondition`

NewAccountIdCondition instantiates a new AccountIdCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdConditionWithDefaults

`func NewAccountIdConditionWithDefaults() *AccountIdCondition`

NewAccountIdConditionWithDefaults instantiates a new AccountIdCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *AccountIdCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AccountIdCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AccountIdCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetAccountIds

`func (o *AccountIdCondition) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AccountIdCondition) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AccountIdCondition) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


