# SearchRequest

Request object for search operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | [default to ""]
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**FulltextFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**AttrFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**MaxMatches** | Pointer to **int32** |  | [optional] 
**Sort** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Aggs** | Pointer to [**map[string]Aggregation**](Aggregation.md) |  | [optional] 
**Expressions** | Pointer to **map[string]string** |  | [optional] 
**Highlight** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Profile** | Pointer to **bool** |  | [optional] 
**TrackScores** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(index string, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SearchRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SearchRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SearchRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetQuery

`func (o *SearchRequest) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFulltextFilter

`func (o *SearchRequest) GetFulltextFilter() map[string]interface{}`

GetFulltextFilter returns the FulltextFilter field if non-nil, zero value otherwise.

### GetFulltextFilterOk

`func (o *SearchRequest) GetFulltextFilterOk() (*map[string]interface{}, bool)`

GetFulltextFilterOk returns a tuple with the FulltextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextFilter

`func (o *SearchRequest) SetFulltextFilter(v map[string]interface{})`

SetFulltextFilter sets FulltextFilter field to given value.

### HasFulltextFilter

`func (o *SearchRequest) HasFulltextFilter() bool`

HasFulltextFilter returns a boolean if a field has been set.

### GetAttrFilter

`func (o *SearchRequest) GetAttrFilter() map[string]interface{}`

GetAttrFilter returns the AttrFilter field if non-nil, zero value otherwise.

### GetAttrFilterOk

`func (o *SearchRequest) GetAttrFilterOk() (*map[string]interface{}, bool)`

GetAttrFilterOk returns a tuple with the AttrFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrFilter

`func (o *SearchRequest) SetAttrFilter(v map[string]interface{})`

SetAttrFilter sets AttrFilter field to given value.

### HasAttrFilter

`func (o *SearchRequest) HasAttrFilter() bool`

HasAttrFilter returns a boolean if a field has been set.

### GetLimit

`func (o *SearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMaxMatches

`func (o *SearchRequest) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *SearchRequest) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *SearchRequest) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *SearchRequest) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetSort

`func (o *SearchRequest) GetSort() []map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchRequest) GetSortOk() (*[]map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchRequest) SetSort(v []map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetAggs

`func (o *SearchRequest) GetAggs() map[string]Aggregation`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *SearchRequest) GetAggsOk() (*map[string]Aggregation, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *SearchRequest) SetAggs(v map[string]Aggregation)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *SearchRequest) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExpressions

`func (o *SearchRequest) GetExpressions() map[string]string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *SearchRequest) GetExpressionsOk() (*map[string]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *SearchRequest) SetExpressions(v map[string]string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *SearchRequest) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetHighlight

`func (o *SearchRequest) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *SearchRequest) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *SearchRequest) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *SearchRequest) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetSource

`func (o *SearchRequest) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchRequest) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchRequest) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOptions

`func (o *SearchRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SearchRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SearchRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SearchRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProfile

`func (o *SearchRequest) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchRequest) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchRequest) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTrackScores

`func (o *SearchRequest) GetTrackScores() bool`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *SearchRequest) GetTrackScoresOk() (*bool, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *SearchRequest) SetTrackScores(v bool)`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *SearchRequest) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


