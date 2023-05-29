# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**Parallel** | Pointer to **bool** |  | [optional] 
**Steps** | Pointer to [**map[string]PipelineStep**](PipelineStep.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPipeline

`func NewPipeline() *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *Pipeline) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Pipeline) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Pipeline) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Pipeline) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetName

`func (o *Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *Pipeline) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Pipeline) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Pipeline) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Pipeline) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetParallel

`func (o *Pipeline) GetParallel() bool`

GetParallel returns the Parallel field if non-nil, zero value otherwise.

### GetParallelOk

`func (o *Pipeline) GetParallelOk() (*bool, bool)`

GetParallelOk returns a tuple with the Parallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel

`func (o *Pipeline) SetParallel(v bool)`

SetParallel sets Parallel field to given value.

### HasParallel

`func (o *Pipeline) HasParallel() bool`

HasParallel returns a boolean if a field has been set.

### GetSteps

`func (o *Pipeline) GetSteps() map[string]PipelineStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Pipeline) GetStepsOk() (*map[string]PipelineStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Pipeline) SetSteps(v map[string]PipelineStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Pipeline) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetType

`func (o *Pipeline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Pipeline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Pipeline) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Pipeline) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


