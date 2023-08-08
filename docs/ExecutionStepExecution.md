# ExecutionStepExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | The name of the step in the pipeline definition | [optional] 
**NextStepAction** | Pointer to [**NextStepAction**](NextStepAction.md) |  | [optional] 
**Output** | Pointer to [**ExecutionStepExecutionOutput**](ExecutionStepExecutionOutput.md) |  | [optional] 
**PipelineExecutionId** | Pointer to **string** | Unique identifier for this step execution | [optional] 
**Status** | Pointer to **string** | The status of the step execution: \&quot;started\&quot;, \&quot;finished\&quot;, \&quot;failed\&quot;, \&quot;skipped\&quot; | [optional] 
**StepForEach** | Pointer to [**ExecutionStepExecutionStepForEach**](ExecutionStepExecutionStepForEach.md) |  | [optional] 

## Methods

### NewExecutionStepExecution

`func NewExecutionStepExecution() *ExecutionStepExecution`

NewExecutionStepExecution instantiates a new ExecutionStepExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepExecutionWithDefaults

`func NewExecutionStepExecutionWithDefaults() *ExecutionStepExecution`

NewExecutionStepExecutionWithDefaults instantiates a new ExecutionStepExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionStepExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionStepExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionStepExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionStepExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *ExecutionStepExecution) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ExecutionStepExecution) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ExecutionStepExecution) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ExecutionStepExecution) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetName

`func (o *ExecutionStepExecution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionStepExecution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionStepExecution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecutionStepExecution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextStepAction

`func (o *ExecutionStepExecution) GetNextStepAction() NextStepAction`

GetNextStepAction returns the NextStepAction field if non-nil, zero value otherwise.

### GetNextStepActionOk

`func (o *ExecutionStepExecution) GetNextStepActionOk() (*NextStepAction, bool)`

GetNextStepActionOk returns a tuple with the NextStepAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStepAction

`func (o *ExecutionStepExecution) SetNextStepAction(v NextStepAction)`

SetNextStepAction sets NextStepAction field to given value.

### HasNextStepAction

`func (o *ExecutionStepExecution) HasNextStepAction() bool`

HasNextStepAction returns a boolean if a field has been set.

### GetOutput

`func (o *ExecutionStepExecution) GetOutput() ExecutionStepExecutionOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ExecutionStepExecution) GetOutputOk() (*ExecutionStepExecutionOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ExecutionStepExecution) SetOutput(v ExecutionStepExecutionOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ExecutionStepExecution) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPipelineExecutionId

`func (o *ExecutionStepExecution) GetPipelineExecutionId() string`

GetPipelineExecutionId returns the PipelineExecutionId field if non-nil, zero value otherwise.

### GetPipelineExecutionIdOk

`func (o *ExecutionStepExecution) GetPipelineExecutionIdOk() (*string, bool)`

GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineExecutionId

`func (o *ExecutionStepExecution) SetPipelineExecutionId(v string)`

SetPipelineExecutionId sets PipelineExecutionId field to given value.

### HasPipelineExecutionId

`func (o *ExecutionStepExecution) HasPipelineExecutionId() bool`

HasPipelineExecutionId returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionStepExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionStepExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionStepExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionStepExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepForEach

`func (o *ExecutionStepExecution) GetStepForEach() ExecutionStepExecutionStepForEach`

GetStepForEach returns the StepForEach field if non-nil, zero value otherwise.

### GetStepForEachOk

`func (o *ExecutionStepExecution) GetStepForEachOk() (*ExecutionStepExecutionStepForEach, bool)`

GetStepForEachOk returns a tuple with the StepForEach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepForEach

`func (o *ExecutionStepExecution) SetStepForEach(v ExecutionStepExecutionStepForEach)`

SetStepForEach sets StepForEach field to given value.

### HasStepForEach

`func (o *ExecutionStepExecution) HasStepForEach() bool`

HasStepForEach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


