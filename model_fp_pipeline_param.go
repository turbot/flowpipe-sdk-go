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

// checks if the FpPipelineParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FpPipelineParam{}

// FpPipelineParam struct for FpPipelineParam
type FpPipelineParam struct {
	Default map[string]interface{} `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Optional *bool `json:"optional,omitempty"`
	Type *string `json:"type,omitempty"`
}
func (o FpPipelineParam) GetResourceType() string {
	return "FpPipelineParam"
}
// NewFpPipelineParam instantiates a new FpPipelineParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFpPipelineParam() *FpPipelineParam {
	this := FpPipelineParam{}
	return &this
}

// NewFpPipelineParamWithDefaults instantiates a new FpPipelineParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFpPipelineParamWithDefaults() *FpPipelineParam {
	this := FpPipelineParam{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *FpPipelineParam) GetDefault() map[string]interface{} {
	if o == nil || IsNil(o.Default) {
		var ret map[string]interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipelineParam) GetDefaultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return map[string]interface{}{}, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *FpPipelineParam) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given map[string]interface{} and assigns it to the Default field.
func (o *FpPipelineParam) SetDefault(v map[string]interface{}) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FpPipelineParam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipelineParam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FpPipelineParam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FpPipelineParam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FpPipelineParam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipelineParam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FpPipelineParam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FpPipelineParam) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *FpPipelineParam) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipelineParam) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *FpPipelineParam) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *FpPipelineParam) SetOptional(v bool) {
	o.Optional = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FpPipelineParam) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpPipelineParam) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FpPipelineParam) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FpPipelineParam) SetType(v string) {
	o.Type = &v
}

func (o FpPipelineParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FpPipelineParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFpPipelineParam struct {
	value *FpPipelineParam
	isSet bool
}

func (v NullableFpPipelineParam) Get() *FpPipelineParam {
	return v.value
}

func (v *NullableFpPipelineParam) Set(val *FpPipelineParam) {
	v.value = val
	v.isSet = true
}

func (v NullableFpPipelineParam) IsSet() bool {
	return v.isSet
}

func (v *NullableFpPipelineParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFpPipelineParam(val *FpPipelineParam) *NullableFpPipelineParam {
	return &NullableFpPipelineParam{value: val, isSet: true}
}

func (v NullableFpPipelineParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFpPipelineParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


