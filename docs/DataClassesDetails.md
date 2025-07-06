# DataClassesDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataClasses** | Pointer to [**[]DataClass**](DataClass.md) | List of data classes. | [optional] 
**IsOverridden** | Pointer to **bool** | Whether the data classes are manually overridden. If &#x60;true&#x60;, the list of data classes is user-defined and remains static. If &#x60;false&#x60;, the data classes are automatically detected and listed by Eon.  | [optional] 

## Methods

### NewDataClassesDetails

`func NewDataClassesDetails() *DataClassesDetails`

NewDataClassesDetails instantiates a new DataClassesDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataClassesDetailsWithDefaults

`func NewDataClassesDetailsWithDefaults() *DataClassesDetails`

NewDataClassesDetailsWithDefaults instantiates a new DataClassesDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataClasses

`func (o *DataClassesDetails) GetDataClasses() []DataClass`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *DataClassesDetails) GetDataClassesOk() (*[]DataClass, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *DataClassesDetails) SetDataClasses(v []DataClass)`

SetDataClasses sets DataClasses field to given value.

### HasDataClasses

`func (o *DataClassesDetails) HasDataClasses() bool`

HasDataClasses returns a boolean if a field has been set.

### GetIsOverridden

`func (o *DataClassesDetails) GetIsOverridden() bool`

GetIsOverridden returns the IsOverridden field if non-nil, zero value otherwise.

### GetIsOverriddenOk

`func (o *DataClassesDetails) GetIsOverriddenOk() (*bool, bool)`

GetIsOverriddenOk returns a tuple with the IsOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridden

`func (o *DataClassesDetails) SetIsOverridden(v bool)`

SetIsOverridden sets IsOverridden field to given value.

### HasIsOverridden

`func (o *DataClassesDetails) HasIsOverridden() bool`

HasIsOverridden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


