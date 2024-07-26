# TriggerExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flowpipe** | Pointer to [**FlowpipeTriggerResponseMetadata**](FlowpipeTriggerResponseMetadata.md) |  | [optional] 
**Results** | Pointer to [**[]PipelineExecutionResponse**](PipelineExecutionResponse.md) |  | [optional] 

## Methods

### NewTriggerExecutionResponse

`func NewTriggerExecutionResponse() *TriggerExecutionResponse`

NewTriggerExecutionResponse instantiates a new TriggerExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerExecutionResponseWithDefaults

`func NewTriggerExecutionResponseWithDefaults() *TriggerExecutionResponse`

NewTriggerExecutionResponseWithDefaults instantiates a new TriggerExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowpipe

`func (o *TriggerExecutionResponse) GetFlowpipe() FlowpipeTriggerResponseMetadata`

GetFlowpipe returns the Flowpipe field if non-nil, zero value otherwise.

### GetFlowpipeOk

`func (o *TriggerExecutionResponse) GetFlowpipeOk() (*FlowpipeTriggerResponseMetadata, bool)`

GetFlowpipeOk returns a tuple with the Flowpipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowpipe

`func (o *TriggerExecutionResponse) SetFlowpipe(v FlowpipeTriggerResponseMetadata)`

SetFlowpipe sets Flowpipe field to given value.

### HasFlowpipe

`func (o *TriggerExecutionResponse) HasFlowpipe() bool`

HasFlowpipe returns a boolean if a field has been set.

### GetResults

`func (o *TriggerExecutionResponse) GetResults() []PipelineExecutionResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TriggerExecutionResponse) GetResultsOk() (*[]PipelineExecutionResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TriggerExecutionResponse) SetResults(v []PipelineExecutionResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *TriggerExecutionResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


