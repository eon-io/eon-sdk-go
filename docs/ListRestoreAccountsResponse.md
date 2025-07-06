# ListRestoreAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]RestoreAccount**](RestoreAccount.md) | List of restore accounts. | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of restore accounts that matched the filter options. | [optional] 

## Methods

### NewListRestoreAccountsResponse

`func NewListRestoreAccountsResponse(accounts []RestoreAccount, ) *ListRestoreAccountsResponse`

NewListRestoreAccountsResponse instantiates a new ListRestoreAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestoreAccountsResponseWithDefaults

`func NewListRestoreAccountsResponseWithDefaults() *ListRestoreAccountsResponse`

NewListRestoreAccountsResponseWithDefaults instantiates a new ListRestoreAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ListRestoreAccountsResponse) GetAccounts() []RestoreAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListRestoreAccountsResponse) GetAccountsOk() (*[]RestoreAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListRestoreAccountsResponse) SetAccounts(v []RestoreAccount)`

SetAccounts sets Accounts field to given value.


### GetNextToken

`func (o *ListRestoreAccountsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListRestoreAccountsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListRestoreAccountsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListRestoreAccountsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListRestoreAccountsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListRestoreAccountsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListRestoreAccountsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListRestoreAccountsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


