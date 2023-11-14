# ExecutionPipelineExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ModconfigStepError**](ModconfigStepError.md) | All errors from the step execution + any errors that can be added to the pipeline execution manually | [optional] 
**Id** | Pointer to **string** | Unique identifier for this pipeline execution | [optional] 
**Name** | Pointer to **string** | The name of the pipeline | [optional] 
**ParentExecutionId** | Pointer to **string** |  | [optional] 
**ParentStepExecutionId** | Pointer to **string** | If this is a child pipeline, then track it&#39;s parent | [optional] 
**PipelineOutput** | Pointer to **map[string]interface{}** | The output of the pipeline | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The status of the pipeline execution: queued, planned, started, completed, failed | [optional] 
**StepStatus** | Pointer to [**map[string]map[string]ExecutionStepStatus**](map.md) | Status of each step on a per-step index basis. Used to determine if dependencies have been met etc. Note that each step may have multiple executions, the status of which are not tracked here. dependencies have been met, etc.  The Step Status used to be per-step, however the addition of for_each means that we now need to expand this tracking to include the \&quot;index\&quot; of the step  for_each have 2 type of results: list or map, however in Flowpipe they are both treated as a map, the list is simply a map that the key happens to be a string of \&quot;0\&quot;, \&quot;1\&quot;, \&quot;2\&quot;    The data structure of StepStatus is as follow:   {    \&quot;echo.echo\&quot;: {     \&quot;0\&quot;: {      xyz     },     \&quot;1\&quot;: {      xyz     }    },    \&quot;http.one\&quot;: {     \&quot;foo\&quot;: {      zzz     },     \&quot;bar\&quot;: {      yyy     }    }   }    echo.echo has a for_each which is a list, so the key is the index of the list    http.one has a for_each which is a map, so the key is the key of the map    LOOP    Loop will be recorded in StepStatus.StepExecution, it&#39;s an array   * | [optional] 

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

### GetEndTime

`func (o *ExecutionPipelineExecution) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExecutionPipelineExecution) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExecutionPipelineExecution) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ExecutionPipelineExecution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

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

### GetStartTime

`func (o *ExecutionPipelineExecution) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExecutionPipelineExecution) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExecutionPipelineExecution) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ExecutionPipelineExecution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

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

### GetStepStatus

`func (o *ExecutionPipelineExecution) GetStepStatus() map[string]map[string]ExecutionStepStatus`

GetStepStatus returns the StepStatus field if non-nil, zero value otherwise.

### GetStepStatusOk

`func (o *ExecutionPipelineExecution) GetStepStatusOk() (*map[string]map[string]ExecutionStepStatus, bool)`

GetStepStatusOk returns a tuple with the StepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepStatus

`func (o *ExecutionPipelineExecution) SetStepStatus(v map[string]map[string]ExecutionStepStatus)`

SetStepStatus sets StepStatus field to given value.

### HasStepStatus

`func (o *ExecutionPipelineExecution) HasStepStatus() bool`

HasStepStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


