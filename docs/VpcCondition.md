# VpcCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**Vpcs** | **[]string** |  | 

## Methods

### NewVpcCondition

`func NewVpcCondition(operator ScalarOperators, vpcs []string, ) *VpcCondition`

NewVpcCondition instantiates a new VpcCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcConditionWithDefaults

`func NewVpcConditionWithDefaults() *VpcCondition`

NewVpcConditionWithDefaults instantiates a new VpcCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *VpcCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *VpcCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *VpcCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetVpcs

`func (o *VpcCondition) GetVpcs() []string`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *VpcCondition) GetVpcsOk() (*[]string, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *VpcCondition) SetVpcs(v []string)`

SetVpcs sets Vpcs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


