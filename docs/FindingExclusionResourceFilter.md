# FindingExclusionResourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAccountWide** | Pointer to **NullableBool** | When &#x60;true&#x60;, matches only exclusions that apply to every resource in the account. When &#x60;false&#x60;, matches only exclusions for a single resource.  | [optional] 
**In** | Pointer to **[]string** | Matches exclusions for any of these resources, by Eon-assigned resource ID. An empty list applies no filter. | [optional] 

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

### GetIsAccountWide

`func (o *FindingExclusionResourceFilter) GetIsAccountWide() bool`

GetIsAccountWide returns the IsAccountWide field if non-nil, zero value otherwise.

### GetIsAccountWideOk

`func (o *FindingExclusionResourceFilter) GetIsAccountWideOk() (*bool, bool)`

GetIsAccountWideOk returns a tuple with the IsAccountWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountWide

`func (o *FindingExclusionResourceFilter) SetIsAccountWide(v bool)`

SetIsAccountWide sets IsAccountWide field to given value.

### HasIsAccountWide

`func (o *FindingExclusionResourceFilter) HasIsAccountWide() bool`

HasIsAccountWide returns a boolean if a field has been set.

### SetIsAccountWideNil

`func (o *FindingExclusionResourceFilter) SetIsAccountWideNil(b bool)`

 SetIsAccountWideNil sets the value for IsAccountWide to be an explicit nil

### UnsetIsAccountWide
`func (o *FindingExclusionResourceFilter) UnsetIsAccountWide()`

UnsetIsAccountWide ensures that no value is present for IsAccountWide, not even an explicit nil
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


