# RestoreAwsOrganizationalUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned organizational unit ID. | 
**RoleArn** | **string** | ARN of the role Eon assumes to access the organizational unit in AWS. | 
**Name** | **string** | Organizational unit display name in Eon. | 
**ProviderOrganizationalUnitId** | **string** | AWS-assigned organizational unit ID. | 
**ProviderManagementAccountId** | **string** | AWS-assigned ID of the organization&#39;s management account. | 
**Status** | [**AccountState**](AccountState.md) |  | 

## Methods

### NewRestoreAwsOrganizationalUnit

`func NewRestoreAwsOrganizationalUnit(id string, roleArn string, name string, providerOrganizationalUnitId string, providerManagementAccountId string, status AccountState, ) *RestoreAwsOrganizationalUnit`

NewRestoreAwsOrganizationalUnit instantiates a new RestoreAwsOrganizationalUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAwsOrganizationalUnitWithDefaults

`func NewRestoreAwsOrganizationalUnitWithDefaults() *RestoreAwsOrganizationalUnit`

NewRestoreAwsOrganizationalUnitWithDefaults instantiates a new RestoreAwsOrganizationalUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreAwsOrganizationalUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreAwsOrganizationalUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreAwsOrganizationalUnit) SetId(v string)`

SetId sets Id field to given value.


### GetRoleArn

`func (o *RestoreAwsOrganizationalUnit) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *RestoreAwsOrganizationalUnit) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *RestoreAwsOrganizationalUnit) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetName

`func (o *RestoreAwsOrganizationalUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreAwsOrganizationalUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreAwsOrganizationalUnit) SetName(v string)`

SetName sets Name field to given value.


### GetProviderOrganizationalUnitId

`func (o *RestoreAwsOrganizationalUnit) GetProviderOrganizationalUnitId() string`

GetProviderOrganizationalUnitId returns the ProviderOrganizationalUnitId field if non-nil, zero value otherwise.

### GetProviderOrganizationalUnitIdOk

`func (o *RestoreAwsOrganizationalUnit) GetProviderOrganizationalUnitIdOk() (*string, bool)`

GetProviderOrganizationalUnitIdOk returns a tuple with the ProviderOrganizationalUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrganizationalUnitId

`func (o *RestoreAwsOrganizationalUnit) SetProviderOrganizationalUnitId(v string)`

SetProviderOrganizationalUnitId sets ProviderOrganizationalUnitId field to given value.


### GetProviderManagementAccountId

`func (o *RestoreAwsOrganizationalUnit) GetProviderManagementAccountId() string`

GetProviderManagementAccountId returns the ProviderManagementAccountId field if non-nil, zero value otherwise.

### GetProviderManagementAccountIdOk

`func (o *RestoreAwsOrganizationalUnit) GetProviderManagementAccountIdOk() (*string, bool)`

GetProviderManagementAccountIdOk returns a tuple with the ProviderManagementAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderManagementAccountId

`func (o *RestoreAwsOrganizationalUnit) SetProviderManagementAccountId(v string)`

SetProviderManagementAccountId sets ProviderManagementAccountId field to given value.


### GetStatus

`func (o *RestoreAwsOrganizationalUnit) GetStatus() AccountState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreAwsOrganizationalUnit) GetStatusOk() (*AccountState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreAwsOrganizationalUnit) SetStatus(v AccountState)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


