# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionType** | [**PermissionType**](PermissionType.md) |  | 
**Description** | **string** | Description of the actions the permission allows. | 
**AllowConditions** | **bool** | Whether the permission can be restricted with access conditions. Relevant when the permission is used in a user role.  | [default to false]

## Methods

### NewPermission

`func NewPermission(permissionType PermissionType, description string, allowConditions bool, ) *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionType

`func (o *Permission) GetPermissionType() PermissionType`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *Permission) GetPermissionTypeOk() (*PermissionType, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *Permission) SetPermissionType(v PermissionType)`

SetPermissionType sets PermissionType field to given value.


### GetDescription

`func (o *Permission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Permission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Permission) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAllowConditions

`func (o *Permission) GetAllowConditions() bool`

GetAllowConditions returns the AllowConditions field if non-nil, zero value otherwise.

### GetAllowConditionsOk

`func (o *Permission) GetAllowConditionsOk() (*bool, bool)`

GetAllowConditionsOk returns a tuple with the AllowConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConditions

`func (o *Permission) SetAllowConditions(v bool)`

SetAllowConditions sets AllowConditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


