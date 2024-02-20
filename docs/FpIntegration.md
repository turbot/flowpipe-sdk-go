# FpIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**EndLineNumber** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartLineNumber** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewFpIntegration

`func NewFpIntegration() *FpIntegration`

NewFpIntegration instantiates a new FpIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpIntegrationWithDefaults

`func NewFpIntegrationWithDefaults() *FpIntegration`

NewFpIntegrationWithDefaults instantiates a new FpIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FpIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *FpIntegration) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *FpIntegration) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *FpIntegration) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *FpIntegration) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *FpIntegration) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *FpIntegration) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *FpIntegration) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *FpIntegration) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FpIntegration) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FpIntegration) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FpIntegration) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FpIntegration) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetName

`func (o *FpIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FpIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FpIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FpIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartLineNumber

`func (o *FpIntegration) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *FpIntegration) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *FpIntegration) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *FpIntegration) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetTags

`func (o *FpIntegration) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FpIntegration) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FpIntegration) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FpIntegration) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *FpIntegration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FpIntegration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FpIntegration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FpIntegration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *FpIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FpIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FpIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FpIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *FpIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FpIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FpIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FpIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


