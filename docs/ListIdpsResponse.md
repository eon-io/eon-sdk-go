# ListIdpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idps** | [**[]Idp**](Idp.md) | List of Identity Providers. | 
**TotalCount** | **int32** | Total number of IDPs. | 
**NextToken** | Pointer to **string** | Token to retrieve the next page of results. | [optional] 

## Methods

### NewListIdpsResponse

`func NewListIdpsResponse(idps []Idp, totalCount int32, ) *ListIdpsResponse`

NewListIdpsResponse instantiates a new ListIdpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdpsResponseWithDefaults

`func NewListIdpsResponseWithDefaults() *ListIdpsResponse`

NewListIdpsResponseWithDefaults instantiates a new ListIdpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdps

`func (o *ListIdpsResponse) GetIdps() []Idp`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *ListIdpsResponse) GetIdpsOk() (*[]Idp, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *ListIdpsResponse) SetIdps(v []Idp)`

SetIdps sets Idps field to given value.


### GetTotalCount

`func (o *ListIdpsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListIdpsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListIdpsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *ListIdpsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListIdpsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListIdpsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListIdpsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


