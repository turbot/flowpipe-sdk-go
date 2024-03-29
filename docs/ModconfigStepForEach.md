# ModconfigStepForEach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForEachStep** | Pointer to **bool** |  | [optional] 
**Key** | **string** |  | 
**Output** | Pointer to [**ModconfigOutput**](ModconfigOutput.md) |  | [optional] 
**TotalCount** | **int32** |  | 

## Methods

### NewModconfigStepForEach

`func NewModconfigStepForEach(key string, totalCount int32, ) *ModconfigStepForEach`

NewModconfigStepForEach instantiates a new ModconfigStepForEach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModconfigStepForEachWithDefaults

`func NewModconfigStepForEachWithDefaults() *ModconfigStepForEach`

NewModconfigStepForEachWithDefaults instantiates a new ModconfigStepForEach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForEachStep

`func (o *ModconfigStepForEach) GetForEachStep() bool`

GetForEachStep returns the ForEachStep field if non-nil, zero value otherwise.

### GetForEachStepOk

`func (o *ModconfigStepForEach) GetForEachStepOk() (*bool, bool)`

GetForEachStepOk returns a tuple with the ForEachStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForEachStep

`func (o *ModconfigStepForEach) SetForEachStep(v bool)`

SetForEachStep sets ForEachStep field to given value.

### HasForEachStep

`func (o *ModconfigStepForEach) HasForEachStep() bool`

HasForEachStep returns a boolean if a field has been set.

### GetKey

`func (o *ModconfigStepForEach) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModconfigStepForEach) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModconfigStepForEach) SetKey(v string)`

SetKey sets Key field to given value.


### GetOutput

`func (o *ModconfigStepForEach) GetOutput() ModconfigOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ModconfigStepForEach) GetOutputOk() (*ModconfigOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ModconfigStepForEach) SetOutput(v ModconfigOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ModconfigStepForEach) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetTotalCount

`func (o *ModconfigStepForEach) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ModconfigStepForEach) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ModconfigStepForEach) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


