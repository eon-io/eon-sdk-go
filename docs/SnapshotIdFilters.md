# SnapshotIdFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches if any string in this list equals &#x60;snapshotId&#x60;. | [optional] 
**NotIn** | Pointer to **[]string** | Matches if no string in this list equals &#x60;snapshotId&#x60;. | [optional] 

## Methods

### NewSnapshotIdFilters

`func NewSnapshotIdFilters() *SnapshotIdFilters`

NewSnapshotIdFilters instantiates a new SnapshotIdFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotIdFiltersWithDefaults

`func NewSnapshotIdFiltersWithDefaults() *SnapshotIdFilters`

NewSnapshotIdFiltersWithDefaults instantiates a new SnapshotIdFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *SnapshotIdFilters) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *SnapshotIdFilters) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *SnapshotIdFilters) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *SnapshotIdFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *SnapshotIdFilters) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *SnapshotIdFilters) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *SnapshotIdFilters) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *SnapshotIdFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


