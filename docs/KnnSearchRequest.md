# KnnSearchRequest

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
**Knn** | [**KnnSearchRequestAllOfKnn**](KnnSearchRequestAllOfKnn.md) |  | 

## Methods

### NewKnnSearchRequest

`func NewKnnSearchRequest(index string, knn KnnSearchRequestAllOfKnn, ) *KnnSearchRequest`

NewKnnSearchRequest instantiates a new KnnSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnSearchRequestWithDefaults

`func NewKnnSearchRequestWithDefaults() *KnnSearchRequest`

NewKnnSearchRequestWithDefaults instantiates a new KnnSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggs

`func (o *KnnSearchRequest) GetAggs() Aggregation`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *KnnSearchRequest) GetAggsOk() (*Aggregation, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *KnnSearchRequest) SetAggs(v Aggregation)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *KnnSearchRequest) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExpressions

`func (o *KnnSearchRequest) GetExpressions() map[string]string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *KnnSearchRequest) GetExpressionsOk() (*map[string]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *KnnSearchRequest) SetExpressions(v map[string]string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *KnnSearchRequest) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetJoin

`func (o *KnnSearchRequest) GetJoin() []JoinInner`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *KnnSearchRequest) GetJoinOk() (*[]JoinInner, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *KnnSearchRequest) SetJoin(v []JoinInner)`

SetJoin sets Join field to given value.

### HasJoin

`func (o *KnnSearchRequest) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

### GetHighlight

`func (o *KnnSearchRequest) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *KnnSearchRequest) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *KnnSearchRequest) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *KnnSearchRequest) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetIndex

`func (o *KnnSearchRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *KnnSearchRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *KnnSearchRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetLimit

`func (o *KnnSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KnnSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KnnSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KnnSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxMatches

`func (o *KnnSearchRequest) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *KnnSearchRequest) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *KnnSearchRequest) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *KnnSearchRequest) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetOffset

`func (o *KnnSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KnnSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KnnSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KnnSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptions

`func (o *KnnSearchRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KnnSearchRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KnnSearchRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KnnSearchRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProfile

`func (o *KnnSearchRequest) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *KnnSearchRequest) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *KnnSearchRequest) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *KnnSearchRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSort

`func (o *KnnSearchRequest) GetSort() []SearchRequestParametersSortInner`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *KnnSearchRequest) GetSortOk() (*[]SearchRequestParametersSortInner, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *KnnSearchRequest) SetSort(v []SearchRequestParametersSortInner)`

SetSort sets Sort field to given value.

### HasSort

`func (o *KnnSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSource

`func (o *KnnSearchRequest) GetSource() SearchRequestParametersSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KnnSearchRequest) GetSourceOk() (*SearchRequestParametersSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KnnSearchRequest) SetSource(v SearchRequestParametersSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *KnnSearchRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTrackScores

`func (o *KnnSearchRequest) GetTrackScores() bool`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *KnnSearchRequest) GetTrackScoresOk() (*bool, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *KnnSearchRequest) SetTrackScores(v bool)`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *KnnSearchRequest) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.

### GetKnn

`func (o *KnnSearchRequest) GetKnn() KnnSearchRequestAllOfKnn`

GetKnn returns the Knn field if non-nil, zero value otherwise.

### GetKnnOk

`func (o *KnnSearchRequest) GetKnnOk() (*KnnSearchRequestAllOfKnn, bool)`

GetKnnOk returns a tuple with the Knn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnn

`func (o *KnnSearchRequest) SetKnn(v KnnSearchRequestAllOfKnn)`

SetKnn sets Knn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


