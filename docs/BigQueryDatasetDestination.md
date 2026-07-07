# BigQueryDatasetDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | **string** | Destination dataset ID. If the dataset doesn&#39;t exist, it&#39;s created.  | 
**Location** | Pointer to **string** | BigQuery location for the dataset. Defaults to the original dataset location.  For multi-region, set to &#x60;US&#x60; or &#x60;EU&#x60;.  | [optional] 
**SelectedTables** | Pointer to **[]string** | List of table names to restore. If empty or not specified, all tables in the snapshot are restored.  | [optional] 
**SelectedViews** | Pointer to **[]string** | List of view names to restore. If empty or not specified, all views in the snapshot are restored.  | [optional] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the restored BigQuery dataset as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon-restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 
**OverwriteExistingData** | Pointer to **bool** | When set to &#x60;true&#x60;, existing tables and metadata (views, routines) in the target dataset will be overwritten by the restored data. Materialized views are never overwritten: an existing materialized view is left as-is, since BigQuery has no &#x60;CREATE OR REPLACE MATERIALIZED VIEW&#x60;. When &#x60;false&#x60; (default), the restore fails if any conflicting objects already exist in the target dataset.  | [optional] [default to false]

## Methods

### NewBigQueryDatasetDestination

`func NewBigQueryDatasetDestination(datasetId string, ) *BigQueryDatasetDestination`

NewBigQueryDatasetDestination instantiates a new BigQueryDatasetDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryDatasetDestinationWithDefaults

`func NewBigQueryDatasetDestinationWithDefaults() *BigQueryDatasetDestination`

NewBigQueryDatasetDestinationWithDefaults instantiates a new BigQueryDatasetDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *BigQueryDatasetDestination) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *BigQueryDatasetDestination) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *BigQueryDatasetDestination) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetLocation

`func (o *BigQueryDatasetDestination) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BigQueryDatasetDestination) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BigQueryDatasetDestination) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BigQueryDatasetDestination) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSelectedTables

`func (o *BigQueryDatasetDestination) GetSelectedTables() []string`

GetSelectedTables returns the SelectedTables field if non-nil, zero value otherwise.

### GetSelectedTablesOk

`func (o *BigQueryDatasetDestination) GetSelectedTablesOk() (*[]string, bool)`

GetSelectedTablesOk returns a tuple with the SelectedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTables

`func (o *BigQueryDatasetDestination) SetSelectedTables(v []string)`

SetSelectedTables sets SelectedTables field to given value.

### HasSelectedTables

`func (o *BigQueryDatasetDestination) HasSelectedTables() bool`

HasSelectedTables returns a boolean if a field has been set.

### GetSelectedViews

`func (o *BigQueryDatasetDestination) GetSelectedViews() []string`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *BigQueryDatasetDestination) GetSelectedViewsOk() (*[]string, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *BigQueryDatasetDestination) SetSelectedViews(v []string)`

SetSelectedViews sets SelectedViews field to given value.

### HasSelectedViews

`func (o *BigQueryDatasetDestination) HasSelectedViews() bool`

HasSelectedViews returns a boolean if a field has been set.

### GetLabels

`func (o *BigQueryDatasetDestination) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BigQueryDatasetDestination) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BigQueryDatasetDestination) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BigQueryDatasetDestination) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOverwriteExistingData

`func (o *BigQueryDatasetDestination) GetOverwriteExistingData() bool`

GetOverwriteExistingData returns the OverwriteExistingData field if non-nil, zero value otherwise.

### GetOverwriteExistingDataOk

`func (o *BigQueryDatasetDestination) GetOverwriteExistingDataOk() (*bool, bool)`

GetOverwriteExistingDataOk returns a tuple with the OverwriteExistingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingData

`func (o *BigQueryDatasetDestination) SetOverwriteExistingData(v bool)`

SetOverwriteExistingData sets OverwriteExistingData field to given value.

### HasOverwriteExistingData

`func (o *BigQueryDatasetDestination) HasOverwriteExistingData() bool`

HasOverwriteExistingData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


