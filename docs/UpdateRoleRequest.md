# UpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role | 
**PermissionGrants** | [**[]PermissionGrantInput**](PermissionGrantInput.md) | The permissions of the role | 
**AccessConditions** | Pointer to [**[]AccessCondition**](AccessCondition.md) | The condition of the role, keyed by name | [optional] 

## Methods

### NewUpdateRoleRequest

`func NewUpdateRoleRequest(name string, permissionGrants []PermissionGrantInput, ) *UpdateRoleRequest`

NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleRequestWithDefaults

`func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest`

NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionGrants

`func (o *UpdateRoleRequest) GetPermissionGrants() []PermissionGrantInput`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *UpdateRoleRequest) GetPermissionGrantsOk() (*[]PermissionGrantInput, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *UpdateRoleRequest) SetPermissionGrants(v []PermissionGrantInput)`

SetPermissionGrants sets PermissionGrants field to given value.


### GetAccessConditions

`func (o *UpdateRoleRequest) GetAccessConditions() []AccessCondition`

GetAccessConditions returns the AccessConditions field if non-nil, zero value otherwise.

### GetAccessConditionsOk

`func (o *UpdateRoleRequest) GetAccessConditionsOk() (*[]AccessCondition, bool)`

GetAccessConditionsOk returns a tuple with the AccessConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConditions

`func (o *UpdateRoleRequest) SetAccessConditions(v []AccessCondition)`

SetAccessConditions sets AccessConditions field to given value.

### HasAccessConditions

`func (o *UpdateRoleRequest) HasAccessConditions() bool`

HasAccessConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


