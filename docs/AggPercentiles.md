# AggPercentiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Numeric field to calculate percentiles for | 
**Values** | Pointer to **[]float32** | Percentile points to compute (0-100). Defaults to 1, 5, 25, 50, 75, 95, and 99 when omitted.  | [optional] 
**Keyed** | Pointer to **bool** | Return an object keyed by percentile when true, or an array when false. Default is false.  | [optional] 
**Tdigest** | Pointer to [**AggTDigest**](AggTDigest.md) |  | [optional] 

## Methods

### NewAggPercentiles

`func NewAggPercentiles(field string, ) *AggPercentiles`

NewAggPercentiles instantiates a new AggPercentiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggPercentilesWithDefaults

`func NewAggPercentilesWithDefaults() *AggPercentiles`

NewAggPercentilesWithDefaults instantiates a new AggPercentiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggPercentiles) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggPercentiles) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggPercentiles) SetField(v string)`

SetField sets Field field to given value.


### GetValues

`func (o *AggPercentiles) GetValues() []float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AggPercentiles) GetValuesOk() (*[]float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AggPercentiles) SetValues(v []float32)`

SetValues sets Values field to given value.

### HasValues

`func (o *AggPercentiles) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetKeyed

`func (o *AggPercentiles) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *AggPercentiles) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *AggPercentiles) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *AggPercentiles) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetTdigest

`func (o *AggPercentiles) GetTdigest() AggTDigest`

GetTdigest returns the Tdigest field if non-nil, zero value otherwise.

### GetTdigestOk

`func (o *AggPercentiles) GetTdigestOk() (*AggTDigest, bool)`

GetTdigestOk returns a tuple with the Tdigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdigest

`func (o *AggPercentiles) SetTdigest(v AggTDigest)`

SetTdigest sets Tdigest field to given value.

### HasTdigest

`func (o *AggPercentiles) HasTdigest() bool`

HasTdigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


