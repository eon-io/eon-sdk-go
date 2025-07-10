# StartTimeDateFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Matches if &#x60;startTime&#x60; is on or later than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 
**EndDate** | Pointer to **string** | Matches if &#x60;startTime&#x60; is on or earlier than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 

## Methods

### NewStartTimeDateFilters

`func NewStartTimeDateFilters() *StartTimeDateFilters`

NewStartTimeDateFilters instantiates a new StartTimeDateFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartTimeDateFiltersWithDefaults

`func NewStartTimeDateFiltersWithDefaults() *StartTimeDateFilters`

NewStartTimeDateFiltersWithDefaults instantiates a new StartTimeDateFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *StartTimeDateFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StartTimeDateFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StartTimeDateFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StartTimeDateFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *StartTimeDateFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StartTimeDateFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StartTimeDateFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StartTimeDateFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


