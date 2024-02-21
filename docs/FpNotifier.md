# FpNotifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**EndLineNumber** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifies** | Pointer to [**[]FpNotify**](FpNotify.md) |  | [optional] 
**StartLineNumber** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewFpNotifier

`func NewFpNotifier() *FpNotifier`

NewFpNotifier instantiates a new FpNotifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpNotifierWithDefaults

`func NewFpNotifierWithDefaults() *FpNotifier`

NewFpNotifierWithDefaults instantiates a new FpNotifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FpNotifier) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpNotifier) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpNotifier) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpNotifier) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *FpNotifier) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *FpNotifier) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *FpNotifier) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *FpNotifier) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FpNotifier) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FpNotifier) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FpNotifier) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FpNotifier) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetName

`func (o *FpNotifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FpNotifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FpNotifier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FpNotifier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifies

`func (o *FpNotifier) GetNotifies() []FpNotify`

GetNotifies returns the Notifies field if non-nil, zero value otherwise.

### GetNotifiesOk

`func (o *FpNotifier) GetNotifiesOk() (*[]FpNotify, bool)`

GetNotifiesOk returns a tuple with the Notifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifies

`func (o *FpNotifier) SetNotifies(v []FpNotify)`

SetNotifies sets Notifies field to given value.

### HasNotifies

`func (o *FpNotifier) HasNotifies() bool`

HasNotifies returns a boolean if a field has been set.

### GetStartLineNumber

`func (o *FpNotifier) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *FpNotifier) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *FpNotifier) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *FpNotifier) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetTitle

`func (o *FpNotifier) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FpNotifier) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FpNotifier) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FpNotifier) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


