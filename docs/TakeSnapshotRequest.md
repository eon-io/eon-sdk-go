# TakeSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaultId** | **string** | Eon-assigned ID of the vault to store the snapshot in.  | 
**RetentionDays** | **int32** | Number of days to retain the snapshot. | 

## Methods

### NewTakeSnapshotRequest

`func NewTakeSnapshotRequest(vaultId string, retentionDays int32, ) *TakeSnapshotRequest`

NewTakeSnapshotRequest instantiates a new TakeSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTakeSnapshotRequestWithDefaults

`func NewTakeSnapshotRequestWithDefaults() *TakeSnapshotRequest`

NewTakeSnapshotRequestWithDefaults instantiates a new TakeSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaultId

`func (o *TakeSnapshotRequest) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *TakeSnapshotRequest) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *TakeSnapshotRequest) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetRetentionDays

`func (o *TakeSnapshotRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *TakeSnapshotRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *TakeSnapshotRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


