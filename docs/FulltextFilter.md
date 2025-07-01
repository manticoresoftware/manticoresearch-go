# FulltextFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Filter object defining a query string | [optional] 
**Match** | Pointer to **map[string]interface{}** | Filter object defining a match keyword passed as a string or in a Match object | [optional] 
**MatchPhrase** | Pointer to **map[string]interface{}** | Filter object defining a match phrase | [optional] 
**MatchAll** | Pointer to **map[string]interface{}** | Filter object to select all documents | [optional] 

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

`func (o *FulltextFilter) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *FulltextFilter) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *FulltextFilter) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *FulltextFilter) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetMatch

`func (o *FulltextFilter) GetMatch() map[string]interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *FulltextFilter) GetMatchOk() (*map[string]interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *FulltextFilter) SetMatch(v map[string]interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *FulltextFilter) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchPhrase

`func (o *FulltextFilter) GetMatchPhrase() map[string]interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *FulltextFilter) GetMatchPhraseOk() (*map[string]interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *FulltextFilter) SetMatchPhrase(v map[string]interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *FulltextFilter) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### GetMatchAll

`func (o *FulltextFilter) GetMatchAll() map[string]interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *FulltextFilter) GetMatchAllOk() (*map[string]interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *FulltextFilter) SetMatchAll(v map[string]interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *FulltextFilter) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


