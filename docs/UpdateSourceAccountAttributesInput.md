# UpdateSourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**NullableUpdateAwsSourceAccountAttributes**](UpdateAwsSourceAccountAttributes.md) |  | [optional] 
**Gcp** | Pointer to [**NullableUpdateGcpSourceAccountAttributes**](UpdateGcpSourceAccountAttributes.md) |  | [optional] 

## Methods

### NewUpdateSourceAccountAttributesInput

`func NewUpdateSourceAccountAttributesInput() *UpdateSourceAccountAttributesInput`

NewUpdateSourceAccountAttributesInput instantiates a new UpdateSourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSourceAccountAttributesInputWithDefaults

`func NewUpdateSourceAccountAttributesInputWithDefaults() *UpdateSourceAccountAttributesInput`

NewUpdateSourceAccountAttributesInputWithDefaults instantiates a new UpdateSourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *UpdateSourceAccountAttributesInput) GetAws() UpdateAwsSourceAccountAttributes`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *UpdateSourceAccountAttributesInput) GetAwsOk() (*UpdateAwsSourceAccountAttributes, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *UpdateSourceAccountAttributesInput) SetAws(v UpdateAwsSourceAccountAttributes)`

SetAws sets Aws field to given value.

### HasAws

`func (o *UpdateSourceAccountAttributesInput) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *UpdateSourceAccountAttributesInput) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *UpdateSourceAccountAttributesInput) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *UpdateSourceAccountAttributesInput) GetGcp() UpdateGcpSourceAccountAttributes`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *UpdateSourceAccountAttributesInput) GetGcpOk() (*UpdateGcpSourceAccountAttributes, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *UpdateSourceAccountAttributesInput) SetGcp(v UpdateGcpSourceAccountAttributes)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *UpdateSourceAccountAttributesInput) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *UpdateSourceAccountAttributesInput) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *UpdateSourceAccountAttributesInput) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


