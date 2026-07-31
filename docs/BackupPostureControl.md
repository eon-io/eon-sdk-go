# BackupPostureControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Server-assigned unique identifier of the control. | 
**Name** | **string** | Human-readable name of the control. | 
**Severity** | [**Severity**](Severity.md) |  | 
**ResourceSelector** | [**NullableBackupPostureControlResourceSelector**](BackupPostureControlResourceSelector.md) |  | 
**Rules** | [**NullableBackupPostureControlRules**](BackupPostureControlRules.md) |  | 

## Methods

### NewBackupPostureControl

`func NewBackupPostureControl(id string, name string, severity Severity, resourceSelector NullableBackupPostureControlResourceSelector, rules NullableBackupPostureControlRules, ) *BackupPostureControl`

NewBackupPostureControl instantiates a new BackupPostureControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPostureControlWithDefaults

`func NewBackupPostureControlWithDefaults() *BackupPostureControl`

NewBackupPostureControlWithDefaults instantiates a new BackupPostureControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupPostureControl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupPostureControl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupPostureControl) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BackupPostureControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupPostureControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupPostureControl) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *BackupPostureControl) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BackupPostureControl) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BackupPostureControl) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.


### GetResourceSelector

`func (o *BackupPostureControl) GetResourceSelector() BackupPostureControlResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *BackupPostureControl) GetResourceSelectorOk() (*BackupPostureControlResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *BackupPostureControl) SetResourceSelector(v BackupPostureControlResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.


### SetResourceSelectorNil

`func (o *BackupPostureControl) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *BackupPostureControl) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetRules

`func (o *BackupPostureControl) GetRules() BackupPostureControlRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BackupPostureControl) GetRulesOk() (*BackupPostureControlRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BackupPostureControl) SetRules(v BackupPostureControlRules)`

SetRules sets Rules field to given value.


### SetRulesNil

`func (o *BackupPostureControl) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BackupPostureControl) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


