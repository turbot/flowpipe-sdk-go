# Process

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** |  | [optional] 
**Pipeline** | Pointer to **string** |  | [optional] 

## Methods

### NewProcess

`func NewProcess() *Process`

NewProcess instantiates a new Process object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessWithDefaults

`func NewProcessWithDefaults() *Process`

NewProcessWithDefaults instantiates a new Process object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *Process) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Process) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Process) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Process) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetPipeline

`func (o *Process) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *Process) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *Process) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *Process) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


