# AwsAccountMetricsDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The CloudWatch region to use for publishing metrics. | [optional] 

## Methods

### NewAwsAccountMetricsDestination

`func NewAwsAccountMetricsDestination() *AwsAccountMetricsDestination`

NewAwsAccountMetricsDestination instantiates a new AwsAccountMetricsDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAccountMetricsDestinationWithDefaults

`func NewAwsAccountMetricsDestinationWithDefaults() *AwsAccountMetricsDestination`

NewAwsAccountMetricsDestinationWithDefaults instantiates a new AwsAccountMetricsDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsAccountMetricsDestination) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsAccountMetricsDestination) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsAccountMetricsDestination) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsAccountMetricsDestination) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


