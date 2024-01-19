# FpTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EndLineNumber** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pipelines** | Pointer to [**[]FpTriggerPipeline**](FpTriggerPipeline.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**StartLineNumber** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewFpTrigger

`func NewFpTrigger() *FpTrigger`

NewFpTrigger instantiates a new FpTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFpTriggerWithDefaults

`func NewFpTriggerWithDefaults() *FpTrigger`

NewFpTriggerWithDefaults instantiates a new FpTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FpTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FpTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FpTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FpTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *FpTrigger) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *FpTrigger) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *FpTrigger) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *FpTrigger) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetEnabled

`func (o *FpTrigger) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FpTrigger) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FpTrigger) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FpTrigger) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEndLineNumber

`func (o *FpTrigger) GetEndLineNumber() int32`

GetEndLineNumber returns the EndLineNumber field if non-nil, zero value otherwise.

### GetEndLineNumberOk

`func (o *FpTrigger) GetEndLineNumberOk() (*int32, bool)`

GetEndLineNumberOk returns a tuple with the EndLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLineNumber

`func (o *FpTrigger) SetEndLineNumber(v int32)`

SetEndLineNumber sets EndLineNumber field to given value.

### HasEndLineNumber

`func (o *FpTrigger) HasEndLineNumber() bool`

HasEndLineNumber returns a boolean if a field has been set.

### GetFileName

`func (o *FpTrigger) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FpTrigger) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FpTrigger) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FpTrigger) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetName

`func (o *FpTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FpTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FpTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FpTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPipelines

`func (o *FpTrigger) GetPipelines() []FpTriggerPipeline`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *FpTrigger) GetPipelinesOk() (*[]FpTriggerPipeline, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *FpTrigger) SetPipelines(v []FpTriggerPipeline)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *FpTrigger) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetQuery

`func (o *FpTrigger) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FpTrigger) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FpTrigger) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *FpTrigger) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSchedule

`func (o *FpTrigger) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *FpTrigger) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *FpTrigger) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *FpTrigger) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStartLineNumber

`func (o *FpTrigger) GetStartLineNumber() int32`

GetStartLineNumber returns the StartLineNumber field if non-nil, zero value otherwise.

### GetStartLineNumberOk

`func (o *FpTrigger) GetStartLineNumberOk() (*int32, bool)`

GetStartLineNumberOk returns a tuple with the StartLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLineNumber

`func (o *FpTrigger) SetStartLineNumber(v int32)`

SetStartLineNumber sets StartLineNumber field to given value.

### HasStartLineNumber

`func (o *FpTrigger) HasStartLineNumber() bool`

HasStartLineNumber returns a boolean if a field has been set.

### GetTags

`func (o *FpTrigger) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FpTrigger) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FpTrigger) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FpTrigger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *FpTrigger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FpTrigger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FpTrigger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FpTrigger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *FpTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FpTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FpTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FpTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *FpTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FpTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FpTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FpTrigger) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


