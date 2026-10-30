# S3TemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetBucketId** | **string** | Name of the target S3 bucket for the restore. | 
**PrefixTemplate** | **string** | Template for the prefix to use when restoring objects. | 
**TargetAccountId** | **string** | Eon-assigned ID of the target restore account. | 
**TargetRegion** | **string** | Target region for the restore. | 
**KmsKeyArn** | Pointer to **string** | ARN of a customer-managed KMS key in the target account and region. It encrypts the restore&#39;s temporary secrets in the restore account: the vault access secret and the container-registry token. Restored objects use the target bucket&#39;s default encryption. AWS-managed keys aren&#39;t accepted.  | [optional] 

## Methods

### NewS3TemplateParams

`func NewS3TemplateParams(targetBucketId string, prefixTemplate string, targetAccountId string, targetRegion string, ) *S3TemplateParams`

NewS3TemplateParams instantiates a new S3TemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3TemplateParamsWithDefaults

`func NewS3TemplateParamsWithDefaults() *S3TemplateParams`

NewS3TemplateParamsWithDefaults instantiates a new S3TemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetBucketId

`func (o *S3TemplateParams) GetTargetBucketId() string`

GetTargetBucketId returns the TargetBucketId field if non-nil, zero value otherwise.

### GetTargetBucketIdOk

`func (o *S3TemplateParams) GetTargetBucketIdOk() (*string, bool)`

GetTargetBucketIdOk returns a tuple with the TargetBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBucketId

`func (o *S3TemplateParams) SetTargetBucketId(v string)`

SetTargetBucketId sets TargetBucketId field to given value.


### GetPrefixTemplate

`func (o *S3TemplateParams) GetPrefixTemplate() string`

GetPrefixTemplate returns the PrefixTemplate field if non-nil, zero value otherwise.

### GetPrefixTemplateOk

`func (o *S3TemplateParams) GetPrefixTemplateOk() (*string, bool)`

GetPrefixTemplateOk returns a tuple with the PrefixTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixTemplate

`func (o *S3TemplateParams) SetPrefixTemplate(v string)`

SetPrefixTemplate sets PrefixTemplate field to given value.


### GetTargetAccountId

`func (o *S3TemplateParams) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *S3TemplateParams) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *S3TemplateParams) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.


### GetTargetRegion

`func (o *S3TemplateParams) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *S3TemplateParams) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *S3TemplateParams) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


### GetKmsKeyArn

`func (o *S3TemplateParams) GetKmsKeyArn() string`

GetKmsKeyArn returns the KmsKeyArn field if non-nil, zero value otherwise.

### GetKmsKeyArnOk

`func (o *S3TemplateParams) GetKmsKeyArnOk() (*string, bool)`

GetKmsKeyArnOk returns a tuple with the KmsKeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyArn

`func (o *S3TemplateParams) SetKmsKeyArn(v string)`

SetKmsKeyArn sets KmsKeyArn field to given value.

### HasKmsKeyArn

`func (o *S3TemplateParams) HasKmsKeyArn() bool`

HasKmsKeyArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


