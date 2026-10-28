# S3Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | The name of the bucket | 
**RegionName** | **string** | The region name where bucket is located | 

## Methods

### NewS3Bucket

`func NewS3Bucket(bucketName string, regionName string, ) *S3Bucket`

NewS3Bucket instantiates a new S3Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketWithDefaults

`func NewS3BucketWithDefaults() *S3Bucket`

NewS3BucketWithDefaults instantiates a new S3Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3Bucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3Bucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3Bucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegionName

`func (o *S3Bucket) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *S3Bucket) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *S3Bucket) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


