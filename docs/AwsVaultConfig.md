# AwsVaultConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | Pointer to **string** | ARN of the KMS key used for encryption. | [optional] 

## Methods

### NewAwsVaultConfig

`func NewAwsVaultConfig() *AwsVaultConfig`

NewAwsVaultConfig instantiates a new AwsVaultConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVaultConfigWithDefaults

`func NewAwsVaultConfigWithDefaults() *AwsVaultConfig`

NewAwsVaultConfigWithDefaults instantiates a new AwsVaultConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *AwsVaultConfig) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *AwsVaultConfig) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *AwsVaultConfig) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *AwsVaultConfig) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


