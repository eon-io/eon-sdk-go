# Ec2InstanceRestoreDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsEc2** | Pointer to [**Ec2InstanceRestoreTarget**](Ec2InstanceRestoreTarget.md) |  | [optional] 

## Methods

### NewEc2InstanceRestoreDestination

`func NewEc2InstanceRestoreDestination() *Ec2InstanceRestoreDestination`

NewEc2InstanceRestoreDestination instantiates a new Ec2InstanceRestoreDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEc2InstanceRestoreDestinationWithDefaults

`func NewEc2InstanceRestoreDestinationWithDefaults() *Ec2InstanceRestoreDestination`

NewEc2InstanceRestoreDestinationWithDefaults instantiates a new Ec2InstanceRestoreDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsEc2

`func (o *Ec2InstanceRestoreDestination) GetAwsEc2() Ec2InstanceRestoreTarget`

GetAwsEc2 returns the AwsEc2 field if non-nil, zero value otherwise.

### GetAwsEc2Ok

`func (o *Ec2InstanceRestoreDestination) GetAwsEc2Ok() (*Ec2InstanceRestoreTarget, bool)`

GetAwsEc2Ok returns a tuple with the AwsEc2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEc2

`func (o *Ec2InstanceRestoreDestination) SetAwsEc2(v Ec2InstanceRestoreTarget)`

SetAwsEc2 sets AwsEc2 field to given value.

### HasAwsEc2

`func (o *Ec2InstanceRestoreDestination) HasAwsEc2() bool`

HasAwsEc2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


