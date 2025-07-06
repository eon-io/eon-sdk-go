# RestoreVolumeToEbsSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderVolumeId** | **string** | Cloud-provider-assigned ID of the volume to restore. | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**EbsSnapshotRestoreDestination**](EbsSnapshotRestoreDestination.md) |  | 

## Methods

### NewRestoreVolumeToEbsSnapshotRequest

`func NewRestoreVolumeToEbsSnapshotRequest(providerVolumeId string, restoreAccountId string, destination EbsSnapshotRestoreDestination, ) *RestoreVolumeToEbsSnapshotRequest`

NewRestoreVolumeToEbsSnapshotRequest instantiates a new RestoreVolumeToEbsSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreVolumeToEbsSnapshotRequestWithDefaults

`func NewRestoreVolumeToEbsSnapshotRequestWithDefaults() *RestoreVolumeToEbsSnapshotRequest`

NewRestoreVolumeToEbsSnapshotRequestWithDefaults instantiates a new RestoreVolumeToEbsSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderVolumeId

`func (o *RestoreVolumeToEbsSnapshotRequest) GetProviderVolumeId() string`

GetProviderVolumeId returns the ProviderVolumeId field if non-nil, zero value otherwise.

### GetProviderVolumeIdOk

`func (o *RestoreVolumeToEbsSnapshotRequest) GetProviderVolumeIdOk() (*string, bool)`

GetProviderVolumeIdOk returns a tuple with the ProviderVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVolumeId

`func (o *RestoreVolumeToEbsSnapshotRequest) SetProviderVolumeId(v string)`

SetProviderVolumeId sets ProviderVolumeId field to given value.


### GetRestoreAccountId

`func (o *RestoreVolumeToEbsSnapshotRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreVolumeToEbsSnapshotRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreVolumeToEbsSnapshotRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreVolumeToEbsSnapshotRequest) GetDestination() EbsSnapshotRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreVolumeToEbsSnapshotRequest) GetDestinationOk() (*EbsSnapshotRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreVolumeToEbsSnapshotRequest) SetDestination(v EbsSnapshotRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


