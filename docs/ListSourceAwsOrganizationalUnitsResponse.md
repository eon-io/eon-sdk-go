# ListSourceAwsOrganizationalUnitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationalUnits** | [**[]SourceAwsOrganizationalUnit**](SourceAwsOrganizationalUnit.md) | List of AWS organizational units. | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | **int32** | Total number of AWS organizational units that matched the filter options. | 

## Methods

### NewListSourceAwsOrganizationalUnitsResponse

`func NewListSourceAwsOrganizationalUnitsResponse(organizationalUnits []SourceAwsOrganizationalUnit, totalCount int32, ) *ListSourceAwsOrganizationalUnitsResponse`

NewListSourceAwsOrganizationalUnitsResponse instantiates a new ListSourceAwsOrganizationalUnitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceAwsOrganizationalUnitsResponseWithDefaults

`func NewListSourceAwsOrganizationalUnitsResponseWithDefaults() *ListSourceAwsOrganizationalUnitsResponse`

NewListSourceAwsOrganizationalUnitsResponseWithDefaults instantiates a new ListSourceAwsOrganizationalUnitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationalUnits

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetOrganizationalUnits() []SourceAwsOrganizationalUnit`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetOrganizationalUnitsOk() (*[]SourceAwsOrganizationalUnit, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *ListSourceAwsOrganizationalUnitsResponse) SetOrganizationalUnits(v []SourceAwsOrganizationalUnit)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.


### GetNextToken

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListSourceAwsOrganizationalUnitsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListSourceAwsOrganizationalUnitsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListSourceAwsOrganizationalUnitsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListSourceAwsOrganizationalUnitsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


