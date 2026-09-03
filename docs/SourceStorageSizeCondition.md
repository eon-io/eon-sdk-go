# SourceStorageSizeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**NumericOperators**](NumericOperators.md) |  | 
**SizeBytes** | **int64** |  | 

## Methods

### NewSourceStorageSizeCondition

`func NewSourceStorageSizeCondition(operator NumericOperators, sizeBytes int64, ) *SourceStorageSizeCondition`

NewSourceStorageSizeCondition instantiates a new SourceStorageSizeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceStorageSizeConditionWithDefaults

`func NewSourceStorageSizeConditionWithDefaults() *SourceStorageSizeCondition`

NewSourceStorageSizeConditionWithDefaults instantiates a new SourceStorageSizeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SourceStorageSizeCondition) GetOperator() NumericOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SourceStorageSizeCondition) GetOperatorOk() (*NumericOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SourceStorageSizeCondition) SetOperator(v NumericOperators)`

SetOperator sets Operator field to given value.


### GetSizeBytes

`func (o *SourceStorageSizeCondition) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SourceStorageSizeCondition) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SourceStorageSizeCondition) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


