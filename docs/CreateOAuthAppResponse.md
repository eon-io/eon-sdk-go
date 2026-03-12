# CreateOAuthAppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthApp** | [**OAuthApp**](OAuthApp.md) |  | 
**ClientSecret** | **string** | The client secret. Only returned at creation time — store it securely. | 

## Methods

### NewCreateOAuthAppResponse

`func NewCreateOAuthAppResponse(oauthApp OAuthApp, clientSecret string, ) *CreateOAuthAppResponse`

NewCreateOAuthAppResponse instantiates a new CreateOAuthAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOAuthAppResponseWithDefaults

`func NewCreateOAuthAppResponseWithDefaults() *CreateOAuthAppResponse`

NewCreateOAuthAppResponseWithDefaults instantiates a new CreateOAuthAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthApp

`func (o *CreateOAuthAppResponse) GetOauthApp() OAuthApp`

GetOauthApp returns the OauthApp field if non-nil, zero value otherwise.

### GetOauthAppOk

`func (o *CreateOAuthAppResponse) GetOauthAppOk() (*OAuthApp, bool)`

GetOauthAppOk returns a tuple with the OauthApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthApp

`func (o *CreateOAuthAppResponse) SetOauthApp(v OAuthApp)`

SetOauthApp sets OauthApp field to given value.


### GetClientSecret

`func (o *CreateOAuthAppResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateOAuthAppResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateOAuthAppResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


