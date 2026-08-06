# MinRetentionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumRetention** | **int32** | Minimum backup retention period, in days. | 
**Frequency** | **string** | Backup cadence the minimum retention applies to. | 

## Methods

### NewMinRetentionRule

`func NewMinRetentionRule(minimumRetention int32, frequency string, ) *MinRetentionRule`

NewMinRetentionRule instantiates a new MinRetentionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinRetentionRuleWithDefaults

`func NewMinRetentionRuleWithDefaults() *MinRetentionRule`

NewMinRetentionRuleWithDefaults instantiates a new MinRetentionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumRetention

`func (o *MinRetentionRule) GetMinimumRetention() int32`

GetMinimumRetention returns the MinimumRetention field if non-nil, zero value otherwise.

### GetMinimumRetentionOk

`func (o *MinRetentionRule) GetMinimumRetentionOk() (*int32, bool)`

GetMinimumRetentionOk returns a tuple with the MinimumRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRetention

`func (o *MinRetentionRule) SetMinimumRetention(v int32)`

SetMinimumRetention sets MinimumRetention field to given value.


### GetFrequency

`func (o *MinRetentionRule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *MinRetentionRule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *MinRetentionRule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


