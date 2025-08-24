# UpdateRestoreAccountConnectivityConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**NullableAwsRestoreAccountConnectivityConfig**](AwsRestoreAccountConnectivityConfig.md) |  | [optional] 
**Gcp** | Pointer to [**NullableGcpRestoreAccountConnectivityConfig**](GcpRestoreAccountConnectivityConfig.md) |  | [optional] 

## Methods

### NewUpdateRestoreAccountConnectivityConfigRequest

`func NewUpdateRestoreAccountConnectivityConfigRequest() *UpdateRestoreAccountConnectivityConfigRequest`

NewUpdateRestoreAccountConnectivityConfigRequest instantiates a new UpdateRestoreAccountConnectivityConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRestoreAccountConnectivityConfigRequestWithDefaults

`func NewUpdateRestoreAccountConnectivityConfigRequestWithDefaults() *UpdateRestoreAccountConnectivityConfigRequest`

NewUpdateRestoreAccountConnectivityConfigRequestWithDefaults instantiates a new UpdateRestoreAccountConnectivityConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *UpdateRestoreAccountConnectivityConfigRequest) GetAws() AwsRestoreAccountConnectivityConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *UpdateRestoreAccountConnectivityConfigRequest) GetAwsOk() (*AwsRestoreAccountConnectivityConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *UpdateRestoreAccountConnectivityConfigRequest) SetAws(v AwsRestoreAccountConnectivityConfig)`

SetAws sets Aws field to given value.

### HasAws

`func (o *UpdateRestoreAccountConnectivityConfigRequest) HasAws() bool`

HasAws returns a boolean if a field has been set.

### SetAwsNil

`func (o *UpdateRestoreAccountConnectivityConfigRequest) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *UpdateRestoreAccountConnectivityConfigRequest) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil
### GetGcp

`func (o *UpdateRestoreAccountConnectivityConfigRequest) GetGcp() GcpRestoreAccountConnectivityConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *UpdateRestoreAccountConnectivityConfigRequest) GetGcpOk() (*GcpRestoreAccountConnectivityConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *UpdateRestoreAccountConnectivityConfigRequest) SetGcp(v GcpRestoreAccountConnectivityConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *UpdateRestoreAccountConnectivityConfigRequest) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### SetGcpNil

`func (o *UpdateRestoreAccountConnectivityConfigRequest) SetGcpNil(b bool)`

 SetGcpNil sets the value for Gcp to be an explicit nil

### UnsetGcp
`func (o *UpdateRestoreAccountConnectivityConfigRequest) UnsetGcp()`

UnsetGcp ensures that no value is present for Gcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


