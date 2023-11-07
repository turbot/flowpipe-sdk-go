# ExecutionStepExecutionStepForEach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Each** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Output** | Pointer to [**ModconfigOutput**](ModconfigOutput.md) |  | [optional] 
**TotalCount** | **int32** |  | 

## Methods

### NewExecutionStepExecutionStepForEach

`func NewExecutionStepExecutionStepForEach(key string, totalCount int32, ) *ExecutionStepExecutionStepForEach`

NewExecutionStepExecutionStepForEach instantiates a new ExecutionStepExecutionStepForEach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepExecutionStepForEachWithDefaults

`func NewExecutionStepExecutionStepForEachWithDefaults() *ExecutionStepExecutionStepForEach`

NewExecutionStepExecutionStepForEachWithDefaults instantiates a new ExecutionStepExecutionStepForEach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEach

`func (o *ExecutionStepExecutionStepForEach) GetEach() string`

GetEach returns the Each field if non-nil, zero value otherwise.

### GetEachOk

`func (o *ExecutionStepExecutionStepForEach) GetEachOk() (*string, bool)`

GetEachOk returns a tuple with the Each field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEach

`func (o *ExecutionStepExecutionStepForEach) SetEach(v string)`

SetEach sets Each field to given value.

### HasEach

`func (o *ExecutionStepExecutionStepForEach) HasEach() bool`

HasEach returns a boolean if a field has been set.

### GetKey

`func (o *ExecutionStepExecutionStepForEach) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExecutionStepExecutionStepForEach) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExecutionStepExecutionStepForEach) SetKey(v string)`

SetKey sets Key field to given value.


### GetOutput

`func (o *ExecutionStepExecutionStepForEach) GetOutput() ModconfigOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ExecutionStepExecutionStepForEach) GetOutputOk() (*ModconfigOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ExecutionStepExecutionStepForEach) SetOutput(v ModconfigOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ExecutionStepExecutionStepForEach) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTotalCount

`func (o *ExecutionStepExecutionStepForEach) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExecutionStepExecutionStepForEach) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExecutionStepExecutionStepForEach) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


