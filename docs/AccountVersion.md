# AccountVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installed** | **string** | Currently installed permissions version. | 
**InstallationMethod** | Pointer to [**AccountInstallationMethod**](AccountInstallationMethod.md) |  | [optional] [default to ACCOUNT_INSTALLATION_METHOD_UNSPECIFIED]
**Latest** | Pointer to **string** | Latest available permissions version. | [optional] 

## Methods

### NewAccountVersion

`func NewAccountVersion(installed string, ) *AccountVersion`

NewAccountVersion instantiates a new AccountVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountVersionWithDefaults

`func NewAccountVersionWithDefaults() *AccountVersion`

NewAccountVersionWithDefaults instantiates a new AccountVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalled

`func (o *AccountVersion) GetInstalled() string`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *AccountVersion) GetInstalledOk() (*string, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *AccountVersion) SetInstalled(v string)`

SetInstalled sets Installed field to given value.


### GetInstallationMethod

`func (o *AccountVersion) GetInstallationMethod() AccountInstallationMethod`

GetInstallationMethod returns the InstallationMethod field if non-nil, zero value otherwise.

### GetInstallationMethodOk

`func (o *AccountVersion) GetInstallationMethodOk() (*AccountInstallationMethod, bool)`

GetInstallationMethodOk returns a tuple with the InstallationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationMethod

`func (o *AccountVersion) SetInstallationMethod(v AccountInstallationMethod)`

SetInstallationMethod sets InstallationMethod field to given value.

### HasInstallationMethod

`func (o *AccountVersion) HasInstallationMethod() bool`

HasInstallationMethod returns a boolean if a field has been set.

### GetLatest

`func (o *AccountVersion) GetLatest() string`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *AccountVersion) GetLatestOk() (*string, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *AccountVersion) SetLatest(v string)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *AccountVersion) HasLatest() bool`

HasLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


