# ProcessEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** | Setting the type as string for now, as the CLI need to print the payload | [optional] 
**Ts** | Pointer to **string** |  | [optional] 

## Methods

### NewProcessEventLog

`func NewProcessEventLog() *ProcessEventLog`

NewProcessEventLog instantiates a new ProcessEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessEventLogWithDefaults

`func NewProcessEventLogWithDefaults() *ProcessEventLog`

NewProcessEventLogWithDefaults instantiates a new ProcessEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ProcessEventLog) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ProcessEventLog) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ProcessEventLog) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ProcessEventLog) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPayload

`func (o *ProcessEventLog) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProcessEventLog) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProcessEventLog) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ProcessEventLog) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTs

`func (o *ProcessEventLog) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ProcessEventLog) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ProcessEventLog) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ProcessEventLog) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


