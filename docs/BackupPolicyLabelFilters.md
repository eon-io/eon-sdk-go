# BackupPolicyLabelFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainsAnyOf** | Pointer to **[]string** | Returns backup policies that have at least one of these labels. | [optional] 
**ContainsAllOf** | Pointer to **[]string** | Returns backup policies that have all of these labels. | [optional] 
**ContainsNoneOf** | Pointer to **[]string** | Returns backup policies that have none of these labels. | [optional] 

## Methods

### NewBackupPolicyLabelFilters

`func NewBackupPolicyLabelFilters() *BackupPolicyLabelFilters`

NewBackupPolicyLabelFilters instantiates a new BackupPolicyLabelFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyLabelFiltersWithDefaults

`func NewBackupPolicyLabelFiltersWithDefaults() *BackupPolicyLabelFilters`

NewBackupPolicyLabelFiltersWithDefaults instantiates a new BackupPolicyLabelFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainsAnyOf

`func (o *BackupPolicyLabelFilters) GetContainsAnyOf() []string`

GetContainsAnyOf returns the ContainsAnyOf field if non-nil, zero value otherwise.

### GetContainsAnyOfOk

`func (o *BackupPolicyLabelFilters) GetContainsAnyOfOk() (*[]string, bool)`

GetContainsAnyOfOk returns a tuple with the ContainsAnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAnyOf

`func (o *BackupPolicyLabelFilters) SetContainsAnyOf(v []string)`

SetContainsAnyOf sets ContainsAnyOf field to given value.

### HasContainsAnyOf

`func (o *BackupPolicyLabelFilters) HasContainsAnyOf() bool`

HasContainsAnyOf returns a boolean if a field has been set.

### GetContainsAllOf

`func (o *BackupPolicyLabelFilters) GetContainsAllOf() []string`

GetContainsAllOf returns the ContainsAllOf field if non-nil, zero value otherwise.

### GetContainsAllOfOk

`func (o *BackupPolicyLabelFilters) GetContainsAllOfOk() (*[]string, bool)`

GetContainsAllOfOk returns a tuple with the ContainsAllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAllOf

`func (o *BackupPolicyLabelFilters) SetContainsAllOf(v []string)`

SetContainsAllOf sets ContainsAllOf field to given value.

### HasContainsAllOf

`func (o *BackupPolicyLabelFilters) HasContainsAllOf() bool`

HasContainsAllOf returns a boolean if a field has been set.

### GetContainsNoneOf

`func (o *BackupPolicyLabelFilters) GetContainsNoneOf() []string`

GetContainsNoneOf returns the ContainsNoneOf field if non-nil, zero value otherwise.

### GetContainsNoneOfOk

`func (o *BackupPolicyLabelFilters) GetContainsNoneOfOk() (*[]string, bool)`

GetContainsNoneOfOk returns a tuple with the ContainsNoneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsNoneOf

`func (o *BackupPolicyLabelFilters) SetContainsNoneOf(v []string)`

SetContainsNoneOf sets ContainsNoneOf field to given value.

### HasContainsNoneOf

`func (o *BackupPolicyLabelFilters) HasContainsNoneOf() bool`

HasContainsNoneOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


