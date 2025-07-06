# ResourceProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsEc2** | Pointer to [**NullableAwsEc2Resource**](AwsEc2Resource.md) |  | [optional] 

## Methods

### NewResourceProperties

`func NewResourceProperties() *ResourceProperties`

NewResourceProperties instantiates a new ResourceProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePropertiesWithDefaults

`func NewResourcePropertiesWithDefaults() *ResourceProperties`

NewResourcePropertiesWithDefaults instantiates a new ResourceProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsEc2

`func (o *ResourceProperties) GetAwsEc2() AwsEc2Resource`

GetAwsEc2 returns the AwsEc2 field if non-nil, zero value otherwise.

### GetAwsEc2Ok

`func (o *ResourceProperties) GetAwsEc2Ok() (*AwsEc2Resource, bool)`

GetAwsEc2Ok returns a tuple with the AwsEc2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEc2

`func (o *ResourceProperties) SetAwsEc2(v AwsEc2Resource)`

SetAwsEc2 sets AwsEc2 field to given value.

### HasAwsEc2

`func (o *ResourceProperties) HasAwsEc2() bool`

HasAwsEc2 returns a boolean if a field has been set.

### SetAwsEc2Nil

`func (o *ResourceProperties) SetAwsEc2Nil(b bool)`

 SetAwsEc2Nil sets the value for AwsEc2 to be an explicit nil

### UnsetAwsEc2
`func (o *ResourceProperties) UnsetAwsEc2()`

UnsetAwsEc2 ensures that no value is present for AwsEc2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


