# ListFindingExclusionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingExclusions** | [**[]FindingExclusion**](FindingExclusion.md) | Finding exclusions on the current page. | 
**TotalCount** | **int32** | Total number of finding exclusions that matched the filter options. | 
**NextPageToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 

## Methods

### NewListFindingExclusionsResponse

`func NewListFindingExclusionsResponse(findingExclusions []FindingExclusion, totalCount int32, ) *ListFindingExclusionsResponse`

NewListFindingExclusionsResponse instantiates a new ListFindingExclusionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFindingExclusionsResponseWithDefaults

`func NewListFindingExclusionsResponseWithDefaults() *ListFindingExclusionsResponse`

NewListFindingExclusionsResponseWithDefaults instantiates a new ListFindingExclusionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingExclusions

`func (o *ListFindingExclusionsResponse) GetFindingExclusions() []FindingExclusion`

GetFindingExclusions returns the FindingExclusions field if non-nil, zero value otherwise.

### GetFindingExclusionsOk

`func (o *ListFindingExclusionsResponse) GetFindingExclusionsOk() (*[]FindingExclusion, bool)`

GetFindingExclusionsOk returns a tuple with the FindingExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingExclusions

`func (o *ListFindingExclusionsResponse) SetFindingExclusions(v []FindingExclusion)`

SetFindingExclusions sets FindingExclusions field to given value.


### GetTotalCount

`func (o *ListFindingExclusionsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListFindingExclusionsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListFindingExclusionsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPageToken

`func (o *ListFindingExclusionsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListFindingExclusionsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListFindingExclusionsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListFindingExclusionsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


