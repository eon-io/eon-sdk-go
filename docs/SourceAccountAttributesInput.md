# SourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Aws** | Pointer to [**NullableAwsSourceAccountAttributesInput**](AwsSourceAccountAttributesInput.md) |  | [optional] 
**Azure** | Pointer to [**NullableAzureSourceAccountAttributesInput**](AzureSourceAccountAttributesInput.md) |  | [optional] 

## Methods

### NewSourceAccountAttributesInput

`func NewSourceAccountAttributesInput(cloudProvider Provider, ) *SourceAccountAttributesInput`

NewSourceAccountAttributesInput instantiates a new SourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountAttributesInputWithDefaults

`func NewSourceAccountAttributesInputWithDefaults() *SourceAccountAttributesInput`

NewSourceAccountAttributesInputWithDefaults instantiates a new SourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *SourceAccountAttributesInput) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SourceAccountAttributesInput) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SourceAccountAttributesInput) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetAws

`func (o *SourceAccountAttributesInput) GetAws() AwsSourceAccountAttributesInput`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *SourceAccountAttributesInput) GetAwsOk() (*AwsSourceAccountAttributesInput, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *SourceAccountAttributesInput) SetAws(v AwsSourceAccountAttributesInput)`

SetAws sets Aws field to given value.

### HasAws

`func (o *SourceAccountAttributesInput) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *SourceAccountAttributesInput) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *SourceAccountAttributesInput) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetAzure

`func (o *SourceAccountAttributesInput) GetAzure() AzureSourceAccountAttributesInput`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *SourceAccountAttributesInput) GetAzureOk() (*AzureSourceAccountAttributesInput, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *SourceAccountAttributesInput) SetAzure(v AzureSourceAccountAttributesInput)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *SourceAccountAttributesInput) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *SourceAccountAttributesInput) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *SourceAccountAttributesInput) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


