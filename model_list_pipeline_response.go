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

// checks if the ListPipelineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPipelineResponse{}

// ListPipelineResponse struct for ListPipelineResponse
type ListPipelineResponse struct {
	Items []Pipeline `json:"items,omitempty"`
	NextToken *string `json:"next_token,omitempty"`
}

// NewListPipelineResponse instantiates a new ListPipelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPipelineResponse() *ListPipelineResponse {
	this := ListPipelineResponse{}
	return &this
}

// NewListPipelineResponseWithDefaults instantiates a new ListPipelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPipelineResponseWithDefaults() *ListPipelineResponse {
	this := ListPipelineResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListPipelineResponse) GetItems() []Pipeline {
	if o == nil || IsNil(o.Items) {
		var ret []Pipeline
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPipelineResponse) GetItemsOk() ([]Pipeline, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListPipelineResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Pipeline and assigns it to the Items field.
func (o *ListPipelineResponse) SetItems(v []Pipeline) {
	o.Items = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListPipelineResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPipelineResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListPipelineResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListPipelineResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListPipelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPipelineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextToken) {
		toSerialize["next_token"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListPipelineResponse struct {
	value *ListPipelineResponse
	isSet bool
}

func (v NullableListPipelineResponse) Get() *ListPipelineResponse {
	return v.value
}

func (v *NullableListPipelineResponse) Set(val *ListPipelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPipelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPipelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPipelineResponse(val *ListPipelineResponse) *NullableListPipelineResponse {
	return &NullableListPipelineResponse{value: val, isSet: true}
}

func (v NullableListPipelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPipelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


