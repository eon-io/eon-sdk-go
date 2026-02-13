# CreateIdpGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpId** | **string** | The ID of the Identity Provider this group belongs to. | 
**ProviderGroupId** | **string** | The group identifier from the Identity Provider. | 
**RoleIds** | **[]string** | List of role IDs to assign to this IDP group. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


