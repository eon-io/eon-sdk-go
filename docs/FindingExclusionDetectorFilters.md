# FindingExclusionDetectorFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]FindingExclusionDetectorType**](FindingExclusionDetectorType.md) | Matches exclusions for any of these detectors. An empty list applies no filter. | [optional] 
**NotIn** | Pointer to [**[]FindingExclusionDetectorType**](FindingExclusionDetectorType.md) | Matches exclusions for none of these detectors. An empty list applies no filter. | [optional] 

## Methods

### NewFindingExclusionDetectorFilters

`func NewFindingExclusionDetectorFilters() *FindingExclusionDetectorFilters`

NewFindingExclusionDetectorFilters instantiates a new FindingExclusionDetectorFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionDetectorFiltersWithDefaults

`func NewFindingExclusionDetectorFiltersWithDefaults() *FindingExclusionDetectorFilters`

NewFindingExclusionDetectorFiltersWithDefaults instantiates a new FindingExclusionDetectorFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *FindingExclusionDetectorFilters) GetIn() []FindingExclusionDetectorType`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *FindingExclusionDetectorFilters) GetInOk() (*[]FindingExclusionDetectorType, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *FindingExclusionDetectorFilters) SetIn(v []FindingExclusionDetectorType)`

SetIn sets In field to given value.

### HasIn

`func (o *FindingExclusionDetectorFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *FindingExclusionDetectorFilters) GetNotIn() []FindingExclusionDetectorType`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *FindingExclusionDetectorFilters) GetNotInOk() (*[]FindingExclusionDetectorType, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *FindingExclusionDetectorFilters) SetNotIn(v []FindingExclusionDetectorType)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *FindingExclusionDetectorFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


