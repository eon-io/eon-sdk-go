# DataClassesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainsAnyOf** | Pointer to [**[]DataClass**](DataClass.md) | Matches if any value in this list is in the &#x60;dataClasses&#x60; list. | [optional] 
**ContainsNoneOf** | Pointer to [**[]DataClass**](DataClass.md) | Matches if none of the values in this list are in the &#x60;dataClasses&#x60; list. | [optional] 
**ContainsAllOf** | Pointer to **[]string** | Matches if all values in this list are in the &#x60;dataClasses&#x60; list. | [optional] 

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

### GetContainsAnyOf

`func (o *DataClassesFilters) GetContainsAnyOf() []DataClass`

GetContainsAnyOf returns the ContainsAnyOf field if non-nil, zero value otherwise.

### GetContainsAnyOfOk

`func (o *DataClassesFilters) GetContainsAnyOfOk() (*[]DataClass, bool)`

GetContainsAnyOfOk returns a tuple with the ContainsAnyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAnyOf

`func (o *DataClassesFilters) SetContainsAnyOf(v []DataClass)`

SetContainsAnyOf sets ContainsAnyOf field to given value.

### HasContainsAnyOf

`func (o *DataClassesFilters) HasContainsAnyOf() bool`

HasContainsAnyOf returns a boolean if a field has been set.

### GetContainsNoneOf

`func (o *DataClassesFilters) GetContainsNoneOf() []DataClass`

GetContainsNoneOf returns the ContainsNoneOf field if non-nil, zero value otherwise.

### GetContainsNoneOfOk

`func (o *DataClassesFilters) GetContainsNoneOfOk() (*[]DataClass, bool)`

GetContainsNoneOfOk returns a tuple with the ContainsNoneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsNoneOf

`func (o *DataClassesFilters) SetContainsNoneOf(v []DataClass)`

SetContainsNoneOf sets ContainsNoneOf field to given value.

### HasContainsNoneOf

`func (o *DataClassesFilters) HasContainsNoneOf() bool`

HasContainsNoneOf returns a boolean if a field has been set.

### GetContainsAllOf

`func (o *DataClassesFilters) GetContainsAllOf() []string`

GetContainsAllOf returns the ContainsAllOf field if non-nil, zero value otherwise.

### GetContainsAllOfOk

`func (o *DataClassesFilters) GetContainsAllOfOk() (*[]string, bool)`

GetContainsAllOfOk returns a tuple with the ContainsAllOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsAllOf

`func (o *DataClassesFilters) SetContainsAllOf(v []string)`

SetContainsAllOf sets ContainsAllOf field to given value.

### HasContainsAllOf

`func (o *DataClassesFilters) HasContainsAllOf() bool`

HasContainsAllOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


