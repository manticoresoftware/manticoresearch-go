# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | Pointer to [**AggTerms**](AggTerms.md) |  | [optional] 
**Sort** | Pointer to **[]interface{}** |  | [optional] 
**Composite** | Pointer to [**AggComposite**](AggComposite.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


