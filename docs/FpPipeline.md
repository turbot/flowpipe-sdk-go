# FpPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Mod** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]ModconfigPipelineOutput**](ModconfigPipelineOutput.md) |  | [optional] 
**Params** | Pointer to [**[]FpPipelineParam**](FpPipelineParam.md) |  | [optional] 
**Steps** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewFpPipeline

`func NewFpPipeline() *FpPipeline`

NewFpPipeline instantiates a new FpPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpPipelineWithDefaults

`func NewFpPipelineWithDefaults() *FpPipeline`

NewFpPipelineWithDefaults instantiates a new FpPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FpPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *FpPipeline) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *FpPipeline) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *FpPipeline) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *FpPipeline) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetMod

`func (o *FpPipeline) GetMod() string`

GetMod returns the Mod field if non-nil, zero value otherwise.

### GetModOk

`func (o *FpPipeline) GetModOk() (*string, bool)`

GetModOk returns a tuple with the Mod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMod

`func (o *FpPipeline) SetMod(v string)`

SetMod sets Mod field to given value.

### HasMod

`func (o *FpPipeline) HasMod() bool`

HasMod returns a boolean if a field has been set.

### GetName

`func (o *FpPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FpPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FpPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FpPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *FpPipeline) GetOutputs() []ModconfigPipelineOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *FpPipeline) GetOutputsOk() (*[]ModconfigPipelineOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *FpPipeline) SetOutputs(v []ModconfigPipelineOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *FpPipeline) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetParams

`func (o *FpPipeline) GetParams() []FpPipelineParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FpPipeline) GetParamsOk() (*[]FpPipelineParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FpPipeline) SetParams(v []FpPipelineParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *FpPipeline) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSteps

`func (o *FpPipeline) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FpPipeline) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FpPipeline) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *FpPipeline) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *FpPipeline) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FpPipeline) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FpPipeline) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FpPipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *FpPipeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FpPipeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FpPipeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FpPipeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


