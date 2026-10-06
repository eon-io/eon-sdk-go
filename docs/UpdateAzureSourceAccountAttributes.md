# UpdateAzureSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]string** | Azure regions Eon discovers in. Omit to leave the current regions unchanged. Pass an empty list to discover in all supported regions.  | [optional] 

## Methods

### NewUpdateAzureSourceAccountAttributes

`func NewUpdateAzureSourceAccountAttributes() *UpdateAzureSourceAccountAttributes`

NewUpdateAzureSourceAccountAttributes instantiates a new UpdateAzureSourceAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAzureSourceAccountAttributesWithDefaults

`func NewUpdateAzureSourceAccountAttributesWithDefaults() *UpdateAzureSourceAccountAttributes`

NewUpdateAzureSourceAccountAttributesWithDefaults instantiates a new UpdateAzureSourceAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *UpdateAzureSourceAccountAttributes) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *UpdateAzureSourceAccountAttributes) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *UpdateAzureSourceAccountAttributes) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *UpdateAzureSourceAccountAttributes) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


