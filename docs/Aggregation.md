# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | Pointer to [**AggTerms**](AggTerms.md) |  | [optional] 
**Sort** | Pointer to **[]interface{}** |  | [optional] 
**Composite** | Pointer to [**AggComposite**](AggComposite.md) |  | [optional] 
**Histogram** | Pointer to [**AggHistogram**](AggHistogram.md) |  | [optional] 
**DateHistogram** | Pointer to [**AggDateHistogram**](AggDateHistogram.md) |  | [optional] 
**Range** | Pointer to [**AggRange**](AggRange.md) |  | [optional] 
**DateRange** | Pointer to [**AggRange**](AggRange.md) |  | [optional] 
**Percentiles** | Pointer to [**AggPercentiles**](AggPercentiles.md) |  | [optional] 
**PercentileRanks** | Pointer to [**AggPercentileRanks**](AggPercentileRanks.md) |  | [optional] 
**MedianAbsoluteDeviation** | Pointer to [**AggMedianAbsoluteDeviation**](AggMedianAbsoluteDeviation.md) |  | [optional] 
**Min** | Pointer to [**AggMetric**](AggMetric.md) |  | [optional] 
**Max** | Pointer to [**AggMetric**](AggMetric.md) |  | [optional] 
**Sum** | Pointer to [**AggMetric**](AggMetric.md) |  | [optional] 
**Avg** | Pointer to [**AggMetric**](AggMetric.md) |  | [optional] 

## Methods

### NewAggregation

`func NewAggregation() *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerms

`func (o *Aggregation) GetTerms() AggTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Aggregation) GetTermsOk() (*AggTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Aggregation) SetTerms(v AggTerms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Aggregation) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetSort

`func (o *Aggregation) GetSort() []interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Aggregation) GetSortOk() (*[]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Aggregation) SetSort(v []interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *Aggregation) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetComposite

`func (o *Aggregation) GetComposite() AggComposite`

GetComposite returns the Composite field if non-nil, zero value otherwise.

### GetCompositeOk

`func (o *Aggregation) GetCompositeOk() (*AggComposite, bool)`

GetCompositeOk returns a tuple with the Composite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposite

`func (o *Aggregation) SetComposite(v AggComposite)`

SetComposite sets Composite field to given value.

### HasComposite

`func (o *Aggregation) HasComposite() bool`

HasComposite returns a boolean if a field has been set.

### GetHistogram

`func (o *Aggregation) GetHistogram() AggHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *Aggregation) GetHistogramOk() (*AggHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *Aggregation) SetHistogram(v AggHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *Aggregation) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetDateHistogram

`func (o *Aggregation) GetDateHistogram() AggDateHistogram`

GetDateHistogram returns the DateHistogram field if non-nil, zero value otherwise.

### GetDateHistogramOk

`func (o *Aggregation) GetDateHistogramOk() (*AggDateHistogram, bool)`

GetDateHistogramOk returns a tuple with the DateHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateHistogram

`func (o *Aggregation) SetDateHistogram(v AggDateHistogram)`

SetDateHistogram sets DateHistogram field to given value.

### HasDateHistogram

`func (o *Aggregation) HasDateHistogram() bool`

HasDateHistogram returns a boolean if a field has been set.

### GetRange

`func (o *Aggregation) GetRange() AggRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Aggregation) GetRangeOk() (*AggRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Aggregation) SetRange(v AggRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Aggregation) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetDateRange

`func (o *Aggregation) GetDateRange() AggRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Aggregation) GetDateRangeOk() (*AggRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Aggregation) SetDateRange(v AggRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Aggregation) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetPercentiles

`func (o *Aggregation) GetPercentiles() AggPercentiles`

GetPercentiles returns the Percentiles field if non-nil, zero value otherwise.

### GetPercentilesOk

`func (o *Aggregation) GetPercentilesOk() (*AggPercentiles, bool)`

GetPercentilesOk returns a tuple with the Percentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentiles

`func (o *Aggregation) SetPercentiles(v AggPercentiles)`

SetPercentiles sets Percentiles field to given value.

### HasPercentiles

`func (o *Aggregation) HasPercentiles() bool`

HasPercentiles returns a boolean if a field has been set.

### GetPercentileRanks

`func (o *Aggregation) GetPercentileRanks() AggPercentileRanks`

GetPercentileRanks returns the PercentileRanks field if non-nil, zero value otherwise.

### GetPercentileRanksOk

`func (o *Aggregation) GetPercentileRanksOk() (*AggPercentileRanks, bool)`

GetPercentileRanksOk returns a tuple with the PercentileRanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileRanks

`func (o *Aggregation) SetPercentileRanks(v AggPercentileRanks)`

SetPercentileRanks sets PercentileRanks field to given value.

### HasPercentileRanks

`func (o *Aggregation) HasPercentileRanks() bool`

HasPercentileRanks returns a boolean if a field has been set.

### GetMedianAbsoluteDeviation

`func (o *Aggregation) GetMedianAbsoluteDeviation() AggMedianAbsoluteDeviation`

GetMedianAbsoluteDeviation returns the MedianAbsoluteDeviation field if non-nil, zero value otherwise.

### GetMedianAbsoluteDeviationOk

`func (o *Aggregation) GetMedianAbsoluteDeviationOk() (*AggMedianAbsoluteDeviation, bool)`

GetMedianAbsoluteDeviationOk returns a tuple with the MedianAbsoluteDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianAbsoluteDeviation

`func (o *Aggregation) SetMedianAbsoluteDeviation(v AggMedianAbsoluteDeviation)`

SetMedianAbsoluteDeviation sets MedianAbsoluteDeviation field to given value.

### HasMedianAbsoluteDeviation

`func (o *Aggregation) HasMedianAbsoluteDeviation() bool`

HasMedianAbsoluteDeviation returns a boolean if a field has been set.

### GetMin

`func (o *Aggregation) GetMin() AggMetric`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Aggregation) GetMinOk() (*AggMetric, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Aggregation) SetMin(v AggMetric)`

SetMin sets Min field to given value.

### HasMin

`func (o *Aggregation) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *Aggregation) GetMax() AggMetric`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Aggregation) GetMaxOk() (*AggMetric, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Aggregation) SetMax(v AggMetric)`

SetMax sets Max field to given value.

### HasMax

`func (o *Aggregation) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetSum

`func (o *Aggregation) GetSum() AggMetric`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *Aggregation) GetSumOk() (*AggMetric, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *Aggregation) SetSum(v AggMetric)`

SetSum sets Sum field to given value.

### HasSum

`func (o *Aggregation) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetAvg

`func (o *Aggregation) GetAvg() AggMetric`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *Aggregation) GetAvgOk() (*AggMetric, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *Aggregation) SetAvg(v AggMetric)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *Aggregation) HasAvg() bool`

HasAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


