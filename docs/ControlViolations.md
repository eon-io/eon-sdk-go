# ControlViolations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High** | **int32** | Number of high-severity control violations. | 
**Medium** | **int32** | Number of medium-severity control violations. | 
**Low** | **int32** | Number of low-severity control violations. | 
**Muted** | **int32** | Number of muted control violations. | 

## Methods

### NewControlViolations

`func NewControlViolations(high int32, medium int32, low int32, muted int32, ) *ControlViolations`

NewControlViolations instantiates a new ControlViolations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlViolationsWithDefaults

`func NewControlViolationsWithDefaults() *ControlViolations`

NewControlViolationsWithDefaults instantiates a new ControlViolations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh

`func (o *ControlViolations) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ControlViolations) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ControlViolations) SetHigh(v int32)`

SetHigh sets High field to given value.


### GetMedium

`func (o *ControlViolations) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ControlViolations) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ControlViolations) SetMedium(v int32)`

SetMedium sets Medium field to given value.


### GetLow

`func (o *ControlViolations) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ControlViolations) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ControlViolations) SetLow(v int32)`

SetLow sets Low field to given value.


### GetMuted

`func (o *ControlViolations) GetMuted() int32`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ControlViolations) GetMutedOk() (*int32, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ControlViolations) SetMuted(v int32)`

SetMuted sets Muted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


