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

// checks if the ExecutionStepStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionStepStatus{}

// ExecutionStepStatus struct for ExecutionStepStatus
type ExecutionStepStatus struct {
	// Step executions that are failed.
	Failed *map[string]bool `json:"failed,omitempty"`
	// Step executions that are finished.
	Finished *map[string]bool `json:"finished,omitempty"`
	// When the step is initializing it doesn't yet have any executions. We track it as initializing until the first execution is queued.
	Initializing *bool `json:"initializing,omitempty"`
	// Indicate that step is in a loop so we don't mark it as finished
	LoopHold *bool `json:"loop_hold,omitempty"`
	// Step executions that are queued.
	Queued *map[string]bool `json:"queued,omitempty"`
	// Step executions that are started.
	Started *map[string]bool `json:"started,omitempty"`
	// There's the step execution in execution, this is the same but in a list for a given step status The element in this slice should point to the same element in the StepExecutions map (in PipelineExecution)
	StepExecutions []ExecutionStepExecution `json:"step_executions,omitempty"`
}
func (o ExecutionStepStatus) GetResourceType() string {
	return "ExecutionStepStatus"
}
// NewExecutionStepStatus instantiates a new ExecutionStepStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionStepStatus() *ExecutionStepStatus {
	this := ExecutionStepStatus{}
	return &this
}

// NewExecutionStepStatusWithDefaults instantiates a new ExecutionStepStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionStepStatusWithDefaults() *ExecutionStepStatus {
	this := ExecutionStepStatus{}
	return &this
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetFailed() map[string]bool {
	if o == nil || IsNil(o.Failed) {
		var ret map[string]bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetFailedOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given map[string]bool and assigns it to the Failed field.
func (o *ExecutionStepStatus) SetFailed(v map[string]bool) {
	o.Failed = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetFinished() map[string]bool {
	if o == nil || IsNil(o.Finished) {
		var ret map[string]bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetFinishedOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given map[string]bool and assigns it to the Finished field.
func (o *ExecutionStepStatus) SetFinished(v map[string]bool) {
	o.Finished = &v
}

// GetInitializing returns the Initializing field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetInitializing() bool {
	if o == nil || IsNil(o.Initializing) {
		var ret bool
		return ret
	}
	return *o.Initializing
}

// GetInitializingOk returns a tuple with the Initializing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetInitializingOk() (*bool, bool) {
	if o == nil || IsNil(o.Initializing) {
		return nil, false
	}
	return o.Initializing, true
}

// HasInitializing returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasInitializing() bool {
	if o != nil && !IsNil(o.Initializing) {
		return true
	}

	return false
}

// SetInitializing gets a reference to the given bool and assigns it to the Initializing field.
func (o *ExecutionStepStatus) SetInitializing(v bool) {
	o.Initializing = &v
}

// GetLoopHold returns the LoopHold field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetLoopHold() bool {
	if o == nil || IsNil(o.LoopHold) {
		var ret bool
		return ret
	}
	return *o.LoopHold
}

// GetLoopHoldOk returns a tuple with the LoopHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetLoopHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.LoopHold) {
		return nil, false
	}
	return o.LoopHold, true
}

// HasLoopHold returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasLoopHold() bool {
	if o != nil && !IsNil(o.LoopHold) {
		return true
	}

	return false
}

// SetLoopHold gets a reference to the given bool and assigns it to the LoopHold field.
func (o *ExecutionStepStatus) SetLoopHold(v bool) {
	o.LoopHold = &v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetQueued() map[string]bool {
	if o == nil || IsNil(o.Queued) {
		var ret map[string]bool
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetQueuedOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given map[string]bool and assigns it to the Queued field.
func (o *ExecutionStepStatus) SetQueued(v map[string]bool) {
	o.Queued = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetStarted() map[string]bool {
	if o == nil || IsNil(o.Started) {
		var ret map[string]bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetStartedOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given map[string]bool and assigns it to the Started field.
func (o *ExecutionStepStatus) SetStarted(v map[string]bool) {
	o.Started = &v
}

// GetStepExecutions returns the StepExecutions field value if set, zero value otherwise.
func (o *ExecutionStepStatus) GetStepExecutions() []ExecutionStepExecution {
	if o == nil || IsNil(o.StepExecutions) {
		var ret []ExecutionStepExecution
		return ret
	}
	return o.StepExecutions
}

// GetStepExecutionsOk returns a tuple with the StepExecutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStepStatus) GetStepExecutionsOk() ([]ExecutionStepExecution, bool) {
	if o == nil || IsNil(o.StepExecutions) {
		return nil, false
	}
	return o.StepExecutions, true
}

// HasStepExecutions returns a boolean if a field has been set.
func (o *ExecutionStepStatus) HasStepExecutions() bool {
	if o != nil && !IsNil(o.StepExecutions) {
		return true
	}

	return false
}

// SetStepExecutions gets a reference to the given []ExecutionStepExecution and assigns it to the StepExecutions field.
func (o *ExecutionStepStatus) SetStepExecutions(v []ExecutionStepExecution) {
	o.StepExecutions = v
}

func (o ExecutionStepStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionStepStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.Initializing) {
		toSerialize["initializing"] = o.Initializing
	}
	if !IsNil(o.LoopHold) {
		toSerialize["loop_hold"] = o.LoopHold
	}
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if !IsNil(o.StepExecutions) {
		toSerialize["step_executions"] = o.StepExecutions
	}
	return toSerialize, nil
}

type NullableExecutionStepStatus struct {
	value *ExecutionStepStatus
	isSet bool
}

func (v NullableExecutionStepStatus) Get() *ExecutionStepStatus {
	return v.value
}

func (v *NullableExecutionStepStatus) Set(val *ExecutionStepStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionStepStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionStepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionStepStatus(val *ExecutionStepStatus) *NullableExecutionStepStatus {
	return &NullableExecutionStepStatus{value: val, isSet: true}
}

func (v NullableExecutionStepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionStepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


