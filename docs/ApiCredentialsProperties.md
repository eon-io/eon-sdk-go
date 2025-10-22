# ApiCredentialsProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the api credential | 
**ClientId** | **string** | The client ID of the api credential | 
**Name** | **string** | The given name of the api credential | 
**Projects** | **[]string** | The projects the api credentials has access to | 
**RoleId** | **string** | The role ID of the api credential | 
**EnableCurrentSecretRotation** | Pointer to **bool** | Whether the api credentials user can rotate their own secret | [optional] [default to false]

## Methods

### NewApiCredentialsProperties

`func NewApiCredentialsProperties(id string, clientId string, name string, projects []string, roleId string, ) *ApiCredentialsProperties`

NewApiCredentialsProperties instantiates a new ApiCredentialsProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCredentialsPropertiesWithDefaults

`func NewApiCredentialsPropertiesWithDefaults() *ApiCredentialsProperties`

NewApiCredentialsPropertiesWithDefaults instantiates a new ApiCredentialsProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiCredentialsProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCredentialsProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCredentialsProperties) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *ApiCredentialsProperties) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiCredentialsProperties) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiCredentialsProperties) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetName

`func (o *ApiCredentialsProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiCredentialsProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiCredentialsProperties) SetName(v string)`

SetName sets Name field to given value.


### GetProjects

`func (o *ApiCredentialsProperties) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ApiCredentialsProperties) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ApiCredentialsProperties) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetRoleId

`func (o *ApiCredentialsProperties) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiCredentialsProperties) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiCredentialsProperties) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetEnableCurrentSecretRotation

`func (o *ApiCredentialsProperties) GetEnableCurrentSecretRotation() bool`

GetEnableCurrentSecretRotation returns the EnableCurrentSecretRotation field if non-nil, zero value otherwise.

### GetEnableCurrentSecretRotationOk

`func (o *ApiCredentialsProperties) GetEnableCurrentSecretRotationOk() (*bool, bool)`

GetEnableCurrentSecretRotationOk returns a tuple with the EnableCurrentSecretRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCurrentSecretRotation

`func (o *ApiCredentialsProperties) SetEnableCurrentSecretRotation(v bool)`

SetEnableCurrentSecretRotation sets EnableCurrentSecretRotation field to given value.

### HasEnableCurrentSecretRotation

`func (o *ApiCredentialsProperties) HasEnableCurrentSecretRotation() bool`

HasEnableCurrentSecretRotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


