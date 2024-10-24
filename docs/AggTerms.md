# AggTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Name of attribute to aggregate by | 
**Size** | Pointer to **int32** | Maximum number of buckets in the result | [optional] 

## Methods

### NewAggTerms

`func NewAggTerms(field string, ) *AggTerms`

NewAggTerms instantiates a new AggTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggTermsWithDefaults

`func NewAggTermsWithDefaults() *AggTerms`

NewAggTermsWithDefaults instantiates a new AggTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggTerms) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggTerms) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggTerms) SetField(v string)`

SetField sets Field field to given value.


### GetSize

`func (o *AggTerms) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggTerms) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggTerms) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AggTerms) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


