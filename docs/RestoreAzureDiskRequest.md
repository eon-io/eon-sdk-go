# RestoreAzureDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore. | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**AzureDiskRestoreDestination**](AzureDiskRestoreDestination.md) |  | 

## Methods

### NewRestoreAzureDiskRequest

`func NewRestoreAzureDiskRequest(providerDiskId string, restoreAccountId string, destination AzureDiskRestoreDestination, ) *RestoreAzureDiskRequest`

NewRestoreAzureDiskRequest instantiates a new RestoreAzureDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAzureDiskRequestWithDefaults

`func NewRestoreAzureDiskRequestWithDefaults() *RestoreAzureDiskRequest`

NewRestoreAzureDiskRequestWithDefaults instantiates a new RestoreAzureDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderDiskId

`func (o *RestoreAzureDiskRequest) GetProviderDiskId() string`

GetProviderDiskId returns the ProviderDiskId field if non-nil, zero value otherwise.

### GetProviderDiskIdOk

`func (o *RestoreAzureDiskRequest) GetProviderDiskIdOk() (*string, bool)`

GetProviderDiskIdOk returns a tuple with the ProviderDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDiskId

`func (o *RestoreAzureDiskRequest) SetProviderDiskId(v string)`

SetProviderDiskId sets ProviderDiskId field to given value.


### GetRestoreAccountId

`func (o *RestoreAzureDiskRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreAzureDiskRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreAzureDiskRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreAzureDiskRequest) GetDestination() AzureDiskRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreAzureDiskRequest) GetDestinationOk() (*AzureDiskRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreAzureDiskRequest) SetDestination(v AzureDiskRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


