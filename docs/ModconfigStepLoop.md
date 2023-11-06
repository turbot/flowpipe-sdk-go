# ModconfigStepLoop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**LoopCompleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewModconfigStepLoop

`func NewModconfigStepLoop(index int32, ) *ModconfigStepLoop`

NewModconfigStepLoop instantiates a new ModconfigStepLoop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigStepLoopWithDefaults

`func NewModconfigStepLoopWithDefaults() *ModconfigStepLoop`

NewModconfigStepLoopWithDefaults instantiates a new ModconfigStepLoop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ModconfigStepLoop) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ModconfigStepLoop) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ModconfigStepLoop) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInput

`func (o *ModconfigStepLoop) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ModconfigStepLoop) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ModconfigStepLoop) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ModconfigStepLoop) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetLoopCompleted

`func (o *ModconfigStepLoop) GetLoopCompleted() bool`

GetLoopCompleted returns the LoopCompleted field if non-nil, zero value otherwise.

### GetLoopCompletedOk

`func (o *ModconfigStepLoop) GetLoopCompletedOk() (*bool, bool)`

GetLoopCompletedOk returns a tuple with the LoopCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopCompleted

`func (o *ModconfigStepLoop) SetLoopCompleted(v bool)`

SetLoopCompleted sets LoopCompleted field to given value.

### HasLoopCompleted

`func (o *ModconfigStepLoop) HasLoopCompleted() bool`

HasLoopCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


