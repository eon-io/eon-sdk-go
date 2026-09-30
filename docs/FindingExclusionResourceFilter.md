# FindingExclusionResourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches exclusions for any of these resources, by cloud provider resource ID. An empty list applies no filter; IDs that match no resource in the account match no exclusion.  | [optional] 

## Methods

### NewFindingExclusionResourceFilter

`func NewFindingExclusionResourceFilter() *FindingExclusionResourceFilter`

NewFindingExclusionResourceFilter instantiates a new FindingExclusionResourceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionResourceFilterWithDefaults

`func NewFindingExclusionResourceFilterWithDefaults() *FindingExclusionResourceFilter`

NewFindingExclusionResourceFilterWithDefaults instantiates a new FindingExclusionResourceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *FindingExclusionResourceFilter) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *FindingExclusionResourceFilter) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *FindingExclusionResourceFilter) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *FindingExclusionResourceFilter) HasIn() bool`

HasIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


