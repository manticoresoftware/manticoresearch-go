# GeoDistance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationAnchor** | Pointer to [**GeoDistanceLocationAnchor**](GeoDistanceLocationAnchor.md) |  | [optional] 
**LocationSource** | Pointer to **string** | Field name in the document that contains location data | [optional] 
**DistanceType** | Pointer to **string** | Algorithm used to calculate the distance | [optional] 
**Distance** | Pointer to **string** | The distance from the anchor point to filter results by | [optional] 

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

`func (o *GeoDistance) GetLocationSource() string`

GetLocationSource returns the LocationSource field if non-nil, zero value otherwise.

### GetLocationSourceOk

`func (o *GeoDistance) GetLocationSourceOk() (*string, bool)`

GetLocationSourceOk returns a tuple with the LocationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSource

`func (o *GeoDistance) SetLocationSource(v string)`

SetLocationSource sets LocationSource field to given value.

### HasLocationSource

`func (o *GeoDistance) HasLocationSource() bool`

HasLocationSource returns a boolean if a field has been set.

### GetDistanceType

`func (o *GeoDistance) GetDistanceType() string`

GetDistanceType returns the DistanceType field if non-nil, zero value otherwise.

### GetDistanceTypeOk

`func (o *GeoDistance) GetDistanceTypeOk() (*string, bool)`

GetDistanceTypeOk returns a tuple with the DistanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceType

`func (o *GeoDistance) SetDistanceType(v string)`

SetDistanceType sets DistanceType field to given value.

### HasDistanceType

`func (o *GeoDistance) HasDistanceType() bool`

HasDistanceType returns a boolean if a field has been set.

### GetDistance

`func (o *GeoDistance) GetDistance() string`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *GeoDistance) GetDistanceOk() (*string, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *GeoDistance) SetDistance(v string)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *GeoDistance) HasDistance() bool`

HasDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


