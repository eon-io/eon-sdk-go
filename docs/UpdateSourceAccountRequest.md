# UpdateSourceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Display name for the source account in Eon. | [optional] 
**SourceAccountAttributes** | Pointer to [**UpdateSourceAccountAttributesInput**](UpdateSourceAccountAttributesInput.md) |  | [optional] 

## Methods

### NewUpdateSourceAccountRequest

`func NewUpdateSourceAccountRequest() *UpdateSourceAccountRequest`

NewUpdateSourceAccountRequest instantiates a new UpdateSourceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourceAccountRequestWithDefaults

`func NewUpdateSourceAccountRequestWithDefaults() *UpdateSourceAccountRequest`

NewUpdateSourceAccountRequestWithDefaults instantiates a new UpdateSourceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSourceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSourceAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSourceAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSourceAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSourceAccountRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSourceAccountRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceAccountAttributes

`func (o *UpdateSourceAccountRequest) GetSourceAccountAttributes() UpdateSourceAccountAttributesInput`

GetSourceAccountAttributes returns the SourceAccountAttributes field if non-nil, zero value otherwise.

### GetSourceAccountAttributesOk

`func (o *UpdateSourceAccountRequest) GetSourceAccountAttributesOk() (*UpdateSourceAccountAttributesInput, bool)`

GetSourceAccountAttributesOk returns a tuple with the SourceAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountAttributes

`func (o *UpdateSourceAccountRequest) SetSourceAccountAttributes(v UpdateSourceAccountAttributesInput)`

SetSourceAccountAttributes sets SourceAccountAttributes field to given value.

### HasSourceAccountAttributes

`func (o *UpdateSourceAccountRequest) HasSourceAccountAttributes() bool`

HasSourceAccountAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


