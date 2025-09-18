# RestoreGcpDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderDiskId** | **string** | Cloud-provider-assigned ID of the disk to restore. | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**GcpDiskRestoreDestination**](GcpDiskRestoreDestination.md) |  | 

## Methods

### NewRestoreGcpDiskRequest

`func NewRestoreGcpDiskRequest(providerDiskId string, restoreAccountId string, destination GcpDiskRestoreDestination, ) *RestoreGcpDiskRequest`

NewRestoreGcpDiskRequest instantiates a new RestoreGcpDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreGcpDiskRequestWithDefaults

`func NewRestoreGcpDiskRequestWithDefaults() *RestoreGcpDiskRequest`

NewRestoreGcpDiskRequestWithDefaults instantiates a new RestoreGcpDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderDiskId

`func (o *RestoreGcpDiskRequest) GetProviderDiskId() string`

GetProviderDiskId returns the ProviderDiskId field if non-nil, zero value otherwise.

### GetProviderDiskIdOk

`func (o *RestoreGcpDiskRequest) GetProviderDiskIdOk() (*string, bool)`

GetProviderDiskIdOk returns a tuple with the ProviderDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDiskId

`func (o *RestoreGcpDiskRequest) SetProviderDiskId(v string)`

SetProviderDiskId sets ProviderDiskId field to given value.


### GetRestoreAccountId

`func (o *RestoreGcpDiskRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreGcpDiskRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreGcpDiskRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreGcpDiskRequest) GetDestination() GcpDiskRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreGcpDiskRequest) GetDestinationOk() (*GcpDiskRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreGcpDiskRequest) SetDestination(v GcpDiskRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


