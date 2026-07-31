# CreateBackupPostureControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name of the control. | 
**Severity** | [**Severity**](Severity.md) |  | 
**ResourceSelector** | [**NullableBackupPostureControlResourceSelector**](BackupPostureControlResourceSelector.md) |  | 
**Rules** | [**NullableBackupPostureControlRules**](BackupPostureControlRules.md) |  | 

## Methods

### NewCreateBackupPostureControlRequest

`func NewCreateBackupPostureControlRequest(name string, severity Severity, resourceSelector NullableBackupPostureControlResourceSelector, rules NullableBackupPostureControlRules, ) *CreateBackupPostureControlRequest`

NewCreateBackupPostureControlRequest instantiates a new CreateBackupPostureControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupPostureControlRequestWithDefaults

`func NewCreateBackupPostureControlRequestWithDefaults() *CreateBackupPostureControlRequest`

NewCreateBackupPostureControlRequestWithDefaults instantiates a new CreateBackupPostureControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBackupPostureControlRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBackupPostureControlRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBackupPostureControlRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *CreateBackupPostureControlRequest) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateBackupPostureControlRequest) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateBackupPostureControlRequest) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetResourceSelector

`func (o *CreateBackupPostureControlRequest) GetResourceSelector() BackupPostureControlResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *CreateBackupPostureControlRequest) GetResourceSelectorOk() (*BackupPostureControlResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *CreateBackupPostureControlRequest) SetResourceSelector(v BackupPostureControlResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### SetResourceSelectorNil

`func (o *CreateBackupPostureControlRequest) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *CreateBackupPostureControlRequest) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetRules

`func (o *CreateBackupPostureControlRequest) GetRules() BackupPostureControlRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateBackupPostureControlRequest) GetRulesOk() (*BackupPostureControlRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateBackupPostureControlRequest) SetRules(v BackupPostureControlRules)`

SetRules sets Rules field to given value.


### SetRulesNil

`func (o *CreateBackupPostureControlRequest) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *CreateBackupPostureControlRequest) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


