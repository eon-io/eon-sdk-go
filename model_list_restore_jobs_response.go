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

// checks if the ListRestoreJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRestoreJobsResponse{}

// ListRestoreJobsResponse struct for ListRestoreJobsResponse
type ListRestoreJobsResponse struct {
	// List of retrieved restore jobs.
	Jobs []RestoreJob `json:"jobs"`
	// Cursor that points to the first record of the next page of results. Pass this value in the next request. 
	NextToken *string `json:"nextToken,omitempty"`
	// Total number of restore jobs that matched the filter options.
	TotalCount int32 `json:"totalCount"`
}

type _ListRestoreJobsResponse ListRestoreJobsResponse

// NewListRestoreJobsResponse instantiates a new ListRestoreJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRestoreJobsResponse(jobs []RestoreJob, totalCount int32) *ListRestoreJobsResponse {
	this := ListRestoreJobsResponse{}
	this.Jobs = jobs
	this.TotalCount = totalCount
	return &this
}

// NewListRestoreJobsResponseWithDefaults instantiates a new ListRestoreJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRestoreJobsResponseWithDefaults() *ListRestoreJobsResponse {
	this := ListRestoreJobsResponse{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *ListRestoreJobsResponse) GetJobs() []RestoreJob {
	if o == nil {
		var ret []RestoreJob
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *ListRestoreJobsResponse) GetJobsOk() ([]RestoreJob, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *ListRestoreJobsResponse) SetJobs(v []RestoreJob) {
	o.Jobs = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListRestoreJobsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRestoreJobsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListRestoreJobsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListRestoreJobsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTotalCount returns the TotalCount field value
func (o *ListRestoreJobsResponse) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListRestoreJobsResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListRestoreJobsResponse) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o ListRestoreJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRestoreJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ListRestoreJobsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobs",
		"totalCount",
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

	varListRestoreJobsResponse := _ListRestoreJobsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListRestoreJobsResponse)

	if err != nil {
		return err
	}

	*o = ListRestoreJobsResponse(varListRestoreJobsResponse)

	return err
}

type NullableListRestoreJobsResponse struct {
	value *ListRestoreJobsResponse
	isSet bool
}

func (v NullableListRestoreJobsResponse) Get() *ListRestoreJobsResponse {
	return v.value
}

func (v *NullableListRestoreJobsResponse) Set(val *ListRestoreJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRestoreJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRestoreJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRestoreJobsResponse(val *ListRestoreJobsResponse) *NullableListRestoreJobsResponse {
	return &NullableListRestoreJobsResponse{value: val, isSet: true}
}

func (v NullableListRestoreJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRestoreJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


