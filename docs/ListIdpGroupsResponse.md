# ListIdpGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]IdpGroup**](IdpGroup.md) | List of IDP groups. | 
**TotalCount** | Pointer to **int32** | Total number of IDP groups. | [optional] 
**NextToken** | Pointer to **string** | Token to retrieve the next page of results. | [optional] 

## Methods

### NewListIdpGroupsResponse

`func NewListIdpGroupsResponse(groups []IdpGroup, ) *ListIdpGroupsResponse`

NewListIdpGroupsResponse instantiates a new ListIdpGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdpGroupsResponseWithDefaults

`func NewListIdpGroupsResponseWithDefaults() *ListIdpGroupsResponse`

NewListIdpGroupsResponseWithDefaults instantiates a new ListIdpGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListIdpGroupsResponse) GetGroups() []IdpGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListIdpGroupsResponse) GetGroupsOk() (*[]IdpGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListIdpGroupsResponse) SetGroups(v []IdpGroup)`

SetGroups sets Groups field to given value.


### GetTotalCount

`func (o *ListIdpGroupsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListIdpGroupsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListIdpGroupsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListIdpGroupsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextToken

`func (o *ListIdpGroupsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListIdpGroupsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListIdpGroupsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListIdpGroupsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


