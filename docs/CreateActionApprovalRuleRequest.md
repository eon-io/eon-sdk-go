# CreateActionApprovalRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | [**ActionApprovalOperationType**](ActionApprovalOperationType.md) |  | 
**RequiredApprovals** | Pointer to **int32** | Number of approvals required before the action can be executed. | [optional] [default to 1]
**ApprovalWindowHours** | **int32** | Hours the request stays open for approval before expiring. | 
**ExecutionWindowHours** | **int32** | Hours after approval during which the approved action can be executed. | 
**Description** | Pointer to **string** | Optional description explaining the purpose of this rule. | [optional] 
**ResourceSelector** | Pointer to [**NullableActionApprovalRuleResourceSelector**](ActionApprovalRuleResourceSelector.md) |  | [optional] 
**ApproverIdpId** | Pointer to **string** | UUID of the SAML identity provider connection this approver group belongs to. | [optional] 
**ApproverProviderGroupId** | Pointer to **string** | Provider group identifier from the IdP. | [optional] 
**ExemptApiCredentials** | Pointer to **bool** | When true, API credential users bypass this approval rule. | [optional] 

## Methods

### NewCreateActionApprovalRuleRequest

`func NewCreateActionApprovalRuleRequest(operation ActionApprovalOperationType, approvalWindowHours int32, executionWindowHours int32, ) *CreateActionApprovalRuleRequest`

NewCreateActionApprovalRuleRequest instantiates a new CreateActionApprovalRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActionApprovalRuleRequestWithDefaults

`func NewCreateActionApprovalRuleRequestWithDefaults() *CreateActionApprovalRuleRequest`

NewCreateActionApprovalRuleRequestWithDefaults instantiates a new CreateActionApprovalRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *CreateActionApprovalRuleRequest) GetOperation() ActionApprovalOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateActionApprovalRuleRequest) GetOperationOk() (*ActionApprovalOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateActionApprovalRuleRequest) SetOperation(v ActionApprovalOperationType)`

SetOperation sets Operation field to given value.


### GetRequiredApprovals

`func (o *CreateActionApprovalRuleRequest) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *CreateActionApprovalRuleRequest) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *CreateActionApprovalRuleRequest) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.

### HasRequiredApprovals

`func (o *CreateActionApprovalRuleRequest) HasRequiredApprovals() bool`

HasRequiredApprovals returns a boolean if a field has been set.

### GetApprovalWindowHours

`func (o *CreateActionApprovalRuleRequest) GetApprovalWindowHours() int32`

GetApprovalWindowHours returns the ApprovalWindowHours field if non-nil, zero value otherwise.

### GetApprovalWindowHoursOk

`func (o *CreateActionApprovalRuleRequest) GetApprovalWindowHoursOk() (*int32, bool)`

GetApprovalWindowHoursOk returns a tuple with the ApprovalWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWindowHours

`func (o *CreateActionApprovalRuleRequest) SetApprovalWindowHours(v int32)`

SetApprovalWindowHours sets ApprovalWindowHours field to given value.


### GetExecutionWindowHours

`func (o *CreateActionApprovalRuleRequest) GetExecutionWindowHours() int32`

GetExecutionWindowHours returns the ExecutionWindowHours field if non-nil, zero value otherwise.

### GetExecutionWindowHoursOk

`func (o *CreateActionApprovalRuleRequest) GetExecutionWindowHoursOk() (*int32, bool)`

GetExecutionWindowHoursOk returns a tuple with the ExecutionWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionWindowHours

`func (o *CreateActionApprovalRuleRequest) SetExecutionWindowHours(v int32)`

SetExecutionWindowHours sets ExecutionWindowHours field to given value.


### GetDescription

`func (o *CreateActionApprovalRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateActionApprovalRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateActionApprovalRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateActionApprovalRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceSelector

`func (o *CreateActionApprovalRuleRequest) GetResourceSelector() ActionApprovalRuleResourceSelector`

GetResourceSelector returns the ResourceSelector field if non-nil, zero value otherwise.

### GetResourceSelectorOk

`func (o *CreateActionApprovalRuleRequest) GetResourceSelectorOk() (*ActionApprovalRuleResourceSelector, bool)`

GetResourceSelectorOk returns a tuple with the ResourceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelector

`func (o *CreateActionApprovalRuleRequest) SetResourceSelector(v ActionApprovalRuleResourceSelector)`

SetResourceSelector sets ResourceSelector field to given value.

### HasResourceSelector

`func (o *CreateActionApprovalRuleRequest) HasResourceSelector() bool`

HasResourceSelector returns a boolean if a field has been set.

### SetResourceSelectorNil

`func (o *CreateActionApprovalRuleRequest) SetResourceSelectorNil(b bool)`

 SetResourceSelectorNil sets the value for ResourceSelector to be an explicit nil

### UnsetResourceSelector
`func (o *CreateActionApprovalRuleRequest) UnsetResourceSelector()`

UnsetResourceSelector ensures that no value is present for ResourceSelector, not even an explicit nil
### GetApproverIdpId

`func (o *CreateActionApprovalRuleRequest) GetApproverIdpId() string`

GetApproverIdpId returns the ApproverIdpId field if non-nil, zero value otherwise.

### GetApproverIdpIdOk

`func (o *CreateActionApprovalRuleRequest) GetApproverIdpIdOk() (*string, bool)`

GetApproverIdpIdOk returns a tuple with the ApproverIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIdpId

`func (o *CreateActionApprovalRuleRequest) SetApproverIdpId(v string)`

SetApproverIdpId sets ApproverIdpId field to given value.

### HasApproverIdpId

`func (o *CreateActionApprovalRuleRequest) HasApproverIdpId() bool`

HasApproverIdpId returns a boolean if a field has been set.

### GetApproverProviderGroupId

`func (o *CreateActionApprovalRuleRequest) GetApproverProviderGroupId() string`

GetApproverProviderGroupId returns the ApproverProviderGroupId field if non-nil, zero value otherwise.

### GetApproverProviderGroupIdOk

`func (o *CreateActionApprovalRuleRequest) GetApproverProviderGroupIdOk() (*string, bool)`

GetApproverProviderGroupIdOk returns a tuple with the ApproverProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverProviderGroupId

`func (o *CreateActionApprovalRuleRequest) SetApproverProviderGroupId(v string)`

SetApproverProviderGroupId sets ApproverProviderGroupId field to given value.

### HasApproverProviderGroupId

`func (o *CreateActionApprovalRuleRequest) HasApproverProviderGroupId() bool`

HasApproverProviderGroupId returns a boolean if a field has been set.

### GetExemptApiCredentials

`func (o *CreateActionApprovalRuleRequest) GetExemptApiCredentials() bool`

GetExemptApiCredentials returns the ExemptApiCredentials field if non-nil, zero value otherwise.

### GetExemptApiCredentialsOk

`func (o *CreateActionApprovalRuleRequest) GetExemptApiCredentialsOk() (*bool, bool)`

GetExemptApiCredentialsOk returns a tuple with the ExemptApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptApiCredentials

`func (o *CreateActionApprovalRuleRequest) SetExemptApiCredentials(v bool)`

SetExemptApiCredentials sets ExemptApiCredentials field to given value.

### HasExemptApiCredentials

`func (o *CreateActionApprovalRuleRequest) HasExemptApiCredentials() bool`

HasExemptApiCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


