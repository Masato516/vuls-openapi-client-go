/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// RoleApiService RoleApi service
type RoleApiService service

type ApiRoleDeleteRoleRequest struct {
	ctx context.Context
	ApiService *RoleApiService
	roleID int64
	authorization *string
}

// api key auth
func (r ApiRoleDeleteRoleRequest) Authorization(authorization string) ApiRoleDeleteRoleRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRoleDeleteRoleRequest) Execute() (*http.Response, error) {
	return r.ApiService.RoleDeleteRoleExecute(r)
}

/*
RoleDeleteRole deleteRole role

role delete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Role ID
 @return ApiRoleDeleteRoleRequest
*/
func (a *RoleApiService) RoleDeleteRole(ctx context.Context, roleID int64) ApiRoleDeleteRoleRequest {
	return ApiRoleDeleteRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// Execute executes the request
func (a *RoleApiService) RoleDeleteRoleExecute(r ApiRoleDeleteRoleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleApiService.RoleDeleteRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/role/{roleID}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleID"+"}", url.PathEscape(parameterToString(r.roleID, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRoleGetRoleDetailRequest struct {
	ctx context.Context
	ApiService *RoleApiService
	roleID int64
	authorization *string
}

// api key auth
func (r ApiRoleGetRoleDetailRequest) Authorization(authorization string) ApiRoleGetRoleDetailRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRoleGetRoleDetailRequest) Execute() (*RoleGetRoleDetailResponseBody, *http.Response, error) {
	return r.ApiService.RoleGetRoleDetailExecute(r)
}

/*
RoleGetRoleDetail getRoleDetail role

role detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Role ID
 @return ApiRoleGetRoleDetailRequest
*/
func (a *RoleApiService) RoleGetRoleDetail(ctx context.Context, roleID int64) ApiRoleGetRoleDetailRequest {
	return ApiRoleGetRoleDetailRequest{
		ApiService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// Execute executes the request
//  @return RoleGetRoleDetailResponseBody
func (a *RoleApiService) RoleGetRoleDetailExecute(r ApiRoleGetRoleDetailRequest) (*RoleGetRoleDetailResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleGetRoleDetailResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleApiService.RoleGetRoleDetail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/role/{roleID}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleID"+"}", url.PathEscape(parameterToString(r.roleID, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml", "application/gob"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiRoleGetRoleListRequest struct {
	ctx context.Context
	ApiService *RoleApiService
	page *int32
	limit *int32
	offset *int32
	filterCveID *string
	authorization *string
}

// Page Number (default: 1)
func (r ApiRoleGetRoleListRequest) Page(page int32) ApiRoleGetRoleListRequest {
	r.page = &page
	return r
}

// Limit (default: 20, max: 100)
func (r ApiRoleGetRoleListRequest) Limit(limit int32) ApiRoleGetRoleListRequest {
	r.limit = &limit
	return r
}

// Offset
func (r ApiRoleGetRoleListRequest) Offset(offset int32) ApiRoleGetRoleListRequest {
	r.offset = &offset
	return r
}

// CveID filter
func (r ApiRoleGetRoleListRequest) FilterCveID(filterCveID string) ApiRoleGetRoleListRequest {
	r.filterCveID = &filterCveID
	return r
}

// api key auth
func (r ApiRoleGetRoleListRequest) Authorization(authorization string) ApiRoleGetRoleListRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRoleGetRoleListRequest) Execute() (*RoleGetRoleListResponseBody, *http.Response, error) {
	return r.ApiService.RoleGetRoleListExecute(r)
}

/*
RoleGetRoleList getRoleList role

role list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRoleGetRoleListRequest
*/
func (a *RoleApiService) RoleGetRoleList(ctx context.Context) ApiRoleGetRoleListRequest {
	return ApiRoleGetRoleListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RoleGetRoleListResponseBody
func (a *RoleApiService) RoleGetRoleListExecute(r ApiRoleGetRoleListRequest) (*RoleGetRoleListResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleGetRoleListResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleApiService.RoleGetRoleList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.filterCveID != nil {
		localVarQueryParams.Add("filterCveID", parameterToString(*r.filterCveID, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml", "application/gob"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiRoleUpdateRoleRequest struct {
	ctx context.Context
	ApiService *RoleApiService
	roleID int64
	updateRoleRequestBody *RoleUpdateRoleRequestBody
	authorization *string
}

func (r ApiRoleUpdateRoleRequest) UpdateRoleRequestBody(updateRoleRequestBody RoleUpdateRoleRequestBody) ApiRoleUpdateRoleRequest {
	r.updateRoleRequestBody = &updateRoleRequestBody
	return r
}

// api key auth
func (r ApiRoleUpdateRoleRequest) Authorization(authorization string) ApiRoleUpdateRoleRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRoleUpdateRoleRequest) Execute() (*RoleUpdateRoleResponseBody, *http.Response, error) {
	return r.ApiService.RoleUpdateRoleExecute(r)
}

/*
RoleUpdateRole updateRole role

update role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleID Role ID
 @return ApiRoleUpdateRoleRequest
*/
func (a *RoleApiService) RoleUpdateRole(ctx context.Context, roleID int64) ApiRoleUpdateRoleRequest {
	return ApiRoleUpdateRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleID: roleID,
	}
}

// Execute executes the request
//  @return RoleUpdateRoleResponseBody
func (a *RoleApiService) RoleUpdateRoleExecute(r ApiRoleUpdateRoleRequest) (*RoleUpdateRoleResponseBody, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleUpdateRoleResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleApiService.RoleUpdateRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/role/{roleID}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleID"+"}", url.PathEscape(parameterToString(r.roleID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRoleRequestBody == nil {
		return localVarReturnValue, nil, reportError("updateRoleRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml", "application/gob"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml", "application/gob"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	// body params
	localVarPostBody = r.updateRoleRequestBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header_Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
