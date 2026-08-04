# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **string** | The table to perform the search on | [optional] 
**Chat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**Query** | Pointer to [**SearchQuery**](SearchQuery.md) |  | [optional] 
**Join** | Pointer to [**[]Join**](Join.md) | Join clause to combine search data from multiple tables | [optional] 
**Highlight** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 
**Limit** | Pointer to **uint64** | Maximum number of results to return | [optional] 
**Knn** | Pointer to [**Knn**](Knn.md) | K-nearest neighbor search settings. Pass a single &#x60;knn&#x60; object or an array of objects for multi-vector search.  | [optional] 
**Hybrid** | Pointer to [**Hybrid**](Hybrid.md) |  | [optional] 
**FacetFilterMode** | Pointer to [**FacetFilterMode**](FacetFilterMode.md) |  | [optional] 
**Aggs** | Pointer to [**map[string]Aggregation**](Aggregation.md) | Defines aggregation settings for grouping results | [optional] 
**Expressions** | Pointer to **map[string]string** | Expressions to calculate additional values for the result. Simpler alternative to &#x60;script_fields&#x60;; expression names must be lowercase.  | [optional] 
**ScriptFields** | Pointer to [**map[string]ScriptField**](ScriptField.md) | Named expressions computed at search time. Each value defines an inline script whose result is stored under the field name. For more information see [Expressions](https://manual.manticoresearch.com/Searching/Expressions#script_fields)  | [optional] 
**MaxMatches** | Pointer to **uint64** | Maximum number of matches allowed in the result | [optional] 
**Offset** | Pointer to **uint64** | Starting point for pagination of the result | [optional] 
**Options** | Pointer to **map[string]interface{}** | Additional search options | [optional] 
**Profile** | Pointer to **bool** | Enable or disable profiling of the search request | [optional] 
**Sort** | Pointer to **interface{}** |  | [optional] 
**Source** | Pointer to **interface{}** |  | [optional] 
**TrackScores** | Pointer to **bool** | Enable or disable result weight calculation used for sorting | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest() *SearchRequest`

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

`func (o *SearchRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *SearchRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *SearchRequest) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *SearchRequest) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetChat

`func (o *SearchRequest) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *SearchRequest) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *SearchRequest) SetChat(v Chat)`

SetChat sets Chat field to given value.

### HasChat

`func (o *SearchRequest) HasChat() bool`

HasChat returns a boolean if a field has been set.

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

`func (o *SearchRequest) GetJoin() []Join`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *SearchRequest) GetJoinOk() (*[]Join, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *SearchRequest) SetJoin(v []Join)`

SetJoin sets Join field to given value.

### HasJoin

`func (o *SearchRequest) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

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

### GetKnn

`func (o *SearchRequest) GetKnn() Knn`

GetKnn returns the Knn field if non-nil, zero value otherwise.

### GetKnnOk

`func (o *SearchRequest) GetKnnOk() (*Knn, bool)`

GetKnnOk returns a tuple with the Knn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnn

`func (o *SearchRequest) SetKnn(v Knn)`

SetKnn sets Knn field to given value.

### HasKnn

`func (o *SearchRequest) HasKnn() bool`

HasKnn returns a boolean if a field has been set.

### GetHybrid

`func (o *SearchRequest) GetHybrid() Hybrid`

GetHybrid returns the Hybrid field if non-nil, zero value otherwise.

### GetHybridOk

`func (o *SearchRequest) GetHybridOk() (*Hybrid, bool)`

GetHybridOk returns a tuple with the Hybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrid

`func (o *SearchRequest) SetHybrid(v Hybrid)`

SetHybrid sets Hybrid field to given value.

### HasHybrid

`func (o *SearchRequest) HasHybrid() bool`

HasHybrid returns a boolean if a field has been set.

### GetFacetFilterMode

`func (o *SearchRequest) GetFacetFilterMode() FacetFilterMode`

GetFacetFilterMode returns the FacetFilterMode field if non-nil, zero value otherwise.

### GetFacetFilterModeOk

`func (o *SearchRequest) GetFacetFilterModeOk() (*FacetFilterMode, bool)`

GetFacetFilterModeOk returns a tuple with the FacetFilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetFilterMode

`func (o *SearchRequest) SetFacetFilterMode(v FacetFilterMode)`

SetFacetFilterMode sets FacetFilterMode field to given value.

### HasFacetFilterMode

`func (o *SearchRequest) HasFacetFilterMode() bool`

HasFacetFilterMode returns a boolean if a field has been set.

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

### GetScriptFields

`func (o *SearchRequest) GetScriptFields() map[string]ScriptField`

GetScriptFields returns the ScriptFields field if non-nil, zero value otherwise.

### GetScriptFieldsOk

`func (o *SearchRequest) GetScriptFieldsOk() (*map[string]ScriptField, bool)`

GetScriptFieldsOk returns a tuple with the ScriptFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptFields

`func (o *SearchRequest) SetScriptFields(v map[string]ScriptField)`

SetScriptFields sets ScriptFields field to given value.

### HasScriptFields

`func (o *SearchRequest) HasScriptFields() bool`

HasScriptFields returns a boolean if a field has been set.

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


