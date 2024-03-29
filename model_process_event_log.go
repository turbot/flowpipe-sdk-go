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

// checks if the ProcessEventLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessEventLog{}

// ProcessEventLog struct for ProcessEventLog
type ProcessEventLog struct {
	EventType *string `json:"event_type,omitempty"`
	// Setting the type as string for now, as the CLI need to print the payload
	Payload *string `json:"payload,omitempty"`
	Ts *string `json:"ts,omitempty"`
}
func (o ProcessEventLog) GetResourceType() string {
	return "ProcessEventLog"
}
// NewProcessEventLog instantiates a new ProcessEventLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessEventLog() *ProcessEventLog {
	this := ProcessEventLog{}
	return &this
}

// NewProcessEventLogWithDefaults instantiates a new ProcessEventLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessEventLogWithDefaults() *ProcessEventLog {
	this := ProcessEventLog{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ProcessEventLog) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessEventLog) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ProcessEventLog) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *ProcessEventLog) SetEventType(v string) {
	o.EventType = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProcessEventLog) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessEventLog) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProcessEventLog) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *ProcessEventLog) SetPayload(v string) {
	o.Payload = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *ProcessEventLog) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessEventLog) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *ProcessEventLog) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *ProcessEventLog) SetTs(v string) {
	o.Ts = &v
}

func (o ProcessEventLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessEventLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableProcessEventLog struct {
	value *ProcessEventLog
	isSet bool
}

func (v NullableProcessEventLog) Get() *ProcessEventLog {
	return v.value
}

func (v *NullableProcessEventLog) Set(val *ProcessEventLog) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessEventLog) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessEventLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessEventLog(val *ProcessEventLog) *NullableProcessEventLog {
	return &NullableProcessEventLog{value: val, isSet: true}
}

func (v NullableProcessEventLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessEventLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


