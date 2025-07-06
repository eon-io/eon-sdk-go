# RestoreFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**ObjectStorageDestination**](ObjectStorageDestination.md) |  | 
**Files** | [**[]FilePath**](FilePath.md) | List of file paths to restore. | 
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 

## Methods

### NewRestoreFilesRequest

`func NewRestoreFilesRequest(destination ObjectStorageDestination, files []FilePath, restoreAccountId string, ) *RestoreFilesRequest`

NewRestoreFilesRequest instantiates a new RestoreFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesRequestWithDefaults

`func NewRestoreFilesRequestWithDefaults() *RestoreFilesRequest`

NewRestoreFilesRequestWithDefaults instantiates a new RestoreFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *RestoreFilesRequest) GetDestination() ObjectStorageDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreFilesRequest) GetDestinationOk() (*ObjectStorageDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreFilesRequest) SetDestination(v ObjectStorageDestination)`

SetDestination sets Destination field to given value.


### GetFiles

`func (o *RestoreFilesRequest) GetFiles() []FilePath`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RestoreFilesRequest) GetFilesOk() (*[]FilePath, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RestoreFilesRequest) SetFiles(v []FilePath)`

SetFiles sets Files field to given value.


### GetRestoreAccountId

`func (o *RestoreFilesRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreFilesRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreFilesRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


