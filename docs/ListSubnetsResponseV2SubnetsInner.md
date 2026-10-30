# ListSubnetsResponseV2SubnetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Cloud-provider-assigned ID of the subnet. | 
**Vpc** | **string** | Cloud-provider-assigned ID of the VPC the subnet is in. | 
**AvailabilityZone** | **string** | Availability zone the subnet is in. | 
**Name** | Pointer to **string** | Name of the subnet. | [optional] 
**CidrBlock** | Pointer to **string** | Primary IPv4 CIDR range of the subnet. | [optional] 

## Methods

### NewListSubnetsResponseV2SubnetsInner

`func NewListSubnetsResponseV2SubnetsInner(id string, vpc string, availabilityZone string, ) *ListSubnetsResponseV2SubnetsInner`

NewListSubnetsResponseV2SubnetsInner instantiates a new ListSubnetsResponseV2SubnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnetsResponseV2SubnetsInnerWithDefaults

`func NewListSubnetsResponseV2SubnetsInnerWithDefaults() *ListSubnetsResponseV2SubnetsInner`

NewListSubnetsResponseV2SubnetsInnerWithDefaults instantiates a new ListSubnetsResponseV2SubnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSubnetsResponseV2SubnetsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSubnetsResponseV2SubnetsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSubnetsResponseV2SubnetsInner) SetId(v string)`

SetId sets Id field to given value.


### GetVpc

`func (o *ListSubnetsResponseV2SubnetsInner) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *ListSubnetsResponseV2SubnetsInner) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *ListSubnetsResponseV2SubnetsInner) SetVpc(v string)`

SetVpc sets Vpc field to given value.


### GetAvailabilityZone

`func (o *ListSubnetsResponseV2SubnetsInner) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ListSubnetsResponseV2SubnetsInner) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ListSubnetsResponseV2SubnetsInner) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetName

`func (o *ListSubnetsResponseV2SubnetsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSubnetsResponseV2SubnetsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSubnetsResponseV2SubnetsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSubnetsResponseV2SubnetsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCidrBlock

`func (o *ListSubnetsResponseV2SubnetsInner) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ListSubnetsResponseV2SubnetsInner) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ListSubnetsResponseV2SubnetsInner) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *ListSubnetsResponseV2SubnetsInner) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


