# ActionApprovalRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Action approval rule ID. | 
**ProjectId** | **string** | ID of the project that contains the action approval rule. | 
**Operation** | [**ActionApprovalOperationType**](ActionApprovalOperationType.md) |  | 
**RequiredApprovals** | **int32** | Number of approvals required before the action can be executed. | 
**ApprovalWindowHours** | **int32** | Hours the request stays open for approval before expiring. | 
**ExecutionWindowHours** | **int32** | Hours after approval during which the approved action can be executed. | 
**Description** | Pointer to **string** | Optional description explaining the purpose of this rule. | [optional] 
**ResourceSelector** | Pointer to [**NullableActionApprovalRuleResourceSelector**](ActionApprovalRuleResourceSelector.md) |  | [optional] 
**ApproverIdpId** | Pointer to **NullableString** | UUID of the SAML identity provider connection this approver group belongs to. | [optional] 
**ApproverProviderGroupId** | Pointer to **NullableString** | Provider group identifier from the IdP. | [optional] 
**ExemptApiCredentials** | Pointer to **bool** | When true, API credential users bypass this approval rule. | [optional] 

## Methods

### NewActionApprovalRule

`func NewActionApprovalRule(id string, projectId string, operation ActionApprovalOperationType, requiredApprovals int32, approvalWindowHours int32, executionWindowHours int32, ) *ActionApprovalRule`

NewActionApprovalRule instantiates a new ActionApprovalRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionApprovalRuleWithDefaults

`func NewActionApprovalRuleWithDefaults() *ActionApprovalRule`

NewActionApprovalRuleWithDefaults instantiates a new ActionApprovalRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionApprovalRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionApprovalRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionApprovalRule) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *ActionApprovalRule) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ActionApprovalRule) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ActionApprovalRule) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetOperation

`func (o *ActionApprovalRule) GetOperation() ActionApprovalOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ActionApprovalRule) GetOperationOk() (*ActionApprovalOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ActionApprovalRule) SetOperation(v ActionApprovalOperationType)`

SetOperation sets Operation field to given value.


### GetRequiredApprovals

`func (o *ActionApprovalRule) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *ActionApprovalRule) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *ActionApprovalRule) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.


### GetApprovalWindowHours

`func (o *ActionApprovalRule) GetApprovalWindowHours() int32`

GetApprovalWindowHours returns the ApprovalWindowHours field if non-nil, zero value otherwise.

### GetApprovalWindowHoursOk

`func (o *ActionApprovalRule) GetApprovalWindowHoursOk() (*int32, bool)`

GetApprovalWindowHoursOk returns a tuple with the ApprovalWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWindowHours

`func (o *ActionApprovalRule) SetApprovalWindowHours(v int32)`

SetApprovalWindowHours sets ApprovalWindowHours field to given value.


### GetExecutionWindowHours

`func (o *ActionApprovalRule) GetExecutionWindowHours() int32`

GetExecutionWindowHours returns the ExecutionWindowHours field if non-nil, zero value otherwise.

### GetExecutionWindowHoursOk

`func (o *ActionApprovalRule) GetExecutionWindowHoursOk() (*int32, bool)`

GetExecutionWindowHoursOk returns a tuple with the ExecutionWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionWindowHours

`func (o *ActionApprovalRule) SetExecutionWindowHours(v int32)`

SetExecutionWindowHours sets ExecutionWindowHours field to given value.


### GetDescription

`func (o *ActionApprovalRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionApprovalRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionApprovalRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionApprovalRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSelector

`func (o *ActionApprovalRule) GetResourceSelector() ActionApprovalRuleResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *ActionApprovalRule) GetResourceSelectorOk() (*ActionApprovalRuleResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *ActionApprovalRule) SetResourceSelector(v ActionApprovalRuleResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *ActionApprovalRule) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### SetResourceSelectorNil

`func (o *ActionApprovalRule) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *ActionApprovalRule) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetApproverIdpId

`func (o *ActionApprovalRule) GetApproverIdpId() string`

GetApproverIdpId returns the ApproverIdpId field if non-nil, zero value otherwise.

### GetApproverIdpIdOk

`func (o *ActionApprovalRule) GetApproverIdpIdOk() (*string, bool)`

GetApproverIdpIdOk returns a tuple with the ApproverIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIdpId

`func (o *ActionApprovalRule) SetApproverIdpId(v string)`

SetApproverIdpId sets ApproverIdpId field to given value.

### HasApproverIdpId

`func (o *ActionApprovalRule) HasApproverIdpId() bool`

HasApproverIdpId returns a boolean if a field has been set.

### SetApproverIdpIdNil

`func (o *ActionApprovalRule) SetApproverIdpIdNil(b bool)`

 SetApproverIdpIdNil sets the value for ApproverIdpId to be an explicit nil

### UnsetApproverIdpId
`func (o *ActionApprovalRule) UnsetApproverIdpId()`

UnsetApproverIdpId ensures that no value is present for ApproverIdpId, not even an explicit nil
### GetApproverProviderGroupId

`func (o *ActionApprovalRule) GetApproverProviderGroupId() string`

GetApproverProviderGroupId returns the ApproverProviderGroupId field if non-nil, zero value otherwise.

### GetApproverProviderGroupIdOk

`func (o *ActionApprovalRule) GetApproverProviderGroupIdOk() (*string, bool)`

GetApproverProviderGroupIdOk returns a tuple with the ApproverProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverProviderGroupId

`func (o *ActionApprovalRule) SetApproverProviderGroupId(v string)`

SetApproverProviderGroupId sets ApproverProviderGroupId field to given value.

### HasApproverProviderGroupId

`func (o *ActionApprovalRule) HasApproverProviderGroupId() bool`

HasApproverProviderGroupId returns a boolean if a field has been set.

### SetApproverProviderGroupIdNil

`func (o *ActionApprovalRule) SetApproverProviderGroupIdNil(b bool)`

 SetApproverProviderGroupIdNil sets the value for ApproverProviderGroupId to be an explicit nil

### UnsetApproverProviderGroupId
`func (o *ActionApprovalRule) UnsetApproverProviderGroupId()`

UnsetApproverProviderGroupId ensures that no value is present for ApproverProviderGroupId, not even an explicit nil
### GetExemptApiCredentials

`func (o *ActionApprovalRule) GetExemptApiCredentials() bool`

GetExemptApiCredentials returns the ExemptApiCredentials field if non-nil, zero value otherwise.

### GetExemptApiCredentialsOk

`func (o *ActionApprovalRule) GetExemptApiCredentialsOk() (*bool, bool)`

GetExemptApiCredentialsOk returns a tuple with the ExemptApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptApiCredentials

`func (o *ActionApprovalRule) SetExemptApiCredentials(v bool)`

SetExemptApiCredentials sets ExemptApiCredentials field to given value.

### HasExemptApiCredentials

`func (o *ActionApprovalRule) HasExemptApiCredentials() bool`

HasExemptApiCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


