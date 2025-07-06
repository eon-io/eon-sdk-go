# QueryDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** | Query ID. Can be used to call [Get Query Status](./get-query-status) and [Get Query Result](./get-query-result).  | 

## Methods

### NewQueryDBResponse

`func NewQueryDBResponse(queryId string, ) *QueryDBResponse`

NewQueryDBResponse instantiates a new QueryDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDBResponseWithDefaults

`func NewQueryDBResponseWithDefaults() *QueryDBResponse`

NewQueryDBResponseWithDefaults instantiates a new QueryDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *QueryDBResponse) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *QueryDBResponse) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *QueryDBResponse) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


