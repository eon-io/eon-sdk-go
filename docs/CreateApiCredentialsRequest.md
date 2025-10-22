# CreateApiCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the api credential user | 
**RoleId** | **string** | The role ID of the api credential user | 
**Projects** | **[]string** | The projects the api credential user has access to | 
**EnableCurrentSecretRotation** | Pointer to **bool** | Whether the api credentials user can rotate their own secret | [optional] [default to false]

## Methods

### NewCreateApiCredentialsRequest

`func NewCreateApiCredentialsRequest(name string, roleId string, projects []string, ) *CreateApiCredentialsRequest`

NewCreateApiCredentialsRequest instantiates a new CreateApiCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiCredentialsRequestWithDefaults

`func NewCreateApiCredentialsRequestWithDefaults() *CreateApiCredentialsRequest`

NewCreateApiCredentialsRequestWithDefaults instantiates a new CreateApiCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApiCredentialsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiCredentialsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiCredentialsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *CreateApiCredentialsRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateApiCredentialsRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateApiCredentialsRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetProjects

`func (o *CreateApiCredentialsRequest) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateApiCredentialsRequest) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateApiCredentialsRequest) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetEnableCurrentSecretRotation

`func (o *CreateApiCredentialsRequest) GetEnableCurrentSecretRotation() bool`

GetEnableCurrentSecretRotation returns the EnableCurrentSecretRotation field if non-nil, zero value otherwise.

### GetEnableCurrentSecretRotationOk

`func (o *CreateApiCredentialsRequest) GetEnableCurrentSecretRotationOk() (*bool, bool)`

GetEnableCurrentSecretRotationOk returns a tuple with the EnableCurrentSecretRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCurrentSecretRotation

`func (o *CreateApiCredentialsRequest) SetEnableCurrentSecretRotation(v bool)`

SetEnableCurrentSecretRotation sets EnableCurrentSecretRotation field to given value.

### HasEnableCurrentSecretRotation

`func (o *CreateApiCredentialsRequest) HasEnableCurrentSecretRotation() bool`

HasEnableCurrentSecretRotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


