# GetPipelineResponse

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

### NewGetPipelineResponse

`func NewGetPipelineResponse() *GetPipelineResponse`

NewGetPipelineResponse instantiates a new GetPipelineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPipelineResponseWithDefaults

`func NewGetPipelineResponseWithDefaults() *GetPipelineResponse`

NewGetPipelineResponseWithDefaults instantiates a new GetPipelineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetPipelineResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPipelineResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPipelineResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPipelineResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *GetPipelineResponse) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *GetPipelineResponse) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *GetPipelineResponse) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *GetPipelineResponse) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetMod

`func (o *GetPipelineResponse) GetMod() string`

GetMod returns the Mod field if non-nil, zero value otherwise.

### GetModOk

`func (o *GetPipelineResponse) GetModOk() (*string, bool)`

GetModOk returns a tuple with the Mod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMod

`func (o *GetPipelineResponse) SetMod(v string)`

SetMod sets Mod field to given value.

### HasMod

`func (o *GetPipelineResponse) HasMod() bool`

HasMod returns a boolean if a field has been set.

### GetName

`func (o *GetPipelineResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPipelineResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPipelineResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPipelineResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *GetPipelineResponse) GetOutputs() []ModconfigPipelineOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *GetPipelineResponse) GetOutputsOk() (*[]ModconfigPipelineOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *GetPipelineResponse) SetOutputs(v []ModconfigPipelineOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *GetPipelineResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetParams

`func (o *GetPipelineResponse) GetParams() []FpPipelineParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetPipelineResponse) GetParamsOk() (*[]FpPipelineParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetPipelineResponse) SetParams(v []FpPipelineParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *GetPipelineResponse) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSteps

`func (o *GetPipelineResponse) GetSteps() []map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GetPipelineResponse) GetStepsOk() (*[]map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GetPipelineResponse) SetSteps(v []map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GetPipelineResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *GetPipelineResponse) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetPipelineResponse) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetPipelineResponse) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetPipelineResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *GetPipelineResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetPipelineResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetPipelineResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetPipelineResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


