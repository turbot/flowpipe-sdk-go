/*
Flowpipe

Flowpipe is workflow and pipelines for DevSecOps.

API version: 0.1.0
Contact: info@flowpipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FperrErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FperrErrorModel{}

// FperrErrorModel struct for FperrErrorModel
type FperrErrorModel struct {
	Detail *string `json:"detail,omitempty"`
	Instance string `json:"instance"`
	Status int32 `json:"status"`
	Title string `json:"title"`
	Type string `json:"type"`
	ValidationErrors []FperrErrorDetailModel `json:"validation_errors,omitempty"`
}

// NewFperrErrorModel instantiates a new FperrErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFperrErrorModel(instance string, status int32, title string, type_ string) *FperrErrorModel {
	this := FperrErrorModel{}
	this.Instance = instance
	this.Status = status
	this.Title = title
	this.Type = type_
	return &this
}

// NewFperrErrorModelWithDefaults instantiates a new FperrErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFperrErrorModelWithDefaults() *FperrErrorModel {
	this := FperrErrorModel{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *FperrErrorModel) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *FperrErrorModel) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *FperrErrorModel) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value
func (o *FperrErrorModel) GetInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *FperrErrorModel) SetInstance(v string) {
	o.Instance = v
}

// GetStatus returns the Status field value
func (o *FperrErrorModel) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FperrErrorModel) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *FperrErrorModel) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FperrErrorModel) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *FperrErrorModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FperrErrorModel) SetType(v string) {
	o.Type = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *FperrErrorModel) GetValidationErrors() []FperrErrorDetailModel {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []FperrErrorDetailModel
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FperrErrorModel) GetValidationErrorsOk() ([]FperrErrorDetailModel, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *FperrErrorModel) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []FperrErrorDetailModel and assigns it to the ValidationErrors field.
func (o *FperrErrorModel) SetValidationErrors(v []FperrErrorDetailModel) {
	o.ValidationErrors = v
}

func (o FperrErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FperrErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["instance"] = o.Instance
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	if !IsNil(o.ValidationErrors) {
		toSerialize["validation_errors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableFperrErrorModel struct {
	value *FperrErrorModel
	isSet bool
}

func (v NullableFperrErrorModel) Get() *FperrErrorModel {
	return v.value
}

func (v *NullableFperrErrorModel) Set(val *FperrErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFperrErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFperrErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFperrErrorModel(val *FperrErrorModel) *NullableFperrErrorModel {
	return &NullableFperrErrorModel{value: val, isSet: true}
}

func (v NullableFperrErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFperrErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


