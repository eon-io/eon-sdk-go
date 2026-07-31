# BackupPostureControlExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableBackupPostureControlGroupCondition**](BackupPostureControlGroupCondition.md) |  | [optional] 
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

### NewBackupPostureControlExpression

`func NewBackupPostureControlExpression() *BackupPostureControlExpression`

NewBackupPostureControlExpression instantiates a new BackupPostureControlExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPostureControlExpressionWithDefaults

`func NewBackupPostureControlExpressionWithDefaults() *BackupPostureControlExpression`

NewBackupPostureControlExpressionWithDefaults instantiates a new BackupPostureControlExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *BackupPostureControlExpression) GetGroup() BackupPostureControlGroupCondition`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BackupPostureControlExpression) GetGroupOk() (*BackupPostureControlGroupCondition, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BackupPostureControlExpression) SetGroup(v BackupPostureControlGroupCondition)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BackupPostureControlExpression) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BackupPostureControlExpression) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BackupPostureControlExpression) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetResourceType

`func (o *BackupPostureControlExpression) GetResourceType() ResourceTypeCondition`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BackupPostureControlExpression) GetResourceTypeOk() (*ResourceTypeCondition, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BackupPostureControlExpression) SetResourceType(v ResourceTypeCondition)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BackupPostureControlExpression) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BackupPostureControlExpression) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BackupPostureControlExpression) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetDataClasses

`func (o *BackupPostureControlExpression) GetDataClasses() DataClassesCondition`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *BackupPostureControlExpression) GetDataClassesOk() (*DataClassesCondition, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *BackupPostureControlExpression) SetDataClasses(v DataClassesCondition)`

SetDataClasses sets DataClasses field to given value.

### HasDataClasses

`func (o *BackupPostureControlExpression) HasDataClasses() bool`

HasDataClasses returns a boolean if a field has been set.

### SetDataClassesNil

`func (o *BackupPostureControlExpression) SetDataClassesNil(b bool)`

 SetDataClassesNil sets the value for DataClasses to be an explicit nil

### UnsetDataClasses
`func (o *BackupPostureControlExpression) UnsetDataClasses()`

UnsetDataClasses ensures that no value is present for DataClasses, not even an explicit nil
### GetEnvironment

`func (o *BackupPostureControlExpression) GetEnvironment() EnvironmentCondition`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BackupPostureControlExpression) GetEnvironmentOk() (*EnvironmentCondition, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BackupPostureControlExpression) SetEnvironment(v EnvironmentCondition)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BackupPostureControlExpression) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BackupPostureControlExpression) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BackupPostureControlExpression) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApps

`func (o *BackupPostureControlExpression) GetApps() AppsCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BackupPostureControlExpression) GetAppsOk() (*AppsCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BackupPostureControlExpression) SetApps(v AppsCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *BackupPostureControlExpression) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *BackupPostureControlExpression) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *BackupPostureControlExpression) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetCloudProvider

`func (o *BackupPostureControlExpression) GetCloudProvider() CloudProviderCondition`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *BackupPostureControlExpression) GetCloudProviderOk() (*CloudProviderCondition, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *BackupPostureControlExpression) SetCloudProvider(v CloudProviderCondition)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *BackupPostureControlExpression) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProviderNil

`func (o *BackupPostureControlExpression) SetCloudProviderNil(b bool)`

 SetCloudProviderNil sets the value for CloudProvider to be an explicit nil

### UnsetCloudProvider
`func (o *BackupPostureControlExpression) UnsetCloudProvider()`

UnsetCloudProvider ensures that no value is present for CloudProvider, not even an explicit nil
### GetAccountId

`func (o *BackupPostureControlExpression) GetAccountId() AccountIdCondition`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BackupPostureControlExpression) GetAccountIdOk() (*AccountIdCondition, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BackupPostureControlExpression) SetAccountId(v AccountIdCondition)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BackupPostureControlExpression) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *BackupPostureControlExpression) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BackupPostureControlExpression) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetSourceRegion

`func (o *BackupPostureControlExpression) GetSourceRegion() RegionCondition`

GetSourceRegion returns the SourceRegion field if non-nil, zero value otherwise.

### GetSourceRegionOk

`func (o *BackupPostureControlExpression) GetSourceRegionOk() (*RegionCondition, bool)`

GetSourceRegionOk returns a tuple with the SourceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegion

`func (o *BackupPostureControlExpression) SetSourceRegion(v RegionCondition)`

SetSourceRegion sets SourceRegion field to given value.

### HasSourceRegion

`func (o *BackupPostureControlExpression) HasSourceRegion() bool`

HasSourceRegion returns a boolean if a field has been set.

### SetSourceRegionNil

`func (o *BackupPostureControlExpression) SetSourceRegionNil(b bool)`

 SetSourceRegionNil sets the value for SourceRegion to be an explicit nil

### UnsetSourceRegion
`func (o *BackupPostureControlExpression) UnsetSourceRegion()`

UnsetSourceRegion ensures that no value is present for SourceRegion, not even an explicit nil
### GetVpc

`func (o *BackupPostureControlExpression) GetVpc() VpcCondition`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *BackupPostureControlExpression) GetVpcOk() (*VpcCondition, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *BackupPostureControlExpression) SetVpc(v VpcCondition)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *BackupPostureControlExpression) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *BackupPostureControlExpression) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *BackupPostureControlExpression) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
### GetSubnets

`func (o *BackupPostureControlExpression) GetSubnets() SubnetsCondition`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BackupPostureControlExpression) GetSubnetsOk() (*SubnetsCondition, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BackupPostureControlExpression) SetSubnets(v SubnetsCondition)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *BackupPostureControlExpression) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *BackupPostureControlExpression) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *BackupPostureControlExpression) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetResourceGroupName

`func (o *BackupPostureControlExpression) GetResourceGroupName() ResourceGroupNameCondition`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *BackupPostureControlExpression) GetResourceGroupNameOk() (*ResourceGroupNameCondition, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *BackupPostureControlExpression) SetResourceGroupName(v ResourceGroupNameCondition)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *BackupPostureControlExpression) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *BackupPostureControlExpression) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *BackupPostureControlExpression) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetResourceName

`func (o *BackupPostureControlExpression) GetResourceName() ResourceNameCondition`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *BackupPostureControlExpression) GetResourceNameOk() (*ResourceNameCondition, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *BackupPostureControlExpression) SetResourceName(v ResourceNameCondition)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *BackupPostureControlExpression) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *BackupPostureControlExpression) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *BackupPostureControlExpression) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceId

`func (o *BackupPostureControlExpression) GetResourceId() ResourceIdCondition`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BackupPostureControlExpression) GetResourceIdOk() (*ResourceIdCondition, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BackupPostureControlExpression) SetResourceId(v ResourceIdCondition)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BackupPostureControlExpression) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BackupPostureControlExpression) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BackupPostureControlExpression) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetTagKeys

`func (o *BackupPostureControlExpression) GetTagKeys() TagKeysCondition`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *BackupPostureControlExpression) GetTagKeysOk() (*TagKeysCondition, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *BackupPostureControlExpression) SetTagKeys(v TagKeysCondition)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *BackupPostureControlExpression) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeysNil

`func (o *BackupPostureControlExpression) SetTagKeysNil(b bool)`

 SetTagKeysNil sets the value for TagKeys to be an explicit nil

### UnsetTagKeys
`func (o *BackupPostureControlExpression) UnsetTagKeys()`

UnsetTagKeys ensures that no value is present for TagKeys, not even an explicit nil
### GetTagKeyValues

`func (o *BackupPostureControlExpression) GetTagKeyValues() TagKeyValuesCondition`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *BackupPostureControlExpression) GetTagKeyValuesOk() (*TagKeyValuesCondition, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *BackupPostureControlExpression) SetTagKeyValues(v TagKeyValuesCondition)`

SetTagKeyValues sets TagKeyValues field to given value.

### HasTagKeyValues

`func (o *BackupPostureControlExpression) HasTagKeyValues() bool`

HasTagKeyValues returns a boolean if a field has been set.

### SetTagKeyValuesNil

`func (o *BackupPostureControlExpression) SetTagKeyValuesNil(b bool)`

 SetTagKeyValuesNil sets the value for TagKeyValues to be an explicit nil

### UnsetTagKeyValues
`func (o *BackupPostureControlExpression) UnsetTagKeyValues()`

UnsetTagKeyValues ensures that no value is present for TagKeyValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


