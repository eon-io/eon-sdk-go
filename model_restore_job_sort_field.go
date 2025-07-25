/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
	"fmt"
)

// RestoreJobSortField Field to sort by.
type RestoreJobSortField string

// List of RestoreJobSortField
const (
	RESTORE_JOB_SORT_CREATED_TIME RestoreJobSortField = "createdTime"
	RESTORE_JOB_SORT_START_TIME RestoreJobSortField = "startTime"
	RESTORE_JOB_SORT_END_TIME RestoreJobSortField = "endTime"
)

// All allowed values of RestoreJobSortField enum
var AllowedRestoreJobSortFieldEnumValues = []RestoreJobSortField{
	"createdTime",
	"startTime",
	"endTime",
}

func (v *RestoreJobSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RestoreJobSortField(value)
	for _, existing := range AllowedRestoreJobSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RestoreJobSortField", value)
}

// NewRestoreJobSortFieldFromValue returns a pointer to a valid RestoreJobSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRestoreJobSortFieldFromValue(v string) (*RestoreJobSortField, error) {
	ev := RestoreJobSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RestoreJobSortField: valid values are %v", v, AllowedRestoreJobSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RestoreJobSortField) IsValid() bool {
	for _, existing := range AllowedRestoreJobSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestoreJobSortField value
func (v RestoreJobSortField) Ptr() *RestoreJobSortField {
	return &v
}

type NullableRestoreJobSortField struct {
	value *RestoreJobSortField
	isSet bool
}

func (v NullableRestoreJobSortField) Get() *RestoreJobSortField {
	return v.value
}

func (v *NullableRestoreJobSortField) Set(val *RestoreJobSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreJobSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreJobSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreJobSortField(val *RestoreJobSortField) *NullableRestoreJobSortField {
	return &NullableRestoreJobSortField{value: val, isSet: true}
}

func (v NullableRestoreJobSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreJobSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

