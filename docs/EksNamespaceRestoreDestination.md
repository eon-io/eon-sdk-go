# EksNamespaceRestoreDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Bucket** | Pointer to [**S3RestoreTarget**](S3RestoreTarget.md) |  | [optional] 

## Methods

### NewEksNamespaceRestoreDestination

`func NewEksNamespaceRestoreDestination() *EksNamespaceRestoreDestination`

NewEksNamespaceRestoreDestination instantiates a new EksNamespaceRestoreDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEksNamespaceRestoreDestinationWithDefaults

`func NewEksNamespaceRestoreDestinationWithDefaults() *EksNamespaceRestoreDestination`

NewEksNamespaceRestoreDestinationWithDefaults instantiates a new EksNamespaceRestoreDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Bucket

`func (o *EksNamespaceRestoreDestination) GetS3Bucket() S3RestoreTarget`

GetS3Bucket returns the S3Bucket field if non-nil, zero value otherwise.

### GetS3BucketOk

`func (o *EksNamespaceRestoreDestination) GetS3BucketOk() (*S3RestoreTarget, bool)`

GetS3BucketOk returns a tuple with the S3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Bucket

`func (o *EksNamespaceRestoreDestination) SetS3Bucket(v S3RestoreTarget)`

SetS3Bucket sets S3Bucket field to given value.

### HasS3Bucket

`func (o *EksNamespaceRestoreDestination) HasS3Bucket() bool`

HasS3Bucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


