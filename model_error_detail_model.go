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

// checks if the ErrorDetailModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetailModel{}

// ErrorDetailModel struct for ErrorDetailModel
type ErrorDetailModel struct {
	Location *string `json:"location,omitempty"`
	Message *string `json:"message,omitempty"`
}
func (o ErrorDetailModel) GetResourceType() string {
	return "ErrorDetailModel"
}
// NewErrorDetailModel instantiates a new ErrorDetailModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetailModel() *ErrorDetailModel {
	this := ErrorDetailModel{}
	return &this
}

// NewErrorDetailModelWithDefaults instantiates a new ErrorDetailModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailModelWithDefaults() *ErrorDetailModel {
	this := ErrorDetailModel{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ErrorDetailModel) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetailModel) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ErrorDetailModel) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ErrorDetailModel) SetLocation(v string) {
	o.Location = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorDetailModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetailModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorDetailModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorDetailModel) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorDetailModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetailModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorDetailModel struct {
	value *ErrorDetailModel
	isSet bool
}

func (v NullableErrorDetailModel) Get() *ErrorDetailModel {
	return v.value
}

func (v *NullableErrorDetailModel) Set(val *ErrorDetailModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetailModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetailModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetailModel(val *ErrorDetailModel) *NullableErrorDetailModel {
	return &NullableErrorDetailModel{value: val, isSet: true}
}

func (v NullableErrorDetailModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetailModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


