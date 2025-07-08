# ResourceTypeToSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreServer** | Pointer to **[]string** | Security group to use for the restore server. If not specified, Eon attempts to use the default security group.  Currently, a single security group is supported.  | [optional] 
**RestoredRdsInstance** | Pointer to **[]string** | Security group to use for restored RDS instances. If not specified, Eon attempts to use the default security group.  Currently, a single security group is supported.  | [optional] 

## Methods

### NewResourceTypeToSecurityGroup

`func NewResourceTypeToSecurityGroup() *ResourceTypeToSecurityGroup`

NewResourceTypeToSecurityGroup instantiates a new ResourceTypeToSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeToSecurityGroupWithDefaults

`func NewResourceTypeToSecurityGroupWithDefaults() *ResourceTypeToSecurityGroup`

NewResourceTypeToSecurityGroupWithDefaults instantiates a new ResourceTypeToSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreServer

`func (o *ResourceTypeToSecurityGroup) GetRestoreServer() []string`

GetRestoreServer returns the RestoreServer field if non-nil, zero value otherwise.

### GetRestoreServerOk

`func (o *ResourceTypeToSecurityGroup) GetRestoreServerOk() (*[]string, bool)`

GetRestoreServerOk returns a tuple with the RestoreServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServer

`func (o *ResourceTypeToSecurityGroup) SetRestoreServer(v []string)`

SetRestoreServer sets RestoreServer field to given value.

### HasRestoreServer

`func (o *ResourceTypeToSecurityGroup) HasRestoreServer() bool`

HasRestoreServer returns a boolean if a field has been set.

### GetRestoredRdsInstance

`func (o *ResourceTypeToSecurityGroup) GetRestoredRdsInstance() []string`

GetRestoredRdsInstance returns the RestoredRdsInstance field if non-nil, zero value otherwise.

### GetRestoredRdsInstanceOk

`func (o *ResourceTypeToSecurityGroup) GetRestoredRdsInstanceOk() (*[]string, bool)`

GetRestoredRdsInstanceOk returns a tuple with the RestoredRdsInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredRdsInstance

`func (o *ResourceTypeToSecurityGroup) SetRestoredRdsInstance(v []string)`

SetRestoredRdsInstance sets RestoredRdsInstance field to given value.

### HasRestoredRdsInstance

`func (o *ResourceTypeToSecurityGroup) HasRestoredRdsInstance() bool`

HasRestoredRdsInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


