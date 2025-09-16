# SourceAccountCloudAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Aws** | Pointer to [**NullableAwsSourceAccountAttributes**](AwsSourceAccountAttributes.md) |  | [optional] 
**Gcp** | Pointer to [**NullableGcpSourceAccountAttributes**](GcpSourceAccountAttributes.md) |  | [optional] 
**Azure** | Pointer to [**NullableAzureSourceAccountAttributes**](AzureSourceAccountAttributes.md) |  | [optional] 

## Methods

### NewSourceAccountCloudAttributes

`func NewSourceAccountCloudAttributes(cloudProvider Provider, ) *SourceAccountCloudAttributes`

NewSourceAccountCloudAttributes instantiates a new SourceAccountCloudAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAccountCloudAttributesWithDefaults

`func NewSourceAccountCloudAttributesWithDefaults() *SourceAccountCloudAttributes`

NewSourceAccountCloudAttributesWithDefaults instantiates a new SourceAccountCloudAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *SourceAccountCloudAttributes) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SourceAccountCloudAttributes) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SourceAccountCloudAttributes) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetAws

`func (o *SourceAccountCloudAttributes) GetAws() AwsSourceAccountAttributes`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *SourceAccountCloudAttributes) GetAwsOk() (*AwsSourceAccountAttributes, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *SourceAccountCloudAttributes) SetAws(v AwsSourceAccountAttributes)`

SetAws sets Aws field to given value.

### HasAws

`func (o *SourceAccountCloudAttributes) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *SourceAccountCloudAttributes) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *SourceAccountCloudAttributes) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *SourceAccountCloudAttributes) GetGcp() GcpSourceAccountAttributes`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *SourceAccountCloudAttributes) GetGcpOk() (*GcpSourceAccountAttributes, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *SourceAccountCloudAttributes) SetGcp(v GcpSourceAccountAttributes)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *SourceAccountCloudAttributes) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *SourceAccountCloudAttributes) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *SourceAccountCloudAttributes) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil
### GetAzure

`func (o *SourceAccountCloudAttributes) GetAzure() AzureSourceAccountAttributes`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *SourceAccountCloudAttributes) GetAzureOk() (*AzureSourceAccountAttributes, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *SourceAccountCloudAttributes) SetAzure(v AzureSourceAccountAttributes)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *SourceAccountCloudAttributes) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *SourceAccountCloudAttributes) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *SourceAccountCloudAttributes) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


