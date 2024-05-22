# ModconfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]ModconfigStepError**](ModconfigStepError.md) |  | [optional] 
**FailureMode** | Pointer to **string** |  | [optional] 
**Flowpipe** | Pointer to **map[string]interface{}** | Flowpipe metadata, contains started_at, finished_at | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewModconfigOutput

`func NewModconfigOutput() *ModconfigOutput`

NewModconfigOutput instantiates a new ModconfigOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigOutputWithDefaults

`func NewModconfigOutputWithDefaults() *ModconfigOutput`

NewModconfigOutputWithDefaults instantiates a new ModconfigOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModconfigOutput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModconfigOutput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModconfigOutput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ModconfigOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *ModconfigOutput) GetErrors() []ModconfigStepError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModconfigOutput) GetErrorsOk() (*[]ModconfigStepError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModconfigOutput) SetErrors(v []ModconfigStepError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ModconfigOutput) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFailureMode

`func (o *ModconfigOutput) GetFailureMode() string`

GetFailureMode returns the FailureMode field if non-nil, zero value otherwise.

### GetFailureModeOk

`func (o *ModconfigOutput) GetFailureModeOk() (*string, bool)`

GetFailureModeOk returns a tuple with the FailureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMode

`func (o *ModconfigOutput) SetFailureMode(v string)`

SetFailureMode sets FailureMode field to given value.

### HasFailureMode

`func (o *ModconfigOutput) HasFailureMode() bool`

HasFailureMode returns a boolean if a field has been set.

### GetFlowpipe

`func (o *ModconfigOutput) GetFlowpipe() map[string]interface{}`

GetFlowpipe returns the Flowpipe field if non-nil, zero value otherwise.

### GetFlowpipeOk

`func (o *ModconfigOutput) GetFlowpipeOk() (*map[string]interface{}, bool)`

GetFlowpipeOk returns a tuple with the Flowpipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowpipe

`func (o *ModconfigOutput) SetFlowpipe(v map[string]interface{})`

SetFlowpipe sets Flowpipe field to given value.

### HasFlowpipe

`func (o *ModconfigOutput) HasFlowpipe() bool`

HasFlowpipe returns a boolean if a field has been set.

### GetStatus

`func (o *ModconfigOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModconfigOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModconfigOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModconfigOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


