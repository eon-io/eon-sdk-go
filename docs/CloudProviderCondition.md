# CloudProviderCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**ScalarOperators**](ScalarOperators.md) |  | 
**CloudProviders** | [**[]Provider**](Provider.md) |  | 

## Methods

### NewCloudProviderCondition

`func NewCloudProviderCondition(operator ScalarOperators, cloudProviders []Provider, ) *CloudProviderCondition`

NewCloudProviderCondition instantiates a new CloudProviderCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderConditionWithDefaults

`func NewCloudProviderConditionWithDefaults() *CloudProviderCondition`

NewCloudProviderConditionWithDefaults instantiates a new CloudProviderCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CloudProviderCondition) GetOperator() ScalarOperators`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudProviderCondition) GetOperatorOk() (*ScalarOperators, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudProviderCondition) SetOperator(v ScalarOperators)`

SetOperator sets Operator field to given value.


### GetCloudProviders

`func (o *CloudProviderCondition) GetCloudProviders() []Provider`

GetCloudProviders returns the CloudProviders field if non-nil, zero value otherwise.

### GetCloudProvidersOk

`func (o *CloudProviderCondition) GetCloudProvidersOk() (*[]Provider, bool)`

GetCloudProvidersOk returns a tuple with the CloudProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviders

`func (o *CloudProviderCondition) SetCloudProviders(v []Provider)`

SetCloudProviders sets CloudProviders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


