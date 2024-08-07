# CmdPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**ArgsString** | Pointer to **map[string]string** |  | [optional] 
**Command** | **string** |  | 
**ExecutionId** | Pointer to **string** | Sepcify execution id, if not specified, a new execution id will be created | [optional] 
**ExecutionMode** | Pointer to **string** |  | [optional] 
**WaitRetry** | Pointer to **int32** |  | [optional] 

## Methods

### NewCmdPipeline

`func NewCmdPipeline(command string, ) *CmdPipeline`

NewCmdPipeline instantiates a new CmdPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmdPipelineWithDefaults

`func NewCmdPipelineWithDefaults() *CmdPipeline`

NewCmdPipelineWithDefaults instantiates a new CmdPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *CmdPipeline) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CmdPipeline) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CmdPipeline) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CmdPipeline) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetArgsString

`func (o *CmdPipeline) GetArgsString() map[string]string`

GetArgsString returns the ArgsString field if non-nil, zero value otherwise.

### GetArgsStringOk

`func (o *CmdPipeline) GetArgsStringOk() (*map[string]string, bool)`

GetArgsStringOk returns a tuple with the ArgsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsString

`func (o *CmdPipeline) SetArgsString(v map[string]string)`

SetArgsString sets ArgsString field to given value.

### HasArgsString

`func (o *CmdPipeline) HasArgsString() bool`

HasArgsString returns a boolean if a field has been set.

### GetCommand

`func (o *CmdPipeline) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CmdPipeline) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CmdPipeline) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetExecutionId

`func (o *CmdPipeline) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *CmdPipeline) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *CmdPipeline) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *CmdPipeline) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CmdPipeline) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CmdPipeline) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CmdPipeline) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CmdPipeline) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetWaitRetry

`func (o *CmdPipeline) GetWaitRetry() int32`

GetWaitRetry returns the WaitRetry field if non-nil, zero value otherwise.

### GetWaitRetryOk

`func (o *CmdPipeline) GetWaitRetryOk() (*int32, bool)`

GetWaitRetryOk returns a tuple with the WaitRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitRetry

`func (o *CmdPipeline) SetWaitRetry(v int32)`

SetWaitRetry sets WaitRetry field to given value.

### HasWaitRetry

`func (o *CmdPipeline) HasWaitRetry() bool`

HasWaitRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


