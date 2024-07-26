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

// checks if the PipelineExecutionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineExecutionResponse{}

// PipelineExecutionResponse struct for PipelineExecutionResponse
type PipelineExecutionResponse struct {
	Errors []ModconfigStepError `json:"errors,omitempty"`
	Flowpipe *FlowpipeResponseMetadata `json:"flowpipe,omitempty"`
	Output map[string]interface{} `json:"output,omitempty"`
}
func (o PipelineExecutionResponse) GetResourceType() string {
	return "PipelineExecutionResponse"
}
// NewPipelineExecutionResponse instantiates a new PipelineExecutionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineExecutionResponse() *PipelineExecutionResponse {
	this := PipelineExecutionResponse{}
	return &this
}

// NewPipelineExecutionResponseWithDefaults instantiates a new PipelineExecutionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineExecutionResponseWithDefaults() *PipelineExecutionResponse {
	this := PipelineExecutionResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PipelineExecutionResponse) GetErrors() []ModconfigStepError {
	if o == nil || IsNil(o.Errors) {
		var ret []ModconfigStepError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineExecutionResponse) GetErrorsOk() ([]ModconfigStepError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PipelineExecutionResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ModconfigStepError and assigns it to the Errors field.
func (o *PipelineExecutionResponse) SetErrors(v []ModconfigStepError) {
	o.Errors = v
}

// GetFlowpipe returns the Flowpipe field value if set, zero value otherwise.
func (o *PipelineExecutionResponse) GetFlowpipe() FlowpipeResponseMetadata {
	if o == nil || IsNil(o.Flowpipe) {
		var ret FlowpipeResponseMetadata
		return ret
	}
	return *o.Flowpipe
}

// GetFlowpipeOk returns a tuple with the Flowpipe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineExecutionResponse) GetFlowpipeOk() (*FlowpipeResponseMetadata, bool) {
	if o == nil || IsNil(o.Flowpipe) {
		return nil, false
	}
	return o.Flowpipe, true
}

// HasFlowpipe returns a boolean if a field has been set.
func (o *PipelineExecutionResponse) HasFlowpipe() bool {
	if o != nil && !IsNil(o.Flowpipe) {
		return true
	}

	return false
}

// SetFlowpipe gets a reference to the given FlowpipeResponseMetadata and assigns it to the Flowpipe field.
func (o *PipelineExecutionResponse) SetFlowpipe(v FlowpipeResponseMetadata) {
	o.Flowpipe = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *PipelineExecutionResponse) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineExecutionResponse) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *PipelineExecutionResponse) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *PipelineExecutionResponse) SetOutput(v map[string]interface{}) {
	o.Output = v
}

func (o PipelineExecutionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineExecutionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Flowpipe) {
		toSerialize["flowpipe"] = o.Flowpipe
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullablePipelineExecutionResponse struct {
	value *PipelineExecutionResponse
	isSet bool
}

func (v NullablePipelineExecutionResponse) Get() *PipelineExecutionResponse {
	return v.value
}

func (v *NullablePipelineExecutionResponse) Set(val *PipelineExecutionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineExecutionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineExecutionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineExecutionResponse(val *PipelineExecutionResponse) *NullablePipelineExecutionResponse {
	return &NullablePipelineExecutionResponse{value: val, isSet: true}
}

func (v NullablePipelineExecutionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineExecutionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

