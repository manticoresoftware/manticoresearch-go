# KnnSearchRequestByVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** |  | [default to ""]
**Field** | **string** |  | [default to ""]
**QueryVector** | **[]float32** |  | 
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

### NewKnnSearchRequestByVector

`func NewKnnSearchRequestByVector(index string, field string, queryVector []float32, k int32, ) *KnnSearchRequestByVector`

NewKnnSearchRequestByVector instantiates a new KnnSearchRequestByVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnSearchRequestByVectorWithDefaults

`func NewKnnSearchRequestByVectorWithDefaults() *KnnSearchRequestByVector`

NewKnnSearchRequestByVectorWithDefaults instantiates a new KnnSearchRequestByVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *KnnSearchRequestByVector) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *KnnSearchRequestByVector) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *KnnSearchRequestByVector) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetField

`func (o *KnnSearchRequestByVector) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnSearchRequestByVector) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnSearchRequestByVector) SetField(v string)`

SetField sets Field field to given value.


### GetQueryVector

`func (o *KnnSearchRequestByVector) GetQueryVector() []float32`

GetQueryVector returns the QueryVector field if non-nil, zero value otherwise.

### GetQueryVectorOk

`func (o *KnnSearchRequestByVector) GetQueryVectorOk() (*[]float32, bool)`

GetQueryVectorOk returns a tuple with the QueryVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVector

`func (o *KnnSearchRequestByVector) SetQueryVector(v []float32)`

SetQueryVector sets QueryVector field to given value.


### GetK

`func (o *KnnSearchRequestByVector) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnSearchRequestByVector) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnSearchRequestByVector) SetK(v int32)`

SetK sets K field to given value.


### GetFulltextFilter

`func (o *KnnSearchRequestByVector) GetFulltextFilter() map[string]interface{}`

GetFulltextFilter returns the FulltextFilter field if non-nil, zero value otherwise.

### GetFulltextFilterOk

`func (o *KnnSearchRequestByVector) GetFulltextFilterOk() (*map[string]interface{}, bool)`

GetFulltextFilterOk returns a tuple with the FulltextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltextFilter

`func (o *KnnSearchRequestByVector) SetFulltextFilter(v map[string]interface{})`

SetFulltextFilter sets FulltextFilter field to given value.

### HasFulltextFilter

`func (o *KnnSearchRequestByVector) HasFulltextFilter() bool`

HasFulltextFilter returns a boolean if a field has been set.

### GetAttrFilter

`func (o *KnnSearchRequestByVector) GetAttrFilter() map[string]interface{}`

GetAttrFilter returns the AttrFilter field if non-nil, zero value otherwise.

### GetAttrFilterOk

`func (o *KnnSearchRequestByVector) GetAttrFilterOk() (*map[string]interface{}, bool)`

GetAttrFilterOk returns a tuple with the AttrFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrFilter

`func (o *KnnSearchRequestByVector) SetAttrFilter(v map[string]interface{})`

SetAttrFilter sets AttrFilter field to given value.

### HasAttrFilter

`func (o *KnnSearchRequestByVector) HasAttrFilter() bool`

HasAttrFilter returns a boolean if a field has been set.

### GetLimit

`func (o *KnnSearchRequestByVector) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KnnSearchRequestByVector) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KnnSearchRequestByVector) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KnnSearchRequestByVector) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *KnnSearchRequestByVector) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *KnnSearchRequestByVector) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *KnnSearchRequestByVector) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *KnnSearchRequestByVector) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMaxMatches

`func (o *KnnSearchRequestByVector) GetMaxMatches() int32`

GetMaxMatches returns the MaxMatches field if non-nil, zero value otherwise.

### GetMaxMatchesOk

`func (o *KnnSearchRequestByVector) GetMaxMatchesOk() (*int32, bool)`

GetMaxMatchesOk returns a tuple with the MaxMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMatches

`func (o *KnnSearchRequestByVector) SetMaxMatches(v int32)`

SetMaxMatches sets MaxMatches field to given value.

### HasMaxMatches

`func (o *KnnSearchRequestByVector) HasMaxMatches() bool`

HasMaxMatches returns a boolean if a field has been set.

### GetSort

`func (o *KnnSearchRequestByVector) GetSort() []map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *KnnSearchRequestByVector) GetSortOk() (*[]map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *KnnSearchRequestByVector) SetSort(v []map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *KnnSearchRequestByVector) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetAggs

`func (o *KnnSearchRequestByVector) GetAggs() map[string]Aggregation`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *KnnSearchRequestByVector) GetAggsOk() (*map[string]Aggregation, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *KnnSearchRequestByVector) SetAggs(v map[string]Aggregation)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *KnnSearchRequestByVector) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetExpressions

`func (o *KnnSearchRequestByVector) GetExpressions() map[string]string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *KnnSearchRequestByVector) GetExpressionsOk() (*map[string]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *KnnSearchRequestByVector) SetExpressions(v map[string]string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *KnnSearchRequestByVector) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetHighlight

`func (o *KnnSearchRequestByVector) GetHighlight() Highlight`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *KnnSearchRequestByVector) GetHighlightOk() (*Highlight, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *KnnSearchRequestByVector) SetHighlight(v Highlight)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *KnnSearchRequestByVector) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetSource

`func (o *KnnSearchRequestByVector) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KnnSearchRequestByVector) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KnnSearchRequestByVector) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *KnnSearchRequestByVector) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOptions

`func (o *KnnSearchRequestByVector) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *KnnSearchRequestByVector) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *KnnSearchRequestByVector) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *KnnSearchRequestByVector) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProfile

`func (o *KnnSearchRequestByVector) GetProfile() bool`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *KnnSearchRequestByVector) GetProfileOk() (*bool, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *KnnSearchRequestByVector) SetProfile(v bool)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *KnnSearchRequestByVector) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTrackScores

`func (o *KnnSearchRequestByVector) GetTrackScores() bool`

GetTrackScores returns the TrackScores field if non-nil, zero value otherwise.

### GetTrackScoresOk

`func (o *KnnSearchRequestByVector) GetTrackScoresOk() (*bool, bool)`

GetTrackScoresOk returns a tuple with the TrackScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackScores

`func (o *KnnSearchRequestByVector) SetTrackScores(v bool)`

SetTrackScores sets TrackScores field to given value.

### HasTrackScores

`func (o *KnnSearchRequestByVector) HasTrackScores() bool`

HasTrackScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


