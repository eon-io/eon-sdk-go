# UpdateGcpRestoreAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | Pointer to **NullableString** | Email of the service account Eon impersonates to access the GCP project. Only the service account name portion can be changed; the GCP project ID must remain the same. To use a different GCP project, disconnect and connect a new restore account.  | [optional] 

## Methods

### NewUpdateGcpRestoreAccountAttributes

`func NewUpdateGcpRestoreAccountAttributes() *UpdateGcpRestoreAccountAttributes`

NewUpdateGcpRestoreAccountAttributes instantiates a new UpdateGcpRestoreAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGcpRestoreAccountAttributesWithDefaults

`func NewUpdateGcpRestoreAccountAttributesWithDefaults() *UpdateGcpRestoreAccountAttributes`

NewUpdateGcpRestoreAccountAttributesWithDefaults instantiates a new UpdateGcpRestoreAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *UpdateGcpRestoreAccountAttributes) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *UpdateGcpRestoreAccountAttributes) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *UpdateGcpRestoreAccountAttributes) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *UpdateGcpRestoreAccountAttributes) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *UpdateGcpRestoreAccountAttributes) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *UpdateGcpRestoreAccountAttributes) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


