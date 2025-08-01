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

// checks if the QueryDBResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryDBResultResponse{}

// QueryDBResultResponse struct for QueryDBResultResponse
type QueryDBResultResponse struct {
	// List of the column names in the returned records.
	Columns []string `json:"columns"`
	// List of records, returned as an array of arrays. The inner array contains the values of the columns in the same order as returned in `columns`. 
	Records [][]string `json:"records"`
	// Cursor that points to the first record of the next page of results. Pass this value in the next request. 
	NextToken NullableString `json:"nextToken"`
}

type _QueryDBResultResponse QueryDBResultResponse

// NewQueryDBResultResponse instantiates a new QueryDBResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryDBResultResponse(columns []string, records [][]string, nextToken NullableString) *QueryDBResultResponse {
	this := QueryDBResultResponse{}
	this.Columns = columns
	this.Records = records
	this.NextToken = nextToken
	return &this
}

// NewQueryDBResultResponseWithDefaults instantiates a new QueryDBResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryDBResultResponseWithDefaults() *QueryDBResultResponse {
	this := QueryDBResultResponse{}
	return &this
}

// GetColumns returns the Columns field value
func (o *QueryDBResultResponse) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *QueryDBResultResponse) GetColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *QueryDBResultResponse) SetColumns(v []string) {
	o.Columns = v
}

// GetRecords returns the Records field value
func (o *QueryDBResultResponse) GetRecords() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *QueryDBResultResponse) GetRecordsOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *QueryDBResultResponse) SetRecords(v [][]string) {
	o.Records = v
}

// GetNextToken returns the NextToken field value
// If the value is explicit nil, the zero value for string will be returned
func (o *QueryDBResultResponse) GetNextToken() string {
	if o == nil || o.NextToken.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextToken.Get()
}

// GetNextTokenOk returns a tuple with the NextToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryDBResultResponse) GetNextTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextToken.Get(), o.NextToken.IsSet()
}

// SetNextToken sets field value
func (o *QueryDBResultResponse) SetNextToken(v string) {
	o.NextToken.Set(&v)
}

func (o QueryDBResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryDBResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["columns"] = o.Columns
	toSerialize["records"] = o.Records
	toSerialize["nextToken"] = o.NextToken.Get()
	return toSerialize, nil
}

func (o *QueryDBResultResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"columns",
		"records",
		"nextToken",
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

	varQueryDBResultResponse := _QueryDBResultResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryDBResultResponse)

	if err != nil {
		return err
	}

	*o = QueryDBResultResponse(varQueryDBResultResponse)

	return err
}

type NullableQueryDBResultResponse struct {
	value *QueryDBResultResponse
	isSet bool
}

func (v NullableQueryDBResultResponse) Get() *QueryDBResultResponse {
	return v.value
}

func (v *NullableQueryDBResultResponse) Set(val *QueryDBResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryDBResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryDBResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryDBResultResponse(val *QueryDBResultResponse) *NullableQueryDBResultResponse {
	return &NullableQueryDBResultResponse{value: val, isSet: true}
}

func (v NullableQueryDBResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryDBResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


