# ExecutionSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**ExecutionSnapshotLayout**](ExecutionSnapshotLayout.md) |  | [optional] 
**Panels** | Pointer to [**map[string]ExecutionSnapshotPanel**](ExecutionSnapshotPanel.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionSnapshot

`func NewExecutionSnapshot() *ExecutionSnapshot`

NewExecutionSnapshot instantiates a new ExecutionSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionSnapshotWithDefaults

`func NewExecutionSnapshotWithDefaults() *ExecutionSnapshot`

NewExecutionSnapshotWithDefaults instantiates a new ExecutionSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ExecutionSnapshot) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExecutionSnapshot) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExecutionSnapshot) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ExecutionSnapshot) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLayout

`func (o *ExecutionSnapshot) GetLayout() ExecutionSnapshotLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ExecutionSnapshot) GetLayoutOk() (*ExecutionSnapshotLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ExecutionSnapshot) SetLayout(v ExecutionSnapshotLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ExecutionSnapshot) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPanels

`func (o *ExecutionSnapshot) GetPanels() map[string]ExecutionSnapshotPanel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *ExecutionSnapshot) GetPanelsOk() (*map[string]ExecutionSnapshotPanel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *ExecutionSnapshot) SetPanels(v map[string]ExecutionSnapshotPanel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *ExecutionSnapshot) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ExecutionSnapshot) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ExecutionSnapshot) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ExecutionSnapshot) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ExecutionSnapshot) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStartTime

`func (o *ExecutionSnapshot) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExecutionSnapshot) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExecutionSnapshot) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ExecutionSnapshot) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


