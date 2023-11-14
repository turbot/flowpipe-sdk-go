# ModconfigStepRetry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**RetryCompleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewModconfigStepRetry

`func NewModconfigStepRetry(index int32, ) *ModconfigStepRetry`

NewModconfigStepRetry instantiates a new ModconfigStepRetry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigStepRetryWithDefaults

`func NewModconfigStepRetryWithDefaults() *ModconfigStepRetry`

NewModconfigStepRetryWithDefaults instantiates a new ModconfigStepRetry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ModconfigStepRetry) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ModconfigStepRetry) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ModconfigStepRetry) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInput

`func (o *ModconfigStepRetry) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ModconfigStepRetry) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ModconfigStepRetry) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ModconfigStepRetry) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetRetryCompleted

`func (o *ModconfigStepRetry) GetRetryCompleted() bool`

GetRetryCompleted returns the RetryCompleted field if non-nil, zero value otherwise.

### GetRetryCompletedOk

`func (o *ModconfigStepRetry) GetRetryCompletedOk() (*bool, bool)`

GetRetryCompletedOk returns a tuple with the RetryCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCompleted

`func (o *ModconfigStepRetry) SetRetryCompleted(v bool)`

SetRetryCompleted sets RetryCompleted field to given value.

### HasRetryCompleted

`func (o *ModconfigStepRetry) HasRetryCompleted() bool`

HasRetryCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


