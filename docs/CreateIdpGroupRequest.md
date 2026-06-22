# CreateIdpGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpId** | **string** | [ID of the identity provider](./list-idps) the group belongs to. | 
**ProviderGroupId** | **string** | Identity-provider-assigned group ID. Must match the exact group ID passed by the identity provider to Eon during SAML sign-on.  | 
**RoleIds** | **[]string** | List of [role IDs](../roles/list-roles) to assign to the group. | 
**DisplayName** | Pointer to **NullableString** | Optional human-readable label for the group mapping. For display and management only; it is not used when matching SAML groups during sign-on.  | [optional] 

## Methods

### NewCreateIdpGroupRequest

`func NewCreateIdpGroupRequest(idpId string, providerGroupId string, roleIds []string, ) *CreateIdpGroupRequest`

NewCreateIdpGroupRequest instantiates a new CreateIdpGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdpGroupRequestWithDefaults

`func NewCreateIdpGroupRequestWithDefaults() *CreateIdpGroupRequest`

NewCreateIdpGroupRequestWithDefaults instantiates a new CreateIdpGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpId

`func (o *CreateIdpGroupRequest) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *CreateIdpGroupRequest) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *CreateIdpGroupRequest) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetProviderGroupId

`func (o *CreateIdpGroupRequest) GetProviderGroupId() string`

GetProviderGroupId returns the ProviderGroupId field if non-nil, zero value otherwise.

### GetProviderGroupIdOk

`func (o *CreateIdpGroupRequest) GetProviderGroupIdOk() (*string, bool)`

GetProviderGroupIdOk returns a tuple with the ProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupId

`func (o *CreateIdpGroupRequest) SetProviderGroupId(v string)`

SetProviderGroupId sets ProviderGroupId field to given value.


### GetRoleIds

`func (o *CreateIdpGroupRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateIdpGroupRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateIdpGroupRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.


### GetDisplayName

`func (o *CreateIdpGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateIdpGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateIdpGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateIdpGroupRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateIdpGroupRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateIdpGroupRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


