/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 0.8.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package esi

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type InsuranceApiService service


/* InsuranceApiService List insurance levels
 Return available insurance levels for all ship types  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "acceptLanguage" (string) Language to use in the response
     @param "datasource" (string) The server name you would like data from
     @param "ifNoneMatch" (string) ETag from a previous request. A 304 will be returned if this matches the current ETag
     @param "language" (string) Language to use in the response, takes precedence over Accept-Language
 @return []GetInsurancePrices200Ok*/
func (a *InsuranceApiService) GetInsurancePrices(ctx context.Context, localVarOptionals map[string]interface{}) ([]GetInsurancePrices200Ok,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GetInsurancePrices200Ok
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/insurance/prices/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["acceptLanguage"], "string", "acceptLanguage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["ifNoneMatch"], "string", "ifNoneMatch"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["language"], "string", "language"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["language"].(string); localVarOk {
		localVarQueryParams.Add("language", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["acceptLanguage"].(string); localVarOk {
		localVarHeaderParams["Accept-Language"] = parameterToString(localVarTempParam, "")
	}
	if localVarTempParam, localVarOk := localVarOptionals["ifNoneMatch"].(string); localVarOk {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

