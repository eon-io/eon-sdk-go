# ListActionApprovalRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionApprovalRules** | [**[]ActionApprovalRule**](ActionApprovalRule.md) | Action approval rules in the current page. | 
**TotalCount** | **int32** | Total number of matching action approval rules. | 
**NextPageToken** | Pointer to **string** | Cursor that points to the first rule on the next page. | [optional] 

## Methods

### NewListActionApprovalRulesResponse

`func NewListActionApprovalRulesResponse(actionApprovalRules []ActionApprovalRule, totalCount int32, ) *ListActionApprovalRulesResponse`

NewListActionApprovalRulesResponse instantiates a new ListActionApprovalRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActionApprovalRulesResponseWithDefaults

`func NewListActionApprovalRulesResponseWithDefaults() *ListActionApprovalRulesResponse`

NewListActionApprovalRulesResponseWithDefaults instantiates a new ListActionApprovalRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionApprovalRules

`func (o *ListActionApprovalRulesResponse) GetActionApprovalRules() []ActionApprovalRule`

GetActionApprovalRules returns the ActionApprovalRules field if non-nil, zero value otherwise.

### GetActionApprovalRulesOk

`func (o *ListActionApprovalRulesResponse) GetActionApprovalRulesOk() (*[]ActionApprovalRule, bool)`

GetActionApprovalRulesOk returns a tuple with the ActionApprovalRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApprovalRules

`func (o *ListActionApprovalRulesResponse) SetActionApprovalRules(v []ActionApprovalRule)`

SetActionApprovalRules sets ActionApprovalRules field to given value.


### GetTotalCount

`func (o *ListActionApprovalRulesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListActionApprovalRulesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListActionApprovalRulesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextPageToken

`func (o *ListActionApprovalRulesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListActionApprovalRulesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListActionApprovalRulesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListActionApprovalRulesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


