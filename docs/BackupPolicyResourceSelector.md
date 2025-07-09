# BackupPolicyResourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSelectionMode** | [**ResourceSelectorMode**](ResourceSelectorMode.md) |  | 
**Expression** | Pointer to [**NullableBackupPolicyExpression**](BackupPolicyExpression.md) |  | [optional] 
**ResourceInclusionOverride** | Pointer to **[]string** | List of cloud-provider-assigned resource IDs to include in the backup policy, regardless of whether they&#39;re excluded by &#x60;resourceSelectionMode&#x60; and &#x60;expression&#x60;.  | [optional] 
**ResourceExclusionOverride** | Pointer to **[]string** | List of cloud-provider-assigned resource IDs to exclude from the backup policy, regardless of whether they&#39;re included by &#x60;resourceSelectionMode&#x60; and &#x60;expression&#x60;.  | [optional] 

## Methods

### NewBackupPolicyResourceSelector

`func NewBackupPolicyResourceSelector(resourceSelectionMode ResourceSelectorMode, ) *BackupPolicyResourceSelector`

NewBackupPolicyResourceSelector instantiates a new BackupPolicyResourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyResourceSelectorWithDefaults

`func NewBackupPolicyResourceSelectorWithDefaults() *BackupPolicyResourceSelector`

NewBackupPolicyResourceSelectorWithDefaults instantiates a new BackupPolicyResourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSelectionMode

`func (o *BackupPolicyResourceSelector) GetResourceSelectionMode() ResourceSelectorMode`

GetResourceSelectionMode returns the ResourceSelectionMode field if non-nil, zero value otherwise.

### GetResourceSelectionModeOk

`func (o *BackupPolicyResourceSelector) GetResourceSelectionModeOk() (*ResourceSelectorMode, bool)`

GetResourceSelectionModeOk returns a tuple with the ResourceSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelectionMode

`func (o *BackupPolicyResourceSelector) SetResourceSelectionMode(v ResourceSelectorMode)`

SetResourceSelectionMode sets ResourceSelectionMode field to given value.


### GetExpression

`func (o *BackupPolicyResourceSelector) GetExpression() BackupPolicyExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BackupPolicyResourceSelector) GetExpressionOk() (*BackupPolicyExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BackupPolicyResourceSelector) SetExpression(v BackupPolicyExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *BackupPolicyResourceSelector) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *BackupPolicyResourceSelector) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *BackupPolicyResourceSelector) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil
### GetResourceInclusionOverride

`func (o *BackupPolicyResourceSelector) GetResourceInclusionOverride() []string`

GetResourceInclusionOverride returns the ResourceInclusionOverride field if non-nil, zero value otherwise.

### GetResourceInclusionOverrideOk

`func (o *BackupPolicyResourceSelector) GetResourceInclusionOverrideOk() (*[]string, bool)`

GetResourceInclusionOverrideOk returns a tuple with the ResourceInclusionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInclusionOverride

`func (o *BackupPolicyResourceSelector) SetResourceInclusionOverride(v []string)`

SetResourceInclusionOverride sets ResourceInclusionOverride field to given value.

### HasResourceInclusionOverride

`func (o *BackupPolicyResourceSelector) HasResourceInclusionOverride() bool`

HasResourceInclusionOverride returns a boolean if a field has been set.

### GetResourceExclusionOverride

`func (o *BackupPolicyResourceSelector) GetResourceExclusionOverride() []string`

GetResourceExclusionOverride returns the ResourceExclusionOverride field if non-nil, zero value otherwise.

### GetResourceExclusionOverrideOk

`func (o *BackupPolicyResourceSelector) GetResourceExclusionOverrideOk() (*[]string, bool)`

GetResourceExclusionOverrideOk returns a tuple with the ResourceExclusionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExclusionOverride

`func (o *BackupPolicyResourceSelector) SetResourceExclusionOverride(v []string)`

SetResourceExclusionOverride sets ResourceExclusionOverride field to given value.

### HasResourceExclusionOverride

`func (o *BackupPolicyResourceSelector) HasResourceExclusionOverride() bool`

HasResourceExclusionOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


