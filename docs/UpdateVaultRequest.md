# UpdateVaultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Vault display name. | 
**Labels** | Pointer to **[]string** | User-defined metadata labels, used for filtering and search only. They do not affect backup behavior, access control, or the tags Eon sets on your cloud resources.  | [optional] 

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


### GetLabels

`func (o *UpdateVaultRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateVaultRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateVaultRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateVaultRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


