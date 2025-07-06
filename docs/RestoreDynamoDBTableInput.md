# RestoreDynamoDBTableInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**DynamodbTableRestoreDestination**](DynamodbTableRestoreDestination.md) |  | 

## Methods

### NewRestoreDynamoDBTableInput

`func NewRestoreDynamoDBTableInput(restoreAccountId string, destination DynamodbTableRestoreDestination, ) *RestoreDynamoDBTableInput`

NewRestoreDynamoDBTableInput instantiates a new RestoreDynamoDBTableInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreDynamoDBTableInputWithDefaults

`func NewRestoreDynamoDBTableInputWithDefaults() *RestoreDynamoDBTableInput`

NewRestoreDynamoDBTableInputWithDefaults instantiates a new RestoreDynamoDBTableInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *RestoreDynamoDBTableInput) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreDynamoDBTableInput) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreDynamoDBTableInput) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreDynamoDBTableInput) GetDestination() DynamodbTableRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreDynamoDBTableInput) GetDestinationOk() (*DynamodbTableRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreDynamoDBTableInput) SetDestination(v DynamodbTableRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


