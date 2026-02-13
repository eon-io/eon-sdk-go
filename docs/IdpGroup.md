# IdpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | System-generated unique identifier for the IDP group. | 
**IdpId** | **string** | The ID of the Identity Provider this group belongs to. | 
**ProviderGroupId** | **string** | The group identifier from the Identity Provider. | 
**RoleIds** | **[]string** | List of role IDs assigned to this IDP group. | 

## Methods

### NewIdpGroup

`func NewIdpGroup(id string, idpId string, providerGroupId string, roleIds []string, ) *IdpGroup`

NewIdpGroup instantiates a new IdpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpGroupWithDefaults

`func NewIdpGroupWithDefaults() *IdpGroup`

NewIdpGroupWithDefaults instantiates a new IdpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpGroup) SetId(v string)`

SetId sets Id field to given value.


### GetIdpId

`func (o *IdpGroup) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *IdpGroup) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *IdpGroup) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetProviderGroupId

`func (o *IdpGroup) GetProviderGroupId() string`

GetProviderGroupId returns the ProviderGroupId field if non-nil, zero value otherwise.

### GetProviderGroupIdOk

`func (o *IdpGroup) GetProviderGroupIdOk() (*string, bool)`

GetProviderGroupIdOk returns a tuple with the ProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderGroupId

`func (o *IdpGroup) SetProviderGroupId(v string)`

SetProviderGroupId sets ProviderGroupId field to given value.


### GetRoleIds

`func (o *IdpGroup) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *IdpGroup) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *IdpGroup) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


