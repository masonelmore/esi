# \LoyaltyApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdLoyaltyPoints**](LoyaltyApi.md#GetCharactersCharacterIdLoyaltyPoints) | **Get** /v1/characters/{character_id}/loyalty/points/ | Get loyalty points
[**GetLoyaltyStoresCorporationIdOffers**](LoyaltyApi.md#GetLoyaltyStoresCorporationIdOffers) | **Get** /v1/loyalty/stores/{corporation_id}/offers/ | List loyalty store offers


# **GetCharactersCharacterIdLoyaltyPoints**
> []GetCharactersCharacterIdLoyaltyPoints200Ok GetCharactersCharacterIdLoyaltyPoints(ctx, characterId, optional)
Get loyalty points

Return a list of loyalty points for all corporations the character has worked for  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdLoyaltyPoints200Ok**](get_characters_character_id_loyalty_points_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoyaltyStoresCorporationIdOffers**
> []GetLoyaltyStoresCorporationIdOffers200Ok GetLoyaltyStoresCorporationIdOffers(ctx, corporationId, optional)
List loyalty store offers

Return a list of offers from a specific corporation's loyalty store  ---  This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetLoyaltyStoresCorporationIdOffers200Ok**](get_loyalty_stores_corporation_id_offers_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

