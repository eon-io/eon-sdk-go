# ResourceIdCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**ResourceIds** | **[]string** |  | 

## Methods

### NewResourceIdCondition

`func NewResourceIdCondition(operator ScalarOperators, resourceIds []string, ) *ResourceIdCondition`

NewResourceIdCondition instantiates a new ResourceIdCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceIdConditionWithDefaults

`func NewResourceIdConditionWithDefaults() *ResourceIdCondition`

NewResourceIdConditionWithDefaults instantiates a new ResourceIdCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ResourceIdCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ResourceIdCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ResourceIdCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetResourceIds

`func (o *ResourceIdCondition) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *ResourceIdCondition) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *ResourceIdCondition) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


