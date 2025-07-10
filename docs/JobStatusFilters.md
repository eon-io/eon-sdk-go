# JobStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]JobStatus**](JobStatus.md) |  | [optional] 
**NotIn** | Pointer to [**[]JobStatus**](JobStatus.md) |  | [optional] 

## Methods

### NewJobStatusFilters

`func NewJobStatusFilters() *JobStatusFilters`

NewJobStatusFilters instantiates a new JobStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatusFiltersWithDefaults

`func NewJobStatusFiltersWithDefaults() *JobStatusFilters`

NewJobStatusFiltersWithDefaults instantiates a new JobStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *JobStatusFilters) GetIn() []JobStatus`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *JobStatusFilters) GetInOk() (*[]JobStatus, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *JobStatusFilters) SetIn(v []JobStatus)`

SetIn sets In field to given value.

### HasIn

`func (o *JobStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *JobStatusFilters) GetNotIn() []JobStatus`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *JobStatusFilters) GetNotInOk() (*[]JobStatus, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *JobStatusFilters) SetNotIn(v []JobStatus)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *JobStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


