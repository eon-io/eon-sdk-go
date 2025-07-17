# RestoreAccountMetricsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned restore account ID. | 
**Enabled** | **bool** | Indicates whether job metrics are enabled for the restore account. If &#x60;true&#x60;, job metrics are collected and can be queried.  | 
**Destination** | [**AccountMetricsDestination**](AccountMetricsDestination.md) |  | 

## Methods

### NewRestoreAccountMetricsConfig

`func NewRestoreAccountMetricsConfig(restoreAccountId string, enabled bool, destination AccountMetricsDestination, ) *RestoreAccountMetricsConfig`

NewRestoreAccountMetricsConfig instantiates a new RestoreAccountMetricsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAccountMetricsConfigWithDefaults

`func NewRestoreAccountMetricsConfigWithDefaults() *RestoreAccountMetricsConfig`

NewRestoreAccountMetricsConfigWithDefaults instantiates a new RestoreAccountMetricsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *RestoreAccountMetricsConfig) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreAccountMetricsConfig) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreAccountMetricsConfig) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetEnabled

`func (o *RestoreAccountMetricsConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RestoreAccountMetricsConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RestoreAccountMetricsConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDestination

`func (o *RestoreAccountMetricsConfig) GetDestination() AccountMetricsDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreAccountMetricsConfig) GetDestinationOk() (*AccountMetricsDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreAccountMetricsConfig) SetDestination(v AccountMetricsDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


