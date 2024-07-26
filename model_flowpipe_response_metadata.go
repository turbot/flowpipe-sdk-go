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

// checks if the FlowpipeResponseMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowpipeResponseMetadata{}

// FlowpipeResponseMetadata struct for FlowpipeResponseMetadata
type FlowpipeResponseMetadata struct {
	ExecutionId *string `json:"execution_id,omitempty"`
	IsStale *bool `json:"is_stale,omitempty"`
	LastLoaded *string `json:"last_loaded,omitempty"`
	Pipeline *string `json:"pipeline,omitempty"`
	PipelineExecutionId *string `json:"pipeline_execution_id,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
}
func (o FlowpipeResponseMetadata) GetResourceType() string {
	return "FlowpipeResponseMetadata"
}
// NewFlowpipeResponseMetadata instantiates a new FlowpipeResponseMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowpipeResponseMetadata() *FlowpipeResponseMetadata {
	this := FlowpipeResponseMetadata{}
	return &this
}

// NewFlowpipeResponseMetadataWithDefaults instantiates a new FlowpipeResponseMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowpipeResponseMetadataWithDefaults() *FlowpipeResponseMetadata {
	this := FlowpipeResponseMetadata{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *FlowpipeResponseMetadata) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetIsStale returns the IsStale field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetIsStale() bool {
	if o == nil || IsNil(o.IsStale) {
		var ret bool
		return ret
	}
	return *o.IsStale
}

// GetIsStaleOk returns a tuple with the IsStale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetIsStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStale) {
		return nil, false
	}
	return o.IsStale, true
}

// HasIsStale returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasIsStale() bool {
	if o != nil && !IsNil(o.IsStale) {
		return true
	}

	return false
}

// SetIsStale gets a reference to the given bool and assigns it to the IsStale field.
func (o *FlowpipeResponseMetadata) SetIsStale(v bool) {
	o.IsStale = &v
}

// GetLastLoaded returns the LastLoaded field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetLastLoaded() string {
	if o == nil || IsNil(o.LastLoaded) {
		var ret string
		return ret
	}
	return *o.LastLoaded
}

// GetLastLoadedOk returns a tuple with the LastLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetLastLoadedOk() (*string, bool) {
	if o == nil || IsNil(o.LastLoaded) {
		return nil, false
	}
	return o.LastLoaded, true
}

// HasLastLoaded returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasLastLoaded() bool {
	if o != nil && !IsNil(o.LastLoaded) {
		return true
	}

	return false
}

// SetLastLoaded gets a reference to the given string and assigns it to the LastLoaded field.
func (o *FlowpipeResponseMetadata) SetLastLoaded(v string) {
	o.LastLoaded = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *FlowpipeResponseMetadata) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetPipelineExecutionId returns the PipelineExecutionId field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetPipelineExecutionId() string {
	if o == nil || IsNil(o.PipelineExecutionId) {
		var ret string
		return ret
	}
	return *o.PipelineExecutionId
}

// GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetPipelineExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PipelineExecutionId) {
		return nil, false
	}
	return o.PipelineExecutionId, true
}

// HasPipelineExecutionId returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasPipelineExecutionId() bool {
	if o != nil && !IsNil(o.PipelineExecutionId) {
		return true
	}

	return false
}

// SetPipelineExecutionId gets a reference to the given string and assigns it to the PipelineExecutionId field.
func (o *FlowpipeResponseMetadata) SetPipelineExecutionId(v string) {
	o.PipelineExecutionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FlowpipeResponseMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlowpipeResponseMetadata) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowpipeResponseMetadata) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlowpipeResponseMetadata) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlowpipeResponseMetadata) SetType(v string) {
	o.Type = &v
}

func (o FlowpipeResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowpipeResponseMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.IsStale) {
		toSerialize["is_stale"] = o.IsStale
	}
	if !IsNil(o.LastLoaded) {
		toSerialize["last_loaded"] = o.LastLoaded
	}
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	if !IsNil(o.PipelineExecutionId) {
		toSerialize["pipeline_execution_id"] = o.PipelineExecutionId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFlowpipeResponseMetadata struct {
	value *FlowpipeResponseMetadata
	isSet bool
}

func (v NullableFlowpipeResponseMetadata) Get() *FlowpipeResponseMetadata {
	return v.value
}

func (v *NullableFlowpipeResponseMetadata) Set(val *FlowpipeResponseMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowpipeResponseMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowpipeResponseMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowpipeResponseMetadata(val *FlowpipeResponseMetadata) *NullableFlowpipeResponseMetadata {
	return &NullableFlowpipeResponseMetadata{value: val, isSet: true}
}

func (v NullableFlowpipeResponseMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowpipeResponseMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

