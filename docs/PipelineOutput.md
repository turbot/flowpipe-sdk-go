# PipelineOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependsOn** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**Sensitive** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPipelineOutput

`func NewPipelineOutput() *PipelineOutput`

NewPipelineOutput instantiates a new PipelineOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineOutputWithDefaults

`func NewPipelineOutputWithDefaults() *PipelineOutput`

NewPipelineOutputWithDefaults instantiates a new PipelineOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependsOn

`func (o *PipelineOutput) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *PipelineOutput) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *PipelineOutput) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *PipelineOutput) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetName

`func (o *PipelineOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolved

`func (o *PipelineOutput) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *PipelineOutput) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *PipelineOutput) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *PipelineOutput) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetSensitive

`func (o *PipelineOutput) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PipelineOutput) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PipelineOutput) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *PipelineOutput) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *PipelineOutput) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PipelineOutput) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PipelineOutput) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PipelineOutput) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


