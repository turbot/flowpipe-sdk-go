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

// checks if the PerrErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerrErrorModel{}

// PerrErrorModel struct for PerrErrorModel
type PerrErrorModel struct {
	Detail *string `json:"detail,omitempty"`
	Instance string `json:"instance"`
	// All errors are fatal unless specified
	Retryable *bool `json:"retryable,omitempty"`
	Status int32 `json:"status"`
	Title string `json:"title"`
	Type string `json:"type"`
	ValidationErrors []PerrErrorDetailModel `json:"validation_errors,omitempty"`
}
func (o PerrErrorModel) GetResourceType() string {
	return "PerrErrorModel"
}
// NewPerrErrorModel instantiates a new PerrErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerrErrorModel(instance string, status int32, title string, type_ string) *PerrErrorModel {
	this := PerrErrorModel{}
	this.Instance = instance
	this.Status = status
	this.Title = title
	this.Type = type_
	return &this
}

// NewPerrErrorModelWithDefaults instantiates a new PerrErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerrErrorModelWithDefaults() *PerrErrorModel {
	this := PerrErrorModel{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *PerrErrorModel) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *PerrErrorModel) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *PerrErrorModel) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value
func (o *PerrErrorModel) GetInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *PerrErrorModel) SetInstance(v string) {
	o.Instance = v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *PerrErrorModel) GetRetryable() bool {
	if o == nil || IsNil(o.Retryable) {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetRetryableOk() (*bool, bool) {
	if o == nil || IsNil(o.Retryable) {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *PerrErrorModel) HasRetryable() bool {
	if o != nil && !IsNil(o.Retryable) {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *PerrErrorModel) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetStatus returns the Status field value
func (o *PerrErrorModel) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PerrErrorModel) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *PerrErrorModel) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PerrErrorModel) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *PerrErrorModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PerrErrorModel) SetType(v string) {
	o.Type = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *PerrErrorModel) GetValidationErrors() []PerrErrorDetailModel {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []PerrErrorDetailModel
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerrErrorModel) GetValidationErrorsOk() ([]PerrErrorDetailModel, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *PerrErrorModel) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []PerrErrorDetailModel and assigns it to the ValidationErrors field.
func (o *PerrErrorModel) SetValidationErrors(v []PerrErrorDetailModel) {
	o.ValidationErrors = v
}

func (o PerrErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerrErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["instance"] = o.Instance
	if !IsNil(o.Retryable) {
		toSerialize["retryable"] = o.Retryable
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	if !IsNil(o.ValidationErrors) {
		toSerialize["validation_errors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullablePerrErrorModel struct {
	value *PerrErrorModel
	isSet bool
}

func (v NullablePerrErrorModel) Get() *PerrErrorModel {
	return v.value
}

func (v *NullablePerrErrorModel) Set(val *PerrErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePerrErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePerrErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerrErrorModel(val *PerrErrorModel) *NullablePerrErrorModel {
	return &NullablePerrErrorModel{value: val, isSet: true}
}

func (v NullablePerrErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerrErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


