# PipelineStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOn** | Pointer to **[]string** |  | [optional] 
**For** | Pointer to **string** |  | [optional] 
**InputTemplate** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPipelineStep

`func NewPipelineStep() *PipelineStep`

NewPipelineStep instantiates a new PipelineStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineStepWithDefaults

`func NewPipelineStepWithDefaults() *PipelineStep`

NewPipelineStepWithDefaults instantiates a new PipelineStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependsOn

`func (o *PipelineStep) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *PipelineStep) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *PipelineStep) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *PipelineStep) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetFor

`func (o *PipelineStep) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *PipelineStep) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *PipelineStep) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *PipelineStep) HasFor() bool`

HasFor returns a boolean if a field has been set.

### GetInputTemplate

`func (o *PipelineStep) GetInputTemplate() string`

GetInputTemplate returns the InputTemplate field if non-nil, zero value otherwise.

### GetInputTemplateOk

`func (o *PipelineStep) GetInputTemplateOk() (*string, bool)`

GetInputTemplateOk returns a tuple with the InputTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTemplate

`func (o *PipelineStep) SetInputTemplate(v string)`

SetInputTemplate sets InputTemplate field to given value.

### HasInputTemplate

`func (o *PipelineStep) HasInputTemplate() bool`

HasInputTemplate returns a boolean if a field has been set.

### GetName

`func (o *PipelineStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PipelineStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineStep) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PipelineStep) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


