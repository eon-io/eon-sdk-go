# ListBackupVaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vaults** | [**[]BackupVault**](BackupVault.md) |  | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | **int32** | Total number of restore jobs that matched the filter options. | 

## Methods

### NewListBackupVaultResponse

`func NewListBackupVaultResponse(vaults []BackupVault, totalCount int32, ) *ListBackupVaultResponse`

NewListBackupVaultResponse instantiates a new ListBackupVaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupVaultResponseWithDefaults

`func NewListBackupVaultResponseWithDefaults() *ListBackupVaultResponse`

NewListBackupVaultResponseWithDefaults instantiates a new ListBackupVaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaults

`func (o *ListBackupVaultResponse) GetVaults() []BackupVault`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *ListBackupVaultResponse) GetVaultsOk() (*[]BackupVault, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *ListBackupVaultResponse) SetVaults(v []BackupVault)`

SetVaults sets Vaults field to given value.


### GetNextToken

`func (o *ListBackupVaultResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListBackupVaultResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListBackupVaultResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListBackupVaultResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListBackupVaultResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListBackupVaultResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListBackupVaultResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


