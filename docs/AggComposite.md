# AggComposite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Maximum number of composite buckets in the result | [optional] 
**Sources** | Pointer to [**[]map[string]AggCompositeSource**](map[string]AggCompositeSource.md) |  | [optional] 

## Methods

### NewAggComposite

`func NewAggComposite() *AggComposite`

NewAggComposite instantiates a new AggComposite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggCompositeWithDefaults

`func NewAggCompositeWithDefaults() *AggComposite`

NewAggCompositeWithDefaults instantiates a new AggComposite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *AggComposite) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggComposite) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggComposite) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AggComposite) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSources

`func (o *AggComposite) GetSources() []map[string]AggCompositeSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AggComposite) GetSourcesOk() (*[]map[string]AggCompositeSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AggComposite) SetSources(v []map[string]AggCompositeSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AggComposite) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


