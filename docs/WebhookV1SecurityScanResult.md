# WebhookV1SecurityScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidence** | Pointer to [**WebhookV1SecurityEvidence**](WebhookV1SecurityEvidence.md) |  | [optional] 
**Justification** | Pointer to **string** | Justification of the finding. | [optional] 
**Score** | Pointer to **float32** | Score of the finding. | [optional] 
**AnomalyGraph** | Pointer to [**WebhookV1AnomalyGraphData**](WebhookV1AnomalyGraphData.md) |  | [optional] 
**Path** | Pointer to **string** | Path of the object the finding refers to. | [optional] 
**Type** | Pointer to [**WebhookV1FindingObjectType**](WebhookV1FindingObjectType.md) |  | [optional] 

## Methods

### NewWebhookV1SecurityScanResult

`func NewWebhookV1SecurityScanResult() *WebhookV1SecurityScanResult`

NewWebhookV1SecurityScanResult instantiates a new WebhookV1SecurityScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1SecurityScanResultWithDefaults

`func NewWebhookV1SecurityScanResultWithDefaults() *WebhookV1SecurityScanResult`

NewWebhookV1SecurityScanResultWithDefaults instantiates a new WebhookV1SecurityScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidence

`func (o *WebhookV1SecurityScanResult) GetEvidence() WebhookV1SecurityEvidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *WebhookV1SecurityScanResult) GetEvidenceOk() (*WebhookV1SecurityEvidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *WebhookV1SecurityScanResult) SetEvidence(v WebhookV1SecurityEvidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *WebhookV1SecurityScanResult) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetJustification

`func (o *WebhookV1SecurityScanResult) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *WebhookV1SecurityScanResult) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *WebhookV1SecurityScanResult) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *WebhookV1SecurityScanResult) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetScore

`func (o *WebhookV1SecurityScanResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *WebhookV1SecurityScanResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *WebhookV1SecurityScanResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *WebhookV1SecurityScanResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetAnomalyGraph

`func (o *WebhookV1SecurityScanResult) GetAnomalyGraph() WebhookV1AnomalyGraphData`

GetAnomalyGraph returns the AnomalyGraph field if non-nil, zero value otherwise.

### GetAnomalyGraphOk

`func (o *WebhookV1SecurityScanResult) GetAnomalyGraphOk() (*WebhookV1AnomalyGraphData, bool)`

GetAnomalyGraphOk returns a tuple with the AnomalyGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyGraph

`func (o *WebhookV1SecurityScanResult) SetAnomalyGraph(v WebhookV1AnomalyGraphData)`

SetAnomalyGraph sets AnomalyGraph field to given value.

### HasAnomalyGraph

`func (o *WebhookV1SecurityScanResult) HasAnomalyGraph() bool`

HasAnomalyGraph returns a boolean if a field has been set.

### GetPath

`func (o *WebhookV1SecurityScanResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WebhookV1SecurityScanResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WebhookV1SecurityScanResult) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WebhookV1SecurityScanResult) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *WebhookV1SecurityScanResult) GetType() WebhookV1FindingObjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookV1SecurityScanResult) GetTypeOk() (*WebhookV1FindingObjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookV1SecurityScanResult) SetType(v WebhookV1FindingObjectType)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookV1SecurityScanResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


