# ListBackupPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**BackupPolicyFilterConditions**](BackupPolicyFilterConditions.md) |  | [optional] 

## Methods

### NewListBackupPoliciesRequest

`func NewListBackupPoliciesRequest() *ListBackupPoliciesRequest`

NewListBackupPoliciesRequest instantiates a new ListBackupPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupPoliciesRequestWithDefaults

`func NewListBackupPoliciesRequestWithDefaults() *ListBackupPoliciesRequest`

NewListBackupPoliciesRequestWithDefaults instantiates a new ListBackupPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ListBackupPoliciesRequest) GetFilters() BackupPolicyFilterConditions`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListBackupPoliciesRequest) GetFiltersOk() (*BackupPolicyFilterConditions, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListBackupPoliciesRequest) SetFilters(v BackupPolicyFilterConditions)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListBackupPoliciesRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


