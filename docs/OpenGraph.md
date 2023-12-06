# OpenGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The opengraph description (og:description) of the mod, for use in social media applications | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenGraph

`func NewOpenGraph() *OpenGraph`

NewOpenGraph instantiates a new OpenGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenGraphWithDefaults

`func NewOpenGraphWithDefaults() *OpenGraph`

NewOpenGraphWithDefaults instantiates a new OpenGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OpenGraph) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenGraph) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenGraph) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenGraph) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *OpenGraph) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OpenGraph) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OpenGraph) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OpenGraph) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


