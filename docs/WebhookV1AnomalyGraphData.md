# WebhookV1AnomalyGraphData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPoints** | Pointer to [**[]WebhookV1AnomalyDataPoint**](WebhookV1AnomalyDataPoint.md) | Points of the time series. | [optional] 

## Methods

### NewWebhookV1AnomalyGraphData

`func NewWebhookV1AnomalyGraphData() *WebhookV1AnomalyGraphData`

NewWebhookV1AnomalyGraphData instantiates a new WebhookV1AnomalyGraphData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1AnomalyGraphDataWithDefaults

`func NewWebhookV1AnomalyGraphDataWithDefaults() *WebhookV1AnomalyGraphData`

NewWebhookV1AnomalyGraphDataWithDefaults instantiates a new WebhookV1AnomalyGraphData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPoints

`func (o *WebhookV1AnomalyGraphData) GetDataPoints() []WebhookV1AnomalyDataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *WebhookV1AnomalyGraphData) GetDataPointsOk() (*[]WebhookV1AnomalyDataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *WebhookV1AnomalyGraphData) SetDataPoints(v []WebhookV1AnomalyDataPoint)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *WebhookV1AnomalyGraphData) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


