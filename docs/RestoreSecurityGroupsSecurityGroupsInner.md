# RestoreSecurityGroupsSecurityGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the security group. | 
**Id** | **string** | Cloud-provider-assigned ID of the security group. | 
**Vpc** | **string** | Cloud-provider-assigned ID of the VPC the security group is in. | 

## Methods

### NewRestoreSecurityGroupsSecurityGroupsInner

`func NewRestoreSecurityGroupsSecurityGroupsInner(name string, id string, vpc string, ) *RestoreSecurityGroupsSecurityGroupsInner`

NewRestoreSecurityGroupsSecurityGroupsInner instantiates a new RestoreSecurityGroupsSecurityGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSecurityGroupsSecurityGroupsInnerWithDefaults

`func NewRestoreSecurityGroupsSecurityGroupsInnerWithDefaults() *RestoreSecurityGroupsSecurityGroupsInner`

NewRestoreSecurityGroupsSecurityGroupsInnerWithDefaults instantiates a new RestoreSecurityGroupsSecurityGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreSecurityGroupsSecurityGroupsInner) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreSecurityGroupsSecurityGroupsInner) SetId(v string)`

SetId sets Id field to given value.


### GetVpc

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *RestoreSecurityGroupsSecurityGroupsInner) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *RestoreSecurityGroupsSecurityGroupsInner) SetVpc(v string)`

SetVpc sets Vpc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


