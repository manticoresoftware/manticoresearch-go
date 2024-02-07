# SortMultiple

Query sort expression for multiple attributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attrs** | **map[string]interface{}** |  | 
**Replace** | **bool** |  | 

## Methods

### NewSortMultiple

`func NewSortMultiple(attrs map[string]interface{}, replace bool, ) *SortMultiple`

NewSortMultiple instantiates a new SortMultiple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortMultipleWithDefaults

`func NewSortMultipleWithDefaults() *SortMultiple`

NewSortMultipleWithDefaults instantiates a new SortMultiple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrs

`func (o *SortMultiple) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *SortMultiple) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *SortMultiple) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.


### GetReplace

`func (o *SortMultiple) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *SortMultiple) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *SortMultiple) SetReplace(v bool)`

SetReplace sets Replace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


