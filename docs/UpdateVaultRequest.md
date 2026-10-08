# UpdateVaultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Vault display name. | 
**Tags** | Pointer to **map[string]string** | User-defined metadata tags, used for filtering and search only. They do not affect backup behavior, access control, or the tags Eon sets on your cloud resources.  | [optional] 

## Methods

### NewUpdateVaultRequest

`func NewUpdateVaultRequest(name string, ) *UpdateVaultRequest`

NewUpdateVaultRequest instantiates a new UpdateVaultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVaultRequestWithDefaults

`func NewUpdateVaultRequestWithDefaults() *UpdateVaultRequest`

NewUpdateVaultRequestWithDefaults instantiates a new UpdateVaultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVaultRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVaultRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVaultRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *UpdateVaultRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVaultRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVaultRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVaultRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


