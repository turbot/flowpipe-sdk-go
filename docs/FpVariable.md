# FpVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EndLineNumber** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ModName** | Pointer to **string** |  | [optional] 
**QualifiedName** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**StartLineNumber** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TypeString** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ValueDefault** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFpVariable

`func NewFpVariable() *FpVariable`

NewFpVariable instantiates a new FpVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpVariableWithDefaults

`func NewFpVariableWithDefaults() *FpVariable`

NewFpVariableWithDefaults instantiates a new FpVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FpVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *FpVariable) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *FpVariable) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *FpVariable) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *FpVariable) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FpVariable) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FpVariable) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FpVariable) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FpVariable) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetModName

`func (o *FpVariable) GetModName() string`

GetModName returns the ModName field if non-nil, zero value otherwise.

### GetModNameOk

`func (o *FpVariable) GetModNameOk() (*string, bool)`

GetModNameOk returns a tuple with the ModName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModName

`func (o *FpVariable) SetModName(v string)`

SetModName sets ModName field to given value.

### HasModName

`func (o *FpVariable) HasModName() bool`

HasModName returns a boolean if a field has been set.

### GetQualifiedName

`func (o *FpVariable) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *FpVariable) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *FpVariable) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *FpVariable) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### GetResourceName

`func (o *FpVariable) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *FpVariable) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *FpVariable) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *FpVariable) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetStartLineNumber

`func (o *FpVariable) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *FpVariable) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *FpVariable) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *FpVariable) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetType

`func (o *FpVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FpVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FpVariable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FpVariable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeString

`func (o *FpVariable) GetTypeString() string`

GetTypeString returns the TypeString field if non-nil, zero value otherwise.

### GetTypeStringOk

`func (o *FpVariable) GetTypeStringOk() (*string, bool)`

GetTypeStringOk returns a tuple with the TypeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeString

`func (o *FpVariable) SetTypeString(v string)`

SetTypeString sets TypeString field to given value.

### HasTypeString

`func (o *FpVariable) HasTypeString() bool`

HasTypeString returns a boolean if a field has been set.

### GetValue

`func (o *FpVariable) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FpVariable) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FpVariable) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *FpVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueDefault

`func (o *FpVariable) GetValueDefault() interface{}`

GetValueDefault returns the ValueDefault field if non-nil, zero value otherwise.

### GetValueDefaultOk

`func (o *FpVariable) GetValueDefaultOk() (*interface{}, bool)`

GetValueDefaultOk returns a tuple with the ValueDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDefault

`func (o *FpVariable) SetValueDefault(v interface{})`

SetValueDefault sets ValueDefault field to given value.

### HasValueDefault

`func (o *FpVariable) HasValueDefault() bool`

HasValueDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


