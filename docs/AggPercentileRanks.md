# AggPercentileRanks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Numeric field to calculate percentile ranks for | 
**Values** | **[]float32** | Input values to rank | 
**Keyed** | Pointer to **bool** | Return an object keyed by input value when true, or an array when false. Default is false.  | [optional] 
**Tdigest** | Pointer to [**AggTDigest**](AggTDigest.md) |  | [optional] 

## Methods

### NewAggPercentileRanks

`func NewAggPercentileRanks(field string, values []float32, ) *AggPercentileRanks`

NewAggPercentileRanks instantiates a new AggPercentileRanks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggPercentileRanksWithDefaults

`func NewAggPercentileRanksWithDefaults() *AggPercentileRanks`

NewAggPercentileRanksWithDefaults instantiates a new AggPercentileRanks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggPercentileRanks) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggPercentileRanks) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggPercentileRanks) SetField(v string)`

SetField sets Field field to given value.


### GetValues

`func (o *AggPercentileRanks) GetValues() []float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AggPercentileRanks) GetValuesOk() (*[]float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AggPercentileRanks) SetValues(v []float32)`

SetValues sets Values field to given value.


### GetKeyed

`func (o *AggPercentileRanks) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *AggPercentileRanks) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *AggPercentileRanks) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *AggPercentileRanks) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetTdigest

`func (o *AggPercentileRanks) GetTdigest() AggTDigest`

GetTdigest returns the Tdigest field if non-nil, zero value otherwise.

### GetTdigestOk

`func (o *AggPercentileRanks) GetTdigestOk() (*AggTDigest, bool)`

GetTdigestOk returns a tuple with the Tdigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdigest

`func (o *AggPercentileRanks) SetTdigest(v AggTDigest)`

SetTdigest sets Tdigest field to given value.

### HasTdigest

`func (o *AggPercentileRanks) HasTdigest() bool`

HasTdigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


