# CostDataFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to [**CloudAccountIdFilters**](CloudAccountIdFilters.md) |  | [optional] 
**SourceAccountProviderId** | Pointer to [**ProviderAccountIdFilters**](ProviderAccountIdFilters.md) |  | [optional] 
**CloudProvider** | Pointer to [**CloudProviderFilters**](CloudProviderFilters.md) |  | [optional] 
**ResourceType** | Pointer to [**CostDataResourceTypeFilters**](CostDataResourceTypeFilters.md) |  | [optional] 
**ResourceId** | Pointer to [**IdFilters**](IdFilters.md) |  | [optional] 

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

### GetAccountId

`func (o *CostDataFilters) GetAccountId() CloudAccountIdFilters`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CostDataFilters) GetAccountIdOk() (*CloudAccountIdFilters, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CostDataFilters) SetAccountId(v CloudAccountIdFilters)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CostDataFilters) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSourceAccountProviderId

`func (o *CostDataFilters) GetSourceAccountProviderId() ProviderAccountIdFilters`

GetSourceAccountProviderId returns the SourceAccountProviderId field if non-nil, zero value otherwise.

### GetSourceAccountProviderIdOk

`func (o *CostDataFilters) GetSourceAccountProviderIdOk() (*ProviderAccountIdFilters, bool)`

GetSourceAccountProviderIdOk returns a tuple with the SourceAccountProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountProviderId

`func (o *CostDataFilters) SetSourceAccountProviderId(v ProviderAccountIdFilters)`

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

`func (o *CostDataFilters) GetResourceType() CostDataResourceTypeFilters`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CostDataFilters) GetResourceTypeOk() (*CostDataResourceTypeFilters, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CostDataFilters) SetResourceType(v CostDataResourceTypeFilters)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


