# AwsDynamoDbResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingMode** | [**DynamoDbBillingMode**](DynamoDbBillingMode.md) |  | 
**ReadCapacityUnits** | **int64** | Provisioned read capacity units of the table. | 
**WriteCapacityUnits** | **int64** | Provisioned write capacity units of the table. | 
**GlobalSecondaryIndexes** | Pointer to [**[]DynamoDbGlobalSecondaryIndex**](DynamoDbGlobalSecondaryIndex.md) | Global secondary indexes on the table. To minimize response size, not returned by default.  To include in the response, set &#x60;include&#x60; to &#x60;DYNAMODB_INDEXES&#x60; in the request.  | [optional] 
**LocalSecondaryIndexes** | Pointer to [**[]DynamoDbLocalSecondaryIndex**](DynamoDbLocalSecondaryIndex.md) | Local secondary indexes on the table. To minimize response size, not returned by default.  To include in the response, set &#x60;include&#x60; to &#x60;DYNAMODB_INDEXES&#x60; in the request.  | [optional] 

## Methods

### NewAwsDynamoDbResource

`func NewAwsDynamoDbResource(billingMode DynamoDbBillingMode, readCapacityUnits int64, writeCapacityUnits int64, ) *AwsDynamoDbResource`

NewAwsDynamoDbResource instantiates a new AwsDynamoDbResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsDynamoDbResourceWithDefaults

`func NewAwsDynamoDbResourceWithDefaults() *AwsDynamoDbResource`

NewAwsDynamoDbResourceWithDefaults instantiates a new AwsDynamoDbResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingMode

`func (o *AwsDynamoDbResource) GetBillingMode() DynamoDbBillingMode`

GetBillingMode returns the BillingMode field if non-nil, zero value otherwise.

### GetBillingModeOk

`func (o *AwsDynamoDbResource) GetBillingModeOk() (*DynamoDbBillingMode, bool)`

GetBillingModeOk returns a tuple with the BillingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMode

`func (o *AwsDynamoDbResource) SetBillingMode(v DynamoDbBillingMode)`

SetBillingMode sets BillingMode field to given value.


### GetReadCapacityUnits

`func (o *AwsDynamoDbResource) GetReadCapacityUnits() int64`

GetReadCapacityUnits returns the ReadCapacityUnits field if non-nil, zero value otherwise.

### GetReadCapacityUnitsOk

`func (o *AwsDynamoDbResource) GetReadCapacityUnitsOk() (*int64, bool)`

GetReadCapacityUnitsOk returns a tuple with the ReadCapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCapacityUnits

`func (o *AwsDynamoDbResource) SetReadCapacityUnits(v int64)`

SetReadCapacityUnits sets ReadCapacityUnits field to given value.


### GetWriteCapacityUnits

`func (o *AwsDynamoDbResource) GetWriteCapacityUnits() int64`

GetWriteCapacityUnits returns the WriteCapacityUnits field if non-nil, zero value otherwise.

### GetWriteCapacityUnitsOk

`func (o *AwsDynamoDbResource) GetWriteCapacityUnitsOk() (*int64, bool)`

GetWriteCapacityUnitsOk returns a tuple with the WriteCapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCapacityUnits

`func (o *AwsDynamoDbResource) SetWriteCapacityUnits(v int64)`

SetWriteCapacityUnits sets WriteCapacityUnits field to given value.


### GetGlobalSecondaryIndexes

`func (o *AwsDynamoDbResource) GetGlobalSecondaryIndexes() []DynamoDbGlobalSecondaryIndex`

GetGlobalSecondaryIndexes returns the GlobalSecondaryIndexes field if non-nil, zero value otherwise.

### GetGlobalSecondaryIndexesOk

`func (o *AwsDynamoDbResource) GetGlobalSecondaryIndexesOk() (*[]DynamoDbGlobalSecondaryIndex, bool)`

GetGlobalSecondaryIndexesOk returns a tuple with the GlobalSecondaryIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSecondaryIndexes

`func (o *AwsDynamoDbResource) SetGlobalSecondaryIndexes(v []DynamoDbGlobalSecondaryIndex)`

SetGlobalSecondaryIndexes sets GlobalSecondaryIndexes field to given value.

### HasGlobalSecondaryIndexes

`func (o *AwsDynamoDbResource) HasGlobalSecondaryIndexes() bool`

HasGlobalSecondaryIndexes returns a boolean if a field has been set.

### GetLocalSecondaryIndexes

`func (o *AwsDynamoDbResource) GetLocalSecondaryIndexes() []DynamoDbLocalSecondaryIndex`

GetLocalSecondaryIndexes returns the LocalSecondaryIndexes field if non-nil, zero value otherwise.

### GetLocalSecondaryIndexesOk

`func (o *AwsDynamoDbResource) GetLocalSecondaryIndexesOk() (*[]DynamoDbLocalSecondaryIndex, bool)`

GetLocalSecondaryIndexesOk returns a tuple with the LocalSecondaryIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecondaryIndexes

`func (o *AwsDynamoDbResource) SetLocalSecondaryIndexes(v []DynamoDbLocalSecondaryIndex)`

SetLocalSecondaryIndexes sets LocalSecondaryIndexes field to given value.

### HasLocalSecondaryIndexes

`func (o *AwsDynamoDbResource) HasLocalSecondaryIndexes() bool`

HasLocalSecondaryIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


