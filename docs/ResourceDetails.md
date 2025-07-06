# ResourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned resource ID. | 
**ProviderResourceId** | **string** | Cloud-provider-assigned resource ID. | 
**ResourceName** | **string** | Resource display name. | 
**ResourceType** | [**ResourceType**](ResourceType.md) |  | 
**ProviderAccountId** | **string** | Cloud-provider-assigned account ID. | 
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Region** | **string** | Region the resource is hosted in. | 
**SourceStorageSizeBytes** | **int64** | Total storage size at the source, in bytes. | 

## Methods

### NewResourceDetails

`func NewResourceDetails(id string, providerResourceId string, resourceName string, resourceType ResourceType, providerAccountId string, cloudProvider Provider, region string, sourceStorageSizeBytes int64, ) *ResourceDetails`

NewResourceDetails instantiates a new ResourceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDetailsWithDefaults

`func NewResourceDetailsWithDefaults() *ResourceDetails`

NewResourceDetailsWithDefaults instantiates a new ResourceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceDetails) SetId(v string)`

SetId sets Id field to given value.


### GetProviderResourceId

`func (o *ResourceDetails) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *ResourceDetails) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *ResourceDetails) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.


### GetResourceName

`func (o *ResourceDetails) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceDetails) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceDetails) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *ResourceDetails) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceDetails) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceDetails) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetProviderAccountId

`func (o *ResourceDetails) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *ResourceDetails) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *ResourceDetails) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetCloudProvider

`func (o *ResourceDetails) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceDetails) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceDetails) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ResourceDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ResourceDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ResourceDetails) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSourceStorageSizeBytes

`func (o *ResourceDetails) GetSourceStorageSizeBytes() int64`

GetSourceStorageSizeBytes returns the SourceStorageSizeBytes field if non-nil, zero value otherwise.

### GetSourceStorageSizeBytesOk

`func (o *ResourceDetails) GetSourceStorageSizeBytesOk() (*int64, bool)`

GetSourceStorageSizeBytesOk returns a tuple with the SourceStorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStorageSizeBytes

`func (o *ResourceDetails) SetSourceStorageSizeBytes(v int64)`

SetSourceStorageSizeBytes sets SourceStorageSizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


