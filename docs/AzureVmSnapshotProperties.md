# AzureVmSnapshotProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the resource | [optional] 
**SizeBytes** | Pointer to **string** | The size of the virtual machine | [optional] 
**Disks** | Pointer to [**[]AzureDisk**](AzureDisk.md) |  | [optional] 

## Methods

### NewAzureVmSnapshotProperties

`func NewAzureVmSnapshotProperties() *AzureVmSnapshotProperties`

NewAzureVmSnapshotProperties instantiates a new AzureVmSnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVmSnapshotPropertiesWithDefaults

`func NewAzureVmSnapshotPropertiesWithDefaults() *AzureVmSnapshotProperties`

NewAzureVmSnapshotPropertiesWithDefaults instantiates a new AzureVmSnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureVmSnapshotProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVmSnapshotProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVmSnapshotProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVmSnapshotProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeBytes

`func (o *AzureVmSnapshotProperties) GetSizeBytes() string`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *AzureVmSnapshotProperties) GetSizeBytesOk() (*string, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *AzureVmSnapshotProperties) SetSizeBytes(v string)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *AzureVmSnapshotProperties) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetDisks

`func (o *AzureVmSnapshotProperties) GetDisks() []AzureDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *AzureVmSnapshotProperties) GetDisksOk() (*[]AzureDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *AzureVmSnapshotProperties) SetDisks(v []AzureDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *AzureVmSnapshotProperties) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


