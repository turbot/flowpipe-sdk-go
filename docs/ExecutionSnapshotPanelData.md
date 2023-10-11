# ExecutionSnapshotPanelData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]ExecutionSnapshotPanelDataColumn**](ExecutionSnapshotPanelDataColumn.md) |  | [optional] 
**Rows** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewExecutionSnapshotPanelData

`func NewExecutionSnapshotPanelData() *ExecutionSnapshotPanelData`

NewExecutionSnapshotPanelData instantiates a new ExecutionSnapshotPanelData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionSnapshotPanelDataWithDefaults

`func NewExecutionSnapshotPanelDataWithDefaults() *ExecutionSnapshotPanelData`

NewExecutionSnapshotPanelDataWithDefaults instantiates a new ExecutionSnapshotPanelData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ExecutionSnapshotPanelData) GetColumns() []ExecutionSnapshotPanelDataColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ExecutionSnapshotPanelData) GetColumnsOk() (*[]ExecutionSnapshotPanelDataColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ExecutionSnapshotPanelData) SetColumns(v []ExecutionSnapshotPanelDataColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ExecutionSnapshotPanelData) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *ExecutionSnapshotPanelData) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ExecutionSnapshotPanelData) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ExecutionSnapshotPanelData) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *ExecutionSnapshotPanelData) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


