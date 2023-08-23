# ExecutionPipelineExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]ModconfigStepError**](ModconfigStepError.md) | All errors from the step execution + any errors that can be added to the pipeline execution manually | [optional] 
**Id** | Pointer to **string** | Unique identifier for this pipeline execution | [optional] 
**Name** | Pointer to **string** | The name of the pipeline | [optional] 
**ParentExecutionId** | Pointer to **string** |  | [optional] 
**ParentStepExecutionId** | Pointer to **string** | If this is a child pipeline, then track it&#39;s parent | [optional] 
**PipelineOutput** | Pointer to **map[string]interface{}** | The output of the pipeline | [optional] 
**Status** | Pointer to **string** | The status of the pipeline execution: queued, planned, started, completed, failed | [optional] 
**StepExecutions** | Pointer to [**map[string]ExecutionStepExecution**](ExecutionStepExecution.md) | Steps triggered by pipelines in the execution. | [optional] 

## Methods

### NewExecutionPipelineExecution

`func NewExecutionPipelineExecution() *ExecutionPipelineExecution`

NewExecutionPipelineExecution instantiates a new ExecutionPipelineExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionPipelineExecutionWithDefaults

`func NewExecutionPipelineExecutionWithDefaults() *ExecutionPipelineExecution`

NewExecutionPipelineExecutionWithDefaults instantiates a new ExecutionPipelineExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *ExecutionPipelineExecution) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ExecutionPipelineExecution) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ExecutionPipelineExecution) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ExecutionPipelineExecution) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetErrors

`func (o *ExecutionPipelineExecution) GetErrors() []ModconfigStepError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExecutionPipelineExecution) GetErrorsOk() (*[]ModconfigStepError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExecutionPipelineExecution) SetErrors(v []ModconfigStepError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExecutionPipelineExecution) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetId

`func (o *ExecutionPipelineExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionPipelineExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionPipelineExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionPipelineExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExecutionPipelineExecution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionPipelineExecution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionPipelineExecution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecutionPipelineExecution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentExecutionId

`func (o *ExecutionPipelineExecution) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *ExecutionPipelineExecution) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *ExecutionPipelineExecution) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *ExecutionPipelineExecution) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### GetParentStepExecutionId

`func (o *ExecutionPipelineExecution) GetParentStepExecutionId() string`

GetParentStepExecutionId returns the ParentStepExecutionId field if non-nil, zero value otherwise.

### GetParentStepExecutionIdOk

`func (o *ExecutionPipelineExecution) GetParentStepExecutionIdOk() (*string, bool)`

GetParentStepExecutionIdOk returns a tuple with the ParentStepExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepExecutionId

`func (o *ExecutionPipelineExecution) SetParentStepExecutionId(v string)`

SetParentStepExecutionId sets ParentStepExecutionId field to given value.

### HasParentStepExecutionId

`func (o *ExecutionPipelineExecution) HasParentStepExecutionId() bool`

HasParentStepExecutionId returns a boolean if a field has been set.

### GetPipelineOutput

`func (o *ExecutionPipelineExecution) GetPipelineOutput() map[string]interface{}`

GetPipelineOutput returns the PipelineOutput field if non-nil, zero value otherwise.

### GetPipelineOutputOk

`func (o *ExecutionPipelineExecution) GetPipelineOutputOk() (*map[string]interface{}, bool)`

GetPipelineOutputOk returns a tuple with the PipelineOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineOutput

`func (o *ExecutionPipelineExecution) SetPipelineOutput(v map[string]interface{})`

SetPipelineOutput sets PipelineOutput field to given value.

### HasPipelineOutput

`func (o *ExecutionPipelineExecution) HasPipelineOutput() bool`

HasPipelineOutput returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionPipelineExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionPipelineExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionPipelineExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionPipelineExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepExecutions

`func (o *ExecutionPipelineExecution) GetStepExecutions() map[string]ExecutionStepExecution`

GetStepExecutions returns the StepExecutions field if non-nil, zero value otherwise.

### GetStepExecutionsOk

`func (o *ExecutionPipelineExecution) GetStepExecutionsOk() (*map[string]ExecutionStepExecution, bool)`

GetStepExecutionsOk returns a tuple with the StepExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutions

`func (o *ExecutionPipelineExecution) SetStepExecutions(v map[string]ExecutionStepExecution)`

SetStepExecutions sets StepExecutions field to given value.

### HasStepExecutions

`func (o *ExecutionPipelineExecution) HasStepExecutions() bool`

HasStepExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


