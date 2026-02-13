# \IamAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdpGroup**](IamAPI.md#CreateIdpGroup) | **Post** /v1/idp-groups | Create IDP Group
[**CreateRole**](IamAPI.md#CreateRole) | **Post** /v1/roles | Create Role
[**DeleteIdpGroup**](IamAPI.md#DeleteIdpGroup) | **Delete** /v1/idp-groups/{groupId} | Delete IDP Group
[**DeleteRole**](IamAPI.md#DeleteRole) | **Delete** /v1/roles/{roleId} | Delete Role
[**GetIdpGroup**](IamAPI.md#GetIdpGroup) | **Get** /v1/idp-groups/{groupId} | Get IDP Group
[**GetRole**](IamAPI.md#GetRole) | **Get** /v1/roles/{roleId} | Get Role
[**ListIdpGroups**](IamAPI.md#ListIdpGroups) | **Post** /v1/idp-groups/list | List IDP Groups
[**ListIdps**](IamAPI.md#ListIdps) | **Post** /v1/idps/list | List Identity Providers
[**ListPermissions**](IamAPI.md#ListPermissions) | **Get** /v1/permissions | List Permissions
[**ListRoles**](IamAPI.md#ListRoles) | **Post** /v1/roles/list | List Roles
[**UpdateIdpGroup**](IamAPI.md#UpdateIdpGroup) | **Put** /v1/idp-groups/{groupId} | Update IDP Group
[**UpdateRole**](IamAPI.md#UpdateRole) | **Put** /v1/roles/{roleId} | Update Role



## CreateIdpGroup

> CreateIdpGroupResponse CreateIdpGroup(ctx).CreateIdpGroupRequest(createIdpGroupRequest).Execute()

Create IDP Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	createIdpGroupRequest := *openapiclient.NewCreateIdpGroupRequest("8f14e45f-ceea-467f-a8dc-e5e47d61e5d7", "engineering-team", []string{"RoleIds_example"}) // CreateIdpGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.CreateIdpGroup(context.Background()).CreateIdpGroupRequest(createIdpGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.CreateIdpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdpGroup`: CreateIdpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.CreateIdpGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdpGroupRequest** | [**CreateIdpGroupRequest**](CreateIdpGroupRequest.md) |  | 

### Return type

[**CreateIdpGroupResponse**](CreateIdpGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> CreateRoleResponse CreateRole(ctx).CreateRoleRequest(createRoleRequest).Execute()

Create Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	createRoleRequest := *openapiclient.NewCreateRoleRequest("Name_example", []openapiclient.PermissionGrantInput{*openapiclient.NewPermissionGrantInput(openapiclient.PermissionType("permission_id_unspecified"))}) // CreateRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.CreateRole(context.Background()).CreateRoleRequest(createRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: CreateRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleRequest** | [**CreateRoleRequest**](CreateRoleRequest.md) |  | 

### Return type

[**CreateRoleResponse**](CreateRoleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdpGroup

> DeleteIdpGroup(ctx, groupId).Execute()

Delete IDP Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	groupId := "ceaf2281-bf04-542d-a801-a6a4865373ad" // string | IDP Group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteIdpGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteIdpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | IDP Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Delete Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	roleId := "ceaf2281-bf04-542d-a801-a6a4865373ad" // string | Role ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IamAPI.DeleteRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpGroup

> GetIdpGroupResponse GetIdpGroup(ctx, groupId).Execute()

Get IDP Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	groupId := "daa42222-6ad2-5d37-bb26-4f2a9cba73da" // string | IDP Group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetIdpGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetIdpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdpGroup`: GetIdpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetIdpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | IDP Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIdpGroupResponse**](GetIdpGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> GetRoleResponse GetRole(ctx, roleId).Execute()

Get Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	roleId := "daa42222-6ad2-5d37-bb26-4f2a9cba73da" // string | Role ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.GetRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: GetRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRoleResponse**](GetRoleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdpGroups

> ListIdpGroupsResponse ListIdpGroups(ctx).PageToken(pageToken).PageSize(pageSize).Execute()

List IDP Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListIdpGroups(context.Background()).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListIdpGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdpGroups`: ListIdpGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListIdpGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdpGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]

### Return type

[**ListIdpGroupsResponse**](ListIdpGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdps

> ListIdpsResponse ListIdps(ctx).PageToken(pageToken).PageSize(pageSize).Execute()

List Identity Providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListIdps(context.Background()).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListIdps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdps`: ListIdpsResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListIdps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]

### Return type

[**ListIdpsResponse**](ListIdpsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> ListPermissionsResponse ListPermissions(ctx).Execute()

List Permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissions`: ListPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


### Return type

[**ListPermissionsResponse**](ListPermissionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> ListRolesResponse ListRoles(ctx).PageToken(pageToken).PageSize(pageSize).Execute()

List Roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	pageToken := "Yjk3ODZjNjktZTIwZC00NjAxLWE1MzktZjg2NGExM2IxYTZlfDE=" // string | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  (optional)
	pageSize := int32(10) // int32 | Maximum number of items to return in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.ListRoles(context.Background()).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: ListRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** | Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests.  | 
 **pageSize** | **int32** | Maximum number of items to return in the response. | [default to 50]

### Return type

[**ListRolesResponse**](ListRolesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpGroup

> UpdateIdpGroupResponse UpdateIdpGroup(ctx, groupId).UpdateIdpGroupRequest(updateIdpGroupRequest).Execute()

Update IDP Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | IDP Group ID.
	updateIdpGroupRequest := *openapiclient.NewUpdateIdpGroupRequest([]string{"RoleIds_example"}) // UpdateIdpGroupRequest | The request body for updating an IDP group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateIdpGroup(context.Background(), groupId).UpdateIdpGroupRequest(updateIdpGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateIdpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdpGroup`: UpdateIdpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateIdpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | IDP Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdpGroupRequest** | [**UpdateIdpGroupRequest**](UpdateIdpGroupRequest.md) | The request body for updating an IDP group | 

### Return type

[**UpdateIdpGroupResponse**](UpdateIdpGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRoleResponse UpdateRole(ctx, roleId).UpdateRoleRequest(updateRoleRequest).Execute()

Update Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon.io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the role
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest("Name_example", []openapiclient.PermissionGrantInput{*openapiclient.NewPermissionGrantInput(openapiclient.PermissionType("permission_id_unspecified"))}) // UpdateRoleRequest | The request body for updating a user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamAPI.UpdateRole(context.Background(), roleId).UpdateRoleRequest(updateRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRole`: UpdateRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `IamAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | ID of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) | The request body for updating a user | 

### Return type

[**UpdateRoleResponse**](UpdateRoleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

