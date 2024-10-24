# AttrFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equals** | Pointer to [**EqualsFilterEquals**](EqualsFilterEquals.md) |  | [optional] 
**In** | Pointer to [**map[string][]EqualsFilterEquals**](array.md) |  | [optional] 
**Range** | Pointer to [**map[string]RangeFilterRangeValue**](RangeFilterRangeValue.md) |  | [optional] 
**GeoDistance** | Pointer to [**GeoFilterGeoDistance**](GeoFilterGeoDistance.md) |  | [optional] 

## Methods

### NewAttrFilter

`func NewAttrFilter() *AttrFilter`

NewAttrFilter instantiates a new AttrFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttrFilterWithDefaults

`func NewAttrFilterWithDefaults() *AttrFilter`

NewAttrFilterWithDefaults instantiates a new AttrFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquals

`func (o *AttrFilter) GetEquals() EqualsFilterEquals`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *AttrFilter) GetEqualsOk() (*EqualsFilterEquals, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *AttrFilter) SetEquals(v EqualsFilterEquals)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *AttrFilter) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetIn

`func (o *AttrFilter) GetIn() map[string][]EqualsFilterEquals`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *AttrFilter) GetInOk() (*map[string][]EqualsFilterEquals, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *AttrFilter) SetIn(v map[string][]EqualsFilterEquals)`

SetIn sets In field to given value.

### HasIn

`func (o *AttrFilter) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetRange

`func (o *AttrFilter) GetRange() map[string]RangeFilterRangeValue`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AttrFilter) GetRangeOk() (*map[string]RangeFilterRangeValue, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AttrFilter) SetRange(v map[string]RangeFilterRangeValue)`

SetRange sets Range field to given value.

### HasRange

`func (o *AttrFilter) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetGeoDistance

`func (o *AttrFilter) GetGeoDistance() GeoFilterGeoDistance`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *AttrFilter) GetGeoDistanceOk() (*GeoFilterGeoDistance, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *AttrFilter) SetGeoDistance(v GeoFilterGeoDistance)`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *AttrFilter) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


