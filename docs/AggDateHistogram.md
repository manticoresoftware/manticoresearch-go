# AggDateHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **interface{}** | Field to group by | 
**Interval** | **interface{}** | Interval of the histogram values | 
**Offset** | Pointer to **interface{}** | Offset of the histogram values. Default value is 0. | [optional] 
**Keyed** | Pointer to **interface{}** | Flag that defines if a search response will be a dictionary with the bucket keys. Default value is false. | [optional] 

## Methods

### NewAggDateHistogram

`func NewAggDateHistogram(field interface{}, interval interface{}, ) *AggDateHistogram`

NewAggDateHistogram instantiates a new AggDateHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggDateHistogramWithDefaults

`func NewAggDateHistogramWithDefaults() *AggDateHistogram`

NewAggDateHistogramWithDefaults instantiates a new AggDateHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggDateHistogram) GetField() interface{}`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggDateHistogram) GetFieldOk() (*interface{}, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggDateHistogram) SetField(v interface{})`

SetField sets Field field to given value.


### SetFieldNil

`func (o *AggDateHistogram) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *AggDateHistogram) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetInterval

`func (o *AggDateHistogram) GetInterval() interface{}`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AggDateHistogram) GetIntervalOk() (*interface{}, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AggDateHistogram) SetInterval(v interface{})`

SetInterval sets Interval field to given value.


### SetIntervalNil

`func (o *AggDateHistogram) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *AggDateHistogram) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetOffset

`func (o *AggDateHistogram) GetOffset() interface{}`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AggDateHistogram) GetOffsetOk() (*interface{}, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AggDateHistogram) SetOffset(v interface{})`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AggDateHistogram) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *AggDateHistogram) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *AggDateHistogram) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetKeyed

`func (o *AggDateHistogram) GetKeyed() interface{}`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *AggDateHistogram) GetKeyedOk() (*interface{}, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *AggDateHistogram) SetKeyed(v interface{})`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *AggDateHistogram) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### SetKeyedNil

`func (o *AggDateHistogram) SetKeyedNil(b bool)`

 SetKeyedNil sets the value for Keyed to be an explicit nil

### UnsetKeyed
`func (o *AggDateHistogram) UnsetKeyed()`

UnsetKeyed ensures that no value is present for Keyed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


