# \SourcesApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceById**](SourcesApi.md#GetSourceById) | **Get** /sources/{sourceId} | Get a source by ID.
[**GetSources**](SourcesApi.md#GetSources) | **Get** /sources | List all sources.



## GetSourceById

> Source GetSourceById(ctx, sourceId).Execute()

Get a source by ID.

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
    sourceId := "sourceId_example" // string | The ID of the source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesApi.GetSourceById(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceById`: Source
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The ID of the source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSources

> ListSourcesResponse GetSources(ctx).Execute()

List all sources.

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
    resp, r, err := apiClient.SourcesApi.GetSources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesApi.GetSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSources`: ListSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcesApi.GetSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcesRequest struct via the builder pattern


### Return type

[**ListSourcesResponse**](ListSourcesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

