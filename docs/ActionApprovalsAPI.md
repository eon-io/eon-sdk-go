# \ActionApprovalsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelActionApprovalRequest**](ActionApprovalsAPI.md#CancelActionApprovalRequest) | **Post** /v1/projects/{projectId}/action-approvals/my-requests/{requestId}/cancel | Cancel My Action Request
[**CreateActionApprovalRequest**](ActionApprovalsAPI.md#CreateActionApprovalRequest) | **Post** /v1/projects/{projectId}/action-approvals/my-requests/{requestId}/submit | Submit My Action Request
[**CreateActionApprovalRule**](ActionApprovalsAPI.md#CreateActionApprovalRule) | **Post** /v1/projects/{projectId}/action-approvals/rules | Create Action Approval Rule
[**DeleteActionApprovalRule**](ActionApprovalsAPI.md#DeleteActionApprovalRule) | **Delete** /v1/projects/{projectId}/action-approvals/rules/{actionApprovalRuleId} | Delete Action Approval Rule
[**GetActionApprovalRule**](ActionApprovalsAPI.md#GetActionApprovalRule) | **Get** /v1/projects/{projectId}/action-approvals/rules/{actionApprovalRuleId} | Get Action Approval Rule
[**GetMyActionApprovalRequest**](ActionApprovalsAPI.md#GetMyActionApprovalRequest) | **Get** /v1/projects/{projectId}/action-approvals/my-requests/{requestId} | Get My Action Request
[**ListActionApprovalRules**](ActionApprovalsAPI.md#ListActionApprovalRules) | **Get** /v1/projects/{projectId}/action-approvals/rules | List Action Approval Rules
[**UpdateActionApprovalRule**](ActionApprovalsAPI.md#UpdateActionApprovalRule) | **Put** /v1/projects/{projectId}/action-approvals/rules/{actionApprovalRuleId} | Update Action Approval Rule



## CancelActionApprovalRequest

> CancelMyMPARequestResponse CancelActionApprovalRequest(ctx, projectId, requestId).Execute()

Cancel My Action Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "8ea94e63-894f-5779-aa8f-06ad93121abb" // string | ID of the project.
	requestId := "0fc514a3-61fe-5d39-a079-9d0df1820956" // string | ID of the action request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.CancelActionApprovalRequest(context.Background(), projectId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.CancelActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActionApprovalRequest`: CancelMyMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.CancelActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**requestId** | **string** | ID of the action request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CancelMyMPARequestResponse**](CancelMyMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActionApprovalRequest

> SubmitMyMPARequestResponse CreateActionApprovalRequest(ctx, projectId, requestId).SubmitMyMPARequestRequest(submitMyMPARequestRequest).Execute()

Submit My Action Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "8ea94e63-894f-5779-aa8f-06ad93121abb" // string | ID of the project.
	requestId := "0fc514a3-61fe-5d39-a079-9d0df1820956" // string | ID of the action request.
	submitMyMPARequestRequest := *openapiclient.NewSubmitMyMPARequestRequest(openapiclient.MPASubmitAction("MPA_SUBMIT_ACTION_UNSPECIFIED")) // SubmitMyMPARequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.CreateActionApprovalRequest(context.Background(), projectId, requestId).SubmitMyMPARequestRequest(submitMyMPARequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.CreateActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActionApprovalRequest`: SubmitMyMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.CreateActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**requestId** | **string** | ID of the action request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **submitMyMPARequestRequest** | [**SubmitMyMPARequestRequest**](SubmitMyMPARequestRequest.md) |  | 

### Return type

[**SubmitMyMPARequestResponse**](SubmitMyMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActionApprovalRule

> CreateActionApprovalRuleResponse CreateActionApprovalRule(ctx, projectId).CreateActionApprovalRuleRequest(createActionApprovalRuleRequest).Execute()

Create Action Approval Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project in which you want to create an action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	createActionApprovalRuleRequest := *openapiclient.NewCreateActionApprovalRuleRequest(openapiclient.ActionApprovalOperationType("ADD_RESTORE_ACCOUNT"), int32(123), int32(123)) // CreateActionApprovalRuleRequest | Action approval rule to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.CreateActionApprovalRule(context.Background(), projectId).CreateActionApprovalRuleRequest(createActionApprovalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.CreateActionApprovalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActionApprovalRule`: CreateActionApprovalRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.CreateActionApprovalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project in which you want to create an action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionApprovalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createActionApprovalRuleRequest** | [**CreateActionApprovalRuleRequest**](CreateActionApprovalRuleRequest.md) | Action approval rule to create. | 

### Return type

[**CreateActionApprovalRuleResponse**](CreateActionApprovalRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActionApprovalRule

> DeleteActionApprovalRule(ctx, projectId, actionApprovalRuleId).Execute()

Delete Action Approval Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	actionApprovalRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action approval rule ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionApprovalsAPI.DeleteActionApprovalRule(context.Background(), projectId, actionApprovalRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.DeleteActionApprovalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**actionApprovalRuleId** | **string** | Action approval rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActionApprovalRuleRequest struct via the builder pattern


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


## GetActionApprovalRule

> GetActionApprovalRuleResponse GetActionApprovalRule(ctx, projectId, actionApprovalRuleId).Execute()

Get Action Approval Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	actionApprovalRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action approval rule ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.GetActionApprovalRule(context.Background(), projectId, actionApprovalRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.GetActionApprovalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionApprovalRule`: GetActionApprovalRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.GetActionApprovalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**actionApprovalRuleId** | **string** | Action approval rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionApprovalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetActionApprovalRuleResponse**](GetActionApprovalRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyActionApprovalRequest

> GetMPARequestResponse GetMyActionApprovalRequest(ctx, projectId, requestId).Execute()

Get My Action Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "8ea94e63-894f-5779-aa8f-06ad93121abb" // string | ID of the project.
	requestId := "0fc514a3-61fe-5d39-a079-9d0df1820956" // string | ID of the action request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.GetMyActionApprovalRequest(context.Background(), projectId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.GetMyActionApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyActionApprovalRequest`: GetMPARequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.GetMyActionApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project. | 
**requestId** | **string** | ID of the action request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyActionApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetMPARequestResponse**](GetMPARequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionApprovalRules

> ListActionApprovalRulesResponse ListActionApprovalRules(ctx, projectId).Operation(operation).PageToken(pageToken).PageSize(pageSize).Execute()

List Action Approval Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project whose action approval rules you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	operation := openapiclient.ActionApprovalOperationType("ADD_RESTORE_ACCOUNT") // ActionApprovalOperationType | Filter rules by the action they protect. (optional)
	pageToken := "pageToken_example" // string | Cursor that points to the first rule on the next page. Get this value from the previous response.  (optional)
	pageSize := int32(10) // int32 | Maximum number of rules to return. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.ListActionApprovalRules(context.Background(), projectId).Operation(operation).PageToken(pageToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.ListActionApprovalRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActionApprovalRules`: ListActionApprovalRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.ListActionApprovalRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project whose action approval rules you want to retrieve. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActionApprovalRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operation** | [**ActionApprovalOperationType**](ActionApprovalOperationType.md) | Filter rules by the action they protect. | 
 **pageToken** | **string** | Cursor that points to the first rule on the next page. Get this value from the previous response.  | 
 **pageSize** | **int32** | Maximum number of rules to return. | [default to 50]

### Return type

[**ListActionApprovalRulesResponse**](ListActionApprovalRulesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActionApprovalRule

> UpdateActionApprovalRuleResponse UpdateActionApprovalRule(ctx, projectId, actionApprovalRuleId).UpdateActionApprovalRuleRequest(updateActionApprovalRuleRequest).Execute()

Update Action Approval Rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/eon-io/eon-service/services/frontend/api-gateway/sdk/external-go"
)

func main() {
	projectId := "733888d8-2573-5f9a-b81d-21f051d24fda" // string | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page. 
	actionApprovalRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Action approval rule ID.
	updateActionApprovalRuleRequest := *openapiclient.NewUpdateActionApprovalRuleRequest() // UpdateActionApprovalRuleRequest | Fields to update on the action approval rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionApprovalsAPI.UpdateActionApprovalRule(context.Background(), projectId, actionApprovalRuleId).UpdateActionApprovalRuleRequest(updateActionApprovalRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionApprovalsAPI.UpdateActionApprovalRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActionApprovalRule`: UpdateActionApprovalRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ActionApprovalsAPI.UpdateActionApprovalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project that contains the action approval rule. You can get your project ID from the [API Credentials](https://console.eon.io/global-management/api-credentials) page.  | 
**actionApprovalRuleId** | **string** | Action approval rule ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActionApprovalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateActionApprovalRuleRequest** | [**UpdateActionApprovalRuleRequest**](UpdateActionApprovalRuleRequest.md) | Fields to update on the action approval rule. | 

### Return type

[**UpdateActionApprovalRuleResponse**](UpdateActionApprovalRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

