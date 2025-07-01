# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **interface{}** | The table to perform the search on | 
**Query** | Pointer to [**SearchQuery**](SearchQuery.md) |  | [optional] 
**Join** | Pointer to **interface{}** | Join clause to combine search data from multiple tables | [optional] 
**Highlight** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 
**Limit** | Pointer to **interface{}** | Maximum number of results to return | [optional] 
**Knn** | Pointer to [**KnnQuery**](KnnQuery.md) |  | [optional] 
**Aggs** | Pointer to  | Defines aggregation settings for grouping results | [optional] 
**Expressions** | Pointer to  | Expressions to calculate additional values for the result | [optional] 
**MaxMatches** | Pointer to **interface{}** | Maximum number of matches allowed in the result | [optional] 
**Offset** | Pointer to **interface{}** | Starting point for pagination of the result | [optional] 
**Options** | Pointer to **interface{}** | Additional search options | [optional] 
**Profile** | Pointer to **interface{}** | Enable or disable profiling of the search request | [optional] 
**Sort** | Pointer to **interface{}** | Sorting criteria for the search results | [optional] 
**Source** | Pointer to **interface{}** | Specify which fields to include or exclude in the response | [optional] 
**TrackScores** | Pointer to **interface{}** | Enable or disable result weight calculation used for sorting | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(table interface{}, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *SearchRequest) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *SearchRequest) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *SearchRequest) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *SearchRequest) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *SearchRequest) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetQuery

`func (o *SearchRequest) GetQuery() SearchQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*SearchQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v SearchQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetJoin

`func (o *SearchRequest) GetJoin() interface{}`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *SearchRequest) GetJoinOk() (*interface{}, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *SearchRequest) SetJoin(v interface{})`

SetJoin sets Join field to given value.

### HasJoin

`func (o *SearchRequest) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

### SetJoinNil

`func (o *SearchRequest) SetJoinNil(b bool)`

 SetJoinNil sets the value for Join to be an explicit nil

### UnsetJoin
`func (o *SearchRequest) UnsetJoin()`

UnsetJoin ensures that no value is present for Join, not even an explicit nil
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

### GetLimit

`func (o *SearchRequest) GetLimit() interface{}`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchRequest) GetLimitOk() (*interface{}, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchRequest) SetLimit(v interface{})`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *SearchRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *SearchRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetKnn

`func (o *SearchRequest) GetKnn() KnnQuery`

GetKnn returns the Knn field if non-nil, zero value otherwise.

### GetKnnOk

`func (o *SearchRequest) GetKnnOk() (*KnnQuery, bool)`

GetKnnOk returns a tuple with the Knn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnn

`func (o *SearchRequest) SetKnn(v KnnQuery)`

SetKnn sets Knn field to given value.

### HasKnn

`func (o *SearchRequest) HasKnn() bool`

HasKnn returns a boolean if a field has been set.

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

### SetAggsNil

`func (o *SearchRequest) SetAggsNil(b bool)`

 SetAggsNil sets the value for Aggs to be an explicit nil

### UnsetAggs
`func (o *SearchRequest) UnsetAggs()`

UnsetAggs ensures that no value is present for Aggs, not even an explicit nil
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

### SetExpressionsNil

`func (o *SearchRequest) SetExpressionsNil(b bool)`

 SetExpressionsNil sets the value for Expressions to be an explicit nil

### UnsetExpressions
`func (o *SearchRequest) UnsetExpressions()`

UnsetExpressions ensures that no value is present for Expressions, not even an explicit nil
### GetMaxMatches

`func (o *SearchRequest) GetMaxMatches() interface{}`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *SearchRequest) GetMaxMatchesOk() (*interface{}, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *SearchRequest) SetMaxMatches(v interface{})`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *SearchRequest) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### SetMaxMatchesNil

`func (o *SearchRequest) SetMaxMatchesNil(b bool)`

 SetMaxMatchesNil sets the value for MaxMatches to be an explicit nil

### UnsetMaxMatches
`func (o *SearchRequest) UnsetMaxMatches()`

UnsetMaxMatches ensures that no value is present for MaxMatches, not even an explicit nil
### GetOffset

`func (o *SearchRequest) GetOffset() interface{}`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchRequest) GetOffsetOk() (*interface{}, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchRequest) SetOffset(v interface{})`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *SearchRequest) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *SearchRequest) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOptions

`func (o *SearchRequest) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SearchRequest) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SearchRequest) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SearchRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SearchRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SearchRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProfile

`func (o *SearchRequest) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchRequest) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchRequest) SetProfile(v interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *SearchRequest) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *SearchRequest) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetSort

`func (o *SearchRequest) GetSort() interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchRequest) GetSortOk() (*interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchRequest) SetSort(v interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *SearchRequest) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *SearchRequest) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetSource

`func (o *SearchRequest) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchRequest) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchRequest) SetSource(v interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *SearchRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SearchRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetTrackScores

`func (o *SearchRequest) GetTrackScores() interface{}`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *SearchRequest) GetTrackScoresOk() (*interface{}, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *SearchRequest) SetTrackScores(v interface{})`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *SearchRequest) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.

### SetTrackScoresNil

`func (o *SearchRequest) SetTrackScoresNil(b bool)`

 SetTrackScoresNil sets the value for TrackScores to be an explicit nil

### UnsetTrackScores
`func (o *SearchRequest) UnsetTrackScores()`

UnsetTrackScores ensures that no value is present for TrackScores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


