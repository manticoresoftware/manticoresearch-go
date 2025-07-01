# HighlightAllOfHighlightQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **interface{}** | Filter object defining a query string | [optional] 
**Match** | Pointer to **interface{}** | Filter object defining a match keyword passed as a string or in a Match object | [optional] 
**MatchPhrase** | Pointer to **interface{}** | Filter object defining a match phrase | [optional] 
**MatchAll** | Pointer to **interface{}** | Filter object to select all documents | [optional] 
**Bool** | Pointer to [**BoolFilter**](BoolFilter.md) |  | [optional] 
**Equals** | Pointer to **interface{}** | Filter to match exact attribute values. | [optional] 
**In** | Pointer to **interface{}** | Filter to match a given set of attribute values. | [optional] 
**Range** | Pointer to **interface{}** | Filter to match a given range of attribute values passed in Range objects | [optional] 
**GeoDistance** | Pointer to [**GeoDistance**](GeoDistance.md) |  | [optional] 

## Methods

### NewHighlightAllOfHighlightQuery

`func NewHighlightAllOfHighlightQuery() *HighlightAllOfHighlightQuery`

NewHighlightAllOfHighlightQuery instantiates a new HighlightAllOfHighlightQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightAllOfHighlightQueryWithDefaults

`func NewHighlightAllOfHighlightQueryWithDefaults() *HighlightAllOfHighlightQuery`

NewHighlightAllOfHighlightQueryWithDefaults instantiates a new HighlightAllOfHighlightQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *HighlightAllOfHighlightQuery) GetQueryString() interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *HighlightAllOfHighlightQuery) GetQueryStringOk() (*interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *HighlightAllOfHighlightQuery) SetQueryString(v interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *HighlightAllOfHighlightQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *HighlightAllOfHighlightQuery) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *HighlightAllOfHighlightQuery) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetMatch

`func (o *HighlightAllOfHighlightQuery) GetMatch() interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *HighlightAllOfHighlightQuery) GetMatchOk() (*interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *HighlightAllOfHighlightQuery) SetMatch(v interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *HighlightAllOfHighlightQuery) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *HighlightAllOfHighlightQuery) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *HighlightAllOfHighlightQuery) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetMatchPhrase

`func (o *HighlightAllOfHighlightQuery) GetMatchPhrase() interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *HighlightAllOfHighlightQuery) GetMatchPhraseOk() (*interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *HighlightAllOfHighlightQuery) SetMatchPhrase(v interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *HighlightAllOfHighlightQuery) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### SetMatchPhraseNil

`func (o *HighlightAllOfHighlightQuery) SetMatchPhraseNil(b bool)`

 SetMatchPhraseNil sets the value for MatchPhrase to be an explicit nil

### UnsetMatchPhrase
`func (o *HighlightAllOfHighlightQuery) UnsetMatchPhrase()`

UnsetMatchPhrase ensures that no value is present for MatchPhrase, not even an explicit nil
### GetMatchAll

`func (o *HighlightAllOfHighlightQuery) GetMatchAll() interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *HighlightAllOfHighlightQuery) GetMatchAllOk() (*interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *HighlightAllOfHighlightQuery) SetMatchAll(v interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *HighlightAllOfHighlightQuery) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### SetMatchAllNil

`func (o *HighlightAllOfHighlightQuery) SetMatchAllNil(b bool)`

 SetMatchAllNil sets the value for MatchAll to be an explicit nil

### UnsetMatchAll
`func (o *HighlightAllOfHighlightQuery) UnsetMatchAll()`

UnsetMatchAll ensures that no value is present for MatchAll, not even an explicit nil
### GetBool

`func (o *HighlightAllOfHighlightQuery) GetBool() BoolFilter`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *HighlightAllOfHighlightQuery) GetBoolOk() (*BoolFilter, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *HighlightAllOfHighlightQuery) SetBool(v BoolFilter)`

SetBool sets Bool field to given value.

### HasBool

`func (o *HighlightAllOfHighlightQuery) HasBool() bool`

HasBool returns a boolean if a field has been set.

### GetEquals

`func (o *HighlightAllOfHighlightQuery) GetEquals() interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *HighlightAllOfHighlightQuery) GetEqualsOk() (*interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *HighlightAllOfHighlightQuery) SetEquals(v interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *HighlightAllOfHighlightQuery) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### SetEqualsNil

`func (o *HighlightAllOfHighlightQuery) SetEqualsNil(b bool)`

 SetEqualsNil sets the value for Equals to be an explicit nil

### UnsetEquals
`func (o *HighlightAllOfHighlightQuery) UnsetEquals()`

UnsetEquals ensures that no value is present for Equals, not even an explicit nil
### GetIn

`func (o *HighlightAllOfHighlightQuery) GetIn() interface{}`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *HighlightAllOfHighlightQuery) GetInOk() (*interface{}, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *HighlightAllOfHighlightQuery) SetIn(v interface{})`

SetIn sets In field to given value.

### HasIn

`func (o *HighlightAllOfHighlightQuery) HasIn() bool`

HasIn returns a boolean if a field has been set.

### SetInNil

`func (o *HighlightAllOfHighlightQuery) SetInNil(b bool)`

 SetInNil sets the value for In to be an explicit nil

### UnsetIn
`func (o *HighlightAllOfHighlightQuery) UnsetIn()`

UnsetIn ensures that no value is present for In, not even an explicit nil
### GetRange

`func (o *HighlightAllOfHighlightQuery) GetRange() interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *HighlightAllOfHighlightQuery) GetRangeOk() (*interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *HighlightAllOfHighlightQuery) SetRange(v interface{})`

SetRange sets Range field to given value.

### HasRange

`func (o *HighlightAllOfHighlightQuery) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRangeNil

`func (o *HighlightAllOfHighlightQuery) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *HighlightAllOfHighlightQuery) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil
### GetGeoDistance

`func (o *HighlightAllOfHighlightQuery) GetGeoDistance() GeoDistance`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *HighlightAllOfHighlightQuery) GetGeoDistanceOk() (*GeoDistance, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *HighlightAllOfHighlightQuery) SetGeoDistance(v GeoDistance)`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *HighlightAllOfHighlightQuery) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


