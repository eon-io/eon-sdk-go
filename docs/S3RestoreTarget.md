# S3RestoreTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Name of an existing bucket to restore the files to. | 
**EncryptionKeyId** | Pointer to **string** | ID of the key you want Eon to use for encrypting the restored files. | [optional] 
**Prefix** | Pointer to **string** | Prefix to add to the restore path. If you don&#39;t specify a prefix, the files are restored to their respective folders in the original file tree, starting from the root of the bucket.  | [optional] 

## Methods

### NewS3RestoreTarget

`func NewS3RestoreTarget(bucketName string, ) *S3RestoreTarget`

NewS3RestoreTarget instantiates a new S3RestoreTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3RestoreTargetWithDefaults

`func NewS3RestoreTargetWithDefaults() *S3RestoreTarget`

NewS3RestoreTargetWithDefaults instantiates a new S3RestoreTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3RestoreTarget) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3RestoreTarget) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3RestoreTarget) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetEncryptionKeyId

`func (o *S3RestoreTarget) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *S3RestoreTarget) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *S3RestoreTarget) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.

### HasEncryptionKeyId

`func (o *S3RestoreTarget) HasEncryptionKeyId() bool`

HasEncryptionKeyId returns a boolean if a field has been set.

### GetPrefix

`func (o *S3RestoreTarget) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3RestoreTarget) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3RestoreTarget) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3RestoreTarget) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


