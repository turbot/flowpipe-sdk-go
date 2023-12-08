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

// checks if the OpenGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenGraph{}

// OpenGraph struct for OpenGraph
type OpenGraph struct {
	// The opengraph description (og:description) of the mod, for use in social media applications
	Description *string `json:"description,omitempty"`
	Title *string `json:"title,omitempty"`
}
func (o OpenGraph) GetResourceType() string {
	return "OpenGraph"
}
// NewOpenGraph instantiates a new OpenGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraph() *OpenGraph {
	this := OpenGraph{}
	return &this
}

// NewOpenGraphWithDefaults instantiates a new OpenGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphWithDefaults() *OpenGraph {
	this := OpenGraph{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenGraph) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraph) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenGraph) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenGraph) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OpenGraph) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraph) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OpenGraph) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OpenGraph) SetTitle(v string) {
	o.Title = &v
}

func (o OpenGraph) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableOpenGraph struct {
	value *OpenGraph
	isSet bool
}

func (v NullableOpenGraph) Get() *OpenGraph {
	return v.value
}

func (v *NullableOpenGraph) Set(val *OpenGraph) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraph) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraph(val *OpenGraph) *NullableOpenGraph {
	return &NullableOpenGraph{value: val, isSet: true}
}

func (v NullableOpenGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

