# EbsSnapshotRestoreDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsEbs** | Pointer to [**EbsSnapshotTarget**](EbsSnapshotTarget.md) |  | [optional] 

## Methods

### NewEbsSnapshotRestoreDestination

`func NewEbsSnapshotRestoreDestination() *EbsSnapshotRestoreDestination`

NewEbsSnapshotRestoreDestination instantiates a new EbsSnapshotRestoreDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbsSnapshotRestoreDestinationWithDefaults

`func NewEbsSnapshotRestoreDestinationWithDefaults() *EbsSnapshotRestoreDestination`

NewEbsSnapshotRestoreDestinationWithDefaults instantiates a new EbsSnapshotRestoreDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsEbs

`func (o *EbsSnapshotRestoreDestination) GetAwsEbs() EbsSnapshotTarget`

GetAwsEbs returns the AwsEbs field if non-nil, zero value otherwise.

### GetAwsEbsOk

`func (o *EbsSnapshotRestoreDestination) GetAwsEbsOk() (*EbsSnapshotTarget, bool)`

GetAwsEbsOk returns a tuple with the AwsEbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEbs

`func (o *EbsSnapshotRestoreDestination) SetAwsEbs(v EbsSnapshotTarget)`

SetAwsEbs sets AwsEbs field to given value.

### HasAwsEbs

`func (o *EbsSnapshotRestoreDestination) HasAwsEbs() bool`

HasAwsEbs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


