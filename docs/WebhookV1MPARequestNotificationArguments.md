# WebhookV1MPARequestNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Approval request ID. | [optional] 
**Operation** | Pointer to **string** | Operation that requires approval. | [optional] 
**RequesterEmail** | Pointer to **string** | Email address of the requester. | [optional] 
**PolicyDescription** | Pointer to **string** | Description of the approval policy. | [optional] 
**ApprovalExpiresAt** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 

## Methods

### NewWebhookV1MPARequestNotificationArguments

`func NewWebhookV1MPARequestNotificationArguments() *WebhookV1MPARequestNotificationArguments`

NewWebhookV1MPARequestNotificationArguments instantiates a new WebhookV1MPARequestNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1MPARequestNotificationArgumentsWithDefaults

`func NewWebhookV1MPARequestNotificationArgumentsWithDefaults() *WebhookV1MPARequestNotificationArguments`

NewWebhookV1MPARequestNotificationArgumentsWithDefaults instantiates a new WebhookV1MPARequestNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *WebhookV1MPARequestNotificationArguments) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookV1MPARequestNotificationArguments) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookV1MPARequestNotificationArguments) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WebhookV1MPARequestNotificationArguments) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetOperation

`func (o *WebhookV1MPARequestNotificationArguments) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *WebhookV1MPARequestNotificationArguments) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *WebhookV1MPARequestNotificationArguments) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *WebhookV1MPARequestNotificationArguments) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRequesterEmail

`func (o *WebhookV1MPARequestNotificationArguments) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *WebhookV1MPARequestNotificationArguments) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *WebhookV1MPARequestNotificationArguments) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.

### HasRequesterEmail

`func (o *WebhookV1MPARequestNotificationArguments) HasRequesterEmail() bool`

HasRequesterEmail returns a boolean if a field has been set.

### GetPolicyDescription

`func (o *WebhookV1MPARequestNotificationArguments) GetPolicyDescription() string`

GetPolicyDescription returns the PolicyDescription field if non-nil, zero value otherwise.

### GetPolicyDescriptionOk

`func (o *WebhookV1MPARequestNotificationArguments) GetPolicyDescriptionOk() (*string, bool)`

GetPolicyDescriptionOk returns a tuple with the PolicyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDescription

`func (o *WebhookV1MPARequestNotificationArguments) SetPolicyDescription(v string)`

SetPolicyDescription sets PolicyDescription field to given value.

### HasPolicyDescription

`func (o *WebhookV1MPARequestNotificationArguments) HasPolicyDescription() bool`

HasPolicyDescription returns a boolean if a field has been set.

### GetApprovalExpiresAt

`func (o *WebhookV1MPARequestNotificationArguments) GetApprovalExpiresAt() WebhookV1ProtobufTimestamp`

GetApprovalExpiresAt returns the ApprovalExpiresAt field if non-nil, zero value otherwise.

### GetApprovalExpiresAtOk

`func (o *WebhookV1MPARequestNotificationArguments) GetApprovalExpiresAtOk() (*WebhookV1ProtobufTimestamp, bool)`

GetApprovalExpiresAtOk returns a tuple with the ApprovalExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalExpiresAt

`func (o *WebhookV1MPARequestNotificationArguments) SetApprovalExpiresAt(v WebhookV1ProtobufTimestamp)`

SetApprovalExpiresAt sets ApprovalExpiresAt field to given value.

### HasApprovalExpiresAt

`func (o *WebhookV1MPARequestNotificationArguments) HasApprovalExpiresAt() bool`

HasApprovalExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


