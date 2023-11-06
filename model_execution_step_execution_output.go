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

// checks if the ExecutionStepExecutionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionStepExecutionOutput{}

// ExecutionStepExecutionOutput Native/primitive output of the step
type ExecutionStepExecutionOutput struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Errors []ModconfigStepError `json:"errors,omitempty"`
	Status *string `json:"status,omitempty"`
}
func (o ExecutionStepExecutionOutput) GetResourceType() string {
	return "ExecutionStepExecutionOutput"
}
// NewExecutionStepExecutionOutput instantiates a new ExecutionStepExecutionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionStepExecutionOutput() *ExecutionStepExecutionOutput {
	this := ExecutionStepExecutionOutput{}
	return &this
}

// NewExecutionStepExecutionOutputWithDefaults instantiates a new ExecutionStepExecutionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionStepExecutionOutputWithDefaults() *ExecutionStepExecutionOutput {
	this := ExecutionStepExecutionOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExecutionStepExecutionOutput) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepExecutionOutput) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExecutionStepExecutionOutput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ExecutionStepExecutionOutput) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ExecutionStepExecutionOutput) GetErrors() []ModconfigStepError {
	if o == nil || IsNil(o.Errors) {
		var ret []ModconfigStepError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepExecutionOutput) GetErrorsOk() ([]ModconfigStepError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ExecutionStepExecutionOutput) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ModconfigStepError and assigns it to the Errors field.
func (o *ExecutionStepExecutionOutput) SetErrors(v []ModconfigStepError) {
	o.Errors = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExecutionStepExecutionOutput) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepExecutionOutput) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExecutionStepExecutionOutput) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExecutionStepExecutionOutput) SetStatus(v string) {
	o.Status = &v
}

func (o ExecutionStepExecutionOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionStepExecutionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableExecutionStepExecutionOutput struct {
	value *ExecutionStepExecutionOutput
	isSet bool
}

func (v NullableExecutionStepExecutionOutput) Get() *ExecutionStepExecutionOutput {
	return v.value
}

func (v *NullableExecutionStepExecutionOutput) Set(val *ExecutionStepExecutionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionStepExecutionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionStepExecutionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionStepExecutionOutput(val *ExecutionStepExecutionOutput) *NullableExecutionStepExecutionOutput {
	return &NullableExecutionStepExecutionOutput{value: val, isSet: true}
}

func (v NullableExecutionStepExecutionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionStepExecutionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


