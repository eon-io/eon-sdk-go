# ResourceGroupNameCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**ResourceGroupNames** | **[]string** |  | 

## Methods

### NewResourceGroupNameCondition

`func NewResourceGroupNameCondition(operator ScalarOperators, resourceGroupNames []string, ) *ResourceGroupNameCondition`

NewResourceGroupNameCondition instantiates a new ResourceGroupNameCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupNameConditionWithDefaults

`func NewResourceGroupNameConditionWithDefaults() *ResourceGroupNameCondition`

NewResourceGroupNameConditionWithDefaults instantiates a new ResourceGroupNameCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ResourceGroupNameCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ResourceGroupNameCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ResourceGroupNameCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetResourceGroupNames

`func (o *ResourceGroupNameCondition) GetResourceGroupNames() []string`

GetResourceGroupNames returns the ResourceGroupNames field if non-nil, zero value otherwise.

### GetResourceGroupNamesOk

`func (o *ResourceGroupNameCondition) GetResourceGroupNamesOk() (*[]string, bool)`

GetResourceGroupNamesOk returns a tuple with the ResourceGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupNames

`func (o *ResourceGroupNameCondition) SetResourceGroupNames(v []string)`

SetResourceGroupNames sets ResourceGroupNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


