# ExecutionSnapshotPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboard** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**ExecutionSnapshotPanelData**](ExecutionSnapshotPanelData.md) |  | [optional] 
**DisplayType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PanelType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### NewExecutionSnapshotPanel

`func NewExecutionSnapshotPanel() *ExecutionSnapshotPanel`

NewExecutionSnapshotPanel instantiates a new ExecutionSnapshotPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionSnapshotPanelWithDefaults

`func NewExecutionSnapshotPanelWithDefaults() *ExecutionSnapshotPanel`

NewExecutionSnapshotPanelWithDefaults instantiates a new ExecutionSnapshotPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboard

`func (o *ExecutionSnapshotPanel) GetDashboard() string`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *ExecutionSnapshotPanel) GetDashboardOk() (*string, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *ExecutionSnapshotPanel) SetDashboard(v string)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *ExecutionSnapshotPanel) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### GetData

`func (o *ExecutionSnapshotPanel) GetData() ExecutionSnapshotPanelData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecutionSnapshotPanel) GetDataOk() (*ExecutionSnapshotPanelData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecutionSnapshotPanel) SetData(v ExecutionSnapshotPanelData)`

SetData sets Data field to given value.

### HasData

`func (o *ExecutionSnapshotPanel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisplayType

`func (o *ExecutionSnapshotPanel) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *ExecutionSnapshotPanel) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *ExecutionSnapshotPanel) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *ExecutionSnapshotPanel) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetName

`func (o *ExecutionSnapshotPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecutionSnapshotPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecutionSnapshotPanel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecutionSnapshotPanel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPanelType

`func (o *ExecutionSnapshotPanel) GetPanelType() string`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *ExecutionSnapshotPanel) GetPanelTypeOk() (*string, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *ExecutionSnapshotPanel) SetPanelType(v string)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *ExecutionSnapshotPanel) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetProperties

`func (o *ExecutionSnapshotPanel) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExecutionSnapshotPanel) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExecutionSnapshotPanel) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ExecutionSnapshotPanel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionSnapshotPanel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionSnapshotPanel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionSnapshotPanel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionSnapshotPanel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ExecutionSnapshotPanel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExecutionSnapshotPanel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExecutionSnapshotPanel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExecutionSnapshotPanel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWidth

`func (o *ExecutionSnapshotPanel) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ExecutionSnapshotPanel) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ExecutionSnapshotPanel) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ExecutionSnapshotPanel) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


