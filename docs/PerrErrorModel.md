# PerrErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Instance** | **string** |  | 
**Retryable** | Pointer to **bool** | All errors are fatal unless specified | [optional] 
**Status** | **int32** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**ValidationErrors** | Pointer to [**[]PerrErrorDetailModel**](PerrErrorDetailModel.md) |  | [optional] 

## Methods

### NewPerrErrorModel

`func NewPerrErrorModel(instance string, status int32, title string, type_ string, ) *PerrErrorModel`

NewPerrErrorModel instantiates a new PerrErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerrErrorModelWithDefaults

`func NewPerrErrorModelWithDefaults() *PerrErrorModel`

NewPerrErrorModelWithDefaults instantiates a new PerrErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *PerrErrorModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PerrErrorModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PerrErrorModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PerrErrorModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *PerrErrorModel) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *PerrErrorModel) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *PerrErrorModel) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetRetryable

`func (o *PerrErrorModel) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *PerrErrorModel) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *PerrErrorModel) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *PerrErrorModel) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetStatus

`func (o *PerrErrorModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PerrErrorModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PerrErrorModel) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *PerrErrorModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PerrErrorModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PerrErrorModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *PerrErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PerrErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PerrErrorModel) SetType(v string)`

SetType sets Type field to given value.


### GetValidationErrors

`func (o *PerrErrorModel) GetValidationErrors() []PerrErrorDetailModel`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *PerrErrorModel) GetValidationErrorsOk() (*[]PerrErrorDetailModel, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *PerrErrorModel) SetValidationErrors(v []PerrErrorDetailModel)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *PerrErrorModel) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


