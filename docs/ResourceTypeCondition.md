# ResourceTypeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**ResourceTypes** | [**[]ResourceType**](ResourceType.md) |  | 

## Methods

### NewResourceTypeCondition

`func NewResourceTypeCondition(operator ScalarOperators, resourceTypes []ResourceType, ) *ResourceTypeCondition`

NewResourceTypeCondition instantiates a new ResourceTypeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeConditionWithDefaults

`func NewResourceTypeConditionWithDefaults() *ResourceTypeCondition`

NewResourceTypeConditionWithDefaults instantiates a new ResourceTypeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ResourceTypeCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ResourceTypeCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ResourceTypeCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetResourceTypes

`func (o *ResourceTypeCondition) GetResourceTypes() []ResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *ResourceTypeCondition) GetResourceTypesOk() (*[]ResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *ResourceTypeCondition) SetResourceTypes(v []ResourceType)`

SetResourceTypes sets ResourceTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


