# DynamoDbIndexProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectionType** | Pointer to [**DynamoDbProjectionType**](DynamoDbProjectionType.md) |  | [optional] 
**NonKeyAttributes** | Pointer to **[]string** | Non-key attribute names projected into the index. | [optional] 

## Methods

### NewDynamoDbIndexProjection

`func NewDynamoDbIndexProjection() *DynamoDbIndexProjection`

NewDynamoDbIndexProjection instantiates a new DynamoDbIndexProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamoDbIndexProjectionWithDefaults

`func NewDynamoDbIndexProjectionWithDefaults() *DynamoDbIndexProjection`

NewDynamoDbIndexProjectionWithDefaults instantiates a new DynamoDbIndexProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectionType

`func (o *DynamoDbIndexProjection) GetProjectionType() DynamoDbProjectionType`

GetProjectionType returns the ProjectionType field if non-nil, zero value otherwise.

### GetProjectionTypeOk

`func (o *DynamoDbIndexProjection) GetProjectionTypeOk() (*DynamoDbProjectionType, bool)`

GetProjectionTypeOk returns a tuple with the ProjectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionType

`func (o *DynamoDbIndexProjection) SetProjectionType(v DynamoDbProjectionType)`

SetProjectionType sets ProjectionType field to given value.

### HasProjectionType

`func (o *DynamoDbIndexProjection) HasProjectionType() bool`

HasProjectionType returns a boolean if a field has been set.

### GetNonKeyAttributes

`func (o *DynamoDbIndexProjection) GetNonKeyAttributes() []string`

GetNonKeyAttributes returns the NonKeyAttributes field if non-nil, zero value otherwise.

### GetNonKeyAttributesOk

`func (o *DynamoDbIndexProjection) GetNonKeyAttributesOk() (*[]string, bool)`

GetNonKeyAttributesOk returns a tuple with the NonKeyAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonKeyAttributes

`func (o *DynamoDbIndexProjection) SetNonKeyAttributes(v []string)`

SetNonKeyAttributes sets NonKeyAttributes field to given value.

### HasNonKeyAttributes

`func (o *DynamoDbIndexProjection) HasNonKeyAttributes() bool`

HasNonKeyAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


