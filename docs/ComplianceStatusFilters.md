# ComplianceStatusFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]ComplianceStatus**](ComplianceStatus.md) | Matches if any value in this list equals &#x60;complianceStatus&#x60;. | [optional] 
**NotIn** | Pointer to [**[]ComplianceStatus**](ComplianceStatus.md) | Matches if no value in this list equals &#x60;complianceStatus&#x60;. | [optional] 

## Methods

### NewComplianceStatusFilters

`func NewComplianceStatusFilters() *ComplianceStatusFilters`

NewComplianceStatusFilters instantiates a new ComplianceStatusFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceStatusFiltersWithDefaults

`func NewComplianceStatusFiltersWithDefaults() *ComplianceStatusFilters`

NewComplianceStatusFiltersWithDefaults instantiates a new ComplianceStatusFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *ComplianceStatusFilters) GetIn() []ComplianceStatus`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ComplianceStatusFilters) GetInOk() (*[]ComplianceStatus, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ComplianceStatusFilters) SetIn(v []ComplianceStatus)`

SetIn sets In field to given value.

### HasIn

`func (o *ComplianceStatusFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *ComplianceStatusFilters) GetNotIn() []ComplianceStatus`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *ComplianceStatusFilters) GetNotInOk() (*[]ComplianceStatus, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *ComplianceStatusFilters) SetNotIn(v []ComplianceStatus)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *ComplianceStatusFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


