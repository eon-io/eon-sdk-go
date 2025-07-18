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

// checks if the QueryDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryDBResponse{}

// QueryDBResponse struct for QueryDBResponse
type QueryDBResponse struct {
	// Query ID. Can be used to call [Get Query Status](./get-query-status) and [Get Query Result](./get-query-result). 
	QueryId string `json:"queryId"`
}

type _QueryDBResponse QueryDBResponse

// NewQueryDBResponse instantiates a new QueryDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryDBResponse(queryId string) *QueryDBResponse {
	this := QueryDBResponse{}
	this.QueryId = queryId
	return &this
}

// NewQueryDBResponseWithDefaults instantiates a new QueryDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryDBResponseWithDefaults() *QueryDBResponse {
	this := QueryDBResponse{}
	return &this
}

// GetQueryId returns the QueryId field value
func (o *QueryDBResponse) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *QueryDBResponse) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *QueryDBResponse) SetQueryId(v string) {
	o.QueryId = v
}

func (o QueryDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	return toSerialize, nil
}

func (o *QueryDBResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryId",
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

	varQueryDBResponse := _QueryDBResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryDBResponse)

	if err != nil {
		return err
	}

	*o = QueryDBResponse(varQueryDBResponse)

	return err
}

type NullableQueryDBResponse struct {
	value *QueryDBResponse
	isSet bool
}

func (v NullableQueryDBResponse) Get() *QueryDBResponse {
	return v.value
}

func (v *NullableQueryDBResponse) Set(val *QueryDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryDBResponse(val *QueryDBResponse) *NullableQueryDBResponse {
	return &NullableQueryDBResponse{value: val, isSet: true}
}

func (v NullableQueryDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


