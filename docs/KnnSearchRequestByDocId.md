# KnnSearchRequestByDocId

Request object for knn search operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | [default to ""]
**Field** | **string** |  | [default to ""]
**DocId** | **int32** |  | 
**K** | **int32** |  | 
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

### NewKnnSearchRequestByDocId

`func NewKnnSearchRequestByDocId(index string, field string, docId int32, k int32, ) *KnnSearchRequestByDocId`

NewKnnSearchRequestByDocId instantiates a new KnnSearchRequestByDocId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnSearchRequestByDocIdWithDefaults

`func NewKnnSearchRequestByDocIdWithDefaults() *KnnSearchRequestByDocId`

NewKnnSearchRequestByDocIdWithDefaults instantiates a new KnnSearchRequestByDocId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *KnnSearchRequestByDocId) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *KnnSearchRequestByDocId) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *KnnSearchRequestByDocId) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetField

`func (o *KnnSearchRequestByDocId) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnSearchRequestByDocId) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnSearchRequestByDocId) SetField(v string)`

SetField sets Field field to given value.


### GetDocId

`func (o *KnnSearchRequestByDocId) GetDocId() int32`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *KnnSearchRequestByDocId) GetDocIdOk() (*int32, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *KnnSearchRequestByDocId) SetDocId(v int32)`

SetDocId sets DocId field to given value.


### GetK

`func (o *KnnSearchRequestByDocId) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnSearchRequestByDocId) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnSearchRequestByDocId) SetK(v int32)`

SetK sets K field to given value.


### GetFulltextFilter

`func (o *KnnSearchRequestByDocId) GetFulltextFilter() map[string]interface{}`

GetFulltextFilter returns the FulltextFilter field if non-nil, zero value otherwise.

### GetFulltextFilterOk

`func (o *KnnSearchRequestByDocId) GetFulltextFilterOk() (*map[string]interface{}, bool)`

GetFulltextFilterOk returns a tuple with the FulltextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextFilter

`func (o *KnnSearchRequestByDocId) SetFulltextFilter(v map[string]interface{})`

SetFulltextFilter sets FulltextFilter field to given value.

### HasFulltextFilter

`func (o *KnnSearchRequestByDocId) HasFulltextFilter() bool`

HasFulltextFilter returns a boolean if a field has been set.

### GetAttrFilter

`func (o *KnnSearchRequestByDocId) GetAttrFilter() map[string]interface{}`

GetAttrFilter returns the AttrFilter field if non-nil, zero value otherwise.

### GetAttrFilterOk

`func (o *KnnSearchRequestByDocId) GetAttrFilterOk() (*map[string]interface{}, bool)`

GetAttrFilterOk returns a tuple with the AttrFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrFilter

`func (o *KnnSearchRequestByDocId) SetAttrFilter(v map[string]interface{})`

SetAttrFilter sets AttrFilter field to given value.

### HasAttrFilter

`func (o *KnnSearchRequestByDocId) HasAttrFilter() bool`

HasAttrFilter returns a boolean if a field has been set.

### GetLimit

`func (o *KnnSearchRequestByDocId) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KnnSearchRequestByDocId) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KnnSearchRequestByDocId) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KnnSearchRequestByDocId) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *KnnSearchRequestByDocId) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KnnSearchRequestByDocId) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KnnSearchRequestByDocId) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KnnSearchRequestByDocId) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMaxMatches

`func (o *KnnSearchRequestByDocId) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *KnnSearchRequestByDocId) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *KnnSearchRequestByDocId) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *KnnSearchRequestByDocId) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetSort

`func (o *KnnSearchRequestByDocId) GetSort() []map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *KnnSearchRequestByDocId) GetSortOk() (*[]map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *KnnSearchRequestByDocId) SetSort(v []map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *KnnSearchRequestByDocId) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetAggs

`func (o *KnnSearchRequestByDocId) GetAggs() map[string]Aggregation`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *KnnSearchRequestByDocId) GetAggsOk() (*map[string]Aggregation, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *KnnSearchRequestByDocId) SetAggs(v map[string]Aggregation)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *KnnSearchRequestByDocId) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExpressions

`func (o *KnnSearchRequestByDocId) GetExpressions() map[string]string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *KnnSearchRequestByDocId) GetExpressionsOk() (*map[string]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *KnnSearchRequestByDocId) SetExpressions(v map[string]string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *KnnSearchRequestByDocId) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetHighlight

`func (o *KnnSearchRequestByDocId) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *KnnSearchRequestByDocId) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *KnnSearchRequestByDocId) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *KnnSearchRequestByDocId) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetSource

`func (o *KnnSearchRequestByDocId) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KnnSearchRequestByDocId) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KnnSearchRequestByDocId) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *KnnSearchRequestByDocId) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOptions

`func (o *KnnSearchRequestByDocId) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KnnSearchRequestByDocId) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KnnSearchRequestByDocId) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KnnSearchRequestByDocId) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProfile

`func (o *KnnSearchRequestByDocId) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *KnnSearchRequestByDocId) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *KnnSearchRequestByDocId) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *KnnSearchRequestByDocId) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTrackScores

`func (o *KnnSearchRequestByDocId) GetTrackScores() bool`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *KnnSearchRequestByDocId) GetTrackScoresOk() (*bool, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *KnnSearchRequestByDocId) SetTrackScores(v bool)`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *KnnSearchRequestByDocId) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


