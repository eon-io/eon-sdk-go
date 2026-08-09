# ActionApprovalRuleResourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSelectionMode** | [**ResourceSelectorMode**](ResourceSelectorMode.md) |  | 
**ResourceInclusionOverride** | Pointer to **[]string** | Resource identifiers to protect, regardless of resourceSelectionMode. When resourceSelectionMode is NONE, only these resources are protected. For cloud-resource-scoped operations these are provider resource IDs (e.g. AWS instance IDs); for backup-policy-scoped operations these are Eon backup policy UUIDs.  | [optional] 
**ResourceExclusionOverride** | Pointer to **[]string** | Resource identifiers to exclude from protection, regardless of resourceSelectionMode. When resourceSelectionMode is ALL, all resources except these are protected. For cloud-resource-scoped operations these are provider resource IDs; for backup-policy-scoped operations these are Eon backup policy UUIDs.  | [optional] 
**Expression** | Pointer to [**NullableActionApprovalRuleExpression**](ActionApprovalRuleExpression.md) |  | [optional] 

## Methods

### NewActionApprovalRuleResourceSelector

`func NewActionApprovalRuleResourceSelector(resourceSelectionMode ResourceSelectorMode, ) *ActionApprovalRuleResourceSelector`

NewActionApprovalRuleResourceSelector instantiates a new ActionApprovalRuleResourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionApprovalRuleResourceSelectorWithDefaults

`func NewActionApprovalRuleResourceSelectorWithDefaults() *ActionApprovalRuleResourceSelector`

NewActionApprovalRuleResourceSelectorWithDefaults instantiates a new ActionApprovalRuleResourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSelectionMode

`func (o *ActionApprovalRuleResourceSelector) GetResourceSelectionMode() ResourceSelectorMode`

GetResourceSelectionMode returns the ResourceSelectionMode field if non-nil, zero value otherwise.

### GetResourceSelectionModeOk

`func (o *ActionApprovalRuleResourceSelector) GetResourceSelectionModeOk() (*ResourceSelectorMode, bool)`

GetResourceSelectionModeOk returns a tuple with the ResourceSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelectionMode

`func (o *ActionApprovalRuleResourceSelector) SetResourceSelectionMode(v ResourceSelectorMode)`

SetResourceSelectionMode sets ResourceSelectionMode field to given value.


### GetResourceInclusionOverride

`func (o *ActionApprovalRuleResourceSelector) GetResourceInclusionOverride() []string`

GetResourceInclusionOverride returns the ResourceInclusionOverride field if non-nil, zero value otherwise.

### GetResourceInclusionOverrideOk

`func (o *ActionApprovalRuleResourceSelector) GetResourceInclusionOverrideOk() (*[]string, bool)`

GetResourceInclusionOverrideOk returns a tuple with the ResourceInclusionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInclusionOverride

`func (o *ActionApprovalRuleResourceSelector) SetResourceInclusionOverride(v []string)`

SetResourceInclusionOverride sets ResourceInclusionOverride field to given value.

### HasResourceInclusionOverride

`func (o *ActionApprovalRuleResourceSelector) HasResourceInclusionOverride() bool`

HasResourceInclusionOverride returns a boolean if a field has been set.

### GetResourceExclusionOverride

`func (o *ActionApprovalRuleResourceSelector) GetResourceExclusionOverride() []string`

GetResourceExclusionOverride returns the ResourceExclusionOverride field if non-nil, zero value otherwise.

### GetResourceExclusionOverrideOk

`func (o *ActionApprovalRuleResourceSelector) GetResourceExclusionOverrideOk() (*[]string, bool)`

GetResourceExclusionOverrideOk returns a tuple with the ResourceExclusionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExclusionOverride

`func (o *ActionApprovalRuleResourceSelector) SetResourceExclusionOverride(v []string)`

SetResourceExclusionOverride sets ResourceExclusionOverride field to given value.

### HasResourceExclusionOverride

`func (o *ActionApprovalRuleResourceSelector) HasResourceExclusionOverride() bool`

HasResourceExclusionOverride returns a boolean if a field has been set.

### GetExpression

`func (o *ActionApprovalRuleResourceSelector) GetExpression() ActionApprovalRuleExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ActionApprovalRuleResourceSelector) GetExpressionOk() (*ActionApprovalRuleExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ActionApprovalRuleResourceSelector) SetExpression(v ActionApprovalRuleExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ActionApprovalRuleResourceSelector) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *ActionApprovalRuleResourceSelector) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *ActionApprovalRuleResourceSelector) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


