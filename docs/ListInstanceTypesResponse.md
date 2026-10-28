# ListInstanceTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypes** | Pointer to **[]string** | The instance types list | [optional] 
**InstantRestoreMapping** | Pointer to **map[string]string** | Map from post-promotion instance type → instant-restore host instance type. Keys present in &#x60;instanceTypes&#x60; but absent from this map have no instant-restore option in this AZ.  | [optional] 

## Methods

### NewListInstanceTypesResponse

`func NewListInstanceTypesResponse() *ListInstanceTypesResponse`

NewListInstanceTypesResponse instantiates a new ListInstanceTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceTypesResponseWithDefaults

`func NewListInstanceTypesResponseWithDefaults() *ListInstanceTypesResponse`

NewListInstanceTypesResponseWithDefaults instantiates a new ListInstanceTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypes

`func (o *ListInstanceTypesResponse) GetInstanceTypes() []string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *ListInstanceTypesResponse) GetInstanceTypesOk() (*[]string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *ListInstanceTypesResponse) SetInstanceTypes(v []string)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *ListInstanceTypesResponse) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetInstantRestoreMapping

`func (o *ListInstanceTypesResponse) GetInstantRestoreMapping() map[string]string`

GetInstantRestoreMapping returns the InstantRestoreMapping field if non-nil, zero value otherwise.

### GetInstantRestoreMappingOk

`func (o *ListInstanceTypesResponse) GetInstantRestoreMappingOk() (*map[string]string, bool)`

GetInstantRestoreMappingOk returns a tuple with the InstantRestoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRestoreMapping

`func (o *ListInstanceTypesResponse) SetInstantRestoreMapping(v map[string]string)`

SetInstantRestoreMapping sets InstantRestoreMapping field to given value.

### HasInstantRestoreMapping

`func (o *ListInstanceTypesResponse) HasInstantRestoreMapping() bool`

HasInstantRestoreMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


