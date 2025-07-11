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

// checks if the SourceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAccount{}

// SourceAccount struct for SourceAccount
type SourceAccount struct {
	// Eon-assigned account ID.
	Id string `json:"id"`
	// Account display name in Eon.
	Name string `json:"name"`
	// Cloud-provider-assigned account ID.
	ProviderAccountId string `json:"providerAccountId"`
	Status AccountState `json:"status"`
	SourceAccountAttributes *AccountConfig `json:"sourceAccountAttributes,omitempty"`
}

type _SourceAccount SourceAccount

// NewSourceAccount instantiates a new SourceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAccount(id string, name string, providerAccountId string, status AccountState) *SourceAccount {
	this := SourceAccount{}
	this.Id = id
	this.Name = name
	this.ProviderAccountId = providerAccountId
	this.Status = status
	return &this
}

// NewSourceAccountWithDefaults instantiates a new SourceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAccountWithDefaults() *SourceAccount {
	this := SourceAccount{}
	return &this
}

// GetId returns the Id field value
func (o *SourceAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceAccount) SetName(v string) {
	o.Name = v
}

// GetProviderAccountId returns the ProviderAccountId field value
func (o *SourceAccount) GetProviderAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderAccountId
}

// GetProviderAccountIdOk returns a tuple with the ProviderAccountId field value
// and a boolean to check if the value has been set.
func (o *SourceAccount) GetProviderAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderAccountId, true
}

// SetProviderAccountId sets field value
func (o *SourceAccount) SetProviderAccountId(v string) {
	o.ProviderAccountId = v
}

// GetStatus returns the Status field value
func (o *SourceAccount) GetStatus() AccountState {
	if o == nil {
		var ret AccountState
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SourceAccount) GetStatusOk() (*AccountState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SourceAccount) SetStatus(v AccountState) {
	o.Status = v
}

// GetSourceAccountAttributes returns the SourceAccountAttributes field value if set, zero value otherwise.
func (o *SourceAccount) GetSourceAccountAttributes() AccountConfig {
	if o == nil || IsNil(o.SourceAccountAttributes) {
		var ret AccountConfig
		return ret
	}
	return *o.SourceAccountAttributes
}

// GetSourceAccountAttributesOk returns a tuple with the SourceAccountAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAccount) GetSourceAccountAttributesOk() (*AccountConfig, bool) {
	if o == nil || IsNil(o.SourceAccountAttributes) {
		return nil, false
	}
	return o.SourceAccountAttributes, true
}

// HasSourceAccountAttributes returns a boolean if a field has been set.
func (o *SourceAccount) HasSourceAccountAttributes() bool {
	if o != nil && !IsNil(o.SourceAccountAttributes) {
		return true
	}

	return false
}

// SetSourceAccountAttributes gets a reference to the given AccountConfig and assigns it to the SourceAccountAttributes field.
func (o *SourceAccount) SetSourceAccountAttributes(v AccountConfig) {
	o.SourceAccountAttributes = &v
}

func (o SourceAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["providerAccountId"] = o.ProviderAccountId
	toSerialize["status"] = o.Status
	if !IsNil(o.SourceAccountAttributes) {
		toSerialize["sourceAccountAttributes"] = o.SourceAccountAttributes
	}
	return toSerialize, nil
}

func (o *SourceAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"providerAccountId",
		"status",
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

	varSourceAccount := _SourceAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSourceAccount)

	if err != nil {
		return err
	}

	*o = SourceAccount(varSourceAccount)

	return err
}

type NullableSourceAccount struct {
	value *SourceAccount
	isSet bool
}

func (v NullableSourceAccount) Get() *SourceAccount {
	return v.value
}

func (v *NullableSourceAccount) Set(val *SourceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAccount(val *SourceAccount) *NullableSourceAccount {
	return &NullableSourceAccount{value: val, isSet: true}
}

func (v NullableSourceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


