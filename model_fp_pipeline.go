/*
Flowpipe

Flowpipe is workflow and pipelines for DevSecOps.

API version: 0.1.0
Contact: info@flowpipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flowpipeapi

import (
	"encoding/json"
)

// checks if the FpPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FpPipeline{}

// FpPipeline struct for FpPipeline
type FpPipeline struct {
	Description *string `json:"description,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Mod *string `json:"mod,omitempty"`
	Name *string `json:"name,omitempty"`
	Outputs []ModconfigPipelineOutput `json:"outputs,omitempty"`
	Params []FpPipelineParam `json:"params,omitempty"`
	Steps []map[string]interface{} `json:"steps,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
}
func (o FpPipeline) GetResourceType() string {
	return "FpPipeline"
}
// NewFpPipeline instantiates a new FpPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFpPipeline() *FpPipeline {
	this := FpPipeline{}
	return &this
}

// NewFpPipelineWithDefaults instantiates a new FpPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFpPipelineWithDefaults() *FpPipeline {
	this := FpPipeline{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FpPipeline) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FpPipeline) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FpPipeline) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *FpPipeline) GetDocumentation() string {
	if o == nil || IsNil(o.Documentation) {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *FpPipeline) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *FpPipeline) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetMod returns the Mod field value if set, zero value otherwise.
func (o *FpPipeline) GetMod() string {
	if o == nil || IsNil(o.Mod) {
		var ret string
		return ret
	}
	return *o.Mod
}

// GetModOk returns a tuple with the Mod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetModOk() (*string, bool) {
	if o == nil || IsNil(o.Mod) {
		return nil, false
	}
	return o.Mod, true
}

// HasMod returns a boolean if a field has been set.
func (o *FpPipeline) HasMod() bool {
	if o != nil && !IsNil(o.Mod) {
		return true
	}

	return false
}

// SetMod gets a reference to the given string and assigns it to the Mod field.
func (o *FpPipeline) SetMod(v string) {
	o.Mod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FpPipeline) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FpPipeline) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FpPipeline) SetName(v string) {
	o.Name = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *FpPipeline) GetOutputs() []ModconfigPipelineOutput {
	if o == nil || IsNil(o.Outputs) {
		var ret []ModconfigPipelineOutput
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetOutputsOk() ([]ModconfigPipelineOutput, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *FpPipeline) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []ModconfigPipelineOutput and assigns it to the Outputs field.
func (o *FpPipeline) SetOutputs(v []ModconfigPipelineOutput) {
	o.Outputs = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *FpPipeline) GetParams() []FpPipelineParam {
	if o == nil || IsNil(o.Params) {
		var ret []FpPipelineParam
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetParamsOk() ([]FpPipelineParam, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *FpPipeline) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []FpPipelineParam and assigns it to the Params field.
func (o *FpPipeline) SetParams(v []FpPipelineParam) {
	o.Params = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *FpPipeline) GetSteps() []map[string]interface{} {
	if o == nil || IsNil(o.Steps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetStepsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *FpPipeline) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []map[string]interface{} and assigns it to the Steps field.
func (o *FpPipeline) SetSteps(v []map[string]interface{}) {
	o.Steps = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FpPipeline) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FpPipeline) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *FpPipeline) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FpPipeline) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipeline) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FpPipeline) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FpPipeline) SetTitle(v string) {
	o.Title = &v
}

func (o FpPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FpPipeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	if !IsNil(o.Mod) {
		toSerialize["mod"] = o.Mod
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableFpPipeline struct {
	value *FpPipeline
	isSet bool
}

func (v NullableFpPipeline) Get() *FpPipeline {
	return v.value
}

func (v *NullableFpPipeline) Set(val *FpPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableFpPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableFpPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFpPipeline(val *FpPipeline) *NullableFpPipeline {
	return &NullableFpPipeline{value: val, isSet: true}
}

func (v NullableFpPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFpPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


