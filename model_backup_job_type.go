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

// BackupJobType Backup Job type. 
type BackupJobType string

// List of BackupJobType
const (
	BACKUP_JOB_TYPE_UNSPECIFIED BackupJobType = "BACKUP_JOB_TYPE_UNSPECIFIED"
	AWS_EC2_BACKUP BackupJobType = "AWS_EC2_BACKUP"
	AWS_EC2_SNAPSHOT_CONVERSION BackupJobType = "AWS_EC2_SNAPSHOT_CONVERSION"
	AWS_RDS_BACKUP BackupJobType = "AWS_RDS_BACKUP"
	AWS_S3_BACKUP BackupJobType = "AWS_S3_BACKUP"
	AWS_DYNAMO_DB_BACKUP BackupJobType = "AWS_DYNAMO_DB_BACKUP"
	AWS_DYNAMO_DB_BACKUP_FULL BackupJobType = "AWS_DYNAMO_DB_BACKUP_FULL"
	AWS_EKS_NAMESPACE_BACKUP BackupJobType = "AWS_EKS_NAMESPACE_BACKUP"
	AZURE_STORAGE_ACCOUNT_BACKUP BackupJobType = "AZURE_STORAGE_ACCOUNT_BACKUP"
	AZURE_VM_BACKUP BackupJobType = "AZURE_VM_BACKUP"
	AZURE_SQL_DATABASE_BACKUP BackupJobType = "AZURE_SQL_DATABASE_BACKUP"
	AZURE_SQL_MANAGED_INSTANCE_BACKUP BackupJobType = "AZURE_SQL_MANAGED_INSTANCE_BACKUP"
	AZURE_SQL_VIRTUAL_MACHINE_BACKUP BackupJobType = "AZURE_SQL_VIRTUAL_MACHINE_BACKUP"
	AZURE_SQL_VIRTUAL_MACHINE_BACKUP_FULL BackupJobType = "AZURE_SQL_VIRTUAL_MACHINE_BACKUP_FULL"
	AZURE_POSTGRESQL_BACKUP BackupJobType = "AZURE_POSTGRESQL_BACKUP"
	AZURE_MYSQL_BACKUP BackupJobType = "AZURE_MYSQL_BACKUP"
	GCP_COMPUTE_ENGINE_INSTANCE_BACKUP BackupJobType = "GCP_COMPUTE_ENGINE_INSTANCE_BACKUP"
	GCP_CLOUD_STORAGE_BACKUP BackupJobType = "GCP_CLOUD_STORAGE_BACKUP"
	GCP_CLOUD_SQL_BACKUP BackupJobType = "GCP_CLOUD_SQL_BACKUP"
	ATLAS_MONGODB_CLUSTER_BACKUP BackupJobType = "ATLAS_MONGODB_CLUSTER_BACKUP"
)

// All allowed values of BackupJobType enum
var AllowedBackupJobTypeEnumValues = []BackupJobType{
	"BACKUP_JOB_TYPE_UNSPECIFIED",
	"AWS_EC2_BACKUP",
	"AWS_EC2_SNAPSHOT_CONVERSION",
	"AWS_RDS_BACKUP",
	"AWS_S3_BACKUP",
	"AWS_DYNAMO_DB_BACKUP",
	"AWS_DYNAMO_DB_BACKUP_FULL",
	"AWS_EKS_NAMESPACE_BACKUP",
	"AZURE_STORAGE_ACCOUNT_BACKUP",
	"AZURE_VM_BACKUP",
	"AZURE_SQL_DATABASE_BACKUP",
	"AZURE_SQL_MANAGED_INSTANCE_BACKUP",
	"AZURE_SQL_VIRTUAL_MACHINE_BACKUP",
	"AZURE_SQL_VIRTUAL_MACHINE_BACKUP_FULL",
	"AZURE_POSTGRESQL_BACKUP",
	"AZURE_MYSQL_BACKUP",
	"GCP_COMPUTE_ENGINE_INSTANCE_BACKUP",
	"GCP_CLOUD_STORAGE_BACKUP",
	"GCP_CLOUD_SQL_BACKUP",
	"ATLAS_MONGODB_CLUSTER_BACKUP",
}

func (v *BackupJobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupJobType(value)
	for _, existing := range AllowedBackupJobTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupJobType", value)
}

// NewBackupJobTypeFromValue returns a pointer to a valid BackupJobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupJobTypeFromValue(v string) (*BackupJobType, error) {
	ev := BackupJobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupJobType: valid values are %v", v, AllowedBackupJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupJobType) IsValid() bool {
	for _, existing := range AllowedBackupJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupJobType value
func (v BackupJobType) Ptr() *BackupJobType {
	return &v
}

type NullableBackupJobType struct {
	value *BackupJobType
	isSet bool
}

func (v NullableBackupJobType) Get() *BackupJobType {
	return v.value
}

func (v *NullableBackupJobType) Set(val *BackupJobType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobType(val *BackupJobType) *NullableBackupJobType {
	return &NullableBackupJobType{value: val, isSet: true}
}

func (v NullableBackupJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

