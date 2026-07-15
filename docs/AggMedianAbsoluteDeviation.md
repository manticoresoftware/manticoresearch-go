# AggMedianAbsoluteDeviation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Numeric field to calculate MAD for | 
**Tdigest** | Pointer to [**AggTDigest**](AggTDigest.md) |  | [optional] 

## Methods

### NewAggMedianAbsoluteDeviation

`func NewAggMedianAbsoluteDeviation(field string, ) *AggMedianAbsoluteDeviation`

NewAggMedianAbsoluteDeviation instantiates a new AggMedianAbsoluteDeviation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggMedianAbsoluteDeviationWithDefaults

`func NewAggMedianAbsoluteDeviationWithDefaults() *AggMedianAbsoluteDeviation`

NewAggMedianAbsoluteDeviationWithDefaults instantiates a new AggMedianAbsoluteDeviation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggMedianAbsoluteDeviation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggMedianAbsoluteDeviation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggMedianAbsoluteDeviation) SetField(v string)`

SetField sets Field field to given value.


### GetTdigest

`func (o *AggMedianAbsoluteDeviation) GetTdigest() AggTDigest`

GetTdigest returns the Tdigest field if non-nil, zero value otherwise.

### GetTdigestOk

`func (o *AggMedianAbsoluteDeviation) GetTdigestOk() (*AggTDigest, bool)`

GetTdigestOk returns a tuple with the Tdigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdigest

`func (o *AggMedianAbsoluteDeviation) SetTdigest(v AggTDigest)`

SetTdigest sets Tdigest field to given value.

### HasTdigest

`func (o *AggMedianAbsoluteDeviation) HasTdigest() bool`

HasTdigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


