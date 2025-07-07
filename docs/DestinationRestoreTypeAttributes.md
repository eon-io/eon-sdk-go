# DestinationRestoreTypeAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectStorage** | Pointer to [**ObjectStorageDestination**](ObjectStorageDestination.md) |  | [optional] 
**AzureVm** | Pointer to [**NullableVirtualMachineRestoreDestination**](VirtualMachineRestoreDestination.md) |  | [optional] 
**DynamoDbTable** | Pointer to [**AwsDynamoDBDestination**](AwsDynamoDBDestination.md) |  | [optional] 
**EksNamespace** | Pointer to [**NullableEksNamespaceRestoreDestination**](EksNamespaceRestoreDestination.md) |  | [optional] 
**RdsInstance** | Pointer to [**AwsDatabaseDestination**](AwsDatabaseDestination.md) |  | [optional] 

## Methods

### NewDestinationRestoreTypeAttributes

`func NewDestinationRestoreTypeAttributes() *DestinationRestoreTypeAttributes`

NewDestinationRestoreTypeAttributes instantiates a new DestinationRestoreTypeAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationRestoreTypeAttributesWithDefaults

`func NewDestinationRestoreTypeAttributesWithDefaults() *DestinationRestoreTypeAttributes`

NewDestinationRestoreTypeAttributesWithDefaults instantiates a new DestinationRestoreTypeAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectStorage

`func (o *DestinationRestoreTypeAttributes) GetObjectStorage() ObjectStorageDestination`

GetObjectStorage returns the ObjectStorage field if non-nil, zero value otherwise.

### GetObjectStorageOk

`func (o *DestinationRestoreTypeAttributes) GetObjectStorageOk() (*ObjectStorageDestination, bool)`

GetObjectStorageOk returns a tuple with the ObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorage

`func (o *DestinationRestoreTypeAttributes) SetObjectStorage(v ObjectStorageDestination)`

SetObjectStorage sets ObjectStorage field to given value.

### HasObjectStorage

`func (o *DestinationRestoreTypeAttributes) HasObjectStorage() bool`

HasObjectStorage returns a boolean if a field has been set.

### GetAzureVm

`func (o *DestinationRestoreTypeAttributes) GetAzureVm() VirtualMachineRestoreDestination`

GetAzureVm returns the AzureVm field if non-nil, zero value otherwise.

### GetAzureVmOk

`func (o *DestinationRestoreTypeAttributes) GetAzureVmOk() (*VirtualMachineRestoreDestination, bool)`

GetAzureVmOk returns a tuple with the AzureVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVm

`func (o *DestinationRestoreTypeAttributes) SetAzureVm(v VirtualMachineRestoreDestination)`

SetAzureVm sets AzureVm field to given value.

### HasAzureVm

`func (o *DestinationRestoreTypeAttributes) HasAzureVm() bool`

HasAzureVm returns a boolean if a field has been set.

### SetAzureVmNil

`func (o *DestinationRestoreTypeAttributes) SetAzureVmNil(b bool)`

 SetAzureVmNil sets the value for AzureVm to be an explicit nil

### UnsetAzureVm
`func (o *DestinationRestoreTypeAttributes) UnsetAzureVm()`

UnsetAzureVm ensures that no value is present for AzureVm, not even an explicit nil
### GetDynamoDbTable

`func (o *DestinationRestoreTypeAttributes) GetDynamoDbTable() AwsDynamoDBDestination`

GetDynamoDbTable returns the DynamoDbTable field if non-nil, zero value otherwise.

### GetDynamoDbTableOk

`func (o *DestinationRestoreTypeAttributes) GetDynamoDbTableOk() (*AwsDynamoDBDestination, bool)`

GetDynamoDbTableOk returns a tuple with the DynamoDbTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamoDbTable

`func (o *DestinationRestoreTypeAttributes) SetDynamoDbTable(v AwsDynamoDBDestination)`

SetDynamoDbTable sets DynamoDbTable field to given value.

### HasDynamoDbTable

`func (o *DestinationRestoreTypeAttributes) HasDynamoDbTable() bool`

HasDynamoDbTable returns a boolean if a field has been set.

### GetEksNamespace

`func (o *DestinationRestoreTypeAttributes) GetEksNamespace() EksNamespaceRestoreDestination`

GetEksNamespace returns the EksNamespace field if non-nil, zero value otherwise.

### GetEksNamespaceOk

`func (o *DestinationRestoreTypeAttributes) GetEksNamespaceOk() (*EksNamespaceRestoreDestination, bool)`

GetEksNamespaceOk returns a tuple with the EksNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEksNamespace

`func (o *DestinationRestoreTypeAttributes) SetEksNamespace(v EksNamespaceRestoreDestination)`

SetEksNamespace sets EksNamespace field to given value.

### HasEksNamespace

`func (o *DestinationRestoreTypeAttributes) HasEksNamespace() bool`

HasEksNamespace returns a boolean if a field has been set.

### SetEksNamespaceNil

`func (o *DestinationRestoreTypeAttributes) SetEksNamespaceNil(b bool)`

 SetEksNamespaceNil sets the value for EksNamespace to be an explicit nil

### UnsetEksNamespace
`func (o *DestinationRestoreTypeAttributes) UnsetEksNamespace()`

UnsetEksNamespace ensures that no value is present for EksNamespace, not even an explicit nil
### GetRdsInstance

`func (o *DestinationRestoreTypeAttributes) GetRdsInstance() AwsDatabaseDestination`

GetRdsInstance returns the RdsInstance field if non-nil, zero value otherwise.

### GetRdsInstanceOk

`func (o *DestinationRestoreTypeAttributes) GetRdsInstanceOk() (*AwsDatabaseDestination, bool)`

GetRdsInstanceOk returns a tuple with the RdsInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsInstance

`func (o *DestinationRestoreTypeAttributes) SetRdsInstance(v AwsDatabaseDestination)`

SetRdsInstance sets RdsInstance field to given value.

### HasRdsInstance

`func (o *DestinationRestoreTypeAttributes) HasRdsInstance() bool`

HasRdsInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


