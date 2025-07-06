# DatabaseDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRds** | Pointer to [**AwsDatabaseDestination**](AwsDatabaseDestination.md) |  | [optional] 

## Methods

### NewDatabaseDestination

`func NewDatabaseDestination() *DatabaseDestination`

NewDatabaseDestination instantiates a new DatabaseDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseDestinationWithDefaults

`func NewDatabaseDestinationWithDefaults() *DatabaseDestination`

NewDatabaseDestinationWithDefaults instantiates a new DatabaseDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRds

`func (o *DatabaseDestination) GetAwsRds() AwsDatabaseDestination`

GetAwsRds returns the AwsRds field if non-nil, zero value otherwise.

### GetAwsRdsOk

`func (o *DatabaseDestination) GetAwsRdsOk() (*AwsDatabaseDestination, bool)`

GetAwsRdsOk returns a tuple with the AwsRds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRds

`func (o *DatabaseDestination) SetAwsRds(v AwsDatabaseDestination)`

SetAwsRds sets AwsRds field to given value.

### HasAwsRds

`func (o *DatabaseDestination) HasAwsRds() bool`

HasAwsRds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


