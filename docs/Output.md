# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]StepError**](StepError.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOutput

`func NewOutput() *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Output) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Output) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Output) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Output) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *Output) GetErrors() []StepError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Output) GetErrorsOk() (*[]StepError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Output) SetErrors(v []StepError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Output) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *Output) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Output) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Output) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Output) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


