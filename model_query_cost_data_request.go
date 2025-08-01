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

// checks if the QueryCostDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryCostDataRequest{}

// QueryCostDataRequest struct for QueryCostDataRequest
type QueryCostDataRequest struct {
	TimePeriod TimePeriod `json:"timePeriod"`
	// Granularity for cost aggregation
	Granularity *string `json:"granularity,omitempty"`
	Filters *CostDataFilters `json:"filters,omitempty"`
	// Group results by specified dimensions
	GroupBy *string `json:"groupBy,omitempty"`
}

type _QueryCostDataRequest QueryCostDataRequest

// NewQueryCostDataRequest instantiates a new QueryCostDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCostDataRequest(timePeriod TimePeriod) *QueryCostDataRequest {
	this := QueryCostDataRequest{}
	this.TimePeriod = timePeriod
	var granularity string = "MONTHLY"
	this.Granularity = &granularity
	return &this
}

// NewQueryCostDataRequestWithDefaults instantiates a new QueryCostDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCostDataRequestWithDefaults() *QueryCostDataRequest {
	this := QueryCostDataRequest{}
	var granularity string = "MONTHLY"
	this.Granularity = &granularity
	return &this
}

// GetTimePeriod returns the TimePeriod field value
func (o *QueryCostDataRequest) GetTimePeriod() TimePeriod {
	if o == nil {
		var ret TimePeriod
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *QueryCostDataRequest) GetTimePeriodOk() (*TimePeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *QueryCostDataRequest) SetTimePeriod(v TimePeriod) {
	o.TimePeriod = v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *QueryCostDataRequest) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCostDataRequest) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *QueryCostDataRequest) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *QueryCostDataRequest) SetGranularity(v string) {
	o.Granularity = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *QueryCostDataRequest) GetFilters() CostDataFilters {
	if o == nil || IsNil(o.Filters) {
		var ret CostDataFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCostDataRequest) GetFiltersOk() (*CostDataFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *QueryCostDataRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given CostDataFilters and assigns it to the Filters field.
func (o *QueryCostDataRequest) SetFilters(v CostDataFilters) {
	o.Filters = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *QueryCostDataRequest) GetGroupBy() string {
	if o == nil || IsNil(o.GroupBy) {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCostDataRequest) GetGroupByOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *QueryCostDataRequest) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *QueryCostDataRequest) SetGroupBy(v string) {
	o.GroupBy = &v
}

func (o QueryCostDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCostDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timePeriod"] = o.TimePeriod
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	return toSerialize, nil
}

func (o *QueryCostDataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timePeriod",
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

	varQueryCostDataRequest := _QueryCostDataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryCostDataRequest)

	if err != nil {
		return err
	}

	*o = QueryCostDataRequest(varQueryCostDataRequest)

	return err
}

type NullableQueryCostDataRequest struct {
	value *QueryCostDataRequest
	isSet bool
}

func (v NullableQueryCostDataRequest) Get() *QueryCostDataRequest {
	return v.value
}

func (v *NullableQueryCostDataRequest) Set(val *QueryCostDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCostDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCostDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCostDataRequest(val *QueryCostDataRequest) *NullableQueryCostDataRequest {
	return &NullableQueryCostDataRequest{value: val, isSet: true}
}

func (v NullableQueryCostDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCostDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


