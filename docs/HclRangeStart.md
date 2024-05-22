# HclRangeStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Byte** | Pointer to **int32** | Byte is the byte offset into the file where the indicated character begins. This is a zero-based offset to the first byte of the first UTF-8 codepoint sequence in the character, and thus gives a position that can be resolved _without_ awareness of Unicode characters. | [optional] 
**Column** | Pointer to **int32** | Column is the source code column where this position points, in unicode characters, with counting starting at 1.  Column counts characters as they appear visually, so for example a latin letter with a combining diacritic mark counts as one character. This is intended for rendering visual markers against source code in contexts where these diacritics would be rendered in a single character cell. Technically speaking, Column is counting grapheme clusters as used in unicode normalization. | [optional] 
**Line** | Pointer to **int32** | Line is the source code line where this position points. Lines are counted starting at 1 and incremented for each newline character encountered. | [optional] 

## Methods

### NewHclRangeStart

`func NewHclRangeStart() *HclRangeStart`

NewHclRangeStart instantiates a new HclRangeStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclRangeStartWithDefaults

`func NewHclRangeStartWithDefaults() *HclRangeStart`

NewHclRangeStartWithDefaults instantiates a new HclRangeStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByte

`func (o *HclRangeStart) GetByte() int32`

GetByte returns the Byte field if non-nil, zero value otherwise.

### GetByteOk

`func (o *HclRangeStart) GetByteOk() (*int32, bool)`

GetByteOk returns a tuple with the Byte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByte

`func (o *HclRangeStart) SetByte(v int32)`

SetByte sets Byte field to given value.

### HasByte

`func (o *HclRangeStart) HasByte() bool`

HasByte returns a boolean if a field has been set.

### GetColumn

`func (o *HclRangeStart) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *HclRangeStart) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *HclRangeStart) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *HclRangeStart) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetLine

`func (o *HclRangeStart) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *HclRangeStart) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *HclRangeStart) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *HclRangeStart) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


