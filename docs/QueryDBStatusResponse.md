# QueryDBStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningTimeMs** | Pointer to **int32** | Time the query took to run, in milliseconds. | [optional] 
**Status** | [**QueryDBStatus**](QueryDBStatus.md) |  | 
**ErrorMessage** | Pointer to **NullableString** | Error message. Present only if the status is &#x60;FAILED&#x60;.  | [optional] 
**OutputLocations** | Pointer to **[]string** | Locations where the query results are saved to. This is set by &#x60;destination&#x60; in the [Run Query](./run-query) request body.  | [optional] 

## Methods

### NewQueryDBStatusResponse

`func NewQueryDBStatusResponse(status QueryDBStatus, ) *QueryDBStatusResponse`

NewQueryDBStatusResponse instantiates a new QueryDBStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDBStatusResponseWithDefaults

`func NewQueryDBStatusResponseWithDefaults() *QueryDBStatusResponse`

NewQueryDBStatusResponseWithDefaults instantiates a new QueryDBStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunningTimeMs

`func (o *QueryDBStatusResponse) GetRunningTimeMs() int32`

GetRunningTimeMs returns the RunningTimeMs field if non-nil, zero value otherwise.

### GetRunningTimeMsOk

`func (o *QueryDBStatusResponse) GetRunningTimeMsOk() (*int32, bool)`

GetRunningTimeMsOk returns a tuple with the RunningTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTimeMs

`func (o *QueryDBStatusResponse) SetRunningTimeMs(v int32)`

SetRunningTimeMs sets RunningTimeMs field to given value.

### HasRunningTimeMs

`func (o *QueryDBStatusResponse) HasRunningTimeMs() bool`

HasRunningTimeMs returns a boolean if a field has been set.

### GetStatus

`func (o *QueryDBStatusResponse) GetStatus() QueryDBStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryDBStatusResponse) GetStatusOk() (*QueryDBStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryDBStatusResponse) SetStatus(v QueryDBStatus)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *QueryDBStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QueryDBStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QueryDBStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QueryDBStatusResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *QueryDBStatusResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *QueryDBStatusResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetOutputLocations

`func (o *QueryDBStatusResponse) GetOutputLocations() []string`

GetOutputLocations returns the OutputLocations field if non-nil, zero value otherwise.

### GetOutputLocationsOk

`func (o *QueryDBStatusResponse) GetOutputLocationsOk() (*[]string, bool)`

GetOutputLocationsOk returns a tuple with the OutputLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocations

`func (o *QueryDBStatusResponse) SetOutputLocations(v []string)`

SetOutputLocations sets OutputLocations field to given value.

### HasOutputLocations

`func (o *QueryDBStatusResponse) HasOutputLocations() bool`

HasOutputLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


