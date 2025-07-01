# FulltextFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **interface{}** | Filter object defining a query string | [optional] 
**Match** | Pointer to **interface{}** | Filter object defining a match keyword passed as a string or in a Match object | [optional] 
**MatchPhrase** | Pointer to **interface{}** | Filter object defining a match phrase | [optional] 
**MatchAll** | Pointer to **interface{}** | Filter object to select all documents | [optional] 

## Methods

### NewFulltextFilter

`func NewFulltextFilter() *FulltextFilter`

NewFulltextFilter instantiates a new FulltextFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulltextFilterWithDefaults

`func NewFulltextFilterWithDefaults() *FulltextFilter`

NewFulltextFilterWithDefaults instantiates a new FulltextFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *FulltextFilter) GetQueryString() interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *FulltextFilter) GetQueryStringOk() (*interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *FulltextFilter) SetQueryString(v interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *FulltextFilter) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *FulltextFilter) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *FulltextFilter) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetMatch

`func (o *FulltextFilter) GetMatch() interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *FulltextFilter) GetMatchOk() (*interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *FulltextFilter) SetMatch(v interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *FulltextFilter) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *FulltextFilter) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *FulltextFilter) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetMatchPhrase

`func (o *FulltextFilter) GetMatchPhrase() interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *FulltextFilter) GetMatchPhraseOk() (*interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *FulltextFilter) SetMatchPhrase(v interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *FulltextFilter) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### SetMatchPhraseNil

`func (o *FulltextFilter) SetMatchPhraseNil(b bool)`

 SetMatchPhraseNil sets the value for MatchPhrase to be an explicit nil

### UnsetMatchPhrase
`func (o *FulltextFilter) UnsetMatchPhrase()`

UnsetMatchPhrase ensures that no value is present for MatchPhrase, not even an explicit nil
### GetMatchAll

`func (o *FulltextFilter) GetMatchAll() interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *FulltextFilter) GetMatchAllOk() (*interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *FulltextFilter) SetMatchAll(v interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *FulltextFilter) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### SetMatchAllNil

`func (o *FulltextFilter) SetMatchAllNil(b bool)`

 SetMatchAllNil sets the value for MatchAll to be an explicit nil

### UnsetMatchAll
`func (o *FulltextFilter) UnsetMatchAll()`

UnsetMatchAll ensures that no value is present for MatchAll, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


