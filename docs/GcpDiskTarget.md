# GcpDiskTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | Zone to restore the disk to. | 
**Settings** | [**GcpDiskSettings**](GcpDiskSettings.md) |  | 

## Methods

### NewGcpDiskTarget

`func NewGcpDiskTarget(zone string, settings GcpDiskSettings, ) *GcpDiskTarget`

NewGcpDiskTarget instantiates a new GcpDiskTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDiskTargetWithDefaults

`func NewGcpDiskTargetWithDefaults() *GcpDiskTarget`

NewGcpDiskTargetWithDefaults instantiates a new GcpDiskTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *GcpDiskTarget) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GcpDiskTarget) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GcpDiskTarget) SetZone(v string)`

SetZone sets Zone field to given value.


### GetSettings

`func (o *GcpDiskTarget) GetSettings() GcpDiskSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GcpDiskTarget) GetSettingsOk() (*GcpDiskSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GcpDiskTarget) SetSettings(v GcpDiskSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


