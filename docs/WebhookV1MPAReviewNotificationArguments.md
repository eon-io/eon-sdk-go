# WebhookV1MPAReviewNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | Approval request ID. | [optional] 
**Operation** | Pointer to **string** | Operation that requires approval. | [optional] 
**RequesterEmail** | Pointer to **string** | Email address of the requester. | [optional] 
**ReviewerEmail** | Pointer to **string** | Email address of the reviewer. | [optional] 
**Decision** | Pointer to **string** | Review decision. | [optional] 
**Comment** | Pointer to **string** | Reviewer comment. | [optional] 
**NewStatus** | Pointer to **string** | Status of the request after the review. | [optional] 

## Methods

### NewWebhookV1MPAReviewNotificationArguments

`func NewWebhookV1MPAReviewNotificationArguments() *WebhookV1MPAReviewNotificationArguments`

NewWebhookV1MPAReviewNotificationArguments instantiates a new WebhookV1MPAReviewNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1MPAReviewNotificationArgumentsWithDefaults

`func NewWebhookV1MPAReviewNotificationArgumentsWithDefaults() *WebhookV1MPAReviewNotificationArguments`

NewWebhookV1MPAReviewNotificationArgumentsWithDefaults instantiates a new WebhookV1MPAReviewNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *WebhookV1MPAReviewNotificationArguments) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookV1MPAReviewNotificationArguments) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WebhookV1MPAReviewNotificationArguments) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetOperation

`func (o *WebhookV1MPAReviewNotificationArguments) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *WebhookV1MPAReviewNotificationArguments) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *WebhookV1MPAReviewNotificationArguments) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRequesterEmail

`func (o *WebhookV1MPAReviewNotificationArguments) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *WebhookV1MPAReviewNotificationArguments) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.

### HasRequesterEmail

`func (o *WebhookV1MPAReviewNotificationArguments) HasRequesterEmail() bool`

HasRequesterEmail returns a boolean if a field has been set.

### GetReviewerEmail

`func (o *WebhookV1MPAReviewNotificationArguments) GetReviewerEmail() string`

GetReviewerEmail returns the ReviewerEmail field if non-nil, zero value otherwise.

### GetReviewerEmailOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetReviewerEmailOk() (*string, bool)`

GetReviewerEmailOk returns a tuple with the ReviewerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerEmail

`func (o *WebhookV1MPAReviewNotificationArguments) SetReviewerEmail(v string)`

SetReviewerEmail sets ReviewerEmail field to given value.

### HasReviewerEmail

`func (o *WebhookV1MPAReviewNotificationArguments) HasReviewerEmail() bool`

HasReviewerEmail returns a boolean if a field has been set.

### GetDecision

`func (o *WebhookV1MPAReviewNotificationArguments) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *WebhookV1MPAReviewNotificationArguments) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *WebhookV1MPAReviewNotificationArguments) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetComment

`func (o *WebhookV1MPAReviewNotificationArguments) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WebhookV1MPAReviewNotificationArguments) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WebhookV1MPAReviewNotificationArguments) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetNewStatus

`func (o *WebhookV1MPAReviewNotificationArguments) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *WebhookV1MPAReviewNotificationArguments) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *WebhookV1MPAReviewNotificationArguments) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *WebhookV1MPAReviewNotificationArguments) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


