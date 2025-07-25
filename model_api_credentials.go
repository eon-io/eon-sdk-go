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

// checks if the ApiCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCredentials{}

// ApiCredentials struct for ApiCredentials
type ApiCredentials struct {
	// Your integration's client ID.
	ClientId string `json:"clientId"`
	// Your integration's client secret.
	ClientSecret string `json:"clientSecret"`
}

type _ApiCredentials ApiCredentials

// NewApiCredentials instantiates a new ApiCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCredentials(clientId string, clientSecret string) *ApiCredentials {
	this := ApiCredentials{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewApiCredentialsWithDefaults instantiates a new ApiCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCredentialsWithDefaults() *ApiCredentials {
	this := ApiCredentials{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ApiCredentials) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ApiCredentials) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ApiCredentials) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ApiCredentials) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ApiCredentials) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ApiCredentials) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o ApiCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	return toSerialize, nil
}

func (o *ApiCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"clientSecret",
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

	varApiCredentials := _ApiCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCredentials)

	if err != nil {
		return err
	}

	*o = ApiCredentials(varApiCredentials)

	return err
}

type NullableApiCredentials struct {
	value *ApiCredentials
	isSet bool
}

func (v NullableApiCredentials) Get() *ApiCredentials {
	return v.value
}

func (v *NullableApiCredentials) Set(val *ApiCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCredentials(val *ApiCredentials) *NullableApiCredentials {
	return &NullableApiCredentials{value: val, isSet: true}
}

func (v NullableApiCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


