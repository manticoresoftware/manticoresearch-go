# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | Pointer to [**AggregationTerms**](AggregationTerms.md) |  | [optional] 
**Sort** | Pointer to [**[]map[string]AggregationSortInnerValue**](map[string]AggregationSortInnerValue.md) |  | [optional] 
**Composite** | Pointer to [**AggregationComposite**](AggregationComposite.md) |  | [optional] 

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

`func (o *Aggregation) GetTerms() AggregationTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Aggregation) GetTermsOk() (*AggregationTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Aggregation) SetTerms(v AggregationTerms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Aggregation) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetSort

`func (o *Aggregation) GetSort() []map[string]AggregationSortInnerValue`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Aggregation) GetSortOk() (*[]map[string]AggregationSortInnerValue, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Aggregation) SetSort(v []map[string]AggregationSortInnerValue)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Aggregation) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetComposite

`func (o *Aggregation) GetComposite() AggregationComposite`

GetComposite returns the Composite field if non-nil, zero value otherwise.

### GetCompositeOk

`func (o *Aggregation) GetCompositeOk() (*AggregationComposite, bool)`

GetCompositeOk returns a tuple with the Composite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposite

`func (o *Aggregation) SetComposite(v AggregationComposite)`

SetComposite sets Composite field to given value.

### HasComposite

`func (o *Aggregation) HasComposite() bool`

HasComposite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


