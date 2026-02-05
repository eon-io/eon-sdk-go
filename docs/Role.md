# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Role ID. | 
**Name** | **string** | Role display name. | 
**IsBuiltInRole** | **bool** | Whether the role is a default role. If &#x60;true&#x60;, the role is a default role provided by Eon and can&#39;t be modified or deleted. If &#x60;false&#x60;, the role is a custom, user-created role.  | 
**PermissionGrants** | [**[]PermissionGrant**](PermissionGrant.md) | List of permissions granted by the role. | 
**AccessConditions** | Pointer to [**[]AccessCondition**](AccessCondition.md) | Sets of access conditions that restrict the resources a permission is granted for. IDs are set by you and are applied to the relevant permission in &#x60;permissionGrants&#x60;. An access condition can be applied to more than one permission grant.  | [optional] 

## Methods

### NewRole

`func NewRole(id string, name string, isBuiltInRole bool, permissionGrants []PermissionGrant, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetIsBuiltInRole

`func (o *Role) GetIsBuiltInRole() bool`

GetIsBuiltInRole returns the IsBuiltInRole field if non-nil, zero value otherwise.

### GetIsBuiltInRoleOk

`func (o *Role) GetIsBuiltInRoleOk() (*bool, bool)`

GetIsBuiltInRoleOk returns a tuple with the IsBuiltInRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltInRole

`func (o *Role) SetIsBuiltInRole(v bool)`

SetIsBuiltInRole sets IsBuiltInRole field to given value.


### GetPermissionGrants

`func (o *Role) GetPermissionGrants() []PermissionGrant`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *Role) GetPermissionGrantsOk() (*[]PermissionGrant, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *Role) SetPermissionGrants(v []PermissionGrant)`

SetPermissionGrants sets PermissionGrants field to given value.


### GetAccessConditions

`func (o *Role) GetAccessConditions() []AccessCondition`

GetAccessConditions returns the AccessConditions field if non-nil, zero value otherwise.

### GetAccessConditionsOk

`func (o *Role) GetAccessConditionsOk() (*[]AccessCondition, bool)`

GetAccessConditionsOk returns a tuple with the AccessConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConditions

`func (o *Role) SetAccessConditions(v []AccessCondition)`

SetAccessConditions sets AccessConditions field to given value.

### HasAccessConditions

`func (o *Role) HasAccessConditions() bool`

HasAccessConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


