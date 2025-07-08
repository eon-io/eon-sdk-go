# DataClassesCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ListOperators**](ListOperators.md) |  | 
**DataClasses** | [**[]DataClass**](DataClass.md) |  | 

## Methods

### NewDataClassesCondition

`func NewDataClassesCondition(operator ListOperators, dataClasses []DataClass, ) *DataClassesCondition`

NewDataClassesCondition instantiates a new DataClassesCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataClassesConditionWithDefaults

`func NewDataClassesConditionWithDefaults() *DataClassesCondition`

NewDataClassesConditionWithDefaults instantiates a new DataClassesCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *DataClassesCondition) GetOperator() ListOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DataClassesCondition) GetOperatorOk() (*ListOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DataClassesCondition) SetOperator(v ListOperators)`

SetOperator sets Operator field to given value.


### GetDataClasses

`func (o *DataClassesCondition) GetDataClasses() []DataClass`

GetDataClasses returns the DataClasses field if non-nil, zero value otherwise.

### GetDataClassesOk

`func (o *DataClassesCondition) GetDataClassesOk() (*[]DataClass, bool)`

GetDataClassesOk returns a tuple with the DataClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClasses

`func (o *DataClassesCondition) SetDataClasses(v []DataClass)`

SetDataClasses sets DataClasses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


