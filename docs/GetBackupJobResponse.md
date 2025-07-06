# GetBackupJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**BackupJob**](BackupJob.md) |  | 

## Methods

### NewGetBackupJobResponse

`func NewGetBackupJobResponse(job BackupJob, ) *GetBackupJobResponse`

NewGetBackupJobResponse instantiates a new GetBackupJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackupJobResponseWithDefaults

`func NewGetBackupJobResponseWithDefaults() *GetBackupJobResponse`

NewGetBackupJobResponseWithDefaults instantiates a new GetBackupJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *GetBackupJobResponse) GetJob() BackupJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *GetBackupJobResponse) GetJobOk() (*BackupJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *GetBackupJobResponse) SetJob(v BackupJob)`

SetJob sets Job field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


