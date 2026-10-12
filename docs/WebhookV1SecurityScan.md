# WebhookV1SecurityScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | Project ID. | [optional] 
**SnapshotId** | Pointer to **string** | Eon-assigned snapshot ID. | [optional] 
**ResourceId** | Pointer to **string** | Eon-assigned resource ID. | [optional] 
**DateCreated** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**UpdatedAt** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**Conclusion** | Pointer to [**WebhookV1ConclusionTypes**](WebhookV1ConclusionTypes.md) |  | [optional] 
**Results** | Pointer to [**[]WebhookV1SecurityScanResult**](WebhookV1SecurityScanResult.md) | Individual findings of the scan. | [optional] 
**Approval** | Pointer to [**WebhookV1SecurityScanApproval**](WebhookV1SecurityScanApproval.md) |  | [optional] 
**ConfidenceScore** | Pointer to **float32** | Confidence score of the conclusion. | [optional] 
**Justification** | Pointer to **string** | Justification of the conclusion. | [optional] 
**DetectorType** | Pointer to [**WebhookV1DetectorType**](WebhookV1DetectorType.md) |  | [optional] 
**ReviewerType** | Pointer to [**WebhookV1ReviewerType**](WebhookV1ReviewerType.md) |  | [optional] 
**ReviewerIsRansomware** | Pointer to **bool** | Whether the reviewer classified the finding as ransomware. Omitted when the finding wasn&#39;t reviewed. | [optional] 
**ReviewerConfidenceScore** | Pointer to **int32** | Confidence score of the reviewer. Omitted when the finding wasn&#39;t reviewed. | [optional] 
**ReviewerReason** | Pointer to **string** | Reviewer&#39;s reasoning. | [optional] 
**ReviewerPromptVersion** | Pointer to **string** | Version of the reviewer prompt. | [optional] 
**RiskLevel** | Pointer to [**WebhookV1RiskLevel**](WebhookV1RiskLevel.md) |  | [optional] 
**RiskReason** | Pointer to **string** | Reason for the risk level. | [optional] 
**RiskModelVersion** | Pointer to **string** | Version of the risk model. | [optional] 
**TotalResultsCount** | Pointer to **int32** | Total number of findings, which can exceed the number of entries in &#x60;results&#x60;. | [optional] 

## Methods

### NewWebhookV1SecurityScan

`func NewWebhookV1SecurityScan() *WebhookV1SecurityScan`

NewWebhookV1SecurityScan instantiates a new WebhookV1SecurityScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1SecurityScanWithDefaults

`func NewWebhookV1SecurityScanWithDefaults() *WebhookV1SecurityScan`

NewWebhookV1SecurityScanWithDefaults instantiates a new WebhookV1SecurityScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *WebhookV1SecurityScan) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookV1SecurityScan) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookV1SecurityScan) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WebhookV1SecurityScan) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *WebhookV1SecurityScan) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *WebhookV1SecurityScan) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *WebhookV1SecurityScan) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *WebhookV1SecurityScan) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetResourceId

`func (o *WebhookV1SecurityScan) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *WebhookV1SecurityScan) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *WebhookV1SecurityScan) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *WebhookV1SecurityScan) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetDateCreated

`func (o *WebhookV1SecurityScan) GetDateCreated() WebhookV1ProtobufTimestamp`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WebhookV1SecurityScan) GetDateCreatedOk() (*WebhookV1ProtobufTimestamp, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WebhookV1SecurityScan) SetDateCreated(v WebhookV1ProtobufTimestamp)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WebhookV1SecurityScan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookV1SecurityScan) GetUpdatedAt() WebhookV1ProtobufTimestamp`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookV1SecurityScan) GetUpdatedAtOk() (*WebhookV1ProtobufTimestamp, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookV1SecurityScan) SetUpdatedAt(v WebhookV1ProtobufTimestamp)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookV1SecurityScan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetConclusion

`func (o *WebhookV1SecurityScan) GetConclusion() WebhookV1ConclusionTypes`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *WebhookV1SecurityScan) GetConclusionOk() (*WebhookV1ConclusionTypes, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *WebhookV1SecurityScan) SetConclusion(v WebhookV1ConclusionTypes)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *WebhookV1SecurityScan) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetResults

`func (o *WebhookV1SecurityScan) GetResults() []WebhookV1SecurityScanResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebhookV1SecurityScan) GetResultsOk() (*[]WebhookV1SecurityScanResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebhookV1SecurityScan) SetResults(v []WebhookV1SecurityScanResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *WebhookV1SecurityScan) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetApproval

`func (o *WebhookV1SecurityScan) GetApproval() WebhookV1SecurityScanApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *WebhookV1SecurityScan) GetApprovalOk() (*WebhookV1SecurityScanApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *WebhookV1SecurityScan) SetApproval(v WebhookV1SecurityScanApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *WebhookV1SecurityScan) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetConfidenceScore

`func (o *WebhookV1SecurityScan) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *WebhookV1SecurityScan) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *WebhookV1SecurityScan) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.

### HasConfidenceScore

`func (o *WebhookV1SecurityScan) HasConfidenceScore() bool`

HasConfidenceScore returns a boolean if a field has been set.

### GetJustification

`func (o *WebhookV1SecurityScan) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *WebhookV1SecurityScan) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *WebhookV1SecurityScan) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *WebhookV1SecurityScan) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetDetectorType

`func (o *WebhookV1SecurityScan) GetDetectorType() WebhookV1DetectorType`

GetDetectorType returns the DetectorType field if non-nil, zero value otherwise.

### GetDetectorTypeOk

`func (o *WebhookV1SecurityScan) GetDetectorTypeOk() (*WebhookV1DetectorType, bool)`

GetDetectorTypeOk returns a tuple with the DetectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectorType

`func (o *WebhookV1SecurityScan) SetDetectorType(v WebhookV1DetectorType)`

SetDetectorType sets DetectorType field to given value.

### HasDetectorType

`func (o *WebhookV1SecurityScan) HasDetectorType() bool`

HasDetectorType returns a boolean if a field has been set.

### GetReviewerType

`func (o *WebhookV1SecurityScan) GetReviewerType() WebhookV1ReviewerType`

GetReviewerType returns the ReviewerType field if non-nil, zero value otherwise.

### GetReviewerTypeOk

`func (o *WebhookV1SecurityScan) GetReviewerTypeOk() (*WebhookV1ReviewerType, bool)`

GetReviewerTypeOk returns a tuple with the ReviewerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerType

`func (o *WebhookV1SecurityScan) SetReviewerType(v WebhookV1ReviewerType)`

SetReviewerType sets ReviewerType field to given value.

### HasReviewerType

`func (o *WebhookV1SecurityScan) HasReviewerType() bool`

HasReviewerType returns a boolean if a field has been set.

### GetReviewerIsRansomware

`func (o *WebhookV1SecurityScan) GetReviewerIsRansomware() bool`

GetReviewerIsRansomware returns the ReviewerIsRansomware field if non-nil, zero value otherwise.

### GetReviewerIsRansomwareOk

`func (o *WebhookV1SecurityScan) GetReviewerIsRansomwareOk() (*bool, bool)`

GetReviewerIsRansomwareOk returns a tuple with the ReviewerIsRansomware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerIsRansomware

`func (o *WebhookV1SecurityScan) SetReviewerIsRansomware(v bool)`

SetReviewerIsRansomware sets ReviewerIsRansomware field to given value.

### HasReviewerIsRansomware

`func (o *WebhookV1SecurityScan) HasReviewerIsRansomware() bool`

HasReviewerIsRansomware returns a boolean if a field has been set.

### GetReviewerConfidenceScore

`func (o *WebhookV1SecurityScan) GetReviewerConfidenceScore() int32`

GetReviewerConfidenceScore returns the ReviewerConfidenceScore field if non-nil, zero value otherwise.

### GetReviewerConfidenceScoreOk

`func (o *WebhookV1SecurityScan) GetReviewerConfidenceScoreOk() (*int32, bool)`

GetReviewerConfidenceScoreOk returns a tuple with the ReviewerConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerConfidenceScore

`func (o *WebhookV1SecurityScan) SetReviewerConfidenceScore(v int32)`

SetReviewerConfidenceScore sets ReviewerConfidenceScore field to given value.

### HasReviewerConfidenceScore

`func (o *WebhookV1SecurityScan) HasReviewerConfidenceScore() bool`

HasReviewerConfidenceScore returns a boolean if a field has been set.

### GetReviewerReason

`func (o *WebhookV1SecurityScan) GetReviewerReason() string`

GetReviewerReason returns the ReviewerReason field if non-nil, zero value otherwise.

### GetReviewerReasonOk

`func (o *WebhookV1SecurityScan) GetReviewerReasonOk() (*string, bool)`

GetReviewerReasonOk returns a tuple with the ReviewerReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerReason

`func (o *WebhookV1SecurityScan) SetReviewerReason(v string)`

SetReviewerReason sets ReviewerReason field to given value.

### HasReviewerReason

`func (o *WebhookV1SecurityScan) HasReviewerReason() bool`

HasReviewerReason returns a boolean if a field has been set.

### GetReviewerPromptVersion

`func (o *WebhookV1SecurityScan) GetReviewerPromptVersion() string`

GetReviewerPromptVersion returns the ReviewerPromptVersion field if non-nil, zero value otherwise.

### GetReviewerPromptVersionOk

`func (o *WebhookV1SecurityScan) GetReviewerPromptVersionOk() (*string, bool)`

GetReviewerPromptVersionOk returns a tuple with the ReviewerPromptVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerPromptVersion

`func (o *WebhookV1SecurityScan) SetReviewerPromptVersion(v string)`

SetReviewerPromptVersion sets ReviewerPromptVersion field to given value.

### HasReviewerPromptVersion

`func (o *WebhookV1SecurityScan) HasReviewerPromptVersion() bool`

HasReviewerPromptVersion returns a boolean if a field has been set.

### GetRiskLevel

`func (o *WebhookV1SecurityScan) GetRiskLevel() WebhookV1RiskLevel`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *WebhookV1SecurityScan) GetRiskLevelOk() (*WebhookV1RiskLevel, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *WebhookV1SecurityScan) SetRiskLevel(v WebhookV1RiskLevel)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *WebhookV1SecurityScan) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetRiskReason

`func (o *WebhookV1SecurityScan) GetRiskReason() string`

GetRiskReason returns the RiskReason field if non-nil, zero value otherwise.

### GetRiskReasonOk

`func (o *WebhookV1SecurityScan) GetRiskReasonOk() (*string, bool)`

GetRiskReasonOk returns a tuple with the RiskReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskReason

`func (o *WebhookV1SecurityScan) SetRiskReason(v string)`

SetRiskReason sets RiskReason field to given value.

### HasRiskReason

`func (o *WebhookV1SecurityScan) HasRiskReason() bool`

HasRiskReason returns a boolean if a field has been set.

### GetRiskModelVersion

`func (o *WebhookV1SecurityScan) GetRiskModelVersion() string`

GetRiskModelVersion returns the RiskModelVersion field if non-nil, zero value otherwise.

### GetRiskModelVersionOk

`func (o *WebhookV1SecurityScan) GetRiskModelVersionOk() (*string, bool)`

GetRiskModelVersionOk returns a tuple with the RiskModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskModelVersion

`func (o *WebhookV1SecurityScan) SetRiskModelVersion(v string)`

SetRiskModelVersion sets RiskModelVersion field to given value.

### HasRiskModelVersion

`func (o *WebhookV1SecurityScan) HasRiskModelVersion() bool`

HasRiskModelVersion returns a boolean if a field has been set.

### GetTotalResultsCount

`func (o *WebhookV1SecurityScan) GetTotalResultsCount() int32`

GetTotalResultsCount returns the TotalResultsCount field if non-nil, zero value otherwise.

### GetTotalResultsCountOk

`func (o *WebhookV1SecurityScan) GetTotalResultsCountOk() (*int32, bool)`

GetTotalResultsCountOk returns a tuple with the TotalResultsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultsCount

`func (o *WebhookV1SecurityScan) SetTotalResultsCount(v int32)`

SetTotalResultsCount sets TotalResultsCount field to given value.

### HasTotalResultsCount

`func (o *WebhookV1SecurityScan) HasTotalResultsCount() bool`

HasTotalResultsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


