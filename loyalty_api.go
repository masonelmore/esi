/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.2.9
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

type LoyaltyApiService service


/* LoyaltyApiService Get loyalty points
 Return a list of loyalty points for all corporations the character has worked for  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param characterId An EVE character ID
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "datasource" (string) The server name you would like data from
     @param "ifNoneMatch" (string) ETag from a previous request. A 304 will be returned if this matches the current ETag
     @param "token" (string) Access token to use if unable to set a header
 @return []GetCharactersCharacterIdLoyaltyPoints200Ok*/
func (a *LoyaltyApiService) GetCharactersCharacterIdLoyaltyPoints(ctx context.Context, characterId int32, localVarOptionals map[string]interface{}) ([]GetCharactersCharacterIdLoyaltyPoints200Ok,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GetCharactersCharacterIdLoyaltyPoints200Ok
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/characters/{character_id}/loyalty/points/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if characterId < 1 {
		return successPayload, nil, reportError("characterId must be greater than 1")
	}
	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["ifNoneMatch"], "string", "ifNoneMatch"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
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

/* LoyaltyApiService List loyalty store offers
 Return a list of offers from a specific corporation&#39;s loyalty store  ---  This route expires daily at 11:05
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param corporationId An EVE corporation ID
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "datasource" (string) The server name you would like data from
     @param "ifNoneMatch" (string) ETag from a previous request. A 304 will be returned if this matches the current ETag
 @return []GetLoyaltyStoresCorporationIdOffers200Ok*/
func (a *LoyaltyApiService) GetLoyaltyStoresCorporationIdOffers(ctx context.Context, corporationId int32, localVarOptionals map[string]interface{}) ([]GetLoyaltyStoresCorporationIdOffers200Ok,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GetLoyaltyStoresCorporationIdOffers200Ok
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/loyalty/stores/{corporation_id}/offers/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if corporationId < 1 {
		return successPayload, nil, reportError("corporationId must be greater than 1")
	}
	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["ifNoneMatch"], "string", "ifNoneMatch"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
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

