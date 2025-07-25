/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the DailyStorageSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailyStorageSummary{}

// DailyStorageSummary struct for DailyStorageSummary
type DailyStorageSummary struct {
	// Date the storage is calculated for.
	SummaryDay time.Time `json:"summaryDay"`
	// Total snapshot storage for the day, in bytes.
	DailyStorageBytes int64 `json:"dailyStorageBytes"`
}

type _DailyStorageSummary DailyStorageSummary

// NewDailyStorageSummary instantiates a new DailyStorageSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyStorageSummary(summaryDay time.Time, dailyStorageBytes int64) *DailyStorageSummary {
	this := DailyStorageSummary{}
	this.SummaryDay = summaryDay
	this.DailyStorageBytes = dailyStorageBytes
	return &this
}

// NewDailyStorageSummaryWithDefaults instantiates a new DailyStorageSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyStorageSummaryWithDefaults() *DailyStorageSummary {
	this := DailyStorageSummary{}
	return &this
}

// GetSummaryDay returns the SummaryDay field value
func (o *DailyStorageSummary) GetSummaryDay() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SummaryDay
}

// GetSummaryDayOk returns a tuple with the SummaryDay field value
// and a boolean to check if the value has been set.
func (o *DailyStorageSummary) GetSummaryDayOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryDay, true
}

// SetSummaryDay sets field value
func (o *DailyStorageSummary) SetSummaryDay(v time.Time) {
	o.SummaryDay = v
}

// GetDailyStorageBytes returns the DailyStorageBytes field value
func (o *DailyStorageSummary) GetDailyStorageBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DailyStorageBytes
}

// GetDailyStorageBytesOk returns a tuple with the DailyStorageBytes field value
// and a boolean to check if the value has been set.
func (o *DailyStorageSummary) GetDailyStorageBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyStorageBytes, true
}

// SetDailyStorageBytes sets field value
func (o *DailyStorageSummary) SetDailyStorageBytes(v int64) {
	o.DailyStorageBytes = v
}

func (o DailyStorageSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyStorageSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["summaryDay"] = o.SummaryDay
	toSerialize["dailyStorageBytes"] = o.DailyStorageBytes
	return toSerialize, nil
}

func (o *DailyStorageSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"summaryDay",
		"dailyStorageBytes",
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

	varDailyStorageSummary := _DailyStorageSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDailyStorageSummary)

	if err != nil {
		return err
	}

	*o = DailyStorageSummary(varDailyStorageSummary)

	return err
}

type NullableDailyStorageSummary struct {
	value *DailyStorageSummary
	isSet bool
}

func (v NullableDailyStorageSummary) Get() *DailyStorageSummary {
	return v.value
}

func (v *NullableDailyStorageSummary) Set(val *DailyStorageSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyStorageSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyStorageSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyStorageSummary(val *DailyStorageSummary) *NullableDailyStorageSummary {
	return &NullableDailyStorageSummary{value: val, isSet: true}
}

func (v NullableDailyStorageSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyStorageSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


