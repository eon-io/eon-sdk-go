# AzureSqlDatabaseRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region to restore to. | 
**ResourceGroupName** | **string** | Name of the resource group to restore to. | 
**ServerName** | **string** | Restored SQL server name. | 
**AdminUserName** | **string** | Username to set for admin of user of the restored database. | 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings. If not provided, defaults to an empty object, with no tags applied.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 

## Methods

### NewAzureSqlDatabaseRestoreTarget

`func NewAzureSqlDatabaseRestoreTarget(region string, resourceGroupName string, serverName string, adminUserName string, ) *AzureSqlDatabaseRestoreTarget`

NewAzureSqlDatabaseRestoreTarget instantiates a new AzureSqlDatabaseRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlDatabaseRestoreTargetWithDefaults

`func NewAzureSqlDatabaseRestoreTargetWithDefaults() *AzureSqlDatabaseRestoreTarget`

NewAzureSqlDatabaseRestoreTargetWithDefaults instantiates a new AzureSqlDatabaseRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AzureSqlDatabaseRestoreTarget) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureSqlDatabaseRestoreTarget) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureSqlDatabaseRestoreTarget) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceGroupName

`func (o *AzureSqlDatabaseRestoreTarget) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureSqlDatabaseRestoreTarget) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureSqlDatabaseRestoreTarget) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetServerName

`func (o *AzureSqlDatabaseRestoreTarget) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AzureSqlDatabaseRestoreTarget) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AzureSqlDatabaseRestoreTarget) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetAdminUserName

`func (o *AzureSqlDatabaseRestoreTarget) GetAdminUserName() string`

GetAdminUserName returns the AdminUserName field if non-nil, zero value otherwise.

### GetAdminUserNameOk

`func (o *AzureSqlDatabaseRestoreTarget) GetAdminUserNameOk() (*string, bool)`

GetAdminUserNameOk returns a tuple with the AdminUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserName

`func (o *AzureSqlDatabaseRestoreTarget) SetAdminUserName(v string)`

SetAdminUserName sets AdminUserName field to given value.


### GetTags

`func (o *AzureSqlDatabaseRestoreTarget) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureSqlDatabaseRestoreTarget) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureSqlDatabaseRestoreTarget) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureSqlDatabaseRestoreTarget) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


