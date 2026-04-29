# UpdateRestoreAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**NullableUpdateAwsRestoreAccountAttributes**](UpdateAwsRestoreAccountAttributes.md) |  | [optional] 
**Gcp** | Pointer to [**NullableUpdateGcpRestoreAccountAttributes**](UpdateGcpRestoreAccountAttributes.md) |  | [optional] 

## Methods

### NewUpdateRestoreAccountAttributesInput

`func NewUpdateRestoreAccountAttributesInput() *UpdateRestoreAccountAttributesInput`

NewUpdateRestoreAccountAttributesInput instantiates a new UpdateRestoreAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRestoreAccountAttributesInputWithDefaults

`func NewUpdateRestoreAccountAttributesInputWithDefaults() *UpdateRestoreAccountAttributesInput`

NewUpdateRestoreAccountAttributesInputWithDefaults instantiates a new UpdateRestoreAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *UpdateRestoreAccountAttributesInput) GetAws() UpdateAwsRestoreAccountAttributes`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *UpdateRestoreAccountAttributesInput) GetAwsOk() (*UpdateAwsRestoreAccountAttributes, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *UpdateRestoreAccountAttributesInput) SetAws(v UpdateAwsRestoreAccountAttributes)`

SetAws sets Aws field to given value.

### HasAws

`func (o *UpdateRestoreAccountAttributesInput) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *UpdateRestoreAccountAttributesInput) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *UpdateRestoreAccountAttributesInput) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *UpdateRestoreAccountAttributesInput) GetGcp() UpdateGcpRestoreAccountAttributes`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *UpdateRestoreAccountAttributesInput) GetGcpOk() (*UpdateGcpRestoreAccountAttributes, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *UpdateRestoreAccountAttributesInput) SetGcp(v UpdateGcpRestoreAccountAttributes)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *UpdateRestoreAccountAttributesInput) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *UpdateRestoreAccountAttributesInput) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *UpdateRestoreAccountAttributesInput) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


