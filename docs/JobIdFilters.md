# JobIdFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches if any string in this list equals &#x60;jobId&#x60;. | [optional] 
**NotIn** | Pointer to **[]string** | Matches if no string in this list equals &#x60;jobId&#x60;. | [optional] 

## Methods

### NewJobIdFilters

`func NewJobIdFilters() *JobIdFilters`

NewJobIdFilters instantiates a new JobIdFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobIdFiltersWithDefaults

`func NewJobIdFiltersWithDefaults() *JobIdFilters`

NewJobIdFiltersWithDefaults instantiates a new JobIdFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *JobIdFilters) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *JobIdFilters) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *JobIdFilters) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *JobIdFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *JobIdFilters) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *JobIdFilters) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *JobIdFilters) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *JobIdFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


