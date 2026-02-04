# PermissionGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | [**PermissionType**](PermissionType.md) |  | 
**AccessConditionId** | Pointer to **string** | Optional ID of the access condition associated with this permission grant. MUST be null for permissions where allowConditions&#x3D;false. | [optional] 

## Methods

### NewPermissionGrant

`func NewPermissionGrant(permission PermissionType, ) *PermissionGrant`

NewPermissionGrant instantiates a new PermissionGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantWithDefaults

`func NewPermissionGrantWithDefaults() *PermissionGrant`

NewPermissionGrantWithDefaults instantiates a new PermissionGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionGrant) GetPermission() PermissionType`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionGrant) GetPermissionOk() (*PermissionType, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionGrant) SetPermission(v PermissionType)`

SetPermission sets Permission field to given value.


### GetAccessConditionId

`func (o *PermissionGrant) GetAccessConditionId() string`

GetAccessConditionId returns the AccessConditionId field if non-nil, zero value otherwise.

### GetAccessConditionIdOk

`func (o *PermissionGrant) GetAccessConditionIdOk() (*string, bool)`

GetAccessConditionIdOk returns a tuple with the AccessConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConditionId

`func (o *PermissionGrant) SetAccessConditionId(v string)`

SetAccessConditionId sets AccessConditionId field to given value.

### HasAccessConditionId

`func (o *PermissionGrant) HasAccessConditionId() bool`

HasAccessConditionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


