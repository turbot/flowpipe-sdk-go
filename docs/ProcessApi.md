# \ProcessApi

All URIs are relative to *https://localhost/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](ProcessApi.md#Get) | **Get** /process/{process_id} | Get process
[**GetExecution**](ProcessApi.md#GetExecution) | **Get** /process/{process_id}/execution | Get process execution
[**GetLog**](ProcessApi.md#GetLog) | **Get** /process/{process_id}/log/process.json | Get process log
[**List**](ProcessApi.md#List) | **Get** /process | List processs



## Get

> Process Get(ctx, processId).Execute()

Get process



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/turbot/flowpipe-sdk-go"
)

func main() {
    processId := "processId_example" // string | The name of the process

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessApi.Get(context.Background(), processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Process
    fmt.Fprintf(os.Stdout, "Response from `ProcessApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | The name of the process | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Process**](Process.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> ExecutionExecution GetExecution(ctx, processId).Execute()

Get process execution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/turbot/flowpipe-sdk-go"
)

func main() {
    processId := "processId_example" // string | The name of the process

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessApi.GetExecution(context.Background(), processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessApi.GetExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecution`: ExecutionExecution
    fmt.Fprintf(os.Stdout, "Response from `ProcessApi.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | The name of the process | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecutionExecution**](ExecutionExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLog

> ListProcessLogJSONResponse GetLog(ctx, processId).Execute()

Get process log



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/turbot/flowpipe-sdk-go"
)

func main() {
    processId := "processId_example" // string | The id of the process

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessApi.GetLog(context.Background(), processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessApi.GetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLog`: ListProcessLogJSONResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessApi.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processId** | **string** | The id of the process | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListProcessLogJSONResponse**](ListProcessLogJSONResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListProcessResponse List(ctx).Limit(limit).NextToken(nextToken).Execute()

List processs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/turbot/flowpipe-sdk-go"
)

func main() {
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProcessApi.List(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProcessApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListProcessResponse
    fmt.Fprintf(os.Stdout, "Response from `ProcessApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListProcessResponse**](ListProcessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

