# ActionApprovalRuleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableActionApprovalRuleGroupCondition**](ActionApprovalRuleGroupCondition.md) |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceTypeCondition**](ResourceTypeCondition.md) |  | [optional] 
**DataClasses** | Pointer to [**NullableDataClassesCondition**](DataClassesCondition.md) |  | [optional] 
**Environment** | Pointer to [**NullableEnvironmentCondition**](EnvironmentCondition.md) |  | [optional] 
**Apps** | Pointer to [**NullableAppsCondition**](AppsCondition.md) |  | [optional] 
**SensitivityAnnotations** | Pointer to [**NullableSensitivityAnnotationsCondition**](SensitivityAnnotationsCondition.md) |  | [optional] 
**SecurityScanConclusion** | Pointer to [**NullableSecurityScanConclusionCondition**](SecurityScanConclusionCondition.md) |  | [optional] 
**CloudProvider** | Pointer to [**NullableCloudProviderCondition**](CloudProviderCondition.md) |  | [optional] 
**AccountId** | Pointer to [**NullableAccountIdCondition**](AccountIdCondition.md) |  | [optional] 
**SourceRegion** | Pointer to [**NullableRegionCondition**](RegionCondition.md) |  | [optional] 
**Vpc** | Pointer to [**NullableVpcCondition**](VpcCondition.md) |  | [optional] 
**Subnets** | Pointer to [**NullableSubnetsCondition**](SubnetsCondition.md) |  | [optional] 
**ResourceGroupName** | Pointer to [**NullableResourceGroupNameCondition**](ResourceGroupNameCondition.md) |  | [optional] 
**EncryptionType** | Pointer to [**NullableEncryptionTypeCondition**](EncryptionTypeCondition.md) |  | [optional] 
**ResourceName** | Pointer to [**NullableResourceNameCondition**](ResourceNameCondition.md) |  | [optional] 
**ResourceId** | Pointer to [**NullableResourceIdCondition**](ResourceIdCondition.md) |  | [optional] 
**TagKeys** | Pointer to [**NullableTagKeysCondition**](TagKeysCondition.md) |  | [optional] 
**TagKeyValues** | Pointer to [**NullableTagKeyValuesCondition**](TagKeyValuesCondition.md) |  | [optional] 

## Methods

### NewActionApprovalRuleExpression

`func NewActionApprovalRuleExpression() *ActionApprovalRuleExpression`

NewActionApprovalRuleExpression instantiates a new ActionApprovalRuleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionApprovalRuleExpressionWithDefaults

`func NewActionApprovalRuleExpressionWithDefaults() *ActionApprovalRuleExpression`

NewActionApprovalRuleExpressionWithDefaults instantiates a new ActionApprovalRuleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ActionApprovalRuleExpression) GetGroup() ActionApprovalRuleGroupCondition`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ActionApprovalRuleExpression) GetGroupOk() (*ActionApprovalRuleGroupCondition, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ActionApprovalRuleExpression) SetGroup(v ActionApprovalRuleGroupCondition)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ActionApprovalRuleExpression) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *ActionApprovalRuleExpression) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *ActionApprovalRuleExpression) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetResourceType

`func (o *ActionApprovalRuleExpression) GetResourceType() ResourceTypeCondition`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ActionApprovalRuleExpression) GetResourceTypeOk() (*ResourceTypeCondition, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ActionApprovalRuleExpression) SetResourceType(v ResourceTypeCondition)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ActionApprovalRuleExpression) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ActionApprovalRuleExpression) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *ActionApprovalRuleExpression) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetDataClasses

`func (o *ActionApprovalRuleExpression) GetDataClasses() DataClassesCondition`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *ActionApprovalRuleExpression) GetDataClassesOk() (*DataClassesCondition, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *ActionApprovalRuleExpression) SetDataClasses(v DataClassesCondition)`

SetDataClasses sets DataClasses field to given value.

### HasDataClasses

`func (o *ActionApprovalRuleExpression) HasDataClasses() bool`

HasDataClasses returns a boolean if a field has been set.

### SetDataClassesNil

`func (o *ActionApprovalRuleExpression) SetDataClassesNil(b bool)`

 SetDataClassesNil sets the value for DataClasses to be an explicit nil

### UnsetDataClasses
`func (o *ActionApprovalRuleExpression) UnsetDataClasses()`

UnsetDataClasses ensures that no value is present for DataClasses, not even an explicit nil
### GetEnvironment

`func (o *ActionApprovalRuleExpression) GetEnvironment() EnvironmentCondition`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ActionApprovalRuleExpression) GetEnvironmentOk() (*EnvironmentCondition, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ActionApprovalRuleExpression) SetEnvironment(v EnvironmentCondition)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ActionApprovalRuleExpression) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ActionApprovalRuleExpression) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ActionApprovalRuleExpression) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApps

`func (o *ActionApprovalRuleExpression) GetApps() AppsCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ActionApprovalRuleExpression) GetAppsOk() (*AppsCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ActionApprovalRuleExpression) SetApps(v AppsCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ActionApprovalRuleExpression) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetAppsNil

`func (o *ActionApprovalRuleExpression) SetAppsNil(b bool)`

 SetAppsNil sets the value for Apps to be an explicit nil

### UnsetApps
`func (o *ActionApprovalRuleExpression) UnsetApps()`

UnsetApps ensures that no value is present for Apps, not even an explicit nil
### GetSensitivityAnnotations

`func (o *ActionApprovalRuleExpression) GetSensitivityAnnotations() SensitivityAnnotationsCondition`

GetSensitivityAnnotations returns the SensitivityAnnotations field if non-nil, zero value otherwise.

### GetSensitivityAnnotationsOk

`func (o *ActionApprovalRuleExpression) GetSensitivityAnnotationsOk() (*SensitivityAnnotationsCondition, bool)`

GetSensitivityAnnotationsOk returns a tuple with the SensitivityAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivityAnnotations

`func (o *ActionApprovalRuleExpression) SetSensitivityAnnotations(v SensitivityAnnotationsCondition)`

SetSensitivityAnnotations sets SensitivityAnnotations field to given value.

### HasSensitivityAnnotations

`func (o *ActionApprovalRuleExpression) HasSensitivityAnnotations() bool`

HasSensitivityAnnotations returns a boolean if a field has been set.

### SetSensitivityAnnotationsNil

`func (o *ActionApprovalRuleExpression) SetSensitivityAnnotationsNil(b bool)`

 SetSensitivityAnnotationsNil sets the value for SensitivityAnnotations to be an explicit nil

### UnsetSensitivityAnnotations
`func (o *ActionApprovalRuleExpression) UnsetSensitivityAnnotations()`

UnsetSensitivityAnnotations ensures that no value is present for SensitivityAnnotations, not even an explicit nil
### GetSecurityScanConclusion

`func (o *ActionApprovalRuleExpression) GetSecurityScanConclusion() SecurityScanConclusionCondition`

GetSecurityScanConclusion returns the SecurityScanConclusion field if non-nil, zero value otherwise.

### GetSecurityScanConclusionOk

`func (o *ActionApprovalRuleExpression) GetSecurityScanConclusionOk() (*SecurityScanConclusionCondition, bool)`

GetSecurityScanConclusionOk returns a tuple with the SecurityScanConclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanConclusion

`func (o *ActionApprovalRuleExpression) SetSecurityScanConclusion(v SecurityScanConclusionCondition)`

SetSecurityScanConclusion sets SecurityScanConclusion field to given value.

### HasSecurityScanConclusion

`func (o *ActionApprovalRuleExpression) HasSecurityScanConclusion() bool`

HasSecurityScanConclusion returns a boolean if a field has been set.

### SetSecurityScanConclusionNil

`func (o *ActionApprovalRuleExpression) SetSecurityScanConclusionNil(b bool)`

 SetSecurityScanConclusionNil sets the value for SecurityScanConclusion to be an explicit nil

### UnsetSecurityScanConclusion
`func (o *ActionApprovalRuleExpression) UnsetSecurityScanConclusion()`

UnsetSecurityScanConclusion ensures that no value is present for SecurityScanConclusion, not even an explicit nil
### GetCloudProvider

`func (o *ActionApprovalRuleExpression) GetCloudProvider() CloudProviderCondition`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ActionApprovalRuleExpression) GetCloudProviderOk() (*CloudProviderCondition, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ActionApprovalRuleExpression) SetCloudProvider(v CloudProviderCondition)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ActionApprovalRuleExpression) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProviderNil

`func (o *ActionApprovalRuleExpression) SetCloudProviderNil(b bool)`

 SetCloudProviderNil sets the value for CloudProvider to be an explicit nil

### UnsetCloudProvider
`func (o *ActionApprovalRuleExpression) UnsetCloudProvider()`

UnsetCloudProvider ensures that no value is present for CloudProvider, not even an explicit nil
### GetAccountId

`func (o *ActionApprovalRuleExpression) GetAccountId() AccountIdCondition`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActionApprovalRuleExpression) GetAccountIdOk() (*AccountIdCondition, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActionApprovalRuleExpression) SetAccountId(v AccountIdCondition)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActionApprovalRuleExpression) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *ActionApprovalRuleExpression) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ActionApprovalRuleExpression) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetSourceRegion

`func (o *ActionApprovalRuleExpression) GetSourceRegion() RegionCondition`

GetSourceRegion returns the SourceRegion field if non-nil, zero value otherwise.

### GetSourceRegionOk

`func (o *ActionApprovalRuleExpression) GetSourceRegionOk() (*RegionCondition, bool)`

GetSourceRegionOk returns a tuple with the SourceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegion

`func (o *ActionApprovalRuleExpression) SetSourceRegion(v RegionCondition)`

SetSourceRegion sets SourceRegion field to given value.

### HasSourceRegion

`func (o *ActionApprovalRuleExpression) HasSourceRegion() bool`

HasSourceRegion returns a boolean if a field has been set.

### SetSourceRegionNil

`func (o *ActionApprovalRuleExpression) SetSourceRegionNil(b bool)`

 SetSourceRegionNil sets the value for SourceRegion to be an explicit nil

### UnsetSourceRegion
`func (o *ActionApprovalRuleExpression) UnsetSourceRegion()`

UnsetSourceRegion ensures that no value is present for SourceRegion, not even an explicit nil
### GetVpc

`func (o *ActionApprovalRuleExpression) GetVpc() VpcCondition`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ActionApprovalRuleExpression) GetVpcOk() (*VpcCondition, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ActionApprovalRuleExpression) SetVpc(v VpcCondition)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *ActionApprovalRuleExpression) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *ActionApprovalRuleExpression) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *ActionApprovalRuleExpression) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
### GetSubnets

`func (o *ActionApprovalRuleExpression) GetSubnets() SubnetsCondition`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ActionApprovalRuleExpression) GetSubnetsOk() (*SubnetsCondition, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ActionApprovalRuleExpression) SetSubnets(v SubnetsCondition)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *ActionApprovalRuleExpression) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *ActionApprovalRuleExpression) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *ActionApprovalRuleExpression) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetResourceGroupName

`func (o *ActionApprovalRuleExpression) GetResourceGroupName() ResourceGroupNameCondition`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ActionApprovalRuleExpression) GetResourceGroupNameOk() (*ResourceGroupNameCondition, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ActionApprovalRuleExpression) SetResourceGroupName(v ResourceGroupNameCondition)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ActionApprovalRuleExpression) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ActionApprovalRuleExpression) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ActionApprovalRuleExpression) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetEncryptionType

`func (o *ActionApprovalRuleExpression) GetEncryptionType() EncryptionTypeCondition`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *ActionApprovalRuleExpression) GetEncryptionTypeOk() (*EncryptionTypeCondition, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *ActionApprovalRuleExpression) SetEncryptionType(v EncryptionTypeCondition)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *ActionApprovalRuleExpression) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### SetEncryptionTypeNil

`func (o *ActionApprovalRuleExpression) SetEncryptionTypeNil(b bool)`

 SetEncryptionTypeNil sets the value for EncryptionType to be an explicit nil

### UnsetEncryptionType
`func (o *ActionApprovalRuleExpression) UnsetEncryptionType()`

UnsetEncryptionType ensures that no value is present for EncryptionType, not even an explicit nil
### GetResourceName

`func (o *ActionApprovalRuleExpression) GetResourceName() ResourceNameCondition`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ActionApprovalRuleExpression) GetResourceNameOk() (*ResourceNameCondition, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ActionApprovalRuleExpression) SetResourceName(v ResourceNameCondition)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ActionApprovalRuleExpression) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ActionApprovalRuleExpression) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ActionApprovalRuleExpression) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceId

`func (o *ActionApprovalRuleExpression) GetResourceId() ResourceIdCondition`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ActionApprovalRuleExpression) GetResourceIdOk() (*ResourceIdCondition, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ActionApprovalRuleExpression) SetResourceId(v ResourceIdCondition)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ActionApprovalRuleExpression) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ActionApprovalRuleExpression) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ActionApprovalRuleExpression) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetTagKeys

`func (o *ActionApprovalRuleExpression) GetTagKeys() TagKeysCondition`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *ActionApprovalRuleExpression) GetTagKeysOk() (*TagKeysCondition, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *ActionApprovalRuleExpression) SetTagKeys(v TagKeysCondition)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *ActionApprovalRuleExpression) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeysNil

`func (o *ActionApprovalRuleExpression) SetTagKeysNil(b bool)`

 SetTagKeysNil sets the value for TagKeys to be an explicit nil

### UnsetTagKeys
`func (o *ActionApprovalRuleExpression) UnsetTagKeys()`

UnsetTagKeys ensures that no value is present for TagKeys, not even an explicit nil
### GetTagKeyValues

`func (o *ActionApprovalRuleExpression) GetTagKeyValues() TagKeyValuesCondition`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *ActionApprovalRuleExpression) GetTagKeyValuesOk() (*TagKeyValuesCondition, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *ActionApprovalRuleExpression) SetTagKeyValues(v TagKeyValuesCondition)`

SetTagKeyValues sets TagKeyValues field to given value.

### HasTagKeyValues

`func (o *ActionApprovalRuleExpression) HasTagKeyValues() bool`

HasTagKeyValues returns a boolean if a field has been set.

### SetTagKeyValuesNil

`func (o *ActionApprovalRuleExpression) SetTagKeyValuesNil(b bool)`

 SetTagKeyValuesNil sets the value for TagKeyValues to be an explicit nil

### UnsetTagKeyValues
`func (o *ActionApprovalRuleExpression) UnsetTagKeyValues()`

UnsetTagKeyValues ensures that no value is present for TagKeyValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


