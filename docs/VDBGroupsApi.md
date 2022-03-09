# \VDBGroupsApi

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVdbGroup**](VDBGroupsApi.md#CreateVdbGroup) | **Post** /vdb-groups | Create a new VDBGroup.
[**DeleteVdbGroup**](VDBGroupsApi.md#DeleteVdbGroup) | **Delete** /vdb-groups/{vdbGroupId} | Delete a VDBGoup.
[**GetBookmarksByVdbGroup**](VDBGroupsApi.md#GetBookmarksByVdbGroup) | **Get** /vdb-groups/{vdbGroupId}/bookmarks | List bookmarks compatible with this VDB Group.
[**GetVdbGroup**](VDBGroupsApi.md#GetVdbGroup) | **Get** /vdb-groups/{vdbGroupId} | Get a VDBGroup by name.
[**GetVdbGroups**](VDBGroupsApi.md#GetVdbGroups) | **Get** /vdb-groups | List all VDBGroups.



## CreateVdbGroup

> CreateVDBGroupResponse CreateVdbGroup(ctx).CreateVDBGroupRequest(createVDBGroupRequest).Execute()

Create a new VDBGroup.

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
    createVDBGroupRequest := *openapiclient.NewCreateVDBGroupRequest("Name_example", []string{"VdbIds_example"}) // CreateVDBGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.CreateVdbGroup(context.Background()).CreateVDBGroupRequest(createVDBGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.CreateVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVdbGroup`: CreateVDBGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.CreateVdbGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVDBGroupRequest** | [**CreateVDBGroupRequest**](CreateVDBGroupRequest.md) |  | 

### Return type

[**CreateVDBGroupResponse**](CreateVDBGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVdbGroup

> DeleteVdbGroup(ctx, vdbGroupId).Execute()

Delete a VDBGoup.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.DeleteVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.DeleteVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookmarksByVdbGroup

> ListBookmarksByVDBGroupsResponse GetBookmarksByVdbGroup(ctx, vdbGroupId).Execute()

List bookmarks compatible with this VDB Group.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetBookmarksByVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetBookmarksByVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookmarksByVdbGroup`: ListBookmarksByVDBGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetBookmarksByVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookmarksByVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBookmarksByVDBGroupsResponse**](ListBookmarksByVDBGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbGroup

> VDBGroup GetVdbGroup(ctx, vdbGroupId).Execute()

Get a VDBGroup by name.

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
    vdbGroupId := "vdbGroupId_example" // string | The ID or name of the VDBGroup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VDBGroupsApi.GetVdbGroup(context.Background(), vdbGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetVdbGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbGroup`: VDBGroup
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetVdbGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdbGroupId** | **string** | The ID or name of the VDBGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VDBGroup**](VDBGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdbGroups

> ListVDBGroupsResponse GetVdbGroups(ctx).Execute()

List all VDBGroups.

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
    resp, r, err := apiClient.VDBGroupsApi.GetVdbGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VDBGroupsApi.GetVdbGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdbGroups`: ListVDBGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `VDBGroupsApi.GetVdbGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdbGroupsRequest struct via the builder pattern


### Return type

[**ListVDBGroupsResponse**](ListVDBGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

