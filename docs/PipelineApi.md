# \PipelineApi

All URIs are relative to *https://localhost/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cmd**](PipelineApi.md#Cmd) | **Post** /pipeline/{pipeline_name}/cmd | Execute a pipeline command
[**Get**](PipelineApi.md#Get) | **Get** /pipeline/{pipeline_name} | Get pipeline
[**List**](PipelineApi.md#List) | **Get** /pipeline | List pipelines



## Cmd

> map[string]interface{} Cmd(ctx, pipelineName).Request(request).Execute()

Execute a pipeline command

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
    pipelineName := "pipelineName_example" // string | The name of the pipeline
    request := *openapiclient.NewCmdPipeline("Command_example") // CmdPipeline | Pipeline command.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.Cmd(context.Background(), pipelineName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.Cmd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Cmd`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.Cmd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineName** | **string** | The name of the pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiCmdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CmdPipeline**](CmdPipeline.md) | Pipeline command. | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> GetPipelineResponse Get(ctx, pipelineName).Execute()

Get pipeline



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
    pipelineName := "pipelineName_example" // string | The name of the pipeline

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.Get(context.Background(), pipelineName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: GetPipelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineName** | **string** | The name of the pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPipelineResponse**](GetPipelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListPipelineResponse List(ctx).Limit(limit).NextToken(nextToken).Execute()

List pipelines



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
    resp, r, err := apiClient.PipelineApi.List(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListPipelineResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.List`: %v\n", resp)
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

[**ListPipelineResponse**](ListPipelineResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

