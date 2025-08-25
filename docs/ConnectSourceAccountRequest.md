# ConnectSourceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Account display name in Eon. | [optional] 
**SourceAccountAttributes** | [**SourceAccountAttributesInput**](SourceAccountAttributesInput.md) |  | 

## Methods

### NewConnectSourceAccountRequest

`func NewConnectSourceAccountRequest(sourceAccountAttributes SourceAccountAttributesInput, ) *ConnectSourceAccountRequest`

NewConnectSourceAccountRequest instantiates a new ConnectSourceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectSourceAccountRequestWithDefaults

`func NewConnectSourceAccountRequestWithDefaults() *ConnectSourceAccountRequest`

NewConnectSourceAccountRequestWithDefaults instantiates a new ConnectSourceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectSourceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectSourceAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectSourceAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectSourceAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceAccountAttributes

`func (o *ConnectSourceAccountRequest) GetSourceAccountAttributes() SourceAccountAttributesInput`

GetSourceAccountAttributes returns the SourceAccountAttributes field if non-nil, zero value otherwise.

### GetSourceAccountAttributesOk

`func (o *ConnectSourceAccountRequest) GetSourceAccountAttributesOk() (*SourceAccountAttributesInput, bool)`

GetSourceAccountAttributesOk returns a tuple with the SourceAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountAttributes

`func (o *ConnectSourceAccountRequest) SetSourceAccountAttributes(v SourceAccountAttributesInput)`

SetSourceAccountAttributes sets SourceAccountAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


