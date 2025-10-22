# CreateApiCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCredentialsProperties** | [**ApiCredentialsProperties**](ApiCredentialsProperties.md) |  | 
**ClientSecret** | **string** | The secret of the api credentials | 

## Methods

### NewCreateApiCredentialsResponse

`func NewCreateApiCredentialsResponse(apiCredentialsProperties ApiCredentialsProperties, clientSecret string, ) *CreateApiCredentialsResponse`

NewCreateApiCredentialsResponse instantiates a new CreateApiCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiCredentialsResponseWithDefaults

`func NewCreateApiCredentialsResponseWithDefaults() *CreateApiCredentialsResponse`

NewCreateApiCredentialsResponseWithDefaults instantiates a new CreateApiCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCredentialsProperties

`func (o *CreateApiCredentialsResponse) GetApiCredentialsProperties() ApiCredentialsProperties`

GetApiCredentialsProperties returns the ApiCredentialsProperties field if non-nil, zero value otherwise.

### GetApiCredentialsPropertiesOk

`func (o *CreateApiCredentialsResponse) GetApiCredentialsPropertiesOk() (*ApiCredentialsProperties, bool)`

GetApiCredentialsPropertiesOk returns a tuple with the ApiCredentialsProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredentialsProperties

`func (o *CreateApiCredentialsResponse) SetApiCredentialsProperties(v ApiCredentialsProperties)`

SetApiCredentialsProperties sets ApiCredentialsProperties field to given value.


### GetClientSecret

`func (o *CreateApiCredentialsResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateApiCredentialsResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateApiCredentialsResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


