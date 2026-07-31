# BackupPostureControlResourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSelectionMode** | [**ResourceSelectorMode**](ResourceSelectorMode.md) |  | 
**Expression** | Pointer to [**NullableBackupPostureControlExpression**](BackupPostureControlExpression.md) |  | [optional] 

## Methods

### NewBackupPostureControlResourceSelector

`func NewBackupPostureControlResourceSelector(resourceSelectionMode ResourceSelectorMode, ) *BackupPostureControlResourceSelector`

NewBackupPostureControlResourceSelector instantiates a new BackupPostureControlResourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPostureControlResourceSelectorWithDefaults

`func NewBackupPostureControlResourceSelectorWithDefaults() *BackupPostureControlResourceSelector`

NewBackupPostureControlResourceSelectorWithDefaults instantiates a new BackupPostureControlResourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSelectionMode

`func (o *BackupPostureControlResourceSelector) GetResourceSelectionMode() ResourceSelectorMode`

GetResourceSelectionMode returns the ResourceSelectionMode field if non-nil, zero value otherwise.

### GetResourceSelectionModeOk

`func (o *BackupPostureControlResourceSelector) GetResourceSelectionModeOk() (*ResourceSelectorMode, bool)`

GetResourceSelectionModeOk returns a tuple with the ResourceSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelectionMode

`func (o *BackupPostureControlResourceSelector) SetResourceSelectionMode(v ResourceSelectorMode)`

SetResourceSelectionMode sets ResourceSelectionMode field to given value.


### GetExpression

`func (o *BackupPostureControlResourceSelector) GetExpression() BackupPostureControlExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BackupPostureControlResourceSelector) GetExpressionOk() (*BackupPostureControlExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BackupPostureControlResourceSelector) SetExpression(v BackupPostureControlExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *BackupPostureControlResourceSelector) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *BackupPostureControlResourceSelector) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *BackupPostureControlResourceSelector) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


