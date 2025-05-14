# Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FragmentSize** | Pointer to **interface{}** | Maximum size of the text fragments in highlighted snippets per field | [optional] 
**Limit** | Pointer to **interface{}** | Maximum size of snippets per field | [optional] 
**LimitSnippets** | Pointer to **interface{}** | Maximum number of snippets per field | [optional] 
**LimitWords** | Pointer to **interface{}** | Maximum number of words per field | [optional] 
**NumberOfFragments** | Pointer to **interface{}** | Total number of highlighted fragments per field | [optional] 
**AfterMatch** | Pointer to **string** | Text inserted after the matched term, typically used for HTML formatting | [optional] [default to "</strong>"]
**AllowEmpty** | Pointer to **bool** | Permits an empty string to be returned as the highlighting result. Otherwise, the beginning of the original text would be returned | [optional] 
**Around** | Pointer to **int32** | Number of words around the match to include in the highlight | [optional] 
**BeforeMatch** | Pointer to **string** | Text inserted before the match, typically used for HTML formatting | [optional] [default to "<strong>"]
**EmitZones** | Pointer to **bool** | Emits an HTML tag with the enclosing zone name before each highlighted snippet | [optional] 
**Encoder** | Pointer to **string** | If set to &#39;html&#39;, retains HTML markup when highlighting | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**ForceAllWords** | Pointer to **bool** | Ignores the length limit until the result includes all keywords | [optional] 
**ForceSnippets** | Pointer to **bool** | Forces snippet generation even if limits allow highlighting the entire text | [optional] 
**HighlightQuery** | Pointer to [**NullableQueryFilter**](QueryFilter.md) |  | [optional] 
**HtmlStripMode** | Pointer to **string** | Defines the mode for handling HTML markup in the highlight | [optional] 
**LimitsPerField** | Pointer to **bool** | Determines whether the &#39;limit&#39;, &#39;limit_words&#39;, and &#39;limit_snippets&#39; options operate as individual limits in each field of the document | [optional] 
**NoMatchSize** | Pointer to **int32** | If set to 1, allows an empty string to be returned as a highlighting result | [optional] 
**Order** | Pointer to **string** | Sets the sorting order of highlighted snippets | [optional] 
**PreTags** | Pointer to **string** | Text inserted before each highlighted snippet | [optional] [default to "<strong>"]
**PostTags** | Pointer to **string** | Text inserted after each highlighted snippet | [optional] [default to "</strong>"]
**StartSnippetId** | Pointer to **int32** | Sets the starting value of the %SNIPPET_ID% macro | [optional] 
**UseBoundaries** | Pointer to **bool** | Defines whether to additionally break snippets by phrase boundary characters | [optional] 

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

### GetFragmentSize

`func (o *Highlight) GetFragmentSize() interface{}`

GetFragmentSize returns the FragmentSize field if non-nil, zero value otherwise.

### GetFragmentSizeOk

`func (o *Highlight) GetFragmentSizeOk() (*interface{}, bool)`

GetFragmentSizeOk returns a tuple with the FragmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentSize

`func (o *Highlight) SetFragmentSize(v interface{})`

SetFragmentSize sets FragmentSize field to given value.

### HasFragmentSize

`func (o *Highlight) HasFragmentSize() bool`

HasFragmentSize returns a boolean if a field has been set.

### SetFragmentSizeNil

`func (o *Highlight) SetFragmentSizeNil(b bool)`

 SetFragmentSizeNil sets the value for FragmentSize to be an explicit nil

### UnsetFragmentSize
`func (o *Highlight) UnsetFragmentSize()`

UnsetFragmentSize ensures that no value is present for FragmentSize, not even an explicit nil
### GetLimit

`func (o *Highlight) GetLimit() interface{}`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Highlight) GetLimitOk() (*interface{}, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Highlight) SetLimit(v interface{})`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Highlight) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Highlight) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Highlight) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetLimitSnippets

`func (o *Highlight) GetLimitSnippets() interface{}`

GetLimitSnippets returns the LimitSnippets field if non-nil, zero value otherwise.

### GetLimitSnippetsOk

`func (o *Highlight) GetLimitSnippetsOk() (*interface{}, bool)`

GetLimitSnippetsOk returns a tuple with the LimitSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitSnippets

`func (o *Highlight) SetLimitSnippets(v interface{})`

SetLimitSnippets sets LimitSnippets field to given value.

### HasLimitSnippets

`func (o *Highlight) HasLimitSnippets() bool`

HasLimitSnippets returns a boolean if a field has been set.

### SetLimitSnippetsNil

`func (o *Highlight) SetLimitSnippetsNil(b bool)`

 SetLimitSnippetsNil sets the value for LimitSnippets to be an explicit nil

### UnsetLimitSnippets
`func (o *Highlight) UnsetLimitSnippets()`

UnsetLimitSnippets ensures that no value is present for LimitSnippets, not even an explicit nil
### GetLimitWords

`func (o *Highlight) GetLimitWords() interface{}`

GetLimitWords returns the LimitWords field if non-nil, zero value otherwise.

### GetLimitWordsOk

`func (o *Highlight) GetLimitWordsOk() (*interface{}, bool)`

GetLimitWordsOk returns a tuple with the LimitWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitWords

`func (o *Highlight) SetLimitWords(v interface{})`

SetLimitWords sets LimitWords field to given value.

### HasLimitWords

`func (o *Highlight) HasLimitWords() bool`

HasLimitWords returns a boolean if a field has been set.

### SetLimitWordsNil

`func (o *Highlight) SetLimitWordsNil(b bool)`

 SetLimitWordsNil sets the value for LimitWords to be an explicit nil

### UnsetLimitWords
`func (o *Highlight) UnsetLimitWords()`

UnsetLimitWords ensures that no value is present for LimitWords, not even an explicit nil
### GetNumberOfFragments

`func (o *Highlight) GetNumberOfFragments() interface{}`

GetNumberOfFragments returns the NumberOfFragments field if non-nil, zero value otherwise.

### GetNumberOfFragmentsOk

`func (o *Highlight) GetNumberOfFragmentsOk() (*interface{}, bool)`

GetNumberOfFragmentsOk returns a tuple with the NumberOfFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFragments

`func (o *Highlight) SetNumberOfFragments(v interface{})`

SetNumberOfFragments sets NumberOfFragments field to given value.

### HasNumberOfFragments

`func (o *Highlight) HasNumberOfFragments() bool`

HasNumberOfFragments returns a boolean if a field has been set.

### SetNumberOfFragmentsNil

`func (o *Highlight) SetNumberOfFragmentsNil(b bool)`

 SetNumberOfFragmentsNil sets the value for NumberOfFragments to be an explicit nil

### UnsetNumberOfFragments
`func (o *Highlight) UnsetNumberOfFragments()`

UnsetNumberOfFragments ensures that no value is present for NumberOfFragments, not even an explicit nil
### GetAfterMatch

`func (o *Highlight) GetAfterMatch() string`

GetAfterMatch returns the AfterMatch field if non-nil, zero value otherwise.

### GetAfterMatchOk

`func (o *Highlight) GetAfterMatchOk() (*string, bool)`

GetAfterMatchOk returns a tuple with the AfterMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterMatch

`func (o *Highlight) SetAfterMatch(v string)`

SetAfterMatch sets AfterMatch field to given value.

### HasAfterMatch

`func (o *Highlight) HasAfterMatch() bool`

HasAfterMatch returns a boolean if a field has been set.

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

### GetBeforeMatch

`func (o *Highlight) GetBeforeMatch() string`

GetBeforeMatch returns the BeforeMatch field if non-nil, zero value otherwise.

### GetBeforeMatchOk

`func (o *Highlight) GetBeforeMatchOk() (*string, bool)`

GetBeforeMatchOk returns a tuple with the BeforeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeMatch

`func (o *Highlight) SetBeforeMatch(v string)`

SetBeforeMatch sets BeforeMatch field to given value.

### HasBeforeMatch

`func (o *Highlight) HasBeforeMatch() bool`

HasBeforeMatch returns a boolean if a field has been set.

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

### GetFields

`func (o *Highlight) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Highlight) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Highlight) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *Highlight) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *Highlight) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *Highlight) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
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

### GetHighlightQuery

`func (o *Highlight) GetHighlightQuery() QueryFilter`

GetHighlightQuery returns the HighlightQuery field if non-nil, zero value otherwise.

### GetHighlightQueryOk

`func (o *Highlight) GetHighlightQueryOk() (*QueryFilter, bool)`

GetHighlightQueryOk returns a tuple with the HighlightQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightQuery

`func (o *Highlight) SetHighlightQuery(v QueryFilter)`

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

### GetOrder

`func (o *Highlight) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Highlight) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Highlight) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Highlight) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


