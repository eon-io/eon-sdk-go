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

// checks if the ConnectSourceAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectSourceAccountResponse{}

// ConnectSourceAccountResponse struct for ConnectSourceAccountResponse
type ConnectSourceAccountResponse struct {
	SourceAccount SourceAccount `json:"sourceAccount"`
}

type _ConnectSourceAccountResponse ConnectSourceAccountResponse

// NewConnectSourceAccountResponse instantiates a new ConnectSourceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectSourceAccountResponse(sourceAccount SourceAccount) *ConnectSourceAccountResponse {
	this := ConnectSourceAccountResponse{}
	this.SourceAccount = sourceAccount
	return &this
}

// NewConnectSourceAccountResponseWithDefaults instantiates a new ConnectSourceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectSourceAccountResponseWithDefaults() *ConnectSourceAccountResponse {
	this := ConnectSourceAccountResponse{}
	return &this
}

// GetSourceAccount returns the SourceAccount field value
func (o *ConnectSourceAccountResponse) GetSourceAccount() SourceAccount {
	if o == nil {
		var ret SourceAccount
		return ret
	}

	return o.SourceAccount
}

// GetSourceAccountOk returns a tuple with the SourceAccount field value
// and a boolean to check if the value has been set.
func (o *ConnectSourceAccountResponse) GetSourceAccountOk() (*SourceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAccount, true
}

// SetSourceAccount sets field value
func (o *ConnectSourceAccountResponse) SetSourceAccount(v SourceAccount) {
	o.SourceAccount = v
}

func (o ConnectSourceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectSourceAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceAccount"] = o.SourceAccount
	return toSerialize, nil
}

func (o *ConnectSourceAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceAccount",
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

	varConnectSourceAccountResponse := _ConnectSourceAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectSourceAccountResponse)

	if err != nil {
		return err
	}

	*o = ConnectSourceAccountResponse(varConnectSourceAccountResponse)

	return err
}

type NullableConnectSourceAccountResponse struct {
	value *ConnectSourceAccountResponse
	isSet bool
}

func (v NullableConnectSourceAccountResponse) Get() *ConnectSourceAccountResponse {
	return v.value
}

func (v *NullableConnectSourceAccountResponse) Set(val *ConnectSourceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectSourceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectSourceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectSourceAccountResponse(val *ConnectSourceAccountResponse) *NullableConnectSourceAccountResponse {
	return &NullableConnectSourceAccountResponse{value: val, isSet: true}
}

func (v NullableConnectSourceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectSourceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


