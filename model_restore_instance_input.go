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

// checks if the RestoreInstanceInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreInstanceInput{}

// RestoreInstanceInput struct for RestoreInstanceInput
type RestoreInstanceInput struct {
	// Eon-assigned ID of the restore account.
	RestoreAccountId string `json:"restoreAccountId"`
	Destination Ec2InstanceRestoreDestination `json:"destination"`
}

type _RestoreInstanceInput RestoreInstanceInput

// NewRestoreInstanceInput instantiates a new RestoreInstanceInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreInstanceInput(restoreAccountId string, destination Ec2InstanceRestoreDestination) *RestoreInstanceInput {
	this := RestoreInstanceInput{}
	this.RestoreAccountId = restoreAccountId
	this.Destination = destination
	return &this
}

// NewRestoreInstanceInputWithDefaults instantiates a new RestoreInstanceInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreInstanceInputWithDefaults() *RestoreInstanceInput {
	this := RestoreInstanceInput{}
	return &this
}

// GetRestoreAccountId returns the RestoreAccountId field value
func (o *RestoreInstanceInput) GetRestoreAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestoreAccountId
}

// GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field value
// and a boolean to check if the value has been set.
func (o *RestoreInstanceInput) GetRestoreAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestoreAccountId, true
}

// SetRestoreAccountId sets field value
func (o *RestoreInstanceInput) SetRestoreAccountId(v string) {
	o.RestoreAccountId = v
}

// GetDestination returns the Destination field value
func (o *RestoreInstanceInput) GetDestination() Ec2InstanceRestoreDestination {
	if o == nil {
		var ret Ec2InstanceRestoreDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *RestoreInstanceInput) GetDestinationOk() (*Ec2InstanceRestoreDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *RestoreInstanceInput) SetDestination(v Ec2InstanceRestoreDestination) {
	o.Destination = v
}

func (o RestoreInstanceInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreInstanceInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restoreAccountId"] = o.RestoreAccountId
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

func (o *RestoreInstanceInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"restoreAccountId",
		"destination",
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

	varRestoreInstanceInput := _RestoreInstanceInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestoreInstanceInput)

	if err != nil {
		return err
	}

	*o = RestoreInstanceInput(varRestoreInstanceInput)

	return err
}

type NullableRestoreInstanceInput struct {
	value *RestoreInstanceInput
	isSet bool
}

func (v NullableRestoreInstanceInput) Get() *RestoreInstanceInput {
	return v.value
}

func (v *NullableRestoreInstanceInput) Set(val *RestoreInstanceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreInstanceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreInstanceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreInstanceInput(val *RestoreInstanceInput) *NullableRestoreInstanceInput {
	return &NullableRestoreInstanceInput{value: val, isSet: true}
}

func (v NullableRestoreInstanceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreInstanceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


