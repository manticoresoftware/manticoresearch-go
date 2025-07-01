# AggTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **interface{}** | Name of attribute to aggregate by | 
**Size** | Pointer to **interface{}** | Maximum number of buckets in the result | [optional] 

## Methods

### NewAggTerms

`func NewAggTerms(field interface{}, ) *AggTerms`

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

`func (o *AggTerms) GetField() interface{}`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggTerms) GetFieldOk() (*interface{}, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggTerms) SetField(v interface{})`

SetField sets Field field to given value.


### SetFieldNil

`func (o *AggTerms) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *AggTerms) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetSize

`func (o *AggTerms) GetSize() interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggTerms) GetSizeOk() (*interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggTerms) SetSize(v interface{})`

SetSize sets Size field to given value.

### HasSize

`func (o *AggTerms) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *AggTerms) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *AggTerms) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


