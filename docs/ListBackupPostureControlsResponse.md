# ListBackupPostureControlsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPostureControls** | [**[]BackupPostureControl**](BackupPostureControl.md) | List of retrieved backup posture controls. | 
**TotalCount** | Pointer to **int32** | Total number of backup posture controls that matched the filter options. | [optional] 
**NextPageToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 

## Methods

### NewListBackupPostureControlsResponse

`func NewListBackupPostureControlsResponse(backupPostureControls []BackupPostureControl, ) *ListBackupPostureControlsResponse`

NewListBackupPostureControlsResponse instantiates a new ListBackupPostureControlsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupPostureControlsResponseWithDefaults

`func NewListBackupPostureControlsResponseWithDefaults() *ListBackupPostureControlsResponse`

NewListBackupPostureControlsResponseWithDefaults instantiates a new ListBackupPostureControlsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPostureControls

`func (o *ListBackupPostureControlsResponse) GetBackupPostureControls() []BackupPostureControl`

GetBackupPostureControls returns the BackupPostureControls field if non-nil, zero value otherwise.

### GetBackupPostureControlsOk

`func (o *ListBackupPostureControlsResponse) GetBackupPostureControlsOk() (*[]BackupPostureControl, bool)`

GetBackupPostureControlsOk returns a tuple with the BackupPostureControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPostureControls

`func (o *ListBackupPostureControlsResponse) SetBackupPostureControls(v []BackupPostureControl)`

SetBackupPostureControls sets BackupPostureControls field to given value.


### GetTotalCount

`func (o *ListBackupPostureControlsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListBackupPostureControlsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListBackupPostureControlsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListBackupPostureControlsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListBackupPostureControlsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListBackupPostureControlsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListBackupPostureControlsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListBackupPostureControlsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


