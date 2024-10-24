# BasicSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggs** | Pointer to [**Aggregation**](Aggregation.md) |  | [optional] 
**Expressions** | Pointer to **map[string]string** |  | [optional] 
**Join** | Pointer to [**[]JoinInner**](JoinInner.md) |  | [optional] 
**Highlight** | Pointer to [**Highlight**](Highlight.md) |  | [optional] 
**Index** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**MaxMatches** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Profile** | Pointer to **bool** |  | [optional] 
**Sort** | Pointer to [**[]SearchRequestParametersSortInner**](SearchRequestParametersSortInner.md) |  | [optional] 
**Source** | Pointer to [**SearchRequestParametersSource**](SearchRequestParametersSource.md) |  | [optional] 
**TrackScores** | Pointer to **bool** |  | [optional] 
**Query** | [**QueryFilter**](QueryFilter.md) |  | 

## Methods

### NewBasicSearchRequest

`func NewBasicSearchRequest(index string, query QueryFilter, ) *BasicSearchRequest`

NewBasicSearchRequest instantiates a new BasicSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicSearchRequestWithDefaults

`func NewBasicSearchRequestWithDefaults() *BasicSearchRequest`

NewBasicSearchRequestWithDefaults instantiates a new BasicSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggs

`func (o *BasicSearchRequest) GetAggs() Aggregation`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *BasicSearchRequest) GetAggsOk() (*Aggregation, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *BasicSearchRequest) SetAggs(v Aggregation)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *BasicSearchRequest) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExpressions

`func (o *BasicSearchRequest) GetExpressions() map[string]string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *BasicSearchRequest) GetExpressionsOk() (*map[string]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *BasicSearchRequest) SetExpressions(v map[string]string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *BasicSearchRequest) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetJoin

`func (o *BasicSearchRequest) GetJoin() []JoinInner`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *BasicSearchRequest) GetJoinOk() (*[]JoinInner, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *BasicSearchRequest) SetJoin(v []JoinInner)`

SetJoin sets Join field to given value.

### HasJoin

`func (o *BasicSearchRequest) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

### GetHighlight

`func (o *BasicSearchRequest) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *BasicSearchRequest) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *BasicSearchRequest) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *BasicSearchRequest) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetIndex

`func (o *BasicSearchRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BasicSearchRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BasicSearchRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetLimit

`func (o *BasicSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BasicSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BasicSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BasicSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxMatches

`func (o *BasicSearchRequest) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *BasicSearchRequest) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *BasicSearchRequest) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *BasicSearchRequest) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetOffset

`func (o *BasicSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BasicSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BasicSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BasicSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptions

`func (o *BasicSearchRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BasicSearchRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BasicSearchRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BasicSearchRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProfile

`func (o *BasicSearchRequest) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BasicSearchRequest) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BasicSearchRequest) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BasicSearchRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSort

`func (o *BasicSearchRequest) GetSort() []SearchRequestParametersSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BasicSearchRequest) GetSortOk() (*[]SearchRequestParametersSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BasicSearchRequest) SetSort(v []SearchRequestParametersSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BasicSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSource

`func (o *BasicSearchRequest) GetSource() SearchRequestParametersSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BasicSearchRequest) GetSourceOk() (*SearchRequestParametersSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BasicSearchRequest) SetSource(v SearchRequestParametersSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *BasicSearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTrackScores

`func (o *BasicSearchRequest) GetTrackScores() bool`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *BasicSearchRequest) GetTrackScoresOk() (*bool, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *BasicSearchRequest) SetTrackScores(v bool)`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *BasicSearchRequest) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.

### GetQuery

`func (o *BasicSearchRequest) GetQuery() QueryFilter`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BasicSearchRequest) GetQueryOk() (*QueryFilter, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BasicSearchRequest) SetQuery(v QueryFilter)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


