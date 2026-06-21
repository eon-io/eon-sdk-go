# MPARequestPolicyDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredApprovals** | **int32** | Number of approvals required by this policy. | 
**CurrentApprovals** | **int32** | Number of qualifying approvals received for this policy. | 
**ExecutionWindowHours** | **int32** | Hours the requester has to execute the approved action after full approval. | 
**ApprovalWindowHours** | **int32** | Hours allowed for gathering all required approvals after submission. | 
**Description** | Pointer to **string** | Human-readable description of the policy. | [optional] 
**ApproverIdpId** | Pointer to **NullableString** | UUID of the SAML identity provider connection this approver group belongs to. | [optional] 
**ApproverProviderGroupId** | Pointer to **NullableString** | Provider group identifier from the IdP. | [optional] 

## Methods

### NewMPARequestPolicyDetail

`func NewMPARequestPolicyDetail(requiredApprovals int32, currentApprovals int32, executionWindowHours int32, approvalWindowHours int32, ) *MPARequestPolicyDetail`

NewMPARequestPolicyDetail instantiates a new MPARequestPolicyDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPARequestPolicyDetailWithDefaults

`func NewMPARequestPolicyDetailWithDefaults() *MPARequestPolicyDetail`

NewMPARequestPolicyDetailWithDefaults instantiates a new MPARequestPolicyDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredApprovals

`func (o *MPARequestPolicyDetail) GetRequiredApprovals() int32`

GetRequiredApprovals returns the RequiredApprovals field if non-nil, zero value otherwise.

### GetRequiredApprovalsOk

`func (o *MPARequestPolicyDetail) GetRequiredApprovalsOk() (*int32, bool)`

GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovals

`func (o *MPARequestPolicyDetail) SetRequiredApprovals(v int32)`

SetRequiredApprovals sets RequiredApprovals field to given value.


### GetCurrentApprovals

`func (o *MPARequestPolicyDetail) GetCurrentApprovals() int32`

GetCurrentApprovals returns the CurrentApprovals field if non-nil, zero value otherwise.

### GetCurrentApprovalsOk

`func (o *MPARequestPolicyDetail) GetCurrentApprovalsOk() (*int32, bool)`

GetCurrentApprovalsOk returns a tuple with the CurrentApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentApprovals

`func (o *MPARequestPolicyDetail) SetCurrentApprovals(v int32)`

SetCurrentApprovals sets CurrentApprovals field to given value.


### GetExecutionWindowHours

`func (o *MPARequestPolicyDetail) GetExecutionWindowHours() int32`

GetExecutionWindowHours returns the ExecutionWindowHours field if non-nil, zero value otherwise.

### GetExecutionWindowHoursOk

`func (o *MPARequestPolicyDetail) GetExecutionWindowHoursOk() (*int32, bool)`

GetExecutionWindowHoursOk returns a tuple with the ExecutionWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionWindowHours

`func (o *MPARequestPolicyDetail) SetExecutionWindowHours(v int32)`

SetExecutionWindowHours sets ExecutionWindowHours field to given value.


### GetApprovalWindowHours

`func (o *MPARequestPolicyDetail) GetApprovalWindowHours() int32`

GetApprovalWindowHours returns the ApprovalWindowHours field if non-nil, zero value otherwise.

### GetApprovalWindowHoursOk

`func (o *MPARequestPolicyDetail) GetApprovalWindowHoursOk() (*int32, bool)`

GetApprovalWindowHoursOk returns a tuple with the ApprovalWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalWindowHours

`func (o *MPARequestPolicyDetail) SetApprovalWindowHours(v int32)`

SetApprovalWindowHours sets ApprovalWindowHours field to given value.


### GetDescription

`func (o *MPARequestPolicyDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MPARequestPolicyDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MPARequestPolicyDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MPARequestPolicyDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApproverIdpId

`func (o *MPARequestPolicyDetail) GetApproverIdpId() string`

GetApproverIdpId returns the ApproverIdpId field if non-nil, zero value otherwise.

### GetApproverIdpIdOk

`func (o *MPARequestPolicyDetail) GetApproverIdpIdOk() (*string, bool)`

GetApproverIdpIdOk returns a tuple with the ApproverIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIdpId

`func (o *MPARequestPolicyDetail) SetApproverIdpId(v string)`

SetApproverIdpId sets ApproverIdpId field to given value.

### HasApproverIdpId

`func (o *MPARequestPolicyDetail) HasApproverIdpId() bool`

HasApproverIdpId returns a boolean if a field has been set.

### SetApproverIdpIdNil

`func (o *MPARequestPolicyDetail) SetApproverIdpIdNil(b bool)`

 SetApproverIdpIdNil sets the value for ApproverIdpId to be an explicit nil

### UnsetApproverIdpId
`func (o *MPARequestPolicyDetail) UnsetApproverIdpId()`

UnsetApproverIdpId ensures that no value is present for ApproverIdpId, not even an explicit nil
### GetApproverProviderGroupId

`func (o *MPARequestPolicyDetail) GetApproverProviderGroupId() string`

GetApproverProviderGroupId returns the ApproverProviderGroupId field if non-nil, zero value otherwise.

### GetApproverProviderGroupIdOk

`func (o *MPARequestPolicyDetail) GetApproverProviderGroupIdOk() (*string, bool)`

GetApproverProviderGroupIdOk returns a tuple with the ApproverProviderGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverProviderGroupId

`func (o *MPARequestPolicyDetail) SetApproverProviderGroupId(v string)`

SetApproverProviderGroupId sets ApproverProviderGroupId field to given value.

### HasApproverProviderGroupId

`func (o *MPARequestPolicyDetail) HasApproverProviderGroupId() bool`

HasApproverProviderGroupId returns a boolean if a field has been set.

### SetApproverProviderGroupIdNil

`func (o *MPARequestPolicyDetail) SetApproverProviderGroupIdNil(b bool)`

 SetApproverProviderGroupIdNil sets the value for ApproverProviderGroupId to be an explicit nil

### UnsetApproverProviderGroupId
`func (o *MPARequestPolicyDetail) UnsetApproverProviderGroupId()`

UnsetApproverProviderGroupId ensures that no value is present for ApproverProviderGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


