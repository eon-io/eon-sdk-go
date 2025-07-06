# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Eon snapshot ID. | 
**ProjectId** | Pointer to **string** | ID of the snapshot&#39;s parent project. | [optional] 
**CreatedTime** | **time.Time** | Date and time the snapshot creation was started. This doesn&#39;t represent the point in time the resource is backed up from, which is instead represented by the &#x60;pointInTime&#x60; property.  | 
**PointInTime** | Pointer to **time.Time** | Date and time of the resource that&#39;s preserved by the snapshot. | [optional] 
**VaultId** | **string** | ID of the vault the snapshot is stored in. | 
**ResourceId** | **string** | Eon-assigned ID of the resource the snapshot is backing up. | 
**ExpirationTime** | Pointer to **time.Time** | Date and time the snapshot&#39;s retention is expected to expire, after which it&#39;s marked for deletion. | [optional] 
**Resource** | Pointer to [**ResourceSnapshot**](ResourceSnapshot.md) |  | [optional] 

## Methods

### NewSnapshot

`func NewSnapshot(id string, createdTime time.Time, vaultId string, resourceId string, ) *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Snapshot) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Snapshot) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Snapshot) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Snapshot) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Snapshot) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Snapshot) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Snapshot) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetPointInTime

`func (o *Snapshot) GetPointInTime() time.Time`

GetPointInTime returns the PointInTime field if non-nil, zero value otherwise.

### GetPointInTimeOk

`func (o *Snapshot) GetPointInTimeOk() (*time.Time, bool)`

GetPointInTimeOk returns a tuple with the PointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTime

`func (o *Snapshot) SetPointInTime(v time.Time)`

SetPointInTime sets PointInTime field to given value.

### HasPointInTime

`func (o *Snapshot) HasPointInTime() bool`

HasPointInTime returns a boolean if a field has been set.

### GetVaultId

`func (o *Snapshot) GetVaultId() string`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *Snapshot) GetVaultIdOk() (*string, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *Snapshot) SetVaultId(v string)`

SetVaultId sets VaultId field to given value.


### GetResourceId

`func (o *Snapshot) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Snapshot) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Snapshot) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetExpirationTime

`func (o *Snapshot) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *Snapshot) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *Snapshot) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *Snapshot) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetResource

`func (o *Snapshot) GetResource() ResourceSnapshot`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Snapshot) GetResourceOk() (*ResourceSnapshot, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Snapshot) SetResource(v ResourceSnapshot)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Snapshot) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


