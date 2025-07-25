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

// checks if the UpdateRestoreAccountConnectivityConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRestoreAccountConnectivityConfigRequest{}

// UpdateRestoreAccountConnectivityConfigRequest struct for UpdateRestoreAccountConnectivityConfigRequest
type UpdateRestoreAccountConnectivityConfigRequest struct {
	Aws NullableAwsRestoreAccountConnectivityConfig `json:"aws,omitempty"`
}

// NewUpdateRestoreAccountConnectivityConfigRequest instantiates a new UpdateRestoreAccountConnectivityConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRestoreAccountConnectivityConfigRequest() *UpdateRestoreAccountConnectivityConfigRequest {
	this := UpdateRestoreAccountConnectivityConfigRequest{}
	return &this
}

// NewUpdateRestoreAccountConnectivityConfigRequestWithDefaults instantiates a new UpdateRestoreAccountConnectivityConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRestoreAccountConnectivityConfigRequestWithDefaults() *UpdateRestoreAccountConnectivityConfigRequest {
	this := UpdateRestoreAccountConnectivityConfigRequest{}
	return &this
}

// GetAws returns the Aws field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRestoreAccountConnectivityConfigRequest) GetAws() AwsRestoreAccountConnectivityConfig {
	if o == nil || IsNil(o.Aws.Get()) {
		var ret AwsRestoreAccountConnectivityConfig
		return ret
	}
	return *o.Aws.Get()
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRestoreAccountConnectivityConfigRequest) GetAwsOk() (*AwsRestoreAccountConnectivityConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aws.Get(), o.Aws.IsSet()
}

// HasAws returns a boolean if a field has been set.
func (o *UpdateRestoreAccountConnectivityConfigRequest) HasAws() bool {
	if o != nil && o.Aws.IsSet() {
		return true
	}

	return false
}

// SetAws gets a reference to the given NullableAwsRestoreAccountConnectivityConfig and assigns it to the Aws field.
func (o *UpdateRestoreAccountConnectivityConfigRequest) SetAws(v AwsRestoreAccountConnectivityConfig) {
	o.Aws.Set(&v)
}
// SetAwsNil sets the value for Aws to be an explicit nil
func (o *UpdateRestoreAccountConnectivityConfigRequest) SetAwsNil() {
	o.Aws.Set(nil)
}

// UnsetAws ensures that no value is present for Aws, not even an explicit nil
func (o *UpdateRestoreAccountConnectivityConfigRequest) UnsetAws() {
	o.Aws.Unset()
}

func (o UpdateRestoreAccountConnectivityConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRestoreAccountConnectivityConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Aws.IsSet() {
		toSerialize["aws"] = o.Aws.Get()
	}
	return toSerialize, nil
}

type NullableUpdateRestoreAccountConnectivityConfigRequest struct {
	value *UpdateRestoreAccountConnectivityConfigRequest
	isSet bool
}

func (v NullableUpdateRestoreAccountConnectivityConfigRequest) Get() *UpdateRestoreAccountConnectivityConfigRequest {
	return v.value
}

func (v *NullableUpdateRestoreAccountConnectivityConfigRequest) Set(val *UpdateRestoreAccountConnectivityConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRestoreAccountConnectivityConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRestoreAccountConnectivityConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRestoreAccountConnectivityConfigRequest(val *UpdateRestoreAccountConnectivityConfigRequest) *NullableUpdateRestoreAccountConnectivityConfigRequest {
	return &NullableUpdateRestoreAccountConnectivityConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateRestoreAccountConnectivityConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRestoreAccountConnectivityConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


