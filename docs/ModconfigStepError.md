# ModconfigStepError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**PerrErrorModel**](PerrErrorModel.md) |  | [optional] 
**Pipeline** | Pointer to **string** |  | [optional] 
**PipelineExecutionId** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**StepExecutionId** | Pointer to **string** |  | [optional] 

## Methods

### NewModconfigStepError

`func NewModconfigStepError() *ModconfigStepError`

NewModconfigStepError instantiates a new ModconfigStepError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigStepErrorWithDefaults

`func NewModconfigStepErrorWithDefaults() *ModconfigStepError`

NewModconfigStepErrorWithDefaults instantiates a new ModconfigStepError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ModconfigStepError) GetError() PerrErrorModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ModconfigStepError) GetErrorOk() (*PerrErrorModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ModconfigStepError) SetError(v PerrErrorModel)`

SetError sets Error field to given value.

### HasError

`func (o *ModconfigStepError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPipeline

`func (o *ModconfigStepError) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *ModconfigStepError) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *ModconfigStepError) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *ModconfigStepError) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineExecutionId

`func (o *ModconfigStepError) GetPipelineExecutionId() string`

GetPipelineExecutionId returns the PipelineExecutionId field if non-nil, zero value otherwise.

### GetPipelineExecutionIdOk

`func (o *ModconfigStepError) GetPipelineExecutionIdOk() (*string, bool)`

GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineExecutionId

`func (o *ModconfigStepError) SetPipelineExecutionId(v string)`

SetPipelineExecutionId sets PipelineExecutionId field to given value.

### HasPipelineExecutionId

`func (o *ModconfigStepError) HasPipelineExecutionId() bool`

HasPipelineExecutionId returns a boolean if a field has been set.

### GetStep

`func (o *ModconfigStepError) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ModconfigStepError) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ModconfigStepError) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *ModconfigStepError) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetStepExecutionId

`func (o *ModconfigStepError) GetStepExecutionId() string`

GetStepExecutionId returns the StepExecutionId field if non-nil, zero value otherwise.

### GetStepExecutionIdOk

`func (o *ModconfigStepError) GetStepExecutionIdOk() (*string, bool)`

GetStepExecutionIdOk returns a tuple with the StepExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutionId

`func (o *ModconfigStepError) SetStepExecutionId(v string)`

SetStepExecutionId sets StepExecutionId field to given value.

### HasStepExecutionId

`func (o *ModconfigStepError) HasStepExecutionId() bool`

HasStepExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


