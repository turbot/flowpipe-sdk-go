# FperrErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Instance** | **string** |  | 
**Status** | **int32** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**ValidationErrors** | Pointer to [**[]FperrErrorDetailModel**](FperrErrorDetailModel.md) |  | [optional] 

## Methods

### NewFperrErrorModel

`func NewFperrErrorModel(instance string, status int32, title string, type_ string, ) *FperrErrorModel`

NewFperrErrorModel instantiates a new FperrErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFperrErrorModelWithDefaults

`func NewFperrErrorModelWithDefaults() *FperrErrorModel`

NewFperrErrorModelWithDefaults instantiates a new FperrErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *FperrErrorModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *FperrErrorModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *FperrErrorModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *FperrErrorModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *FperrErrorModel) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *FperrErrorModel) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *FperrErrorModel) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetStatus

`func (o *FperrErrorModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FperrErrorModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FperrErrorModel) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *FperrErrorModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FperrErrorModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FperrErrorModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *FperrErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FperrErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FperrErrorModel) SetType(v string)`

SetType sets Type field to given value.


### GetValidationErrors

`func (o *FperrErrorModel) GetValidationErrors() []FperrErrorDetailModel`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *FperrErrorModel) GetValidationErrorsOk() (*[]FperrErrorDetailModel, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *FperrErrorModel) SetValidationErrors(v []FperrErrorDetailModel)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *FperrErrorModel) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


