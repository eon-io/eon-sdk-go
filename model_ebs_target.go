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

// checks if the EbsTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EbsTarget{}

// EbsTarget struct for EbsTarget
type EbsTarget struct {
	// Description to apply to the restored volume.  Default: No description. 
	Description *string `json:"description,omitempty"`
	// Tags to apply to the restored volume as key-value pairs, where key and value are both strings.  **Example:** `{\"eon_api_restore\": \"true\"}` 
	Tags *map[string]string `json:"tags,omitempty"`
	// ID of the key you want Eon to use for encrypting the restored volume.
	VolumeEncryptionKeyId string `json:"volumeEncryptionKeyId"`
	// Availability zone to restore the volume to.
	AvailabilityZone string `json:"availabilityZone"`
	VolumeSettings VolumeSettings `json:"volumeSettings"`
}

type _EbsTarget EbsTarget

// NewEbsTarget instantiates a new EbsTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEbsTarget(volumeEncryptionKeyId string, availabilityZone string, volumeSettings VolumeSettings) *EbsTarget {
	this := EbsTarget{}
	this.VolumeEncryptionKeyId = volumeEncryptionKeyId
	this.AvailabilityZone = availabilityZone
	this.VolumeSettings = volumeSettings
	return &this
}

// NewEbsTargetWithDefaults instantiates a new EbsTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEbsTargetWithDefaults() *EbsTarget {
	this := EbsTarget{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EbsTarget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbsTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EbsTarget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EbsTarget) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EbsTarget) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbsTarget) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EbsTarget) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *EbsTarget) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetVolumeEncryptionKeyId returns the VolumeEncryptionKeyId field value
func (o *EbsTarget) GetVolumeEncryptionKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeEncryptionKeyId
}

// GetVolumeEncryptionKeyIdOk returns a tuple with the VolumeEncryptionKeyId field value
// and a boolean to check if the value has been set.
func (o *EbsTarget) GetVolumeEncryptionKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeEncryptionKeyId, true
}

// SetVolumeEncryptionKeyId sets field value
func (o *EbsTarget) SetVolumeEncryptionKeyId(v string) {
	o.VolumeEncryptionKeyId = v
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *EbsTarget) GetAvailabilityZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *EbsTarget) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *EbsTarget) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetVolumeSettings returns the VolumeSettings field value
func (o *EbsTarget) GetVolumeSettings() VolumeSettings {
	if o == nil {
		var ret VolumeSettings
		return ret
	}

	return o.VolumeSettings
}

// GetVolumeSettingsOk returns a tuple with the VolumeSettings field value
// and a boolean to check if the value has been set.
func (o *EbsTarget) GetVolumeSettingsOk() (*VolumeSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeSettings, true
}

// SetVolumeSettings sets field value
func (o *EbsTarget) SetVolumeSettings(v VolumeSettings) {
	o.VolumeSettings = v
}

func (o EbsTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EbsTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["volumeEncryptionKeyId"] = o.VolumeEncryptionKeyId
	toSerialize["availabilityZone"] = o.AvailabilityZone
	toSerialize["volumeSettings"] = o.VolumeSettings
	return toSerialize, nil
}

func (o *EbsTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"volumeEncryptionKeyId",
		"availabilityZone",
		"volumeSettings",
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

	varEbsTarget := _EbsTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEbsTarget)

	if err != nil {
		return err
	}

	*o = EbsTarget(varEbsTarget)

	return err
}

type NullableEbsTarget struct {
	value *EbsTarget
	isSet bool
}

func (v NullableEbsTarget) Get() *EbsTarget {
	return v.value
}

func (v *NullableEbsTarget) Set(val *EbsTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableEbsTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableEbsTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEbsTarget(val *EbsTarget) *NullableEbsTarget {
	return &NullableEbsTarget{value: val, isSet: true}
}

func (v NullableEbsTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEbsTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


