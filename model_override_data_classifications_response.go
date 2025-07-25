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

// checks if the OverrideDataClassificationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OverrideDataClassificationsResponse{}

// OverrideDataClassificationsResponse struct for OverrideDataClassificationsResponse
type OverrideDataClassificationsResponse struct {
	// List of data classes that were overridden..
	DataClasses []DataClass `json:"dataClasses,omitempty"`
}

// NewOverrideDataClassificationsResponse instantiates a new OverrideDataClassificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverrideDataClassificationsResponse() *OverrideDataClassificationsResponse {
	this := OverrideDataClassificationsResponse{}
	return &this
}

// NewOverrideDataClassificationsResponseWithDefaults instantiates a new OverrideDataClassificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverrideDataClassificationsResponseWithDefaults() *OverrideDataClassificationsResponse {
	this := OverrideDataClassificationsResponse{}
	return &this
}

// GetDataClasses returns the DataClasses field value if set, zero value otherwise.
func (o *OverrideDataClassificationsResponse) GetDataClasses() []DataClass {
	if o == nil || IsNil(o.DataClasses) {
		var ret []DataClass
		return ret
	}
	return o.DataClasses
}

// GetDataClassesOk returns a tuple with the DataClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OverrideDataClassificationsResponse) GetDataClassesOk() ([]DataClass, bool) {
	if o == nil || IsNil(o.DataClasses) {
		return nil, false
	}
	return o.DataClasses, true
}

// HasDataClasses returns a boolean if a field has been set.
func (o *OverrideDataClassificationsResponse) HasDataClasses() bool {
	if o != nil && !IsNil(o.DataClasses) {
		return true
	}

	return false
}

// SetDataClasses gets a reference to the given []DataClass and assigns it to the DataClasses field.
func (o *OverrideDataClassificationsResponse) SetDataClasses(v []DataClass) {
	o.DataClasses = v
}

func (o OverrideDataClassificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OverrideDataClassificationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataClasses) {
		toSerialize["dataClasses"] = o.DataClasses
	}
	return toSerialize, nil
}

type NullableOverrideDataClassificationsResponse struct {
	value *OverrideDataClassificationsResponse
	isSet bool
}

func (v NullableOverrideDataClassificationsResponse) Get() *OverrideDataClassificationsResponse {
	return v.value
}

func (v *NullableOverrideDataClassificationsResponse) Set(val *OverrideDataClassificationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideDataClassificationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideDataClassificationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideDataClassificationsResponse(val *OverrideDataClassificationsResponse) *NullableOverrideDataClassificationsResponse {
	return &NullableOverrideDataClassificationsResponse{value: val, isSet: true}
}

func (v NullableOverrideDataClassificationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideDataClassificationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


