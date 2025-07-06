# DailyStorageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryDay** | **time.Time** | Date the storage is calculated for. | 
**DailyStorageBytes** | **int64** | Total snapshot storage for the day, in bytes. | 

## Methods

### NewDailyStorageSummary

`func NewDailyStorageSummary(summaryDay time.Time, dailyStorageBytes int64, ) *DailyStorageSummary`

NewDailyStorageSummary instantiates a new DailyStorageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyStorageSummaryWithDefaults

`func NewDailyStorageSummaryWithDefaults() *DailyStorageSummary`

NewDailyStorageSummaryWithDefaults instantiates a new DailyStorageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryDay

`func (o *DailyStorageSummary) GetSummaryDay() time.Time`

GetSummaryDay returns the SummaryDay field if non-nil, zero value otherwise.

### GetSummaryDayOk

`func (o *DailyStorageSummary) GetSummaryDayOk() (*time.Time, bool)`

GetSummaryDayOk returns a tuple with the SummaryDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryDay

`func (o *DailyStorageSummary) SetSummaryDay(v time.Time)`

SetSummaryDay sets SummaryDay field to given value.


### GetDailyStorageBytes

`func (o *DailyStorageSummary) GetDailyStorageBytes() int64`

GetDailyStorageBytes returns the DailyStorageBytes field if non-nil, zero value otherwise.

### GetDailyStorageBytesOk

`func (o *DailyStorageSummary) GetDailyStorageBytesOk() (*int64, bool)`

GetDailyStorageBytesOk returns a tuple with the DailyStorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyStorageBytes

`func (o *DailyStorageSummary) SetDailyStorageBytes(v int64)`

SetDailyStorageBytes sets DailyStorageBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


