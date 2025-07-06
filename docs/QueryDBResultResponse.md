# QueryDBResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | **[]string** | List of the column names in the returned records. | 
**Records** | **[][]string** | List of records, returned as an array of arrays. The inner array contains the values of the columns in the same order as returned in &#x60;columns&#x60;.  | 
**NextToken** | **NullableString** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | 

## Methods

### NewQueryDBResultResponse

`func NewQueryDBResultResponse(columns []string, records [][]string, nextToken NullableString, ) *QueryDBResultResponse`

NewQueryDBResultResponse instantiates a new QueryDBResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDBResultResponseWithDefaults

`func NewQueryDBResultResponseWithDefaults() *QueryDBResultResponse`

NewQueryDBResultResponseWithDefaults instantiates a new QueryDBResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *QueryDBResultResponse) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryDBResultResponse) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryDBResultResponse) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetRecords

`func (o *QueryDBResultResponse) GetRecords() [][]string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *QueryDBResultResponse) GetRecordsOk() (*[][]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *QueryDBResultResponse) SetRecords(v [][]string)`

SetRecords sets Records field to given value.


### GetNextToken

`func (o *QueryDBResultResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *QueryDBResultResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *QueryDBResultResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *QueryDBResultResponse) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *QueryDBResultResponse) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


