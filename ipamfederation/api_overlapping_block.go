/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type OverlappingBlockAPI interface {
	/*
			Create Create the overlapping block.

			Use this method to create an __OverlappingBlock__ object.
		The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return OverlappingBlockAPICreateRequest
	*/
	Create(ctx context.Context) OverlappingBlockAPICreateRequest

	// CreateExecute executes the request
	//  @return FederationCreateOverlappingBlockResponse
	CreateExecute(r OverlappingBlockAPICreateRequest) (*FederationCreateOverlappingBlockResponse, *http.Response, error)
	/*
			Delete Delete the overlapping block.

			Use this method to delete an __OverlappingBlock__ object.
		The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return OverlappingBlockAPIDeleteRequest
	*/
	Delete(ctx context.Context, id string) OverlappingBlockAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r OverlappingBlockAPIDeleteRequest) (*http.Response, error)
	/*
			List Retrieve the overlapping block.

			Use this method to retrieve __OverlappingBlock__ objects.
		The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return OverlappingBlockAPIListRequest
	*/
	List(ctx context.Context) OverlappingBlockAPIListRequest

	// ListExecute executes the request
	//  @return FederationListOverlappingBlockResponse
	ListExecute(r OverlappingBlockAPIListRequest) (*FederationListOverlappingBlockResponse, *http.Response, error)
	/*
			Read Retrieve the overlapping block.

			Use this method to retrieve an __OverlappingBlock__ object.
		The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return OverlappingBlockAPIReadRequest
	*/
	Read(ctx context.Context, id string) OverlappingBlockAPIReadRequest

	// ReadExecute executes the request
	//  @return FederationReadOverlappingBlockResponse
	ReadExecute(r OverlappingBlockAPIReadRequest) (*FederationReadOverlappingBlockResponse, *http.Response, error)
	/*
			Update Update the overlapping block.

			Use this method to update an __OverlappingBlock__ object.
		The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return OverlappingBlockAPIUpdateRequest
	*/
	Update(ctx context.Context, id string) OverlappingBlockAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return FederationUpdateOverlappingBlockResponse
	UpdateExecute(r OverlappingBlockAPIUpdateRequest) (*FederationUpdateOverlappingBlockResponse, *http.Response, error)
}

// OverlappingBlockAPIService OverlappingBlockAPI service
type OverlappingBlockAPIService internal.Service

type OverlappingBlockAPICreateRequest struct {
	ctx        context.Context
	ApiService OverlappingBlockAPI
	body       *FederationOverlappingBlock
}

func (r OverlappingBlockAPICreateRequest) Body(body FederationOverlappingBlock) OverlappingBlockAPICreateRequest {
	r.body = &body
	return r
}

func (r OverlappingBlockAPICreateRequest) Execute() (*FederationCreateOverlappingBlockResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create the overlapping block.

Use this method to create an __OverlappingBlock__ object.
The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OverlappingBlockAPICreateRequest
*/
func (a *OverlappingBlockAPIService) Create(ctx context.Context) OverlappingBlockAPICreateRequest {
	return OverlappingBlockAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FederationCreateOverlappingBlockResponse
func (a *OverlappingBlockAPIService) CreateExecute(r OverlappingBlockAPICreateRequest) (*FederationCreateOverlappingBlockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *FederationCreateOverlappingBlockResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OverlappingBlockAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/overlapping_block"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type OverlappingBlockAPIDeleteRequest struct {
	ctx        context.Context
	ApiService OverlappingBlockAPI
	id         string
}

func (r OverlappingBlockAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete the overlapping block.

Use this method to delete an __OverlappingBlock__ object.
The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return OverlappingBlockAPIDeleteRequest
*/
func (a *OverlappingBlockAPIService) Delete(ctx context.Context, id string) OverlappingBlockAPIDeleteRequest {
	return OverlappingBlockAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OverlappingBlockAPIService) DeleteExecute(r OverlappingBlockAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OverlappingBlockAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/overlapping_block/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OverlappingBlockAPIListRequest struct {
	ctx        context.Context
	ApiService OverlappingBlockAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	torderBy   *string
	tfilter    *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r OverlappingBlockAPIListRequest) Fields(fields string) OverlappingBlockAPIListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r OverlappingBlockAPIListRequest) Filter(filter string) OverlappingBlockAPIListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r OverlappingBlockAPIListRequest) Offset(offset int32) OverlappingBlockAPIListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r OverlappingBlockAPIListRequest) Limit(limit int32) OverlappingBlockAPIListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r OverlappingBlockAPIListRequest) PageToken(pageToken string) OverlappingBlockAPIListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r OverlappingBlockAPIListRequest) OrderBy(orderBy string) OverlappingBlockAPIListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for sorting by tags.
func (r OverlappingBlockAPIListRequest) TorderBy(torderBy string) OverlappingBlockAPIListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r OverlappingBlockAPIListRequest) Tfilter(tfilter string) OverlappingBlockAPIListRequest {
	r.tfilter = &tfilter
	return r
}

func (r OverlappingBlockAPIListRequest) Execute() (*FederationListOverlappingBlockResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve the overlapping block.

Use this method to retrieve __OverlappingBlock__ objects.
The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OverlappingBlockAPIListRequest
*/
func (a *OverlappingBlockAPIService) List(ctx context.Context) OverlappingBlockAPIListRequest {
	return OverlappingBlockAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FederationListOverlappingBlockResponse
func (a *OverlappingBlockAPIService) ListExecute(r OverlappingBlockAPIListRequest) (*FederationListOverlappingBlockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *FederationListOverlappingBlockResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OverlappingBlockAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/overlapping_block"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type OverlappingBlockAPIReadRequest struct {
	ctx        context.Context
	ApiService OverlappingBlockAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r OverlappingBlockAPIReadRequest) Fields(fields string) OverlappingBlockAPIReadRequest {
	r.fields = &fields
	return r
}

func (r OverlappingBlockAPIReadRequest) Execute() (*FederationReadOverlappingBlockResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Retrieve the overlapping block.

Use this method to retrieve an __OverlappingBlock__ object.
The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return OverlappingBlockAPIReadRequest
*/
func (a *OverlappingBlockAPIService) Read(ctx context.Context, id string) OverlappingBlockAPIReadRequest {
	return OverlappingBlockAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return FederationReadOverlappingBlockResponse
func (a *OverlappingBlockAPIService) ReadExecute(r OverlappingBlockAPIReadRequest) (*FederationReadOverlappingBlockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *FederationReadOverlappingBlockResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OverlappingBlockAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/overlapping_block/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type OverlappingBlockAPIUpdateRequest struct {
	ctx        context.Context
	ApiService OverlappingBlockAPI
	id         string
	body       *FederationOverlappingBlock
}

func (r OverlappingBlockAPIUpdateRequest) Body(body FederationOverlappingBlock) OverlappingBlockAPIUpdateRequest {
	r.body = &body
	return r
}

func (r OverlappingBlockAPIUpdateRequest) Execute() (*FederationUpdateOverlappingBlockResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update the overlapping block.

Use this method to update an __OverlappingBlock__ object.
The __OverlappingBlock__ indicates an address range that may be managed independently by all participating IPAM services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return OverlappingBlockAPIUpdateRequest
*/
func (a *OverlappingBlockAPIService) Update(ctx context.Context, id string) OverlappingBlockAPIUpdateRequest {
	return OverlappingBlockAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return FederationUpdateOverlappingBlockResponse
func (a *OverlappingBlockAPIService) UpdateExecute(r OverlappingBlockAPIUpdateRequest) (*FederationUpdateOverlappingBlockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *FederationUpdateOverlappingBlockResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OverlappingBlockAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/federation/overlapping_block/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}