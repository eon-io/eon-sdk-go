# AwsEc2Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]InventoryVolume**](InventoryVolume.md) | List of volumes attached to the resource. | [optional] 

## Methods

### NewAwsEc2Resource

`func NewAwsEc2Resource() *AwsEc2Resource`

NewAwsEc2Resource instantiates a new AwsEc2Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEc2ResourceWithDefaults

`func NewAwsEc2ResourceWithDefaults() *AwsEc2Resource`

NewAwsEc2ResourceWithDefaults instantiates a new AwsEc2Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *AwsEc2Resource) GetVolumes() []InventoryVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AwsEc2Resource) GetVolumesOk() (*[]InventoryVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AwsEc2Resource) SetVolumes(v []InventoryVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AwsEc2Resource) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


