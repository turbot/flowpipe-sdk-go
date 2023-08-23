# ListTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]FpTrigger**](FpTrigger.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListTriggerResponse

`func NewListTriggerResponse() *ListTriggerResponse`

NewListTriggerResponse instantiates a new ListTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTriggerResponseWithDefaults

`func NewListTriggerResponseWithDefaults() *ListTriggerResponse`

NewListTriggerResponseWithDefaults instantiates a new ListTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListTriggerResponse) GetItems() []FpTrigger`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListTriggerResponse) GetItemsOk() (*[]FpTrigger, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListTriggerResponse) SetItems(v []FpTrigger)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListTriggerResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListTriggerResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListTriggerResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListTriggerResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListTriggerResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


