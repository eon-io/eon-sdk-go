# RestoreVolumeToEbsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderVolumeId** | **string** | Cloud-provider-assigned ID of the volume to restore. | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**EbsRestoreDestination**](EbsRestoreDestination.md) |  | 

## Methods

### NewRestoreVolumeToEbsRequest

`func NewRestoreVolumeToEbsRequest(providerVolumeId string, restoreAccountId string, destination EbsRestoreDestination, ) *RestoreVolumeToEbsRequest`

NewRestoreVolumeToEbsRequest instantiates a new RestoreVolumeToEbsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreVolumeToEbsRequestWithDefaults

`func NewRestoreVolumeToEbsRequestWithDefaults() *RestoreVolumeToEbsRequest`

NewRestoreVolumeToEbsRequestWithDefaults instantiates a new RestoreVolumeToEbsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderVolumeId

`func (o *RestoreVolumeToEbsRequest) GetProviderVolumeId() string`

GetProviderVolumeId returns the ProviderVolumeId field if non-nil, zero value otherwise.

### GetProviderVolumeIdOk

`func (o *RestoreVolumeToEbsRequest) GetProviderVolumeIdOk() (*string, bool)`

GetProviderVolumeIdOk returns a tuple with the ProviderVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVolumeId

`func (o *RestoreVolumeToEbsRequest) SetProviderVolumeId(v string)`

SetProviderVolumeId sets ProviderVolumeId field to given value.


### GetRestoreAccountId

`func (o *RestoreVolumeToEbsRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreVolumeToEbsRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreVolumeToEbsRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreVolumeToEbsRequest) GetDestination() EbsRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreVolumeToEbsRequest) GetDestinationOk() (*EbsRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreVolumeToEbsRequest) SetDestination(v EbsRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


