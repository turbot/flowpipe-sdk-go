# PipelineExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ModconfigStepError**](ModconfigStepError.md) |  | [optional] 
**Flowpipe** | Pointer to [**FlowpipeResponseMetadata**](FlowpipeResponseMetadata.md) |  | [optional] 
**Results** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPipelineExecutionResponse

`func NewPipelineExecutionResponse() *PipelineExecutionResponse`

NewPipelineExecutionResponse instantiates a new PipelineExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineExecutionResponseWithDefaults

`func NewPipelineExecutionResponseWithDefaults() *PipelineExecutionResponse`

NewPipelineExecutionResponseWithDefaults instantiates a new PipelineExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *PipelineExecutionResponse) GetErrors() []ModconfigStepError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PipelineExecutionResponse) GetErrorsOk() (*[]ModconfigStepError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PipelineExecutionResponse) SetErrors(v []ModconfigStepError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PipelineExecutionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFlowpipe

`func (o *PipelineExecutionResponse) GetFlowpipe() FlowpipeResponseMetadata`

GetFlowpipe returns the Flowpipe field if non-nil, zero value otherwise.

### GetFlowpipeOk

`func (o *PipelineExecutionResponse) GetFlowpipeOk() (*FlowpipeResponseMetadata, bool)`

GetFlowpipeOk returns a tuple with the Flowpipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowpipe

`func (o *PipelineExecutionResponse) SetFlowpipe(v FlowpipeResponseMetadata)`

SetFlowpipe sets Flowpipe field to given value.

### HasFlowpipe

`func (o *PipelineExecutionResponse) HasFlowpipe() bool`

HasFlowpipe returns a boolean if a field has been set.

### GetResults

`func (o *PipelineExecutionResponse) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PipelineExecutionResponse) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PipelineExecutionResponse) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *PipelineExecutionResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


