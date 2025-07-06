# StorageAccountRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the storage account to restore the files to. | 
**ResourceGroup** | Pointer to **string** | Name of the resource group that contains the storage account. | [optional] 
**Container** | **string** | Name of the container in the storage account to restore the files to. | 
**Prefix** | Pointer to **string** | Prefix to add to the restore path. If you don&#39;t specify a prefix, the files are restored to their respective folders in the original file tree, starting from the root of the container.  | [optional] 

## Methods

### NewStorageAccountRestoreTarget

`func NewStorageAccountRestoreTarget(name string, container string, ) *StorageAccountRestoreTarget`

NewStorageAccountRestoreTarget instantiates a new StorageAccountRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAccountRestoreTargetWithDefaults

`func NewStorageAccountRestoreTargetWithDefaults() *StorageAccountRestoreTarget`

NewStorageAccountRestoreTargetWithDefaults instantiates a new StorageAccountRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageAccountRestoreTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageAccountRestoreTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageAccountRestoreTarget) SetName(v string)`

SetName sets Name field to given value.


### GetResourceGroup

`func (o *StorageAccountRestoreTarget) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *StorageAccountRestoreTarget) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *StorageAccountRestoreTarget) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *StorageAccountRestoreTarget) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetContainer

`func (o *StorageAccountRestoreTarget) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *StorageAccountRestoreTarget) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *StorageAccountRestoreTarget) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetPrefix

`func (o *StorageAccountRestoreTarget) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *StorageAccountRestoreTarget) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *StorageAccountRestoreTarget) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *StorageAccountRestoreTarget) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


