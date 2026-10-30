# S3Buckets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]S3Bucket**](S3Bucket.md) | List of buckets. | 

## Methods

### NewS3Buckets

`func NewS3Buckets(buckets []S3Bucket, ) *S3Buckets`

NewS3Buckets instantiates a new S3Buckets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketsWithDefaults

`func NewS3BucketsWithDefaults() *S3Buckets`

NewS3BucketsWithDefaults instantiates a new S3Buckets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *S3Buckets) GetBuckets() []S3Bucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *S3Buckets) GetBucketsOk() (*[]S3Bucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *S3Buckets) SetBuckets(v []S3Bucket)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


