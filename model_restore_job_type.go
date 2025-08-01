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

// RestoreJobType Type of restore. This value is a combination of the cloud provider, resource type, and restore method. 
type RestoreJobType string

// List of RestoreJobType
const (
	RESTORE_JOB_TYPE_UNSPECIFIED RestoreJobType = "RESTORE_JOB_TYPE_UNSPECIFIED"
	AWS_S3_OBJECT_RESTORE RestoreJobType = "AWS_S3_OBJECT_RESTORE"
	AWS_S3_BUCKET_RESTORE RestoreJobType = "AWS_S3_BUCKET_RESTORE"
	AWS_EC2_FILE_RESTORE RestoreJobType = "AWS_EC2_FILE_RESTORE"
	AWS_EC2_INSTANCE_RESTORE RestoreJobType = "AWS_EC2_INSTANCE_RESTORE"
	AWS_EC2_AMI_CONVERT RestoreJobType = "AWS_EC2_AMI_CONVERT"
	AWS_EC2_EBS_VOLUME_RESTORE RestoreJobType = "AWS_EC2_EBS_VOLUME_RESTORE"
	AWS_EC2_EBS_SNAPSHOT_CONVERT RestoreJobType = "AWS_EC2_EBS_SNAPSHOT_CONVERT"
	AWS_EC2_RECORD_RESTORE RestoreJobType = "AWS_EC2_RECORD_RESTORE"
	AWS_RDS_INSTANCE_RESTORE RestoreJobType = "AWS_RDS_INSTANCE_RESTORE"
	AWS_RDS_RECORD_RESTORE RestoreJobType = "AWS_RDS_RECORD_RESTORE"
	AWS_EKS_NAMESPACE_RESTORE RestoreJobType = "AWS_EKS_NAMESPACE_RESTORE"
	AWS_EKS_NAMESPACE_RECORD_RESTORE RestoreJobType = "AWS_EKS_NAMESPACE_RECORD_RESTORE"
	AWS_EKS_FILE_RESTORE RestoreJobType = "AWS_EKS_FILE_RESTORE"
	AWS_DYNAMO_DB_TABLE_RESTORE RestoreJobType = "AWS_DYNAMO_DB_TABLE_RESTORE"
	AWS_DYNAMO_DB_RECORD_RESTORE RestoreJobType = "AWS_DYNAMO_DB_RECORD_RESTORE"
	AZURE_STORAGE_ACCOUNT_RESTORE RestoreJobType = "AZURE_STORAGE_ACCOUNT_RESTORE"
	AZURE_STORAGE_ACCOUNT_BLOB_RESTORE RestoreJobType = "AZURE_STORAGE_ACCOUNT_BLOB_RESTORE"
	AZURE_VM_FILE_RESTORE RestoreJobType = "AZURE_VM_FILE_RESTORE"
	AZURE_VM_DISK_RESTORE RestoreJobType = "AZURE_VM_DISK_RESTORE"
	AZURE_VM_INSTANCE_RESTORE RestoreJobType = "AZURE_VM_INSTANCE_RESTORE"
	AZURE_VM_RECORD_RESTORE RestoreJobType = "AZURE_VM_RECORD_RESTORE"
	AZURE_SQL_DATABASE_RECORD_RESTORE RestoreJobType = "AZURE_SQL_DATABASE_RECORD_RESTORE"
	GCP_CLOUD_STORAGE_BUCKET_RESTORE RestoreJobType = "GCP_CLOUD_STORAGE_BUCKET_RESTORE"
	GCP_CLOUD_STORAGE_OBJECT_RESTORE RestoreJobType = "GCP_CLOUD_STORAGE_OBJECT_RESTORE"
	GCP_COMPUTE_ENGINE_FILE_RESTORE RestoreJobType = "GCP_COMPUTE_ENGINE_FILE_RESTORE"
	GCP_COMPUTE_ENGINE_DISK_RESTORE RestoreJobType = "GCP_COMPUTE_ENGINE_DISK_RESTORE"
	GCP_COMPUTE_ENGINE_INSTANCE_RESTORE RestoreJobType = "GCP_COMPUTE_ENGINE_INSTANCE_RESTORE"
	GCP_COMPUTE_ENGINE_RECORD_RESTORE RestoreJobType = "GCP_COMPUTE_ENGINE_RECORD_RESTORE"
	GCP_CLOUD_SQL_DATABASE_RESTORE RestoreJobType = "GCP_CLOUD_SQL_DATABASE_RESTORE"
	GCP_CLOUD_SQL_RECORD_RESTORE RestoreJobType = "GCP_CLOUD_SQL_RECORD_RESTORE"
)

// All allowed values of RestoreJobType enum
var AllowedRestoreJobTypeEnumValues = []RestoreJobType{
	"RESTORE_JOB_TYPE_UNSPECIFIED",
	"AWS_S3_OBJECT_RESTORE",
	"AWS_S3_BUCKET_RESTORE",
	"AWS_EC2_FILE_RESTORE",
	"AWS_EC2_INSTANCE_RESTORE",
	"AWS_EC2_AMI_CONVERT",
	"AWS_EC2_EBS_VOLUME_RESTORE",
	"AWS_EC2_EBS_SNAPSHOT_CONVERT",
	"AWS_EC2_RECORD_RESTORE",
	"AWS_RDS_INSTANCE_RESTORE",
	"AWS_RDS_RECORD_RESTORE",
	"AWS_EKS_NAMESPACE_RESTORE",
	"AWS_EKS_NAMESPACE_RECORD_RESTORE",
	"AWS_EKS_FILE_RESTORE",
	"AWS_DYNAMO_DB_TABLE_RESTORE",
	"AWS_DYNAMO_DB_RECORD_RESTORE",
	"AZURE_STORAGE_ACCOUNT_RESTORE",
	"AZURE_STORAGE_ACCOUNT_BLOB_RESTORE",
	"AZURE_VM_FILE_RESTORE",
	"AZURE_VM_DISK_RESTORE",
	"AZURE_VM_INSTANCE_RESTORE",
	"AZURE_VM_RECORD_RESTORE",
	"AZURE_SQL_DATABASE_RECORD_RESTORE",
	"GCP_CLOUD_STORAGE_BUCKET_RESTORE",
	"GCP_CLOUD_STORAGE_OBJECT_RESTORE",
	"GCP_COMPUTE_ENGINE_FILE_RESTORE",
	"GCP_COMPUTE_ENGINE_DISK_RESTORE",
	"GCP_COMPUTE_ENGINE_INSTANCE_RESTORE",
	"GCP_COMPUTE_ENGINE_RECORD_RESTORE",
	"GCP_CLOUD_SQL_DATABASE_RESTORE",
	"GCP_CLOUD_SQL_RECORD_RESTORE",
}

func (v *RestoreJobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RestoreJobType(value)
	for _, existing := range AllowedRestoreJobTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RestoreJobType", value)
}

// NewRestoreJobTypeFromValue returns a pointer to a valid RestoreJobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRestoreJobTypeFromValue(v string) (*RestoreJobType, error) {
	ev := RestoreJobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RestoreJobType: valid values are %v", v, AllowedRestoreJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RestoreJobType) IsValid() bool {
	for _, existing := range AllowedRestoreJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestoreJobType value
func (v RestoreJobType) Ptr() *RestoreJobType {
	return &v
}

type NullableRestoreJobType struct {
	value *RestoreJobType
	isSet bool
}

func (v NullableRestoreJobType) Get() *RestoreJobType {
	return v.value
}

func (v *NullableRestoreJobType) Set(val *RestoreJobType) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreJobType(val *RestoreJobType) *NullableRestoreJobType {
	return &NullableRestoreJobType{value: val, isSet: true}
}

func (v NullableRestoreJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

