# Oauth2TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Access token, which you pass in the &#x60;Authorization&#x60; header when calling API methods. | [optional] 
**TokenType** | Pointer to **string** | Type of the token, which is always &#x60;Bearer&#x60;. | [optional] 
**ExpiresIn** | Pointer to **int32** | Length of time until the token expires, in seconds. Default: &#x60;43200&#x60; (12 hours).  | [optional] 

## Methods

### NewOauth2TokenResponse

`func NewOauth2TokenResponse() *Oauth2TokenResponse`

NewOauth2TokenResponse instantiates a new Oauth2TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenResponseWithDefaults

`func NewOauth2TokenResponseWithDefaults() *Oauth2TokenResponse`

NewOauth2TokenResponseWithDefaults instantiates a new Oauth2TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Oauth2TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Oauth2TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Oauth2TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Oauth2TokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *Oauth2TokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Oauth2TokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Oauth2TokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Oauth2TokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *Oauth2TokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Oauth2TokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Oauth2TokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Oauth2TokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


