# \ContractsApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdContracts**](ContractsApi.md#GetCharactersCharacterIdContracts) | **Get** /v1/characters/{character_id}/contracts/ | Get contracts
[**GetCharactersCharacterIdContractsContractIdBids**](ContractsApi.md#GetCharactersCharacterIdContractsContractIdBids) | **Get** /v1/characters/{character_id}/contracts/{contract_id}/bids/ | Get contract bids
[**GetCharactersCharacterIdContractsContractIdItems**](ContractsApi.md#GetCharactersCharacterIdContractsContractIdItems) | **Get** /v1/characters/{character_id}/contracts/{contract_id}/items/ | Get contract items
[**GetContractsPublicBidsContractId**](ContractsApi.md#GetContractsPublicBidsContractId) | **Get** /v1/contracts/public/bids/{contract_id}/ | Get public contract bids
[**GetContractsPublicItemsContractId**](ContractsApi.md#GetContractsPublicItemsContractId) | **Get** /v1/contracts/public/items/{contract_id}/ | Get public contract items
[**GetContractsPublicRegionId**](ContractsApi.md#GetContractsPublicRegionId) | **Get** /v1/contracts/public/{region_id}/ | Get public contracts
[**GetCorporationsCorporationIdContracts**](ContractsApi.md#GetCorporationsCorporationIdContracts) | **Get** /v1/corporations/{corporation_id}/contracts/ | Get corporation contracts
[**GetCorporationsCorporationIdContractsContractIdBids**](ContractsApi.md#GetCorporationsCorporationIdContractsContractIdBids) | **Get** /v1/corporations/{corporation_id}/contracts/{contract_id}/bids/ | Get corporation contract bids
[**GetCorporationsCorporationIdContractsContractIdItems**](ContractsApi.md#GetCorporationsCorporationIdContractsContractIdItems) | **Get** /v1/corporations/{corporation_id}/contracts/{contract_id}/items/ | Get corporation contract items


# **GetCharactersCharacterIdContracts**
> []GetCharactersCharacterIdContracts200Ok GetCharactersCharacterIdContracts(ctx, characterId, optional)
Get contracts

Returns contracts available to a character, only if the character is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is \"in_progress\".  ---  This route is cached for up to 300 seconds

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
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContracts200Ok**](get_characters_character_id_contracts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContractsContractIdBids**
> []GetCharactersCharacterIdContractsContractIdBids200Ok GetCharactersCharacterIdContractsContractIdBids(ctx, characterId, contractId, optional)
Get contract bids

Lists bids on a particular auction contract  ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **characterId** | **int32**| An EVE character ID | 
  **contractId** | **int32**| ID of a contract | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contractId** | **int32**| ID of a contract | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContractsContractIdBids200Ok**](get_characters_character_id_contracts_contract_id_bids_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContractsContractIdItems**
> []GetCharactersCharacterIdContractsContractIdItems200Ok GetCharactersCharacterIdContractsContractIdItems(ctx, characterId, contractId, optional)
Get contract items

Lists items of a particular contract  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **characterId** | **int32**| An EVE character ID | 
  **contractId** | **int32**| ID of a contract | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int32**| An EVE character ID | 
 **contractId** | **int32**| ID of a contract | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContractsContractIdItems200Ok**](get_characters_character_id_contracts_contract_id_items_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractsPublicBidsContractId**
> []GetContractsPublicBidsContractId200Ok GetContractsPublicBidsContractId(ctx, contractId, optional)
Get public contract bids

Lists bids on a public auction contract  ---  This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contractId** | **int32**| ID of a contract | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int32**| ID of a contract | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

[**[]GetContractsPublicBidsContractId200Ok**](get_contracts_public_bids_contract_id_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractsPublicItemsContractId**
> []GetContractsPublicItemsContractId200Ok GetContractsPublicItemsContractId(ctx, contractId, optional)
Get public contract items

Lists items of a public contract  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contractId** | **int32**| ID of a contract | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int32**| ID of a contract | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

[**[]GetContractsPublicItemsContractId200Ok**](get_contracts_public_items_contract_id_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractsPublicRegionId**
> []GetContractsPublicRegionId200Ok GetContractsPublicRegionId(ctx, regionId, optional)
Get public contracts

Returns a paginated list of all public contracts in the given region  ---  This route is cached for up to 1800 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **regionId** | **int32**| An EVE region id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32**| An EVE region id | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

[**[]GetContractsPublicRegionId200Ok**](get_contracts_public_region_id_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContracts**
> []GetCorporationsCorporationIdContracts200Ok GetCorporationsCorporationIdContracts(ctx, corporationId, optional)
Get corporation contracts

Returns contracts available to a corporation, only if the corporation is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is \"in_progress\".  ---  This route is cached for up to 300 seconds

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
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContracts200Ok**](get_corporations_corporation_id_contracts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContractsContractIdBids**
> []GetCorporationsCorporationIdContractsContractIdBids200Ok GetCorporationsCorporationIdContractsContractIdBids(ctx, contractId, corporationId, optional)
Get corporation contract bids

Lists bids on a particular auction contract  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contractId** | **int32**| ID of a contract | 
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int32**| ID of a contract | 
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContractsContractIdBids200Ok**](get_corporations_corporation_id_contracts_contract_id_bids_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContractsContractIdItems**
> []GetCorporationsCorporationIdContractsContractIdItems200Ok GetCorporationsCorporationIdContractsContractIdItems(ctx, contractId, corporationId, optional)
Get corporation contract items

Lists items of a particular contract  ---  This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contractId** | **int32**| ID of a contract | 
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **int32**| ID of a contract | 
 **corporationId** | **int32**| An EVE corporation ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContractsContractIdItems200Ok**](get_corporations_corporation_id_contracts_contract_id_items_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

