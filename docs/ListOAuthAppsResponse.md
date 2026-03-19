# ListOAuthAppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthApps** | [**[]OAuthApp**](OAuthApp.md) | List of OAuth apps registered in the tenant&#39;s Cognito user pool. | 

## Methods

### NewListOAuthAppsResponse

`func NewListOAuthAppsResponse(oauthApps []OAuthApp, ) *ListOAuthAppsResponse`

NewListOAuthAppsResponse instantiates a new ListOAuthAppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOAuthAppsResponseWithDefaults

`func NewListOAuthAppsResponseWithDefaults() *ListOAuthAppsResponse`

NewListOAuthAppsResponseWithDefaults instantiates a new ListOAuthAppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthApps

`func (o *ListOAuthAppsResponse) GetOauthApps() []OAuthApp`

GetOauthApps returns the OauthApps field if non-nil, zero value otherwise.

### GetOauthAppsOk

`func (o *ListOAuthAppsResponse) GetOauthAppsOk() (*[]OAuthApp, bool)`

GetOauthAppsOk returns a tuple with the OauthApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthApps

`func (o *ListOAuthAppsResponse) SetOauthApps(v []OAuthApp)`

SetOauthApps sets OauthApps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


