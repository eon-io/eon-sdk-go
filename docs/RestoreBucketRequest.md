# RestoreBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**ObjectStorageDestination**](ObjectStorageDestination.md) |  | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 

## Methods

### NewRestoreBucketRequest

`func NewRestoreBucketRequest(destination ObjectStorageDestination, restoreAccountId string, ) *RestoreBucketRequest`

NewRestoreBucketRequest instantiates a new RestoreBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreBucketRequestWithDefaults

`func NewRestoreBucketRequestWithDefaults() *RestoreBucketRequest`

NewRestoreBucketRequestWithDefaults instantiates a new RestoreBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *RestoreBucketRequest) GetDestination() ObjectStorageDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreBucketRequest) GetDestinationOk() (*ObjectStorageDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreBucketRequest) SetDestination(v ObjectStorageDestination)`

SetDestination sets Destination field to given value.


### GetRestoreAccountId

`func (o *RestoreBucketRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreBucketRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreBucketRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


