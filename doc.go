/*
Package eon provides a Go client library for the Eon.io REST API.

Eon is a comprehensive cloud backup and restore service that supports multiple cloud providers
including AWS, Azure, and GCP.

# Quick Start

	import "github.com/eon-io/eon-sdk-go"

	// Create a client
	config := eon.NewConfiguration()
	config.Servers[0].URL = "https://your-eon-instance.com/api"
	client := eon.NewAPIClient(config)

	// Authenticate
	ctx := context.WithValue(context.Background(), eon.ContextAccessToken, "your-token")

	// List snapshots
	resp, _, err := client.SnapshotsAPI.ListSnapshots(ctx, "project-id").Execute()

# Authentication

The SDK supports OAuth 2.0 client credentials flow for authentication:

	authResp, _, err := client.AuthAPI.GetAccessToken(ctx).
		ApiCredentials(eon.ApiCredentials{
			ClientId:  "your-client-id",
			ClientSecret: "your-client-secret",
		}).Execute()

	// Use the token for subsequent requests
	ctx = context.WithValue(ctx, eon.ContextAccessToken, authResp.GetAccessToken())

# Available APIs

The SDK provides the following API interfaces:

  - AuthAPI: Authentication and token management
  - SnapshotsAPI: Snapshot management and restore operations
  - AccountsAPI: Source and restore account management
  - ResourcesAPI: Resource inventory and management
  - JobsAPI: Backup and restore job monitoring
  - DashboardAPI: Dashboard and analytics data

# Error Handling

All API methods return a standard Go error. API-specific errors can be checked:

	resp, httpResp, err := client.SomeAPI.SomeMethod(ctx).Execute()
	if err != nil {
		if apiErr, ok := err.(*eon.GenericOpenAPIError); ok {
			fmt.Printf("API Error: %s, Status: %d\n", apiErr.Error(), httpResp.StatusCode)
		}
	}

# Configuration

The client can be configured with custom HTTP clients, timeouts, and base URLs:

	config := eon.NewConfiguration()
	config.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	config.Servers[0].URL = "https://custom-eon-instance.com/api"
	client := eon.NewAPIClient(config)

For more examples and detailed documentation, see the README.md file.
*/
package eon
