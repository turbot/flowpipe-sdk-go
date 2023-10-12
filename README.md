# Go SDK for Flowpipe

Flowpipe is a low-code workflow automation tool that aims to be simple yet powerful.

For help on getting started with Flowpipe, please visit https://flowpipe.io

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://www.flowpipe.io](http://www.flowpipe.io)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import flowpipeapi "github.com/turbot/flowpipe-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), flowpipeapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), flowpipeapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), flowpipeapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), flowpipeapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api/v0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PipelineApi* | [**Cmd**](docs/PipelineApi.md#cmd) | **Post** /pipeline/{pipeline_name}/cmd | Execute a pipeline command
*PipelineApi* | [**Get**](docs/PipelineApi.md#get) | **Get** /pipeline/{pipeline_name} | Get pipeline
*PipelineApi* | [**List**](docs/PipelineApi.md#list) | **Get** /pipeline | List pipelines
*ProcessApi* | [**Get**](docs/ProcessApi.md#get) | **Get** /process/{process_id} | Get process
*ProcessApi* | [**GetOutput**](docs/ProcessApi.md#getoutput) | **Get** /process/{process_id}/output | Get process output
*ProcessApi* | [**GetSnapshot**](docs/ProcessApi.md#getsnapshot) | **Get** /process/:process_id/log/process.sps | Get process snapshot
*ProcessApi* | [**List**](docs/ProcessApi.md#list) | **Get** /process | List processs
*TriggerApi* | [**Get**](docs/TriggerApi.md#get) | **Get** /trigger/{trigger_name} | Get trigger
*TriggerApi* | [**List**](docs/TriggerApi.md#list) | **Get** /trigger | List triggers
*VariableApi* | [**Get**](docs/VariableApi.md#get) | **Get** /variable/{variable_name} | Get variable
*VariableApi* | [**List**](docs/VariableApi.md#list) | **Get** /variable | List variables


## Documentation For Models

 - [CmdPipeline](docs/CmdPipeline.md)
 - [ExecutionSnapshot](docs/ExecutionSnapshot.md)
 - [ExecutionSnapshotLayout](docs/ExecutionSnapshotLayout.md)
 - [ExecutionSnapshotPanel](docs/ExecutionSnapshotPanel.md)
 - [ExecutionSnapshotPanelData](docs/ExecutionSnapshotPanelData.md)
 - [ExecutionSnapshotPanelDataColumn](docs/ExecutionSnapshotPanelDataColumn.md)
 - [FpPipelineParam](docs/FpPipelineParam.md)
 - [FpTrigger](docs/FpTrigger.md)
 - [GetPipelineResponse](docs/GetPipelineResponse.md)
 - [ListPipelineResponse](docs/ListPipelineResponse.md)
 - [ListPipelineResponseItem](docs/ListPipelineResponseItem.md)
 - [ListProcessResponse](docs/ListProcessResponse.md)
 - [ListTriggerResponse](docs/ListTriggerResponse.md)
 - [ListVariableResponse](docs/ListVariableResponse.md)
 - [ModconfigPipelineOutput](docs/ModconfigPipelineOutput.md)
 - [PerrErrorDetailModel](docs/PerrErrorDetailModel.md)
 - [PerrErrorModel](docs/PerrErrorModel.md)
 - [Process](docs/Process.md)
 - [ProcessOutputData](docs/ProcessOutputData.md)
 - [Variable](docs/Variable.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@flowpipe.io

