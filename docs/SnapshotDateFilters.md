# SnapshotDateFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Matches if &#x60;pointInTime&#x60; is on or later than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 
**EndDate** | Pointer to **string** | Matches if &#x60;pointInTime&#x60; is on or earlier than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 

## Methods

### NewSnapshotDateFilters

`func NewSnapshotDateFilters() *SnapshotDateFilters`

NewSnapshotDateFilters instantiates a new SnapshotDateFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotDateFiltersWithDefaults

`func NewSnapshotDateFiltersWithDefaults() *SnapshotDateFilters`

NewSnapshotDateFiltersWithDefaults instantiates a new SnapshotDateFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *SnapshotDateFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SnapshotDateFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SnapshotDateFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SnapshotDateFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SnapshotDateFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SnapshotDateFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SnapshotDateFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SnapshotDateFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


