# \UniverseApi

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniverseAncestries**](UniverseApi.md#GetUniverseAncestries) | **Get** /universe/ancestries/ | Get ancestries
[**GetUniverseAsteroidBeltsAsteroidBeltId**](UniverseApi.md#GetUniverseAsteroidBeltsAsteroidBeltId) | **Get** /universe/asteroid_belts/{asteroid_belt_id}/ | Get asteroid belt information
[**GetUniverseBloodlines**](UniverseApi.md#GetUniverseBloodlines) | **Get** /universe/bloodlines/ | Get bloodlines
[**GetUniverseCategories**](UniverseApi.md#GetUniverseCategories) | **Get** /universe/categories/ | Get item categories
[**GetUniverseCategoriesCategoryId**](UniverseApi.md#GetUniverseCategoriesCategoryId) | **Get** /universe/categories/{category_id}/ | Get item category information
[**GetUniverseConstellations**](UniverseApi.md#GetUniverseConstellations) | **Get** /universe/constellations/ | Get constellations
[**GetUniverseConstellationsConstellationId**](UniverseApi.md#GetUniverseConstellationsConstellationId) | **Get** /universe/constellations/{constellation_id}/ | Get constellation information
[**GetUniverseFactions**](UniverseApi.md#GetUniverseFactions) | **Get** /universe/factions/ | Get factions
[**GetUniverseGraphics**](UniverseApi.md#GetUniverseGraphics) | **Get** /universe/graphics/ | Get graphics
[**GetUniverseGraphicsGraphicId**](UniverseApi.md#GetUniverseGraphicsGraphicId) | **Get** /universe/graphics/{graphic_id}/ | Get graphic information
[**GetUniverseGroups**](UniverseApi.md#GetUniverseGroups) | **Get** /universe/groups/ | Get item groups
[**GetUniverseGroupsGroupId**](UniverseApi.md#GetUniverseGroupsGroupId) | **Get** /universe/groups/{group_id}/ | Get item group information
[**GetUniverseMoonsMoonId**](UniverseApi.md#GetUniverseMoonsMoonId) | **Get** /universe/moons/{moon_id}/ | Get moon information
[**GetUniversePlanetsPlanetId**](UniverseApi.md#GetUniversePlanetsPlanetId) | **Get** /universe/planets/{planet_id}/ | Get planet information
[**GetUniverseRaces**](UniverseApi.md#GetUniverseRaces) | **Get** /universe/races/ | Get character races
[**GetUniverseRegions**](UniverseApi.md#GetUniverseRegions) | **Get** /universe/regions/ | Get regions
[**GetUniverseRegionsRegionId**](UniverseApi.md#GetUniverseRegionsRegionId) | **Get** /universe/regions/{region_id}/ | Get region information
[**GetUniverseStargatesStargateId**](UniverseApi.md#GetUniverseStargatesStargateId) | **Get** /universe/stargates/{stargate_id}/ | Get stargate information
[**GetUniverseStarsStarId**](UniverseApi.md#GetUniverseStarsStarId) | **Get** /universe/stars/{star_id}/ | Get star information
[**GetUniverseStationsStationId**](UniverseApi.md#GetUniverseStationsStationId) | **Get** /universe/stations/{station_id}/ | Get station information
[**GetUniverseStructures**](UniverseApi.md#GetUniverseStructures) | **Get** /universe/structures/ | List all public structures
[**GetUniverseStructuresStructureId**](UniverseApi.md#GetUniverseStructuresStructureId) | **Get** /universe/structures/{structure_id}/ | Get structure information
[**GetUniverseSystemJumps**](UniverseApi.md#GetUniverseSystemJumps) | **Get** /universe/system_jumps/ | Get system jumps
[**GetUniverseSystemKills**](UniverseApi.md#GetUniverseSystemKills) | **Get** /universe/system_kills/ | Get system kills
[**GetUniverseSystems**](UniverseApi.md#GetUniverseSystems) | **Get** /universe/systems/ | Get solar systems
[**GetUniverseSystemsSystemId**](UniverseApi.md#GetUniverseSystemsSystemId) | **Get** /universe/systems/{system_id}/ | Get solar system information
[**GetUniverseTypes**](UniverseApi.md#GetUniverseTypes) | **Get** /universe/types/ | Get types
[**GetUniverseTypesTypeId**](UniverseApi.md#GetUniverseTypesTypeId) | **Get** /universe/types/{type_id}/ | Get type information
[**PostUniverseIds**](UniverseApi.md#PostUniverseIds) | **Post** /universe/ids/ | Bulk names to IDs
[**PostUniverseNames**](UniverseApi.md#PostUniverseNames) | **Post** /universe/names/ | Get names and categories for a set of ID&#39;s


# **GetUniverseAncestries**
> []GetUniverseAncestries200Ok GetUniverseAncestries(ctx, optional)
Get ancestries

Get all character ancestries  --- Alternate route: `/dev/universe/ancestries/`  Alternate route: `/legacy/universe/ancestries/`  Alternate route: `/v1/universe/ancestries/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**[]GetUniverseAncestries200Ok**](get_universe_ancestries_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseAsteroidBeltsAsteroidBeltId**
> GetUniverseAsteroidBeltsAsteroidBeltIdOk GetUniverseAsteroidBeltsAsteroidBeltId(ctx, asteroidBeltId, optional)
Get asteroid belt information

Get information on an asteroid belt  --- Alternate route: `/dev/universe/asteroid_belts/{asteroid_belt_id}/`  Alternate route: `/legacy/universe/asteroid_belts/{asteroid_belt_id}/`  Alternate route: `/v1/universe/asteroid_belts/{asteroid_belt_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **asteroidBeltId** | **int32**| asteroid_belt_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asteroidBeltId** | **int32**| asteroid_belt_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseAsteroidBeltsAsteroidBeltIdOk**](get_universe_asteroid_belts_asteroid_belt_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseBloodlines**
> []GetUniverseBloodlines200Ok GetUniverseBloodlines(ctx, optional)
Get bloodlines

Get a list of bloodlines  --- Alternate route: `/dev/universe/bloodlines/`  Alternate route: `/legacy/universe/bloodlines/`  Alternate route: `/v1/universe/bloodlines/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**[]GetUniverseBloodlines200Ok**](get_universe_bloodlines_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseCategories**
> []int32 GetUniverseCategories(ctx, optional)
Get item categories

Get a list of item categories  --- Alternate route: `/dev/universe/categories/`  Alternate route: `/legacy/universe/categories/`  Alternate route: `/v1/universe/categories/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseCategoriesCategoryId**
> GetUniverseCategoriesCategoryIdOk GetUniverseCategoriesCategoryId(ctx, categoryId, optional)
Get item category information

Get information of an item category  --- Alternate route: `/dev/universe/categories/{category_id}/`  Alternate route: `/legacy/universe/categories/{category_id}/`  Alternate route: `/v1/universe/categories/{category_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **categoryId** | **int32**| An Eve item category ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **int32**| An Eve item category ID | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseCategoriesCategoryIdOk**](get_universe_categories_category_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseConstellations**
> []int32 GetUniverseConstellations(ctx, optional)
Get constellations

Get a list of constellations  --- Alternate route: `/dev/universe/constellations/`  Alternate route: `/legacy/universe/constellations/`  Alternate route: `/v1/universe/constellations/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseConstellationsConstellationId**
> GetUniverseConstellationsConstellationIdOk GetUniverseConstellationsConstellationId(ctx, constellationId, optional)
Get constellation information

Get information on a constellation  --- Alternate route: `/dev/universe/constellations/{constellation_id}/`  Alternate route: `/legacy/universe/constellations/{constellation_id}/`  Alternate route: `/v1/universe/constellations/{constellation_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **constellationId** | **int32**| constellation_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constellationId** | **int32**| constellation_id integer | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseConstellationsConstellationIdOk**](get_universe_constellations_constellation_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseFactions**
> []GetUniverseFactions200Ok GetUniverseFactions(ctx, optional)
Get factions

Get a list of factions  --- Alternate route: `/dev/universe/factions/`  Alternate route: `/v2/universe/factions/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**[]GetUniverseFactions200Ok**](get_universe_factions_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGraphics**
> []int32 GetUniverseGraphics(ctx, optional)
Get graphics

Get a list of graphics  --- Alternate route: `/dev/universe/graphics/`  Alternate route: `/legacy/universe/graphics/`  Alternate route: `/v1/universe/graphics/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGraphicsGraphicId**
> GetUniverseGraphicsGraphicIdOk GetUniverseGraphicsGraphicId(ctx, graphicId, optional)
Get graphic information

Get information on a graphic  --- Alternate route: `/dev/universe/graphics/{graphic_id}/`  Alternate route: `/legacy/universe/graphics/{graphic_id}/`  Alternate route: `/v1/universe/graphics/{graphic_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **graphicId** | **int32**| graphic_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphicId** | **int32**| graphic_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseGraphicsGraphicIdOk**](get_universe_graphics_graphic_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroups**
> []int32 GetUniverseGroups(ctx, optional)
Get item groups

Get a list of item groups  --- Alternate route: `/dev/universe/groups/`  Alternate route: `/legacy/universe/groups/`  Alternate route: `/v1/universe/groups/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroupsGroupId**
> GetUniverseGroupsGroupIdOk GetUniverseGroupsGroupId(ctx, groupId, optional)
Get item group information

Get information on an item group  --- Alternate route: `/dev/universe/groups/{group_id}/`  Alternate route: `/legacy/universe/groups/{group_id}/`  Alternate route: `/v1/universe/groups/{group_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int32**| An Eve item group ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32**| An Eve item group ID | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseGroupsGroupIdOk**](get_universe_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseMoonsMoonId**
> GetUniverseMoonsMoonIdOk GetUniverseMoonsMoonId(ctx, moonId, optional)
Get moon information

Get information on a moon  --- Alternate route: `/dev/universe/moons/{moon_id}/`  Alternate route: `/legacy/universe/moons/{moon_id}/`  Alternate route: `/v1/universe/moons/{moon_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **moonId** | **int32**| moon_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moonId** | **int32**| moon_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseMoonsMoonIdOk**](get_universe_moons_moon_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniversePlanetsPlanetId**
> GetUniversePlanetsPlanetIdOk GetUniversePlanetsPlanetId(ctx, planetId, optional)
Get planet information

Get information on a planet  --- Alternate route: `/dev/universe/planets/{planet_id}/`  Alternate route: `/legacy/universe/planets/{planet_id}/`  Alternate route: `/v1/universe/planets/{planet_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **planetId** | **int32**| planet_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planetId** | **int32**| planet_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniversePlanetsPlanetIdOk**](get_universe_planets_planet_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseRaces**
> []GetUniverseRaces200Ok GetUniverseRaces(ctx, optional)
Get character races

Get a list of character races  --- Alternate route: `/dev/universe/races/`  Alternate route: `/legacy/universe/races/`  Alternate route: `/v1/universe/races/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**[]GetUniverseRaces200Ok**](get_universe_races_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseRegions**
> []int32 GetUniverseRegions(ctx, optional)
Get regions

Get a list of regions  --- Alternate route: `/dev/universe/regions/`  Alternate route: `/legacy/universe/regions/`  Alternate route: `/v1/universe/regions/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseRegionsRegionId**
> GetUniverseRegionsRegionIdOk GetUniverseRegionsRegionId(ctx, regionId, optional)
Get region information

Get information on a region  --- Alternate route: `/dev/universe/regions/{region_id}/`  Alternate route: `/legacy/universe/regions/{region_id}/`  Alternate route: `/v1/universe/regions/{region_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **regionId** | **int32**| region_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int32**| region_id integer | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseRegionsRegionIdOk**](get_universe_regions_region_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStargatesStargateId**
> GetUniverseStargatesStargateIdOk GetUniverseStargatesStargateId(ctx, stargateId, optional)
Get stargate information

Get information on a stargate  --- Alternate route: `/dev/universe/stargates/{stargate_id}/`  Alternate route: `/legacy/universe/stargates/{stargate_id}/`  Alternate route: `/v1/universe/stargates/{stargate_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **stargateId** | **int32**| stargate_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stargateId** | **int32**| stargate_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStargatesStargateIdOk**](get_universe_stargates_stargate_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStarsStarId**
> GetUniverseStarsStarIdOk GetUniverseStarsStarId(ctx, starId, optional)
Get star information

Get information on a star  --- Alternate route: `/dev/universe/stars/{star_id}/`  Alternate route: `/legacy/universe/stars/{star_id}/`  Alternate route: `/v1/universe/stars/{star_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **starId** | **int32**| star_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **starId** | **int32**| star_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStarsStarIdOk**](get_universe_stars_star_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStationsStationId**
> GetUniverseStationsStationIdOk GetUniverseStationsStationId(ctx, stationId, optional)
Get station information

Get information on a station  --- Alternate route: `/dev/universe/stations/{station_id}/`  Alternate route: `/v2/universe/stations/{station_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **stationId** | **int32**| station_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stationId** | **int32**| station_id integer | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStationsStationIdOk**](get_universe_stations_station_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructures**
> []int64 GetUniverseStructures(ctx, optional)
List all public structures

List all public structures  --- Alternate route: `/dev/universe/structures/`  Alternate route: `/legacy/universe/structures/`  Alternate route: `/v1/universe/structures/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructuresStructureId**
> GetUniverseStructuresStructureIdOk GetUniverseStructuresStructureId(ctx, structureId, optional)
Get structure information

Returns information on requested structure if you are on the ACL. Otherwise, returns \"Forbidden\" for all inputs.  --- Alternate route: `/dev/universe/structures/{structure_id}/`  Alternate route: `/v2/universe/structures/{structure_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **structureId** | **int64**| An Eve structure ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **structureId** | **int64**| An Eve structure ID | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **string**| Access token to use if unable to set a header | 

### Return type

[**GetUniverseStructuresStructureIdOk**](get_universe_structures_structure_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemJumps**
> []GetUniverseSystemJumps200Ok GetUniverseSystemJumps(ctx, optional)
Get system jumps

Get the number of jumps in solar systems within the last hour ending at the timestamp of the Last-Modified header, excluding wormhole space. Only systems with jumps will be listed  --- Alternate route: `/dev/universe/system_jumps/`  Alternate route: `/legacy/universe/system_jumps/`  Alternate route: `/v1/universe/system_jumps/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetUniverseSystemJumps200Ok**](get_universe_system_jumps_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemKills**
> []GetUniverseSystemKills200Ok GetUniverseSystemKills(ctx, optional)
Get system kills

Get the number of ship, pod and NPC kills per solar system within the last hour ending at the timestamp of the Last-Modified header, excluding wormhole space. Only systems with kills will be listed  --- Alternate route: `/dev/universe/system_kills/`  Alternate route: `/v2/universe/system_kills/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetUniverseSystemKills200Ok**](get_universe_system_kills_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystems**
> []int32 GetUniverseSystems(ctx, optional)
Get solar systems

Get a list of solar systems  --- Alternate route: `/dev/universe/systems/`  Alternate route: `/legacy/universe/systems/`  Alternate route: `/v1/universe/systems/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemsSystemId**
> GetUniverseSystemsSystemIdOk GetUniverseSystemsSystemId(ctx, systemId, optional)
Get solar system information

Get information on a solar system.  --- Alternate route: `/dev/universe/systems/{system_id}/`  Alternate route: `/v4/universe/systems/{system_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **systemId** | **int32**| system_id integer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemId** | **int32**| system_id integer | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseSystemsSystemIdOk**](get_universe_systems_system_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypes**
> []int32 GetUniverseTypes(ctx, optional)
Get types

Get a list of type ids  --- Alternate route: `/dev/universe/types/`  Alternate route: `/legacy/universe/types/`  Alternate route: `/v1/universe/types/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **int32**| Which page of results to return | [default to 1]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypesTypeId**
> GetUniverseTypesTypeIdOk GetUniverseTypesTypeId(ctx, typeId, optional)
Get type information

Get information on a type  --- Alternate route: `/dev/universe/types/{type_id}/`  Alternate route: `/v3/universe/types/{type_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **typeId** | **int32**| An Eve item type ID | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeId** | **int32**| An Eve item type ID | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **string**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**GetUniverseTypesTypeIdOk**](get_universe_types_type_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUniverseIds**
> PostUniverseIdsOk PostUniverseIds(ctx, names, optional)
Bulk names to IDs

Resolve a set of names to IDs in the following categories: agents, alliances, characters, constellations, corporations factions, inventory_types, regions, stations, and systems. Only exact matches will be returned. All names searched for are cached for 12 hours.  --- Alternate route: `/dev/universe/ids/`  Alternate route: `/legacy/universe/ids/`  Alternate route: `/v1/universe/ids/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **names** | **[]string**| The names to resolve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string**| The names to resolve | 
 **acceptLanguage** | **string**| Language to use in the response | [default to en-us]
 **datasource** | **string**| The server name you would like data from | [default to tranquility]
 **language** | **string**| Language to use in the response, takes precedence over Accept-Language | [default to en-us]

### Return type

[**PostUniverseIdsOk**](post_universe_ids_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUniverseNames**
> []PostUniverseNames200Ok PostUniverseNames(ctx, ids, optional)
Get names and categories for a set of ID's

Resolve a set of IDs to names and categories. Supported ID's for resolving are: Characters, Corporations, Alliances, Stations, Solar Systems, Constellations, Regions, Types.  --- Alternate route: `/dev/universe/names/`  Alternate route: `/v2/universe/names/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ids** | **[]int32**| The ids to resolve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int32**| The ids to resolve | 
 **datasource** | **string**| The server name you would like data from | [default to tranquility]

### Return type

[**[]PostUniverseNames200Ok**](post_universe_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

