# \StatusApi

All URIs are relative to *https://esi.evetech.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatus**](StatusApi.md#GetStatus) | **Get** /v1/status/ | Retrieve the uptime and player counts


# **GetStatus**
> GetStatusOk GetStatus(ctx, optional)
Retrieve the uptime and player counts

EVE Server status  ---  This route is cached for up to 30 seconds

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

[**GetStatusOk**](get_status_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

