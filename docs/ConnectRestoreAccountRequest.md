# ConnectRestoreAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Account display name in Eon. | 
**RestoreAccountConfig** | [**AccountConfigInput**](AccountConfigInput.md) |  | 

## Methods

### NewConnectRestoreAccountRequest

`func NewConnectRestoreAccountRequest(name string, restoreAccountConfig AccountConfigInput, ) *ConnectRestoreAccountRequest`

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


### GetRestoreAccountConfig

`func (o *ConnectRestoreAccountRequest) GetRestoreAccountConfig() AccountConfigInput`

GetRestoreAccountConfig returns the RestoreAccountConfig field if non-nil, zero value otherwise.

### GetRestoreAccountConfigOk

`func (o *ConnectRestoreAccountRequest) GetRestoreAccountConfigOk() (*AccountConfigInput, bool)`

GetRestoreAccountConfigOk returns a tuple with the RestoreAccountConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountConfig

`func (o *ConnectRestoreAccountRequest) SetRestoreAccountConfig(v AccountConfigInput)`

SetRestoreAccountConfig sets RestoreAccountConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


