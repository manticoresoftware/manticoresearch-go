# AggregationTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Attribute Name to Aggregate | [optional] 
**Size** | Pointer to **int32** | Maximum Number of Buckets in the Result | [optional] 

## Methods

### NewAggregationTerms

`func NewAggregationTerms() *AggregationTerms`

NewAggregationTerms instantiates a new AggregationTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationTermsWithDefaults

`func NewAggregationTermsWithDefaults() *AggregationTerms`

NewAggregationTermsWithDefaults instantiates a new AggregationTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggregationTerms) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregationTerms) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregationTerms) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *AggregationTerms) HasField() bool`

HasField returns a boolean if a field has been set.

### GetSize

`func (o *AggregationTerms) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggregationTerms) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggregationTerms) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AggregationTerms) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


