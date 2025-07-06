package main

import (
	"context"
	"fmt"
	"log"
	"os"

	eon "github.com/eon-io/eon-sdk-go"
)

func main() {
	// Configuration
	endpoint := os.Getenv("EON_ENDPOINT")
	clientID := os.Getenv("EON_CLIENT_ID")
	clientSecret := os.Getenv("EON_CLIENT_SECRET")
	projectID := os.Getenv("EON_PROJECT_ID")

	if endpoint == "" || clientID == "" || clientSecret == "" || projectID == "" {
		log.Fatal("Please set EON_ENDPOINT, EON_CLIENT_ID, EON_CLIENT_SECRET, and EON_PROJECT_ID environment variables")
	}

	// Create client configuration
	config := eon.NewConfiguration()
	config.Servers[0].URL = fmt.Sprintf("%s/api", endpoint)

	// Create API client
	client := eon.NewAPIClient(config)
	ctx := context.Background()

	// Authenticate using OAuth 2.0 client credentials
	fmt.Println("Authenticating...")
	authResp, httpResp, err := client.AuthAPI.GetAccessToken(ctx).
		ApiCredentials(eon.ApiCredentials{
			ClientId:     clientID,
			ClientSecret: clientSecret,
		}).Execute()

	if err != nil {
		log.Fatalf("Authentication failed: %v (Status: %d)", err, httpResp.StatusCode)
	}

	// Set up authenticated context
	authCtx := context.WithValue(ctx, eon.ContextAccessToken, authResp.GetAccessToken())
	fmt.Printf("Authentication successful! Token expires in %d seconds\n", authResp.GetExpirationSeconds())

	// Example 1: List source accounts
	fmt.Println("\n=== Listing Source Accounts ===")
	sourceResp, _, err := client.AccountsAPI.ListSourceAccounts(authCtx, projectID).
		ListSourceAccountsRequest(eon.ListSourceAccountsRequest{}).Execute()
	if err != nil {
		log.Printf("Error listing source accounts: %v", err)
	} else {
		accounts := sourceResp.GetAccounts()
		fmt.Printf("Found %d source accounts:\n", len(accounts))
		for _, account := range accounts {
			fmt.Printf("  - %s (%s) - Status: %s\n", 
				account.GetName(), 
				account.GetProviderAccountId(), 
				account.GetStatus())
		}
	}

	// Example 2: List restore accounts
	fmt.Println("\n=== Listing Restore Accounts ===")
	restoreResp, _, err := client.AccountsAPI.ListRestoreAccounts(authCtx, projectID).
		ListRestoreAccountsRequest(eon.ListRestoreAccountsRequest{}).Execute()
	if err != nil {
		log.Printf("Error listing restore accounts: %v", err)
	} else {
		accounts := restoreResp.GetAccounts()
		fmt.Printf("Found %d restore accounts:\n", len(accounts))
		for _, account := range accounts {
			fmt.Printf("  - %s - Status: %s\n", 
				account.GetProviderAccountId(), 
				account.GetStatus())
		}
	}

	// Example 3: List resources (limited to first 10)
	fmt.Println("\n=== Listing Resources (first 10) ===")
	pageSize := int32(10)
	resourcesResp, _, err := client.ResourcesAPI.ListResources(authCtx, projectID).
		ListInventoryRequest(eon.ListInventoryRequest{}).Execute()
	if err != nil {
		log.Printf("Error listing resources: %v", err)
	} else {
		resources := resourcesResp.GetResources()
		fmt.Printf("Found %d resources (showing first %d):\n", len(resources), pageSize)
		for _, resource := range resources {
			fmt.Printf("  - %s (%s) - Type: %s\n", 
				resource.GetResourceName(), 
				resource.GetId(), 
				resource.GetResourceType())
		}
	}

	// Example 4: Get snapshots if resource ID is provided
	resourceID := os.Getenv("EON_RESOURCE_ID")
	if resourceID != "" {
		fmt.Printf("\n=== Getting Snapshots for Resource %s ===\n", resourceID)
		snapshotsResp, _, err := client.SnapshotsAPI.ListResourceSnapshots(authCtx, projectID, resourceID).
			ListInventorySnapshotsRequest(eon.ListInventorySnapshotsRequest{}).Execute()
		if err != nil {
			log.Printf("Error listing snapshots: %v", err)
		} else {
			snapshots := snapshotsResp.GetSnapshots()
			fmt.Printf("Found %d snapshots:\n", len(snapshots))
			for _, snapshot := range snapshots {
				fmt.Printf("  - %s (created: %s)\n", 
					snapshot.GetId(), 
					snapshot.GetCreatedTime().Format("2006-01-02 15:04:05"))
			}
		}
	}

	fmt.Println("\n=== Example completed successfully! ===")
} 