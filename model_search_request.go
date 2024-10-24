/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 5.0.0
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	_"bytes"
	_"fmt"
)

// checks if the SearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequest{}

// SearchRequest Request object for search operation
type SearchRequest struct {
	// The index to perform the search on
	Index string
	Query *SearchQuery
	// Join clause to combine search data from multiple tables
	Join []Join
	Highlight *Highlight
	// Maximum number of results to return
	Limit *int32
	Knn *KnnQuery
	// Defines aggregation settings for grouping results
	Aggs map[string]Aggregation
	// Expressions to calculate additional values for the result
	Expressions map[string]string
	// Maximum number of matches allowed in the result
	MaxMatches *int32
	// Starting point for pagination of the result
	Offset *int32
	// Additional search options
	Options map[string]interface{}
	// Enable or disable profiling of the search request
	Profile *bool
	Sort interface{}
	Source interface{}
	// Enable or disable result weight calculation used for sorting
	TrackScores *bool
}

type _SearchRequest SearchRequest

// NewSearchRequest instantiates a new SearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequest(index string) *SearchRequest {
	this := SearchRequest{}
	this.Index = index
	return &this
}

// NewSearchRequestWithDefaults instantiates a new SearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestWithDefaults() *SearchRequest {
	this := SearchRequest{}
	return &this
}

// GetIndex returns the Index field value
func (o *SearchRequest) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SearchRequest) SetIndex(v string) {
	o.Index = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchRequest) GetQuery() SearchQuery {
	if o == nil || IsNil(o.Query) {
		var ret SearchQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetQueryOk() (*SearchQuery, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given SearchQuery and assigns it to the Query field.
func (o *SearchRequest) SetQuery(v SearchQuery) {
	o.Query = &v
}

// GetJoin returns the Join field value if set, zero value otherwise.
func (o *SearchRequest) GetJoin() []Join {
	if o == nil || IsNil(o.Join) {
		var ret []Join
		return ret
	}
	return o.Join
}

// GetJoinOk returns a tuple with the Join field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetJoinOk() ([]Join, bool) {
	if o == nil || IsNil(o.Join) {
		return nil, false
	}
	return o.Join, true
}

// HasJoin returns a boolean if a field has been set.
func (o *SearchRequest) HasJoin() bool {
	if o != nil && !IsNil(o.Join) {
		return true
	}

	return false
}

// SetJoin gets a reference to the given []Join and assigns it to the Join field.
func (o *SearchRequest) SetJoin(v []Join) {
	o.Join = v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *SearchRequest) GetHighlight() Highlight {
	if o == nil || IsNil(o.Highlight) {
		var ret Highlight
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetHighlightOk() (*Highlight, bool) {
	if o == nil || IsNil(o.Highlight) {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *SearchRequest) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given Highlight and assigns it to the Highlight field.
func (o *SearchRequest) SetHighlight(v Highlight) {
	o.Highlight = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SearchRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SearchRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SearchRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetKnn returns the Knn field value if set, zero value otherwise.
func (o *SearchRequest) GetKnn() KnnQuery {
	if o == nil || IsNil(o.Knn) {
		var ret KnnQuery
		return ret
	}
	return *o.Knn
}

// GetKnnOk returns a tuple with the Knn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetKnnOk() (*KnnQuery, bool) {
	if o == nil || IsNil(o.Knn) {
		return nil, false
	}
	return o.Knn, true
}

// HasKnn returns a boolean if a field has been set.
func (o *SearchRequest) HasKnn() bool {
	if o != nil && !IsNil(o.Knn) {
		return true
	}

	return false
}

// SetKnn gets a reference to the given KnnQuery and assigns it to the Knn field.
func (o *SearchRequest) SetKnn(v KnnQuery) {
	o.Knn = &v
}

// GetAggs returns the Aggs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetAggs() map[string]Aggregation {
	if o == nil {
		var ret map[string]Aggregation
		return ret
	}
	return o.Aggs
}

// GetAggsOk returns a tuple with the Aggs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetAggsOk() (*map[string]Aggregation, bool) {
	if o == nil || IsNil(o.Aggs) {
		return nil, false
	}
	return &o.Aggs, true
}

// HasAggs returns a boolean if a field has been set.
func (o *SearchRequest) HasAggs() bool {
	if o != nil && !IsNil(o.Aggs) {
		return true
	}

	return false
}

// SetAggs gets a reference to the given map[string]Aggregation and assigns it to the Aggs field.
func (o *SearchRequest) SetAggs(v map[string]Aggregation) {
	o.Aggs = v
}

// GetExpressions returns the Expressions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetExpressions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetExpressionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Expressions) {
		return nil, false
	}
	return &o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SearchRequest) HasExpressions() bool {
	if o != nil && !IsNil(o.Expressions) {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given map[string]string and assigns it to the Expressions field.
func (o *SearchRequest) SetExpressions(v map[string]string) {
	o.Expressions = v
}

// GetMaxMatches returns the MaxMatches field value if set, zero value otherwise.
func (o *SearchRequest) GetMaxMatches() int32 {
	if o == nil || IsNil(o.MaxMatches) {
		var ret int32
		return ret
	}
	return *o.MaxMatches
}

// GetMaxMatchesOk returns a tuple with the MaxMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetMaxMatchesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMatches) {
		return nil, false
	}
	return o.MaxMatches, true
}

// HasMaxMatches returns a boolean if a field has been set.
func (o *SearchRequest) HasMaxMatches() bool {
	if o != nil && !IsNil(o.MaxMatches) {
		return true
	}

	return false
}

// SetMaxMatches gets a reference to the given int32 and assigns it to the MaxMatches field.
func (o *SearchRequest) SetMaxMatches(v int32) {
	o.MaxMatches = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SearchRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SearchRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SearchRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SearchRequest) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SearchRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *SearchRequest) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SearchRequest) GetProfile() bool {
	if o == nil || IsNil(o.Profile) {
		var ret bool
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SearchRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given bool and assigns it to the Profile field.
func (o *SearchRequest) SetProfile(v bool) {
	o.Profile = &v
}

// GetSort returns the Sort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetSort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetSortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SearchRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given interface{} and assigns it to the Sort field.
func (o *SearchRequest) SetSort(v interface{}) {
	o.Sort = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchRequest) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchRequest) GetSourceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SearchRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *SearchRequest) SetSource(v interface{}) {
	o.Source = v
}

// GetTrackScores returns the TrackScores field value if set, zero value otherwise.
func (o *SearchRequest) GetTrackScores() bool {
	if o == nil || IsNil(o.TrackScores) {
		var ret bool
		return ret
	}
	return *o.TrackScores
}

// GetTrackScoresOk returns a tuple with the TrackScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetTrackScoresOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackScores) {
		return nil, false
	}
	return o.TrackScores, true
}

// HasTrackScores returns a boolean if a field has been set.
func (o *SearchRequest) HasTrackScores() bool {
	if o != nil && !IsNil(o.TrackScores) {
		return true
	}

	return false
}

// SetTrackScores gets a reference to the given bool and assigns it to the TrackScores field.
func (o *SearchRequest) SetTrackScores(v bool) {
	o.TrackScores = &v
}

func (o SearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Join) {
		toSerialize["join"] = o.Join
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Knn) {
		toSerialize["knn"] = o.Knn
	}
	if o.Aggs != nil {
		toSerialize["aggs"] = o.Aggs
	}
	if o.Expressions != nil {
		toSerialize["expressions"] = o.Expressions
	}
	if !IsNil(o.MaxMatches) {
		toSerialize["max_matches"] = o.MaxMatches
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Source != nil {
		toSerialize["_source"] = o.Source
	}
	if !IsNil(o.TrackScores) {
		toSerialize["track_scores"] = o.TrackScores
	}
	return toSerialize, nil
}

type NullableSearchRequest struct {
	value *SearchRequest
	isSet bool
}

func (v NullableSearchRequest) Get() *SearchRequest {
	return v.value
}

func (v *NullableSearchRequest) Set(val *SearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequest(val *SearchRequest) *NullableSearchRequest {
	return &NullableSearchRequest{value: val, isSet: true}
}

func (v NullableSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


