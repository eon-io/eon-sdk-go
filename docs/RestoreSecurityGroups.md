# RestoreSecurityGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The region to restore | 
**SecurityGroups** | [**[]RestoreSecurityGroupsSecurityGroupsInner**](RestoreSecurityGroupsSecurityGroupsInner.md) | The available zones for the region | 

## Methods

### NewRestoreSecurityGroups

`func NewRestoreSecurityGroups(region string, securityGroups []RestoreSecurityGroupsSecurityGroupsInner, ) *RestoreSecurityGroups`

NewRestoreSecurityGroups instantiates a new RestoreSecurityGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSecurityGroupsWithDefaults

`func NewRestoreSecurityGroupsWithDefaults() *RestoreSecurityGroups`

NewRestoreSecurityGroupsWithDefaults instantiates a new RestoreSecurityGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *RestoreSecurityGroups) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RestoreSecurityGroups) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RestoreSecurityGroups) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSecurityGroups

`func (o *RestoreSecurityGroups) GetSecurityGroups() []RestoreSecurityGroupsSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RestoreSecurityGroups) GetSecurityGroupsOk() (*[]RestoreSecurityGroupsSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RestoreSecurityGroups) SetSecurityGroups(v []RestoreSecurityGroupsSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


