# DynamoDbIndexThroughput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadCapacityUnits** | Pointer to **NullableInt64** | Maximum number of strongly consistent reads consumed per second before DynamoDB returns a ThrottlingException. | [optional] 
**WriteCapacityUnits** | Pointer to **NullableInt64** | Maximum number of writes consumed per second before DynamoDB returns a ThrottlingException. | [optional] 

## Methods

### NewDynamoDbIndexThroughput

`func NewDynamoDbIndexThroughput() *DynamoDbIndexThroughput`

NewDynamoDbIndexThroughput instantiates a new DynamoDbIndexThroughput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDbIndexThroughputWithDefaults

`func NewDynamoDbIndexThroughputWithDefaults() *DynamoDbIndexThroughput`

NewDynamoDbIndexThroughputWithDefaults instantiates a new DynamoDbIndexThroughput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadCapacityUnits

`func (o *DynamoDbIndexThroughput) GetReadCapacityUnits() int64`

GetReadCapacityUnits returns the ReadCapacityUnits field if non-nil, zero value otherwise.

### GetReadCapacityUnitsOk

`func (o *DynamoDbIndexThroughput) GetReadCapacityUnitsOk() (*int64, bool)`

GetReadCapacityUnitsOk returns a tuple with the ReadCapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCapacityUnits

`func (o *DynamoDbIndexThroughput) SetReadCapacityUnits(v int64)`

SetReadCapacityUnits sets ReadCapacityUnits field to given value.

### HasReadCapacityUnits

`func (o *DynamoDbIndexThroughput) HasReadCapacityUnits() bool`

HasReadCapacityUnits returns a boolean if a field has been set.

### SetReadCapacityUnitsNil

`func (o *DynamoDbIndexThroughput) SetReadCapacityUnitsNil(b bool)`

 SetReadCapacityUnitsNil sets the value for ReadCapacityUnits to be an explicit nil

### UnsetReadCapacityUnits
`func (o *DynamoDbIndexThroughput) UnsetReadCapacityUnits()`

UnsetReadCapacityUnits ensures that no value is present for ReadCapacityUnits, not even an explicit nil
### GetWriteCapacityUnits

`func (o *DynamoDbIndexThroughput) GetWriteCapacityUnits() int64`

GetWriteCapacityUnits returns the WriteCapacityUnits field if non-nil, zero value otherwise.

### GetWriteCapacityUnitsOk

`func (o *DynamoDbIndexThroughput) GetWriteCapacityUnitsOk() (*int64, bool)`

GetWriteCapacityUnitsOk returns a tuple with the WriteCapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCapacityUnits

`func (o *DynamoDbIndexThroughput) SetWriteCapacityUnits(v int64)`

SetWriteCapacityUnits sets WriteCapacityUnits field to given value.

### HasWriteCapacityUnits

`func (o *DynamoDbIndexThroughput) HasWriteCapacityUnits() bool`

HasWriteCapacityUnits returns a boolean if a field has been set.

### SetWriteCapacityUnitsNil

`func (o *DynamoDbIndexThroughput) SetWriteCapacityUnitsNil(b bool)`

 SetWriteCapacityUnitsNil sets the value for WriteCapacityUnits to be an explicit nil

### UnsetWriteCapacityUnits
`func (o *DynamoDbIndexThroughput) UnsetWriteCapacityUnits()`

UnsetWriteCapacityUnits ensures that no value is present for WriteCapacityUnits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


