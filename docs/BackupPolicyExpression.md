# BackupPolicyExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableBackupPolicyGroupCondition**](BackupPolicyGroupCondition.md) |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceTypeCondition**](ResourceTypeCondition.md) |  | [optional] 
**DataClasses** | Pointer to [**NullableDataClassesCondition**](DataClassesCondition.md) |  | [optional] 
**Environment** | Pointer to [**NullableEnvironmentCondition**](EnvironmentCondition.md) |  | [optional] 
**Apps** | Pointer to [**NullableAppsCondition**](AppsCondition.md) |  | [optional] 
**CloudProvider** | Pointer to [**NullableCloudProviderCondition**](CloudProviderCondition.md) |  | [optional] 
**AccountId** | Pointer to [**NullableAccountIdCondition**](AccountIdCondition.md) |  | [optional] 
**SourceRegion** | Pointer to [**NullableRegionCondition**](RegionCondition.md) |  | [optional] 
**Vpc** | Pointer to [**NullableVpcCondition**](VpcCondition.md) |  | [optional] 
**Subnets** | Pointer to [**NullableSubnetsCondition**](SubnetsCondition.md) |  | [optional] 
**ResourceGroupName** | Pointer to [**NullableResourceGroupNameCondition**](ResourceGroupNameCondition.md) |  | [optional] 
**ResourceName** | Pointer to [**NullableResourceNameCondition**](ResourceNameCondition.md) |  | [optional] 
**ResourceId** | Pointer to [**NullableResourceIdCondition**](ResourceIdCondition.md) |  | [optional] 
**TagKeys** | Pointer to [**NullableTagKeysCondition**](TagKeysCondition.md) |  | [optional] 
**TagKeyValues** | Pointer to [**NullableTagKeyValuesCondition**](TagKeyValuesCondition.md) |  | [optional] 

## Methods

### NewBackupPolicyExpression

`func NewBackupPolicyExpression() *BackupPolicyExpression`

NewBackupPolicyExpression instantiates a new BackupPolicyExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyExpressionWithDefaults

`func NewBackupPolicyExpressionWithDefaults() *BackupPolicyExpression`

NewBackupPolicyExpressionWithDefaults instantiates a new BackupPolicyExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *BackupPolicyExpression) GetGroup() BackupPolicyGroupCondition`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BackupPolicyExpression) GetGroupOk() (*BackupPolicyGroupCondition, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BackupPolicyExpression) SetGroup(v BackupPolicyGroupCondition)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BackupPolicyExpression) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BackupPolicyExpression) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BackupPolicyExpression) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetResourceType

`func (o *BackupPolicyExpression) GetResourceType() ResourceTypeCondition`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BackupPolicyExpression) GetResourceTypeOk() (*ResourceTypeCondition, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BackupPolicyExpression) SetResourceType(v ResourceTypeCondition)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BackupPolicyExpression) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BackupPolicyExpression) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BackupPolicyExpression) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetDataClasses

`func (o *BackupPolicyExpression) GetDataClasses() DataClassesCondition`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *BackupPolicyExpression) GetDataClassesOk() (*DataClassesCondition, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *BackupPolicyExpression) SetDataClasses(v DataClassesCondition)`

SetDataClasses sets DataClasses field to given value.

### HasDataClasses

`func (o *BackupPolicyExpression) HasDataClasses() bool`

HasDataClasses returns a boolean if a field has been set.

### SetDataClassesNil

`func (o *BackupPolicyExpression) SetDataClassesNil(b bool)`

 SetDataClassesNil sets the value for DataClasses to be an explicit nil

### UnsetDataClasses
`func (o *BackupPolicyExpression) UnsetDataClasses()`

UnsetDataClasses ensures that no value is present for DataClasses, not even an explicit nil
### GetEnvironment

`func (o *BackupPolicyExpression) GetEnvironment() EnvironmentCondition`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BackupPolicyExpression) GetEnvironmentOk() (*EnvironmentCondition, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BackupPolicyExpression) SetEnvironment(v EnvironmentCondition)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BackupPolicyExpression) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BackupPolicyExpression) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BackupPolicyExpression) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApps

`func (o *BackupPolicyExpression) GetApps() AppsCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BackupPolicyExpression) GetAppsOk() (*AppsCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BackupPolicyExpression) SetApps(v AppsCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *BackupPolicyExpression) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *BackupPolicyExpression) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *BackupPolicyExpression) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetCloudProvider

`func (o *BackupPolicyExpression) GetCloudProvider() CloudProviderCondition`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *BackupPolicyExpression) GetCloudProviderOk() (*CloudProviderCondition, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *BackupPolicyExpression) SetCloudProvider(v CloudProviderCondition)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *BackupPolicyExpression) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProviderNil

`func (o *BackupPolicyExpression) SetCloudProviderNil(b bool)`

 SetCloudProviderNil sets the value for CloudProvider to be an explicit nil

### UnsetCloudProvider
`func (o *BackupPolicyExpression) UnsetCloudProvider()`

UnsetCloudProvider ensures that no value is present for CloudProvider, not even an explicit nil
### GetAccountId

`func (o *BackupPolicyExpression) GetAccountId() AccountIdCondition`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BackupPolicyExpression) GetAccountIdOk() (*AccountIdCondition, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BackupPolicyExpression) SetAccountId(v AccountIdCondition)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BackupPolicyExpression) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *BackupPolicyExpression) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BackupPolicyExpression) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetSourceRegion

`func (o *BackupPolicyExpression) GetSourceRegion() RegionCondition`

GetSourceRegion returns the SourceRegion field if non-nil, zero value otherwise.

### GetSourceRegionOk

`func (o *BackupPolicyExpression) GetSourceRegionOk() (*RegionCondition, bool)`

GetSourceRegionOk returns a tuple with the SourceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegion

`func (o *BackupPolicyExpression) SetSourceRegion(v RegionCondition)`

SetSourceRegion sets SourceRegion field to given value.

### HasSourceRegion

`func (o *BackupPolicyExpression) HasSourceRegion() bool`

HasSourceRegion returns a boolean if a field has been set.

### SetSourceRegionNil

`func (o *BackupPolicyExpression) SetSourceRegionNil(b bool)`

 SetSourceRegionNil sets the value for SourceRegion to be an explicit nil

### UnsetSourceRegion
`func (o *BackupPolicyExpression) UnsetSourceRegion()`

UnsetSourceRegion ensures that no value is present for SourceRegion, not even an explicit nil
### GetVpc

`func (o *BackupPolicyExpression) GetVpc() VpcCondition`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *BackupPolicyExpression) GetVpcOk() (*VpcCondition, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *BackupPolicyExpression) SetVpc(v VpcCondition)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *BackupPolicyExpression) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *BackupPolicyExpression) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *BackupPolicyExpression) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
### GetSubnets

`func (o *BackupPolicyExpression) GetSubnets() SubnetsCondition`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BackupPolicyExpression) GetSubnetsOk() (*SubnetsCondition, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BackupPolicyExpression) SetSubnets(v SubnetsCondition)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *BackupPolicyExpression) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *BackupPolicyExpression) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *BackupPolicyExpression) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetResourceGroupName

`func (o *BackupPolicyExpression) GetResourceGroupName() ResourceGroupNameCondition`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *BackupPolicyExpression) GetResourceGroupNameOk() (*ResourceGroupNameCondition, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *BackupPolicyExpression) SetResourceGroupName(v ResourceGroupNameCondition)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *BackupPolicyExpression) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *BackupPolicyExpression) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *BackupPolicyExpression) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetResourceName

`func (o *BackupPolicyExpression) GetResourceName() ResourceNameCondition`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *BackupPolicyExpression) GetResourceNameOk() (*ResourceNameCondition, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *BackupPolicyExpression) SetResourceName(v ResourceNameCondition)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *BackupPolicyExpression) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *BackupPolicyExpression) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *BackupPolicyExpression) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceId

`func (o *BackupPolicyExpression) GetResourceId() ResourceIdCondition`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BackupPolicyExpression) GetResourceIdOk() (*ResourceIdCondition, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BackupPolicyExpression) SetResourceId(v ResourceIdCondition)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BackupPolicyExpression) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BackupPolicyExpression) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BackupPolicyExpression) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetTagKeys

`func (o *BackupPolicyExpression) GetTagKeys() TagKeysCondition`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *BackupPolicyExpression) GetTagKeysOk() (*TagKeysCondition, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *BackupPolicyExpression) SetTagKeys(v TagKeysCondition)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *BackupPolicyExpression) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeysNil

`func (o *BackupPolicyExpression) SetTagKeysNil(b bool)`

 SetTagKeysNil sets the value for TagKeys to be an explicit nil

### UnsetTagKeys
`func (o *BackupPolicyExpression) UnsetTagKeys()`

UnsetTagKeys ensures that no value is present for TagKeys, not even an explicit nil
### GetTagKeyValues

`func (o *BackupPolicyExpression) GetTagKeyValues() TagKeyValuesCondition`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *BackupPolicyExpression) GetTagKeyValuesOk() (*TagKeyValuesCondition, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *BackupPolicyExpression) SetTagKeyValues(v TagKeyValuesCondition)`

SetTagKeyValues sets TagKeyValues field to given value.

### HasTagKeyValues

`func (o *BackupPolicyExpression) HasTagKeyValues() bool`

HasTagKeyValues returns a boolean if a field has been set.

### SetTagKeyValuesNil

`func (o *BackupPolicyExpression) SetTagKeyValuesNil(b bool)`

 SetTagKeyValuesNil sets the value for TagKeyValues to be an explicit nil

### UnsetTagKeyValues
`func (o *BackupPolicyExpression) UnsetTagKeyValues()`

UnsetTagKeyValues ensures that no value is present for TagKeyValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


