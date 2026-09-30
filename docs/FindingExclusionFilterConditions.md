# FindingExclusionFilterConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**NullableFindingExclusionScopeFilters**](FindingExclusionScopeFilters.md) |  | [optional] 
**Resource** | Pointer to [**NullableFindingExclusionResourceFilter**](FindingExclusionResourceFilter.md) |  | [optional] 
**Type** | Pointer to [**NullableFindingExclusionTypeFilters**](FindingExclusionTypeFilters.md) |  | [optional] 
**Detector** | Pointer to [**NullableFindingExclusionDetectorFilters**](FindingExclusionDetectorFilters.md) |  | [optional] 
**Value** | Pointer to [**NullableFindingExclusionValueFilters**](FindingExclusionValueFilters.md) |  | [optional] 

## Methods

### NewFindingExclusionFilterConditions

`func NewFindingExclusionFilterConditions() *FindingExclusionFilterConditions`

NewFindingExclusionFilterConditions instantiates a new FindingExclusionFilterConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionFilterConditionsWithDefaults

`func NewFindingExclusionFilterConditionsWithDefaults() *FindingExclusionFilterConditions`

NewFindingExclusionFilterConditionsWithDefaults instantiates a new FindingExclusionFilterConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *FindingExclusionFilterConditions) GetScope() FindingExclusionScopeFilters`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FindingExclusionFilterConditions) GetScopeOk() (*FindingExclusionScopeFilters, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FindingExclusionFilterConditions) SetScope(v FindingExclusionScopeFilters)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FindingExclusionFilterConditions) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *FindingExclusionFilterConditions) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *FindingExclusionFilterConditions) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetResource

`func (o *FindingExclusionFilterConditions) GetResource() FindingExclusionResourceFilter`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *FindingExclusionFilterConditions) GetResourceOk() (*FindingExclusionResourceFilter, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *FindingExclusionFilterConditions) SetResource(v FindingExclusionResourceFilter)`

SetResource sets Resource field to given value.

### HasResource

`func (o *FindingExclusionFilterConditions) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *FindingExclusionFilterConditions) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *FindingExclusionFilterConditions) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetType

`func (o *FindingExclusionFilterConditions) GetType() FindingExclusionTypeFilters`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindingExclusionFilterConditions) GetTypeOk() (*FindingExclusionTypeFilters, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindingExclusionFilterConditions) SetType(v FindingExclusionTypeFilters)`

SetType sets Type field to given value.

### HasType

`func (o *FindingExclusionFilterConditions) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FindingExclusionFilterConditions) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FindingExclusionFilterConditions) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDetector

`func (o *FindingExclusionFilterConditions) GetDetector() FindingExclusionDetectorFilters`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *FindingExclusionFilterConditions) GetDetectorOk() (*FindingExclusionDetectorFilters, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *FindingExclusionFilterConditions) SetDetector(v FindingExclusionDetectorFilters)`

SetDetector sets Detector field to given value.

### HasDetector

`func (o *FindingExclusionFilterConditions) HasDetector() bool`

HasDetector returns a boolean if a field has been set.

### SetDetectorNil

`func (o *FindingExclusionFilterConditions) SetDetectorNil(b bool)`

 SetDetectorNil sets the value for Detector to be an explicit nil

### UnsetDetector
`func (o *FindingExclusionFilterConditions) UnsetDetector()`

UnsetDetector ensures that no value is present for Detector, not even an explicit nil
### GetValue

`func (o *FindingExclusionFilterConditions) GetValue() FindingExclusionValueFilters`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FindingExclusionFilterConditions) GetValueOk() (*FindingExclusionValueFilters, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FindingExclusionFilterConditions) SetValue(v FindingExclusionValueFilters)`

SetValue sets Value field to given value.

### HasValue

`func (o *FindingExclusionFilterConditions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FindingExclusionFilterConditions) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FindingExclusionFilterConditions) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


