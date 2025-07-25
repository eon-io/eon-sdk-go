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

// checks if the SnapshotDateFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotDateFilters{}

// SnapshotDateFilters struct for SnapshotDateFilters
type SnapshotDateFilters struct {
	// Matches if `pointInTime` is on or later than this date. Must be in ISO 8601 YYYY-MM-DD format. 
	StartDate *string `json:"startDate,omitempty"`
	// Matches if `pointInTime` is on or earlier than this date. Must be in ISO 8601 YYYY-MM-DD format. 
	EndDate *string `json:"endDate,omitempty"`
}

// NewSnapshotDateFilters instantiates a new SnapshotDateFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotDateFilters() *SnapshotDateFilters {
	this := SnapshotDateFilters{}
	return &this
}

// NewSnapshotDateFiltersWithDefaults instantiates a new SnapshotDateFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotDateFiltersWithDefaults() *SnapshotDateFilters {
	this := SnapshotDateFilters{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SnapshotDateFilters) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDateFilters) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SnapshotDateFilters) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SnapshotDateFilters) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SnapshotDateFilters) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDateFilters) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SnapshotDateFilters) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SnapshotDateFilters) SetEndDate(v string) {
	o.EndDate = &v
}

func (o SnapshotDateFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotDateFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableSnapshotDateFilters struct {
	value *SnapshotDateFilters
	isSet bool
}

func (v NullableSnapshotDateFilters) Get() *SnapshotDateFilters {
	return v.value
}

func (v *NullableSnapshotDateFilters) Set(val *SnapshotDateFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotDateFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotDateFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotDateFilters(val *SnapshotDateFilters) *NullableSnapshotDateFilters {
	return &NullableSnapshotDateFilters{value: val, isSet: true}
}

func (v NullableSnapshotDateFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotDateFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


