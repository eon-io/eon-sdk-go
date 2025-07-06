# ListRestoreJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]RestoreJob**](RestoreJob.md) | List of retrieved restore jobs. | [optional] 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of restore jobs that matched the filter options. | [optional] 

## Methods

### NewListRestoreJobsResponse

`func NewListRestoreJobsResponse() *ListRestoreJobsResponse`

NewListRestoreJobsResponse instantiates a new ListRestoreJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestoreJobsResponseWithDefaults

`func NewListRestoreJobsResponseWithDefaults() *ListRestoreJobsResponse`

NewListRestoreJobsResponseWithDefaults instantiates a new ListRestoreJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListRestoreJobsResponse) GetJobs() []RestoreJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListRestoreJobsResponse) GetJobsOk() (*[]RestoreJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListRestoreJobsResponse) SetJobs(v []RestoreJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListRestoreJobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetNextToken

`func (o *ListRestoreJobsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListRestoreJobsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListRestoreJobsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListRestoreJobsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListRestoreJobsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListRestoreJobsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListRestoreJobsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListRestoreJobsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


