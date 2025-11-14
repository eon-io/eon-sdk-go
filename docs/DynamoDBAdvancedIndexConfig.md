# DynamoDBAdvancedIndexConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSecondaryIndexMode** | Pointer to [**IndexRestoreMode**](IndexRestoreMode.md) |  | [optional] [default to KEEP]
**GlobalSecondaryIndexMode** | Pointer to [**IndexRestoreMode**](IndexRestoreMode.md) |  | [optional] [default to KEEP]

## Methods

### NewDynamoDBAdvancedIndexConfig

`func NewDynamoDBAdvancedIndexConfig() *DynamoDBAdvancedIndexConfig`

NewDynamoDBAdvancedIndexConfig instantiates a new DynamoDBAdvancedIndexConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDBAdvancedIndexConfigWithDefaults

`func NewDynamoDBAdvancedIndexConfigWithDefaults() *DynamoDBAdvancedIndexConfig`

NewDynamoDBAdvancedIndexConfigWithDefaults instantiates a new DynamoDBAdvancedIndexConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) GetLocalSecondaryIndexMode() IndexRestoreMode`

GetLocalSecondaryIndexMode returns the LocalSecondaryIndexMode field if non-nil, zero value otherwise.

### GetLocalSecondaryIndexModeOk

`func (o *DynamoDBAdvancedIndexConfig) GetLocalSecondaryIndexModeOk() (*IndexRestoreMode, bool)`

GetLocalSecondaryIndexModeOk returns a tuple with the LocalSecondaryIndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) SetLocalSecondaryIndexMode(v IndexRestoreMode)`

SetLocalSecondaryIndexMode sets LocalSecondaryIndexMode field to given value.

### HasLocalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) HasLocalSecondaryIndexMode() bool`

HasLocalSecondaryIndexMode returns a boolean if a field has been set.

### GetGlobalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) GetGlobalSecondaryIndexMode() IndexRestoreMode`

GetGlobalSecondaryIndexMode returns the GlobalSecondaryIndexMode field if non-nil, zero value otherwise.

### GetGlobalSecondaryIndexModeOk

`func (o *DynamoDBAdvancedIndexConfig) GetGlobalSecondaryIndexModeOk() (*IndexRestoreMode, bool)`

GetGlobalSecondaryIndexModeOk returns a tuple with the GlobalSecondaryIndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) SetGlobalSecondaryIndexMode(v IndexRestoreMode)`

SetGlobalSecondaryIndexMode sets GlobalSecondaryIndexMode field to given value.

### HasGlobalSecondaryIndexMode

`func (o *DynamoDBAdvancedIndexConfig) HasGlobalSecondaryIndexMode() bool`

HasGlobalSecondaryIndexMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


