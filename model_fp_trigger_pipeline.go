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

// checks if the FpTriggerPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FpTriggerPipeline{}

// FpTriggerPipeline struct for FpTriggerPipeline
type FpTriggerPipeline struct {
	CaptureGroup *string `json:"capture_group,omitempty"`
	Pipeline *string `json:"pipeline,omitempty"`
}
func (o FpTriggerPipeline) GetResourceType() string {
	return "FpTriggerPipeline"
}
// NewFpTriggerPipeline instantiates a new FpTriggerPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFpTriggerPipeline() *FpTriggerPipeline {
	this := FpTriggerPipeline{}
	return &this
}

// NewFpTriggerPipelineWithDefaults instantiates a new FpTriggerPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFpTriggerPipelineWithDefaults() *FpTriggerPipeline {
	this := FpTriggerPipeline{}
	return &this
}

// GetCaptureGroup returns the CaptureGroup field value if set, zero value otherwise.
func (o *FpTriggerPipeline) GetCaptureGroup() string {
	if o == nil || IsNil(o.CaptureGroup) {
		var ret string
		return ret
	}
	return *o.CaptureGroup
}

// GetCaptureGroupOk returns a tuple with the CaptureGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpTriggerPipeline) GetCaptureGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CaptureGroup) {
		return nil, false
	}
	return o.CaptureGroup, true
}

// HasCaptureGroup returns a boolean if a field has been set.
func (o *FpTriggerPipeline) HasCaptureGroup() bool {
	if o != nil && !IsNil(o.CaptureGroup) {
		return true
	}

	return false
}

// SetCaptureGroup gets a reference to the given string and assigns it to the CaptureGroup field.
func (o *FpTriggerPipeline) SetCaptureGroup(v string) {
	o.CaptureGroup = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *FpTriggerPipeline) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpTriggerPipeline) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *FpTriggerPipeline) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *FpTriggerPipeline) SetPipeline(v string) {
	o.Pipeline = &v
}

func (o FpTriggerPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FpTriggerPipeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaptureGroup) {
		toSerialize["capture_group"] = o.CaptureGroup
	}
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	return toSerialize, nil
}

type NullableFpTriggerPipeline struct {
	value *FpTriggerPipeline
	isSet bool
}

func (v NullableFpTriggerPipeline) Get() *FpTriggerPipeline {
	return v.value
}

func (v *NullableFpTriggerPipeline) Set(val *FpTriggerPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableFpTriggerPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableFpTriggerPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFpTriggerPipeline(val *FpTriggerPipeline) *NullableFpTriggerPipeline {
	return &NullableFpTriggerPipeline{value: val, isSet: true}
}

func (v NullableFpTriggerPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFpTriggerPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

