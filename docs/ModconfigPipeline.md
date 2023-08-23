# ModconfigPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]ModconfigPipelineOutput**](ModconfigPipelineOutput.md) |  | [optional] 
**PipelineName** | Pointer to **string** | TODO: hack to serialise pipeline name because HclResourceImpl is not serialised | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewModconfigPipeline

`func NewModconfigPipeline() *ModconfigPipeline`

NewModconfigPipeline instantiates a new ModconfigPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigPipelineWithDefaults

`func NewModconfigPipelineWithDefaults() *ModconfigPipeline`

NewModconfigPipelineWithDefaults instantiates a new ModconfigPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModconfigPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModconfigPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModconfigPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModconfigPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *ModconfigPipeline) GetOutputs() []ModconfigPipelineOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ModconfigPipeline) GetOutputsOk() (*[]ModconfigPipelineOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ModconfigPipeline) SetOutputs(v []ModconfigPipelineOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ModconfigPipeline) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPipelineName

`func (o *ModconfigPipeline) GetPipelineName() string`

GetPipelineName returns the PipelineName field if non-nil, zero value otherwise.

### GetPipelineNameOk

`func (o *ModconfigPipeline) GetPipelineNameOk() (*string, bool)`

GetPipelineNameOk returns a tuple with the PipelineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineName

`func (o *ModconfigPipeline) SetPipelineName(v string)`

SetPipelineName sets PipelineName field to given value.

### HasPipelineName

`func (o *ModconfigPipeline) HasPipelineName() bool`

HasPipelineName returns a boolean if a field has been set.

### GetSteps

`func (o *ModconfigPipeline) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ModconfigPipeline) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ModconfigPipeline) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ModconfigPipeline) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


