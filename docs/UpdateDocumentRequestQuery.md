# UpdateDocumentRequestQuery

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

### NewUpdateDocumentRequestQuery

`func NewUpdateDocumentRequestQuery() *UpdateDocumentRequestQuery`

NewUpdateDocumentRequestQuery instantiates a new UpdateDocumentRequestQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentRequestQueryWithDefaults

`func NewUpdateDocumentRequestQueryWithDefaults() *UpdateDocumentRequestQuery`

NewUpdateDocumentRequestQueryWithDefaults instantiates a new UpdateDocumentRequestQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *UpdateDocumentRequestQuery) GetQueryString() interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *UpdateDocumentRequestQuery) GetQueryStringOk() (*interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *UpdateDocumentRequestQuery) SetQueryString(v interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *UpdateDocumentRequestQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *UpdateDocumentRequestQuery) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *UpdateDocumentRequestQuery) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetMatch

`func (o *UpdateDocumentRequestQuery) GetMatch() interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *UpdateDocumentRequestQuery) GetMatchOk() (*interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *UpdateDocumentRequestQuery) SetMatch(v interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *UpdateDocumentRequestQuery) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *UpdateDocumentRequestQuery) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *UpdateDocumentRequestQuery) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetMatchPhrase

`func (o *UpdateDocumentRequestQuery) GetMatchPhrase() interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *UpdateDocumentRequestQuery) GetMatchPhraseOk() (*interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *UpdateDocumentRequestQuery) SetMatchPhrase(v interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *UpdateDocumentRequestQuery) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### SetMatchPhraseNil

`func (o *UpdateDocumentRequestQuery) SetMatchPhraseNil(b bool)`

 SetMatchPhraseNil sets the value for MatchPhrase to be an explicit nil

### UnsetMatchPhrase
`func (o *UpdateDocumentRequestQuery) UnsetMatchPhrase()`

UnsetMatchPhrase ensures that no value is present for MatchPhrase, not even an explicit nil
### GetMatchAll

`func (o *UpdateDocumentRequestQuery) GetMatchAll() interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *UpdateDocumentRequestQuery) GetMatchAllOk() (*interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *UpdateDocumentRequestQuery) SetMatchAll(v interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *UpdateDocumentRequestQuery) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### SetMatchAllNil

`func (o *UpdateDocumentRequestQuery) SetMatchAllNil(b bool)`

 SetMatchAllNil sets the value for MatchAll to be an explicit nil

### UnsetMatchAll
`func (o *UpdateDocumentRequestQuery) UnsetMatchAll()`

UnsetMatchAll ensures that no value is present for MatchAll, not even an explicit nil
### GetBool

`func (o *UpdateDocumentRequestQuery) GetBool() BoolFilter`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *UpdateDocumentRequestQuery) GetBoolOk() (*BoolFilter, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *UpdateDocumentRequestQuery) SetBool(v BoolFilter)`

SetBool sets Bool field to given value.

### HasBool

`func (o *UpdateDocumentRequestQuery) HasBool() bool`

HasBool returns a boolean if a field has been set.

### GetEquals

`func (o *UpdateDocumentRequestQuery) GetEquals() interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *UpdateDocumentRequestQuery) GetEqualsOk() (*interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *UpdateDocumentRequestQuery) SetEquals(v interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *UpdateDocumentRequestQuery) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### SetEqualsNil

`func (o *UpdateDocumentRequestQuery) SetEqualsNil(b bool)`

 SetEqualsNil sets the value for Equals to be an explicit nil

### UnsetEquals
`func (o *UpdateDocumentRequestQuery) UnsetEquals()`

UnsetEquals ensures that no value is present for Equals, not even an explicit nil
### GetIn

`func (o *UpdateDocumentRequestQuery) GetIn() interface{}`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *UpdateDocumentRequestQuery) GetInOk() (*interface{}, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *UpdateDocumentRequestQuery) SetIn(v interface{})`

SetIn sets In field to given value.

### HasIn

`func (o *UpdateDocumentRequestQuery) HasIn() bool`

HasIn returns a boolean if a field has been set.

### SetInNil

`func (o *UpdateDocumentRequestQuery) SetInNil(b bool)`

 SetInNil sets the value for In to be an explicit nil

### UnsetIn
`func (o *UpdateDocumentRequestQuery) UnsetIn()`

UnsetIn ensures that no value is present for In, not even an explicit nil
### GetRange

`func (o *UpdateDocumentRequestQuery) GetRange() interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *UpdateDocumentRequestQuery) GetRangeOk() (*interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *UpdateDocumentRequestQuery) SetRange(v interface{})`

SetRange sets Range field to given value.

### HasRange

`func (o *UpdateDocumentRequestQuery) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRangeNil

`func (o *UpdateDocumentRequestQuery) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *UpdateDocumentRequestQuery) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil
### GetGeoDistance

`func (o *UpdateDocumentRequestQuery) GetGeoDistance() GeoDistance`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *UpdateDocumentRequestQuery) GetGeoDistanceOk() (*GeoDistance, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *UpdateDocumentRequestQuery) SetGeoDistance(v GeoDistance)`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *UpdateDocumentRequestQuery) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


