# ConnectRestoreAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Account display name in Eon. | [optional] 
**RestoreAccountAttributes** | [**RestoreAccountConfigInput**](RestoreAccountConfigInput.md) |  | 

## Methods

### NewConnectRestoreAccountRequest

`func NewConnectRestoreAccountRequest(restoreAccountAttributes RestoreAccountConfigInput, ) *ConnectRestoreAccountRequest`

NewConnectRestoreAccountRequest instantiates a new ConnectRestoreAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectRestoreAccountRequestWithDefaults

`func NewConnectRestoreAccountRequestWithDefaults() *ConnectRestoreAccountRequest`

NewConnectRestoreAccountRequestWithDefaults instantiates a new ConnectRestoreAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectRestoreAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectRestoreAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectRestoreAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectRestoreAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestoreAccountAttributes

`func (o *ConnectRestoreAccountRequest) GetRestoreAccountAttributes() RestoreAccountConfigInput`

GetRestoreAccountAttributes returns the RestoreAccountAttributes field if non-nil, zero value otherwise.

### GetRestoreAccountAttributesOk

`func (o *ConnectRestoreAccountRequest) GetRestoreAccountAttributesOk() (*RestoreAccountConfigInput, bool)`

GetRestoreAccountAttributesOk returns a tuple with the RestoreAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountAttributes

`func (o *ConnectRestoreAccountRequest) SetRestoreAccountAttributes(v RestoreAccountConfigInput)`

SetRestoreAccountAttributes sets RestoreAccountAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


