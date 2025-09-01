# RestoreAzureVmInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**AzureVmInstanceRestoreDestination**](AzureVmInstanceRestoreDestination.md) |  | 

## Methods

### NewRestoreAzureVmInstanceRequest

`func NewRestoreAzureVmInstanceRequest(restoreAccountId string, destination AzureVmInstanceRestoreDestination, ) *RestoreAzureVmInstanceRequest`

NewRestoreAzureVmInstanceRequest instantiates a new RestoreAzureVmInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAzureVmInstanceRequestWithDefaults

`func NewRestoreAzureVmInstanceRequestWithDefaults() *RestoreAzureVmInstanceRequest`

NewRestoreAzureVmInstanceRequestWithDefaults instantiates a new RestoreAzureVmInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *RestoreAzureVmInstanceRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreAzureVmInstanceRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreAzureVmInstanceRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreAzureVmInstanceRequest) GetDestination() AzureVmInstanceRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreAzureVmInstanceRequest) GetDestinationOk() (*AzureVmInstanceRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreAzureVmInstanceRequest) SetDestination(v AzureVmInstanceRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


