# ResourceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**NullableResourceSnapshotProperties**](ResourceSnapshotProperties.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Resource tags as key-value pairs. Both keys and values are strings. If a tag is a key with no value, the value is presented as an empty string.  **Example:** &#x60;{\&quot;Name\&quot;: \&quot;customers\&quot;}&#x60;  | [optional] 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 

## Methods

### NewResourceSnapshot

`func NewResourceSnapshot() *ResourceSnapshot`

NewResourceSnapshot instantiates a new ResourceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSnapshotWithDefaults

`func NewResourceSnapshotWithDefaults() *ResourceSnapshot`

NewResourceSnapshotWithDefaults instantiates a new ResourceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ResourceSnapshot) GetProperties() ResourceSnapshotProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ResourceSnapshot) GetPropertiesOk() (*ResourceSnapshotProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ResourceSnapshot) SetProperties(v ResourceSnapshotProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ResourceSnapshot) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ResourceSnapshot) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ResourceSnapshot) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetTags

`func (o *ResourceSnapshot) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceSnapshot) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceSnapshot) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceSnapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceSnapshot) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceSnapshot) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceSnapshot) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceSnapshot) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


