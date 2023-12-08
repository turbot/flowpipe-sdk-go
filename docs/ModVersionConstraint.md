# ModVersionConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | the fully qualified mod name, e.g. github.com/turbot/mod1 | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModVersionConstraint

`func NewModVersionConstraint() *ModVersionConstraint`

NewModVersionConstraint instantiates a new ModVersionConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModVersionConstraintWithDefaults

`func NewModVersionConstraintWithDefaults() *ModVersionConstraint`

NewModVersionConstraintWithDefaults instantiates a new ModVersionConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModVersionConstraint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModVersionConstraint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModVersionConstraint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModVersionConstraint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ModVersionConstraint) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModVersionConstraint) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModVersionConstraint) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModVersionConstraint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


