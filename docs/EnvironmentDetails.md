# EnvironmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**IsOverridden** | Pointer to **bool** | Whether the environment is manually overridden. If &#x60;true&#x60;, the environment is user-defined and remains static. If &#x60;false&#x60;, the environment is automatically detected and set by Eon.  | [optional] 

## Methods

### NewEnvironmentDetails

`func NewEnvironmentDetails() *EnvironmentDetails`

NewEnvironmentDetails instantiates a new EnvironmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDetailsWithDefaults

`func NewEnvironmentDetailsWithDefaults() *EnvironmentDetails`

NewEnvironmentDetailsWithDefaults instantiates a new EnvironmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *EnvironmentDetails) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentDetails) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentDetails) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentDetails) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsOverridden

`func (o *EnvironmentDetails) GetIsOverridden() bool`

GetIsOverridden returns the IsOverridden field if non-nil, zero value otherwise.

### GetIsOverriddenOk

`func (o *EnvironmentDetails) GetIsOverriddenOk() (*bool, bool)`

GetIsOverriddenOk returns a tuple with the IsOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridden

`func (o *EnvironmentDetails) SetIsOverridden(v bool)`

SetIsOverridden sets IsOverridden field to given value.

### HasIsOverridden

`func (o *EnvironmentDetails) HasIsOverridden() bool`

HasIsOverridden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


