# ConnectRestoreAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccount** | Pointer to [**RestoreAccount**](RestoreAccount.md) |  | [optional] 
**ActionApprovalRequest** | Pointer to [**NullableMPARequest**](MPARequest.md) |  | [optional] 

## Methods

### NewConnectRestoreAccountResponse

`func NewConnectRestoreAccountResponse() *ConnectRestoreAccountResponse`

NewConnectRestoreAccountResponse instantiates a new ConnectRestoreAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectRestoreAccountResponseWithDefaults

`func NewConnectRestoreAccountResponseWithDefaults() *ConnectRestoreAccountResponse`

NewConnectRestoreAccountResponseWithDefaults instantiates a new ConnectRestoreAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccount

`func (o *ConnectRestoreAccountResponse) GetRestoreAccount() RestoreAccount`

GetRestoreAccount returns the RestoreAccount field if non-nil, zero value otherwise.

### GetRestoreAccountOk

`func (o *ConnectRestoreAccountResponse) GetRestoreAccountOk() (*RestoreAccount, bool)`

GetRestoreAccountOk returns a tuple with the RestoreAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccount

`func (o *ConnectRestoreAccountResponse) SetRestoreAccount(v RestoreAccount)`

SetRestoreAccount sets RestoreAccount field to given value.

### HasRestoreAccount

`func (o *ConnectRestoreAccountResponse) HasRestoreAccount() bool`

HasRestoreAccount returns a boolean if a field has been set.

### GetActionApprovalRequest

`func (o *ConnectRestoreAccountResponse) GetActionApprovalRequest() MPARequest`

GetActionApprovalRequest returns the ActionApprovalRequest field if non-nil, zero value otherwise.

### GetActionApprovalRequestOk

`func (o *ConnectRestoreAccountResponse) GetActionApprovalRequestOk() (*MPARequest, bool)`

GetActionApprovalRequestOk returns a tuple with the ActionApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApprovalRequest

`func (o *ConnectRestoreAccountResponse) SetActionApprovalRequest(v MPARequest)`

SetActionApprovalRequest sets ActionApprovalRequest field to given value.

### HasActionApprovalRequest

`func (o *ConnectRestoreAccountResponse) HasActionApprovalRequest() bool`

HasActionApprovalRequest returns a boolean if a field has been set.

### SetActionApprovalRequestNil

`func (o *ConnectRestoreAccountResponse) SetActionApprovalRequestNil(b bool)`

 SetActionApprovalRequestNil sets the value for ActionApprovalRequest to be an explicit nil

### UnsetActionApprovalRequest
`func (o *ConnectRestoreAccountResponse) UnsetActionApprovalRequest()`

UnsetActionApprovalRequest ensures that no value is present for ActionApprovalRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


