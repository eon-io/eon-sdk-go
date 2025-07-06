# AwsDynamoDBDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreRegion** | **string** | Region to restore to. | 
**EncryptionKeyId** | **string** | ID of the key you want Eon to use for encrypting the restored resource. | 
**RestoredName** | **string** | Name to assign to the restored resource. | 
**WriteCapacityUnits** | Pointer to **int32** | The number of write capacity units for the restored table. If not specified, the default is 5. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 

## Methods

### NewAwsDynamoDBDestination

`func NewAwsDynamoDBDestination(restoreRegion string, encryptionKeyId string, restoredName string, ) *AwsDynamoDBDestination`

NewAwsDynamoDBDestination instantiates a new AwsDynamoDBDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsDynamoDBDestinationWithDefaults

`func NewAwsDynamoDBDestinationWithDefaults() *AwsDynamoDBDestination`

NewAwsDynamoDBDestinationWithDefaults instantiates a new AwsDynamoDBDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreRegion

`func (o *AwsDynamoDBDestination) GetRestoreRegion() string`

GetRestoreRegion returns the RestoreRegion field if non-nil, zero value otherwise.

### GetRestoreRegionOk

`func (o *AwsDynamoDBDestination) GetRestoreRegionOk() (*string, bool)`

GetRestoreRegionOk returns a tuple with the RestoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreRegion

`func (o *AwsDynamoDBDestination) SetRestoreRegion(v string)`

SetRestoreRegion sets RestoreRegion field to given value.


### GetEncryptionKeyId

`func (o *AwsDynamoDBDestination) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *AwsDynamoDBDestination) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *AwsDynamoDBDestination) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.


### GetRestoredName

`func (o *AwsDynamoDBDestination) GetRestoredName() string`

GetRestoredName returns the RestoredName field if non-nil, zero value otherwise.

### GetRestoredNameOk

`func (o *AwsDynamoDBDestination) GetRestoredNameOk() (*string, bool)`

GetRestoredNameOk returns a tuple with the RestoredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredName

`func (o *AwsDynamoDBDestination) SetRestoredName(v string)`

SetRestoredName sets RestoredName field to given value.


### GetWriteCapacityUnits

`func (o *AwsDynamoDBDestination) GetWriteCapacityUnits() int32`

GetWriteCapacityUnits returns the WriteCapacityUnits field if non-nil, zero value otherwise.

### GetWriteCapacityUnitsOk

`func (o *AwsDynamoDBDestination) GetWriteCapacityUnitsOk() (*int32, bool)`

GetWriteCapacityUnitsOk returns a tuple with the WriteCapacityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCapacityUnits

`func (o *AwsDynamoDBDestination) SetWriteCapacityUnits(v int32)`

SetWriteCapacityUnits sets WriteCapacityUnits field to given value.

### HasWriteCapacityUnits

`func (o *AwsDynamoDBDestination) HasWriteCapacityUnits() bool`

HasWriteCapacityUnits returns a boolean if a field has been set.

### GetTags

`func (o *AwsDynamoDBDestination) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsDynamoDBDestination) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsDynamoDBDestination) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsDynamoDBDestination) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


