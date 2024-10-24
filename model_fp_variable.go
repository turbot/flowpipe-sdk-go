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

// checks if the FpVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FpVariable{}

// FpVariable struct for FpVariable
type FpVariable struct {
	Description *string `json:"description,omitempty"`
	EndLineNumber *int32 `json:"end_line_number,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	ModName *string `json:"mod_name,omitempty"`
	QualifiedName *string `json:"qualified_name,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	StartLineNumber *int32 `json:"start_line_number,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Type *interface{} `json:"type,omitempty"`
	TypeString *string `json:"type_string,omitempty"`
	Value *interface{} `json:"value,omitempty"`
	ValueDefault *interface{} `json:"value_default,omitempty"`
}
func (o FpVariable) GetResourceType() string {
	return "FpVariable"
}
// NewFpVariable instantiates a new FpVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFpVariable() *FpVariable {
	this := FpVariable{}
	return &this
}

// NewFpVariableWithDefaults instantiates a new FpVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFpVariableWithDefaults() *FpVariable {
	this := FpVariable{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FpVariable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FpVariable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FpVariable) SetDescription(v string) {
	o.Description = &v
}

// GetEndLineNumber returns the EndLineNumber field value if set, zero value otherwise.
func (o *FpVariable) GetEndLineNumber() int32 {
	if o == nil || IsNil(o.EndLineNumber) {
		var ret int32
		return ret
	}
	return *o.EndLineNumber
}

// GetEndLineNumberOk returns a tuple with the EndLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetEndLineNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.EndLineNumber) {
		return nil, false
	}
	return o.EndLineNumber, true
}

// HasEndLineNumber returns a boolean if a field has been set.
func (o *FpVariable) HasEndLineNumber() bool {
	if o != nil && !IsNil(o.EndLineNumber) {
		return true
	}

	return false
}

// SetEndLineNumber gets a reference to the given int32 and assigns it to the EndLineNumber field.
func (o *FpVariable) SetEndLineNumber(v int32) {
	o.EndLineNumber = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *FpVariable) GetEnum() []interface{} {
	if o == nil || IsNil(o.Enum) {
		var ret []interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetEnumOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *FpVariable) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []interface{} and assigns it to the Enum field.
func (o *FpVariable) SetEnum(v []interface{}) {
	o.Enum = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FpVariable) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FpVariable) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FpVariable) SetFileName(v string) {
	o.FileName = &v
}

// GetModName returns the ModName field value if set, zero value otherwise.
func (o *FpVariable) GetModName() string {
	if o == nil || IsNil(o.ModName) {
		var ret string
		return ret
	}
	return *o.ModName
}

// GetModNameOk returns a tuple with the ModName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetModNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModName) {
		return nil, false
	}
	return o.ModName, true
}

// HasModName returns a boolean if a field has been set.
func (o *FpVariable) HasModName() bool {
	if o != nil && !IsNil(o.ModName) {
		return true
	}

	return false
}

// SetModName gets a reference to the given string and assigns it to the ModName field.
func (o *FpVariable) SetModName(v string) {
	o.ModName = &v
}

// GetQualifiedName returns the QualifiedName field value if set, zero value otherwise.
func (o *FpVariable) GetQualifiedName() string {
	if o == nil || IsNil(o.QualifiedName) {
		var ret string
		return ret
	}
	return *o.QualifiedName
}

// GetQualifiedNameOk returns a tuple with the QualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetQualifiedNameOk() (*string, bool) {
	if o == nil || IsNil(o.QualifiedName) {
		return nil, false
	}
	return o.QualifiedName, true
}

// HasQualifiedName returns a boolean if a field has been set.
func (o *FpVariable) HasQualifiedName() bool {
	if o != nil && !IsNil(o.QualifiedName) {
		return true
	}

	return false
}

// SetQualifiedName gets a reference to the given string and assigns it to the QualifiedName field.
func (o *FpVariable) SetQualifiedName(v string) {
	o.QualifiedName = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *FpVariable) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *FpVariable) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *FpVariable) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetStartLineNumber returns the StartLineNumber field value if set, zero value otherwise.
func (o *FpVariable) GetStartLineNumber() int32 {
	if o == nil || IsNil(o.StartLineNumber) {
		var ret int32
		return ret
	}
	return *o.StartLineNumber
}

// GetStartLineNumberOk returns a tuple with the StartLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetStartLineNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.StartLineNumber) {
		return nil, false
	}
	return o.StartLineNumber, true
}

// HasStartLineNumber returns a boolean if a field has been set.
func (o *FpVariable) HasStartLineNumber() bool {
	if o != nil && !IsNil(o.StartLineNumber) {
		return true
	}

	return false
}

// SetStartLineNumber gets a reference to the given int32 and assigns it to the StartLineNumber field.
func (o *FpVariable) SetStartLineNumber(v int32) {
	o.StartLineNumber = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FpVariable) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FpVariable) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *FpVariable) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FpVariable) GetType() interface{} {
	if o == nil || IsNil(o.Type) {
		var ret interface{}
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FpVariable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *FpVariable) SetType(v interface{}) {
	o.Type = &v
}

// GetTypeString returns the TypeString field value if set, zero value otherwise.
func (o *FpVariable) GetTypeString() string {
	if o == nil || IsNil(o.TypeString) {
		var ret string
		return ret
	}
	return *o.TypeString
}

// GetTypeStringOk returns a tuple with the TypeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetTypeStringOk() (*string, bool) {
	if o == nil || IsNil(o.TypeString) {
		return nil, false
	}
	return o.TypeString, true
}

// HasTypeString returns a boolean if a field has been set.
func (o *FpVariable) HasTypeString() bool {
	if o != nil && !IsNil(o.TypeString) {
		return true
	}

	return false
}

// SetTypeString gets a reference to the given string and assigns it to the TypeString field.
func (o *FpVariable) SetTypeString(v string) {
	o.TypeString = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FpVariable) GetValue() interface{} {
	if o == nil || IsNil(o.Value) {
		var ret interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FpVariable) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *FpVariable) SetValue(v interface{}) {
	o.Value = &v
}

// GetValueDefault returns the ValueDefault field value if set, zero value otherwise.
func (o *FpVariable) GetValueDefault() interface{} {
	if o == nil || IsNil(o.ValueDefault) {
		var ret interface{}
		return ret
	}
	return *o.ValueDefault
}

// GetValueDefaultOk returns a tuple with the ValueDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FpVariable) GetValueDefaultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ValueDefault) {
		return nil, false
	}
	return o.ValueDefault, true
}

// HasValueDefault returns a boolean if a field has been set.
func (o *FpVariable) HasValueDefault() bool {
	if o != nil && !IsNil(o.ValueDefault) {
		return true
	}

	return false
}

// SetValueDefault gets a reference to the given interface{} and assigns it to the ValueDefault field.
func (o *FpVariable) SetValueDefault(v interface{}) {
	o.ValueDefault = &v
}

func (o FpVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FpVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EndLineNumber) {
		toSerialize["end_line_number"] = o.EndLineNumber
	}
	if !IsNil(o.Enum) {
		toSerialize["enum"] = o.Enum
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.ModName) {
		toSerialize["mod_name"] = o.ModName
	}
	if !IsNil(o.QualifiedName) {
		toSerialize["qualified_name"] = o.QualifiedName
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !IsNil(o.StartLineNumber) {
		toSerialize["start_line_number"] = o.StartLineNumber
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeString) {
		toSerialize["type_string"] = o.TypeString
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueDefault) {
		toSerialize["value_default"] = o.ValueDefault
	}
	return toSerialize, nil
}

type NullableFpVariable struct {
	value *FpVariable
	isSet bool
}

func (v NullableFpVariable) Get() *FpVariable {
	return v.value
}

func (v *NullableFpVariable) Set(val *FpVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableFpVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableFpVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFpVariable(val *FpVariable) *NullableFpVariable {
	return &NullableFpVariable{value: val, isSet: true}
}

func (v NullableFpVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFpVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


