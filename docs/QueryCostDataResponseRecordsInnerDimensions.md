# QueryCostDataResponseRecordsInnerDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Eon-assigned resource ID | [optional] 
**SourceAccountId** | Pointer to **string** | Eon-assigned source account ID | [optional] 
**SourceAccountProviderId** | Pointer to **string** | Source account provider id | [optional] 
**SourceAccountName** | Pointer to **string** | Source account alias | [optional] 
**SourceStorageSize** | Pointer to **int64** | Source storage size in bytes | [optional] 
**ProviderResourceId** | Pointer to **string** | Cloud provider resource ID | [optional] 
**ResourceName** | Pointer to **string** | Cloud provider resource name | [optional] 
**ResourceTags** | Pointer to **map[string]string** | Cloud provider resource tags | [optional] 
**ResourceSourceRegion** | Pointer to **string** | Cloud provider resource region | [optional] 

## Methods

### NewQueryCostDataResponseRecordsInnerDimensions

`func NewQueryCostDataResponseRecordsInnerDimensions() *QueryCostDataResponseRecordsInnerDimensions`

NewQueryCostDataResponseRecordsInnerDimensions instantiates a new QueryCostDataResponseRecordsInnerDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCostDataResponseRecordsInnerDimensionsWithDefaults

`func NewQueryCostDataResponseRecordsInnerDimensionsWithDefaults() *QueryCostDataResponseRecordsInnerDimensions`

NewQueryCostDataResponseRecordsInnerDimensionsWithDefaults instantiates a new QueryCostDataResponseRecordsInnerDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetResourceType

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetSourceAccountProviderId

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountProviderId() string`

GetSourceAccountProviderId returns the SourceAccountProviderId field if non-nil, zero value otherwise.

### GetSourceAccountProviderIdOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountProviderIdOk() (*string, bool)`

GetSourceAccountProviderIdOk returns a tuple with the SourceAccountProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountProviderId

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetSourceAccountProviderId(v string)`

SetSourceAccountProviderId sets SourceAccountProviderId field to given value.

### HasSourceAccountProviderId

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasSourceAccountProviderId() bool`

HasSourceAccountProviderId returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceStorageSize

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceStorageSize() int64`

GetSourceStorageSize returns the SourceStorageSize field if non-nil, zero value otherwise.

### GetSourceStorageSizeOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetSourceStorageSizeOk() (*int64, bool)`

GetSourceStorageSizeOk returns a tuple with the SourceStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStorageSize

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetSourceStorageSize(v int64)`

SetSourceStorageSize sets SourceStorageSize field to given value.

### HasSourceStorageSize

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasSourceStorageSize() bool`

HasSourceStorageSize returns a boolean if a field has been set.

### GetProviderResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### GetResourceName

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceTags

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceTags() map[string]string`

GetResourceTags returns the ResourceTags field if non-nil, zero value otherwise.

### GetResourceTagsOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceTagsOk() (*map[string]string, bool)`

GetResourceTagsOk returns a tuple with the ResourceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTags

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetResourceTags(v map[string]string)`

SetResourceTags sets ResourceTags field to given value.

### HasResourceTags

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasResourceTags() bool`

HasResourceTags returns a boolean if a field has been set.

### GetResourceSourceRegion

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceSourceRegion() string`

GetResourceSourceRegion returns the ResourceSourceRegion field if non-nil, zero value otherwise.

### GetResourceSourceRegionOk

`func (o *QueryCostDataResponseRecordsInnerDimensions) GetResourceSourceRegionOk() (*string, bool)`

GetResourceSourceRegionOk returns a tuple with the ResourceSourceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSourceRegion

`func (o *QueryCostDataResponseRecordsInnerDimensions) SetResourceSourceRegion(v string)`

SetResourceSourceRegion sets ResourceSourceRegion field to given value.

### HasResourceSourceRegion

`func (o *QueryCostDataResponseRecordsInnerDimensions) HasResourceSourceRegion() bool`

HasResourceSourceRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


