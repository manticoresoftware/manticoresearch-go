# HighlightField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] [default to 256]
**LimitWords** | Pointer to **int32** |  | [optional] [default to 0]
**LimitSnippets** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewHighlightField

`func NewHighlightField(name string, ) *HighlightField`

NewHighlightField instantiates a new HighlightField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightFieldWithDefaults

`func NewHighlightFieldWithDefaults() *HighlightField`

NewHighlightFieldWithDefaults instantiates a new HighlightField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HighlightField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HighlightField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HighlightField) SetName(v string)`

SetName sets Name field to given value.


### GetLimit

`func (o *HighlightField) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *HighlightField) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *HighlightField) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *HighlightField) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitWords

`func (o *HighlightField) GetLimitWords() int32`

GetLimitWords returns the LimitWords field if non-nil, zero value otherwise.

### GetLimitWordsOk

`func (o *HighlightField) GetLimitWordsOk() (*int32, bool)`

GetLimitWordsOk returns a tuple with the LimitWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWords

`func (o *HighlightField) SetLimitWords(v int32)`

SetLimitWords sets LimitWords field to given value.

### HasLimitWords

`func (o *HighlightField) HasLimitWords() bool`

HasLimitWords returns a boolean if a field has been set.

### GetLimitSnippets

`func (o *HighlightField) GetLimitSnippets() int32`

GetLimitSnippets returns the LimitSnippets field if non-nil, zero value otherwise.

### GetLimitSnippetsOk

`func (o *HighlightField) GetLimitSnippetsOk() (*int32, bool)`

GetLimitSnippetsOk returns a tuple with the LimitSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSnippets

`func (o *HighlightField) SetLimitSnippets(v int32)`

SetLimitSnippets sets LimitSnippets field to given value.

### HasLimitSnippets

`func (o *HighlightField) HasLimitSnippets() bool`

HasLimitSnippets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


