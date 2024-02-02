# Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fieldnames** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to [**[]HighlightField**](HighlightField.md) |  | [optional] 
**Encoder** | Pointer to **string** |  | [optional] 
**HighlightQuery** | Pointer to **map[string]interface{}** |  | [optional] 
**PreTags** | Pointer to **string** |  | [optional] [default to "<strong>"]
**PostTags** | Pointer to **string** |  | [optional] [default to "</strong>"]
**NoMatchSize** | Pointer to **int32** |  | [optional] 
**FragmentSize** | Pointer to **int32** |  | [optional] [default to 256]
**NumberOfFragments** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 256]
**LimitWords** | Pointer to **int32** |  | [optional] [default to 0]
**LimitSnippets** | Pointer to **int32** |  | [optional] [default to 0]
**LimitsPerField** | Pointer to **bool** |  | [optional] [default to false]
**UseBoundaries** | Pointer to **bool** |  | [optional] [default to false]
**ForceAllWords** | Pointer to **bool** |  | [optional] [default to false]
**AllowEmpty** | Pointer to **bool** |  | [optional] [default to false]
**EmitZones** | Pointer to **bool** |  | [optional] [default to false]
**ForceSnippets** | Pointer to **bool** |  | [optional] [default to false]
**Around** | Pointer to **int32** |  | [optional] [default to 5]
**StartSnippetId** | Pointer to **int32** |  | [optional] [default to 1]
**HtmlStripMode** | Pointer to **string** |  | [optional] 
**SnippetBoundary** | Pointer to **string** |  | [optional] 

## Methods

### NewHighlight

`func NewHighlight() *Highlight`

NewHighlight instantiates a new Highlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightWithDefaults

`func NewHighlightWithDefaults() *Highlight`

NewHighlightWithDefaults instantiates a new Highlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldnames

`func (o *Highlight) GetFieldnames() []string`

GetFieldnames returns the Fieldnames field if non-nil, zero value otherwise.

### GetFieldnamesOk

`func (o *Highlight) GetFieldnamesOk() (*[]string, bool)`

GetFieldnamesOk returns a tuple with the Fieldnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldnames

`func (o *Highlight) SetFieldnames(v []string)`

SetFieldnames sets Fieldnames field to given value.

### HasFieldnames

`func (o *Highlight) HasFieldnames() bool`

HasFieldnames returns a boolean if a field has been set.

### GetFields

`func (o *Highlight) GetFields() []HighlightField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Highlight) GetFieldsOk() (*[]HighlightField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Highlight) SetFields(v []HighlightField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Highlight) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetEncoder

`func (o *Highlight) GetEncoder() string`

GetEncoder returns the Encoder field if non-nil, zero value otherwise.

### GetEncoderOk

`func (o *Highlight) GetEncoderOk() (*string, bool)`

GetEncoderOk returns a tuple with the Encoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoder

`func (o *Highlight) SetEncoder(v string)`

SetEncoder sets Encoder field to given value.

### HasEncoder

`func (o *Highlight) HasEncoder() bool`

HasEncoder returns a boolean if a field has been set.

### GetHighlightQuery

`func (o *Highlight) GetHighlightQuery() map[string]interface{}`

GetHighlightQuery returns the HighlightQuery field if non-nil, zero value otherwise.

### GetHighlightQueryOk

`func (o *Highlight) GetHighlightQueryOk() (*map[string]interface{}, bool)`

GetHighlightQueryOk returns a tuple with the HighlightQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightQuery

`func (o *Highlight) SetHighlightQuery(v map[string]interface{})`

SetHighlightQuery sets HighlightQuery field to given value.

### HasHighlightQuery

`func (o *Highlight) HasHighlightQuery() bool`

HasHighlightQuery returns a boolean if a field has been set.

### SetHighlightQueryNil

`func (o *Highlight) SetHighlightQueryNil(b bool)`

 SetHighlightQueryNil sets the value for HighlightQuery to be an explicit nil

### UnsetHighlightQuery
`func (o *Highlight) UnsetHighlightQuery()`

UnsetHighlightQuery ensures that no value is present for HighlightQuery, not even an explicit nil
### GetPreTags

`func (o *Highlight) GetPreTags() string`

GetPreTags returns the PreTags field if non-nil, zero value otherwise.

### GetPreTagsOk

`func (o *Highlight) GetPreTagsOk() (*string, bool)`

GetPreTagsOk returns a tuple with the PreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreTags

`func (o *Highlight) SetPreTags(v string)`

SetPreTags sets PreTags field to given value.

### HasPreTags

`func (o *Highlight) HasPreTags() bool`

HasPreTags returns a boolean if a field has been set.

### GetPostTags

`func (o *Highlight) GetPostTags() string`

GetPostTags returns the PostTags field if non-nil, zero value otherwise.

### GetPostTagsOk

`func (o *Highlight) GetPostTagsOk() (*string, bool)`

GetPostTagsOk returns a tuple with the PostTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTags

`func (o *Highlight) SetPostTags(v string)`

SetPostTags sets PostTags field to given value.

### HasPostTags

`func (o *Highlight) HasPostTags() bool`

HasPostTags returns a boolean if a field has been set.

### GetNoMatchSize

`func (o *Highlight) GetNoMatchSize() int32`

GetNoMatchSize returns the NoMatchSize field if non-nil, zero value otherwise.

### GetNoMatchSizeOk

`func (o *Highlight) GetNoMatchSizeOk() (*int32, bool)`

GetNoMatchSizeOk returns a tuple with the NoMatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMatchSize

`func (o *Highlight) SetNoMatchSize(v int32)`

SetNoMatchSize sets NoMatchSize field to given value.

### HasNoMatchSize

`func (o *Highlight) HasNoMatchSize() bool`

HasNoMatchSize returns a boolean if a field has been set.

### GetFragmentSize

`func (o *Highlight) GetFragmentSize() int32`

GetFragmentSize returns the FragmentSize field if non-nil, zero value otherwise.

### GetFragmentSizeOk

`func (o *Highlight) GetFragmentSizeOk() (*int32, bool)`

GetFragmentSizeOk returns a tuple with the FragmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentSize

`func (o *Highlight) SetFragmentSize(v int32)`

SetFragmentSize sets FragmentSize field to given value.

### HasFragmentSize

`func (o *Highlight) HasFragmentSize() bool`

HasFragmentSize returns a boolean if a field has been set.

### GetNumberOfFragments

`func (o *Highlight) GetNumberOfFragments() int32`

GetNumberOfFragments returns the NumberOfFragments field if non-nil, zero value otherwise.

### GetNumberOfFragmentsOk

`func (o *Highlight) GetNumberOfFragmentsOk() (*int32, bool)`

GetNumberOfFragmentsOk returns a tuple with the NumberOfFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFragments

`func (o *Highlight) SetNumberOfFragments(v int32)`

SetNumberOfFragments sets NumberOfFragments field to given value.

### HasNumberOfFragments

`func (o *Highlight) HasNumberOfFragments() bool`

HasNumberOfFragments returns a boolean if a field has been set.

### GetLimit

`func (o *Highlight) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Highlight) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Highlight) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Highlight) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitWords

`func (o *Highlight) GetLimitWords() int32`

GetLimitWords returns the LimitWords field if non-nil, zero value otherwise.

### GetLimitWordsOk

`func (o *Highlight) GetLimitWordsOk() (*int32, bool)`

GetLimitWordsOk returns a tuple with the LimitWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWords

`func (o *Highlight) SetLimitWords(v int32)`

SetLimitWords sets LimitWords field to given value.

### HasLimitWords

`func (o *Highlight) HasLimitWords() bool`

HasLimitWords returns a boolean if a field has been set.

### GetLimitSnippets

`func (o *Highlight) GetLimitSnippets() int32`

GetLimitSnippets returns the LimitSnippets field if non-nil, zero value otherwise.

### GetLimitSnippetsOk

`func (o *Highlight) GetLimitSnippetsOk() (*int32, bool)`

GetLimitSnippetsOk returns a tuple with the LimitSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSnippets

`func (o *Highlight) SetLimitSnippets(v int32)`

SetLimitSnippets sets LimitSnippets field to given value.

### HasLimitSnippets

`func (o *Highlight) HasLimitSnippets() bool`

HasLimitSnippets returns a boolean if a field has been set.

### GetLimitsPerField

`func (o *Highlight) GetLimitsPerField() bool`

GetLimitsPerField returns the LimitsPerField field if non-nil, zero value otherwise.

### GetLimitsPerFieldOk

`func (o *Highlight) GetLimitsPerFieldOk() (*bool, bool)`

GetLimitsPerFieldOk returns a tuple with the LimitsPerField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsPerField

`func (o *Highlight) SetLimitsPerField(v bool)`

SetLimitsPerField sets LimitsPerField field to given value.

### HasLimitsPerField

`func (o *Highlight) HasLimitsPerField() bool`

HasLimitsPerField returns a boolean if a field has been set.

### GetUseBoundaries

`func (o *Highlight) GetUseBoundaries() bool`

GetUseBoundaries returns the UseBoundaries field if non-nil, zero value otherwise.

### GetUseBoundariesOk

`func (o *Highlight) GetUseBoundariesOk() (*bool, bool)`

GetUseBoundariesOk returns a tuple with the UseBoundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBoundaries

`func (o *Highlight) SetUseBoundaries(v bool)`

SetUseBoundaries sets UseBoundaries field to given value.

### HasUseBoundaries

`func (o *Highlight) HasUseBoundaries() bool`

HasUseBoundaries returns a boolean if a field has been set.

### GetForceAllWords

`func (o *Highlight) GetForceAllWords() bool`

GetForceAllWords returns the ForceAllWords field if non-nil, zero value otherwise.

### GetForceAllWordsOk

`func (o *Highlight) GetForceAllWordsOk() (*bool, bool)`

GetForceAllWordsOk returns a tuple with the ForceAllWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAllWords

`func (o *Highlight) SetForceAllWords(v bool)`

SetForceAllWords sets ForceAllWords field to given value.

### HasForceAllWords

`func (o *Highlight) HasForceAllWords() bool`

HasForceAllWords returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *Highlight) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *Highlight) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *Highlight) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *Highlight) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetEmitZones

`func (o *Highlight) GetEmitZones() bool`

GetEmitZones returns the EmitZones field if non-nil, zero value otherwise.

### GetEmitZonesOk

`func (o *Highlight) GetEmitZonesOk() (*bool, bool)`

GetEmitZonesOk returns a tuple with the EmitZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitZones

`func (o *Highlight) SetEmitZones(v bool)`

SetEmitZones sets EmitZones field to given value.

### HasEmitZones

`func (o *Highlight) HasEmitZones() bool`

HasEmitZones returns a boolean if a field has been set.

### GetForceSnippets

`func (o *Highlight) GetForceSnippets() bool`

GetForceSnippets returns the ForceSnippets field if non-nil, zero value otherwise.

### GetForceSnippetsOk

`func (o *Highlight) GetForceSnippetsOk() (*bool, bool)`

GetForceSnippetsOk returns a tuple with the ForceSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSnippets

`func (o *Highlight) SetForceSnippets(v bool)`

SetForceSnippets sets ForceSnippets field to given value.

### HasForceSnippets

`func (o *Highlight) HasForceSnippets() bool`

HasForceSnippets returns a boolean if a field has been set.

### GetAround

`func (o *Highlight) GetAround() int32`

GetAround returns the Around field if non-nil, zero value otherwise.

### GetAroundOk

`func (o *Highlight) GetAroundOk() (*int32, bool)`

GetAroundOk returns a tuple with the Around field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAround

`func (o *Highlight) SetAround(v int32)`

SetAround sets Around field to given value.

### HasAround

`func (o *Highlight) HasAround() bool`

HasAround returns a boolean if a field has been set.

### GetStartSnippetId

`func (o *Highlight) GetStartSnippetId() int32`

GetStartSnippetId returns the StartSnippetId field if non-nil, zero value otherwise.

### GetStartSnippetIdOk

`func (o *Highlight) GetStartSnippetIdOk() (*int32, bool)`

GetStartSnippetIdOk returns a tuple with the StartSnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSnippetId

`func (o *Highlight) SetStartSnippetId(v int32)`

SetStartSnippetId sets StartSnippetId field to given value.

### HasStartSnippetId

`func (o *Highlight) HasStartSnippetId() bool`

HasStartSnippetId returns a boolean if a field has been set.

### GetHtmlStripMode

`func (o *Highlight) GetHtmlStripMode() string`

GetHtmlStripMode returns the HtmlStripMode field if non-nil, zero value otherwise.

### GetHtmlStripModeOk

`func (o *Highlight) GetHtmlStripModeOk() (*string, bool)`

GetHtmlStripModeOk returns a tuple with the HtmlStripMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlStripMode

`func (o *Highlight) SetHtmlStripMode(v string)`

SetHtmlStripMode sets HtmlStripMode field to given value.

### HasHtmlStripMode

`func (o *Highlight) HasHtmlStripMode() bool`

HasHtmlStripMode returns a boolean if a field has been set.

### GetSnippetBoundary

`func (o *Highlight) GetSnippetBoundary() string`

GetSnippetBoundary returns the SnippetBoundary field if non-nil, zero value otherwise.

### GetSnippetBoundaryOk

`func (o *Highlight) GetSnippetBoundaryOk() (*string, bool)`

GetSnippetBoundaryOk returns a tuple with the SnippetBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetBoundary

`func (o *Highlight) SetSnippetBoundary(v string)`

SetSnippetBoundary sets SnippetBoundary field to given value.

### HasSnippetBoundary

`func (o *Highlight) HasSnippetBoundary() bool`

HasSnippetBoundary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


