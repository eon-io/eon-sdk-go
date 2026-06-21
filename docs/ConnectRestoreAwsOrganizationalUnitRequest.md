# ConnectRestoreAwsOrganizationalUnitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the role Eon assumes to access the organizational unit in AWS. | 
**ProviderOrganizationalUnitId** | **string** | AWS-assigned organizational unit ID. | 

## Methods

### NewConnectRestoreAwsOrganizationalUnitRequest

`func NewConnectRestoreAwsOrganizationalUnitRequest(roleArn string, providerOrganizationalUnitId string, ) *ConnectRestoreAwsOrganizationalUnitRequest`

NewConnectRestoreAwsOrganizationalUnitRequest instantiates a new ConnectRestoreAwsOrganizationalUnitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectRestoreAwsOrganizationalUnitRequestWithDefaults

`func NewConnectRestoreAwsOrganizationalUnitRequestWithDefaults() *ConnectRestoreAwsOrganizationalUnitRequest`

NewConnectRestoreAwsOrganizationalUnitRequestWithDefaults instantiates a new ConnectRestoreAwsOrganizationalUnitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetProviderOrganizationalUnitId

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) GetProviderOrganizationalUnitId() string`

GetProviderOrganizationalUnitId returns the ProviderOrganizationalUnitId field if non-nil, zero value otherwise.

### GetProviderOrganizationalUnitIdOk

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) GetProviderOrganizationalUnitIdOk() (*string, bool)`

GetProviderOrganizationalUnitIdOk returns a tuple with the ProviderOrganizationalUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrganizationalUnitId

`func (o *ConnectRestoreAwsOrganizationalUnitRequest) SetProviderOrganizationalUnitId(v string)`

SetProviderOrganizationalUnitId sets ProviderOrganizationalUnitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


