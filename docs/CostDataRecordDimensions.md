# CostDataRecordDimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Eon-assigned resource ID | [optional] 
**SourceAccountProviderId** | Pointer to **string** | Source account provider id | [optional] 

## Methods

### NewCostDataRecordDimensions

`func NewCostDataRecordDimensions() *CostDataRecordDimensions`

NewCostDataRecordDimensions instantiates a new CostDataRecordDimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataRecordDimensionsWithDefaults

`func NewCostDataRecordDimensionsWithDefaults() *CostDataRecordDimensions`

NewCostDataRecordDimensionsWithDefaults instantiates a new CostDataRecordDimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CostDataRecordDimensions) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CostDataRecordDimensions) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CostDataRecordDimensions) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CostDataRecordDimensions) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetResourceType

`func (o *CostDataRecordDimensions) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CostDataRecordDimensions) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CostDataRecordDimensions) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CostDataRecordDimensions) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceId

`func (o *CostDataRecordDimensions) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CostDataRecordDimensions) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CostDataRecordDimensions) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CostDataRecordDimensions) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetSourceAccountProviderId

`func (o *CostDataRecordDimensions) GetSourceAccountProviderId() string`

GetSourceAccountProviderId returns the SourceAccountProviderId field if non-nil, zero value otherwise.

### GetSourceAccountProviderIdOk

`func (o *CostDataRecordDimensions) GetSourceAccountProviderIdOk() (*string, bool)`

GetSourceAccountProviderIdOk returns a tuple with the SourceAccountProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountProviderId

`func (o *CostDataRecordDimensions) SetSourceAccountProviderId(v string)`

SetSourceAccountProviderId sets SourceAccountProviderId field to given value.

### HasSourceAccountProviderId

`func (o *CostDataRecordDimensions) HasSourceAccountProviderId() bool`

HasSourceAccountProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


