/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
)

// checks if the ListRestoreJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRestoreJobsResponse{}

// ListRestoreJobsResponse struct for ListRestoreJobsResponse
type ListRestoreJobsResponse struct {
	// List of retrieved restore jobs.
	Jobs []RestoreJob `json:"jobs,omitempty"`
	// Cursor that points to the first record of the next page of results. Pass this value in the next request. 
	NextToken *string `json:"nextToken,omitempty"`
	// Total number of restore jobs that matched the filter options.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewListRestoreJobsResponse instantiates a new ListRestoreJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRestoreJobsResponse() *ListRestoreJobsResponse {
	this := ListRestoreJobsResponse{}
	return &this
}

// NewListRestoreJobsResponseWithDefaults instantiates a new ListRestoreJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRestoreJobsResponseWithDefaults() *ListRestoreJobsResponse {
	this := ListRestoreJobsResponse{}
	return &this
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *ListRestoreJobsResponse) GetJobs() []RestoreJob {
	if o == nil || IsNil(o.Jobs) {
		var ret []RestoreJob
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRestoreJobsResponse) GetJobsOk() ([]RestoreJob, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *ListRestoreJobsResponse) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []RestoreJob and assigns it to the Jobs field.
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

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListRestoreJobsResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRestoreJobsResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListRestoreJobsResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListRestoreJobsResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
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
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
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


