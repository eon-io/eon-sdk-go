# UpdateIdpGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleIds** | **[]string** | List of [role IDs](../roles/list-roles) to assign to the group. This replaces all existing role assignments.  | 
**DisplayName** | Pointer to **NullableString** | Optional human-readable label for the group mapping. For display and management only; it is not used when matching SAML groups during sign-on. Omit (or send null) to leave the existing label unchanged; send an empty string to clear it.  | [optional] 

## Methods

### NewUpdateIdpGroupRequest

`func NewUpdateIdpGroupRequest(roleIds []string, ) *UpdateIdpGroupRequest`

NewUpdateIdpGroupRequest instantiates a new UpdateIdpGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdpGroupRequestWithDefaults

`func NewUpdateIdpGroupRequestWithDefaults() *UpdateIdpGroupRequest`

NewUpdateIdpGroupRequestWithDefaults instantiates a new UpdateIdpGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleIds

`func (o *UpdateIdpGroupRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UpdateIdpGroupRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UpdateIdpGroupRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.


### GetDisplayName

`func (o *UpdateIdpGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateIdpGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateIdpGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateIdpGroupRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateIdpGroupRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateIdpGroupRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


