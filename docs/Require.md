# Require

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flowpipe** | Pointer to [**FlowpipeRequire**](FlowpipeRequire.md) |  | [optional] 
**Mods** | Pointer to [**[]ModVersionConstraint**](ModVersionConstraint.md) |  | [optional] 

## Methods

### NewRequire

`func NewRequire() *Require`

NewRequire instantiates a new Require object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequireWithDefaults

`func NewRequireWithDefaults() *Require`

NewRequireWithDefaults instantiates a new Require object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowpipe

`func (o *Require) GetFlowpipe() FlowpipeRequire`

GetFlowpipe returns the Flowpipe field if non-nil, zero value otherwise.

### GetFlowpipeOk

`func (o *Require) GetFlowpipeOk() (*FlowpipeRequire, bool)`

GetFlowpipeOk returns a tuple with the Flowpipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowpipe

`func (o *Require) SetFlowpipe(v FlowpipeRequire)`

SetFlowpipe sets Flowpipe field to given value.

### HasFlowpipe

`func (o *Require) HasFlowpipe() bool`

HasFlowpipe returns a boolean if a field has been set.

### GetMods

`func (o *Require) GetMods() []ModVersionConstraint`

GetMods returns the Mods field if non-nil, zero value otherwise.

### GetModsOk

`func (o *Require) GetModsOk() (*[]ModVersionConstraint, bool)`

GetModsOk returns a tuple with the Mods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMods

`func (o *Require) SetMods(v []ModVersionConstraint)`

SetMods sets Mods field to given value.

### HasMods

`func (o *Require) HasMods() bool`

HasMods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


