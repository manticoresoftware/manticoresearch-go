# QueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Filter object defining a query string | [optional] 
**Match** | Pointer to **map[string]interface{}** | Filter object defining a match keyword passed as a string or in a Match object | [optional] 
**MatchPhrase** | Pointer to **map[string]interface{}** | Filter object defining a match phrase | [optional] 
**MatchAll** | Pointer to **map[string]interface{}** | Filter object to select all documents | [optional] 
**Bool** | Pointer to [**BoolFilter**](BoolFilter.md) |  | [optional] 
**Equals** | Pointer to **interface{}** |  | [optional] 
**In** | Pointer to **map[string]interface{}** | Filter to match a given set of attribute values. | [optional] 
**Range** | Pointer to **map[string]interface{}** | Filter to match a given range of attribute values passed in Range objects | [optional] 
**GeoDistance** | Pointer to [**GeoDistance**](GeoDistance.md) |  | [optional] 

## Methods

### NewQueryFilter

`func NewQueryFilter() *QueryFilter`

NewQueryFilter instantiates a new QueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFilterWithDefaults

`func NewQueryFilterWithDefaults() *QueryFilter`

NewQueryFilterWithDefaults instantiates a new QueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *QueryFilter) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *QueryFilter) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *QueryFilter) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *QueryFilter) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetMatch

`func (o *QueryFilter) GetMatch() map[string]interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *QueryFilter) GetMatchOk() (*map[string]interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *QueryFilter) SetMatch(v map[string]interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *QueryFilter) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchPhrase

`func (o *QueryFilter) GetMatchPhrase() map[string]interface{}`

GetMatchPhrase returns the MatchPhrase field if non-nil, zero value otherwise.

### GetMatchPhraseOk

`func (o *QueryFilter) GetMatchPhraseOk() (*map[string]interface{}, bool)`

GetMatchPhraseOk returns a tuple with the MatchPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPhrase

`func (o *QueryFilter) SetMatchPhrase(v map[string]interface{})`

SetMatchPhrase sets MatchPhrase field to given value.

### HasMatchPhrase

`func (o *QueryFilter) HasMatchPhrase() bool`

HasMatchPhrase returns a boolean if a field has been set.

### GetMatchAll

`func (o *QueryFilter) GetMatchAll() map[string]interface{}`

GetMatchAll returns the MatchAll field if non-nil, zero value otherwise.

### GetMatchAllOk

`func (o *QueryFilter) GetMatchAllOk() (*map[string]interface{}, bool)`

GetMatchAllOk returns a tuple with the MatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAll

`func (o *QueryFilter) SetMatchAll(v map[string]interface{})`

SetMatchAll sets MatchAll field to given value.

### HasMatchAll

`func (o *QueryFilter) HasMatchAll() bool`

HasMatchAll returns a boolean if a field has been set.

### GetBool

`func (o *QueryFilter) GetBool() BoolFilter`

GetBool returns the Bool field if non-nil, zero value otherwise.

### GetBoolOk

`func (o *QueryFilter) GetBoolOk() (*BoolFilter, bool)`

GetBoolOk returns a tuple with the Bool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBool

`func (o *QueryFilter) SetBool(v BoolFilter)`

SetBool sets Bool field to given value.

### HasBool

`func (o *QueryFilter) HasBool() bool`

HasBool returns a boolean if a field has been set.

### GetEquals

`func (o *QueryFilter) GetEquals() interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *QueryFilter) GetEqualsOk() (*interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *QueryFilter) SetEquals(v interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *QueryFilter) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### SetEqualsNil

`func (o *QueryFilter) SetEqualsNil(b bool)`

 SetEqualsNil sets the value for Equals to be an explicit nil

### UnsetEquals
`func (o *QueryFilter) UnsetEquals()`

UnsetEquals ensures that no value is present for Equals, not even an explicit nil
### GetIn

`func (o *QueryFilter) GetIn() map[string]interface{}`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *QueryFilter) GetInOk() (*map[string]interface{}, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *QueryFilter) SetIn(v map[string]interface{})`

SetIn sets In field to given value.

### HasIn

`func (o *QueryFilter) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetRange

`func (o *QueryFilter) GetRange() map[string]interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *QueryFilter) GetRangeOk() (*map[string]interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *QueryFilter) SetRange(v map[string]interface{})`

SetRange sets Range field to given value.

### HasRange

`func (o *QueryFilter) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetGeoDistance

`func (o *QueryFilter) GetGeoDistance() GeoDistance`

GetGeoDistance returns the GeoDistance field if non-nil, zero value otherwise.

### GetGeoDistanceOk

`func (o *QueryFilter) GetGeoDistanceOk() (*GeoDistance, bool)`

GetGeoDistanceOk returns a tuple with the GeoDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDistance

`func (o *QueryFilter) SetGeoDistance(v GeoDistance)`

SetGeoDistance sets GeoDistance field to given value.

### HasGeoDistance

`func (o *QueryFilter) HasGeoDistance() bool`

HasGeoDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


