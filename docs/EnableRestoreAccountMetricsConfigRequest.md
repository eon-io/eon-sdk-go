# EnableRestoreAccountMetricsConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**NullableAwsAccountMetricsDestination**](AwsAccountMetricsDestination.md) |  | [optional] 

## Methods

### NewEnableRestoreAccountMetricsConfigRequest

`func NewEnableRestoreAccountMetricsConfigRequest() *EnableRestoreAccountMetricsConfigRequest`

NewEnableRestoreAccountMetricsConfigRequest instantiates a new EnableRestoreAccountMetricsConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableRestoreAccountMetricsConfigRequestWithDefaults

`func NewEnableRestoreAccountMetricsConfigRequestWithDefaults() *EnableRestoreAccountMetricsConfigRequest`

NewEnableRestoreAccountMetricsConfigRequestWithDefaults instantiates a new EnableRestoreAccountMetricsConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *EnableRestoreAccountMetricsConfigRequest) GetAws() AwsAccountMetricsDestination`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *EnableRestoreAccountMetricsConfigRequest) GetAwsOk() (*AwsAccountMetricsDestination, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *EnableRestoreAccountMetricsConfigRequest) SetAws(v AwsAccountMetricsDestination)`

SetAws sets Aws field to given value.

### HasAws

`func (o *EnableRestoreAccountMetricsConfigRequest) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *EnableRestoreAccountMetricsConfigRequest) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *EnableRestoreAccountMetricsConfigRequest) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


