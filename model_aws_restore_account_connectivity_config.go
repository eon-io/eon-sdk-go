/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
)

// checks if the AwsRestoreAccountConnectivityConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsRestoreAccountConnectivityConfig{}

// AwsRestoreAccountConnectivityConfig struct for AwsRestoreAccountConnectivityConfig
type AwsRestoreAccountConnectivityConfig struct {
	// VPCs to configure.
	VpcConfigs []AwsVpcConnectivityConfig `json:"vpcConfigs,omitempty"`
}

// NewAwsRestoreAccountConnectivityConfig instantiates a new AwsRestoreAccountConnectivityConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsRestoreAccountConnectivityConfig() *AwsRestoreAccountConnectivityConfig {
	this := AwsRestoreAccountConnectivityConfig{}
	return &this
}

// NewAwsRestoreAccountConnectivityConfigWithDefaults instantiates a new AwsRestoreAccountConnectivityConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsRestoreAccountConnectivityConfigWithDefaults() *AwsRestoreAccountConnectivityConfig {
	this := AwsRestoreAccountConnectivityConfig{}
	return &this
}

// GetVpcConfigs returns the VpcConfigs field value if set, zero value otherwise.
func (o *AwsRestoreAccountConnectivityConfig) GetVpcConfigs() []AwsVpcConnectivityConfig {
	if o == nil || IsNil(o.VpcConfigs) {
		var ret []AwsVpcConnectivityConfig
		return ret
	}
	return o.VpcConfigs
}

// GetVpcConfigsOk returns a tuple with the VpcConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsRestoreAccountConnectivityConfig) GetVpcConfigsOk() ([]AwsVpcConnectivityConfig, bool) {
	if o == nil || IsNil(o.VpcConfigs) {
		return nil, false
	}
	return o.VpcConfigs, true
}

// HasVpcConfigs returns a boolean if a field has been set.
func (o *AwsRestoreAccountConnectivityConfig) HasVpcConfigs() bool {
	if o != nil && !IsNil(o.VpcConfigs) {
		return true
	}

	return false
}

// SetVpcConfigs gets a reference to the given []AwsVpcConnectivityConfig and assigns it to the VpcConfigs field.
func (o *AwsRestoreAccountConnectivityConfig) SetVpcConfigs(v []AwsVpcConnectivityConfig) {
	o.VpcConfigs = v
}

func (o AwsRestoreAccountConnectivityConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsRestoreAccountConnectivityConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VpcConfigs) {
		toSerialize["vpcConfigs"] = o.VpcConfigs
	}
	return toSerialize, nil
}

type NullableAwsRestoreAccountConnectivityConfig struct {
	value *AwsRestoreAccountConnectivityConfig
	isSet bool
}

func (v NullableAwsRestoreAccountConnectivityConfig) Get() *AwsRestoreAccountConnectivityConfig {
	return v.value
}

func (v *NullableAwsRestoreAccountConnectivityConfig) Set(val *AwsRestoreAccountConnectivityConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsRestoreAccountConnectivityConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsRestoreAccountConnectivityConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsRestoreAccountConnectivityConfig(val *AwsRestoreAccountConnectivityConfig) *NullableAwsRestoreAccountConnectivityConfig {
	return &NullableAwsRestoreAccountConnectivityConfig{value: val, isSet: true}
}

func (v NullableAwsRestoreAccountConnectivityConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsRestoreAccountConnectivityConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


