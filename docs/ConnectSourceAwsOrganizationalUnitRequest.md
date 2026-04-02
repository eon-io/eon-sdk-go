# ConnectSourceAwsOrganizationalUnitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleArn** | **string** | ARN of the role Eon assumes to access the organizational unit in AWS. | 
**ProviderOrganizationalUnitId** | **string** | AWS-assigned organizational unit ID. | 
**SelectedAccountIds** | Pointer to **[]string** | Optional list of AWS account IDs to selectively onboard. When provided, only these accounts within the OU will be discovered and backed up. When omitted, all accounts in the OU are onboarded (legacy behavior).  | [optional] 

## Methods

### NewConnectSourceAwsOrganizationalUnitRequest

`func NewConnectSourceAwsOrganizationalUnitRequest(roleArn string, providerOrganizationalUnitId string, ) *ConnectSourceAwsOrganizationalUnitRequest`

NewConnectSourceAwsOrganizationalUnitRequest instantiates a new ConnectSourceAwsOrganizationalUnitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectSourceAwsOrganizationalUnitRequestWithDefaults

`func NewConnectSourceAwsOrganizationalUnitRequestWithDefaults() *ConnectSourceAwsOrganizationalUnitRequest`

NewConnectSourceAwsOrganizationalUnitRequestWithDefaults instantiates a new ConnectSourceAwsOrganizationalUnitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleArn

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *ConnectSourceAwsOrganizationalUnitRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetProviderOrganizationalUnitId

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetProviderOrganizationalUnitId() string`

GetProviderOrganizationalUnitId returns the ProviderOrganizationalUnitId field if non-nil, zero value otherwise.

### GetProviderOrganizationalUnitIdOk

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetProviderOrganizationalUnitIdOk() (*string, bool)`

GetProviderOrganizationalUnitIdOk returns a tuple with the ProviderOrganizationalUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrganizationalUnitId

`func (o *ConnectSourceAwsOrganizationalUnitRequest) SetProviderOrganizationalUnitId(v string)`

SetProviderOrganizationalUnitId sets ProviderOrganizationalUnitId field to given value.


### GetSelectedAccountIds

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetSelectedAccountIds() []string`

GetSelectedAccountIds returns the SelectedAccountIds field if non-nil, zero value otherwise.

### GetSelectedAccountIdsOk

`func (o *ConnectSourceAwsOrganizationalUnitRequest) GetSelectedAccountIdsOk() (*[]string, bool)`

GetSelectedAccountIdsOk returns a tuple with the SelectedAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedAccountIds

`func (o *ConnectSourceAwsOrganizationalUnitRequest) SetSelectedAccountIds(v []string)`

SetSelectedAccountIds sets SelectedAccountIds field to given value.

### HasSelectedAccountIds

`func (o *ConnectSourceAwsOrganizationalUnitRequest) HasSelectedAccountIds() bool`

HasSelectedAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


