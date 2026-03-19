# OAuthApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The Cognito App Client ID. | [optional] 
**Name** | Pointer to **string** | Display name of the OAuth app. | [optional] 
**CallbackURLs** | Pointer to **[]string** | Allowed redirect URIs. | [optional] 
**AuthUrl** | Pointer to **string** | Cognito hosted UI authorization endpoint URL. | [optional] 
**TokenUrl** | Pointer to **string** | Cognito token endpoint URL. | [optional] 
**Type** | Pointer to **string** | The OAuth app type. | [optional] 

## Methods

### NewOAuthApp

`func NewOAuthApp() *OAuthApp`

NewOAuthApp instantiates a new OAuthApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAppWithDefaults

`func NewOAuthAppWithDefaults() *OAuthApp`

NewOAuthAppWithDefaults instantiates a new OAuthApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthApp) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthApp) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthApp) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthApp) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetName

`func (o *OAuthApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuthApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCallbackURLs

`func (o *OAuthApp) GetCallbackURLs() []string`

GetCallbackURLs returns the CallbackURLs field if non-nil, zero value otherwise.

### GetCallbackURLsOk

`func (o *OAuthApp) GetCallbackURLsOk() (*[]string, bool)`

GetCallbackURLsOk returns a tuple with the CallbackURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackURLs

`func (o *OAuthApp) SetCallbackURLs(v []string)`

SetCallbackURLs sets CallbackURLs field to given value.

### HasCallbackURLs

`func (o *OAuthApp) HasCallbackURLs() bool`

HasCallbackURLs returns a boolean if a field has been set.

### GetAuthUrl

`func (o *OAuthApp) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *OAuthApp) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *OAuthApp) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *OAuthApp) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *OAuthApp) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuthApp) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuthApp) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *OAuthApp) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetType

`func (o *OAuthApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthApp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuthApp) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


