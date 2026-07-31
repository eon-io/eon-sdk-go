# ListBackupPostureControlsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sorts** | Pointer to [**[]BackupPostureControlSort**](BackupPostureControlSort.md) | Sort instructions applied in order. When omitted, controls are returned sorted by severity, most severe first.  | [optional] 

## Methods

### NewListBackupPostureControlsRequest

`func NewListBackupPostureControlsRequest() *ListBackupPostureControlsRequest`

NewListBackupPostureControlsRequest instantiates a new ListBackupPostureControlsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupPostureControlsRequestWithDefaults

`func NewListBackupPostureControlsRequestWithDefaults() *ListBackupPostureControlsRequest`

NewListBackupPostureControlsRequestWithDefaults instantiates a new ListBackupPostureControlsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSorts

`func (o *ListBackupPostureControlsRequest) GetSorts() []BackupPostureControlSort`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *ListBackupPostureControlsRequest) GetSortsOk() (*[]BackupPostureControlSort, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *ListBackupPostureControlsRequest) SetSorts(v []BackupPostureControlSort)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *ListBackupPostureControlsRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


