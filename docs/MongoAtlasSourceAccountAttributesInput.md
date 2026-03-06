# MongoAtlasSourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiPublicKey** | **string** | Public key of the MongoDB Atlas API key with Organization Read Only permissions. | 
**ApiPrivateKey** | **string** | Private key of the MongoDB Atlas API key. | 
**AtlasProjectId** | **string** | The MongoDB Atlas project ID to onboard. | 

## Methods

### NewMongoAtlasSourceAccountAttributesInput

`func NewMongoAtlasSourceAccountAttributesInput(apiPublicKey string, apiPrivateKey string, atlasProjectId string, ) *MongoAtlasSourceAccountAttributesInput`

NewMongoAtlasSourceAccountAttributesInput instantiates a new MongoAtlasSourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoAtlasSourceAccountAttributesInputWithDefaults

`func NewMongoAtlasSourceAccountAttributesInputWithDefaults() *MongoAtlasSourceAccountAttributesInput`

NewMongoAtlasSourceAccountAttributesInputWithDefaults instantiates a new MongoAtlasSourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiPublicKey

`func (o *MongoAtlasSourceAccountAttributesInput) GetApiPublicKey() string`

GetApiPublicKey returns the ApiPublicKey field if non-nil, zero value otherwise.

### GetApiPublicKeyOk

`func (o *MongoAtlasSourceAccountAttributesInput) GetApiPublicKeyOk() (*string, bool)`

GetApiPublicKeyOk returns a tuple with the ApiPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPublicKey

`func (o *MongoAtlasSourceAccountAttributesInput) SetApiPublicKey(v string)`

SetApiPublicKey sets ApiPublicKey field to given value.


### GetApiPrivateKey

`func (o *MongoAtlasSourceAccountAttributesInput) GetApiPrivateKey() string`

GetApiPrivateKey returns the ApiPrivateKey field if non-nil, zero value otherwise.

### GetApiPrivateKeyOk

`func (o *MongoAtlasSourceAccountAttributesInput) GetApiPrivateKeyOk() (*string, bool)`

GetApiPrivateKeyOk returns a tuple with the ApiPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPrivateKey

`func (o *MongoAtlasSourceAccountAttributesInput) SetApiPrivateKey(v string)`

SetApiPrivateKey sets ApiPrivateKey field to given value.


### GetAtlasProjectId

`func (o *MongoAtlasSourceAccountAttributesInput) GetAtlasProjectId() string`

GetAtlasProjectId returns the AtlasProjectId field if non-nil, zero value otherwise.

### GetAtlasProjectIdOk

`func (o *MongoAtlasSourceAccountAttributesInput) GetAtlasProjectIdOk() (*string, bool)`

GetAtlasProjectIdOk returns a tuple with the AtlasProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlasProjectId

`func (o *MongoAtlasSourceAccountAttributesInput) SetAtlasProjectId(v string)`

SetAtlasProjectId sets AtlasProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


