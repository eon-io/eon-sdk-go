# JobSnapshotDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Snapshot ID. | [optional] 
**PointInTime** | **time.Time** | Date and time of the resource that&#39;s preserved by the snapshot. | 
**ExpirationTime** | Pointer to **time.Time** | Date and time the snapshot&#39;s retention is expected to expire, after which it&#39;s marked for deletion. | [optional] 

## Methods

### NewJobSnapshotDetails

`func NewJobSnapshotDetails(pointInTime time.Time, ) *JobSnapshotDetails`

NewJobSnapshotDetails instantiates a new JobSnapshotDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSnapshotDetailsWithDefaults

`func NewJobSnapshotDetailsWithDefaults() *JobSnapshotDetails`

NewJobSnapshotDetailsWithDefaults instantiates a new JobSnapshotDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobSnapshotDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobSnapshotDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobSnapshotDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobSnapshotDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPointInTime

`func (o *JobSnapshotDetails) GetPointInTime() time.Time`

GetPointInTime returns the PointInTime field if non-nil, zero value otherwise.

### GetPointInTimeOk

`func (o *JobSnapshotDetails) GetPointInTimeOk() (*time.Time, bool)`

GetPointInTimeOk returns a tuple with the PointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTime

`func (o *JobSnapshotDetails) SetPointInTime(v time.Time)`

SetPointInTime sets PointInTime field to given value.


### GetExpirationTime

`func (o *JobSnapshotDetails) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *JobSnapshotDetails) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *JobSnapshotDetails) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *JobSnapshotDetails) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


