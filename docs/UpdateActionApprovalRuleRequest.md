# UpdateActionApprovalRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredApprovals** | Pointer to **int32** | Number of approvals required before the action can be executed. | [optional] 
**ApprovalWindowHours** | Pointer to **int32** | Hours the request stays open for approval before expiring. | [optional] 
**ExecutionWindowHours** | Pointer to **int32** | Hours after approval during which the approved action can be executed. | [optional] 
**Description** | Pointer to **string** | Optional description explaining the purpose of this rule. | [optional] 
**ResourceSelector** | Pointer to [**NullableActionApprovalRuleResourceSelector**](ActionApprovalRuleResourceSelector.md) |  | [optional] 
**ApproverIdpId** | Pointer to **NullableString** | UUID of the SAML identity provider connection this approver group belongs to. Set to null to remove the restriction. | [optional] 
**ApproverProviderGroupId** | Pointer to **NullableString** | Provider group identifier from the IdP. Set to null to remove. | [optional] 
**ExemptApiCredentials** | Pointer to **bool** | When true, API credential users bypass this approval rule. | [optional] 

## Methods

### NewUpdateActionApprovalRuleRequest

`func NewUpdateActionApprovalRuleRequest() *UpdateActionApprovalRuleRequest`

NewUpdateActionApprovalRuleRequest instantiates a new UpdateActionApprovalRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActionApprovalRuleRequestWithDefaults

`func NewUpdateActionApprovalRuleRequestWithDefaults() *UpdateActionApprovalRuleRequest`

NewUpdateActionApprovalRuleRequestWithDefaults instantiates a new UpdateActionApprovalRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredApprovals

`func (o *UpdateActionApprovalRuleRequest) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *UpdateActionApprovalRuleRequest) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *UpdateActionApprovalRuleRequest) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *UpdateActionApprovalRuleRequest) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.

### GetApprovalWindowHours

`func (o *UpdateActionApprovalRuleRequest) GetApprovalWindowHours() int32`

GetApprovalWindowHours returns the ApprovalWindowHours field if non-nil, zero value otherwise.

### GetApprovalWindowHoursOk

`func (o *UpdateActionApprovalRuleRequest) GetApprovalWindowHoursOk() (*int32, bool)`

GetApprovalWindowHoursOk returns a tuple with the ApprovalWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWindowHours

`func (o *UpdateActionApprovalRuleRequest) SetApprovalWindowHours(v int32)`

SetApprovalWindowHours sets ApprovalWindowHours field to given value.

### HasApprovalWindowHours

`func (o *UpdateActionApprovalRuleRequest) HasApprovalWindowHours() bool`

HasApprovalWindowHours returns a boolean if a field has been set.

### GetExecutionWindowHours

`func (o *UpdateActionApprovalRuleRequest) GetExecutionWindowHours() int32`

GetExecutionWindowHours returns the ExecutionWindowHours field if non-nil, zero value otherwise.

### GetExecutionWindowHoursOk

`func (o *UpdateActionApprovalRuleRequest) GetExecutionWindowHoursOk() (*int32, bool)`

GetExecutionWindowHoursOk returns a tuple with the ExecutionWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionWindowHours

`func (o *UpdateActionApprovalRuleRequest) SetExecutionWindowHours(v int32)`

SetExecutionWindowHours sets ExecutionWindowHours field to given value.

### HasExecutionWindowHours

`func (o *UpdateActionApprovalRuleRequest) HasExecutionWindowHours() bool`

HasExecutionWindowHours returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateActionApprovalRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateActionApprovalRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateActionApprovalRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateActionApprovalRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSelector

`func (o *UpdateActionApprovalRuleRequest) GetResourceSelector() ActionApprovalRuleResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *UpdateActionApprovalRuleRequest) GetResourceSelectorOk() (*ActionApprovalRuleResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *UpdateActionApprovalRuleRequest) SetResourceSelector(v ActionApprovalRuleResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *UpdateActionApprovalRuleRequest) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### SetResourceSelectorNil

`func (o *UpdateActionApprovalRuleRequest) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *UpdateActionApprovalRuleRequest) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetApproverIdpId

`func (o *UpdateActionApprovalRuleRequest) GetApproverIdpId() string`

GetApproverIdpId returns the ApproverIdpId field if non-nil, zero value otherwise.

### GetApproverIdpIdOk

`func (o *UpdateActionApprovalRuleRequest) GetApproverIdpIdOk() (*string, bool)`

GetApproverIdpIdOk returns a tuple with the ApproverIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIdpId

`func (o *UpdateActionApprovalRuleRequest) SetApproverIdpId(v string)`

SetApproverIdpId sets ApproverIdpId field to given value.

### HasApproverIdpId

`func (o *UpdateActionApprovalRuleRequest) HasApproverIdpId() bool`

HasApproverIdpId returns a boolean if a field has been set.

### SetApproverIdpIdNil

`func (o *UpdateActionApprovalRuleRequest) SetApproverIdpIdNil(b bool)`

 SetApproverIdpIdNil sets the value for ApproverIdpId to be an explicit nil

### UnsetApproverIdpId
`func (o *UpdateActionApprovalRuleRequest) UnsetApproverIdpId()`

UnsetApproverIdpId ensures that no value is present for ApproverIdpId, not even an explicit nil
### GetApproverProviderGroupId

`func (o *UpdateActionApprovalRuleRequest) GetApproverProviderGroupId() string`

GetApproverProviderGroupId returns the ApproverProviderGroupId field if non-nil, zero value otherwise.

### GetApproverProviderGroupIdOk

`func (o *UpdateActionApprovalRuleRequest) GetApproverProviderGroupIdOk() (*string, bool)`

GetApproverProviderGroupIdOk returns a tuple with the ApproverProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverProviderGroupId

`func (o *UpdateActionApprovalRuleRequest) SetApproverProviderGroupId(v string)`

SetApproverProviderGroupId sets ApproverProviderGroupId field to given value.

### HasApproverProviderGroupId

`func (o *UpdateActionApprovalRuleRequest) HasApproverProviderGroupId() bool`

HasApproverProviderGroupId returns a boolean if a field has been set.

### SetApproverProviderGroupIdNil

`func (o *UpdateActionApprovalRuleRequest) SetApproverProviderGroupIdNil(b bool)`

 SetApproverProviderGroupIdNil sets the value for ApproverProviderGroupId to be an explicit nil

### UnsetApproverProviderGroupId
`func (o *UpdateActionApprovalRuleRequest) UnsetApproverProviderGroupId()`

UnsetApproverProviderGroupId ensures that no value is present for ApproverProviderGroupId, not even an explicit nil
### GetExemptApiCredentials

`func (o *UpdateActionApprovalRuleRequest) GetExemptApiCredentials() bool`

GetExemptApiCredentials returns the ExemptApiCredentials field if non-nil, zero value otherwise.

### GetExemptApiCredentialsOk

`func (o *UpdateActionApprovalRuleRequest) GetExemptApiCredentialsOk() (*bool, bool)`

GetExemptApiCredentialsOk returns a tuple with the ExemptApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptApiCredentials

`func (o *UpdateActionApprovalRuleRequest) SetExemptApiCredentials(v bool)`

SetExemptApiCredentials sets ExemptApiCredentials field to given value.

### HasExemptApiCredentials

`func (o *UpdateActionApprovalRuleRequest) HasExemptApiCredentials() bool`

HasExemptApiCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


