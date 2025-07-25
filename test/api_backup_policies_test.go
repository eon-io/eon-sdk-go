/*
Eon API

Testing BackupPoliciesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package eon

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func Test_eon_BackupPoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BackupPoliciesAPIService CreateBackupPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.BackupPoliciesAPI.CreateBackupPolicy(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPoliciesAPIService DeleteBackupPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var backupPolicyId string
		var projectId string

		httpRes, err := apiClient.BackupPoliciesAPI.DeleteBackupPolicy(context.Background(), backupPolicyId, projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPoliciesAPIService GetBackupPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var backupPolicyId string
		var projectId string

		resp, httpRes, err := apiClient.BackupPoliciesAPI.GetBackupPolicy(context.Background(), backupPolicyId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPoliciesAPIService ListBackupPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.BackupPoliciesAPI.ListBackupPolicies(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPoliciesAPIService UpdateBackupPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var backupPolicyId string
		var projectId string

		resp, httpRes, err := apiClient.BackupPoliciesAPI.UpdateBackupPolicy(context.Background(), backupPolicyId, projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
