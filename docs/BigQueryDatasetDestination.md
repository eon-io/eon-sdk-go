# BigQueryDatasetDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | **string** | Target dataset ID for the restore. If the dataset doesn&#39;t exist, it will be created.  | 
**Location** | Pointer to **string** | BigQuery location for the dataset (e.g., \&quot;US\&quot;, \&quot;EU\&quot;, \&quot;us-central1\&quot;). Defaults to the original dataset location if not specified.  | [optional] 
**SelectedTables** | Pointer to **[]string** | List of table names to restore. If empty or not specified, all tables in the snapshot will be restored.  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


