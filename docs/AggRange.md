# AggRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Attribute to group by | 
**Ranges** | [**[]ModelRange**](ModelRange.md) | Ordered list of ranges | 
**Keyed** | Pointer to **bool** | Return buckets as an object keyed by range label when true, or as an array when false. Default is false.  | [optional] 

## Methods

### NewAggRange

`func NewAggRange(field string, ranges []ModelRange, ) *AggRange`

NewAggRange instantiates a new AggRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggRangeWithDefaults

`func NewAggRangeWithDefaults() *AggRange`

NewAggRangeWithDefaults instantiates a new AggRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggRange) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggRange) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggRange) SetField(v string)`

SetField sets Field field to given value.


### GetRanges

`func (o *AggRange) GetRanges() []ModelRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AggRange) GetRangesOk() (*[]ModelRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AggRange) SetRanges(v []ModelRange)`

SetRanges sets Ranges field to given value.


### GetKeyed

`func (o *AggRange) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *AggRange) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *AggRange) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *AggRange) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


