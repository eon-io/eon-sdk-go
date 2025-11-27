# GCSRestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Name of an existing bucket to restore the files to. | 
**Prefix** | Pointer to **string** | Prefix to add to the restore path. If you don&#39;t specify a prefix, the files are restored to their respective folders in the original file tree, starting from the root of the bucket.  | [optional] 

## Methods

### NewGCSRestoreTarget

`func NewGCSRestoreTarget(bucketName string, ) *GCSRestoreTarget`

NewGCSRestoreTarget instantiates a new GCSRestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSRestoreTargetWithDefaults

`func NewGCSRestoreTargetWithDefaults() *GCSRestoreTarget`

NewGCSRestoreTargetWithDefaults instantiates a new GCSRestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *GCSRestoreTarget) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GCSRestoreTarget) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GCSRestoreTarget) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetPrefix

`func (o *GCSRestoreTarget) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GCSRestoreTarget) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GCSRestoreTarget) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GCSRestoreTarget) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


