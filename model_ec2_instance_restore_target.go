/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Ec2InstanceRestoreTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ec2InstanceRestoreTarget{}

// Ec2InstanceRestoreTarget struct for Ec2InstanceRestoreTarget
type Ec2InstanceRestoreTarget struct {
	// Region to restore the instance to.
	Region string `json:"region"`
	// Instance type to use for the restored instance.
	InstanceType string `json:"instanceType"`
	// Subnet ID to associate with the restored instance.
	SubnetId string `json:"subnetId"`
	// List of security group IDs to associate with the restored instance.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`
	// Tags to apply to the restored instance as key-value pairs, where key and value are both strings.  **Example:** `{\"eon_api_restore\": \"true\"}` 
	Tags *map[string]string `json:"tags,omitempty"`
	// Volumes to restore and attach to the restored instance. Each item in the list corresponds to a volume to be restored, where `providerVolumeId` matches the volume's ID at the time of the snapshot. The root volume must be present in the list. 
	VolumeRestoreParameters []RestoreInstanceVolumeInput `json:"volumeRestoreParameters"`
}

type _Ec2InstanceRestoreTarget Ec2InstanceRestoreTarget

// NewEc2InstanceRestoreTarget instantiates a new Ec2InstanceRestoreTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEc2InstanceRestoreTarget(region string, instanceType string, subnetId string, volumeRestoreParameters []RestoreInstanceVolumeInput) *Ec2InstanceRestoreTarget {
	this := Ec2InstanceRestoreTarget{}
	this.Region = region
	this.InstanceType = instanceType
	this.SubnetId = subnetId
	this.VolumeRestoreParameters = volumeRestoreParameters
	return &this
}

// NewEc2InstanceRestoreTargetWithDefaults instantiates a new Ec2InstanceRestoreTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEc2InstanceRestoreTargetWithDefaults() *Ec2InstanceRestoreTarget {
	this := Ec2InstanceRestoreTarget{}
	return &this
}

// GetRegion returns the Region field value
func (o *Ec2InstanceRestoreTarget) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Ec2InstanceRestoreTarget) SetRegion(v string) {
	o.Region = v
}

// GetInstanceType returns the InstanceType field value
func (o *Ec2InstanceRestoreTarget) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *Ec2InstanceRestoreTarget) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetSubnetId returns the SubnetId field value
func (o *Ec2InstanceRestoreTarget) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *Ec2InstanceRestoreTarget) SetSubnetId(v string) {
	o.SubnetId = v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *Ec2InstanceRestoreTarget) GetSecurityGroupIds() []string {
	if o == nil || IsNil(o.SecurityGroupIds) {
		var ret []string
		return ret
	}
	return o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SecurityGroupIds) {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *Ec2InstanceRestoreTarget) HasSecurityGroupIds() bool {
	if o != nil && !IsNil(o.SecurityGroupIds) {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *Ec2InstanceRestoreTarget) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Ec2InstanceRestoreTarget) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Ec2InstanceRestoreTarget) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Ec2InstanceRestoreTarget) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetVolumeRestoreParameters returns the VolumeRestoreParameters field value
func (o *Ec2InstanceRestoreTarget) GetVolumeRestoreParameters() []RestoreInstanceVolumeInput {
	if o == nil {
		var ret []RestoreInstanceVolumeInput
		return ret
	}

	return o.VolumeRestoreParameters
}

// GetVolumeRestoreParametersOk returns a tuple with the VolumeRestoreParameters field value
// and a boolean to check if the value has been set.
func (o *Ec2InstanceRestoreTarget) GetVolumeRestoreParametersOk() ([]RestoreInstanceVolumeInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeRestoreParameters, true
}

// SetVolumeRestoreParameters sets field value
func (o *Ec2InstanceRestoreTarget) SetVolumeRestoreParameters(v []RestoreInstanceVolumeInput) {
	o.VolumeRestoreParameters = v
}

func (o Ec2InstanceRestoreTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ec2InstanceRestoreTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["instanceType"] = o.InstanceType
	toSerialize["subnetId"] = o.SubnetId
	if !IsNil(o.SecurityGroupIds) {
		toSerialize["securityGroupIds"] = o.SecurityGroupIds
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["volumeRestoreParameters"] = o.VolumeRestoreParameters
	return toSerialize, nil
}

func (o *Ec2InstanceRestoreTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"instanceType",
		"subnetId",
		"volumeRestoreParameters",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEc2InstanceRestoreTarget := _Ec2InstanceRestoreTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEc2InstanceRestoreTarget)

	if err != nil {
		return err
	}

	*o = Ec2InstanceRestoreTarget(varEc2InstanceRestoreTarget)

	return err
}

type NullableEc2InstanceRestoreTarget struct {
	value *Ec2InstanceRestoreTarget
	isSet bool
}

func (v NullableEc2InstanceRestoreTarget) Get() *Ec2InstanceRestoreTarget {
	return v.value
}

func (v *NullableEc2InstanceRestoreTarget) Set(val *Ec2InstanceRestoreTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableEc2InstanceRestoreTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableEc2InstanceRestoreTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEc2InstanceRestoreTarget(val *Ec2InstanceRestoreTarget) *NullableEc2InstanceRestoreTarget {
	return &NullableEc2InstanceRestoreTarget{value: val, isSet: true}
}

func (v NullableEc2InstanceRestoreTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEc2InstanceRestoreTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


