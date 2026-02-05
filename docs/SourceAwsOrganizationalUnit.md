# SourceAwsOrganizationalUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon-assigned organizational unit ID. | 
**RoleArn** | **string** | ARN of the role Eon assumes to access the organizational unit in AWS. | 
**Name** | **string** | The organizational unit display name in Eon. | 
**ProviderOrganizationalUnitId** | **string** | AWS Organizational Unit ID. | 
**ProviderManagementAccountId** | **string** | AWS Organization management account ID. | 
**Status** | [**AccountState**](AccountState.md) |  | 

## Methods

### NewSourceAwsOrganizationalUnit

`func NewSourceAwsOrganizationalUnit(id string, roleArn string, name string, providerOrganizationalUnitId string, providerManagementAccountId string, status AccountState, ) *SourceAwsOrganizationalUnit`

NewSourceAwsOrganizationalUnit instantiates a new SourceAwsOrganizationalUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAwsOrganizationalUnitWithDefaults

`func NewSourceAwsOrganizationalUnitWithDefaults() *SourceAwsOrganizationalUnit`

NewSourceAwsOrganizationalUnitWithDefaults instantiates a new SourceAwsOrganizationalUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceAwsOrganizationalUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceAwsOrganizationalUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceAwsOrganizationalUnit) SetId(v string)`

SetId sets Id field to given value.


### GetRoleArn

`func (o *SourceAwsOrganizationalUnit) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SourceAwsOrganizationalUnit) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SourceAwsOrganizationalUnit) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetName

`func (o *SourceAwsOrganizationalUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceAwsOrganizationalUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceAwsOrganizationalUnit) SetName(v string)`

SetName sets Name field to given value.


### GetProviderOrganizationalUnitId

`func (o *SourceAwsOrganizationalUnit) GetProviderOrganizationalUnitId() string`

GetProviderOrganizationalUnitId returns the ProviderOrganizationalUnitId field if non-nil, zero value otherwise.

### GetProviderOrganizationalUnitIdOk

`func (o *SourceAwsOrganizationalUnit) GetProviderOrganizationalUnitIdOk() (*string, bool)`

GetProviderOrganizationalUnitIdOk returns a tuple with the ProviderOrganizationalUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOrganizationalUnitId

`func (o *SourceAwsOrganizationalUnit) SetProviderOrganizationalUnitId(v string)`

SetProviderOrganizationalUnitId sets ProviderOrganizationalUnitId field to given value.


### GetProviderManagementAccountId

`func (o *SourceAwsOrganizationalUnit) GetProviderManagementAccountId() string`

GetProviderManagementAccountId returns the ProviderManagementAccountId field if non-nil, zero value otherwise.

### GetProviderManagementAccountIdOk

`func (o *SourceAwsOrganizationalUnit) GetProviderManagementAccountIdOk() (*string, bool)`

GetProviderManagementAccountIdOk returns a tuple with the ProviderManagementAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderManagementAccountId

`func (o *SourceAwsOrganizationalUnit) SetProviderManagementAccountId(v string)`

SetProviderManagementAccountId sets ProviderManagementAccountId field to given value.


### GetStatus

`func (o *SourceAwsOrganizationalUnit) GetStatus() AccountState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceAwsOrganizationalUnit) GetStatusOk() (*AccountState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceAwsOrganizationalUnit) SetStatus(v AccountState)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


