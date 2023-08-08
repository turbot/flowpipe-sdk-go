# ExecutionExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this execution. | [optional] 
**PipelineExecutions** | Pointer to [**map[string]ExecutionPipelineExecution**](ExecutionPipelineExecution.md) | Pipelines triggered by the execution. Even if the pipelines are nested, we maintain a flat list of all pipelines for easy lookup and querying. | [optional] 

## Methods

### NewExecutionExecution

`func NewExecutionExecution() *ExecutionExecution`

NewExecutionExecution instantiates a new ExecutionExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionExecutionWithDefaults

`func NewExecutionExecutionWithDefaults() *ExecutionExecution`

NewExecutionExecutionWithDefaults instantiates a new ExecutionExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPipelineExecutions

`func (o *ExecutionExecution) GetPipelineExecutions() map[string]ExecutionPipelineExecution`

GetPipelineExecutions returns the PipelineExecutions field if non-nil, zero value otherwise.

### GetPipelineExecutionsOk

`func (o *ExecutionExecution) GetPipelineExecutionsOk() (*map[string]ExecutionPipelineExecution, bool)`

GetPipelineExecutionsOk returns a tuple with the PipelineExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineExecutions

`func (o *ExecutionExecution) SetPipelineExecutions(v map[string]ExecutionPipelineExecution)`

SetPipelineExecutions sets PipelineExecutions field to given value.

### HasPipelineExecutions

`func (o *ExecutionExecution) HasPipelineExecutions() bool`

HasPipelineExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


