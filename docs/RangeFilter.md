# RangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to [**map[string]RangeFilterRangeValue**](RangeFilterRangeValue.md) |  | [optional] 

## Methods

### NewRangeFilter

`func NewRangeFilter() *RangeFilter`

NewRangeFilter instantiates a new RangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeFilterWithDefaults

`func NewRangeFilterWithDefaults() *RangeFilter`

NewRangeFilterWithDefaults instantiates a new RangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *RangeFilter) GetRange() map[string]RangeFilterRangeValue`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *RangeFilter) GetRangeOk() (*map[string]RangeFilterRangeValue, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *RangeFilter) SetRange(v map[string]RangeFilterRangeValue)`

SetRange sets Range field to given value.

### HasRange

`func (o *RangeFilter) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


