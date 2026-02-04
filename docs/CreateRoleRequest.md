# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role | 
**PermissionGrants** | [**[]PermissionGrantInput**](PermissionGrantInput.md) | The permissions of the role | 
**AccessConditions** | Pointer to [**[]AccessCondition**](AccessCondition.md) | The condition of the role, keyed by name | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest(name string, permissionGrants []PermissionGrantInput, ) *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionGrants

`func (o *CreateRoleRequest) GetPermissionGrants() []PermissionGrantInput`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *CreateRoleRequest) GetPermissionGrantsOk() (*[]PermissionGrantInput, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *CreateRoleRequest) SetPermissionGrants(v []PermissionGrantInput)`

SetPermissionGrants sets PermissionGrants field to given value.


### GetAccessConditions

`func (o *CreateRoleRequest) GetAccessConditions() []AccessCondition`

GetAccessConditions returns the AccessConditions field if non-nil, zero value otherwise.

### GetAccessConditionsOk

`func (o *CreateRoleRequest) GetAccessConditionsOk() (*[]AccessCondition, bool)`

GetAccessConditionsOk returns a tuple with the AccessConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConditions

`func (o *CreateRoleRequest) SetAccessConditions(v []AccessCondition)`

SetAccessConditions sets AccessConditions field to given value.

### HasAccessConditions

`func (o *CreateRoleRequest) HasAccessConditions() bool`

HasAccessConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


