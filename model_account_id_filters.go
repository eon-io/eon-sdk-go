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

// checks if the AccountIdFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountIdFilters{}

// AccountIdFilters struct for AccountIdFilters
type AccountIdFilters struct {
	// Matches if any value in this list equals `accountId`.
	In []string `json:"in,omitempty"`
	// Matches if none of the values in this list equal `accountId`.
	NotIn []string `json:"notIn,omitempty"`
}

// NewAccountIdFilters instantiates a new AccountIdFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdFilters() *AccountIdFilters {
	this := AccountIdFilters{}
	return &this
}

// NewAccountIdFiltersWithDefaults instantiates a new AccountIdFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdFiltersWithDefaults() *AccountIdFilters {
	this := AccountIdFilters{}
	return &this
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *AccountIdFilters) GetIn() []string {
	if o == nil || IsNil(o.In) {
		var ret []string
		return ret
	}
	return o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdFilters) GetInOk() ([]string, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *AccountIdFilters) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given []string and assigns it to the In field.
func (o *AccountIdFilters) SetIn(v []string) {
	o.In = v
}

// GetNotIn returns the NotIn field value if set, zero value otherwise.
func (o *AccountIdFilters) GetNotIn() []string {
	if o == nil || IsNil(o.NotIn) {
		var ret []string
		return ret
	}
	return o.NotIn
}

// GetNotInOk returns a tuple with the NotIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdFilters) GetNotInOk() ([]string, bool) {
	if o == nil || IsNil(o.NotIn) {
		return nil, false
	}
	return o.NotIn, true
}

// HasNotIn returns a boolean if a field has been set.
func (o *AccountIdFilters) HasNotIn() bool {
	if o != nil && !IsNil(o.NotIn) {
		return true
	}

	return false
}

// SetNotIn gets a reference to the given []string and assigns it to the NotIn field.
func (o *AccountIdFilters) SetNotIn(v []string) {
	o.NotIn = v
}

func (o AccountIdFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountIdFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !IsNil(o.NotIn) {
		toSerialize["notIn"] = o.NotIn
	}
	return toSerialize, nil
}

type NullableAccountIdFilters struct {
	value *AccountIdFilters
	isSet bool
}

func (v NullableAccountIdFilters) Get() *AccountIdFilters {
	return v.value
}

func (v *NullableAccountIdFilters) Set(val *AccountIdFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdFilters(val *AccountIdFilters) *NullableAccountIdFilters {
	return &NullableAccountIdFilters{value: val, isSet: true}
}

func (v NullableAccountIdFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


