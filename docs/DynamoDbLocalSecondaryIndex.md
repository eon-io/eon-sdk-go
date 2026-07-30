# DynamoDbLocalSecondaryIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexName** | Pointer to **string** | Name of the local secondary index. | [optional] 
**KeySchema** | Pointer to [**[]DynamoDbIndexKeySchemaElement**](DynamoDbIndexKeySchemaElement.md) | Complete key schema for the local secondary index. | [optional] 
**Projection** | Pointer to [**NullableDynamoDbIndexProjection**](DynamoDbIndexProjection.md) |  | [optional] 
**IndexSizeBytes** | Pointer to **NullableInt64** | Total size of the index, in bytes. | [optional] 

## Methods

### NewDynamoDbLocalSecondaryIndex

`func NewDynamoDbLocalSecondaryIndex() *DynamoDbLocalSecondaryIndex`

NewDynamoDbLocalSecondaryIndex instantiates a new DynamoDbLocalSecondaryIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDbLocalSecondaryIndexWithDefaults

`func NewDynamoDbLocalSecondaryIndexWithDefaults() *DynamoDbLocalSecondaryIndex`

NewDynamoDbLocalSecondaryIndexWithDefaults instantiates a new DynamoDbLocalSecondaryIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexName

`func (o *DynamoDbLocalSecondaryIndex) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *DynamoDbLocalSecondaryIndex) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *DynamoDbLocalSecondaryIndex) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *DynamoDbLocalSecondaryIndex) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetKeySchema

`func (o *DynamoDbLocalSecondaryIndex) GetKeySchema() []DynamoDbIndexKeySchemaElement`

GetKeySchema returns the KeySchema field if non-nil, zero value otherwise.

### GetKeySchemaOk

`func (o *DynamoDbLocalSecondaryIndex) GetKeySchemaOk() (*[]DynamoDbIndexKeySchemaElement, bool)`

GetKeySchemaOk returns a tuple with the KeySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySchema

`func (o *DynamoDbLocalSecondaryIndex) SetKeySchema(v []DynamoDbIndexKeySchemaElement)`

SetKeySchema sets KeySchema field to given value.

### HasKeySchema

`func (o *DynamoDbLocalSecondaryIndex) HasKeySchema() bool`

HasKeySchema returns a boolean if a field has been set.

### GetProjection

`func (o *DynamoDbLocalSecondaryIndex) GetProjection() DynamoDbIndexProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *DynamoDbLocalSecondaryIndex) GetProjectionOk() (*DynamoDbIndexProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *DynamoDbLocalSecondaryIndex) SetProjection(v DynamoDbIndexProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *DynamoDbLocalSecondaryIndex) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### SetProjectionNil

`func (o *DynamoDbLocalSecondaryIndex) SetProjectionNil(b bool)`

 SetProjectionNil sets the value for Projection to be an explicit nil

### UnsetProjection
`func (o *DynamoDbLocalSecondaryIndex) UnsetProjection()`

UnsetProjection ensures that no value is present for Projection, not even an explicit nil
### GetIndexSizeBytes

`func (o *DynamoDbLocalSecondaryIndex) GetIndexSizeBytes() int64`

GetIndexSizeBytes returns the IndexSizeBytes field if non-nil, zero value otherwise.

### GetIndexSizeBytesOk

`func (o *DynamoDbLocalSecondaryIndex) GetIndexSizeBytesOk() (*int64, bool)`

GetIndexSizeBytesOk returns a tuple with the IndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexSizeBytes

`func (o *DynamoDbLocalSecondaryIndex) SetIndexSizeBytes(v int64)`

SetIndexSizeBytes sets IndexSizeBytes field to given value.

### HasIndexSizeBytes

`func (o *DynamoDbLocalSecondaryIndex) HasIndexSizeBytes() bool`

HasIndexSizeBytes returns a boolean if a field has been set.

### SetIndexSizeBytesNil

`func (o *DynamoDbLocalSecondaryIndex) SetIndexSizeBytesNil(b bool)`

 SetIndexSizeBytesNil sets the value for IndexSizeBytes to be an explicit nil

### UnsetIndexSizeBytes
`func (o *DynamoDbLocalSecondaryIndex) UnsetIndexSizeBytes()`

UnsetIndexSizeBytes ensures that no value is present for IndexSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


