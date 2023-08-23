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
	"fmt"
)

// ModconfigNextStepAction the model 'ModconfigNextStepAction'
type ModconfigNextStepAction string

// List of modconfig.NextStepAction
const (
	NextStepActionStart ModconfigNextStepAction = "start"
	NextStepActionInaccessible ModconfigNextStepAction = "inaccessible"
	NextStepActionSkip ModconfigNextStepAction = "skip"
)

// All allowed values of ModconfigNextStepAction enum
var AllowedModconfigNextStepActionEnumValues = []ModconfigNextStepAction{
	"start",
	"inaccessible",
	"skip",
}

func (v *ModconfigNextStepAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModconfigNextStepAction(value)
	for _, existing := range AllowedModconfigNextStepActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModconfigNextStepAction", value)
}

// NewModconfigNextStepActionFromValue returns a pointer to a valid ModconfigNextStepAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModconfigNextStepActionFromValue(v string) (*ModconfigNextStepAction, error) {
	ev := ModconfigNextStepAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModconfigNextStepAction: valid values are %v", v, AllowedModconfigNextStepActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModconfigNextStepAction) IsValid() bool {
	for _, existing := range AllowedModconfigNextStepActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modconfig.NextStepAction value
func (v ModconfigNextStepAction) Ptr() *ModconfigNextStepAction {
	return &v
}

type NullableModconfigNextStepAction struct {
	value *ModconfigNextStepAction
	isSet bool
}

func (v NullableModconfigNextStepAction) Get() *ModconfigNextStepAction {
	return v.value
}

func (v *NullableModconfigNextStepAction) Set(val *ModconfigNextStepAction) {
	v.value = val
	v.isSet = true
}

func (v NullableModconfigNextStepAction) IsSet() bool {
	return v.isSet
}

func (v *NullableModconfigNextStepAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModconfigNextStepAction(val *ModconfigNextStepAction) *NullableModconfigNextStepAction {
	return &NullableModconfigNextStepAction{value: val, isSet: true}
}

func (v NullableModconfigNextStepAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModconfigNextStepAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

