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

// checks if the SubnetsCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubnetsCondition{}

// SubnetsCondition struct for SubnetsCondition
type SubnetsCondition struct {
	Operator ListOperators `json:"operator"`
	Subnets []string `json:"subnets"`
}

type _SubnetsCondition SubnetsCondition

// NewSubnetsCondition instantiates a new SubnetsCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetsCondition(operator ListOperators, subnets []string) *SubnetsCondition {
	this := SubnetsCondition{}
	this.Operator = operator
	this.Subnets = subnets
	return &this
}

// NewSubnetsConditionWithDefaults instantiates a new SubnetsCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetsConditionWithDefaults() *SubnetsCondition {
	this := SubnetsCondition{}
	return &this
}

// GetOperator returns the Operator field value
func (o *SubnetsCondition) GetOperator() ListOperators {
	if o == nil {
		var ret ListOperators
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SubnetsCondition) GetOperatorOk() (*ListOperators, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SubnetsCondition) SetOperator(v ListOperators) {
	o.Operator = v
}

// GetSubnets returns the Subnets field value
func (o *SubnetsCondition) GetSubnets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value
// and a boolean to check if the value has been set.
func (o *SubnetsCondition) GetSubnetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subnets, true
}

// SetSubnets sets field value
func (o *SubnetsCondition) SetSubnets(v []string) {
	o.Subnets = v
}

func (o SubnetsCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubnetsCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	toSerialize["subnets"] = o.Subnets
	return toSerialize, nil
}

func (o *SubnetsCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operator",
		"subnets",
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

	varSubnetsCondition := _SubnetsCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubnetsCondition)

	if err != nil {
		return err
	}

	*o = SubnetsCondition(varSubnetsCondition)

	return err
}

type NullableSubnetsCondition struct {
	value *SubnetsCondition
	isSet bool
}

func (v NullableSubnetsCondition) Get() *SubnetsCondition {
	return v.value
}

func (v *NullableSubnetsCondition) Set(val *SubnetsCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetsCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetsCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetsCondition(val *SubnetsCondition) *NullableSubnetsCondition {
	return &NullableSubnetsCondition{value: val, isSet: true}
}

func (v NullableSubnetsCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetsCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


