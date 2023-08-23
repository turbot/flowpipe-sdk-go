# PcerrErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Instance** | **string** |  | 
**Retryable** | Pointer to **bool** | All errors are fatal unless specified | [optional] 
**Status** | **int32** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**ValidationErrors** | Pointer to [**[]PcerrErrorDetailModel**](PcerrErrorDetailModel.md) |  | [optional] 

## Methods

### NewPcerrErrorModel

`func NewPcerrErrorModel(instance string, status int32, title string, type_ string, ) *PcerrErrorModel`

NewPcerrErrorModel instantiates a new PcerrErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcerrErrorModelWithDefaults

`func NewPcerrErrorModelWithDefaults() *PcerrErrorModel`

NewPcerrErrorModelWithDefaults instantiates a new PcerrErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *PcerrErrorModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PcerrErrorModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PcerrErrorModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PcerrErrorModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *PcerrErrorModel) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *PcerrErrorModel) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *PcerrErrorModel) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetRetryable

`func (o *PcerrErrorModel) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *PcerrErrorModel) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *PcerrErrorModel) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *PcerrErrorModel) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetStatus

`func (o *PcerrErrorModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PcerrErrorModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PcerrErrorModel) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *PcerrErrorModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PcerrErrorModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PcerrErrorModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *PcerrErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PcerrErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PcerrErrorModel) SetType(v string)`

SetType sets Type field to given value.


### GetValidationErrors

`func (o *PcerrErrorModel) GetValidationErrors() []PcerrErrorDetailModel`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PcerrErrorModel) GetValidationErrorsOk() (*[]PcerrErrorDetailModel, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PcerrErrorModel) SetValidationErrors(v []PcerrErrorDetailModel)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *PcerrErrorModel) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


