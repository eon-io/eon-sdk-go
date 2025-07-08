# ResourceNameCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**ResourceNames** | **[]string** |  | 

## Methods

### NewResourceNameCondition

`func NewResourceNameCondition(operator ScalarOperators, resourceNames []string, ) *ResourceNameCondition`

NewResourceNameCondition instantiates a new ResourceNameCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNameConditionWithDefaults

`func NewResourceNameConditionWithDefaults() *ResourceNameCondition`

NewResourceNameConditionWithDefaults instantiates a new ResourceNameCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ResourceNameCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ResourceNameCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ResourceNameCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetResourceNames

`func (o *ResourceNameCondition) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *ResourceNameCondition) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *ResourceNameCondition) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


