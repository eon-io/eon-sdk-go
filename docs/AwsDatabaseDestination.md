# AwsDatabaseDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreRegion** | **string** | Region to restore to. | 
**EncryptionKeyId** | **string** | ID of the key you want Eon to use for encrypting the restored resource. | 
**RestoredName** | **string** | Name to assign to the restored resource. | 
**SecurityGroups** | Pointer to **[]string** | List of security group IDs to associate with the restored resource. Must be in the same VPC of &#x60;subnetGroup&#x60;.  | [optional] 
**SubnetGroup** | Pointer to **string** | Subnet group ID to associate with the restored resource. Must be in the same VPC of &#x60;securityGroup&#x60;.  | [optional] 
**DbInstanceClass** | Pointer to **string** | Instance class to use for the restored resource. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to apply to the restored instance as key-value pairs, where key and value are both strings.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 

## Methods

### NewAwsDatabaseDestination

`func NewAwsDatabaseDestination(restoreRegion string, encryptionKeyId string, restoredName string, ) *AwsDatabaseDestination`

NewAwsDatabaseDestination instantiates a new AwsDatabaseDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsDatabaseDestinationWithDefaults

`func NewAwsDatabaseDestinationWithDefaults() *AwsDatabaseDestination`

NewAwsDatabaseDestinationWithDefaults instantiates a new AwsDatabaseDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreRegion

`func (o *AwsDatabaseDestination) GetRestoreRegion() string`

GetRestoreRegion returns the RestoreRegion field if non-nil, zero value otherwise.

### GetRestoreRegionOk

`func (o *AwsDatabaseDestination) GetRestoreRegionOk() (*string, bool)`

GetRestoreRegionOk returns a tuple with the RestoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreRegion

`func (o *AwsDatabaseDestination) SetRestoreRegion(v string)`

SetRestoreRegion sets RestoreRegion field to given value.


### GetEncryptionKeyId

`func (o *AwsDatabaseDestination) GetEncryptionKeyId() string`

GetEncryptionKeyId returns the EncryptionKeyId field if non-nil, zero value otherwise.

### GetEncryptionKeyIdOk

`func (o *AwsDatabaseDestination) GetEncryptionKeyIdOk() (*string, bool)`

GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyId

`func (o *AwsDatabaseDestination) SetEncryptionKeyId(v string)`

SetEncryptionKeyId sets EncryptionKeyId field to given value.


### GetRestoredName

`func (o *AwsDatabaseDestination) GetRestoredName() string`

GetRestoredName returns the RestoredName field if non-nil, zero value otherwise.

### GetRestoredNameOk

`func (o *AwsDatabaseDestination) GetRestoredNameOk() (*string, bool)`

GetRestoredNameOk returns a tuple with the RestoredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredName

`func (o *AwsDatabaseDestination) SetRestoredName(v string)`

SetRestoredName sets RestoredName field to given value.


### GetSecurityGroups

`func (o *AwsDatabaseDestination) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AwsDatabaseDestination) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AwsDatabaseDestination) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AwsDatabaseDestination) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSubnetGroup

`func (o *AwsDatabaseDestination) GetSubnetGroup() string`

GetSubnetGroup returns the SubnetGroup field if non-nil, zero value otherwise.

### GetSubnetGroupOk

`func (o *AwsDatabaseDestination) GetSubnetGroupOk() (*string, bool)`

GetSubnetGroupOk returns a tuple with the SubnetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGroup

`func (o *AwsDatabaseDestination) SetSubnetGroup(v string)`

SetSubnetGroup sets SubnetGroup field to given value.

### HasSubnetGroup

`func (o *AwsDatabaseDestination) HasSubnetGroup() bool`

HasSubnetGroup returns a boolean if a field has been set.

### GetDbInstanceClass

`func (o *AwsDatabaseDestination) GetDbInstanceClass() string`

GetDbInstanceClass returns the DbInstanceClass field if non-nil, zero value otherwise.

### GetDbInstanceClassOk

`func (o *AwsDatabaseDestination) GetDbInstanceClassOk() (*string, bool)`

GetDbInstanceClassOk returns a tuple with the DbInstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstanceClass

`func (o *AwsDatabaseDestination) SetDbInstanceClass(v string)`

SetDbInstanceClass sets DbInstanceClass field to given value.

### HasDbInstanceClass

`func (o *AwsDatabaseDestination) HasDbInstanceClass() bool`

HasDbInstanceClass returns a boolean if a field has been set.

### GetTags

`func (o *AwsDatabaseDestination) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsDatabaseDestination) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsDatabaseDestination) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsDatabaseDestination) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


