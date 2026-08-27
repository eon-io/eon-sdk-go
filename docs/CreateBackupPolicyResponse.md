# CreateBackupPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicy** | Pointer to [**BackupPolicy**](BackupPolicy.md) |  | [optional] 
**ActionApprovalRequest** | Pointer to [**NullableMPARequest**](MPARequest.md) |  | [optional] 

## Methods

### NewCreateBackupPolicyResponse

`func NewCreateBackupPolicyResponse() *CreateBackupPolicyResponse`

NewCreateBackupPolicyResponse instantiates a new CreateBackupPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBackupPolicyResponseWithDefaults

`func NewCreateBackupPolicyResponseWithDefaults() *CreateBackupPolicyResponse`

NewCreateBackupPolicyResponseWithDefaults instantiates a new CreateBackupPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *CreateBackupPolicyResponse) GetBackupPolicy() BackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *CreateBackupPolicyResponse) GetBackupPolicyOk() (*BackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *CreateBackupPolicyResponse) SetBackupPolicy(v BackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *CreateBackupPolicyResponse) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetActionApprovalRequest

`func (o *CreateBackupPolicyResponse) GetActionApprovalRequest() MPARequest`

GetActionApprovalRequest returns the ActionApprovalRequest field if non-nil, zero value otherwise.

### GetActionApprovalRequestOk

`func (o *CreateBackupPolicyResponse) GetActionApprovalRequestOk() (*MPARequest, bool)`

GetActionApprovalRequestOk returns a tuple with the ActionApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApprovalRequest

`func (o *CreateBackupPolicyResponse) SetActionApprovalRequest(v MPARequest)`

SetActionApprovalRequest sets ActionApprovalRequest field to given value.

### HasActionApprovalRequest

`func (o *CreateBackupPolicyResponse) HasActionApprovalRequest() bool`

HasActionApprovalRequest returns a boolean if a field has been set.

### SetActionApprovalRequestNil

`func (o *CreateBackupPolicyResponse) SetActionApprovalRequestNil(b bool)`

 SetActionApprovalRequestNil sets the value for ActionApprovalRequest to be an explicit nil

### UnsetActionApprovalRequest
`func (o *CreateBackupPolicyResponse) UnsetActionApprovalRequest()`

UnsetActionApprovalRequest ensures that no value is present for ActionApprovalRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


