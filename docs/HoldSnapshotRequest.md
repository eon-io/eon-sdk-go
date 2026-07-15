# HoldSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional free-text note explaining why the snapshot is being held. Shown on the snapshot&#39;s hold badge and cleared when the hold is removed.  | [optional] 

## Methods

### NewHoldSnapshotRequest

`func NewHoldSnapshotRequest() *HoldSnapshotRequest`

NewHoldSnapshotRequest instantiates a new HoldSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldSnapshotRequestWithDefaults

`func NewHoldSnapshotRequestWithDefaults() *HoldSnapshotRequest`

NewHoldSnapshotRequestWithDefaults instantiates a new HoldSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HoldSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldSnapshotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldSnapshotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


