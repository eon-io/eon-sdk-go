# QueryDBRestoreDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Bucket** | Pointer to [**NullableS3RestoreDestination**](S3RestoreDestination.md) |  | [optional] 

## Methods

### NewQueryDBRestoreDestination

`func NewQueryDBRestoreDestination() *QueryDBRestoreDestination`

NewQueryDBRestoreDestination instantiates a new QueryDBRestoreDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDBRestoreDestinationWithDefaults

`func NewQueryDBRestoreDestinationWithDefaults() *QueryDBRestoreDestination`

NewQueryDBRestoreDestinationWithDefaults instantiates a new QueryDBRestoreDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Bucket

`func (o *QueryDBRestoreDestination) GetS3Bucket() S3RestoreDestination`

GetS3Bucket returns the S3Bucket field if non-nil, zero value otherwise.

### GetS3BucketOk

`func (o *QueryDBRestoreDestination) GetS3BucketOk() (*S3RestoreDestination, bool)`

GetS3BucketOk returns a tuple with the S3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Bucket

`func (o *QueryDBRestoreDestination) SetS3Bucket(v S3RestoreDestination)`

SetS3Bucket sets S3Bucket field to given value.

### HasS3Bucket

`func (o *QueryDBRestoreDestination) HasS3Bucket() bool`

HasS3Bucket returns a boolean if a field has been set.

### SetS3BucketNil

`func (o *QueryDBRestoreDestination) SetS3BucketNil(b bool)`

 SetS3BucketNil sets the value for S3Bucket to be an explicit nil

### UnsetS3Bucket
`func (o *QueryDBRestoreDestination) UnsetS3Bucket()`

UnsetS3Bucket ensures that no value is present for S3Bucket, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


