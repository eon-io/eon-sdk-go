# ListRestoreAwsOrganizationalUnitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationalUnits** | [**[]RestoreAwsOrganizationalUnit**](RestoreAwsOrganizationalUnit.md) | List of AWS organizational units. | 
**NextToken** | Pointer to **string** | Cursor that points to the first record of the next page of results. Pass this value in the next request.  | [optional] 
**TotalCount** | **int32** | Total number of AWS organizational units that matched the filter options. | 

## Methods

### NewListRestoreAwsOrganizationalUnitsResponse

`func NewListRestoreAwsOrganizationalUnitsResponse(organizationalUnits []RestoreAwsOrganizationalUnit, totalCount int32, ) *ListRestoreAwsOrganizationalUnitsResponse`

NewListRestoreAwsOrganizationalUnitsResponse instantiates a new ListRestoreAwsOrganizationalUnitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestoreAwsOrganizationalUnitsResponseWithDefaults

`func NewListRestoreAwsOrganizationalUnitsResponseWithDefaults() *ListRestoreAwsOrganizationalUnitsResponse`

NewListRestoreAwsOrganizationalUnitsResponseWithDefaults instantiates a new ListRestoreAwsOrganizationalUnitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationalUnits

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetOrganizationalUnits() []RestoreAwsOrganizationalUnit`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetOrganizationalUnitsOk() (*[]RestoreAwsOrganizationalUnit, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *ListRestoreAwsOrganizationalUnitsResponse) SetOrganizationalUnits(v []RestoreAwsOrganizationalUnit)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.


### GetNextToken

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListRestoreAwsOrganizationalUnitsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListRestoreAwsOrganizationalUnitsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListRestoreAwsOrganizationalUnitsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListRestoreAwsOrganizationalUnitsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


