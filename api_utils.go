/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 5.0.0
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// UtilsAPIService UtilsAPI service
type UtilsAPIService service

type ApiSqlRequest struct {
	ctx context.Context
	ApiService *UtilsAPIService
	body *string
	rawResponse *bool
}

// A query parameter string. 
func (r ApiSqlRequest) Body(body string) ApiSqlRequest {
	r.body = &body
	return r
}

// Optional parameter, defines a format of response. Can be set to &#x60;False&#x60; for Select only queries and set to &#x60;True&#x60; for any type of queries. Default value is &#39;True&#39;. 
func (r ApiSqlRequest) RawResponse(rawResponse bool) ApiSqlRequest {
	r.rawResponse = &rawResponse
	return r
}

func (r ApiSqlRequest) Execute() (*SqlResponse, *http.Response, error) {
	return r.ApiService.SqlExecute(r)
}

/*
Sql Perform SQL requests

Run a query in SQL format.
Expects a query string passed through `body` parameter and optional `raw_response` parameter that defines a format of response.
`raw_response` can be set to `False` for Select queries only, e.g., `SELECT * FROM myindex`
The query string must stay as it is, no URL encoding is needed.
The response object depends on the query executed. In select mode the response has same format as `/search` operation.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSqlRequest
*/
func (a *UtilsAPIService) Sql(ctx context.Context) ApiSqlRequest {
	return ApiSqlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SqlResponse
func (a *UtilsAPIService) SqlExecute(r ApiSqlRequest) (*SqlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SqlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UtilsAPIService.Sql")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.rawResponse == nil {
		var defaultValue bool = true
		r.rawResponse = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "raw_response", r.rawResponse, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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
			var v ErrorResponse
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
