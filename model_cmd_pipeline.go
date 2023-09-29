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

// checks if the CmdPipeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmdPipeline{}

// CmdPipeline struct for CmdPipeline
type CmdPipeline struct {
	Args map[string]interface{} `json:"args,omitempty"`
	ArgsString *map[string]string `json:"args_string,omitempty"`
	Command string `json:"command"`
}
func (o CmdPipeline) GetResourceType() string {
	return "CmdPipeline"
}
// NewCmdPipeline instantiates a new CmdPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmdPipeline(command string) *CmdPipeline {
	this := CmdPipeline{}
	this.Command = command
	return &this
}

// NewCmdPipelineWithDefaults instantiates a new CmdPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmdPipelineWithDefaults() *CmdPipeline {
	this := CmdPipeline{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CmdPipeline) GetArgs() map[string]interface{} {
	if o == nil || IsNil(o.Args) {
		var ret map[string]interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmdPipeline) GetArgsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return map[string]interface{}{}, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CmdPipeline) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given map[string]interface{} and assigns it to the Args field.
func (o *CmdPipeline) SetArgs(v map[string]interface{}) {
	o.Args = v
}

// GetArgsString returns the ArgsString field value if set, zero value otherwise.
func (o *CmdPipeline) GetArgsString() map[string]string {
	if o == nil || IsNil(o.ArgsString) {
		var ret map[string]string
		return ret
	}
	return *o.ArgsString
}

// GetArgsStringOk returns a tuple with the ArgsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmdPipeline) GetArgsStringOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ArgsString) {
		return nil, false
	}
	return o.ArgsString, true
}

// HasArgsString returns a boolean if a field has been set.
func (o *CmdPipeline) HasArgsString() bool {
	if o != nil && !IsNil(o.ArgsString) {
		return true
	}

	return false
}

// SetArgsString gets a reference to the given map[string]string and assigns it to the ArgsString field.
func (o *CmdPipeline) SetArgsString(v map[string]string) {
	o.ArgsString = &v
}

// GetCommand returns the Command field value
func (o *CmdPipeline) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *CmdPipeline) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *CmdPipeline) SetCommand(v string) {
	o.Command = v
}

func (o CmdPipeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmdPipeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.ArgsString) {
		toSerialize["args_string"] = o.ArgsString
	}
	toSerialize["command"] = o.Command
	return toSerialize, nil
}

type NullableCmdPipeline struct {
	value *CmdPipeline
	isSet bool
}

func (v NullableCmdPipeline) Get() *CmdPipeline {
	return v.value
}

func (v *NullableCmdPipeline) Set(val *CmdPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableCmdPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableCmdPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmdPipeline(val *CmdPipeline) *NullableCmdPipeline {
	return &NullableCmdPipeline{value: val, isSet: true}
}

func (v NullableCmdPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmdPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


