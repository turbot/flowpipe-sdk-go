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

// checks if the ModconfigPipelineParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModconfigPipelineParam{}

// ModconfigPipelineParam struct for ModconfigPipelineParam
type ModconfigPipelineParam struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Optional *bool `json:"optional,omitempty"`
}
func (o ModconfigPipelineParam) GetResourceType() string {
	return "ModconfigPipelineParam"
}
// NewModconfigPipelineParam instantiates a new ModconfigPipelineParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModconfigPipelineParam() *ModconfigPipelineParam {
	this := ModconfigPipelineParam{}
	return &this
}

// NewModconfigPipelineParamWithDefaults instantiates a new ModconfigPipelineParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModconfigPipelineParamWithDefaults() *ModconfigPipelineParam {
	this := ModconfigPipelineParam{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModconfigPipelineParam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModconfigPipelineParam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModconfigPipelineParam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModconfigPipelineParam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModconfigPipelineParam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModconfigPipelineParam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModconfigPipelineParam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModconfigPipelineParam) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *ModconfigPipelineParam) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModconfigPipelineParam) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *ModconfigPipelineParam) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *ModconfigPipelineParam) SetOptional(v bool) {
	o.Optional = &v
}

func (o ModconfigPipelineParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModconfigPipelineParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	return toSerialize, nil
}

type NullableModconfigPipelineParam struct {
	value *ModconfigPipelineParam
	isSet bool
}

func (v NullableModconfigPipelineParam) Get() *ModconfigPipelineParam {
	return v.value
}

func (v *NullableModconfigPipelineParam) Set(val *ModconfigPipelineParam) {
	v.value = val
	v.isSet = true
}

func (v NullableModconfigPipelineParam) IsSet() bool {
	return v.isSet
}

func (v *NullableModconfigPipelineParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModconfigPipelineParam(val *ModconfigPipelineParam) *NullableModconfigPipelineParam {
	return &NullableModconfigPipelineParam{value: val, isSet: true}
}

func (v NullableModconfigPipelineParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModconfigPipelineParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


