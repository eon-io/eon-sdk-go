# ObjectStorageDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Bucket** | Pointer to [**S3RestoreTarget**](S3RestoreTarget.md) |  | [optional] 
**StorageAccount** | Pointer to [**StorageAccountRestoreTarget**](StorageAccountRestoreTarget.md) |  | [optional] 

## Methods

### NewObjectStorageDestination

`func NewObjectStorageDestination() *ObjectStorageDestination`

NewObjectStorageDestination instantiates a new ObjectStorageDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageDestinationWithDefaults

`func NewObjectStorageDestinationWithDefaults() *ObjectStorageDestination`

NewObjectStorageDestinationWithDefaults instantiates a new ObjectStorageDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Bucket

`func (o *ObjectStorageDestination) GetS3Bucket() S3RestoreTarget`

GetS3Bucket returns the S3Bucket field if non-nil, zero value otherwise.

### GetS3BucketOk

`func (o *ObjectStorageDestination) GetS3BucketOk() (*S3RestoreTarget, bool)`

GetS3BucketOk returns a tuple with the S3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Bucket

`func (o *ObjectStorageDestination) SetS3Bucket(v S3RestoreTarget)`

SetS3Bucket sets S3Bucket field to given value.

### HasS3Bucket

`func (o *ObjectStorageDestination) HasS3Bucket() bool`

HasS3Bucket returns a boolean if a field has been set.

### GetStorageAccount

`func (o *ObjectStorageDestination) GetStorageAccount() StorageAccountRestoreTarget`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *ObjectStorageDestination) GetStorageAccountOk() (*StorageAccountRestoreTarget, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *ObjectStorageDestination) SetStorageAccount(v StorageAccountRestoreTarget)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *ObjectStorageDestination) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


