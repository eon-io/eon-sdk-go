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

// checks if the HighFrequencyBackupResourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighFrequencyBackupResourceType{}

// HighFrequencyBackupResourceType struct for HighFrequencyBackupResourceType
type HighFrequencyBackupResourceType struct {
	ResourceType *ResourceType `json:"resourceType,omitempty"`
}

// NewHighFrequencyBackupResourceType instantiates a new HighFrequencyBackupResourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighFrequencyBackupResourceType() *HighFrequencyBackupResourceType {
	this := HighFrequencyBackupResourceType{}
	return &this
}

// NewHighFrequencyBackupResourceTypeWithDefaults instantiates a new HighFrequencyBackupResourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighFrequencyBackupResourceTypeWithDefaults() *HighFrequencyBackupResourceType {
	this := HighFrequencyBackupResourceType{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *HighFrequencyBackupResourceType) GetResourceType() ResourceType {
	if o == nil || IsNil(o.ResourceType) {
		var ret ResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighFrequencyBackupResourceType) GetResourceTypeOk() (*ResourceType, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *HighFrequencyBackupResourceType) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceType and assigns it to the ResourceType field.
func (o *HighFrequencyBackupResourceType) SetResourceType(v ResourceType) {
	o.ResourceType = &v
}

func (o HighFrequencyBackupResourceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighFrequencyBackupResourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	return toSerialize, nil
}

type NullableHighFrequencyBackupResourceType struct {
	value *HighFrequencyBackupResourceType
	isSet bool
}

func (v NullableHighFrequencyBackupResourceType) Get() *HighFrequencyBackupResourceType {
	return v.value
}

func (v *NullableHighFrequencyBackupResourceType) Set(val *HighFrequencyBackupResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableHighFrequencyBackupResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableHighFrequencyBackupResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighFrequencyBackupResourceType(val *HighFrequencyBackupResourceType) *NullableHighFrequencyBackupResourceType {
	return &NullableHighFrequencyBackupResourceType{value: val, isSet: true}
}

func (v NullableHighFrequencyBackupResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighFrequencyBackupResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


