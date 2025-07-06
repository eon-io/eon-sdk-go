# Eon SDK for Go

A Go client library for the Eon.io REST API - The comprehensive cloud backup and restore service.

## Installation

```bash
go get github.com/eon-io/eon-sdk-go
```

## Usage

### Basic Client Setup

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    eon "github.com/eon-io/eon-sdk-go"
)

func main() {
    // Create a new configuration
    config := eon.NewConfiguration()
    config.Servers[0].URL = "https://your-eon-instance.com/api"
    
    // Create API client
    client := eon.NewAPIClient(config)
    
    // Set up authentication (Bearer token)
    ctx := context.WithValue(context.Background(), eon.ContextAccessToken, "your-access-token")
    
    // Example: List snapshots
    projectId := "your-project-id"
    resp, httpResp, err := client.SnapshotsAPI.ListSnapshots(ctx, projectId).Execute()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d snapshots\n", len(resp.GetSnapshots()))
}
```

### Authentication

The SDK supports OAuth 2.0 client credentials flow:

```go
// Get access token using client credentials
authResp, _, err := client.AuthAPI.GetAccessToken(ctx).
    ApiCredentials(eon.ApiCredentials{
        ClientId:     "your-client-id",
        ClientSecret: "your-client-secret",
    }).Execute()

if err != nil {
    log.Fatal(err)
}

// Use the token for subsequent requests
ctx = context.WithValue(ctx, eon.ContextAccessToken, authResp.GetAccessToken())
```

### Available APIs

The SDK provides access to all Eon APIs:

- **AuthAPI**: Authentication and token management
- **SnapshotsAPI**: Snapshot management and restore operations
- **AccountsAPI**: Source and restore account management
- **ResourcesAPI**: Resource inventory and management
- **JobsAPI**: Backup and restore job monitoring
- **DashboardAPI**: Dashboard and analytics data

### Examples

#### List Source Accounts

```go
resp, _, err := client.AccountsAPI.ListSourceAccounts(ctx, projectId).Execute()
if err != nil {
    log.Fatal(err)
}

for _, account := range resp.GetAccounts() {
    fmt.Printf("Account: %s (%s)\n", account.GetName(), account.GetProviderAccountId())
}
```

#### Start a Volume Restore

```go
restoreReq := eon.RestoreVolumeToEbsRequest{
    ProviderVolumeId: "vol-12345678",
    RestoreAccountId: "restore-account-id",
    Destination: eon.EbsRestoreDestination{
        AwsEbs: &eon.EbsTarget{
            AvailabilityZone: "us-east-1a",
            VolumeSettings: eon.VolumeSettings{
                Type:      "gp3",
                SizeBytes: 100 * 1024 * 1024 * 1024, // 100 GB
            },
        },
    },
}

resp, _, err := client.SnapshotsAPI.RestoreEbsVolume(ctx, projectId, resourceId, snapshotId).
    RestoreVolumeToEbsRequest(restoreReq).Execute()

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Restore job started: %s\n", resp.GetJobId())
```

#### Monitor Job Status

```go
jobResp, _, err := client.JobsAPI.GetRestoreJob(ctx, jobId, projectId).Execute()
if err != nil {
    log.Fatal(err)
}

job := jobResp.GetJob()
fmt.Printf("Job Status: %s\n", job.GetJobExecutionDetails().GetStatus())
```

## API Reference

### Core Types

- `Snapshot`: Represents a backup snapshot
- `RestoreJob`: Represents a restore operation
- `SourceAccount`: Source account for backups
- `RestoreAccount`: Target account for restores
- `InventoryResource`: Resource in the inventory

### Error Handling

The SDK uses standard Go error handling patterns:

```go
resp, httpResp, err := client.SomeAPI.SomeMethod(ctx).Execute()
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*eon.GenericOpenAPIError); ok {
        fmt.Printf("API Error: %s\n", apiErr.Error())
        fmt.Printf("Status Code: %d\n", httpResp.StatusCode)
    }
    return err
}
```

## Configuration

### Custom HTTP Client

```go
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

config := eon.NewConfiguration()
config.HTTPClient = httpClient
```

### Custom Base URL

```go
config := eon.NewConfiguration()
config.Servers[0].URL = "https://custom-eon-instance.com/api"
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This SDK is released under the same license as the Eon service.

## Support

For support and questions:
- Create an issue in this repository
- Contact Eon support through your service portal
- Check the official Eon documentation

## Version Compatibility

This SDK is compatible with Eon API version 1.0.0 and later.



