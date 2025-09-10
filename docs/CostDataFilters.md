# CostDataFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountProviderId** | Pointer to [**SourceAccountProviderIdFilters**](SourceAccountProviderIdFilters.md) |  | [optional] 
**CloudProvider** | Pointer to [**CloudProviderFilters**](CloudProviderFilters.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceTypeFilters**](ResourceTypeFilters.md) |  | [optional] 
**ResourceId** | Pointer to [**IdFilters**](IdFilters.md) |  | [optional] 
**TagKeys** | Pointer to [**TagKeysFilters**](TagKeysFilters.md) |  | [optional] 
**TagKeyValues** | Pointer to [**TagKeyValuesFilters**](TagKeyValuesFilters.md) |  | [optional] 

## Methods

### NewCostDataFilters

`func NewCostDataFilters() *CostDataFilters`

NewCostDataFilters instantiates a new CostDataFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataFiltersWithDefaults

`func NewCostDataFiltersWithDefaults() *CostDataFilters`

NewCostDataFiltersWithDefaults instantiates a new CostDataFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountProviderId

`func (o *CostDataFilters) GetSourceAccountProviderId() SourceAccountProviderIdFilters`

GetSourceAccountProviderId returns the SourceAccountProviderId field if non-nil, zero value otherwise.

### GetSourceAccountProviderIdOk

`func (o *CostDataFilters) GetSourceAccountProviderIdOk() (*SourceAccountProviderIdFilters, bool)`

GetSourceAccountProviderIdOk returns a tuple with the SourceAccountProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountProviderId

`func (o *CostDataFilters) SetSourceAccountProviderId(v SourceAccountProviderIdFilters)`

SetSourceAccountProviderId sets SourceAccountProviderId field to given value.

### HasSourceAccountProviderId

`func (o *CostDataFilters) HasSourceAccountProviderId() bool`

HasSourceAccountProviderId returns a boolean if a field has been set.

### GetCloudProvider

`func (o *CostDataFilters) GetCloudProvider() CloudProviderFilters`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CostDataFilters) GetCloudProviderOk() (*CloudProviderFilters, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CostDataFilters) SetCloudProvider(v CloudProviderFilters)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CostDataFilters) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetResourceType

`func (o *CostDataFilters) GetResourceType() ResourceTypeFilters`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CostDataFilters) GetResourceTypeOk() (*ResourceTypeFilters, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CostDataFilters) SetResourceType(v ResourceTypeFilters)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CostDataFilters) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *CostDataFilters) GetResourceId() IdFilters`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CostDataFilters) GetResourceIdOk() (*IdFilters, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CostDataFilters) SetResourceId(v IdFilters)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CostDataFilters) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetTagKeys

`func (o *CostDataFilters) GetTagKeys() TagKeysFilters`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *CostDataFilters) GetTagKeysOk() (*TagKeysFilters, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeys

`func (o *CostDataFilters) SetTagKeys(v TagKeysFilters)`

SetTagKeys sets TagKeys field to given value.

### HasTagKeys

`func (o *CostDataFilters) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### GetTagKeyValues

`func (o *CostDataFilters) GetTagKeyValues() TagKeyValuesFilters`

GetTagKeyValues returns the TagKeyValues field if non-nil, zero value otherwise.

### GetTagKeyValuesOk

`func (o *CostDataFilters) GetTagKeyValuesOk() (*TagKeyValuesFilters, bool)`

GetTagKeyValuesOk returns a tuple with the TagKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKeyValues

`func (o *CostDataFilters) SetTagKeyValues(v TagKeyValuesFilters)`

SetTagKeyValues sets TagKeyValues field to given value.

### HasTagKeyValues

`func (o *CostDataFilters) HasTagKeyValues() bool`

HasTagKeyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


