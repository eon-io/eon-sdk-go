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

// checks if the RestoreFilesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreFilesRequest{}

// RestoreFilesRequest struct for RestoreFilesRequest
type RestoreFilesRequest struct {
	Destination ObjectStorageDestination `json:"destination"`
	// List of file paths to restore.
	Files []FilePath `json:"files"`
	// Eon-assigned ID of the restore account.
	RestoreAccountId string `json:"restoreAccountId"`
}

type _RestoreFilesRequest RestoreFilesRequest

// NewRestoreFilesRequest instantiates a new RestoreFilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreFilesRequest(destination ObjectStorageDestination, files []FilePath, restoreAccountId string) *RestoreFilesRequest {
	this := RestoreFilesRequest{}
	this.Destination = destination
	this.Files = files
	this.RestoreAccountId = restoreAccountId
	return &this
}

// NewRestoreFilesRequestWithDefaults instantiates a new RestoreFilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreFilesRequestWithDefaults() *RestoreFilesRequest {
	this := RestoreFilesRequest{}
	return &this
}

// GetDestination returns the Destination field value
func (o *RestoreFilesRequest) GetDestination() ObjectStorageDestination {
	if o == nil {
		var ret ObjectStorageDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *RestoreFilesRequest) GetDestinationOk() (*ObjectStorageDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *RestoreFilesRequest) SetDestination(v ObjectStorageDestination) {
	o.Destination = v
}

// GetFiles returns the Files field value
func (o *RestoreFilesRequest) GetFiles() []FilePath {
	if o == nil {
		var ret []FilePath
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *RestoreFilesRequest) GetFilesOk() ([]FilePath, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *RestoreFilesRequest) SetFiles(v []FilePath) {
	o.Files = v
}

// GetRestoreAccountId returns the RestoreAccountId field value
func (o *RestoreFilesRequest) GetRestoreAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestoreAccountId
}

// GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field value
// and a boolean to check if the value has been set.
func (o *RestoreFilesRequest) GetRestoreAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestoreAccountId, true
}

// SetRestoreAccountId sets field value
func (o *RestoreFilesRequest) SetRestoreAccountId(v string) {
	o.RestoreAccountId = v
}

func (o RestoreFilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreFilesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["files"] = o.Files
	toSerialize["restoreAccountId"] = o.RestoreAccountId
	return toSerialize, nil
}

func (o *RestoreFilesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"files",
		"restoreAccountId",
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

	varRestoreFilesRequest := _RestoreFilesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestoreFilesRequest)

	if err != nil {
		return err
	}

	*o = RestoreFilesRequest(varRestoreFilesRequest)

	return err
}

type NullableRestoreFilesRequest struct {
	value *RestoreFilesRequest
	isSet bool
}

func (v NullableRestoreFilesRequest) Get() *RestoreFilesRequest {
	return v.value
}

func (v *NullableRestoreFilesRequest) Set(val *RestoreFilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreFilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreFilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreFilesRequest(val *RestoreFilesRequest) *NullableRestoreFilesRequest {
	return &NullableRestoreFilesRequest{value: val, isSet: true}
}

func (v NullableRestoreFilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreFilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


