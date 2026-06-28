# SameSourceAccountRestoreLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the same-source-account restore restriction is active. | 
**DataAccessRuleId** | Pointer to **string** | Optional. The id of an access condition (an entry in &#x60;accessConditions&#x60;) that scopes which resources the restriction applies to. Omit to apply the restriction to all resources.  | [optional] 

## Methods

### NewSameSourceAccountRestoreLimits

`func NewSameSourceAccountRestoreLimits(enabled bool, ) *SameSourceAccountRestoreLimits`

NewSameSourceAccountRestoreLimits instantiates a new SameSourceAccountRestoreLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSameSourceAccountRestoreLimitsWithDefaults

`func NewSameSourceAccountRestoreLimitsWithDefaults() *SameSourceAccountRestoreLimits`

NewSameSourceAccountRestoreLimitsWithDefaults instantiates a new SameSourceAccountRestoreLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SameSourceAccountRestoreLimits) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SameSourceAccountRestoreLimits) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SameSourceAccountRestoreLimits) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDataAccessRuleId

`func (o *SameSourceAccountRestoreLimits) GetDataAccessRuleId() string`

GetDataAccessRuleId returns the DataAccessRuleId field if non-nil, zero value otherwise.

### GetDataAccessRuleIdOk

`func (o *SameSourceAccountRestoreLimits) GetDataAccessRuleIdOk() (*string, bool)`

GetDataAccessRuleIdOk returns a tuple with the DataAccessRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessRuleId

`func (o *SameSourceAccountRestoreLimits) SetDataAccessRuleId(v string)`

SetDataAccessRuleId sets DataAccessRuleId field to given value.

### HasDataAccessRuleId

`func (o *SameSourceAccountRestoreLimits) HasDataAccessRuleId() bool`

HasDataAccessRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


