# ListBackupPoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicies** | Pointer to [**[]BackupPolicy**](BackupPolicy.md) |  | [optional] 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of backup jobs that matched the filter options. | [optional] 

## Methods

### NewListBackupPoliciesResponse

`func NewListBackupPoliciesResponse() *ListBackupPoliciesResponse`

NewListBackupPoliciesResponse instantiates a new ListBackupPoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupPoliciesResponseWithDefaults

`func NewListBackupPoliciesResponseWithDefaults() *ListBackupPoliciesResponse`

NewListBackupPoliciesResponseWithDefaults instantiates a new ListBackupPoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicies

`func (o *ListBackupPoliciesResponse) GetBackupPolicies() []BackupPolicy`

GetBackupPolicies returns the BackupPolicies field if non-nil, zero value otherwise.

### GetBackupPoliciesOk

`func (o *ListBackupPoliciesResponse) GetBackupPoliciesOk() (*[]BackupPolicy, bool)`

GetBackupPoliciesOk returns a tuple with the BackupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicies

`func (o *ListBackupPoliciesResponse) SetBackupPolicies(v []BackupPolicy)`

SetBackupPolicies sets BackupPolicies field to given value.

### HasBackupPolicies

`func (o *ListBackupPoliciesResponse) HasBackupPolicies() bool`

HasBackupPolicies returns a boolean if a field has been set.

### GetNextToken

`func (o *ListBackupPoliciesResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListBackupPoliciesResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListBackupPoliciesResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListBackupPoliciesResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListBackupPoliciesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListBackupPoliciesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListBackupPoliciesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListBackupPoliciesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


