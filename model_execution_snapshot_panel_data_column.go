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

// checks if the ExecutionSnapshotPanelDataColumn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionSnapshotPanelDataColumn{}

// ExecutionSnapshotPanelDataColumn struct for ExecutionSnapshotPanelDataColumn
type ExecutionSnapshotPanelDataColumn struct {
	DataType *string `json:"data_type,omitempty"`
	Name *string `json:"name,omitempty"`
}
func (o ExecutionSnapshotPanelDataColumn) GetResourceType() string {
	return "ExecutionSnapshotPanelDataColumn"
}
// NewExecutionSnapshotPanelDataColumn instantiates a new ExecutionSnapshotPanelDataColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionSnapshotPanelDataColumn() *ExecutionSnapshotPanelDataColumn {
	this := ExecutionSnapshotPanelDataColumn{}
	return &this
}

// NewExecutionSnapshotPanelDataColumnWithDefaults instantiates a new ExecutionSnapshotPanelDataColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionSnapshotPanelDataColumnWithDefaults() *ExecutionSnapshotPanelDataColumn {
	this := ExecutionSnapshotPanelDataColumn{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *ExecutionSnapshotPanelDataColumn) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionSnapshotPanelDataColumn) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *ExecutionSnapshotPanelDataColumn) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *ExecutionSnapshotPanelDataColumn) SetDataType(v string) {
	o.DataType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExecutionSnapshotPanelDataColumn) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionSnapshotPanelDataColumn) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExecutionSnapshotPanelDataColumn) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExecutionSnapshotPanelDataColumn) SetName(v string) {
	o.Name = &v
}

func (o ExecutionSnapshotPanelDataColumn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionSnapshotPanelDataColumn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableExecutionSnapshotPanelDataColumn struct {
	value *ExecutionSnapshotPanelDataColumn
	isSet bool
}

func (v NullableExecutionSnapshotPanelDataColumn) Get() *ExecutionSnapshotPanelDataColumn {
	return v.value
}

func (v *NullableExecutionSnapshotPanelDataColumn) Set(val *ExecutionSnapshotPanelDataColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionSnapshotPanelDataColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionSnapshotPanelDataColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionSnapshotPanelDataColumn(val *ExecutionSnapshotPanelDataColumn) *NullableExecutionSnapshotPanelDataColumn {
	return &NullableExecutionSnapshotPanelDataColumn{value: val, isSet: true}
}

func (v NullableExecutionSnapshotPanelDataColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionSnapshotPanelDataColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


