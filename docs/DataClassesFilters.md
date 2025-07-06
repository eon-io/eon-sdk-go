# DataClassesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]DataClass**](DataClass.md) | Matches if any value in this list equals &#x60;dataClasses&#x60;. | [optional] 
**NotIn** | Pointer to [**[]DataClass**](DataClass.md) | Matches if no value in this list equals &#x60;dataClasses&#x60; list. | [optional] 

## Methods

### NewDataClassesFilters

`func NewDataClassesFilters() *DataClassesFilters`

NewDataClassesFilters instantiates a new DataClassesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataClassesFiltersWithDefaults

`func NewDataClassesFiltersWithDefaults() *DataClassesFilters`

NewDataClassesFiltersWithDefaults instantiates a new DataClassesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *DataClassesFilters) GetIn() []DataClass`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *DataClassesFilters) GetInOk() (*[]DataClass, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *DataClassesFilters) SetIn(v []DataClass)`

SetIn sets In field to given value.

### HasIn

`func (o *DataClassesFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *DataClassesFilters) GetNotIn() []DataClass`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *DataClassesFilters) GetNotInOk() (*[]DataClass, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *DataClassesFilters) SetNotIn(v []DataClass)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *DataClassesFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


