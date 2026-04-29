# UpdateGcpSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | Pointer to **NullableString** | Email of the service account Eon impersonates to access the GCP project. Only the service account name portion can be changed; the GCP project ID must remain the same. To use a different GCP project, disconnect and connect a new source account.  | [optional] 

## Methods

### NewUpdateGcpSourceAccountAttributes

`func NewUpdateGcpSourceAccountAttributes() *UpdateGcpSourceAccountAttributes`

NewUpdateGcpSourceAccountAttributes instantiates a new UpdateGcpSourceAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGcpSourceAccountAttributesWithDefaults

`func NewUpdateGcpSourceAccountAttributesWithDefaults() *UpdateGcpSourceAccountAttributes`

NewUpdateGcpSourceAccountAttributesWithDefaults instantiates a new UpdateGcpSourceAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *UpdateGcpSourceAccountAttributes) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *UpdateGcpSourceAccountAttributes) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *UpdateGcpSourceAccountAttributes) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *UpdateGcpSourceAccountAttributes) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *UpdateGcpSourceAccountAttributes) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *UpdateGcpSourceAccountAttributes) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


