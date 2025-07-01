# Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FragmentSize** | Pointer to **interface{}** | Maximum size of the text fragments in highlighted snippets per field | [optional] 
**Limit** | Pointer to **interface{}** | Maximum size of snippets per field | [optional] 
**LimitSnippets** | Pointer to **interface{}** | Maximum number of snippets per field | [optional] 
**LimitWords** | Pointer to **interface{}** | Maximum number of words per field | [optional] 
**NumberOfFragments** | Pointer to **interface{}** | Total number of highlighted fragments per field | [optional] 
**AfterMatch** | Pointer to **interface{}** | Text inserted after the matched term, typically used for HTML formatting | [optional] [default to </strong>]
**AllowEmpty** | Pointer to **interface{}** | Permits an empty string to be returned as the highlighting result. Otherwise, the beginning of the original text would be returned | [optional] 
**Around** | Pointer to **interface{}** | Number of words around the match to include in the highlight | [optional] 
**BeforeMatch** | Pointer to **interface{}** | Text inserted before the match, typically used for HTML formatting | [optional] [default to <strong>]
**EmitZones** | Pointer to **interface{}** | Emits an HTML tag with the enclosing zone name before each highlighted snippet | [optional] 
**Encoder** | Pointer to **interface{}** | If set to &#39;html&#39;, retains HTML markup when highlighting | [optional] 
**Fields** | Pointer to [**HighlightFields**](HighlightFields.md) |  | [optional] 
**ForceAllWords** | Pointer to **interface{}** | Ignores the length limit until the result includes all keywords | [optional] 
**ForceSnippets** | Pointer to **interface{}** | Forces snippet generation even if limits allow highlighting the entire text | [optional] 
**HighlightQuery** | Pointer to [**HighlightAllOfHighlightQuery**](HighlightAllOfHighlightQuery.md) |  | [optional] 
**HtmlStripMode** | Pointer to **interface{}** | Defines the mode for handling HTML markup in the highlight | [optional] 
**LimitsPerField** | Pointer to **interface{}** | Determines whether the &#39;limit&#39;, &#39;limit_words&#39;, and &#39;limit_snippets&#39; options operate as individual limits in each field of the document | [optional] 
**NoMatchSize** | Pointer to **interface{}** | If set to 1, allows an empty string to be returned as a highlighting result | [optional] 
**Order** | Pointer to **interface{}** | Sets the sorting order of highlighted snippets | [optional] 
**PreTags** | Pointer to **interface{}** | Text inserted before each highlighted snippet | [optional] [default to <strong>]
**PostTags** | Pointer to **interface{}** | Text inserted after each highlighted snippet | [optional] [default to </strong>]
**StartSnippetId** | Pointer to **interface{}** | Sets the starting value of the %SNIPPET_ID% macro | [optional] 
**UseBoundaries** | Pointer to **interface{}** | Defines whether to additionally break snippets by phrase boundary characters | [optional] 

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

`func (o *Highlight) GetAfterMatch() interface{}`

GetAfterMatch returns the AfterMatch field if non-nil, zero value otherwise.

### GetAfterMatchOk

`func (o *Highlight) GetAfterMatchOk() (*interface{}, bool)`

GetAfterMatchOk returns a tuple with the AfterMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterMatch

`func (o *Highlight) SetAfterMatch(v interface{})`

SetAfterMatch sets AfterMatch field to given value.

### HasAfterMatch

`func (o *Highlight) HasAfterMatch() bool`

HasAfterMatch returns a boolean if a field has been set.

### SetAfterMatchNil

`func (o *Highlight) SetAfterMatchNil(b bool)`

 SetAfterMatchNil sets the value for AfterMatch to be an explicit nil

### UnsetAfterMatch
`func (o *Highlight) UnsetAfterMatch()`

UnsetAfterMatch ensures that no value is present for AfterMatch, not even an explicit nil
### GetAllowEmpty

`func (o *Highlight) GetAllowEmpty() interface{}`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *Highlight) GetAllowEmptyOk() (*interface{}, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *Highlight) SetAllowEmpty(v interface{})`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *Highlight) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### SetAllowEmptyNil

`func (o *Highlight) SetAllowEmptyNil(b bool)`

 SetAllowEmptyNil sets the value for AllowEmpty to be an explicit nil

### UnsetAllowEmpty
`func (o *Highlight) UnsetAllowEmpty()`

UnsetAllowEmpty ensures that no value is present for AllowEmpty, not even an explicit nil
### GetAround

`func (o *Highlight) GetAround() interface{}`

GetAround returns the Around field if non-nil, zero value otherwise.

### GetAroundOk

`func (o *Highlight) GetAroundOk() (*interface{}, bool)`

GetAroundOk returns a tuple with the Around field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAround

`func (o *Highlight) SetAround(v interface{})`

SetAround sets Around field to given value.

### HasAround

`func (o *Highlight) HasAround() bool`

HasAround returns a boolean if a field has been set.

### SetAroundNil

`func (o *Highlight) SetAroundNil(b bool)`

 SetAroundNil sets the value for Around to be an explicit nil

### UnsetAround
`func (o *Highlight) UnsetAround()`

UnsetAround ensures that no value is present for Around, not even an explicit nil
### GetBeforeMatch

`func (o *Highlight) GetBeforeMatch() interface{}`

GetBeforeMatch returns the BeforeMatch field if non-nil, zero value otherwise.

### GetBeforeMatchOk

`func (o *Highlight) GetBeforeMatchOk() (*interface{}, bool)`

GetBeforeMatchOk returns a tuple with the BeforeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeMatch

`func (o *Highlight) SetBeforeMatch(v interface{})`

SetBeforeMatch sets BeforeMatch field to given value.

### HasBeforeMatch

`func (o *Highlight) HasBeforeMatch() bool`

HasBeforeMatch returns a boolean if a field has been set.

### SetBeforeMatchNil

`func (o *Highlight) SetBeforeMatchNil(b bool)`

 SetBeforeMatchNil sets the value for BeforeMatch to be an explicit nil

### UnsetBeforeMatch
`func (o *Highlight) UnsetBeforeMatch()`

UnsetBeforeMatch ensures that no value is present for BeforeMatch, not even an explicit nil
### GetEmitZones

`func (o *Highlight) GetEmitZones() interface{}`

GetEmitZones returns the EmitZones field if non-nil, zero value otherwise.

### GetEmitZonesOk

`func (o *Highlight) GetEmitZonesOk() (*interface{}, bool)`

GetEmitZonesOk returns a tuple with the EmitZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitZones

`func (o *Highlight) SetEmitZones(v interface{})`

SetEmitZones sets EmitZones field to given value.

### HasEmitZones

`func (o *Highlight) HasEmitZones() bool`

HasEmitZones returns a boolean if a field has been set.

### SetEmitZonesNil

`func (o *Highlight) SetEmitZonesNil(b bool)`

 SetEmitZonesNil sets the value for EmitZones to be an explicit nil

### UnsetEmitZones
`func (o *Highlight) UnsetEmitZones()`

UnsetEmitZones ensures that no value is present for EmitZones, not even an explicit nil
### GetEncoder

`func (o *Highlight) GetEncoder() interface{}`

GetEncoder returns the Encoder field if non-nil, zero value otherwise.

### GetEncoderOk

`func (o *Highlight) GetEncoderOk() (*interface{}, bool)`

GetEncoderOk returns a tuple with the Encoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoder

`func (o *Highlight) SetEncoder(v interface{})`

SetEncoder sets Encoder field to given value.

### HasEncoder

`func (o *Highlight) HasEncoder() bool`

HasEncoder returns a boolean if a field has been set.

### SetEncoderNil

`func (o *Highlight) SetEncoderNil(b bool)`

 SetEncoderNil sets the value for Encoder to be an explicit nil

### UnsetEncoder
`func (o *Highlight) UnsetEncoder()`

UnsetEncoder ensures that no value is present for Encoder, not even an explicit nil
### GetFields

`func (o *Highlight) GetFields() HighlightFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Highlight) GetFieldsOk() (*HighlightFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Highlight) SetFields(v HighlightFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Highlight) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetForceAllWords

`func (o *Highlight) GetForceAllWords() interface{}`

GetForceAllWords returns the ForceAllWords field if non-nil, zero value otherwise.

### GetForceAllWordsOk

`func (o *Highlight) GetForceAllWordsOk() (*interface{}, bool)`

GetForceAllWordsOk returns a tuple with the ForceAllWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAllWords

`func (o *Highlight) SetForceAllWords(v interface{})`

SetForceAllWords sets ForceAllWords field to given value.

### HasForceAllWords

`func (o *Highlight) HasForceAllWords() bool`

HasForceAllWords returns a boolean if a field has been set.

### SetForceAllWordsNil

`func (o *Highlight) SetForceAllWordsNil(b bool)`

 SetForceAllWordsNil sets the value for ForceAllWords to be an explicit nil

### UnsetForceAllWords
`func (o *Highlight) UnsetForceAllWords()`

UnsetForceAllWords ensures that no value is present for ForceAllWords, not even an explicit nil
### GetForceSnippets

`func (o *Highlight) GetForceSnippets() interface{}`

GetForceSnippets returns the ForceSnippets field if non-nil, zero value otherwise.

### GetForceSnippetsOk

`func (o *Highlight) GetForceSnippetsOk() (*interface{}, bool)`

GetForceSnippetsOk returns a tuple with the ForceSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSnippets

`func (o *Highlight) SetForceSnippets(v interface{})`

SetForceSnippets sets ForceSnippets field to given value.

### HasForceSnippets

`func (o *Highlight) HasForceSnippets() bool`

HasForceSnippets returns a boolean if a field has been set.

### SetForceSnippetsNil

`func (o *Highlight) SetForceSnippetsNil(b bool)`

 SetForceSnippetsNil sets the value for ForceSnippets to be an explicit nil

### UnsetForceSnippets
`func (o *Highlight) UnsetForceSnippets()`

UnsetForceSnippets ensures that no value is present for ForceSnippets, not even an explicit nil
### GetHighlightQuery

`func (o *Highlight) GetHighlightQuery() HighlightAllOfHighlightQuery`

GetHighlightQuery returns the HighlightQuery field if non-nil, zero value otherwise.

### GetHighlightQueryOk

`func (o *Highlight) GetHighlightQueryOk() (*HighlightAllOfHighlightQuery, bool)`

GetHighlightQueryOk returns a tuple with the HighlightQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightQuery

`func (o *Highlight) SetHighlightQuery(v HighlightAllOfHighlightQuery)`

SetHighlightQuery sets HighlightQuery field to given value.

### HasHighlightQuery

`func (o *Highlight) HasHighlightQuery() bool`

HasHighlightQuery returns a boolean if a field has been set.

### GetHtmlStripMode

`func (o *Highlight) GetHtmlStripMode() interface{}`

GetHtmlStripMode returns the HtmlStripMode field if non-nil, zero value otherwise.

### GetHtmlStripModeOk

`func (o *Highlight) GetHtmlStripModeOk() (*interface{}, bool)`

GetHtmlStripModeOk returns a tuple with the HtmlStripMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlStripMode

`func (o *Highlight) SetHtmlStripMode(v interface{})`

SetHtmlStripMode sets HtmlStripMode field to given value.

### HasHtmlStripMode

`func (o *Highlight) HasHtmlStripMode() bool`

HasHtmlStripMode returns a boolean if a field has been set.

### SetHtmlStripModeNil

`func (o *Highlight) SetHtmlStripModeNil(b bool)`

 SetHtmlStripModeNil sets the value for HtmlStripMode to be an explicit nil

### UnsetHtmlStripMode
`func (o *Highlight) UnsetHtmlStripMode()`

UnsetHtmlStripMode ensures that no value is present for HtmlStripMode, not even an explicit nil
### GetLimitsPerField

`func (o *Highlight) GetLimitsPerField() interface{}`

GetLimitsPerField returns the LimitsPerField field if non-nil, zero value otherwise.

### GetLimitsPerFieldOk

`func (o *Highlight) GetLimitsPerFieldOk() (*interface{}, bool)`

GetLimitsPerFieldOk returns a tuple with the LimitsPerField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitsPerField

`func (o *Highlight) SetLimitsPerField(v interface{})`

SetLimitsPerField sets LimitsPerField field to given value.

### HasLimitsPerField

`func (o *Highlight) HasLimitsPerField() bool`

HasLimitsPerField returns a boolean if a field has been set.

### SetLimitsPerFieldNil

`func (o *Highlight) SetLimitsPerFieldNil(b bool)`

 SetLimitsPerFieldNil sets the value for LimitsPerField to be an explicit nil

### UnsetLimitsPerField
`func (o *Highlight) UnsetLimitsPerField()`

UnsetLimitsPerField ensures that no value is present for LimitsPerField, not even an explicit nil
### GetNoMatchSize

`func (o *Highlight) GetNoMatchSize() interface{}`

GetNoMatchSize returns the NoMatchSize field if non-nil, zero value otherwise.

### GetNoMatchSizeOk

`func (o *Highlight) GetNoMatchSizeOk() (*interface{}, bool)`

GetNoMatchSizeOk returns a tuple with the NoMatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMatchSize

`func (o *Highlight) SetNoMatchSize(v interface{})`

SetNoMatchSize sets NoMatchSize field to given value.

### HasNoMatchSize

`func (o *Highlight) HasNoMatchSize() bool`

HasNoMatchSize returns a boolean if a field has been set.

### SetNoMatchSizeNil

`func (o *Highlight) SetNoMatchSizeNil(b bool)`

 SetNoMatchSizeNil sets the value for NoMatchSize to be an explicit nil

### UnsetNoMatchSize
`func (o *Highlight) UnsetNoMatchSize()`

UnsetNoMatchSize ensures that no value is present for NoMatchSize, not even an explicit nil
### GetOrder

`func (o *Highlight) GetOrder() interface{}`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Highlight) GetOrderOk() (*interface{}, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Highlight) SetOrder(v interface{})`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Highlight) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *Highlight) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *Highlight) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetPreTags

`func (o *Highlight) GetPreTags() interface{}`

GetPreTags returns the PreTags field if non-nil, zero value otherwise.

### GetPreTagsOk

`func (o *Highlight) GetPreTagsOk() (*interface{}, bool)`

GetPreTagsOk returns a tuple with the PreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreTags

`func (o *Highlight) SetPreTags(v interface{})`

SetPreTags sets PreTags field to given value.

### HasPreTags

`func (o *Highlight) HasPreTags() bool`

HasPreTags returns a boolean if a field has been set.

### SetPreTagsNil

`func (o *Highlight) SetPreTagsNil(b bool)`

 SetPreTagsNil sets the value for PreTags to be an explicit nil

### UnsetPreTags
`func (o *Highlight) UnsetPreTags()`

UnsetPreTags ensures that no value is present for PreTags, not even an explicit nil
### GetPostTags

`func (o *Highlight) GetPostTags() interface{}`

GetPostTags returns the PostTags field if non-nil, zero value otherwise.

### GetPostTagsOk

`func (o *Highlight) GetPostTagsOk() (*interface{}, bool)`

GetPostTagsOk returns a tuple with the PostTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTags

`func (o *Highlight) SetPostTags(v interface{})`

SetPostTags sets PostTags field to given value.

### HasPostTags

`func (o *Highlight) HasPostTags() bool`

HasPostTags returns a boolean if a field has been set.

### SetPostTagsNil

`func (o *Highlight) SetPostTagsNil(b bool)`

 SetPostTagsNil sets the value for PostTags to be an explicit nil

### UnsetPostTags
`func (o *Highlight) UnsetPostTags()`

UnsetPostTags ensures that no value is present for PostTags, not even an explicit nil
### GetStartSnippetId

`func (o *Highlight) GetStartSnippetId() interface{}`

GetStartSnippetId returns the StartSnippetId field if non-nil, zero value otherwise.

### GetStartSnippetIdOk

`func (o *Highlight) GetStartSnippetIdOk() (*interface{}, bool)`

GetStartSnippetIdOk returns a tuple with the StartSnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSnippetId

`func (o *Highlight) SetStartSnippetId(v interface{})`

SetStartSnippetId sets StartSnippetId field to given value.

### HasStartSnippetId

`func (o *Highlight) HasStartSnippetId() bool`

HasStartSnippetId returns a boolean if a field has been set.

### SetStartSnippetIdNil

`func (o *Highlight) SetStartSnippetIdNil(b bool)`

 SetStartSnippetIdNil sets the value for StartSnippetId to be an explicit nil

### UnsetStartSnippetId
`func (o *Highlight) UnsetStartSnippetId()`

UnsetStartSnippetId ensures that no value is present for StartSnippetId, not even an explicit nil
### GetUseBoundaries

`func (o *Highlight) GetUseBoundaries() interface{}`

GetUseBoundaries returns the UseBoundaries field if non-nil, zero value otherwise.

### GetUseBoundariesOk

`func (o *Highlight) GetUseBoundariesOk() (*interface{}, bool)`

GetUseBoundariesOk returns a tuple with the UseBoundaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBoundaries

`func (o *Highlight) SetUseBoundaries(v interface{})`

SetUseBoundaries sets UseBoundaries field to given value.

### HasUseBoundaries

`func (o *Highlight) HasUseBoundaries() bool`

HasUseBoundaries returns a boolean if a field has been set.

### SetUseBoundariesNil

`func (o *Highlight) SetUseBoundariesNil(b bool)`

 SetUseBoundariesNil sets the value for UseBoundaries to be an explicit nil

### UnsetUseBoundaries
`func (o *Highlight) UnsetUseBoundaries()`

UnsetUseBoundaries ensures that no value is present for UseBoundaries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


