# MPARequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the action approval request. | 
**Operation** | [**MPAOperationType**](MPAOperationType.md) |  | 
**RequesterEmail** | **string** | Email address of the requester who created the approval request. | 
**RequesterName** | Pointer to **string** | Display name of the requester at the time the request was created. | [optional] 
**Comment** | Pointer to **string** | Optional comment from the requester explaining why this operation is needed. | [optional] 
**Status** | [**MPARequestStatus**](MPARequestStatus.md) |  | 
**PolicyDetails** | [**[]MPARequestPolicyDetail**](MPARequestPolicyDetail.md) | Per-policy configuration and approval progress. | 
**ApprovalExpirationTime** | Pointer to **NullableTime** | Deadline for gathering approvals. Set when the request is confirmed (submitted). Null for requests in CREATED or DISCARDED status. | [optional] 
**ExecutionExpirationTime** | Pointer to **NullableTime** | Deadline for executing the approved action. Set when the request is approved. | [optional] 
**ResolvedTime** | Pointer to **NullableTime** | Time when the request reached a terminal state, such as executed, expired, canceled, denied, or discarded. Null while the request is still active. | [optional] 
**CreatedTime** | **time.Time** | Time when the approval request was created. | 

## Methods

### NewMPARequest

`func NewMPARequest(id string, operation MPAOperationType, requesterEmail string, status MPARequestStatus, policyDetails []MPARequestPolicyDetail, createdTime time.Time, ) *MPARequest`

NewMPARequest instantiates a new MPARequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPARequestWithDefaults

`func NewMPARequestWithDefaults() *MPARequest`

NewMPARequestWithDefaults instantiates a new MPARequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MPARequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MPARequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MPARequest) SetId(v string)`

SetId sets Id field to given value.


### GetOperation

`func (o *MPARequest) GetOperation() MPAOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MPARequest) GetOperationOk() (*MPAOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MPARequest) SetOperation(v MPAOperationType)`

SetOperation sets Operation field to given value.


### GetRequesterEmail

`func (o *MPARequest) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *MPARequest) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *MPARequest) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.


### GetRequesterName

`func (o *MPARequest) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *MPARequest) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *MPARequest) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *MPARequest) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetComment

`func (o *MPARequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MPARequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MPARequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MPARequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStatus

`func (o *MPARequest) GetStatus() MPARequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MPARequest) GetStatusOk() (*MPARequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MPARequest) SetStatus(v MPARequestStatus)`

SetStatus sets Status field to given value.


### GetPolicyDetails

`func (o *MPARequest) GetPolicyDetails() []MPARequestPolicyDetail`

GetPolicyDetails returns the PolicyDetails field if non-nil, zero value otherwise.

### GetPolicyDetailsOk

`func (o *MPARequest) GetPolicyDetailsOk() (*[]MPARequestPolicyDetail, bool)`

GetPolicyDetailsOk returns a tuple with the PolicyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDetails

`func (o *MPARequest) SetPolicyDetails(v []MPARequestPolicyDetail)`

SetPolicyDetails sets PolicyDetails field to given value.


### GetApprovalExpirationTime

`func (o *MPARequest) GetApprovalExpirationTime() time.Time`

GetApprovalExpirationTime returns the ApprovalExpirationTime field if non-nil, zero value otherwise.

### GetApprovalExpirationTimeOk

`func (o *MPARequest) GetApprovalExpirationTimeOk() (*time.Time, bool)`

GetApprovalExpirationTimeOk returns a tuple with the ApprovalExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalExpirationTime

`func (o *MPARequest) SetApprovalExpirationTime(v time.Time)`

SetApprovalExpirationTime sets ApprovalExpirationTime field to given value.

### HasApprovalExpirationTime

`func (o *MPARequest) HasApprovalExpirationTime() bool`

HasApprovalExpirationTime returns a boolean if a field has been set.

### SetApprovalExpirationTimeNil

`func (o *MPARequest) SetApprovalExpirationTimeNil(b bool)`

 SetApprovalExpirationTimeNil sets the value for ApprovalExpirationTime to be an explicit nil

### UnsetApprovalExpirationTime
`func (o *MPARequest) UnsetApprovalExpirationTime()`

UnsetApprovalExpirationTime ensures that no value is present for ApprovalExpirationTime, not even an explicit nil
### GetExecutionExpirationTime

`func (o *MPARequest) GetExecutionExpirationTime() time.Time`

GetExecutionExpirationTime returns the ExecutionExpirationTime field if non-nil, zero value otherwise.

### GetExecutionExpirationTimeOk

`func (o *MPARequest) GetExecutionExpirationTimeOk() (*time.Time, bool)`

GetExecutionExpirationTimeOk returns a tuple with the ExecutionExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionExpirationTime

`func (o *MPARequest) SetExecutionExpirationTime(v time.Time)`

SetExecutionExpirationTime sets ExecutionExpirationTime field to given value.

### HasExecutionExpirationTime

`func (o *MPARequest) HasExecutionExpirationTime() bool`

HasExecutionExpirationTime returns a boolean if a field has been set.

### SetExecutionExpirationTimeNil

`func (o *MPARequest) SetExecutionExpirationTimeNil(b bool)`

 SetExecutionExpirationTimeNil sets the value for ExecutionExpirationTime to be an explicit nil

### UnsetExecutionExpirationTime
`func (o *MPARequest) UnsetExecutionExpirationTime()`

UnsetExecutionExpirationTime ensures that no value is present for ExecutionExpirationTime, not even an explicit nil
### GetResolvedTime

`func (o *MPARequest) GetResolvedTime() time.Time`

GetResolvedTime returns the ResolvedTime field if non-nil, zero value otherwise.

### GetResolvedTimeOk

`func (o *MPARequest) GetResolvedTimeOk() (*time.Time, bool)`

GetResolvedTimeOk returns a tuple with the ResolvedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTime

`func (o *MPARequest) SetResolvedTime(v time.Time)`

SetResolvedTime sets ResolvedTime field to given value.

### HasResolvedTime

`func (o *MPARequest) HasResolvedTime() bool`

HasResolvedTime returns a boolean if a field has been set.

### SetResolvedTimeNil

`func (o *MPARequest) SetResolvedTimeNil(b bool)`

 SetResolvedTimeNil sets the value for ResolvedTime to be an explicit nil

### UnsetResolvedTime
`func (o *MPARequest) UnsetResolvedTime()`

UnsetResolvedTime ensures that no value is present for ResolvedTime, not even an explicit nil
### GetCreatedTime

`func (o *MPARequest) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MPARequest) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MPARequest) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


