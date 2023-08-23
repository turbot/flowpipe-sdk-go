# ModconfigOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]ModconfigStepError**](ModconfigStepError.md) |  | [optional] 
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


