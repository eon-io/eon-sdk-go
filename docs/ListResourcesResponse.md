# ListResourcesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]InventoryResource**](InventoryResource.md) | List of retrieved resources. | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of resources that matched the filter options. | [optional] 

## Methods

### NewListResourcesResponse

`func NewListResourcesResponse(resources []InventoryResource, ) *ListResourcesResponse`

NewListResourcesResponse instantiates a new ListResourcesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourcesResponseWithDefaults

`func NewListResourcesResponseWithDefaults() *ListResourcesResponse`

NewListResourcesResponseWithDefaults instantiates a new ListResourcesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ListResourcesResponse) GetResources() []InventoryResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ListResourcesResponse) GetResourcesOk() (*[]InventoryResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ListResourcesResponse) SetResources(v []InventoryResource)`

SetResources sets Resources field to given value.


### GetNextToken

`func (o *ListResourcesResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListResourcesResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListResourcesResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListResourcesResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListResourcesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListResourcesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListResourcesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListResourcesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


