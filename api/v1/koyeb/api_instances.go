/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type InstancesApi interface {

	/*
	ExecCommand Exec Command

	This endpoint opens a websocket. Once open, all frames going through the websocket should be formatted in JSON. Input frames should match the format specified below. Output frames will match the response schema.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExecCommandRequest
	*/
	ExecCommand(ctx context.Context) ApiExecCommandRequest

	// ExecCommandExecute executes the request
	//  @return StreamResultOfExecCommandReply
	ExecCommandExecute(r ApiExecCommandRequest) (*StreamResultOfExecCommandReply, *http.Response, error)

	/*
	GetInstance Get Instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the instance
	@return ApiGetInstanceRequest
	*/
	GetInstance(ctx context.Context, id string) ApiGetInstanceRequest

	// GetInstanceExecute executes the request
	//  @return GetInstanceReply
	GetInstanceExecute(r ApiGetInstanceRequest) (*GetInstanceReply, *http.Response, error)

	/*
	ListInstanceEvents List Instance events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListInstanceEventsRequest
	*/
	ListInstanceEvents(ctx context.Context) ApiListInstanceEventsRequest

	// ListInstanceEventsExecute executes the request
	//  @return ListInstanceEventsReply
	ListInstanceEventsExecute(r ApiListInstanceEventsRequest) (*ListInstanceEventsReply, *http.Response, error)

	/*
	ListInstances List Instances

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListInstancesRequest
	*/
	ListInstances(ctx context.Context) ApiListInstancesRequest

	// ListInstancesExecute executes the request
	//  @return ListInstancesReply
	ListInstancesExecute(r ApiListInstancesRequest) (*ListInstancesReply, *http.Response, error)
}

// InstancesApiService InstancesApi service
type InstancesApiService service

type ApiExecCommandRequest struct {
	ctx context.Context
	ApiService InstancesApi
	id *string
	bodyCommand *[]string
	bodyTtySizeHeight *int32
	bodyTtySizeWidth *int32
	bodyStdinData *string
	idType *string
}

// ID of the resource to exec on.
func (r ApiExecCommandRequest) Id(id string) ApiExecCommandRequest {
	r.id = &id
	return r
}

// Command to exec. Mandatory in the first frame sent
func (r ApiExecCommandRequest) BodyCommand(bodyCommand []string) ApiExecCommandRequest {
	r.bodyCommand = &bodyCommand
	return r
}

func (r ApiExecCommandRequest) BodyTtySizeHeight(bodyTtySizeHeight int32) ApiExecCommandRequest {
	r.bodyTtySizeHeight = &bodyTtySizeHeight
	return r
}

func (r ApiExecCommandRequest) BodyTtySizeWidth(bodyTtySizeWidth int32) ApiExecCommandRequest {
	r.bodyTtySizeWidth = &bodyTtySizeWidth
	return r
}

// Data is base64 encoded
func (r ApiExecCommandRequest) BodyStdinData(bodyStdinData string) ApiExecCommandRequest {
	r.bodyStdinData = &bodyStdinData
	return r
}

// When specified, it is used to determine if the kind of resource the id refers to. If missing, defaults to the instance id.
func (r ApiExecCommandRequest) IdType(idType string) ApiExecCommandRequest {
	r.idType = &idType
	return r
}

func (r ApiExecCommandRequest) Execute() (*StreamResultOfExecCommandReply, *http.Response, error) {
	return r.ApiService.ExecCommandExecute(r)
}

/*
ExecCommand Exec Command

This endpoint opens a websocket. Once open, all frames going through the websocket should be formatted in JSON. Input frames should match the format specified below. Output frames will match the response schema.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExecCommandRequest
*/
func (a *InstancesApiService) ExecCommand(ctx context.Context) ApiExecCommandRequest {
	return ApiExecCommandRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StreamResultOfExecCommandReply
func (a *InstancesApiService) ExecCommandExecute(r ApiExecCommandRequest) (*StreamResultOfExecCommandReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StreamResultOfExecCommandReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ExecCommand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/streams/instances/exec"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.bodyCommand != nil {
		t := *r.bodyCommand
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("body.command", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("body.command", parameterToString(t, "multi"))
		}
	}
	if r.bodyTtySizeHeight != nil {
		localVarQueryParams.Add("body.tty_size.height", parameterToString(*r.bodyTtySizeHeight, ""))
	}
	if r.bodyTtySizeWidth != nil {
		localVarQueryParams.Add("body.tty_size.width", parameterToString(*r.bodyTtySizeWidth, ""))
	}
	if r.bodyStdinData != nil {
		localVarQueryParams.Add("body.stdin.data", parameterToString(*r.bodyStdinData, ""))
	}
	if r.idType != nil {
		localVarQueryParams.Add("id_type", parameterToString(*r.idType, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiGetInstanceRequest struct {
	ctx context.Context
	ApiService InstancesApi
	id string
}

func (r ApiGetInstanceRequest) Execute() (*GetInstanceReply, *http.Response, error) {
	return r.ApiService.GetInstanceExecute(r)
}

/*
GetInstance Get Instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the instance
 @return ApiGetInstanceRequest
*/
func (a *InstancesApiService) GetInstance(ctx context.Context, id string) ApiGetInstanceRequest {
	return ApiGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetInstanceReply
func (a *InstancesApiService) GetInstanceExecute(r ApiGetInstanceRequest) (*GetInstanceReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInstanceReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.GetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiListInstanceEventsRequest struct {
	ctx context.Context
	ApiService InstancesApi
	instanceId *string
	types *[]string
	limit *string
	offset *string
	order *string
}

// (Optional) Filter on instance id
func (r ApiListInstanceEventsRequest) InstanceId(instanceId string) ApiListInstanceEventsRequest {
	r.instanceId = &instanceId
	return r
}

// (Optional) Filter on instance event types
func (r ApiListInstanceEventsRequest) Types(types []string) ApiListInstanceEventsRequest {
	r.types = &types
	return r
}

// (Optional) The number of items to return
func (r ApiListInstanceEventsRequest) Limit(limit string) ApiListInstanceEventsRequest {
	r.limit = &limit
	return r
}

// (Optional) The offset in the list of item to return
func (r ApiListInstanceEventsRequest) Offset(offset string) ApiListInstanceEventsRequest {
	r.offset = &offset
	return r
}

// (Optional) Sorts the list in the ascending or the descending order
func (r ApiListInstanceEventsRequest) Order(order string) ApiListInstanceEventsRequest {
	r.order = &order
	return r
}

func (r ApiListInstanceEventsRequest) Execute() (*ListInstanceEventsReply, *http.Response, error) {
	return r.ApiService.ListInstanceEventsExecute(r)
}

/*
ListInstanceEvents List Instance events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListInstanceEventsRequest
*/
func (a *InstancesApiService) ListInstanceEvents(ctx context.Context) ApiListInstanceEventsRequest {
	return ApiListInstanceEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListInstanceEventsReply
func (a *InstancesApiService) ListInstanceEventsExecute(r ApiListInstanceEventsRequest) (*ListInstanceEventsReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListInstanceEventsReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ListInstanceEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instance_events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.instanceId != nil {
		localVarQueryParams.Add("instance_id", parameterToString(*r.instanceId, ""))
	}
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("types", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("types", parameterToString(t, "multi"))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type ApiListInstancesRequest struct {
	ctx context.Context
	ApiService InstancesApi
	appId *string
	serviceId *string
	deploymentId *string
	regionalDeploymentId *string
	allocationId *string
	statuses *[]string
	limit *string
	offset *string
	order *string
}

// (Optional) Filter on application id
func (r ApiListInstancesRequest) AppId(appId string) ApiListInstancesRequest {
	r.appId = &appId
	return r
}

// (Optional) Filter on service id
func (r ApiListInstancesRequest) ServiceId(serviceId string) ApiListInstancesRequest {
	r.serviceId = &serviceId
	return r
}

// (Optional) Filter on deployment id
func (r ApiListInstancesRequest) DeploymentId(deploymentId string) ApiListInstancesRequest {
	r.deploymentId = &deploymentId
	return r
}

// (Optional) Filter on regional deployment id
func (r ApiListInstancesRequest) RegionalDeploymentId(regionalDeploymentId string) ApiListInstancesRequest {
	r.regionalDeploymentId = &regionalDeploymentId
	return r
}

// (Optional) Filter on allocation id
func (r ApiListInstancesRequest) AllocationId(allocationId string) ApiListInstancesRequest {
	r.allocationId = &allocationId
	return r
}

// (Optional) Filter on instance statuses
func (r ApiListInstancesRequest) Statuses(statuses []string) ApiListInstancesRequest {
	r.statuses = &statuses
	return r
}

// (Optional) The number of items to return
func (r ApiListInstancesRequest) Limit(limit string) ApiListInstancesRequest {
	r.limit = &limit
	return r
}

// (Optional) The offset in the list of item to return
func (r ApiListInstancesRequest) Offset(offset string) ApiListInstancesRequest {
	r.offset = &offset
	return r
}

// (Optional) Sorts the list in the ascending or the descending order
func (r ApiListInstancesRequest) Order(order string) ApiListInstancesRequest {
	r.order = &order
	return r
}

func (r ApiListInstancesRequest) Execute() (*ListInstancesReply, *http.Response, error) {
	return r.ApiService.ListInstancesExecute(r)
}

/*
ListInstances List Instances

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListInstancesRequest
*/
func (a *InstancesApiService) ListInstances(ctx context.Context) ApiListInstancesRequest {
	return ApiListInstancesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListInstancesReply
func (a *InstancesApiService) ListInstancesExecute(r ApiListInstancesRequest) (*ListInstancesReply, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListInstancesReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ListInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.appId != nil {
		localVarQueryParams.Add("app_id", parameterToString(*r.appId, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.deploymentId != nil {
		localVarQueryParams.Add("deployment_id", parameterToString(*r.deploymentId, ""))
	}
	if r.regionalDeploymentId != nil {
		localVarQueryParams.Add("regional_deployment_id", parameterToString(*r.regionalDeploymentId, ""))
	}
	if r.allocationId != nil {
		localVarQueryParams.Add("allocation_id", parameterToString(*r.allocationId, ""))
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("statuses", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("statuses", parameterToString(t, "multi"))
		}
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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
