# InventoryFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**IdFilters**](IdFilters.md) |  | [optional] 
**ProviderResourceId** | Pointer to [**ResourceIdFilters**](ResourceIdFilters.md) |  | [optional] 
**ResourceName** | Pointer to [**ResourceNameFilters**](ResourceNameFilters.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceTypeFilters**](ResourceTypeFilters.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentFilters**](EnvironmentFilters.md) |  | [optional] 
**AccountId** | Pointer to [**AccountIdFilters**](AccountIdFilters.md) |  | [optional] 
**Apps** | Pointer to [**AppFilters**](AppFilters.md) |  | [optional] 
**Subnets** | Pointer to [**SubnetFilters**](SubnetFilters.md) |  | [optional] 
**DataClasses** | Pointer to [**DataClassesFilters**](DataClassesFilters.md) |  | [optional] 
**BackupStatus** | Pointer to [**BackupStatusFilters**](BackupStatusFilters.md) |  | [optional] 
**TagKeys** | Pointer to [**TagKeysFilters**](TagKeysFilters.md) |  | [optional] 
**TagKeyValues** | Pointer to [**TagKeyValuesFilters**](TagKeyValuesFilters.md) |  | [optional] 

## Methods

### NewInventoryFilterConditions

`func NewInventoryFilterConditions() *InventoryFilterConditions`

NewInventoryFilterConditions instantiates a new InventoryFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryFilterConditionsWithDefaults

`func NewInventoryFilterConditionsWithDefaults() *InventoryFilterConditions`

NewInventoryFilterConditionsWithDefaults instantiates a new InventoryFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryFilterConditions) GetId() IdFilters`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryFilterConditions) GetIdOk() (*IdFilters, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryFilterConditions) SetId(v IdFilters)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryFilterConditions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderResourceId

`func (o *InventoryFilterConditions) GetProviderResourceId() ResourceIdFilters`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *InventoryFilterConditions) GetProviderResourceIdOk() (*ResourceIdFilters, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *InventoryFilterConditions) SetProviderResourceId(v ResourceIdFilters)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *InventoryFilterConditions) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *InventoryFilterConditions) GetResourceName() ResourceNameFilters`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InventoryFilterConditions) GetResourceNameOk() (*ResourceNameFilters, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InventoryFilterConditions) SetResourceName(v ResourceNameFilters)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *InventoryFilterConditions) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *InventoryFilterConditions) GetResourceType() ResourceTypeFilters`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *InventoryFilterConditions) GetResourceTypeOk() (*ResourceTypeFilters, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *InventoryFilterConditions) SetResourceType(v ResourceTypeFilters)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *InventoryFilterConditions) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetEnvironment

`func (o *InventoryFilterConditions) GetEnvironment() EnvironmentFilters`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InventoryFilterConditions) GetEnvironmentOk() (*EnvironmentFilters, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InventoryFilterConditions) SetEnvironment(v EnvironmentFilters)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InventoryFilterConditions) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAccountId

`func (o *InventoryFilterConditions) GetAccountId() AccountIdFilters`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InventoryFilterConditions) GetAccountIdOk() (*AccountIdFilters, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InventoryFilterConditions) SetAccountId(v AccountIdFilters)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InventoryFilterConditions) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApps

`func (o *InventoryFilterConditions) GetApps() AppFilters`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *InventoryFilterConditions) GetAppsOk() (*AppFilters, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *InventoryFilterConditions) SetApps(v AppFilters)`

SetApps sets Apps field to given value.

### HasApps

`func (o *InventoryFilterConditions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetSubnets

`func (o *InventoryFilterConditions) GetSubnets() SubnetFilters`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InventoryFilterConditions) GetSubnetsOk() (*SubnetFilters, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InventoryFilterConditions) SetSubnets(v SubnetFilters)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InventoryFilterConditions) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetDataClasses

`func (o *InventoryFilterConditions) GetDataClasses() DataClassesFilters`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *InventoryFilterConditions) GetDataClassesOk() (*DataClassesFilters, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *InventoryFilterConditions) SetDataClasses(v DataClassesFilters)`

SetDataClasses sets DataClasses field to given value.

### HasDataClasses

`func (o *InventoryFilterConditions) HasDataClasses() bool`

HasDataClasses returns a boolean if a field has been set.

### GetBackupStatus

`func (o *InventoryFilterConditions) GetBackupStatus() BackupStatusFilters`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *InventoryFilterConditions) GetBackupStatusOk() (*BackupStatusFilters, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *InventoryFilterConditions) SetBackupStatus(v BackupStatusFilters)`

SetBackupStatus sets BackupStatus field to given value.

### HasBackupStatus

`func (o *InventoryFilterConditions) HasBackupStatus() bool`

HasBackupStatus returns a boolean if a field has been set.

### GetTagKeys

`func (o *InventoryFilterConditions) GetTagKeys() TagKeysFilters`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *InventoryFilterConditions) GetTagKeysOk() (*TagKeysFilters, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *InventoryFilterConditions) SetTagKeys(v TagKeysFilters)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *InventoryFilterConditions) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagKeyValues

`func (o *InventoryFilterConditions) GetTagKeyValues() TagKeyValuesFilters`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *InventoryFilterConditions) GetTagKeyValuesOk() (*TagKeyValuesFilters, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *InventoryFilterConditions) SetTagKeyValues(v TagKeyValuesFilters)`

SetTagKeyValues sets TagKeyValues field to given value.

### HasTagKeyValues

`func (o *InventoryFilterConditions) HasTagKeyValues() bool`

HasTagKeyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


