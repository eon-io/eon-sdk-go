# RestoreJobTypeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]RestoreJobType**](RestoreJobType.md) | Matches if any item in this list equals &#x60;restoreType&#x60;. | [optional] 
**NotIn** | Pointer to [**[]RestoreJobType**](RestoreJobType.md) | Matches if no item in this list equals &#x60;restoreType&#x60;. | [optional] 

## Methods

### NewRestoreJobTypeFilters

`func NewRestoreJobTypeFilters() *RestoreJobTypeFilters`

NewRestoreJobTypeFilters instantiates a new RestoreJobTypeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobTypeFiltersWithDefaults

`func NewRestoreJobTypeFiltersWithDefaults() *RestoreJobTypeFilters`

NewRestoreJobTypeFiltersWithDefaults instantiates a new RestoreJobTypeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *RestoreJobTypeFilters) GetIn() []RestoreJobType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *RestoreJobTypeFilters) GetInOk() (*[]RestoreJobType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *RestoreJobTypeFilters) SetIn(v []RestoreJobType)`

SetIn sets In field to given value.

### HasIn

`func (o *RestoreJobTypeFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *RestoreJobTypeFilters) GetNotIn() []RestoreJobType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *RestoreJobTypeFilters) GetNotInOk() (*[]RestoreJobType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *RestoreJobTypeFilters) SetNotIn(v []RestoreJobType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *RestoreJobTypeFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


