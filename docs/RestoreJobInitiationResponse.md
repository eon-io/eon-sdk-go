# RestoreJobInitiationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | Restore job ID. | [optional] 
**ActionApprovalRequest** | Pointer to [**NullableMPARequest**](MPARequest.md) |  | [optional] 

## Methods

### NewRestoreJobInitiationResponse

`func NewRestoreJobInitiationResponse() *RestoreJobInitiationResponse`

NewRestoreJobInitiationResponse instantiates a new RestoreJobInitiationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobInitiationResponseWithDefaults

`func NewRestoreJobInitiationResponseWithDefaults() *RestoreJobInitiationResponse`

NewRestoreJobInitiationResponseWithDefaults instantiates a new RestoreJobInitiationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *RestoreJobInitiationResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RestoreJobInitiationResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RestoreJobInitiationResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RestoreJobInitiationResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetActionApprovalRequest

`func (o *RestoreJobInitiationResponse) GetActionApprovalRequest() MPARequest`

GetActionApprovalRequest returns the ActionApprovalRequest field if non-nil, zero value otherwise.

### GetActionApprovalRequestOk

`func (o *RestoreJobInitiationResponse) GetActionApprovalRequestOk() (*MPARequest, bool)`

GetActionApprovalRequestOk returns a tuple with the ActionApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApprovalRequest

`func (o *RestoreJobInitiationResponse) SetActionApprovalRequest(v MPARequest)`

SetActionApprovalRequest sets ActionApprovalRequest field to given value.

### HasActionApprovalRequest

`func (o *RestoreJobInitiationResponse) HasActionApprovalRequest() bool`

HasActionApprovalRequest returns a boolean if a field has been set.

### SetActionApprovalRequestNil

`func (o *RestoreJobInitiationResponse) SetActionApprovalRequestNil(b bool)`

 SetActionApprovalRequestNil sets the value for ActionApprovalRequest to be an explicit nil

### UnsetActionApprovalRequest
`func (o *RestoreJobInitiationResponse) UnsetActionApprovalRequest()`

UnsetActionApprovalRequest ensures that no value is present for ActionApprovalRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


