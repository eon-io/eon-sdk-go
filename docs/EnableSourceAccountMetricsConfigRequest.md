# EnableSourceAccountMetricsConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**NullableAwsAccountMetricsDestination**](AwsAccountMetricsDestination.md) |  | [optional] 

## Methods

### NewEnableSourceAccountMetricsConfigRequest

`func NewEnableSourceAccountMetricsConfigRequest() *EnableSourceAccountMetricsConfigRequest`

NewEnableSourceAccountMetricsConfigRequest instantiates a new EnableSourceAccountMetricsConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableSourceAccountMetricsConfigRequestWithDefaults

`func NewEnableSourceAccountMetricsConfigRequestWithDefaults() *EnableSourceAccountMetricsConfigRequest`

NewEnableSourceAccountMetricsConfigRequestWithDefaults instantiates a new EnableSourceAccountMetricsConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *EnableSourceAccountMetricsConfigRequest) GetAws() AwsAccountMetricsDestination`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *EnableSourceAccountMetricsConfigRequest) GetAwsOk() (*AwsAccountMetricsDestination, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *EnableSourceAccountMetricsConfigRequest) SetAws(v AwsAccountMetricsDestination)`

SetAws sets Aws field to given value.

### HasAws

`func (o *EnableSourceAccountMetricsConfigRequest) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *EnableSourceAccountMetricsConfigRequest) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *EnableSourceAccountMetricsConfigRequest) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


