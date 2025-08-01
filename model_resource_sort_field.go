/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"encoding/json"
	"fmt"
)

// ResourceSortField Field to sort by.
type ResourceSortField string

// List of ResourceSortField
const (
	ID ResourceSortField = "id"
	BACKUP_STATUS ResourceSortField = "backupStatus"
	ACCOUNT_ID ResourceSortField = "accountId"
	RESOURCE_NAME ResourceSortField = "resourceName"
	ENVIRONMENT ResourceSortField = "environment"
	PROVIDER_RESOURCE_ID ResourceSortField = "providerResourceId"
	SIZE_BYTES ResourceSortField = "sizeBytes"
	SNAPSHOT_STORAGE_SIZE_BYTES ResourceSortField = "snapshotStorageSizeBytes"
	EON_SNAPSHOT_COUNT ResourceSortField = "eonSnapshotCount"
	NON_EON_SNAPSHOT_COUNT ResourceSortField = "nonEonSnapshotCount"
	LATEST_SNAPSHOT_TIME ResourceSortField = "latestSnapshotTime"
	OLDEST_SNAPSHOT_TIME ResourceSortField = "oldestSnapshotTime"
)

// All allowed values of ResourceSortField enum
var AllowedResourceSortFieldEnumValues = []ResourceSortField{
	"id",
	"backupStatus",
	"accountId",
	"resourceName",
	"environment",
	"providerResourceId",
	"sizeBytes",
	"snapshotStorageSizeBytes",
	"eonSnapshotCount",
	"nonEonSnapshotCount",
	"latestSnapshotTime",
	"oldestSnapshotTime",
}

func (v *ResourceSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceSortField(value)
	for _, existing := range AllowedResourceSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceSortField", value)
}

// NewResourceSortFieldFromValue returns a pointer to a valid ResourceSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceSortFieldFromValue(v string) (*ResourceSortField, error) {
	ev := ResourceSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceSortField: valid values are %v", v, AllowedResourceSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceSortField) IsValid() bool {
	for _, existing := range AllowedResourceSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceSortField value
func (v ResourceSortField) Ptr() *ResourceSortField {
	return &v
}

type NullableResourceSortField struct {
	value *ResourceSortField
	isSet bool
}

func (v NullableResourceSortField) Get() *ResourceSortField {
	return v.value
}

func (v *NullableResourceSortField) Set(val *ResourceSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSortField(val *ResourceSortField) *NullableResourceSortField {
	return &NullableResourceSortField{value: val, isSet: true}
}

func (v NullableResourceSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

