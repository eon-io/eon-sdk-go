# DynamoDbGlobalSecondaryIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexName** | Pointer to **string** | Name of the global secondary index. | [optional] 
**KeySchema** | Pointer to [**[]DynamoDbIndexKeySchemaElement**](DynamoDbIndexKeySchemaElement.md) | Complete key schema for the global secondary index. | [optional] 
**Projection** | Pointer to [**NullableDynamoDbIndexProjection**](DynamoDbIndexProjection.md) |  | [optional] 
**OnDemandThroughput** | Pointer to [**NullableDynamoDbIndexThroughput**](DynamoDbIndexThroughput.md) |  | [optional] 
**ProvisionedThroughput** | Pointer to [**NullableDynamoDbIndexThroughput**](DynamoDbIndexThroughput.md) |  | [optional] 
**WarmThroughput** | Pointer to [**NullableDynamoDbIndexThroughput**](DynamoDbIndexThroughput.md) |  | [optional] 
**IndexSizeBytes** | Pointer to **NullableInt64** | Total size of the index, in bytes. | [optional] 

## Methods

### NewDynamoDbGlobalSecondaryIndex

`func NewDynamoDbGlobalSecondaryIndex() *DynamoDbGlobalSecondaryIndex`

NewDynamoDbGlobalSecondaryIndex instantiates a new DynamoDbGlobalSecondaryIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDbGlobalSecondaryIndexWithDefaults

`func NewDynamoDbGlobalSecondaryIndexWithDefaults() *DynamoDbGlobalSecondaryIndex`

NewDynamoDbGlobalSecondaryIndexWithDefaults instantiates a new DynamoDbGlobalSecondaryIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexName

`func (o *DynamoDbGlobalSecondaryIndex) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *DynamoDbGlobalSecondaryIndex) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *DynamoDbGlobalSecondaryIndex) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *DynamoDbGlobalSecondaryIndex) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetKeySchema

`func (o *DynamoDbGlobalSecondaryIndex) GetKeySchema() []DynamoDbIndexKeySchemaElement`

GetKeySchema returns the KeySchema field if non-nil, zero value otherwise.

### GetKeySchemaOk

`func (o *DynamoDbGlobalSecondaryIndex) GetKeySchemaOk() (*[]DynamoDbIndexKeySchemaElement, bool)`

GetKeySchemaOk returns a tuple with the KeySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySchema

`func (o *DynamoDbGlobalSecondaryIndex) SetKeySchema(v []DynamoDbIndexKeySchemaElement)`

SetKeySchema sets KeySchema field to given value.

### HasKeySchema

`func (o *DynamoDbGlobalSecondaryIndex) HasKeySchema() bool`

HasKeySchema returns a boolean if a field has been set.

### GetProjection

`func (o *DynamoDbGlobalSecondaryIndex) GetProjection() DynamoDbIndexProjection`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *DynamoDbGlobalSecondaryIndex) GetProjectionOk() (*DynamoDbIndexProjection, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *DynamoDbGlobalSecondaryIndex) SetProjection(v DynamoDbIndexProjection)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *DynamoDbGlobalSecondaryIndex) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### SetProjectionNil

`func (o *DynamoDbGlobalSecondaryIndex) SetProjectionNil(b bool)`

 SetProjectionNil sets the value for Projection to be an explicit nil

### UnsetProjection
`func (o *DynamoDbGlobalSecondaryIndex) UnsetProjection()`

UnsetProjection ensures that no value is present for Projection, not even an explicit nil
### GetOnDemandThroughput

`func (o *DynamoDbGlobalSecondaryIndex) GetOnDemandThroughput() DynamoDbIndexThroughput`

GetOnDemandThroughput returns the OnDemandThroughput field if non-nil, zero value otherwise.

### GetOnDemandThroughputOk

`func (o *DynamoDbGlobalSecondaryIndex) GetOnDemandThroughputOk() (*DynamoDbIndexThroughput, bool)`

GetOnDemandThroughputOk returns a tuple with the OnDemandThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandThroughput

`func (o *DynamoDbGlobalSecondaryIndex) SetOnDemandThroughput(v DynamoDbIndexThroughput)`

SetOnDemandThroughput sets OnDemandThroughput field to given value.

### HasOnDemandThroughput

`func (o *DynamoDbGlobalSecondaryIndex) HasOnDemandThroughput() bool`

HasOnDemandThroughput returns a boolean if a field has been set.

### SetOnDemandThroughputNil

`func (o *DynamoDbGlobalSecondaryIndex) SetOnDemandThroughputNil(b bool)`

 SetOnDemandThroughputNil sets the value for OnDemandThroughput to be an explicit nil

### UnsetOnDemandThroughput
`func (o *DynamoDbGlobalSecondaryIndex) UnsetOnDemandThroughput()`

UnsetOnDemandThroughput ensures that no value is present for OnDemandThroughput, not even an explicit nil
### GetProvisionedThroughput

`func (o *DynamoDbGlobalSecondaryIndex) GetProvisionedThroughput() DynamoDbIndexThroughput`

GetProvisionedThroughput returns the ProvisionedThroughput field if non-nil, zero value otherwise.

### GetProvisionedThroughputOk

`func (o *DynamoDbGlobalSecondaryIndex) GetProvisionedThroughputOk() (*DynamoDbIndexThroughput, bool)`

GetProvisionedThroughputOk returns a tuple with the ProvisionedThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedThroughput

`func (o *DynamoDbGlobalSecondaryIndex) SetProvisionedThroughput(v DynamoDbIndexThroughput)`

SetProvisionedThroughput sets ProvisionedThroughput field to given value.

### HasProvisionedThroughput

`func (o *DynamoDbGlobalSecondaryIndex) HasProvisionedThroughput() bool`

HasProvisionedThroughput returns a boolean if a field has been set.

### SetProvisionedThroughputNil

`func (o *DynamoDbGlobalSecondaryIndex) SetProvisionedThroughputNil(b bool)`

 SetProvisionedThroughputNil sets the value for ProvisionedThroughput to be an explicit nil

### UnsetProvisionedThroughput
`func (o *DynamoDbGlobalSecondaryIndex) UnsetProvisionedThroughput()`

UnsetProvisionedThroughput ensures that no value is present for ProvisionedThroughput, not even an explicit nil
### GetWarmThroughput

`func (o *DynamoDbGlobalSecondaryIndex) GetWarmThroughput() DynamoDbIndexThroughput`

GetWarmThroughput returns the WarmThroughput field if non-nil, zero value otherwise.

### GetWarmThroughputOk

`func (o *DynamoDbGlobalSecondaryIndex) GetWarmThroughputOk() (*DynamoDbIndexThroughput, bool)`

GetWarmThroughputOk returns a tuple with the WarmThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmThroughput

`func (o *DynamoDbGlobalSecondaryIndex) SetWarmThroughput(v DynamoDbIndexThroughput)`

SetWarmThroughput sets WarmThroughput field to given value.

### HasWarmThroughput

`func (o *DynamoDbGlobalSecondaryIndex) HasWarmThroughput() bool`

HasWarmThroughput returns a boolean if a field has been set.

### SetWarmThroughputNil

`func (o *DynamoDbGlobalSecondaryIndex) SetWarmThroughputNil(b bool)`

 SetWarmThroughputNil sets the value for WarmThroughput to be an explicit nil

### UnsetWarmThroughput
`func (o *DynamoDbGlobalSecondaryIndex) UnsetWarmThroughput()`

UnsetWarmThroughput ensures that no value is present for WarmThroughput, not even an explicit nil
### GetIndexSizeBytes

`func (o *DynamoDbGlobalSecondaryIndex) GetIndexSizeBytes() int64`

GetIndexSizeBytes returns the IndexSizeBytes field if non-nil, zero value otherwise.

### GetIndexSizeBytesOk

`func (o *DynamoDbGlobalSecondaryIndex) GetIndexSizeBytesOk() (*int64, bool)`

GetIndexSizeBytesOk returns a tuple with the IndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexSizeBytes

`func (o *DynamoDbGlobalSecondaryIndex) SetIndexSizeBytes(v int64)`

SetIndexSizeBytes sets IndexSizeBytes field to given value.

### HasIndexSizeBytes

`func (o *DynamoDbGlobalSecondaryIndex) HasIndexSizeBytes() bool`

HasIndexSizeBytes returns a boolean if a field has been set.

### SetIndexSizeBytesNil

`func (o *DynamoDbGlobalSecondaryIndex) SetIndexSizeBytesNil(b bool)`

 SetIndexSizeBytesNil sets the value for IndexSizeBytes to be an explicit nil

### UnsetIndexSizeBytes
`func (o *DynamoDbGlobalSecondaryIndex) UnsetIndexSizeBytes()`

UnsetIndexSizeBytes ensures that no value is present for IndexSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


