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

// checks if the EbsSnapshotTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EbsSnapshotTarget{}

// EbsSnapshotTarget struct for EbsSnapshotTarget
type EbsSnapshotTarget struct {
	// Region to restore the EBS snapshot to.
	Region string `json:"region"`
	// Description to apply to the EBS snapshot.
	Description *string `json:"description,omitempty"`
	// Tags to apply to the EBS snapshot as key-value pairs, where key and value are both strings.  **Example:** `{\"eon_api_restore\": \"true\"}` 
	Tags *map[string]string `json:"tags,omitempty"`
	// ID of the key you want Eon to use for encrypting the EBS snapshot.
	SnapshotEncryptionKeyId string `json:"snapshotEncryptionKeyId"`
}

type _EbsSnapshotTarget EbsSnapshotTarget

// NewEbsSnapshotTarget instantiates a new EbsSnapshotTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEbsSnapshotTarget(region string, snapshotEncryptionKeyId string) *EbsSnapshotTarget {
	this := EbsSnapshotTarget{}
	this.Region = region
	this.SnapshotEncryptionKeyId = snapshotEncryptionKeyId
	return &this
}

// NewEbsSnapshotTargetWithDefaults instantiates a new EbsSnapshotTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEbsSnapshotTargetWithDefaults() *EbsSnapshotTarget {
	this := EbsSnapshotTarget{}
	return &this
}

// GetRegion returns the Region field value
func (o *EbsSnapshotTarget) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EbsSnapshotTarget) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *EbsSnapshotTarget) SetRegion(v string) {
	o.Region = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EbsSnapshotTarget) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbsSnapshotTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EbsSnapshotTarget) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EbsSnapshotTarget) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EbsSnapshotTarget) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EbsSnapshotTarget) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EbsSnapshotTarget) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *EbsSnapshotTarget) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetSnapshotEncryptionKeyId returns the SnapshotEncryptionKeyId field value
func (o *EbsSnapshotTarget) GetSnapshotEncryptionKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotEncryptionKeyId
}

// GetSnapshotEncryptionKeyIdOk returns a tuple with the SnapshotEncryptionKeyId field value
// and a boolean to check if the value has been set.
func (o *EbsSnapshotTarget) GetSnapshotEncryptionKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotEncryptionKeyId, true
}

// SetSnapshotEncryptionKeyId sets field value
func (o *EbsSnapshotTarget) SetSnapshotEncryptionKeyId(v string) {
	o.SnapshotEncryptionKeyId = v
}

func (o EbsSnapshotTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EbsSnapshotTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["snapshotEncryptionKeyId"] = o.SnapshotEncryptionKeyId
	return toSerialize, nil
}

func (o *EbsSnapshotTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"snapshotEncryptionKeyId",
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

	varEbsSnapshotTarget := _EbsSnapshotTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEbsSnapshotTarget)

	if err != nil {
		return err
	}

	*o = EbsSnapshotTarget(varEbsSnapshotTarget)

	return err
}

type NullableEbsSnapshotTarget struct {
	value *EbsSnapshotTarget
	isSet bool
}

func (v NullableEbsSnapshotTarget) Get() *EbsSnapshotTarget {
	return v.value
}

func (v *NullableEbsSnapshotTarget) Set(val *EbsSnapshotTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableEbsSnapshotTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableEbsSnapshotTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEbsSnapshotTarget(val *EbsSnapshotTarget) *NullableEbsSnapshotTarget {
	return &NullableEbsSnapshotTarget{value: val, isSet: true}
}

func (v NullableEbsSnapshotTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEbsSnapshotTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


