# HclRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Filename is the name of the file into which this range&#39;s positions point. | [optional] 
**Start** | Pointer to [**HclRangeStart**](HclRangeStart.md) |  | [optional] 

## Methods

### NewHclRange

`func NewHclRange() *HclRange`

NewHclRange instantiates a new HclRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclRangeWithDefaults

`func NewHclRangeWithDefaults() *HclRange`

NewHclRangeWithDefaults instantiates a new HclRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *HclRange) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *HclRange) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *HclRange) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *HclRange) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetStart

`func (o *HclRange) GetStart() HclRangeStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HclRange) GetStartOk() (*HclRangeStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HclRange) SetStart(v HclRangeStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *HclRange) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


