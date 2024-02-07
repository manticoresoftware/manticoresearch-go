# GeoDistanceFilter

Geo distance attribute filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationAnchor** | Pointer to [**GeoDistanceFilterLocationAnchor**](GeoDistanceFilterLocationAnchor.md) |  | [optional] 
**LocationSource** | Pointer to **string** | Attribute containing latitude and longitude data | [optional] 
**DistanceType** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **string** |  | [optional] 

## Methods

### NewGeoDistanceFilter

`func NewGeoDistanceFilter() *GeoDistanceFilter`

NewGeoDistanceFilter instantiates a new GeoDistanceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoDistanceFilterWithDefaults

`func NewGeoDistanceFilterWithDefaults() *GeoDistanceFilter`

NewGeoDistanceFilterWithDefaults instantiates a new GeoDistanceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationAnchor

`func (o *GeoDistanceFilter) GetLocationAnchor() GeoDistanceFilterLocationAnchor`

GetLocationAnchor returns the LocationAnchor field if non-nil, zero value otherwise.

### GetLocationAnchorOk

`func (o *GeoDistanceFilter) GetLocationAnchorOk() (*GeoDistanceFilterLocationAnchor, bool)`

GetLocationAnchorOk returns a tuple with the LocationAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAnchor

`func (o *GeoDistanceFilter) SetLocationAnchor(v GeoDistanceFilterLocationAnchor)`

SetLocationAnchor sets LocationAnchor field to given value.

### HasLocationAnchor

`func (o *GeoDistanceFilter) HasLocationAnchor() bool`

HasLocationAnchor returns a boolean if a field has been set.

### GetLocationSource

`func (o *GeoDistanceFilter) GetLocationSource() string`

GetLocationSource returns the LocationSource field if non-nil, zero value otherwise.

### GetLocationSourceOk

`func (o *GeoDistanceFilter) GetLocationSourceOk() (*string, bool)`

GetLocationSourceOk returns a tuple with the LocationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSource

`func (o *GeoDistanceFilter) SetLocationSource(v string)`

SetLocationSource sets LocationSource field to given value.

### HasLocationSource

`func (o *GeoDistanceFilter) HasLocationSource() bool`

HasLocationSource returns a boolean if a field has been set.

### GetDistanceType

`func (o *GeoDistanceFilter) GetDistanceType() string`

GetDistanceType returns the DistanceType field if non-nil, zero value otherwise.

### GetDistanceTypeOk

`func (o *GeoDistanceFilter) GetDistanceTypeOk() (*string, bool)`

GetDistanceTypeOk returns a tuple with the DistanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceType

`func (o *GeoDistanceFilter) SetDistanceType(v string)`

SetDistanceType sets DistanceType field to given value.

### HasDistanceType

`func (o *GeoDistanceFilter) HasDistanceType() bool`

HasDistanceType returns a boolean if a field has been set.

### GetDistance

`func (o *GeoDistanceFilter) GetDistance() string`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *GeoDistanceFilter) GetDistanceOk() (*string, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *GeoDistanceFilter) SetDistance(v string)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *GeoDistanceFilter) HasDistance() bool`

HasDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


