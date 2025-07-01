# GeoDistance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationAnchor** | Pointer to [**GeoDistanceLocationAnchor**](GeoDistanceLocationAnchor.md) |  | [optional] 
**LocationSource** | Pointer to **interface{}** | Field name in the document that contains location data | [optional] 
**DistanceType** | Pointer to **interface{}** | Algorithm used to calculate the distance | [optional] 
**Distance** | Pointer to **interface{}** | The distance from the anchor point to filter results by | [optional] 

## Methods

### NewGeoDistance

`func NewGeoDistance() *GeoDistance`

NewGeoDistance instantiates a new GeoDistance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoDistanceWithDefaults

`func NewGeoDistanceWithDefaults() *GeoDistance`

NewGeoDistanceWithDefaults instantiates a new GeoDistance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationAnchor

`func (o *GeoDistance) GetLocationAnchor() GeoDistanceLocationAnchor`

GetLocationAnchor returns the LocationAnchor field if non-nil, zero value otherwise.

### GetLocationAnchorOk

`func (o *GeoDistance) GetLocationAnchorOk() (*GeoDistanceLocationAnchor, bool)`

GetLocationAnchorOk returns a tuple with the LocationAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnchor

`func (o *GeoDistance) SetLocationAnchor(v GeoDistanceLocationAnchor)`

SetLocationAnchor sets LocationAnchor field to given value.

### HasLocationAnchor

`func (o *GeoDistance) HasLocationAnchor() bool`

HasLocationAnchor returns a boolean if a field has been set.

### GetLocationSource

`func (o *GeoDistance) GetLocationSource() interface{}`

GetLocationSource returns the LocationSource field if non-nil, zero value otherwise.

### GetLocationSourceOk

`func (o *GeoDistance) GetLocationSourceOk() (*interface{}, bool)`

GetLocationSourceOk returns a tuple with the LocationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSource

`func (o *GeoDistance) SetLocationSource(v interface{})`

SetLocationSource sets LocationSource field to given value.

### HasLocationSource

`func (o *GeoDistance) HasLocationSource() bool`

HasLocationSource returns a boolean if a field has been set.

### SetLocationSourceNil

`func (o *GeoDistance) SetLocationSourceNil(b bool)`

 SetLocationSourceNil sets the value for LocationSource to be an explicit nil

### UnsetLocationSource
`func (o *GeoDistance) UnsetLocationSource()`

UnsetLocationSource ensures that no value is present for LocationSource, not even an explicit nil
### GetDistanceType

`func (o *GeoDistance) GetDistanceType() interface{}`

GetDistanceType returns the DistanceType field if non-nil, zero value otherwise.

### GetDistanceTypeOk

`func (o *GeoDistance) GetDistanceTypeOk() (*interface{}, bool)`

GetDistanceTypeOk returns a tuple with the DistanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceType

`func (o *GeoDistance) SetDistanceType(v interface{})`

SetDistanceType sets DistanceType field to given value.

### HasDistanceType

`func (o *GeoDistance) HasDistanceType() bool`

HasDistanceType returns a boolean if a field has been set.

### SetDistanceTypeNil

`func (o *GeoDistance) SetDistanceTypeNil(b bool)`

 SetDistanceTypeNil sets the value for DistanceType to be an explicit nil

### UnsetDistanceType
`func (o *GeoDistance) UnsetDistanceType()`

UnsetDistanceType ensures that no value is present for DistanceType, not even an explicit nil
### GetDistance

`func (o *GeoDistance) GetDistance() interface{}`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *GeoDistance) GetDistanceOk() (*interface{}, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *GeoDistance) SetDistance(v interface{})`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *GeoDistance) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *GeoDistance) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *GeoDistance) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


