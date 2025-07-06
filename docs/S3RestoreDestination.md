# S3RestoreDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Name of an existing bucket to restore query results to. Must be in the same region as the snapshot&#39;s vault.  | 
**Prefix** | Pointer to **string** | Prefix to prepend to the restore path. If omitted, query results are restored to the bucket root.  | [optional] 

## Methods

### NewS3RestoreDestination

`func NewS3RestoreDestination(bucketName string, ) *S3RestoreDestination`

NewS3RestoreDestination instantiates a new S3RestoreDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3RestoreDestinationWithDefaults

`func NewS3RestoreDestinationWithDefaults() *S3RestoreDestination`

NewS3RestoreDestinationWithDefaults instantiates a new S3RestoreDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3RestoreDestination) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3RestoreDestination) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3RestoreDestination) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetPrefix

`func (o *S3RestoreDestination) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3RestoreDestination) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3RestoreDestination) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3RestoreDestination) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


