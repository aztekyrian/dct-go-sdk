# \EnginesApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivityCheck**](EnginesApi.md#ConnectivityCheck) | **Post** /engines/connnectivity/check | Checks connectivity between an engine and a remote host on a given port.
[**GetEngineById**](EnginesApi.md#GetEngineById) | **Get** /engines/{engineId} | Returns an engine by ID.
[**GetEngines**](EnginesApi.md#GetEngines) | **Get** /engines | List all engines.



## ConnectivityCheck

> EngineConnectivityCheckResponse ConnectivityCheck(ctx).EngineConnectivityCheckRequest(engineConnectivityCheckRequest).Execute()

Checks connectivity between an engine and a remote host on a given port.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    engineConnectivityCheckRequest := *openapiclient.NewEngineConnectivityCheckRequest("EngineId_example", "Host_example", NullableInt32(123)) // EngineConnectivityCheckRequest | The api to check connectivity of engine and a remote host on given port.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.ConnectivityCheck(context.Background()).EngineConnectivityCheckRequest(engineConnectivityCheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.ConnectivityCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectivityCheck`: EngineConnectivityCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.ConnectivityCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineConnectivityCheckRequest** | [**EngineConnectivityCheckRequest**](EngineConnectivityCheckRequest.md) | The api to check connectivity of engine and a remote host on given port. | 

### Return type

[**EngineConnectivityCheckResponse**](EngineConnectivityCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngineById

> Engine GetEngineById(ctx, engineId).Execute()

Returns an engine by ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    engineId := "engineId_example" // string | The ID of the engine.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngineById(context.Background(), engineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngineById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngineById`: Engine
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngineById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**engineId** | **string** | The ID of the engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngineByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Engine**](Engine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngines

> ListEnginesResponse GetEngines(ctx).Execute()

List all engines.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnginesApi.GetEngines(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnginesApi.GetEngines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEngines`: ListEnginesResponse
    fmt.Fprintf(os.Stdout, "Response from `EnginesApi.GetEngines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnginesRequest struct via the builder pattern


### Return type

[**ListEnginesResponse**](ListEnginesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

