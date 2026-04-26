# UpdateRestoreAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Display name for the restore account in Eon. | [optional] 
**RestoreAccountAttributes** | Pointer to [**UpdateRestoreAccountAttributesInput**](UpdateRestoreAccountAttributesInput.md) |  | [optional] 

## Methods

### NewUpdateRestoreAccountRequest

`func NewUpdateRestoreAccountRequest() *UpdateRestoreAccountRequest`

NewUpdateRestoreAccountRequest instantiates a new UpdateRestoreAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRestoreAccountRequestWithDefaults

`func NewUpdateRestoreAccountRequestWithDefaults() *UpdateRestoreAccountRequest`

NewUpdateRestoreAccountRequestWithDefaults instantiates a new UpdateRestoreAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRestoreAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRestoreAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRestoreAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRestoreAccountRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateRestoreAccountRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateRestoreAccountRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRestoreAccountAttributes

`func (o *UpdateRestoreAccountRequest) GetRestoreAccountAttributes() UpdateRestoreAccountAttributesInput`

GetRestoreAccountAttributes returns the RestoreAccountAttributes field if non-nil, zero value otherwise.

### GetRestoreAccountAttributesOk

`func (o *UpdateRestoreAccountRequest) GetRestoreAccountAttributesOk() (*UpdateRestoreAccountAttributesInput, bool)`

GetRestoreAccountAttributesOk returns a tuple with the RestoreAccountAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountAttributes

`func (o *UpdateRestoreAccountRequest) SetRestoreAccountAttributes(v UpdateRestoreAccountAttributesInput)`

SetRestoreAccountAttributes sets RestoreAccountAttributes field to given value.

### HasRestoreAccountAttributes

`func (o *UpdateRestoreAccountRequest) HasRestoreAccountAttributes() bool`

HasRestoreAccountAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


