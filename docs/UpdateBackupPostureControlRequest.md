# UpdateBackupPostureControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name of the control. | 
**Severity** | [**Severity**](Severity.md) |  | 
**ResourceSelector** | [**NullableBackupPostureControlResourceSelector**](BackupPostureControlResourceSelector.md) |  | 
**Rules** | [**NullableBackupPostureControlRules**](BackupPostureControlRules.md) |  | 

## Methods

### NewUpdateBackupPostureControlRequest

`func NewUpdateBackupPostureControlRequest(name string, severity Severity, resourceSelector NullableBackupPostureControlResourceSelector, rules NullableBackupPostureControlRules, ) *UpdateBackupPostureControlRequest`

NewUpdateBackupPostureControlRequest instantiates a new UpdateBackupPostureControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupPostureControlRequestWithDefaults

`func NewUpdateBackupPostureControlRequestWithDefaults() *UpdateBackupPostureControlRequest`

NewUpdateBackupPostureControlRequestWithDefaults instantiates a new UpdateBackupPostureControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBackupPostureControlRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBackupPostureControlRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBackupPostureControlRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *UpdateBackupPostureControlRequest) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateBackupPostureControlRequest) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateBackupPostureControlRequest) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetResourceSelector

`func (o *UpdateBackupPostureControlRequest) GetResourceSelector() BackupPostureControlResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *UpdateBackupPostureControlRequest) GetResourceSelectorOk() (*BackupPostureControlResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *UpdateBackupPostureControlRequest) SetResourceSelector(v BackupPostureControlResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### SetResourceSelectorNil

`func (o *UpdateBackupPostureControlRequest) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *UpdateBackupPostureControlRequest) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetRules

`func (o *UpdateBackupPostureControlRequest) GetRules() BackupPostureControlRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateBackupPostureControlRequest) GetRulesOk() (*BackupPostureControlRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateBackupPostureControlRequest) SetRules(v BackupPostureControlRules)`

SetRules sets Rules field to given value.


### SetRulesNil

`func (o *UpdateBackupPostureControlRequest) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *UpdateBackupPostureControlRequest) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


