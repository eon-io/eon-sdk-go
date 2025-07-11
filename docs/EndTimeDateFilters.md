# EndTimeDateFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Matches if &#x60;endTime&#x60; is on or later than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 
**EndDate** | Pointer to **string** | Matches if &#x60;endTime&#x60; is on or earlier than this date. Must be in ISO 8601 YYYY-MM-DD format.  | [optional] 

## Methods

### NewEndTimeDateFilters

`func NewEndTimeDateFilters() *EndTimeDateFilters`

NewEndTimeDateFilters instantiates a new EndTimeDateFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndTimeDateFiltersWithDefaults

`func NewEndTimeDateFiltersWithDefaults() *EndTimeDateFilters`

NewEndTimeDateFiltersWithDefaults instantiates a new EndTimeDateFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *EndTimeDateFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EndTimeDateFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EndTimeDateFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EndTimeDateFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EndTimeDateFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EndTimeDateFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EndTimeDateFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EndTimeDateFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


