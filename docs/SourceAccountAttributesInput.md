# SourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Aws** | Pointer to [**NullableAwsSourceAccountAttributesInput**](AwsSourceAccountAttributesInput.md) |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


