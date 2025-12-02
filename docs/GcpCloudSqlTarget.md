# GcpCloudSqlTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | Zone to restore to. | 
**Name** | **string** | Instance name. | 
**NetworkType** | [**GcpNetworkType**](GcpNetworkType.md) |  | 
**NetworkName** | Pointer to **string** | Name of the VPC network to use. | [optional] 
**NetworkHostProject** | Pointer to **string** | ID of the project that hosts the VPC network. Applicable only when restoring to a shared VPC network.  | [optional] 
**Labels** | Pointer to **map[string]string** | Labels to apply to the restored Cloud SQL instance as key-value pairs, where key and value are both strings. These labels are always applied: &#x60;\&quot;eon-restore\&quot;: \&quot;true\&quot;&#x60;.  **Example:** &#x60;{\&quot;eon_api_restore\&quot;: \&quot;true\&quot;}&#x60;  | [optional] 

## Methods

### NewGcpCloudSqlTarget

`func NewGcpCloudSqlTarget(zone string, name string, networkType GcpNetworkType, ) *GcpCloudSqlTarget`

NewGcpCloudSqlTarget instantiates a new GcpCloudSqlTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCloudSqlTargetWithDefaults

`func NewGcpCloudSqlTargetWithDefaults() *GcpCloudSqlTarget`

NewGcpCloudSqlTargetWithDefaults instantiates a new GcpCloudSqlTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *GcpCloudSqlTarget) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GcpCloudSqlTarget) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GcpCloudSqlTarget) SetZone(v string)`

SetZone sets Zone field to given value.


### GetName

`func (o *GcpCloudSqlTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCloudSqlTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCloudSqlTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkType

`func (o *GcpCloudSqlTarget) GetNetworkType() GcpNetworkType`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *GcpCloudSqlTarget) GetNetworkTypeOk() (*GcpNetworkType, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *GcpCloudSqlTarget) SetNetworkType(v GcpNetworkType)`

SetNetworkType sets NetworkType field to given value.


### GetNetworkName

`func (o *GcpCloudSqlTarget) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GcpCloudSqlTarget) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GcpCloudSqlTarget) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *GcpCloudSqlTarget) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkHostProject

`func (o *GcpCloudSqlTarget) GetNetworkHostProject() string`

GetNetworkHostProject returns the NetworkHostProject field if non-nil, zero value otherwise.

### GetNetworkHostProjectOk

`func (o *GcpCloudSqlTarget) GetNetworkHostProjectOk() (*string, bool)`

GetNetworkHostProjectOk returns a tuple with the NetworkHostProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHostProject

`func (o *GcpCloudSqlTarget) SetNetworkHostProject(v string)`

SetNetworkHostProject sets NetworkHostProject field to given value.

### HasNetworkHostProject

`func (o *GcpCloudSqlTarget) HasNetworkHostProject() bool`

HasNetworkHostProject returns a boolean if a field has been set.

### GetLabels

`func (o *GcpCloudSqlTarget) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpCloudSqlTarget) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpCloudSqlTarget) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpCloudSqlTarget) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


