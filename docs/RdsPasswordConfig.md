# RdsPasswordConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordManagementType** | Pointer to **string** | Password management type for the restored database.  | [optional] 
**Password** | Pointer to **string** | Customer-managed password (required when passwordManagementType is SELF_MANAGED).  | [optional] 
**SecretsManagerKmsKeyId** | Pointer to **string** | KMS key ARN used to encrypt the Secrets Manager secret that stores the database master password. Only applicable when passwordManagementType is AWS_SECRETS_MANAGER. This is distinct from encryptionKeyId which controls RDS storage encryption.  | [optional] 

## Methods

### NewRdsPasswordConfig

`func NewRdsPasswordConfig() *RdsPasswordConfig`

NewRdsPasswordConfig instantiates a new RdsPasswordConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsPasswordConfigWithDefaults

`func NewRdsPasswordConfigWithDefaults() *RdsPasswordConfig`

NewRdsPasswordConfigWithDefaults instantiates a new RdsPasswordConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordManagementType

`func (o *RdsPasswordConfig) GetPasswordManagementType() string`

GetPasswordManagementType returns the PasswordManagementType field if non-nil, zero value otherwise.

### GetPasswordManagementTypeOk

`func (o *RdsPasswordConfig) GetPasswordManagementTypeOk() (*string, bool)`

GetPasswordManagementTypeOk returns a tuple with the PasswordManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordManagementType

`func (o *RdsPasswordConfig) SetPasswordManagementType(v string)`

SetPasswordManagementType sets PasswordManagementType field to given value.

### HasPasswordManagementType

`func (o *RdsPasswordConfig) HasPasswordManagementType() bool`

HasPasswordManagementType returns a boolean if a field has been set.

### GetPassword

`func (o *RdsPasswordConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RdsPasswordConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RdsPasswordConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RdsPasswordConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecretsManagerKmsKeyId

`func (o *RdsPasswordConfig) GetSecretsManagerKmsKeyId() string`

GetSecretsManagerKmsKeyId returns the SecretsManagerKmsKeyId field if non-nil, zero value otherwise.

### GetSecretsManagerKmsKeyIdOk

`func (o *RdsPasswordConfig) GetSecretsManagerKmsKeyIdOk() (*string, bool)`

GetSecretsManagerKmsKeyIdOk returns a tuple with the SecretsManagerKmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsManagerKmsKeyId

`func (o *RdsPasswordConfig) SetSecretsManagerKmsKeyId(v string)`

SetSecretsManagerKmsKeyId sets SecretsManagerKmsKeyId field to given value.

### HasSecretsManagerKmsKeyId

`func (o *RdsPasswordConfig) HasSecretsManagerKmsKeyId() bool`

HasSecretsManagerKmsKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


