# WebhookV1AnomalyDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 
**ActualValue** | Pointer to **float64** | Observed value. | [optional] 
**ForecastValue** | Pointer to **float64** | Forecast value. | [optional] 
**UpperBound** | Pointer to **float64** | Upper bound of the expected range. | [optional] 
**LowerBound** | Pointer to **float64** | Lower bound of the expected range. | [optional] 
**IsAnomaly** | Pointer to **bool** | Whether the point is flagged as an anomaly. | [optional] 
**SnapshotId** | Pointer to **string** | Eon-assigned ID of the snapshot associated with the point. | [optional] 
**PointInTime** | Pointer to [**WebhookV1ProtobufTimestamp**](WebhookV1ProtobufTimestamp.md) |  | [optional] 

## Methods

### NewWebhookV1AnomalyDataPoint

`func NewWebhookV1AnomalyDataPoint() *WebhookV1AnomalyDataPoint`

NewWebhookV1AnomalyDataPoint instantiates a new WebhookV1AnomalyDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1AnomalyDataPointWithDefaults

`func NewWebhookV1AnomalyDataPointWithDefaults() *WebhookV1AnomalyDataPoint`

NewWebhookV1AnomalyDataPointWithDefaults instantiates a new WebhookV1AnomalyDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WebhookV1AnomalyDataPoint) GetTimestamp() WebhookV1ProtobufTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookV1AnomalyDataPoint) GetTimestampOk() (*WebhookV1ProtobufTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookV1AnomalyDataPoint) SetTimestamp(v WebhookV1ProtobufTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookV1AnomalyDataPoint) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActualValue

`func (o *WebhookV1AnomalyDataPoint) GetActualValue() float64`

GetActualValue returns the ActualValue field if non-nil, zero value otherwise.

### GetActualValueOk

`func (o *WebhookV1AnomalyDataPoint) GetActualValueOk() (*float64, bool)`

GetActualValueOk returns a tuple with the ActualValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualValue

`func (o *WebhookV1AnomalyDataPoint) SetActualValue(v float64)`

SetActualValue sets ActualValue field to given value.

### HasActualValue

`func (o *WebhookV1AnomalyDataPoint) HasActualValue() bool`

HasActualValue returns a boolean if a field has been set.

### GetForecastValue

`func (o *WebhookV1AnomalyDataPoint) GetForecastValue() float64`

GetForecastValue returns the ForecastValue field if non-nil, zero value otherwise.

### GetForecastValueOk

`func (o *WebhookV1AnomalyDataPoint) GetForecastValueOk() (*float64, bool)`

GetForecastValueOk returns a tuple with the ForecastValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastValue

`func (o *WebhookV1AnomalyDataPoint) SetForecastValue(v float64)`

SetForecastValue sets ForecastValue field to given value.

### HasForecastValue

`func (o *WebhookV1AnomalyDataPoint) HasForecastValue() bool`

HasForecastValue returns a boolean if a field has been set.

### GetUpperBound

`func (o *WebhookV1AnomalyDataPoint) GetUpperBound() float64`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *WebhookV1AnomalyDataPoint) GetUpperBoundOk() (*float64, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *WebhookV1AnomalyDataPoint) SetUpperBound(v float64)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *WebhookV1AnomalyDataPoint) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetLowerBound

`func (o *WebhookV1AnomalyDataPoint) GetLowerBound() float64`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *WebhookV1AnomalyDataPoint) GetLowerBoundOk() (*float64, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *WebhookV1AnomalyDataPoint) SetLowerBound(v float64)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *WebhookV1AnomalyDataPoint) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetIsAnomaly

`func (o *WebhookV1AnomalyDataPoint) GetIsAnomaly() bool`

GetIsAnomaly returns the IsAnomaly field if non-nil, zero value otherwise.

### GetIsAnomalyOk

`func (o *WebhookV1AnomalyDataPoint) GetIsAnomalyOk() (*bool, bool)`

GetIsAnomalyOk returns a tuple with the IsAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnomaly

`func (o *WebhookV1AnomalyDataPoint) SetIsAnomaly(v bool)`

SetIsAnomaly sets IsAnomaly field to given value.

### HasIsAnomaly

`func (o *WebhookV1AnomalyDataPoint) HasIsAnomaly() bool`

HasIsAnomaly returns a boolean if a field has been set.

### GetSnapshotId

`func (o *WebhookV1AnomalyDataPoint) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *WebhookV1AnomalyDataPoint) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *WebhookV1AnomalyDataPoint) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *WebhookV1AnomalyDataPoint) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetPointInTime

`func (o *WebhookV1AnomalyDataPoint) GetPointInTime() WebhookV1ProtobufTimestamp`

GetPointInTime returns the PointInTime field if non-nil, zero value otherwise.

### GetPointInTimeOk

`func (o *WebhookV1AnomalyDataPoint) GetPointInTimeOk() (*WebhookV1ProtobufTimestamp, bool)`

GetPointInTimeOk returns a tuple with the PointInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTime

`func (o *WebhookV1AnomalyDataPoint) SetPointInTime(v WebhookV1ProtobufTimestamp)`

SetPointInTime sets PointInTime field to given value.

### HasPointInTime

`func (o *WebhookV1AnomalyDataPoint) HasPointInTime() bool`

HasPointInTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


