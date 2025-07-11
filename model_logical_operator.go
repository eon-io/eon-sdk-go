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

// LogicalOperator Logical operator to apply to the list of `operands`. 
type LogicalOperator string

// List of LogicalOperator
const (
	AND_OPERATOR LogicalOperator = "AND"
	OR_OPERATOR LogicalOperator = "OR"
	LOGICAL_OPERATOR_UNSPECIFIED LogicalOperator = "UNSPECIFIED"
)

// All allowed values of LogicalOperator enum
var AllowedLogicalOperatorEnumValues = []LogicalOperator{
	"AND",
	"OR",
	"UNSPECIFIED",
}

func (v *LogicalOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogicalOperator(value)
	for _, existing := range AllowedLogicalOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogicalOperator", value)
}

// NewLogicalOperatorFromValue returns a pointer to a valid LogicalOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogicalOperatorFromValue(v string) (*LogicalOperator, error) {
	ev := LogicalOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogicalOperator: valid values are %v", v, AllowedLogicalOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogicalOperator) IsValid() bool {
	for _, existing := range AllowedLogicalOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogicalOperator value
func (v LogicalOperator) Ptr() *LogicalOperator {
	return &v
}

type NullableLogicalOperator struct {
	value *LogicalOperator
	isSet bool
}

func (v NullableLogicalOperator) Get() *LogicalOperator {
	return v.value
}

func (v *NullableLogicalOperator) Set(val *LogicalOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalOperator(val *LogicalOperator) *NullableLogicalOperator {
	return &NullableLogicalOperator{value: val, isSet: true}
}

func (v NullableLogicalOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

