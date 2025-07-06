# ListBackupJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]BackupJob**](BackupJob.md) | List of retrieved backup jobs. | [optional] 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of backup jobs that matched the filter options. | [optional] 

## Methods

### NewListBackupJobsResponse

`func NewListBackupJobsResponse() *ListBackupJobsResponse`

NewListBackupJobsResponse instantiates a new ListBackupJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupJobsResponseWithDefaults

`func NewListBackupJobsResponseWithDefaults() *ListBackupJobsResponse`

NewListBackupJobsResponseWithDefaults instantiates a new ListBackupJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListBackupJobsResponse) GetJobs() []BackupJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListBackupJobsResponse) GetJobsOk() (*[]BackupJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListBackupJobsResponse) SetJobs(v []BackupJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListBackupJobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetNextToken

`func (o *ListBackupJobsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListBackupJobsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListBackupJobsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListBackupJobsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListBackupJobsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListBackupJobsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListBackupJobsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListBackupJobsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


