# EncryptionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Cloud-provider-assigned ID of the encryption key. | 
**Arn** | **string** | ARN of the encryption key. | 
**AliasName** | **string** | Encryption key alias. | 

## Methods

### NewEncryptionKey

`func NewEncryptionKey(id string, arn string, aliasName string, ) *EncryptionKey`

NewEncryptionKey instantiates a new EncryptionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyWithDefaults

`func NewEncryptionKeyWithDefaults() *EncryptionKey`

NewEncryptionKeyWithDefaults instantiates a new EncryptionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EncryptionKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionKey) SetId(v string)`

SetId sets Id field to given value.


### GetArn

`func (o *EncryptionKey) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *EncryptionKey) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *EncryptionKey) SetArn(v string)`

SetArn sets Arn field to given value.


### GetAliasName

`func (o *EncryptionKey) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *EncryptionKey) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *EncryptionKey) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


