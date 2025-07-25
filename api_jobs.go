/*
Eon API

The Eon.io REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eon

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// JobsAPIService JobsAPI service
type JobsAPIService service

type ApiGetBackupJobRequest struct {
	ctx context.Context
	ApiService *JobsAPIService
	jobId string
	projectId string
}

func (r ApiGetBackupJobRequest) Execute() (*GetBackupJobResponse, *http.Response, error) {
	return r.ApiService.GetBackupJobExecute(r)
}

/*
GetBackupJob Get Backup Job

Description: Retrieves a backup job by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId ID of the job to retrieve.
 @param projectId ID of the project the job is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
 @return ApiGetBackupJobRequest
*/
func (a *JobsAPIService) GetBackupJob(ctx context.Context, jobId string, projectId string) ApiGetBackupJobRequest {
	return ApiGetBackupJobRequest{
		ApiService: a,
		ctx: ctx,
		jobId: jobId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return GetBackupJobResponse
func (a *JobsAPIService) GetBackupJobExecute(r ApiGetBackupJobRequest) (*GetBackupJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBackupJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsAPIService.GetBackupJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/backup-jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRestoreJobRequest struct {
	ctx context.Context
	ApiService *JobsAPIService
	jobId string
	projectId string
}

func (r ApiGetRestoreJobRequest) Execute() (*GetRestoreJobResponse, *http.Response, error) {
	return r.ApiService.GetRestoreJobExecute(r)
}

/*
GetRestoreJob Get Restore Job

Description: Retrieves a restore job by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId ID of the job to retrieve.
 @param projectId ID of the project the job is in. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
 @return ApiGetRestoreJobRequest
*/
func (a *JobsAPIService) GetRestoreJob(ctx context.Context, jobId string, projectId string) ApiGetRestoreJobRequest {
	return ApiGetRestoreJobRequest{
		ApiService: a,
		ctx: ctx,
		jobId: jobId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return GetRestoreJobResponse
func (a *JobsAPIService) GetRestoreJobExecute(r ApiGetRestoreJobRequest) (*GetRestoreJobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetRestoreJobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsAPIService.GetRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/restore-jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListBackupJobsRequest struct {
	ctx context.Context
	ApiService *JobsAPIService
	projectId string
	pageToken *string
	pageSize *int32
	listBackupJobsRequest *ListBackupJobsRequest
}

// Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests. 
func (r ApiListBackupJobsRequest) PageToken(pageToken string) ApiListBackupJobsRequest {
	r.pageToken = &pageToken
	return r
}

// Maximum number of items to return in the response.
func (r ApiListBackupJobsRequest) PageSize(pageSize int32) ApiListBackupJobsRequest {
	r.pageSize = &pageSize
	return r
}

// Filter and sort options.
func (r ApiListBackupJobsRequest) ListBackupJobsRequest(listBackupJobsRequest ListBackupJobsRequest) ApiListBackupJobsRequest {
	r.listBackupJobsRequest = &listBackupJobsRequest
	return r
}

func (r ApiListBackupJobsRequest) Execute() (*ListBackupJobsResponse, *http.Response, error) {
	return r.ApiService.ListBackupJobsExecute(r)
}

/*
ListBackupJobs List Backup Jobs

Description: Retrieves a list of backup jobs for the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
 @return ApiListBackupJobsRequest
*/
func (a *JobsAPIService) ListBackupJobs(ctx context.Context, projectId string) ApiListBackupJobsRequest {
	return ApiListBackupJobsRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return ListBackupJobsResponse
func (a *JobsAPIService) ListBackupJobsExecute(r ApiListBackupJobsRequest) (*ListBackupJobsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListBackupJobsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsAPIService.ListBackupJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/backup-jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.listBackupJobsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListRestoreJobsRequest struct {
	ctx context.Context
	ApiService *JobsAPIService
	projectId string
	pageToken *string
	pageSize *int32
	listRestoreJobsRequest *ListRestoreJobsRequest
}

// Cursor that points to the first record of the next page of results. Get this value from the previous response. To preserve the results in the same order, use the same sorting and filters in the first request as all subsequent requests. 
func (r ApiListRestoreJobsRequest) PageToken(pageToken string) ApiListRestoreJobsRequest {
	r.pageToken = &pageToken
	return r
}

// Maximum number of items to return in the response.
func (r ApiListRestoreJobsRequest) PageSize(pageSize int32) ApiListRestoreJobsRequest {
	r.pageSize = &pageSize
	return r
}

// Filter and sort options. 
func (r ApiListRestoreJobsRequest) ListRestoreJobsRequest(listRestoreJobsRequest ListRestoreJobsRequest) ApiListRestoreJobsRequest {
	r.listRestoreJobsRequest = &listRestoreJobsRequest
	return r
}

func (r ApiListRestoreJobsRequest) Execute() (*ListRestoreJobsResponse, *http.Response, error) {
	return r.ApiService.ListRestoreJobsExecute(r)
}

/*
ListRestoreJobs List Restore Jobs

Description: Retrieves a list of restore jobs for the given project.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId ID of the project whose restore jobs you want to retrieve. You can get your project ID from the [API Credentials](/global-settings/api-credentials) page in your global settings. 
 @return ApiListRestoreJobsRequest
*/
func (a *JobsAPIService) ListRestoreJobs(ctx context.Context, projectId string) ApiListRestoreJobsRequest {
	return ApiListRestoreJobsRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return ListRestoreJobsResponse
func (a *JobsAPIService) ListRestoreJobsExecute(r ApiListRestoreJobsRequest) (*ListRestoreJobsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListRestoreJobsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsAPIService.ListRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/restore-jobs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.listRestoreJobsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
