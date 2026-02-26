# RestoreAccountCloudAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Aws** | Pointer to [**NullableAwsRestoreAccountAttributes**](AwsRestoreAccountAttributes.md) |  | [optional] 
**Gcp** | Pointer to [**NullableGcpRestoreAccountAttributes**](GcpRestoreAccountAttributes.md) |  | [optional] 
**Azure** | Pointer to [**NullableAzureRestoreAccountAttributes**](AzureRestoreAccountAttributes.md) |  | [optional] 

## Methods

### NewRestoreAccountCloudAttributes

`func NewRestoreAccountCloudAttributes(cloudProvider Provider, ) *RestoreAccountCloudAttributes`

NewRestoreAccountCloudAttributes instantiates a new RestoreAccountCloudAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAccountCloudAttributesWithDefaults

`func NewRestoreAccountCloudAttributesWithDefaults() *RestoreAccountCloudAttributes`

NewRestoreAccountCloudAttributesWithDefaults instantiates a new RestoreAccountCloudAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *RestoreAccountCloudAttributes) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *RestoreAccountCloudAttributes) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *RestoreAccountCloudAttributes) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetAws

`func (o *RestoreAccountCloudAttributes) GetAws() AwsRestoreAccountAttributes`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *RestoreAccountCloudAttributes) GetAwsOk() (*AwsRestoreAccountAttributes, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *RestoreAccountCloudAttributes) SetAws(v AwsRestoreAccountAttributes)`

SetAws sets Aws field to given value.

### HasAws

`func (o *RestoreAccountCloudAttributes) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *RestoreAccountCloudAttributes) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *RestoreAccountCloudAttributes) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *RestoreAccountCloudAttributes) GetGcp() GcpRestoreAccountAttributes`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *RestoreAccountCloudAttributes) GetGcpOk() (*GcpRestoreAccountAttributes, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *RestoreAccountCloudAttributes) SetGcp(v GcpRestoreAccountAttributes)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *RestoreAccountCloudAttributes) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *RestoreAccountCloudAttributes) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *RestoreAccountCloudAttributes) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil
### GetAzure

`func (o *RestoreAccountCloudAttributes) GetAzure() AzureRestoreAccountAttributes`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *RestoreAccountCloudAttributes) GetAzureOk() (*AzureRestoreAccountAttributes, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *RestoreAccountCloudAttributes) SetAzure(v AzureRestoreAccountAttributes)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *RestoreAccountCloudAttributes) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *RestoreAccountCloudAttributes) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *RestoreAccountCloudAttributes) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


