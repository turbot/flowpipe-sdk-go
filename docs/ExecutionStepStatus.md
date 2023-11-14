# ExecutionStepStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorHold** | Pointer to **bool** | Indicates that a step is in retry loop so we don&#39;t mark it as failed | [optional] 
**Failed** | Pointer to **map[string]bool** | Step executions that are failed. | [optional] 
**Finished** | Pointer to **map[string]bool** | Step executions that are finished. | [optional] 
**Initializing** | Pointer to **bool** | When the step is initializing it doesn&#39;t yet have any executions. We track it as initializing until the first execution is queued. | [optional] 
**LoopHold** | Pointer to **bool** | Both LoopHold and ErrorHold must be resolved **before** the \&quot;finish\&quot; event is called, i.e. it needs to be calculated at the end of \&quot;step start command\&quot; and \&quot;step pipeline finish\&quot; command.  It can&#39;t be calculated at the \&quot;finish\&quot; event because it&#39;s already too late. If the planner see that it has an finish event without either a LoopHold or ErrorHold, it will mark the step as completed or failed  Indicates that step is in a loop so we don&#39;t mark it as finished | [optional] 
**OverralState** | Pointer to **string** |  | [optional] 
**Queued** | Pointer to **map[string]bool** | Step executions that are queued. | [optional] 
**Started** | Pointer to **map[string]bool** | Step executions that are started. | [optional] 
**StepExecutions** | Pointer to [**[]ExecutionStepExecution**](ExecutionStepExecution.md) | There&#39;s the step execution in execution, this is the same but in a list for a given step status The element in this slice should point to the same element in the StepExecutions map (in PipelineExecution) | [optional] 

## Methods

### NewExecutionStepStatus

`func NewExecutionStepStatus() *ExecutionStepStatus`

NewExecutionStepStatus instantiates a new ExecutionStepStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepStatusWithDefaults

`func NewExecutionStepStatusWithDefaults() *ExecutionStepStatus`

NewExecutionStepStatusWithDefaults instantiates a new ExecutionStepStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorHold

`func (o *ExecutionStepStatus) GetErrorHold() bool`

GetErrorHold returns the ErrorHold field if non-nil, zero value otherwise.

### GetErrorHoldOk

`func (o *ExecutionStepStatus) GetErrorHoldOk() (*bool, bool)`

GetErrorHoldOk returns a tuple with the ErrorHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHold

`func (o *ExecutionStepStatus) SetErrorHold(v bool)`

SetErrorHold sets ErrorHold field to given value.

### HasErrorHold

`func (o *ExecutionStepStatus) HasErrorHold() bool`

HasErrorHold returns a boolean if a field has been set.

### GetFailed

`func (o *ExecutionStepStatus) GetFailed() map[string]bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ExecutionStepStatus) GetFailedOk() (*map[string]bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ExecutionStepStatus) SetFailed(v map[string]bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ExecutionStepStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFinished

`func (o *ExecutionStepStatus) GetFinished() map[string]bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ExecutionStepStatus) GetFinishedOk() (*map[string]bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ExecutionStepStatus) SetFinished(v map[string]bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *ExecutionStepStatus) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetInitializing

`func (o *ExecutionStepStatus) GetInitializing() bool`

GetInitializing returns the Initializing field if non-nil, zero value otherwise.

### GetInitializingOk

`func (o *ExecutionStepStatus) GetInitializingOk() (*bool, bool)`

GetInitializingOk returns a tuple with the Initializing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializing

`func (o *ExecutionStepStatus) SetInitializing(v bool)`

SetInitializing sets Initializing field to given value.

### HasInitializing

`func (o *ExecutionStepStatus) HasInitializing() bool`

HasInitializing returns a boolean if a field has been set.

### GetLoopHold

`func (o *ExecutionStepStatus) GetLoopHold() bool`

GetLoopHold returns the LoopHold field if non-nil, zero value otherwise.

### GetLoopHoldOk

`func (o *ExecutionStepStatus) GetLoopHoldOk() (*bool, bool)`

GetLoopHoldOk returns a tuple with the LoopHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopHold

`func (o *ExecutionStepStatus) SetLoopHold(v bool)`

SetLoopHold sets LoopHold field to given value.

### HasLoopHold

`func (o *ExecutionStepStatus) HasLoopHold() bool`

HasLoopHold returns a boolean if a field has been set.

### GetOverralState

`func (o *ExecutionStepStatus) GetOverralState() string`

GetOverralState returns the OverralState field if non-nil, zero value otherwise.

### GetOverralStateOk

`func (o *ExecutionStepStatus) GetOverralStateOk() (*string, bool)`

GetOverralStateOk returns a tuple with the OverralState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverralState

`func (o *ExecutionStepStatus) SetOverralState(v string)`

SetOverralState sets OverralState field to given value.

### HasOverralState

`func (o *ExecutionStepStatus) HasOverralState() bool`

HasOverralState returns a boolean if a field has been set.

### GetQueued

`func (o *ExecutionStepStatus) GetQueued() map[string]bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ExecutionStepStatus) GetQueuedOk() (*map[string]bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ExecutionStepStatus) SetQueued(v map[string]bool)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *ExecutionStepStatus) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetStarted

`func (o *ExecutionStepStatus) GetStarted() map[string]bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ExecutionStepStatus) GetStartedOk() (*map[string]bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ExecutionStepStatus) SetStarted(v map[string]bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *ExecutionStepStatus) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStepExecutions

`func (o *ExecutionStepStatus) GetStepExecutions() []ExecutionStepExecution`

GetStepExecutions returns the StepExecutions field if non-nil, zero value otherwise.

### GetStepExecutionsOk

`func (o *ExecutionStepStatus) GetStepExecutionsOk() (*[]ExecutionStepExecution, bool)`

GetStepExecutionsOk returns a tuple with the StepExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutions

`func (o *ExecutionStepStatus) SetStepExecutions(v []ExecutionStepExecution)`

SetStepExecutions sets StepExecutions field to given value.

### HasStepExecutions

`func (o *ExecutionStepStatus) HasStepExecutions() bool`

HasStepExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


