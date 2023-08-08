# ExecutionStepExecutionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]StepError**](StepError.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionStepExecutionOutput

`func NewExecutionStepExecutionOutput() *ExecutionStepExecutionOutput`

NewExecutionStepExecutionOutput instantiates a new ExecutionStepExecutionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepExecutionOutputWithDefaults

`func NewExecutionStepExecutionOutputWithDefaults() *ExecutionStepExecutionOutput`

NewExecutionStepExecutionOutputWithDefaults instantiates a new ExecutionStepExecutionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExecutionStepExecutionOutput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecutionStepExecutionOutput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecutionStepExecutionOutput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ExecutionStepExecutionOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *ExecutionStepExecutionOutput) GetErrors() []StepError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExecutionStepExecutionOutput) GetErrorsOk() (*[]StepError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExecutionStepExecutionOutput) SetErrors(v []StepError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExecutionStepExecutionOutput) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionStepExecutionOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionStepExecutionOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionStepExecutionOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionStepExecutionOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


