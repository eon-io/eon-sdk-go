# InventoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned resource ID. | 
**CreatedTime** | Pointer to **time.Time** | Date and time the resource record was created in Eon. | [optional] 
**DiscoveredTime** | Pointer to **time.Time** | Date and time the resource was first discovered. | [optional] 
**LatestSnapshotTime** | Pointer to **time.Time** | Date and time of the resource&#39;s latest Eon snapshot. | [optional] 
**OldestSnapshotTime** | Pointer to **time.Time** | Date and time of the resource&#39;s first Eon snapshot. | [optional] 
**BackupStatus** | [**BackupStatus**](BackupStatus.md) |  | 
**ProviderResourceId** | **string** | Cloud-provider-assigned resource ID. | 
**ResourceName** | **string** | Resource display name. | 
**Classifications** | Pointer to [**Classifications**](Classifications.md) |  | [optional] 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**SnapshotStorage** | [**SnapshotStorage**](SnapshotStorage.md) |  | 
**SourceStorage** | [**SourceStorage**](SourceStorage.md) |  | 
**ControlViolationCounts** | Pointer to [**ControlViolations**](ControlViolations.md) |  | [optional] 
**Tags** | **map[string]string** | Resource tags as key-value pairs. Both keys and values are strings. If a tag is a key with no value, the value is presented as an empty string.  **Example:** &#x60;{\&quot;env\&quot;: \&quot;prod\&quot;, \&quot;app\&quot;: \&quot;web\&quot;}&#x60;  | 
**CloudProvider** | [**Provider**](Provider.md) |  | 
**ResourceType** | [**ResourceType**](ResourceType.md) |  | 
**Region** | **string** | Region the resource is hosted in. | 
**Vpc** | Pointer to **string** | VPC the resource is in. | [optional] 
**Subnets** | Pointer to **[]string** | List of subnets the resource belongs to. | [optional] 
**ResourceProperties** | Pointer to [**NullableResourceProperties**](ResourceProperties.md) |  | [optional] 

## Methods

### NewInventoryResource

`func NewInventoryResource(id string, backupStatus BackupStatus, providerResourceId string, resourceName string, providerAccountId string, snapshotStorage SnapshotStorage, sourceStorage SourceStorage, tags map[string]string, cloudProvider Provider, resourceType ResourceType, region string, ) *InventoryResource`

NewInventoryResource instantiates a new InventoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryResourceWithDefaults

`func NewInventoryResourceWithDefaults() *InventoryResource`

NewInventoryResourceWithDefaults instantiates a new InventoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryResource) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedTime

`func (o *InventoryResource) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InventoryResource) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InventoryResource) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *InventoryResource) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDiscoveredTime

`func (o *InventoryResource) GetDiscoveredTime() time.Time`

GetDiscoveredTime returns the DiscoveredTime field if non-nil, zero value otherwise.

### GetDiscoveredTimeOk

`func (o *InventoryResource) GetDiscoveredTimeOk() (*time.Time, bool)`

GetDiscoveredTimeOk returns a tuple with the DiscoveredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredTime

`func (o *InventoryResource) SetDiscoveredTime(v time.Time)`

SetDiscoveredTime sets DiscoveredTime field to given value.

### HasDiscoveredTime

`func (o *InventoryResource) HasDiscoveredTime() bool`

HasDiscoveredTime returns a boolean if a field has been set.

### GetLatestSnapshotTime

`func (o *InventoryResource) GetLatestSnapshotTime() time.Time`

GetLatestSnapshotTime returns the LatestSnapshotTime field if non-nil, zero value otherwise.

### GetLatestSnapshotTimeOk

`func (o *InventoryResource) GetLatestSnapshotTimeOk() (*time.Time, bool)`

GetLatestSnapshotTimeOk returns a tuple with the LatestSnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshotTime

`func (o *InventoryResource) SetLatestSnapshotTime(v time.Time)`

SetLatestSnapshotTime sets LatestSnapshotTime field to given value.

### HasLatestSnapshotTime

`func (o *InventoryResource) HasLatestSnapshotTime() bool`

HasLatestSnapshotTime returns a boolean if a field has been set.

### GetOldestSnapshotTime

`func (o *InventoryResource) GetOldestSnapshotTime() time.Time`

GetOldestSnapshotTime returns the OldestSnapshotTime field if non-nil, zero value otherwise.

### GetOldestSnapshotTimeOk

`func (o *InventoryResource) GetOldestSnapshotTimeOk() (*time.Time, bool)`

GetOldestSnapshotTimeOk returns a tuple with the OldestSnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestSnapshotTime

`func (o *InventoryResource) SetOldestSnapshotTime(v time.Time)`

SetOldestSnapshotTime sets OldestSnapshotTime field to given value.

### HasOldestSnapshotTime

`func (o *InventoryResource) HasOldestSnapshotTime() bool`

HasOldestSnapshotTime returns a boolean if a field has been set.

### GetBackupStatus

`func (o *InventoryResource) GetBackupStatus() BackupStatus`

GetBackupStatus returns the BackupStatus field if non-nil, zero value otherwise.

### GetBackupStatusOk

`func (o *InventoryResource) GetBackupStatusOk() (*BackupStatus, bool)`

GetBackupStatusOk returns a tuple with the BackupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStatus

`func (o *InventoryResource) SetBackupStatus(v BackupStatus)`

SetBackupStatus sets BackupStatus field to given value.


### GetProviderResourceId

`func (o *InventoryResource) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *InventoryResource) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *InventoryResource) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.


### GetResourceName

`func (o *InventoryResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *InventoryResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *InventoryResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetClassifications

`func (o *InventoryResource) GetClassifications() Classifications`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *InventoryResource) GetClassificationsOk() (*Classifications, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *InventoryResource) SetClassifications(v Classifications)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *InventoryResource) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *InventoryResource) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *InventoryResource) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *InventoryResource) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetSnapshotStorage

`func (o *InventoryResource) GetSnapshotStorage() SnapshotStorage`

GetSnapshotStorage returns the SnapshotStorage field if non-nil, zero value otherwise.

### GetSnapshotStorageOk

`func (o *InventoryResource) GetSnapshotStorageOk() (*SnapshotStorage, bool)`

GetSnapshotStorageOk returns a tuple with the SnapshotStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotStorage

`func (o *InventoryResource) SetSnapshotStorage(v SnapshotStorage)`

SetSnapshotStorage sets SnapshotStorage field to given value.


### GetSourceStorage

`func (o *InventoryResource) GetSourceStorage() SourceStorage`

GetSourceStorage returns the SourceStorage field if non-nil, zero value otherwise.

### GetSourceStorageOk

`func (o *InventoryResource) GetSourceStorageOk() (*SourceStorage, bool)`

GetSourceStorageOk returns a tuple with the SourceStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStorage

`func (o *InventoryResource) SetSourceStorage(v SourceStorage)`

SetSourceStorage sets SourceStorage field to given value.


### GetControlViolationCounts

`func (o *InventoryResource) GetControlViolationCounts() ControlViolations`

GetControlViolationCounts returns the ControlViolationCounts field if non-nil, zero value otherwise.

### GetControlViolationCountsOk

`func (o *InventoryResource) GetControlViolationCountsOk() (*ControlViolations, bool)`

GetControlViolationCountsOk returns a tuple with the ControlViolationCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlViolationCounts

`func (o *InventoryResource) SetControlViolationCounts(v ControlViolations)`

SetControlViolationCounts sets ControlViolationCounts field to given value.

### HasControlViolationCounts

`func (o *InventoryResource) HasControlViolationCounts() bool`

HasControlViolationCounts returns a boolean if a field has been set.

### GetTags

`func (o *InventoryResource) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryResource) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryResource) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetCloudProvider

`func (o *InventoryResource) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *InventoryResource) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *InventoryResource) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetResourceType

`func (o *InventoryResource) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *InventoryResource) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *InventoryResource) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetRegion

`func (o *InventoryResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InventoryResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InventoryResource) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVpc

`func (o *InventoryResource) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *InventoryResource) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *InventoryResource) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *InventoryResource) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetSubnets

`func (o *InventoryResource) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InventoryResource) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InventoryResource) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InventoryResource) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetResourceProperties

`func (o *InventoryResource) GetResourceProperties() ResourceProperties`

GetResourceProperties returns the ResourceProperties field if non-nil, zero value otherwise.

### GetResourcePropertiesOk

`func (o *InventoryResource) GetResourcePropertiesOk() (*ResourceProperties, bool)`

GetResourcePropertiesOk returns a tuple with the ResourceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProperties

`func (o *InventoryResource) SetResourceProperties(v ResourceProperties)`

SetResourceProperties sets ResourceProperties field to given value.

### HasResourceProperties

`func (o *InventoryResource) HasResourceProperties() bool`

HasResourceProperties returns a boolean if a field has been set.

### SetResourcePropertiesNil

`func (o *InventoryResource) SetResourcePropertiesNil(b bool)`

 SetResourcePropertiesNil sets the value for ResourceProperties to be an explicit nil

### UnsetResourceProperties
`func (o *InventoryResource) UnsetResourceProperties()`

UnsetResourceProperties ensures that no value is present for ResourceProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


