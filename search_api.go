/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.5
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
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type SearchApiService service


/* SearchApiService Search on a string
 Search for entities that match a given sub-string.  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param categories Type of entities to search for
 @param characterId An EVE character ID
 @param search The string to search on
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "acceptLanguage" (string) Language to use in the response
     @param "datasource" (string) The server name you would like data from
     @param "ifNoneMatch" (string) ETag from a previous request. A 304 will be returned if this matches the current ETag
     @param "language" (string) Language to use in the response, takes precedence over Accept-Language
     @param "strict" (bool) Whether the search should be a strict match
     @param "token" (string) Access token to use if unable to set a header
 @return GetCharactersCharacterIdSearchOk*/
func (a *SearchApiService) GetCharactersCharacterIdSearch(ctx context.Context, categories []string, characterId int32, search string, localVarOptionals map[string]interface{}) (GetCharactersCharacterIdSearchOk,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetCharactersCharacterIdSearchOk
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v3/characters/{character_id}/search/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if len(categories) < 1 {
		return successPayload, nil, reportError("categories must have at least 1 elements")
	}
	if len(categories) > 11 {
		return successPayload, nil, reportError("categories must have less than 11 elements")
	}
	if characterId < 1 {
		return successPayload, nil, reportError("characterId must be greater than 1")
	}
	if strlen(search) < 3 {
		return successPayload, nil, reportError("search must have at least 3 elements")
	}
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
	if err := typeCheckParameter(localVarOptionals["strict"], "bool", "strict"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("categories", parameterToString(categories, "csv"))
	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["language"].(string); localVarOk {
		localVarQueryParams.Add("language", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("search", parameterToString(search, ""))
	if localVarTempParam, localVarOk := localVarOptionals["strict"].(bool); localVarOk {
		localVarQueryParams.Add("strict", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
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

/* SearchApiService Search on a string
 Search for entities that match a given sub-string.  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param categories Type of entities to search for
 @param search The string to search on
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "acceptLanguage" (string) Language to use in the response
     @param "datasource" (string) The server name you would like data from
     @param "ifNoneMatch" (string) ETag from a previous request. A 304 will be returned if this matches the current ETag
     @param "language" (string) Language to use in the response, takes precedence over Accept-Language
     @param "strict" (bool) Whether the search should be a strict match
 @return GetSearchOk*/
func (a *SearchApiService) GetSearch(ctx context.Context, categories []string, search string, localVarOptionals map[string]interface{}) (GetSearchOk,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetSearchOk
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v2/search/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if len(categories) < 1 {
		return successPayload, nil, reportError("categories must have at least 1 elements")
	}
	if len(categories) > 10 {
		return successPayload, nil, reportError("categories must have less than 10 elements")
	}
	if strlen(search) < 3 {
		return successPayload, nil, reportError("search must have at least 3 elements")
	}
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
	if err := typeCheckParameter(localVarOptionals["strict"], "bool", "strict"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("categories", parameterToString(categories, "csv"))
	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["language"].(string); localVarOk {
		localVarQueryParams.Add("language", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("search", parameterToString(search, ""))
	if localVarTempParam, localVarOk := localVarOptionals["strict"].(bool); localVarOk {
		localVarQueryParams.Add("strict", parameterToString(localVarTempParam, ""))
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

